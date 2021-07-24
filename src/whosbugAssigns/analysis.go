package whosbugAssigns

import (
	javaParser "anrlr4_ast/java"
	"strings"
)

func TestParseCommit() []CommitParsedType {
	releaseDiff := getDiff("C:\\Users\\KevinMatt\\Desktop\\java-test\\", "master", "whosbug_test_1")
	return parseCommit(releaseDiff.Diff, strings.Split(releaseDiff.CommitInfo, "\n"))
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
func analysis(repoPath string, branchName string, projectId string) []CommitParsedType {
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
		analyzeCommitDiff1(projectId, commitDiffs, commitId, commit)
		allCommits = append(allCommits, commit)
	}
	return allCommits
}

/** analyzeCommitDiff
 * @Description:
 * @param projectId
 * @param commitDiffs
 * @param commitId
 * @param Commit
 * @author KevinMatt 2021-07-22 13:24:09
 * @function_mark
 */
func analyzeCommitDiff1(projectId string, commitDiffs []DiffParsedType, commitId string, commit CommitParsedType) {
	//for index := 0; index < len(commitDiffs); index++ {
	//	commitDiff := commitDiffs[index]
	//	commitDiff.Commit = commitId
	//	commitDiff.DiffContent = ""
	//	// 处理后的源码路径
	//	tempFile := commitDiff.DiffFilePath
	//	// diff的原始路径
	//	filePath := commitDiff.DiffFile
	//	antlrAnalyzeRes := antlrAnalysis(tempFile, "java")
	//
	//	ChangeLineNumbers := commitDiff.ChangeLineNumbers
	//	var objects ObjectInfoType
	//	// !注意，此处的changeLineNumbers内的map类型为map[string]string,作为整形使用时需要转型
	//	for _, changeLineNumber := range ChangeLineNumbers {
	//		addObjectFromChangeLineNumber(projectId, filePath, objects, changeLineNumber, antlrAnalyzeRes)
	//	}
	//	//CommitDiff.DiffContent =
	//	// TODO 重构addObjectFromChangeLineNumber()方法，使得commit["commit_diffs"]值作为切片类型生效
	//	//Commit.CommitDiff =
	//
	//}
}

//func addObjectFromChangeLineNumber(projectId string, filePath string, objects map[string]interface{}, changeLineNumber map[string]string, antlrAnalyzeRes javaParser.AnalysisInfoType) {
//	// TODO 重构findChangedMethod
//	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
//	if changeMethod == nil {
//		return
//	}
//	if _, ok := objects[changeMethod["startLine"]]; ok {
//		return
//	}
//	childHashCode := hashCode64(projectId, changeMethod["methodName"], filePath)
//	parent := changeMethod["masterObject"]
//
//	objects[changeMethod["startLine"]] = map[string]interface{}{
//		"Name":        changeMethod["methodName"],
//		"hash":        childHashCode,
//		"parent_name": parent["objectName"],
//		"parent_hash": hashCode64(projectId, parent["objectName"], filePath),
//	}
//}

/** findChangedMethod
 * @Description:
 * @author KevinMatt 2021-07-22 14:47:36
 * @function_mark
 */
func findChangedMethod(changeLineNumber map[string]string, antlrAnalyzeRes javaParser.AnalysisInfoType) map[string]interface{} {
	//TODO 重构
	startLineNumbers := make([]int, len(antlrAnalyzeRes.AstInfoList.Methods))
	for _, part := range antlrAnalyzeRes.AstInfoList.Methods {
		startLineNumbers = append(startLineNumbers, part.StartLine)
	}
	//resIndex := searchInsert(startLineNumbers, changeLineNumber["LineNumber"])
	return nil
}
func searchInsert(nums int, target int) int {
	return 0
}
