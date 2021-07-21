package whosbugAssigns

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

var month_correspond = map[string]string{
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

func toIso8601(timeList []string) string {
	return fmt.Sprintf("%s-%s-%sT%s%s:%s", timeList[3], month_correspond[timeList[0]], timeList[1], timeList[2], timeList[4][3:], timeList[4][3:])
}

func parseCommit(data string, commitInfos []string) []map[string]interface{} {
	patCommit, err := regexp.Compile(`commit\ ([a-f0-9]{40})\n`)
	errorHandler(err)
	rawCommits := patCommit.FindAllStringSubmatch(data, -1)
	var parsedCommits []map[string]interface{}
	for index, commitInfoLine := 0, commitInfos[0]; index < len(rawCommits) && index < len(commitInfos); index++ {
		commitInfoLine = commitInfos[index]
		infoList := strings.Split(commitInfoLine, ",")
		timeList := strings.Split(infoList[0][4:], " ")
		parsedCommit := map[string]interface{}{
			"commit_left_index": patCommit.FindAllStringSubmatchIndex(data, -1)[index][0],
			"commit":            infoList[0],
			"commit_time":       toIso8601(timeList),
			"committer": map[string]string{
				"name":  infoList[2],
				"email": infoList[1],
			},
		}
		parsedCommits = append(parsedCommits, parsedCommit)
	}
	return parsedCommits
}

// @title parseDiff
// @description 将git log的信息的diff部分分解提取
// @param data string
// @author KevinMatt
func parseDiff(data string) []map[string]interface{} {
	patDiff, err := regexp.Compile(`diff\ \-\-git\ a/(.*)\ b/.+\n`)
	errorHandler(err)
	patDiffPart, err := regexp.Compile(`@@\ .*?\ @@\n`)
	errorHandler(err)
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)

	// 输出变量
	diffParsed := make([]map[string]interface{}, 0)
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

		if diffHeadMatch == nil {
			continue
		}
		// @@部分的左下标
		rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)[1]
		tempFileContent := diffPartsContent[rightDiffHeadIndex:]
		lines := strings.Split(tempFileContent, "\n")
		var changeLineNumbers interface{}
		changeLineNumbers = findAllChangedLineNumbers(lines)
		// 循环替换每一变动行的第一位
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

			if _, err := os.Stat(path.Dir(diffFilePath)); os.IsNotExist(err) {
				err := os.MkdirAll(path.Dir(diffFilePath), os.ModePerm)
				errorHandler(err)
			}
			fd, err := os.OpenFile(diffFilePath, os.O_RDWR, os.ModePerm)
			errorHandler(err)
			_, err = fd.WriteString(sourceCode)
			errorHandler(err)
			err = fd.Close()
			errorHandler(err)
			diffParsed[index]["diff_file"] = parts
			diffParsed[index]["diff_file_path"] = diffFilePath
			diffParsed[index]["change_line_numbers"] = changeLineNumbers
		}
	}
	return diffParsed
}
