package whosbugAssigns

import javaParser "anrlr4_ast/java"

// analyzeCommitDiff
/* @Description: 分析commitDiff
 * @param projectId 项目ID
 * @param commitDiffs commitDiff切片
 * @param commitId CommitHash
 * @param commit 解析后的commit信息
 * @return CommitParsedType
 * @author KevinMatt 2021-07-25 13:54:04
 * @function_mark PASS
 */
func analyzeCommitDiff(projectId string, commitDiffs []DiffParsedType, commitId string, commit CommitParsedType) CommitParsedType {
	for index := 0; index < len(commitDiffs); index++ {
		commitDiff := commitDiffs[index]
		commitDiff.Commit = commitId
		// 处理后的源码路径
		tempFile := commitDiff.DiffFilePath
		// diff的原始路径
		filePath := commitDiff.DiffFile
		antlrAnalyzeRes := antlrAnalysis(tempFile, "java")

		changeLineNumbers := commitDiff.ChangeLineNumbers
		objects := make(map[int]map[string]string)
		for _, changeLineNumber := range changeLineNumbers {
			objects = addObjectFromChangeLineNumber(projectId, filePath, objects, changeLineNumber, antlrAnalyzeRes)
		}
		commitDiff.DiffContent = objects
		commit.CommitDiffs = append(commit.CommitDiffs, commitDiff)
	}
	return commit
}

// addObjectFromChangeLineNumber
/* @Description:  存储分析得到的方法改变信息(基于行号索引)
 * @param projectId 项目ID
 * @param filePath	文件目录
 * @param objects 	传入的空参数
 * @param changeLineNumber 改变行行号
 * @param antlrAnalyzeRes  分析结果
 * @return map[int]map[string]string
 * @author KevinMatt 2021-07-25 14:03:36
 * @function_mark PASS
 */
func addObjectFromChangeLineNumber(projectId string, filePath string, objects map[int]map[string]string, changeLineNumber ChangeLineNumberType, antlrAnalyzeRes javaParser.AnalysisInfoType) map[int]map[string]string {
	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
	if len(objects) > 0 {
		if _, ok := objects[changeMethod.StartLine]; ok {
			return objects
		}
	}
	childHashCode := hashCode64(projectId, changeMethod.MethodName, filePath)
	parent := changeMethod.MasterObject
	objects[changeMethod.StartLine] = make(map[string]string)
	objects[changeMethod.StartLine] = map[string]string{
		"name":        changeMethod.MethodName,
		"hash":        childHashCode,
		"parent_name": parent.ObjectName,
		"parent_hash": hashCode64(projectId, parent.ObjectName, filePath),
	}
	return objects
}

// findChangedMethod
/* @Description: 			寻找变更的方法
 * @param changeLineNumber 	变更行信息
 * @param antlrAnalyzeRes 	分析结果
 * @return javaParser.MethodInfoType
 * @author KevinMatt 2021-07-25 14:11:45
 * @function_mark PASS
 */
func findChangedMethod(changeLineNumber ChangeLineNumberType, antlrAnalyzeRes javaParser.AnalysisInfoType) javaParser.MethodInfoType {
	var changeMethodInfo javaParser.MethodInfoType
	startLineNumbers := make([]int, 0)
	for _, part := range antlrAnalyzeRes.AstInfoList.Methods {
		startLineNumbers = append(startLineNumbers, part.StartLine)
	}
	resIndex := findIntervalIndex(startLineNumbers, changeLineNumber.LineNumber)
	if resIndex > -1 {
		changeMethodInfo = antlrAnalyzeRes.AstInfoList.Methods[resIndex]
	}
	return changeMethodInfo
}

// findIntervalIndex
/* @Description: 	寻找插入空隙
 * @param nums		切片lineNumbers
 * @param target	目标lineNumber
 * @return int 		空隙位置
 * @author KevinMatt 2021-07-25 14:17:52
 * @function_mark 	PASS
 */
func findIntervalIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) >= 2 && target > nums[1] {
		return -1
	}
	if target <= nums[0] {
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
