package analyze

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"

	"git.woa.com/bkdevops/whosbug-ci/internal/util"
	"git.woa.com/bkdevops/whosbug-ci/internal/zaplog"
	"github.com/pkg/errors"
)

// Diff 解析后的diff信息
type Diff struct {
	CommitterEmail string
	CommitTime     string
	CommitAuthor   string
	DiffFileName   string
	ChangeLines    []ChangeLine
	CommitHash     string
	DiffText       string
	OldLineCount   int
	NewLineCount   int
	TargetLanguage string
}

// ChangeLine 存储单个改变行的信息
type ChangeLine struct {
	// +0.5用于标识移除行位置(避免移出对象范围)
	LineNumber float64
	ChangeType string
}

// SupportLanguagesMap 支持的语言源码文件后缀到语言名的映射
var SupportLanguagesMap = map[string]string{
	// C源码后缀
	".C": "c",
	".c": "c",
	// C++源码后缀
	".cpp": "cpp",
	".cc":  "cpp",
	".cxx": "cpp",
	".hh":  "cpp",
	// Golang源码后缀
	".go": "golang",
	// Java源码后缀
	".java": "java",
	// JavaScript源码后缀
	".js": "javascript",
	// Kotlin源码后缀
	".kt": "kotlin",
}

// MatchCommit git解析commits的主体过程，最后直接生成结果集，位置在SourceCode下(此部分可做商榷)
// TODO: 优化为go-git
//
//	@param diffPath diff-commit文件目录
//	@param commitPath commit-info文件目录
//	@author KevinMatt 2021-07-29 17:37:10
//	@function_mark PASS
func MatchCommit(diffPath, commitPath string) {
	commitFd, err := os.Open(commitPath)
	if err != nil {
		log.Println("OpenFile Error: ", err)
		os.Exit(1)
	}
	diffFd, err := os.Open(diffPath)
	if err != nil {
		log.Println("OpenFile Error: ", err)
		os.Exit(1)
	}
	lineReaderCommit := bufio.NewReader(commitFd)
	lineReaderDiff := bufio.NewReader(diffFd)
	for {
		line, _, err := lineReaderDiff.ReadLine()
		if err == io.EOF {
			break
		}

		// 匹配tree行
		res := patTree.FindString(string(line))
		if res != "" {
			// 匹配到一个commit的tree行，从commit info读一行
			commitLine, _, err := lineReaderCommit.ReadLine()
			if err == io.EOF {
				break
			}

			commitInfo := GetCommitInfo(string(commitLine))
			// 获取一次完整的commit，使用循环交错读取的方法避免跳过commit
			fullCommit, err := getFullCommit(patCommit, lineReaderDiff)
			if err != nil {
				zaplog.Logger.Error("[MatchCommit] error when parse full commit", zaplog.Error(err))
			}

			// 获取单次commit中的每一次diff，并处理diff，送进协程
			parseDiff(fullCommit, commitInfo)
			ProcessBar.Add(1)
			runtime.GC()
		}
	}
	err = commitFd.Close()
	if err != nil {
		log.Println(errors.WithStack(err))
	}
	err = diffFd.Close()
	if err != nil {
		log.Println(errors.WithStack(err))
	}
}

// getFullCommit
//
//	@Description: 交错读取commit-diff文件
//	@param patCommit 预编译的正则表达式
//	@param lineReaderDiff 全局共享fd
//	@return string 返回完整的commit串
//	@return string 错误信息
//	@author KevinMatt 2021-07-29 17:52:58
//	@function_mark PASS
func getFullCommit(patCommit *regexp.Regexp, lineReaderDiff *bufio.Reader) (string, error) {
	var lines []string
	for {
		line, _, err := lineReaderDiff.ReadLine()
		if err == io.EOF {
			break
		}
		// 匹配commit行，交错读取
		res := patCommit.FindString(string(line))
		if res != "" {
			break
		}
		lines = append(lines, string(line))
	}
	return strings.Join(lines, "\n"), nil
}

// parseDiff 将commit内的diff解析后存入SourceCode中
//
//	@param data string
//	@param commitInfo CommitInfoType
//	@author kevineluo
//	@update 2023-02-27 11:10:20
func parseDiff(data string, commitInfo CommitInfoType) {
	// 匹配所有diffs及子匹配->匹配去除a/ or b/的纯目录
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)

	// 匹配diff行的index列表
	indexList := patDiff.FindAllStringIndex(data, -1)

	// 遍历所有diff
	for index, rawDiff := range rawDiffs {

		// 如果非匹配的语言文件，直接跳过
		isSupportLanguage, targetLanguage := languageFilter(path.Base(rawDiff[2]))
		if !isSupportLanguage {
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
			OldLineCount := util.QuatToNum(temp[0][1:])
			NewlineCount := util.QuatToNum(temp[1][1:])

			// 获取所有行，并按"\n"切分，略去第一行(@@行)
			lines := (strings.Split(diffPartsContent[rightDiffHeadIndex[1]:][0:], "\n"))[1:]

			// 传入行的切片，寻找所有变动行
			changeLineNumbers := findAllChangedLineNumbers(lines)

			// 替换 +/-行，删除-行内容，切片传递，无需返回值
			replaceLines(lines)
			util.ForDebug()
			// 填入到结构体中，准备送入协程
			var diffParsed = Diff{
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
			diffParsed.ChangeLines = append(diffParsed.ChangeLines, changeLineNumbers...)
			// 上传任务到Antlr解析协程池
			err := AntlrAnalysisPool.Invoke(diffParsed)
			if err != nil {
				zaplog.Logger.Error("[ParseDiff] invoke antlr parse task error", zaplog.Error(err))
			}
		}
	}
}

// findAllChangedLineNumbers 找到所有变动行号
//
//	@param lines []string
//	@return changeLineNumbers []config.ChangeLineType
//	@author kevineluo
//	@update 2023-02-27 11:10:16
func findAllChangedLineNumbers(lines []string) (changeLineNumbers []ChangeLine) {
	var lineNumber float64
	var recordLineNumber float64
	for index, line := range lines {
		lineNumber = float64(index + 1)
		res := markCompile.FindString(line)
		if res != "-" {
			// 若不是删去行，记录最新行号
			recordLineNumber = lineNumber
			if res == "+" {
				changeLine := ChangeLine{
					LineNumber: recordLineNumber,
					ChangeType: res,
				}
				changeLineNumbers = append(changeLineNumbers, changeLine)
			}
		} else {
			// 若是删去行，记录为连续删去行前的最新行号
			changeLine := ChangeLine{
				LineNumber: recordLineNumber + 0.5,
				ChangeType: res,
			}
			changeLineNumbers = append(changeLineNumbers, changeLine)
		}
	}
	return
}

// replaceLines
//
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

// languageFilter 语言过滤器，确定目标文件是否为支持的语言
//
//	@param fileName string
//	@return support bool
//	@return targetLanguage string
//	@author kevineluo
//	@update 2023-02-27 11:10:26
func languageFilter(fileName string) (support bool, targetLanguage string) {
	for suffix, language := range SupportLanguagesMap {
		if path.Ext(fileName) == suffix {
			support, targetLanguage = true, language
			return
		}
	}
	return
}
