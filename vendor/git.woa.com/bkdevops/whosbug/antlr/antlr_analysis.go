package antlr

import (
	cpp "git.woa.com/bkdevops/whosbug/antlr/cppLib"
	golang "git.woa.com/bkdevops/whosbug/antlr/goLib"
	java "git.woa.com/bkdevops/whosbug/antlr/javaLib"
	js "git.woa.com/bkdevops/whosbug/antlr/jsLib"
	kt "git.woa.com/bkdevops/whosbug/antlr/kotlinLib"
	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/crypto"
	"git.woa.com/bkdevops/whosbug/util"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// AnalyzeCommitDiff
//	@Description: 使用antlr分析commitDiff信息
//	@param commitDiff diff信息(path)
//	@author KevinMatt 2021-08-03 21:41:08
//	@function_mark PASS
func AnalyzeCommitDiff(commitDiff config.DiffParsedType) {
	//	获取antlr分析结果
	antlrAnalyzeRes := antlrAnalysis(commitDiff.DiffText, commitDiff.TargetLanguage)

	var prevObject config.ObjectInfoType
	var countMinus int = 0
	var countPlus int = 0
	for _, changeLineNumber := range commitDiff.ChangeLineNumbers {
		currentObject := addObjectFromChangeLineNumber(commitDiff, changeLineNumber, antlrAnalyzeRes)
		if currentObject.Equals(config.ObjectInfoType{}) {
			continue
		}
		if prevObject.Equals(config.ObjectInfoType{}) {
			prevObject = currentObject
			continue
		}
		if currentObject.Equals(prevObject) {
			if changeLineNumber.ChangeType == "-" {
				countMinus++
			} else if changeLineNumber.ChangeType == "+" {
				countPlus++
			}
		} else {
			prevObject.OldLineCount = countMinus + prevObject.CurrentLineCount - countPlus
			if prevObject.OldLineCount < 0 {
				prevObject.OldLineCount = 0
			}
			prevObject.RemovedLineCount = countMinus
			prevObject.AddedNewLineCount = countPlus
			if prevObject.CurrentLineCount < countPlus {
				util.ForDebug()
			}
			if prevObject.AddedNewLineCount != 0 || prevObject.RemovedLineCount != 0 {
				// 送入channel
				config.ObjectChan <- prevObject
			}
			countMinus = 0
			countPlus = 0
			prevObject = currentObject
		}
		// 用于比较两次的结构体是否重复(匹配行范围导致的重复结果)
	}
}

// antlrAnalysis
//	@Description: antlr分析过程
//	@param targetFilePath 分析的目标文件
//	@param langMode 分析的语言模式
//	@return astResType 返回分析信息结构体
//	@author KevinMatt 2021-07-29 19:49:37
//	@function_mark  PASS
func antlrAnalysis(diffText string, langMode string) (result astResType) {
	switch langMode {
	case "java":
		result = ExecuteJava(diffText)
	case "python":
		//	TODO: Python Unsupported
		// result = ExecutePython(diffText)
	case "kotlin":
		result = ExecuteKotlin(diffText)
	case "golang":
		result = ExecuteGolang(diffText)
	case "javascript":
		// ! 暂时下线js解析
		// result = ExecuteJavaScript(diffText)
	case "cpp":
		// TODO Really Slow
		result = ExecuteCpp(diffText)
	default:
		break
	}
	return
}

func ExecuteGolang(diffText string) astResType {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := goLexerPool.Get().(*golang.GoLexer)
	defer goLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	lexer.RemoveErrorListeners()
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
	p.RemoveErrorListeners()
	//	解析模式->每个编译单位
	tree := p.SourceFile()
	//	创建listener
	listener := newGoTreeShapeListenerPool.Get().(*GoTreeShapeListener)
	defer newGoTreeShapeListenerPool.Put(listener)
	// 初始化置空
	listener.AstInfoList = astResType{}
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.AstInfoList
}

func ExecuteCpp(diffText string) astResType {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := cppLexerPool.Get().(*cpp.CPP14Lexer)
	defer cppLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := cppParserPool.Get().(*cpp.CPP14Parser)
	p.RemoveErrorListeners()
	defer cppParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	// //	启用SLL两阶段加速解析模式
	// p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.TranslationUnit()

	//	创建listener
	listener := newCppTreeShapeListenerPool.Get().(*CppTreeShapeListener)
	defer newCppTreeShapeListenerPool.Put(listener)
	// 初始化置空
	listener.AstInfoList = astResType{}
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.AstInfoList
}

func ExecuteJavaScript(diffText string) astResType {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := javascriptLexerPool.Get().(*js.JavaScriptLexer)
	defer javaLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := javascriptParserPool.Get().(*js.JavaScriptParser)
	p.RemoveErrorListeners()
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
	return listener.AstInfoList
}

func ExecutePython(diffText string) AnalysisInfoType {
	util.ForDebug(diffText)
	return AnalysisInfoType{}
}

func ExecuteJava(diffText string) astResType {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := javaLexerPool.Get().(*java.JavaLexer)
	defer javaLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := javaParserPool.Get().(*java.JavaParser)
	p.RemoveErrorListeners()
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
	// 初始化置空
	listener.AstInfoList = *new(astResType)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.AstInfoList
}

// ExecuteKotlin
//	@Description: 执行java分析
//	@param targetFilePath 分析目标路径
//	@return javaparser.AnalysisInfoType 返回分析结果结构体
//	@author KevinMatt 2021-07-29 19:51:16
//	@function_mark PASS
func ExecuteKotlin(diffText string) astResType {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := kotlinLexerPool.Get().(*kt.KotlinLexer)
	defer kotlinLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := kotlinParserPool.Get().(*kt.KotlinParser)
	p.RemoveErrorListeners()
	defer kotlinParserPool.Put(p)
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
	// 初始化置空
	listener.AstInfoList = astResType{}
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.AstInfoList
}

// addObjectFromChangeLineNumber
//	@Description: 传入的参数较多，大致功能是构建object的map
//	@param commitDiff
//	@param changeLineNumber 行号变动
//	@param antlrAnalyzeRes antlr分析结果
//	@return objectInfoType
//	@author KevinMatt 2021-08-03 19:26:12
//	@function_mark PASS
func addObjectFromChangeLineNumber(commitDiff config.DiffParsedType, changeLineNumber config.ChangeLineType, antlrAnalyzeRes astResType) (newObject config.ObjectInfoType) {
	//	寻找变动方法
	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
	if changeMethod.MethodName == "" {
		return
	}

	//	TODO Ready for newMethod
	newObject = config.ObjectInfoType{
		CommitHash:       commitDiff.CommitHash, //crypto.Base64Encrypt(commitDiff.CommitHash)
		ID:               crypto.Base64Encrypt(changeMethod.MethodName),
		OldID:            "",
		FilePath:         crypto.Base64Encrypt(commitDiff.DiffFileName),
		Parameters:       crypto.Base64Encrypt(changeMethod.Parameters),
		OldLineCount:     0,
		CurrentLineCount: changeMethod.EndLine - changeMethod.StartLine + 1,
		StartLine:        changeMethod.StartLine,
		EndLine:          changeMethod.EndLine,
	}
	return
}

// findChangedMethod
//	@Description: 寻找变动了的方法
//	@param changeLineNumber 变动行
//	@param antlrAnalyzeRes antlr分析结果
//	@return changeMethodInfo 变动方法信息
//	@author KevinMatt 2021-07-29 19:38:19
//	@function_mark PASS
func findChangedMethod(changeLineNumber config.ChangeLineType, antlrAnalyzeRes astResType) (changeMethodInfo MethodInfoType) {
	var lineRangeList []LineRangeType
	//	遍历匹配到的方法列表，存储其首行
	for index := range antlrAnalyzeRes.Methods {
		lineRangeList = append(lineRangeList, LineRangeType{
			StartLine: antlrAnalyzeRes.Methods[index].StartLine,
			EndLine:   antlrAnalyzeRes.Methods[index].EndLine,
		})
	}
	//	寻找方法行所在的范围位置
	resIndex := FindIntervalIndex(lineRangeList, changeLineNumber.LineNumber)
	//	判断是否有位置插入
	if resIndex > -1 {
		changeMethodInfo = antlrAnalyzeRes.Methods[resIndex]
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
func FindIntervalIndex(lineRangeList []LineRangeType, target float64) int {
	if len(lineRangeList) == 0 {
		return -1
	}

	for index, item := range lineRangeList {
		if target < float64(item.EndLine+1) && target > float64(item.StartLine-1) {
			return index
		}
	}
	return -1
}

// // FindIntervalIndex
// //	@Description: 寻找可插入位置
// //	@param nums 传入的行号切片
// //	@param target 要插入的目标行号
// //	@return int 返回插入位置，-1代表无法插入
// //	@author KevinMatt 2021-07-29 19:42:18
// //	@function_mark PASS
// func FindIntervalIndex_old(nums []int, target float64) int {
// 	var newTarget int = int(target)
// 	if float64(int(target)) < target {
// 		newTarget = int(target) + 1
// 	}

// 	if len(nums) == 0 {
// 		return -1
// 	}
// 	if newTarget < nums[0] || newTarget > nums[len(nums)-1] {
// 		return -1
// 	}
// 	for index := range nums {
// 		if newTarget < nums[index] {
// 			return index - 1
// 		} else if newTarget == nums[index] {
// 			return index
// 		}
// 	}
// 	return -1
// }
