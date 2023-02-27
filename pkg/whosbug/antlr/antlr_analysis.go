package antlr

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
	"strings"

	"git.woa.com/bkdevops/whosbug-ci/internal/util"
	"git.woa.com/bkdevops/whosbug-ci/internal/zaplog"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/crypto"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/logging"

	"github.com/panjf2000/ants"
)

var (
	// processDiffs 已处理的commit数
	processDiffs int
	// AntlrAnalysisPool 解析协程池
	AntlrAnalysisPool, _ = ants.NewPoolWithFunc(runtime.NumCPU()/4, func(commitDiff interface{}) {
		AnalyzeCommitDiff(commitDiff.(DiffParsedType))
		// 指示已经处理的diff数量
		processDiffs++
		log.SetOutput(logging.LogFile)
		log.Println("Diff No.", processDiffs, " From", commitDiff.(DiffParsedType).CommitHash, " Sent Into Channel.")
		// if processDiffs%50 == 0 {
		// runtime.GC()
		// }
	})
)

// AnalyzeCommitDiff //	@Description: 使用antlr分析commitDiff信息
//
//	@param commitDiff config.DiffParsedType
//	@author kevineluo
//	@update 2023-02-25 04:23:05
func AnalyzeCommitDiff(commitDiff DiffParsedType) {
	//	获取antlr分析结果
	// TODO: 抽象ast analyzer interface
	antlrAnalyzeRes := antlrAnalysis(commitDiff.DiffText, commitDiff.TargetLanguage)

	var prevObject ObjectInfoType
	var countMinus int
	var countPlus int
	for _, changeLineNumber := range commitDiff.ChangeLineNumbers {
		currentObject := addObjectFromChangeLineNumber(commitDiff, changeLineNumber, antlrAnalyzeRes)
		if currentObject.Equals(ObjectInfoType{}) {
			continue
		}
		if prevObject.Equals(ObjectInfoType{}) {
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
				ObjectChan <- prevObject
			}
			countMinus = 0
			countPlus = 0
			prevObject = currentObject
		}
		// 用于比较两次的结构体是否重复(匹配行范围导致的重复结果)
	}
}

// 防止语法分析识别出错时，程序终止
func myRecover() {
	if err := recover(); err != nil {
		errMsg := fmt.Sprintf("======== Panic ========\nPanic: %v\nTraceBack:\n%s\n======== Panic ========", err, string(debug.Stack()))
		zaplog.Logger.DPanic(errMsg)
	}
}

// antlrAnalysis
//
//	@Description: antlr分析过程
//	@param targetFilePath 分析的目标文件
//	@param langMode 分析的语言模式
//	@return AstResType 返回分析信息结构体
//	@author KevinMatt 2021-07-29 19:49:37
//	@function_mark  PASS
func antlrAnalysis(diffText string, langMode string) (result AstInfo) {
	defer myRecover()
	switch langMode {
	case "c":
		// TODO: C Unsupported
		// result = ExecuteC(diffText)
	case "java":
		// result = ExecuteJava(diffText)
	case "python":
		// TODO: Python Unsupported
		// result = ExecutePython(diffText)
	case "kotlin":
		// result = ExecuteKotlin(diffText)
	case "golang":
		// result = ExecuteGolang(diffText)
	case "javascript":
		// TODO: Js Unsupported
		// result = ExecuteJavaScript(diffText)
	case "cpp":
		// TODO: too Slow
		// result = ExecuteCpp(diffText)
	default:
		break
	}
	return
}

func ExecutePython(diffText string) AstInfo {
	util.ForDebug(diffText)
	return AstInfo{}
}

// addObjectFromChangeLineNumber
//
//	@Description: 传入的参数较多，大致功能是构建object的map
//	@param commitDiff
//	@param changeLineNumber 行号变动
//	@param antlrAnalyzeRes antlr分析结果
//	@return objectInfoType
//	@author KevinMatt 2021-08-03 19:26:12
//	@function_mark PASS
func addObjectFromChangeLineNumber(commitDiff DiffParsedType, changeLineNumber ChangeLineType, antlrAnalyzeRes AstInfo) (newObject ObjectInfoType) {
	//	寻找变动方法
	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
	if changeMethod.Name == "" {
		return
	}
	oldMethodName := findFather(changeMethod.Name)
	if oldMethodName != "" {
		addClass(commitDiff, oldMethodName, antlrAnalyzeRes)
	}

	//	TODO Ready for newMethod
	newObject = ObjectInfoType{
		CommitHash:       commitDiff.CommitHash, //crypto.Base64Encrypt(commitDiff.CommitHash)
		ID:               crypto.Base64Encrypt(changeMethod.Name),
		OldID:            crypto.Base64Encrypt(oldMethodName),
		FilePath:         crypto.Base64Encrypt(commitDiff.DiffFileName),
		Parameters:       crypto.Base64Encrypt(changeMethod.Parameters),
		OldLineCount:     0,
		CurrentLineCount: changeMethod.EndLine - changeMethod.StartLine + 1,
		StartLine:        changeMethod.StartLine,
		EndLine:          changeMethod.EndLine,
	}
	return
}

// findFather
//
//	@Description: 寻找定义链的上端
//	@param methodName 定义链末尾的名字
//	@return oldMethodName 定义链上端的名字
//	@author Psy 2022-08-17 20:23:21
func findFather(methodName string) (oldMethodName string) {
	oldMethodName = ""
	oldMethodNameIdx := strings.LastIndex(methodName, ".")
	if oldMethodNameIdx != -1 {
		oldMethodName = methodName[:oldMethodNameIdx]
	}
	return
}

// adddClass
//
//	@Description: 寻找类的起始行
//	@param oldMethodName 类的定义链
//	@param antlrAnalyzeRes antlr分析结果
//	@return changeMethodInfo 类信息
//	@author Psy 2022-08-17 15:33:33
func addClass(commitDiff DiffParsedType, preMethodName string, antlrAnalyzeRes AstInfo) {
	idx := strings.LastIndex(preMethodName, ".")
	if idx == -1 {
		return
	}
	newObj := ObjectInfoType{}
	methodName := preMethodName[:idx]

	resIndex := -1
	for index := range antlrAnalyzeRes.Classes {
		if antlrAnalyzeRes.Classes[index].Name == methodName {
			resIndex = index
			break
		}
	}
	if resIndex > -1 {
		oldMethodName := findFather(methodName)
		newObj.CommitHash = commitDiff.CommitHash
		newObj.ID = crypto.Base64Encrypt(methodName)
		newObj.OldID = crypto.Base64Encrypt(oldMethodName)
		newObj.FilePath = crypto.Base64Encrypt(commitDiff.DiffFileName)
		newObj.StartLine = antlrAnalyzeRes.Classes[resIndex].StartLine
		newObj.EndLine = antlrAnalyzeRes.Classes[resIndex].EndLine
		ObjectChan <- newObj
		if oldMethodName != "" {
			addClass(commitDiff, oldMethodName, antlrAnalyzeRes)
		}
	}
}

// findChangedMethod
//
//	@Description: 寻找变动了的方法
//	@param changeLineNumber 变动行
//	@param antlrAnalyzeRes antlr分析结果
//	@return changeMethodInfo 变动方法信息
//	@author KevinMatt 2021-07-29 19:38:19
//	@function_mark PASS
func findChangedMethod(changeLineNumber ChangeLineType, antlrAnalyzeRes AstInfo) (changeMethodInfo Method) {
	var lineRangeList []LineRange
	//	遍历匹配到的方法列表，存储其首行
	for index := range antlrAnalyzeRes.Methods {
		lineRangeList = append(lineRangeList, LineRange{
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
//
//	@Description: 寻找可插入位置
//	@param nums 传入的行号切片
//	@param target 要插入的目标行号
//	@return int 返回插入位置，-1代表无法插入
//	@author KevinMatt 2021-07-29 19:42:18
//	@function_mark PASS
func FindIntervalIndex(lineRangeList []LineRange, target float64) int {
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
