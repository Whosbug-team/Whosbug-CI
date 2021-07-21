package whosbugAssigns

import (
	"strings"
)

func analysis(repoPath string, branchName string, projectId string) []map[string]string {
	releaseDiff := getDiff(repoPath, branchName, projectId)
	commits := parseCommit(releaseDiff["diff"], strings.Split(releaseDiff["commit_info"], "\n"))
	var allCommits []map[string]string

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
		//commt["commit_diffs"] = ""

	}
	return allCommits
}
