package commit_diffpack

import (
	"github.com/pkg/errors"
	"log"
	"path"
	"strings"
	"whosbugPack/global_type"
	"whosbugPack/utility"
)

/* ParseDiffToFile
/* @Description: 将commit内的diff解析后存入SourceCode中
 * @param data 传入的fullCommit字符串
 * @param CommitHash 本次commit的Hash
 * @author KevinMatt 2021-07-29 22:54:33
 * @function_mark PASS
*/
func ParseDiffToFile(data string, commitInfo global_type.CommitInfoType) {
	// 匹配所有diffs及子匹配->匹配去除a/ or b/的纯目录
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)

	// 匹配diff行的index列表
	indexList := patDiff.FindAllStringIndex(data, -1)

	// 遍历所有diff
	for index, rawDiff := range rawDiffs {

		// 如果非匹配的语言文件，直接跳过
		if !lanFilter(path.Base(rawDiff[2])) {
			continue
		} else {
			// 获得左索引
			leftDiffIndex := indexList[index][0]

			var diffPartsContent string
			var rightDiffIndex int
			// 判断是否为最后一项diff，随后获取代码段
			if index == len(rawDiffs)-1 {
				diffPartsContent = data[leftDiffIndex:]
			} else {
				rightDiffIndex = (indexList[index+1])[0]
				diffPartsContent = data[leftDiffIndex:rightDiffIndex]
			}

			// 匹配@@行
			rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)

			// 无有效匹配直接跳过
			if rightDiffHeadIndex == nil {
				continue
			}

			// 获取所有行，并按"\n"切分，略去第一行(@@行)
			lines := (strings.Split(diffPartsContent[rightDiffHeadIndex[1]:][0:], "\n"))[1:]

			// 传入行的切片，寻找所有变动行
			changeLineNumbers := findAllChangedLineNumbers(lines)

			// 替换 +/-行，删除-行内容，切片传递，无需返回值
			replaceLines(lines)

			// 填入到结构体中，准备送入协程
			var diffParsed global_type.DiffParsedType
			diffParsed.DiffText = strings.Join(lines, "\n")
			diffParsed.DiffFileName = rawDiff[2]
			diffParsed.ChangeLineNumbers = append(diffParsed.ChangeLineNumbers, changeLineNumbers...)
			diffParsed.CommitHash = commitInfo.CommitHash
			diffParsed.CommitterName = commitInfo.CommitterName
			diffParsed.CommitTime = commitInfo.CommitTime
			diffParsed.CommitterEmail = commitInfo.CommitterEmail
			// 得到单个diff后直接送入analyze进行分析
			//fmt.Println("pool running: ", pool.Running())
			// 上传任务到协程池
			go func() {
				err := Pool.Invoke(diffParsed)
				if err != nil {
					log.Println(utility.ErrorStack(errors.WithStack(err)))
				}
			}()
		}
	}
}

/* findAllChangedLineNumbers
/* @Description: 找到所有变动行号
 * @param lines 传入的行
 * @return []ChangeLineType 返回变动行信息结构体切片
 * @author KevinMatt 2021-07-29 19:48:01
 * @function_mark PASS
*/
func findAllChangedLineNumbers(lines []string) []global_type.ChangeLineType {
	var changeLineNumbers []global_type.ChangeLineType
	lineNumber := 0
	for index, line := range lines {
		lineNumber = index + 1
		res := markCompile.FindString(line)
		if res != "" {
			var tempStruct global_type.ChangeLineType
			tempStruct.LineNumber = lineNumber
			tempStruct.ChangeType = string(line[0])
			changeLineNumbers = append(changeLineNumbers, tempStruct)
		}
	}
	return changeLineNumbers
}

/* replaceLines
/* @Description: 替换处理传入的行
 * @param lines 传入的行切片
 * @author KevinMatt 2021-07-29 19:07:41
 * @function_mark PASS
*/
func replaceLines(lines []string) {
	for index := range lines {
		if len(lines[index]) >= 1 {
			if string(lines[index][0]) == "+" {
				lines[index] = utility.ConCatStrings("", lines[index][1:])
			} else if string(lines[index][0]) == "-" || lines[index] == "\\ No newline at end of file" {
				lines[index] = ""
			} else {
				lines[index] = utility.ConCatStrings("", lines[index][1:])
			}
		}
	}
}

/* lanFilter
/* @Description: 语言过滤器，确定目标文件是否为支持的语言
 * @param fileName 文件名
 * @return bool 是否支持语言
 * @author KevinMatt 2021-07-26 20:48:57
 * @function_mark PASS
*/
func lanFilter(fileName string) bool {
	for index := range global_type.SupportLans {
		if path.Ext(fileName) == global_type.SupportLans[index] {
			return true
		}
	}
	return false
}
