package whosbugAssigns

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

var monthCorrespond = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

// toIso8601
/* @Description: 时间戳转换
 * @param timeList
 * @return string
 * @author KevinMatt 2021-07-25 13:42:29
 * @function_mark PASS
 */
func toIso8601(timeList []string) string {
	return fmt.Sprintf("%s-%s-%sT%s%s:%s", timeList[3], monthCorrespond[timeList[0]], timeList[1], timeList[2], timeList[4][3:], timeList[4][3:])
}

// parseCommit
/* @Description: 解析commit信息
 * @param data 传入数据的diff部分(git log元数据)
 * @param commitInfos  log元数据分片
 * @return []CommitParsedType
 * @author KevinMatt 2021-07-25 13:36:35
 * @function_mark PASS
 */
func parseCommit(data string, commitInfos []string) []CommitParsedType {
	patCommit, err := regexp.Compile(`(commit\ ([a-f0-9]{40}))`)
	errorHandler(err)
	rawCommits := patCommit.FindAllStringSubmatch(data, -1)
	var parsedCommits []CommitParsedType
	for index, commitInfoLine := 0, commitInfos[0]; index < len(rawCommits) && index < len(commitInfos); index++ {
		commitInfoLine = commitInfos[index]
		infoList := strings.Split(commitInfoLine, ",")
		timeList := strings.Split(infoList[3][4:], " ")
		var parsedCommit CommitParsedType
		parsedCommit.CommitLeftIndex = patCommit.FindAllStringSubmatchIndex(data, -1)[index][0]
		parsedCommit.Commit = infoList[0]
		parsedCommit.CommitTime = toIso8601(timeList)
		parsedCommit.CommitterInfo.Name = infoList[2]
		parsedCommit.CommitterInfo.Email = infoList[1]
		parsedCommits = append(parsedCommits, parsedCommit)
	}
	return parsedCommits
}

// parseDiff
/* @Description: 将git log的信息的diff部分分解提取
 * @param data
 * @return []DiffParsedType
 * @author KevinMatt 2021-07-25 13:43:55
 * @function_mark PASS
 */
func parseDiff(data string) []DiffParsedType {
	patDiff, err := regexp.Compile(`(diff\ \-\-git\ a/(.*)\ b/.+)`)
	errorHandler(err)
	patDiffPart, err := regexp.Compile(`(@@\ .*?\ @@)`)
	errorHandler(err)
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)
	diffParsed := make([]DiffParsedType, 0)

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
		var changeLineNumbers []ChangeLineNumberType
		changeLineNumbers = findAllChangedLineNumbers(lines)
		lines = replaceLines(lines)
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
			var diffSingle DiffParsedType
			diffSingle.DiffFile = parts
			diffSingle.DiffFilePath = diffFilePath
			diffSingle.ChangeLineNumbers = append(diffSingle.ChangeLineNumbers, changeLineNumbers...)
			diffParsed = append(diffParsed, diffSingle)
		} else {
			continue
		}
	}
	return diffParsed
}

// replaceLines
/* @Description: 清除+/-符号并移除-行和No newline提示
 * @param lines 传入行集合
 * @return []string
 * @author KevinMatt 2021-07-25 13:52:57
 * @function_mark PASS
 */
func replaceLines(lines []string) []string {
	for index := 0; index < len(lines); index++ {
		if len(lines[index]) > 1 {
			if string(lines[index][0]) == "+" {
				lines[index] = "" + lines[index][1:]
				//strings.Replace(lines[index], string(lines[index][0]), "", 1)
			} else if string(lines[index][0]) == "-" || lines[index] == "\\ No newline at end of file\r\n" {
				lines[index] = ""
			} else {
				lines[index] = "" + lines[index][1:]
			}
		}
	}
	return lines
}
