package commit

import (
	"log"
	"path"
	"runtime"
	"strconv"
	"strings"

	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/util"

	"github.com/pkg/errors"
)

// ParseDiff
//	@Description: 将commit内的diff解析后存入SourceCode中
//	@param data 传入的fullCommit字符串
//	@param CommitHash 本次commit的Hash
//	@author KevinMatt 2021-07-29 22:54:33
//	@function_mark PASS
func ParseDiff(data string, commitInfo config.CommitInfoType) {
	// 匹配所有diffs及子匹配->匹配去除a/ or b/的纯目录
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)

	// 匹配diff行的index列表
	indexList := patDiff.FindAllStringIndex(data, -1)

	// 遍历所有diff
	for index, rawDiff := range rawDiffs {

		// 如果非匹配的语言文件，直接跳过
		// TODO: 还需要识别目标语言
		support, targetLanguage := languageFilter(path.Base(rawDiff[2]))
		if !support {
			runtime.GC()
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
			rightDiffHeadIndex := patDiffPart.FindStringSubmatchIndex(diffPartsContent)

			// 无有效匹配直接跳过
			if rightDiffHeadIndex == nil {
				continue
			}
			temp := strings.Split(diffPartsContent[rightDiffHeadIndex[4]:rightDiffHeadIndex[5]], " ")
			OldLineCount := QuatToNum(temp[0][1:])
			NewlineCount := QuatToNum(temp[1][1:])

			// 获取所有行，并按"\n"切分，略去第一行(@@行)
			lines := (strings.Split(diffPartsContent[rightDiffHeadIndex[1]:][0:], "\n"))[1:]

			// 传入行的切片，寻找所有变动行
			changeLineNumbers := findAllChangedLineNumbers(lines)

			// 替换 +/-行，删除-行内容，切片传递，无需返回值
			replaceLines(lines)
			util.ForDebug()
			// 填入到结构体中，准备送入协程
			var diffParsed = config.DiffParsedType{
				CommitterEmail: commitInfo.CommitterEmail,
				CommitTime:     commitInfo.CommitTime,
				CommitAuthor:   commitInfo.CommitAuthor,
				DiffFileName:   rawDiff[2],
				CommitHash:     commitInfo.CommitHash,
				DiffText:       strings.Join(lines, "\n"),
				OldLineCount:   OldLineCount,
				NewLineCount:   NewlineCount,
				TargetLanguage: targetLanguage,
			}
			diffParsed.ChangeLineNumbers = append(diffParsed.ChangeLineNumbers, changeLineNumbers...)
			// 上传任务到Antlr解析协程池
			//go func() {
			err := AntlrAnalysisPool.Invoke(diffParsed)
			if err != nil {
				log.Println(util.ErrorStack(errors.WithStack(err)))
			}
			//}()
		}
	}
}

func QuatToNum(text string) (sum int) {
	for index := 0; index < len(text); index++ {
		if text[index] == ',' {
			continue
		}
		temp, _ := strconv.Atoi(string(text[index]))
		sum = sum*10 + temp
	}
	return
}

// findAllChangedLineNumbers
//	@Description: 找到所有变动行号
//	@param lines 传入的行
//	@return []ChangeLineType 返回变动行信息结构体切片
//	@author KevinMatt 2021-07-29 19:48:01
//	@function_mark PASS
func findAllChangedLineNumbers(lines []string) (changeLineNumbers []config.ChangeLineType) {
	var lineNumber float64
	var recordLineNumber float64
	for index, line := range lines {
		lineNumber = float64(index + 1)
		res := markCompile.FindString(line)
		if res != "-" {
			// 若不是删去行，记录最新行号
			recordLineNumber = lineNumber
			if res == "+" {
				changeLine := config.ChangeLineType{
					LineNumber: recordLineNumber,
					ChangeType: res,
				}
				changeLineNumbers = append(changeLineNumbers, changeLine)
			}
		} else {
			// 若是删去行，记录为连续删去行前的最新行号
			changeLine := config.ChangeLineType{
				LineNumber: recordLineNumber + 0.5,
				ChangeType: res,
			}
			changeLineNumbers = append(changeLineNumbers, changeLine)
		}
	}
	return
}

// replaceLines
//	@Description: 替换处理传入的行
//	@param lines 传入的行切片
//	@author KevinMatt 2021-07-29 19:07:41
//	@function_mark PASS
func replaceLines(lines []string) {
	for index := 0; index < len(lines); index++ {
		if len(lines[index]) > 0 {
			if string(lines[index][0]) == "+" {
				lines[index] = util.ConCatStrings("", lines[index][1:])
			} else if string(lines[index][0]) == "-" || lines[index] == "\\ No newline at end of file" {
				if index == len(lines)-1 {
					lines = lines[:index]
				} else {
					lines = append(lines[:index], lines[index+1:]...)
				}
				index--
			}
		}
	}
}

// languageFilter
//	@Description: 语言过滤器，确定目标文件是否为支持的语言
//	@param fileName 文件名
//	@return bool 是否支持语言
//	@author KevinMatt 2021-07-26 20:48:57
//	@function_mark PASS
func languageFilter(fileName string) (support bool, targetLanguage string) {
	for suffix, language := range config.SupportLanguagesMap {
		if path.Ext(fileName) == suffix {
			support, targetLanguage = true, language
			return
		}
	}
	return
}
