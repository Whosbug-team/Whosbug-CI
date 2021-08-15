package antlrpack

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	cpp "whosbugPack/antlrpack/cpp_lib"
	golang "whosbugPack/antlrpack/go_lib"
	javaparser "whosbugPack/antlrpack/java_lib"
	javascript "whosbugPack/antlrpack/js_lib"
	kotlin "whosbugPack/antlrpack/kotlin_lib"
	"whosbugPack/global_type"
	"whosbugPack/utility"
)

// AnalyzeCommitDiff
//	@Description: 使用antlr分析commitDiff信息
//	@param commitDiff diff信息(path)
//	@author KevinMatt 2021-08-03 21:41:08
//	@function_mark PASS
func AnalyzeCommitDiff(commitDiff global_type.DiffParsedType) {
	//	获取antlr分析结果
	antlrAnalyzeRes := antlrAnalysis(commitDiff.DiffText, "java")
	//var tempCompare global_type.ObjectInfoType
	var tempCompare global_type.ObjectInfoType
	var countMinus int = 0
	var countPlus int = 0
	for index := range commitDiff.ChangeLineNumbers {
		temp := addObjectFromChangeLineNumber(commitDiff, commitDiff.ChangeLineNumbers[index], antlrAnalyzeRes)
		//newTemp := addObjectFromChangeLineNumber(commitDiff, commitDiff.ChangeLineNumbers[index], antlrAnalyzeRes)
		if temp.Equals(tempCompare) {
			if commitDiff.ChangeLineNumbers[index].ChangeType == "-" {
				countMinus++
			} else {
				countPlus++
			}
			continue
		}
		if !temp.Equals(global_type.ObjectInfoType{}) {
			temp.ChangedOldLineCount = countMinus
			countMinus = 0
			temp.ChangedNewLineCount = countPlus
			countPlus = 0
			temp.NewLineCount = temp.OldLineCount + temp.ChangedNewLineCount - temp.ChangedOldLineCount
			//	送入channel
			global_type.ObjectChan <- temp
		}
		//	用于比较两次的结构体是否重复(匹配行范围导致的重复结果)
		tempCompare = temp
	}
}

// antlrAnalysis
//	@Description: antlr分析过程
//	@param targetFilePath 分析的目标文件
//	@param langMode 分析的语言模式
//	@return javaparser.AnalysisInfoType 返回分析信息结构体(跨包)
//	@author KevinMatt 2021-07-29 19:49:37
//	@function_mark  PASS
func antlrAnalysis(diffText string, langMode string) AnalysisInfoType {
	var result AnalysisInfoType
	switch langMode {
	case "java":
		result = ExecuteJava(diffText)
	//	TODO 其他语言的适配支持
	case "python":
		result = ExecutePython(diffText)
	case "kotlin":
		result = ExecuteKotlin(diffText)
	case "golang":
		result = ExecuteGolang(diffText)
	case "javascript":
		result = ExecuteJavaScript(diffText)
	default:
		break
	}
	return result
}
func ExecuteGolang(diffText string) AnalysisInfoType{
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := goLexerPool.Get().(*golang.GoLexer)
	defer goLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := goParserPool.Get().(*golang.GoParser)
	defer goParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.SourceFile()
	//	创建listener
	listener := newGoTreeShapeListenerPool.Get().(*GoTreeShapeListener)
	defer newGoTreeShapeListenerPool.Put(listener)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Infos
}
func ExecuteCpp(diffText string) AnalysisInfoType{
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := cppLexerPool.Get().(*cpp.CPP14Lexer)
	defer cppLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := cppParserPool.Get().(*cpp.CPP14Parser)
	defer cppParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.TranslationUnit()
	//	创建listener
	listener := newCppTreeShapeListenerPool.Get().(*CppTreeShapeListener)
	defer newCppTreeShapeListenerPool.Put(listener)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Infos
}

func ExecuteJavaScript(diffText string) AnalysisInfoType{
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := javascriptLexerPool.Get().(*javascript.JavaScriptLexer)
	defer javaLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := javascriptParserPool.Get().(*javascript.JavaScriptParser)
	defer javascriptParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.Program()
	//	创建listener
	listener := newJavaScriptTreeShapeListenerPool.Get().(*JSTreeShapeListener)
	defer newJavaScriptTreeShapeListenerPool.Put(listener)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Infos
}

func ExecutePython(diffText string) AnalysisInfoType {
	utility.ForDebug(diffText)
	return AnalysisInfoType{}
}

func ExecuteJava(diffText string) AnalysisInfoType {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := javaLexerPool.Get().(*javaparser.JavaLexer)
	defer javaLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := javaParserPool.Get().(*javaparser.JavaParser)
	defer javaParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.CompilationUnit()
	//	创建listener
	listener := newJavaTreeShapeListenerPool.Get().(*JavaTreeShapeListener)
	defer newJavaTreeShapeListenerPool.Put(listener)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Infos
}



// ExecuteJava
//	@Description: 执行java分析
//	@param targetFilePath 分析目标路径
//	@return javaparser.AnalysisInfoType 返回分析结果结构体
//	@author KevinMatt 2021-07-29 19:51:16
//	@function_mark PASS
func ExecuteKotlin(diffText string) AnalysisInfoType {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := kotlinLexerPool.Get().(*kotlin.KotlinLexer)
	defer kotlinLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := kotlinParserPool.Get().(*kotlin.KotlinParser)
	defer kotlinParserPool.Put(p)
	p.RemoveErrorListeners()
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.KotlinFile()
	//	创建listener
	listener := newKotlinTreeShapeListenerPool.Get().(*KotlinTreeShapeListener)
	defer newKotlinTreeShapeListenerPool.Put(listener)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Infos
}

// addObjectFromChangeLineNumber
//	@Description: 传入的参数较多，大致功能是构建object的map
//	@param commitDiff
//	@param changeLineNumber 行号变动
//	@param antlrAnalyzeRes antlr分析结果
//	@return objectInfoType
//	@author KevinMatt 2021-08-03 19:26:12
//	@function_mark PASS
func addObjectFromChangeLineNumber(commitDiff global_type.DiffParsedType, changeLineNumber global_type.ChangeLineType, antlrAnalyzeRes AnalysisInfoType) global_type.ObjectInfoType {
	//	寻找变动方法
	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
	if changeMethod.MethodName == "" {
		//	为空直接跳过执行
		return global_type.ObjectInfoType{}
	}

	//	TODO Ready for newMethod
	var newObject global_type.ObjectInfoType
	newObject = global_type.ObjectInfoType{
		CommitHash:   commitDiff.CommitHash, //utility.Base64Encrypt(commitDiff.CommitHash)
		Id:           changeMethod.MasterObject.ObjectName + "." + changeMethod.MethodName,
		OldId:        "",
		FilePath:     utility.Base64Encrypt(commitDiff.CommitHash),
		OldLineCount: changeMethod.EndLine - changeMethod.StartLine,
		NewLineCount: commitDiff.NewLineCount,
		Calling:      changeMethod.CallMethods,
	}
	return newObject
}

// findChangedMethod
//	@Description: 寻找变动了的方法
//	@param changeLineNumber 变动行
//	@param antlrAnalyzeRes antlr分析结果
//	@return changeMethodInfo 变动方法信息
//	@author KevinMatt 2021-07-29 19:38:19
//	@function_mark PASS
func findChangedMethod(changeLineNumber global_type.ChangeLineType, antlrAnalyzeRes AnalysisInfoType) (changeMethodInfo MethodInfoType) {
	var startLineNumbers []int
	//	遍历匹配到的方法列表，存储其首行
	for index := range antlrAnalyzeRes.AstInfoList.Methods {
		startLineNumbers = append(startLineNumbers, antlrAnalyzeRes.AstInfoList.Methods[index].StartLine)
	}
	//	寻找方法行所在的范围位置
	resIndex := FindIntervalIndex(startLineNumbers, changeLineNumber.LineNumber)
	//	判断是否有位置插入
	if resIndex > -1 {
		changeMethodInfo = antlrAnalyzeRes.AstInfoList.Methods[resIndex]
	}
	return
}

// FindIntervalIndex
//	@Description: 寻找可插入位置
//	@param nums 传入的行号切片
//	@param target 要插入的目标行号
//	@return int 返回插入位置，-1代表无法插入
//	@author KevinMatt 2021-07-29 19:42:18
//	@function_mark PASS
func FindIntervalIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}
	for index := range nums {
		if target < nums[index] {
			return index - 1
		} else if target == nums[index] {
			return index
		}
	}
	return -1
}
