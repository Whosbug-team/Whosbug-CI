package antlrpack

import (
	javaparser "anrlr4_ast/java"
	"encoding/base64"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"whosbugPack/global_type"
	"whosbugPack/utility"
)

/* AnalyzeCommitDiff
/* @Description: 使用antlr分析commitDiff信息
 * @param commitDiff diff信息(path)
 * @author KevinMatt 2021-08-03 21:41:08
 * @function_mark PASS
*/
func AnalyzeCommitDiff(commitDiff global_type.DiffParsedType) {

	// 获取antlr分析结果
	antlrAnalyzeRes := antlrAnalysis(commitDiff.DiffText, "java")
	var tempCompare global_type.ObjectInfoType
	for index, _ := range commitDiff.ChangeLineNumbers {
		temp := addObjectFromChangeLineNumber(commitDiff, commitDiff.ChangeLineNumbers[index], antlrAnalyzeRes)
		if temp == tempCompare {
			continue
		}
		if temp != (global_type.ObjectInfoType{}) {
			// 送入channel
			global_type.ObjectChan <- temp
		}
		// 用于比较两次的结构体是否重复(匹配行范围导致的重复结果)
		tempCompare = temp
	}
}

/* antlrAnalysis
/* @Description: antlr分析过程
 * @param targetFilePath 分析的目标文件
 * @param langMode 分析的语言模式
 * @return javaparser.AnalysisInfoType 返回分析信息结构体(跨包)
 * @author KevinMatt 2021-07-29 19:49:37
 * @function_mark  PASS
*/
func antlrAnalysis(diffText string, langMode string) AnalysisInfoType {
	var result AnalysisInfoType
	switch langMode {
	case "java":
		result = ExecuteJava(diffText)
	// TODO 其他语言的适配支持
	default:
		break
	}
	return result
}

/* ExecuteJava
/* @Description: 执行java分析
 * @param targetFilePath 分析目标路径
 * @return javaparser.AnalysisInfoType 返回分析结果结构体
 * @author KevinMatt 2021-07-29 19:51:16
 * @function_mark PASS
*/
func ExecuteJava(diffText string) AnalysisInfoType {
	// 截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	// 初始化lexer
	lexer := lexerPool.Get().(*javaparser.JavaLexer)
	defer lexerPool.Put(lexer)
	lexer.SetInputStream(input)
	// 初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	// 初始化Parser
	p := parserPool.Get().(*javaparser.JavaParser)
	defer parserPool.Put(p)
	p.SetTokenStream(stream)
	// 构建语法解析树
	p.BuildParseTrees = true
	// 启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	// 解析模式->每个编译单位
	tree := p.CompilationUnit()
	// 创建listener
	listener := newTreeShapeListenerPool.Get().(*TreeShapeListener)
	defer newTreeShapeListenerPool.Put(listener)
	// 执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.Infos
}

/* addObjectFromChangeLineNumber
/* @Description: 传入的参数较多，大致功能是构建object的map
 * @param commitDiff
 * @param changeLineNumber 行号变动
 * @param antlrAnalyzeRes antlr分析结果
 * @return objectInfoType
 * @author KevinMatt 2021-08-03 19:26:12
 * @function_mark PASS
*/
func addObjectFromChangeLineNumber(commitDiff global_type.DiffParsedType, changeLineNumber global_type.ChangeLineType, antlrAnalyzeRes AnalysisInfoType) global_type.ObjectInfoType {
	// 寻找变动方法
	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
	if changeMethod.MethodName == "" {
		// 为空直接跳过执行
		return global_type.ObjectInfoType{}
	}
	tempEncrypt := func(text string) string {
		return base64.StdEncoding.EncodeToString([]byte(utility.Encrypt(global_type.Config.ProjectId, "", text)))
	}
	var object global_type.ObjectInfoType
	object.Name = tempEncrypt(changeMethod.MethodName)
	object.Hash = tempEncrypt(commitDiff.CommitHash)
	object.ParentName = tempEncrypt(changeMethod.MasterObject.ObjectName)
	object.ParentHash = tempEncrypt(fmt.Sprintf("%x", hashCode64([]byte(global_type.Config.ProjectId), []byte(changeMethod.MasterObject.ObjectName), []byte(commitDiff.DiffFileName))))
	object.FilePath = tempEncrypt(commitDiff.DiffFileName)
	object.Owner = tempEncrypt(utility.ConCatStrings(commitDiff.CommitterName, "-", commitDiff.CommitterEmail))
	object.CommitTime = commitDiff.CommitTime
	object.OldName = ""
	return object
}

/* findChangedMethod
/* @Description: 寻找变动了的方法
 * @param changeLineNumber 变动行
 * @param antlrAnalyzeRes antlr分析结果
 * @return changeMethodInfo 变动方法信息
 * @author KevinMatt 2021-07-29 19:38:19
 * @function_mark PASS
*/
func findChangedMethod(changeLineNumber global_type.ChangeLineType, antlrAnalyzeRes AnalysisInfoType) (changeMethodInfo MethodInfoType) {
	var startLineNumbers []int
	// 遍历匹配到的方法列表，存储其首行
	for index := range antlrAnalyzeRes.AstInfoList.Methods {
		startLineNumbers = append(startLineNumbers, antlrAnalyzeRes.AstInfoList.Methods[index].StartLine)
	}
	// 寻找方法行所在的范围位置
	resIndex := findIntervalIndex(startLineNumbers, changeLineNumber.LineNumber)
	// 判断是否有位置插入
	if resIndex > -1 {
		changeMethodInfo = antlrAnalyzeRes.AstInfoList.Methods[resIndex]
	}
	return
}

/* findIntervalIndex
/* @Description: 寻找可插入位置
 * @param nums 传入的行号切片
 * @param target 要插入的目标行号
 * @return int 返回插入位置，-1代表无法插入
 * @author KevinMatt 2021-07-29 19:42:18
 * @function_mark PASS
*/
func findIntervalIndex(nums []int, target int) int {
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
