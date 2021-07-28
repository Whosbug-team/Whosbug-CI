package whosbugAssigns

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
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
	// 线性读获取所有的diff起始行
	divideCommitDiff(commitInfoList)
	// diff送入协程
	processMain(commitInfoList)

	//var commits []CommitParsedType
	//var parsedCommits []CommitParsedType
	//for index := range commits {
	//	commit := commits[index]
	//	var diffPark string
	//	if index == len(commits)-1 {
	//		diffPark = releaseDiff.DiffPath[commit.CommitLeftIndex:]
	//	} else {
	//		nextCommitLeftIndex := commits[index+1].CommitLeftIndex
	//		diffPark = releaseDiff.DiffPath[commit.CommitLeftIndex:nextCommitLeftIndex]
	//	}
	//	// 解析diff信息
	//	commitDiffs := parseDiff(diffPark)
	//	// commitDiffs切片引用传递，不需返回值
	//	analyzeCommitDiff(projectId, commitDiffs, commit.Commit)
	//	commit.CommitDiffs = append(commit.CommitDiffs, commitDiffs...)
	//	parsedCommits = append(parsedCommits, commit)
	//}
	fmt.Println("Analysis cost: ", time.Since(t))
	//return parsedCommits
	return nil
}

func processMain(commitInfoList []CommitInfoType) {
	for _, commitInfo := range commitInfoList {
		for index := range commitInfo.DiffInfoList {
			if index == len(commitInfo.DiffInfoList)-1 {
				continue
			}
			processDiff(commitInfo.DiffInfoList[index], commitInfo.DiffInfoList[index+1])
		}
	}
}
func processDiff(diffInfo1, diffInfo2 DiffInfoType) {
	data := readFileByLineNumber(workPath+"\\full-res", diffInfo1.diffHeadLineNumber, diffInfo2.diffHeadLineNumber)
	fmt.Println(data)
	return
}
func divideCommitDiff(commitInfoList []CommitInfoType) {
	patDiff, _ := regexp.Compile(`(diff\ \-\-git\ a/(.*)\ b/.+)`)
	getDiffListPerCommit1(commitInfoList, patDiff)
	return
}

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

//func getDiffListPerCommit(StartLineNumber int, EndLineNumber int, patDiff, patDiffPart *regexp.Regexp) []DiffInfoType {
//	fd, _ := os.Open(workPath + "\\full-res")
//	fileScanner := bufio.NewScanner(fd)
//	lineNumber := 1
//	var diffInfoList []DiffInfoType
//	for fileScanner.Scan() {
//		if lineNumber >= StartLineNumber && lineNumber < EndLineNumber {
//			var diffInfo DiffInfoType
//			diffRes := patDiff.FindStringSubmatch(fileScanner.Text())
//			if diffRes != nil {
//				diffInfo.StartLineNumber = lineNumber
//				diffInfo.diffHeadLineNumber = lineNumber + 5
//				diffInfo.DiffFilePath = diffRes[2]
//				diffInfo.DiffFileName = path.Base(diffRes[2])
//				diffInfoList = append(diffInfoList, diffInfo)
//			}
//		} else if lineNumber >= EndLineNumber {
//			break
//		}
//		lineNumber++
//	}
//	return diffInfoList
//}
func readFileByLineNumber(filePath string, lineNumberStart int, lineNumberEnd int) [][]byte {
	var res [][]byte
	fd, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(fd)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount >= lineNumberStart && lineCount < lineNumberEnd {
			res = append(res, fileScanner.Bytes())
		} else if lineCount >= lineNumberEnd {
			break
		}
		lineCount++
	}
	defer fd.Close()
	return res
}
