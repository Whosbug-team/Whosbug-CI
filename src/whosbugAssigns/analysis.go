package whosbugAssigns

import (
	"strings"
)

func TestParseCommit() []map[string]interface{} {
	releaseDiff := getDiff("C:\\Users\\KevinMatt\\Desktop\\java-test\\", "master", "whosbug_test_1")
	return parseCommit(releaseDiff["diff"], strings.Split(releaseDiff["commit_info"], "\n"))
}

/** analysis
 * @Description: 分析逻辑主函数
 * @param repoPath 仓库地址
 * @param branchName 分支
 * @param projectId projectID
 * @return []map[string]interface{} 分析结果
 * @author KevinMatt 2021-07-22 13:24:31
 * @function_mark
 */
func analysis(repoPath string, branchName string, projectId string) []map[string]interface{} {
	releaseDiff := getDiff(repoPath, branchName, projectId)
	commits := parseCommit(releaseDiff["diff"], strings.Split(releaseDiff["commit_info"], "\n"))
	// allCommits:interface实际存储的内容为string
	var allCommits []map[string]interface{}

	for index := range commits {
		commit := commits[index]
		commitId := commit["commit"]
		var diffPark string
		if index == len(commits)-1 {
			diffPark = releaseDiff["diff"][commit["commit_left_index"].(int):]
		} else {
			nextCommitLeftIndex := commits[index+1]["commit_left_index"].(int)
			diffPark = releaseDiff["diff"][commit["commit_left_index"].(int):nextCommitLeftIndex]
		}

		commitDiffs := parseDiff(diffPark)
		//TODO 重构 analyze_commit_diff函数
		//commit["commit_diffs"] = interface {}
		analyzeCommitDiff(projectId, commitDiffs, commitId.(string), commit)
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
 * @author KevinMatt 2021-07-22 13:24:09
 * @function_mark
 */
func analyzeCommitDiff(projectId string, commitDiffs []map[string]interface{}, commitId string, commit map[string]interface{}) {
	for index := 0; index < len(commitDiffs); index++ {
		commitDiff := commitDiffs[index]
		commitDiff["commit"] = commitId
		commitDiff["diff_content"] = ""
		// 处理后的源码路径
		tempFile := commitDiff["diff_file_path"].(string)
		// diff的原始路径
		filePath := commitDiff["diff_file"].(string)
		antlrAnalyzeRes := antlrAnalysis(tempFile, "java")

		changeLineNumbers := commitDiff["change_line_numbers"].([]map[string]string)
		var objects map[string]interface{}
		// !注意，此处的changeLineNumbers内的map类型为map[string]string,作为整形使用时需要转型
		for _, changeLineNumber := range changeLineNumbers {
			addObjectFromChangeLineNumber(projectId, filePath, objects, changeLineNumber, antlrAnalyzeRes)
		}

		commitDiff["diff_content"] = objects
		// TODO 重构addObjectFromChangeLineNumber()方法，使得commit["commit_diffs"]值作为切片类型生效
		//commit["commit_diffs"] = append(commit["commit_diffs"], commitDiff)
	}
}

func addObjectFromChangeLineNumber(projectId string, filePath string, objects map[string]interface{}, changeLineNumber map[string]string, antlrAnalyzeRes string) {
	// TODO 重构findChangedMethod
	//changeMethod := findChangeMethod()
}

/** findChangedMethod
 * @Description:
 * @author KevinMatt 2021-07-22 14:47:36
 * @function_mark
 */
func findChangedMethod(changeLineNumber map[string]string, antlrAnalyzeRes) {
	//TODO 重构

}
