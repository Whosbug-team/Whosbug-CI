package whosbugAssigns

import (
	javaParser "anrlr4_ast/java"
	"strings"
)

func TestParseCommit() []CommitParsedType {
	releaseDiff := getDiff("C:\\Users\\KevinMatt\\Desktop\\java-test\\", "master", "whosbug_test_1")
	return parseCommit(releaseDiff.Diff, strings.Split(releaseDiff.CommitInfo, "\n"))
}
func AnalysisTest(repoPath, branchName, projectId string) []CommitParsedType {
	return analysis(repoPath, branchName, projectId)
}

/** analysis
 * @Description: 分析逻辑主函数
 * @param repoPath 仓库地址
 * @param BranchName 分支
 * @param projectId projectID
 * @return []map[string]interface{} 分析结果
 * @author KevinMatt 2021-07-22 13:24:31
 * @function_mark
 */
func analysis(repoPath, branchName, projectId string) []CommitParsedType {
	releaseDiff := getDiff(repoPath, branchName, projectId)
	commits := parseCommit(releaseDiff.Diff, strings.Split(releaseDiff.CommitInfo, "\n"))
	// allCommits:interface实际存储的内容为string
	var allCommits []CommitParsedType
	for index := range commits {
		commit := commits[index]
		commitId := commit.Commit
		var diffPark string
		if index == len(commits)-1 {
			diffPark = releaseDiff.Diff[commit.CommitLeftIndex:]
		} else {
			nextCommitLeftIndex := commits[index+1].CommitLeftIndex
			diffPark = releaseDiff.Diff[commit.CommitLeftIndex:nextCommitLeftIndex]
		}
		commitDiffs := parseDiff(diffPark)
		commit = analyzeCommitDiff(projectId, commitDiffs, commitId, commit)
		allCommits = append(allCommits, commit)
	}
	return allCommits
}

/** analyzeCommitDiff
 * @Description:
 * @param projectId
 * @param commitDiffs
 * @param commitId
 * @param commit
 * @return CommitParsedType
 * @author KevinMatt 2021-07-25 03:57:26
 * @function_mark
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
		// !注意，此处的changeLineNumbers内的map类型为map[string]string,作为整形使用时需要转型
		for _, changeLineNumber := range changeLineNumbers {
			objects = addObjectFromChangeLineNumber(projectId, filePath, objects, changeLineNumber, antlrAnalyzeRes)
		}
		commitDiff.DiffContent = objects
		// TODO 重构addObjectFromChangeLineNumber()方法，使得commit["commit_diffs"]值作为切片类型生效
		commit.CommitDiffs = append(commit.CommitDiffs, commitDiff)
	}
	return commit
}

func addObjectFromChangeLineNumber(projectId string, filePath string, objects map[int]map[string]string, changeLineNumber ChangeLineNumberType, antlrAnalyzeRes javaParser.AnalysisInfoType) map[int]map[string]string {
	// TODO 重构findChangedMethod
	//objects := make(map[int]interface{})
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
		"Name":        changeMethod.MethodName,
		"hash":        childHashCode,
		"parent_name": parent.ObjectName,
		"parent_hash": hashCode64(projectId, parent.ObjectName, filePath),
	}
	return objects
}

/** findChangedMethod
 * @Description:
 * @author KevinMatt 2021-07-22 14:47:36
 * @function_mark
 */
func findChangedMethod(changeLineNumber ChangeLineNumberType, antlrAnalyzeRes javaParser.AnalysisInfoType) javaParser.MethodInfoType {
	//TODO 重构
	var changeMethodInfo javaParser.MethodInfoType
	startLineNumbers := make([]int, 0)
	for _, part := range antlrAnalyzeRes.AstInfoList.Methods {
		startLineNumbers = append(startLineNumbers, part.StartLine)
	}
	resIndex := searchInsert(startLineNumbers, changeLineNumber.LineNumber)
	if resIndex > -1 {
		changeMethodInfo = antlrAnalyzeRes.AstInfoList.Methods[resIndex]
	}
	return changeMethodInfo
}
func searchInsert(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	if len(nums) >= 2 && target > nums[1] {
		return -1
	}
	if target < nums[0] {
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
