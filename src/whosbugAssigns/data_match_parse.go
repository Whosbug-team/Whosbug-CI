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

const parCommitPattern = `(commit\ ([a-f0-9]{40}))`
const parDiffPattern = `(diff\ \-\-git\ a/(.*)\ b/.+)`
const parDiffPartPattern = `(@@\ .*?\ @@)`

/* parseDiff
/* @Description: 将git log的信息的diff部分分解提取
 * @param data
 * @return []DiffParsedType
 * @author KevinMatt 2021-07-26 21:32:28
 * @function_mark PASS
*/
func parseDiff(data string) []DiffParsedType {
	t := time.Now()
	patDiff, _ := regexp.Compile(`(diff\ \-\-git\ a/(.*)\ b/.+)`)
	patDiffPart, _ := regexp.Compile(`(@@\ .*?\ @@)`)
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)
	diffParsedList := make([]DiffParsedType, 0)
	indexList := patDiff.FindAllStringIndex(data, -1)

	for index, rawDiff := range rawDiffs {
		leftDiffIndex := indexList[index][0]
		var diffPartsContent string
		var rightDiffIndex int
		if index == len(rawDiffs)-1 {
			diffPartsContent = data[leftDiffIndex:]
		} else {
			rightDiffIndex = (indexList[index+1])[0]
			diffPartsContent = data[leftDiffIndex:rightDiffIndex]
		}

		diffHeadMatch := patDiffPart.FindAllString(diffPartsContent, -1)
		if diffHeadMatch == nil {
			continue
		}

		rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)[1]

		tempFileContent := diffPartsContent[rightDiffHeadIndex:]

		lines := (strings.SplitAfter(tempFileContent[0:], "\n"))[1:]

		var changeLineNumbers []ChangeLineNumberType
		changeLineNumbers = findAllChangedLineNumbers(lines)
		lines = replaceLines(lines)
		sourceCode := strings.Join(lines, "")
		fileName := path.Base(rawDiff[2])

		if lanFilter(fileName) {
			commitDicName := data[7:17]
			diffFilePath := fmt.Sprintf("SourceCode/%s/%s", commitDicName, fileName)

			if _, err := os.Stat(path.Dir(diffFilePath)); os.IsNotExist(err) {
				err = os.MkdirAll(path.Dir(diffFilePath), os.ModePerm)
				errorHandler(err, "mkdir ", diffFilePath)
			}

			fd, err := os.OpenFile(diffFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			errorHandler(err, "open ", diffFilePath)
			_, err = fd.WriteString(sourceCode)
			errorHandler(err, "write ", sourceCode)
			err = fd.Close()
			errorHandler(err, "close fd ", sourceCode)
			var diffParsed DiffParsedType
			diffParsed.DiffFile = rawDiff[2]
			diffParsed.DiffFilePath = diffFilePath
			diffParsed.ChangeLineNumbers = append(diffParsed.ChangeLineNumbers, changeLineNumbers...)
			diffParsedList = append(diffParsedList, diffParsed)
		} else {
			continue
		}
	}
	fmt.Println("parsDiff cost ", time.Since(t))
	return diffParsedList
}

/* parseCommit
/* @Description: 解析commit信息
 * @param data 传入数据的diff部分(git log元数据)
 * @param commitInfos  log元数据分片
 * @return []CommitParsedType
 * @author KevinMatt 2021-07-26 20:52:13
 * @function_mark PASS
*/
//func parseCommit1(data string, commitInfos []string) []CommitParsedType {
//	t := time.Now()
//	patCommit, _ := regexp.Compile(parCommitPattern)
//	rawCommits := patCommit.FindAllStringSubmatch(data, -1)
//	indexList := patCommit.FindAllStringSubmatchIndex(data, -1)
//	var parsedCommitList []CommitParsedType
//	for index, commitInfoLine := 0, commitInfos[0]; index < len(rawCommits) && index < len(commitInfos); index++ {
//		commitInfoLine = commitInfos[index]
//		infoList := strings.Split(commitInfoLine, ",")
//		timeList := strings.Split(infoList[3][4:], " ")
//		var parsedCommit CommitParsedType
//		parsedCommit.CommitLeftIndex = indexList[index][0]
//		parsedCommit.Commit = infoList[0]
//		parsedCommit.CommitTime = toIso8601(timeList)
//		parsedCommit.CommitterInfo.Name = infoList[2]
//		parsedCommit.CommitterInfo.Email = infoList[1]
//		parsedCommitList = append(parsedCommitList, parsedCommit)
//	}
//	fmt.Println("parse COM cost ", time.Since(t))
//	return parsedCommitList
//}

/* parseCommit
/* @Description: 解析commit信息
 * @param diffInfoPath
 * @param commitInfoPath
 * @return []CommitInfoType
 * @author KevinMatt 2021-07-28 14:12:06
 * @function_mark PASS
*/
func parseCommit(diffInfoPath string, commitInfoPath string) []CommitInfoType {
	t := time.Now()
	patCommit, _ := regexp.Compile(parCommitPattern)
	// 读diff
	diffFileHandle, err := os.Open(diffInfoPath)
	commitFileHandle, err := os.Open(commitInfoPath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer diffFileHandle.Close()
	defer commitFileHandle.Close()
	lineReaderCommit := bufio.NewReader(commitFileHandle)
	lineReaderDiff := bufio.NewReader(diffFileHandle)
	lineNumber := 1
	var commitInfoList []CommitInfoType
	for {
		line, _, err := lineReaderDiff.ReadLine()
		if err == io.EOF {
			break
		}
		res := patCommit.FindString(string(line))
		if res != "" {
			// 匹配到一个commit，从commitinfo读一行
			commitLine, _, err := lineReaderCommit.ReadLine()
			if err == io.EOF {
				break
			}
			var commitInfo CommitInfoType
			splitInfo := strings.Split(string(commitLine), ",")
			if splitInfo == nil {
				continue
			}
			commitInfo.StartLineNumber = lineNumber
			commitInfo.commitHash = splitInfo[0]
			commitInfo.CommitterEmail = splitInfo[1]
			commitInfo.committerName = splitInfo[2]
			commitInfo.CommitTime = toIso8601(strings.Split(splitInfo[3][4:], " "))
			commitInfoList = append(commitInfoList, commitInfo)
		}
		lineNumber++
	}
	fmt.Println("time cost: ", time.Since(t))
	return commitInfoList
}

/* findAllChangedLineNumbers
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
	for index, line := range lines {
		lineNumber = index + 1
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
