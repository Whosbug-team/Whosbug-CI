package whosbugAssigns

import (
	"fmt"
	"os"
	"regexp"
)

// findAllChangedLineNumbers
/* @Description: 匹配所有改变行(以+/-开头的行)的行号
 * @param lines 传入diff中的所有代码行(完整文件代码行)
 * @return []ChangeLineNumberType 返回存储所有变更行信息的切片
 * @author KevinMatt 2021-07-25 13:47:42
 * @function_mark PASS
 */
func findAllChangedLineNumbers(lines []string) []ChangeLineNumberType {
	markCompile, err := regexp.Compile(`^[\+\-]`)
	errorHandler(err)
	changeLineNumbers := make([]ChangeLineNumberType, 0)
	lineNumber := 0
	for _, line := range lines {
		lineNumber++
		res := markCompile.FindString(line)
		if res != "" {
			var tempStruct ChangeLineNumberType
			tempStruct.LineNumber = lineNumber
			tempStruct.ChangeType = string(line[0])
			changeLineNumbers = append(changeLineNumbers, tempStruct)
		}
	}
	return changeLineNumbers
}

// getDiff
/* @Description: 获取release的diff信息
 * @param repoPath 仓库目录/url
 * @param branchName 分支名
 * @param projectId 项目id
 * @return ReleaseDiffType 返回releaseDiff结构体
 * @author KevinMatt 2021-07-25 13:12:07
 * @function_mark PASS
 */
func getDiff(repoPath, branchName, projectId string) ReleaseDiffType {

	secret := os.Getenv("WHOSBUG_SECRET")
	originPath, err := os.Getwd()
	errorHandler(err)
	err = os.Chdir(repoPath)
	errorHandler(err)
	fmt.Println("Work path changed to:", repoPath)

	newReleaseCommitHash := execCommandOutput("git", "rev-parse", "HEAD")

	originHash := make([]byte, len(projectId))
	err = encrypt([]byte(projectId), originHash, []byte(secret), []byte(projectId))
	errorHandler(err)
	getLatestRelease(string(originHash))
	lastReleaseCommitHash := make([]byte, len(originHash))

	err = decrypt([]byte(projectId), lastReleaseCommitHash, []byte(secret), originHash)
	if string(lastReleaseCommitHash) != string(originHash) {
		lastReleaseCommitHash = nil
	}
	errorHandler(err)
	fmt.Println("last release's Commit hash: ", string(lastReleaseCommitHash))
	fmt.Println("new release's Commit hash: ", newReleaseCommitHash)

	var diff, commitInfo string
	if string(lastReleaseCommitHash) != "" {
		diff = execCommandOutput("git", "log", "--full-diff", "-p", "-U1000", "--pretty=raw", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
		commitInfo = execCommandOutput("git", "log", "--pretty=format:%H,%ce,%cn,%cd", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
	} else {
		diff = execCommandOutput("git", "log", "--full-diff", "-p", "-U1000", "--pretty=raw")
		commitInfo = execCommandOutput("git", "log", "--pretty=format:%H,%ce,%cn,%cd")
	}
	var releaseDiff ReleaseDiffType
	releaseDiff.CommitInfo = commitInfo
	releaseDiff.Diff = diff
	releaseDiff.BranchName = branchName
	releaseDiff.HeadCommitInfo = newReleaseCommitHash

	// 返回原工作目录
	err = os.Chdir(originPath)
	fmt.Println("Work path changed back to:", originPath)
	errorHandler(err)

	return releaseDiff
}
