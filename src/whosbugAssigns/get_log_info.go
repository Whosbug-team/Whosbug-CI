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
// @return changeLineNumbers []map[string]string:{ {"line_number":xxx,"change_type":xxx}... }
func findAllChangedLineNumbers(lines []string) []map[string]string {

	markCompile, err := regexp.Compile(`^[\+\-]`)
	errorHandler(err)
	changeLineNumbers := make([]map[string]string, 0)
	lineNumber := 0
	for _, line := range lines {
		lineNumber++
		res := markCompile.FindString(line)
		if res != "" {
			tempMap := map[string]string{
				"line_number": strconv.Itoa(lineNumber),
				"change_type": string(line[0]),
			}
			changeLineNumbers = append(changeLineNumbers, tempMap)
		}
	}
	return changeLineNumbers
}

// @title findAllChangedLines
// @description 匹配所有代码变更行并存储到切片，匹配的正则规范为：+/-行被匹配
// @auth KevinMatt
// @param lines []string diff中所有的代码行
// @return ChangedLineInfos []map[string]string 返回map的切片，格式类似：ChangedLineInfos={
//	{"lineNumber": 行号, "changeType": 变动类型},
//	{"lineNumber": 行号, "changeType": 变动类型},
//    ...
//}
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
			tempMap["lineNumber"] = strconv.Itoa(lineNumber)
			tempMap["changeType"] = string(line[0])
			ChangedLineInfos = append(ChangedLineInfos, tempMap)
		}
	}
	return ChangedLineInfos
}

// @title findAllDiffInfo
// @description 匹配并返回所有diff行内容
// @author KevinMatt
// @param lines string 抽取的原始git log数据
// @return 返回release_diff_map
func findAllDiffInfo(repoPath string, branceName string, projectId string) map[string]string {

	secret := os.Getenv("WHOSBUG_SECRET")
	originPath, err := os.Getwd()
	errorHandler(err)
	err = os.Chdir(repoPath)
	errorHandler(err)
	fmt.Println("Work path changed to:", repoPath)

	newReleaseCommitHash := execCommandOutput("git rev-parse HEAD")
	originHash := getLatestRelease(encrypt(projectId, secret, projectId))
	lastReleaseCommitHash := decrypt(projectId, secret, originHash)
	fmt.Println("last release's commit hash: ", lastReleaseCommitHash)
	fmt.Println("new release's commit hash: ", newReleaseCommitHash)

	var diff, commitInfo string
	if lastReleaseCommitHash != "" {
		diff = execCommandOutput(fmt.Sprintf("git log --full-diff -p -U1000 --pretty=raw %s..%s", lastReleaseCommitHash, newReleaseCommitHash))
		commitInfo = execCommandOutput(fmt.Sprintf("git log --pretty=format:%%H,%%ce,%%cn,%%cd %s..%s", lastReleaseCommitHash, newReleaseCommitHash))
	} else {
		diff = execCommandOutput("git log --full-diff -p -U1000 --pretty=raw'")
		commitInfo = execCommandOutput("git log --pretty=format:%H,%ce,%cn,%cd")
	}

	releaseDiff := make(map[string]string)
	releaseDiff["commit_info"] = commitInfo
	releaseDiff["diff"] = diff
	releaseDiff["branch_name"] = branceName
	releaseDiff["head_commit_id"] = newReleaseCommitHash

	// 返回原工作目录
	err = os.Chdir(originPath)
	fmt.Println("Work path changed back to:", originPath)
	errorHandler(err)

	return releaseDiff
}

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
