package whosbugAssigns

import (
	"strings"
)

// Analysis
/* @Description: 分析调用主逻辑函数
 * @param repoPath 仓库地址/url
 * @param branchName 分支名
 * @param projectId 项目id
 * @return []CommitParsedType 返回解析后的commit信息
 * @author KevinMatt 2021-07-25 13:08:45
 * @function_mark PASS
 */
func Analysis(repoPath, branchName, projectId string) []CommitParsedType {
	releaseDiff := getDiff(repoPath, branchName, projectId)
	commits := parseCommit(releaseDiff.Diff, strings.Split(releaseDiff.CommitInfo, "\n"))
	var parsedCommits []CommitParsedType
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
		parsedCommits = append(parsedCommits, commit)
	}
	return parsedCommits
}
