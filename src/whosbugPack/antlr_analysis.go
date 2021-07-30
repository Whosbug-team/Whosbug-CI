package whosbugPack

import (
	javaparser "anrlr4_ast/java"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"sync"
)

type TreeShapeListener struct {
	Infos AnalysisInfoType
}

/* newTreeShapeListener
/* @Description: 创建新的listener
 * @return *TreeShapeListener
 * @author KevinMatt 2021-07-29 20:08:20
 * @function_mark PASS
*/

var (
	lexerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return javaparser.NewJavaLexer(nil)
	}}
	parserPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return javaparser.NewJavaParser(nil)
	}}
	newTreeShapeListener *sync.Pool = &sync.Pool{New: func() interface{} {
		return new(TreeShapeListener)
	}}
)

// AnalyzeCommitDiff
/* @Description: 使用antlr分析commitDiff信息
 * @param commitDiff diff信息(path)
 * @param commitId commit的Hash值
 * @author KevinMatt 2021-07-29 22:48:28
 * @function_mark
 */
func AnalyzeCommitDiff(commitDiff diffParsedType) diffParsedType {
	// 源码路径(仓库路径)
	filePath := commitDiff.diffFileName

	// 获取antlr分析结果
	antlrAnalyzeRes := antlrAnalysis(commitDiff.diffText, "java")

	// 创建要存入的objects
	objects := make(map[int]map[string]string)

	for _, changeLineNumber := range commitDiff.changeLineNumbers {
		// 根据行号添加object
		temp := addObjectFromChangeLineNumber(filePath, objects, changeLineNumber, antlrAnalyzeRes)
		if temp != nil {
			objects = temp
		}
	}
	commitDiff.diffContent = objects
	return commitDiff
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
		//// 解析前置空javaparser的Infos结构体
		//javaparser.Infos.SetEmpty()
		result = executeJava(diffText)
	default:
		break
	}
	return result
}

/* executeJava
/* @Description: 执行java分析
 * @param targetFilePath 分析目标路径
 * @return javaparser.AnalysisInfoType 返回分析结果结构体
 * @author KevinMatt 2021-07-29 19:51:16
 * @function_mark PASS
*/
func executeJava(diffText string) AnalysisInfoType {
	// 截取目标文件的输入流
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
	// 移除错误诊断监听，尝试提高性能
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// 构建语法解析树
	p.BuildParseTrees = true
	// ! 启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	// 解析模式->每个编译单位
	tree := p.CompilationUnit()
	// 创建listener
	listener := newTreeShapeListener.Get().(*TreeShapeListener)
	// 执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	defer newTreeShapeListener.Put(listener)
	return listener.Infos
}

/* addObjectFromChangeLineNumber
/* @Description: 传入的参数较多，大致功能是构建object的map
 * @param filePath 变动的文件名
 * @param objects 如变量名
 * @param changeLineNumber 行号变动
 * @param antlrAnalyzeRes antlr分析结果
 * @return map[int]map[string]string 返回object
 * @author KevinMatt 2021-07-29 19:31:58
 * @function_mark PASS
*/
func addObjectFromChangeLineNumber(fileName string, objects map[int]map[string]string, changeLineNumber changeLineType, antlrAnalyzeRes AnalysisInfoType) map[int]map[string]string {
	// 寻找变动方法
	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
	if changeMethod.MethodName == "" {
		return nil
	}
	// 判断object中是否有重复元素
	if len(objects) > 0 {
		if _, ok := objects[changeMethod.StartLine]; ok {
			return objects
		}
	}

	// 装入变量
	objects[changeMethod.StartLine] = make(map[string]string)
	objects[changeMethod.StartLine] = map[string]string{
		"name":        changeMethod.MethodName,
		"hash":        fmt.Sprintf("%x", hashCode64([]byte(config.ProjectId), []byte(changeMethod.MethodName), []byte(fileName))),
		"parent_name": changeMethod.MasterObject.ObjectName,
		"parent_hash": fmt.Sprintf("%x", hashCode64([]byte(config.ProjectId), []byte(changeMethod.MasterObject.ObjectName), []byte(fileName))),
	}
	return objects
}

/* findChangedMethod
/* @Description: 寻找变动了的方法
 * @param changeLineNumber 变动行
 * @param antlrAnalyzeRes antlr分析结果
 * @return changeMethodInfo 变动方法信息
 * @author KevinMatt 2021-07-29 19:38:19
 * @function_mark PASS
*/
func findChangedMethod(changeLineNumber changeLineType, antlrAnalyzeRes AnalysisInfoType) (changeMethodInfo MethodInfoType) {
	var startLineNumbers []int
	// 遍历匹配到的方法列表，存储其首行
	for index := range antlrAnalyzeRes.AstInfoList.Methods {
		startLineNumbers = append(startLineNumbers, antlrAnalyzeRes.AstInfoList.Methods[index].StartLine)
	}

	// 寻找方法行可以插入的位置
	resIndex := findIntervalIndex(startLineNumbers, changeLineNumber.lineNumber)

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
