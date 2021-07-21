package whosbugAssigns

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

// @title parseDiff
// @description 将git log的信息的diff部分分解提取
// @param data string
// @author KevinMatt
func parseDiff(data string) []map[string]string {

	patDiff, err := regexp.Compile(`diff\ \-\-git\ a/(.*)\ b/.+\n`)
	errorHandler(err)
	patDiffPart, err := regexp.Compile(`@@\ .*?\ @@\n`)
	errorHandler(err)
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)
	for index := 0; index < len(rawDiffs); index++ {
		// 正则匹配的三维结果
		rawCommit := rawDiffs[index]
		// 完整的整行匹配
		fullCommit := rawCommit[0]
		// 子匹配(sub_match)
		parts := rawCommit[2]
		leftDiffIndex := patDiff.FindStringIndex(fullCommit)[0]
		var diffPartsContent string
		var rightDiffIndex int
		if index == len(rawDiffs)-1 {
			diffPartsContent = data[leftDiffIndex:]
		} else {
			rightDiffIndex = patDiff.FindStringIndex(rawDiffs[index+1][0])[0]
			diffPartsContent = data[leftDiffIndex:rightDiffIndex]
		}
		diffHeadMatch := patDiffPart.FindAllString(diffPartsContent, -1)
		var diffHead string
		if diffHeadMatch != nil {
			// @@部分
			diffHead = diffHeadMatch[0]
		} else {
			continue
		}
		rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)[1]
		tempFileContent := diffPartsContent[rightDiffHeadIndex:]
		lines := strings.Split(tempFileContent, "\n")
		changeLineNumbers := findAllChangedLineNumbers(lines)
		// 循环扣去每一行的第一位
		for index := 0; index < len(lines); index++ {
			//strings.Replace(lines[index], string(lines[index][0]), " ", 1)
			lines[index] = " " + lines[index][1:]
		}
		sourceCode := strings.Join(lines, "\n")
		fileName := path.Base(parts)

		if lanFilter(fileName) {
			fmt.Println(fileName)
			commitDicName := data[7:17]
			diffFilePath := fmt.Sprintf("SourceCode/%s/%s", commitDicName, fileName)
			if
		}

	}
}

//func parseDiff(data string) []map[string]string {
//	patDiff, err := regexp.Compile(`diff\ \-\-git\ a/(.*)\ b/.+\n`)
//	errorHandler(err)
//	patDiffPart, err := regexp.Compile(`@@\ .*?\ @@\n`)
//	errorHandler(err)
//	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)
//	for index := 0; index < len(rawDiffs); index++ {
//		// 正则的三维匹配结果
//		rawCommit := rawDiffs[index]
//		// 完整的父匹配
//		fullCommit := rawCommit[0]
//		// 子匹配
//		parts := rawDiffs[index][2]
//		leftDiffIndex := patDiff.FindStringIndex(fullCommit)[0]
//		var diffPartsContent string
//		var rightDiffIndex int
//		if index == len(rawDiffs)-1 {
//			diffPartsContent = data[leftDiffIndex:]
//		} else {
//			rightDiffIndex = patDiff.FindStringIndex(rawDiffs[index+1][0])[0]
//			diffPartsContent = data[leftDiffIndex:rightDiffIndex]
//		}
//		diffHeadMatch := patDiffPart.FindAllString(diffPartsContent, -1)
//		var diffHead string
//		if diffHeadMatch != nil {
//			// @@信息
//			diffHead = diffHeadMatch[0]
//		} else {
//			continue
//		}
//		// diff头部信息右侧下标
//		rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)[1]
//		tempFileContent := diffPartsContent[rightDiffHeadIndex:]
//		lines := strings.Split(tempFileContent, "\n")
//		changeLineNumbers := findAllChangedLines(lines)
//		lines =
//
//	}
