package whosbugAssigns

import "strings"

func analysis(repoPath string, branchName string, projectId string) {
	releaseDiff := getDiff(repoPath, branchName, projectId)
	commits := parseCommit(releaseDiff["diff"], strings.Split(releaseDiff["commit_info"], "\n"))

}
