package whosbugAssigns

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
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
	workPath, _ = os.Getwd()
	t := time.Now()
	secret = os.Getenv("WHOSBUG_SECRET")
	if secret == "" {
		secret = "defaultsecret"
	}
	//获取log
	releaseDiff := getLogInfo(repoPath, branchName, projectId)
	// 解析logCommit
	commitInfoList := parseCommit(releaseDiff.DiffPath, releaseDiff.CommitInfoPath)
	divideCommitDiff(commitInfoList)
	processMainMethod(commitInfoList, projectId)

	//processMain(commitInfoList)

	fmt.Println("Analysis cost: ", time.Since(t))
	//return parsedCommits
	return nil
}

func processACommit(commitInfo CommitInfoType) {

}
func processMainMethod(commitInfoList []CommitInfoType, projectId string) {
	fd, _ := os.Open(workPath + "\\full-res")
	fileReader := bufio.NewReader(fd)
	lineNumber := 1
	index := 0
	var commitFull []string
	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}
		if lineNumber >= commitInfoList[index].StartLineNumber && lineNumber < commitInfoList[index+1].StartLineNumber {
			commitFull = append(commitFull, string(line))
		} else if lineNumber == commitInfoList[index+1].StartLineNumber {
			commitDiffs := parseDiff(strings.Join(commitFull, "\n"))
			analyzeCommitDiff(projectId, commitDiffs, commitInfoList[index].commitHash)

			index++
		}
		lineNumber++
	}
}

func processMain(commitInfoList []CommitInfoType) {
	fd, _ := os.Open(workPath + "\\full-res")
	fileScanner := bufio.NewScanner(fd)

	for _, commitInfo := range commitInfoList {
		for index := range commitInfo.DiffInfoList {
			if index == len(commitInfo.DiffInfoList)-1 {
				continue
			}
			processDiff(fileScanner, commitInfo.DiffInfoList[index], commitInfo.DiffInfoList[index+1], commitInfo.commitHash)
		}
		flag = 1
	}
	fd.Close()
}

func processDiff(fileScanner *bufio.Scanner, diffInfo1, diffInfo2 DiffInfoType, commitHash string) {
	//data := withScanner(strings.NewReader(workPath+"\\full-res"), int64(diffInfo1.diffHeadLineNumber), int64(diffInfo2.diffHeadLineNumber))
	data := readFileByLineNumber(fileScanner, diffInfo1.diffHeadLineNumber, diffInfo2.diffHeadLineNumber)
	fd, _ := os.OpenFile("SourceCode\\"+commitHash+diffInfo1.DiffFileName, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	_, err := fd.WriteString(strings.Join(data, "\n"))
	if err != nil {
		log.Println(err)
	}
}
func divideCommitDiff(commitInfoList []CommitInfoType) {
	patDiff, _ := regexp.Compile(`(diff\ \-\-git\ a/(.*)\ b/.+)`)
	getDiffListPerCommit1(commitInfoList, patDiff)
	return
}

var flag int = 1

func getDiffListPerCommit1(commitInfoList []CommitInfoType, patDiff *regexp.Regexp) {
	fd, _ := os.Open(workPath + "\\full-res")
	fileScanner := bufio.NewScanner(fd)
	lineNumber := 1
	index := 0

	for fileScanner.Scan() && index < len(commitInfoList)-1 {
		if lineNumber >= commitInfoList[index].StartLineNumber && lineNumber < commitInfoList[index+1].StartLineNumber {
			diffRes := patDiff.FindStringSubmatch(fileScanner.Text())
			if diffRes != nil {
				var diffInfo DiffInfoType
				diffInfo.StartLineNumber = lineNumber
				diffInfo.diffHeadLineNumber = lineNumber + 5
				diffInfo.DiffFilePath = diffRes[2]
				diffInfo.DiffFileName = path.Base(diffRes[2])
				commitInfoList[index].DiffInfoList = append(commitInfoList[index].DiffInfoList, diffInfo)
			}
		} else if lineNumber >= commitInfoList[index+1].StartLineNumber {
			index++
		}
		lineNumber++
	}
}

func readFileByLineNumber(fileScanner *bufio.Scanner, lineNumberStart int, lineNumberEnd int) []string {
	var res []string
	lineCount := 0
	if flag == 1 {
		lineCount = lineNumberStart
		flag = 0
	}
	for fileScanner.Scan() {
		if lineCount <= lineNumberEnd-lineNumberStart && flag == 0 {
			res = append(res, fileScanner.Text())
		} else {
			break
		}
		lineCount++
	}
	return res
}
