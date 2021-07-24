package whosbugAssigns

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

type changeLineNumberType struct {
	lineNumber int
	changeType string
}
type diffParsedType struct {
	diffFile          string
	diffFilePath      string
	changeLineNumbers []changeLineNumber
}

func parseDiff1(data string) []diffParsedType {
	patDiff, err := regexp.Compile(`(diff\ \-\-git\ a/(.*)\ b/.+)`)
	errorHandler(err)
	patDiffPart, err := regexp.Compile(`(@@\ .*?\ @@)`)
	errorHandler(err)
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)
	diffParsed := make([]diffParsedType, len(rawDiffs))

	for index, rawCommit := range rawDiffs {
		parts := rawCommit[2]
		leftDiffIndex := patDiff.FindAllStringIndex(data, -1)[index][0]
		var diffPartsContent string
		var rightDiffIndex int
		if index == len(rawDiffs)-1 {
			diffPartsContent = data[leftDiffIndex:]
		} else {
			rightDiffIndex = (patDiff.FindAllStringIndex(data, -1)[index+1])[0]
			diffPartsContent = data[leftDiffIndex:rightDiffIndex]
		}
		diffHeadMatch := patDiffPart.FindAllString(diffPartsContent, -1)

		if diffHeadMatch == nil {
			continue
		}
		rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)[1]
		tempFileContent := diffPartsContent[rightDiffHeadIndex:]
		lines := (strings.SplitAfter(tempFileContent[0:], "\n"))[1:]
		var changLineNumbers changeLineNumberType
		changeLineNumbers = findAllChangedLineNumbers1(lines)
		lines := replaceLines(lines)
		sourceCode := strings.Join(lines, "")
		fileName := path.Base(parts)

		if lanFilter(fileName) {
			commitDicName := data[7:17]
			diffFilePath := fmt.Sprintf("SourceCode/%s/%s", commitDicName, fileName)

			if _, err := os.Stat(path.Dir(diffFilePath)); os.IsNotExist(err) {
				err := os.MkdirAll(path.Dir(diffFilePath), os.ModePerm)
				errorHandler(err)
			}
			fd, err := os.OpenFile(diffFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			errorHandler(err)
			_, err = fd.WriteString(sourceCode)
			errorHandler(err)
			err = fd.Close()
			errorHandler(err)
			diffParsed[index].diffFile = parts
			diffParsed[index].diffFilePath = diffFilePath
			diffParsed[index].changeLineNumbers = append(diffParsed[index].changeLineNumbers, changLineNumbers)
		}
	}
	return diffParsed
}

func findAllChangedLineNumbers1(lines []string) []changeLineNumberType {
	markCompile, err := regexp.Compile(`^[\+\-]`)
	errorHandler(err)
	changeLineNumbers := make([]changeLineNumberType, 0)
	lineNumber := 0
	for _, line := range lines {
		lineNumber++
		res := markCompile.FindString(line)
		if res != "" {
			var tempStruct changeLineNumberType
			tempStruct.lineNumber = lineNumber
			tempStruct.changeType = string(line[0])
			changeLineNumbers = append(changeLineNumbers, tempStruct)
		}
	}
	return changeLineNumbers
}
