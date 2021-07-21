package whosbugAssigns

import (
	"strings"
)

func analysis(repoPath string, branchName string, projectId string) []map[string]string {
	releaseDiff := getDiff(repoPath, branchName, projectId)
	commits := parseCommit(releaseDiff["diff"], strings.Split(releaseDiff["commit_info"], "\n"))
	// allCommits:interface实际存储的内容为
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

// @title analyzeCommitDiff
// @description 分析commit diff情况
// @author KevinMatt
// @param projectId:项目唯一标识
// @param commitDiffs 一次commit的所有变更文件（一个diff对应一个变更文件）
// @param commitId 该次commit的版本号
// @param commit 包含该次commit的所有diff objects的map（commit['commit_diffs']即为该次commit的所有diff objects）
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
		commit["commit_diffs"] = append(commit["commit_diffs"], commitDiff)
	}
}

func addObjectFromChangeLineNumber(projectId string, filePath string, objects map[string]string, changeLineNumber map[string]string, antlrAnalyzeRes string) {
	// TODO 重构findChangedMethod
	//changeMethod := findChangeMethod()
}