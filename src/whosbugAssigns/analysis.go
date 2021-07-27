package whosbugAssigns

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//Analysis
/* Analysis
/* @Description: 分析调用主逻辑函数
 * @param repoPath 仓库地址/url
 * @param branchName 分支名
 * @param projectId 项目id
 * @return []CommitParsedType 返回解析后的commit信息
 * @author KevinMatt 2021-07-25 13:08:45
 * @function_mark PASS
*/
func Analysis(repoPath, branchName, projectId string) []CommitParsedType {
	t := time.Now()
	secret = os.Getenv("WHOSBUG_SECRET")
	if secret == "" {
		secret = "defaultsecret"
	}
	//获取log
	releaseDiff := getLogInfo(repoPath, branchName, projectId)
	//解析logCommit
	commits := parseCommit(releaseDiff.Diff, strings.Split(releaseDiff.CommitInfo, "\n"))
	var parsedCommits []CommitParsedType
	for index := range commits {
		commit := commits[index]
		var diffPark string
		if index == len(commits)-1 {
			diffPark = releaseDiff.Diff[commit.CommitLeftIndex:]
		} else {
			nextCommitLeftIndex := commits[index+1].CommitLeftIndex
			diffPark = releaseDiff.Diff[commit.CommitLeftIndex:nextCommitLeftIndex]
		}
		// 解析diff信息
		commitDiffs := parseDiff(diffPark)
		// commitDiffs切片引用传递，不需返回值
		analyzeCommitDiff(projectId, commitDiffs, commit.Commit)
		commit.CommitDiffs = append(commit.CommitDiffs, commitDiffs...)
		parsedCommits = append(parsedCommits, commit)
	}
	fmt.Println("Analysis cost: ", time.Since(t))
	return parsedCommits
}
