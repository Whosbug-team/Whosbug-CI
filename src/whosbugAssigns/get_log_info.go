package whosbugAssigns

import (
	"fmt"
	"os"
	"time"
)

/* getLogInfo
/* @Description: 获取release的diff信息
 * @param repoPath 仓库目录/url
 * @param branchName 分支名
 * @param projectId 项目id
 * @return ReleaseDiffType 返回releaseDiff结构体
 * @author KevinMatt 2021-07-26 20:48:22
 * @function_mark PASS
*/
func getLogInfo(repoPath, branchName, projectId string) ReleaseDiffType {
	t := time.Now()
	originPath, err := os.Getwd()
	errorHandler(err, "Get present path ")
	err = os.Chdir(repoPath)
	errorHandler(err, "Cd to ", repoPath)
	fmt.Println("Work path changed to:", repoPath)

	newReleaseCommitHash := execCommandOutput("git", "rev-parse", "HEAD")

	originHash := encrypt(projectId, secret, projectId)
	errorHandler(err, "encrypt Error")
	getLatestRelease(originHash)

	lastReleaseCommitHash := decrypt(projectId, secret, originHash)
	if lastReleaseCommitHash == originHash {
		lastReleaseCommitHash = originHash
	}
	// 执行命令获得原始字符串
	diff, commitInfo := getDiffInfo(lastReleaseCommitHash, newReleaseCommitHash)
	var releaseDiff ReleaseDiffType
	releaseDiff.CommitInfo, releaseDiff.Diff, releaseDiff.BranchName, releaseDiff.HeadCommitInfo = commitInfo, diff, branchName, newReleaseCommitHash
	// 返回原工作目录
	err = os.Chdir(originPath)
	fmt.Println("Work path changed back to:", originPath)
	errorHandler(err, "Cd to ", originPath)
	fmt.Println("getLogInfo cost ", time.Since(t))
	return releaseDiff
}

/* getDiffInfo
/* @Description: 根据命令行获取diff信息和commit信息
 * @param lastReleaseCommitHash
 * @param newReleaseCommitHash
 * @return diffInfo
 * @return commitInfo
 * @author KevinMatt 2021-07-26 20:44:07
 * @function_mark
*/
func getDiffInfo(lastReleaseCommitHash, newReleaseCommitHash string) (diffInfo string, commitInfo string) {
	if lastReleaseCommitHash == "" {
		diffInfo = execCommandOutput("git", "log", "--full-diff", "-p", "-U1000", "--pretty=raw", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
		commitInfo = execCommandOutput("git", "log", "--pretty=format:%H,%ce,%cn,%cd", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
	} else {
		diffInfo = execCommandOutput("git", "log", "--full-diff", "-p", "-U1000", "--pretty=raw")
		commitInfo = execCommandOutput("git", "log", "--pretty=format:%H,%ce,%cn,%cd")
		fmt.Println("last release's Commit hash: ", lastReleaseCommitHash)
	}
	fmt.Println("new release's Commit hash: ", newReleaseCommitHash)
	return
}
