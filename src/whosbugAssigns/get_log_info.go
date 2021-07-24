package whosbugAssigns

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

// @title findAllChangedLineNumbers
// @description 匹配所有改变行(以+/-开头的行)的行号
// @author KevinMatt
// @param lines []string 传入diff中的所有代码行(完整文件代码行)
// @return ChangeLineNumbers []map[string]string:{ {"line_number":xxx,"change_type":xxx}... }
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

// @title findAllChangedLines
// @description
// @auth KevinMatt
// @param lines []string diff中所有的代码行
// @return ChangedLineInfos []map[string]string 返回map的切片，格式类似：ChangedLineInfos={
//	{"LineNumber": 行号, "ChangeType": 变动类型},
//	{"LineNumber": 行号, "ChangeType": 变动类型},
//    ...
//}

/** findAllChangedLines
 * @Description: 匹配所有代码变更行并存储到切片，匹配的正则规范为：+/-行被匹配
 * @param lines  diff中所有的代码行
 * @return []map[string]string
 * @author KevinMatt 2021-07-24 23:07:56
 * @function_mark
 */
func findAllChangedLines(lines []string) []map[string]string {

	re, err := regexp.Compile(`^[\+\-]`)
	errorHandler(err)
	// ChangedLineInfos 匹配到的所有改变行信息(行号&改变类型)
	var ChangedLineInfos []map[string]string
	lineNumber := 0
	// 遍历匹配
	for index := 0; index < len(lines); index++ {
		lineNumber++
		match := re.FindString(lines[index])
		line := lines[index]
		if match != "" {
			tempMap := make(map[string]string)
			tempMap["LineNumber"] = strconv.Itoa(lineNumber)
			tempMap["ChangeType"] = string(line[0])
			ChangedLineInfos = append(ChangedLineInfos, tempMap)
		}
	}
	return ChangedLineInfos
}

func GetDiffTest(repoPath string, branchName string, projectId string) ReleaseDiffType {
	return getDiff(repoPath, branchName, projectId)
}

/** getDiff
 * @Description:
 * @param repoPath
 * @param BranchName
 * @param projectId
 * @return map[string]string
 * @author KevinMatt 2021-07-22 13:56:12
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
		diff = execCommandOutput("git", "log", "--full-Diff", "-p", "-U1000", "--pretty=raw", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
		commitInfo = execCommandOutput("git", "log", "--pretty=format:%H,%ce,%cn,%cd", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
	} else {
		diff = execCommandOutput("git", "log", "--full-Diff", "-p", "-U1000", "--pretty=raw")
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

/** getLatestRelease
 * @Description: 获得最新的Release信息
 * @param projectId 项目ID
 * @return string Release信息
 * @author KevinMatt 2021-07-22 16:50:26
 * @function_mark
 */
func getLatestRelease(projectId string) string {
	token := genToken()
	urls := HOST + "/release/last/"
	headers := make(map[string]string)
	headers["token"] = token
	data := make(map[string]string)
	data["pid"] = projectId

	getLatestReleaseRes, err := http.PostForm(urls, url.Values{"token": {token}, "pid": {projectId}})
	errorHandler(err)
	if getLatestReleaseRes.StatusCode == 200 {
		res, err := ioutil.ReadAll(getLatestReleaseRes.Body)
		errorHandler(err)
		return string(res)
	} else {
		fmt.Println(getLatestReleaseRes.Body)
		return ""
	}
}
