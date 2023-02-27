package analyze

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
	"strings"

	"git.woa.com/bkdevops/whosbug-ci/internal/util"
	"git.woa.com/bkdevops/whosbug-ci/internal/zaplog"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/crypto"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/logging"

	"github.com/panjf2000/ants"
)

var (
	// processDiffs 已处理的commit数
	processDiffs int
	// AntlrAnalysisPool 解析协程池
	AntlrAnalysisPool, _ = ants.NewPoolWithFunc(runtime.NumCPU()/4, func(commitDiff any) {
		GenerateObjects(commitDiff.(Diff))
		processDiffs++
		log.SetOutput(logging.LogFile)
		log.Println("Diff No.", processDiffs, " From", commitDiff.(Diff).CommitHash, " Sent Into Channel.")
		// if processDiffs%50 == 0 {
		// runtime.GC()
		// }
	})
)

// GenerateObjects antlr分析commitDiff信息生成objects
//
//	@param commitDiff Diff
//	@author kevineluo
//	@update 2023-02-28 02:45:39
func GenerateObjects(commitDiff Diff) {
	// 进行antlr静态解析
	parser, err := initAstParser(commitDiff.TargetLanguage)
	if err != nil {
		zaplog.Logger.Error("[GenerateObjects] error when initAstParser", zaplog.Error(err))
		return
	}
	astInfo := parser.AstParse(commitDiff.DiffText)

	var prevObject ObjectInfo
	var countMinus int
	var countPlus int
	for _, changeLine := range commitDiff.ChangeLines {
		currentObject := addObject(commitDiff, changeLine, astInfo)
		if currentObject.Equals(ObjectInfo{}) {
			continue
		}
		if prevObject.Equals(ObjectInfo{}) {
			prevObject = currentObject
			continue
		}
		if currentObject.Equals(prevObject) {
			if changeLine.ChangeType == "-" {
				countMinus++
			} else if changeLine.ChangeType == "+" {
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

func initAstParser(langMode string) (parser antlr.AstParser, err error) {
	defer Recover()
	switch langMode {
	case "c":
		// TODO: C Unsupported
		parser = new(antlr.CAstParser)
	case "java":
		parser = new(antlr.JavaAstParser)
	case "python":
		// TODO: Python Unsupported
	case "kotlin":
		parser = new(antlr.KotlinAstParser)
	case "golang":
		parser = new(antlr.GoAstParser)
	case "javascript":
		// TODO: Js Unsupported
		parser = new(antlr.JSAstParser)
	case "cpp":
		// TODO: too Slow
		parser = new(antlr.CppAstParser)
	default:
		err = ErrUnsupportedLanguage
		return
	}
	err = parser.Init()
	return
}

// addObject
//
//	@Description: 传入的参数较多，大致功能是构建object的map
//	@param commitDiff
//	@param changeLineNumber 行号变动
//	@param antlrAnalyzeRes antlr分析结果
//	@return objectInfoType
//	@author KevinMatt 2021-08-03 19:26:12
//	@function_mark PASS
func addObject(commitDiff Diff, changeLine ChangeLine, astInfo antlr.AstInfo) (newObject ObjectInfo) {
	//	寻找变动方法
	changeMethod := findChangedMethod(changeLine, astInfo)
	if changeMethod.Name == "" {
		return
	}
	oldMethodName := findFather(changeMethod.Name)
	if oldMethodName != "" {
		addClass(commitDiff, oldMethodName, astInfo)
	}

	//	TODO Ready for newMethod
	newObject = ObjectInfo{
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

// addClass
//
//	@Description: 寻找类的起始行
//	@param oldMethodName 类的定义链
//	@param antlrAnalyzeRes antlr分析结果
//	@return changeMethodInfo 类信息
//	@author Psy 2022-08-17 15:33:33
func addClass(commitDiff Diff, preMethodName string, antlrAnalyzeRes antlr.AstInfo) {
	idx := strings.LastIndex(preMethodName, ".")
	if idx == -1 {
		return
	}
	newObj := ObjectInfo{}
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

// findChangedMethod 寻找变动了的方法
//
//	@param changeLineNumber ChangeLine
//	@param astInfo antlr.AstInfo
//	@return changeMethodInfo antlr.Method
//	@author kevineluo
//	@update 2023-02-28 03:00:19
func findChangedMethod(changeLineNumber ChangeLine, astInfo antlr.AstInfo) (changeMethodInfo antlr.Method) {
	var lineRangeList []antlr.LineRange
	//	遍历匹配到的方法列表，存储其首行
	for index := range astInfo.Methods {
		lineRangeList = append(lineRangeList, antlr.LineRange{
			StartLine: astInfo.Methods[index].StartLine,
			EndLine:   astInfo.Methods[index].EndLine,
		})
	}
	//	寻找方法行所在的范围位置
	resIndex := findIntervalIndex(lineRangeList, changeLineNumber.LineNumber)
	//	判断是否有位置插入
	if resIndex > -1 {
		changeMethodInfo = astInfo.Methods[resIndex]
	}
	return
}

// findIntervalIndex 寻找可插入位置
//
//	@param lineRangeList []antlr.LineRange
//	@param target float64
//	@return int 返回插入位置，-1代表无法插入
//	@author kevineluo
//	@update 2023-02-28 02:59:57
func findIntervalIndex(lineRangeList []antlr.LineRange, target float64) int {
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

// Recover 防止语法分析识别出错时，程序终止
//
//	@author kevineluo
//	@update 2023-02-28 03:03:26
func Recover() {
	if err := recover(); err != nil {
		errMsg := fmt.Sprintf("======== Panic ========\nPanic: %v\nTraceBack:\n%s\n======== Panic ========", err, string(debug.Stack()))
		zaplog.Logger.DPanic(errMsg)
	}
}
