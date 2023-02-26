package git

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"

	"git.woa.com/bkdevops/whosbug-ci/internal/util"
	"git.woa.com/bkdevops/whosbug-ci/internal/zaplog"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/config"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/crypto"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/logging"
	"github.com/panjf2000/ants"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

const (
	// parCommitPattern 匹配commit行
	parCommitPattern = `(commit\ ([a-f0-9]{40}))`
	// parTreePattern 匹配tree行，用于`交错匹配`
	parTreePattern = `(tree\ ([a-f0-9]{40}))`
	// parDiffPattern 匹配diff行，用于每一次commit信息的处理
	parDiffPattern = `(diff\ \-\-git\ a/(.*)\ b/.+)`
	// parDiffPartPattern 匹配diff段的末行@@行，用于获取diff代码内容的起始位置
	parDiffPartPattern = `(@@\ (.*?)\ @@)`
	// markPattern 匹配+/-变动行
	markPattern = `^[\+\-]`
)

var (
	// patCommit, patTree 预编译正则匹配
	patCommit, _ = regexp.Compile(parCommitPattern)
	patTree, _   = regexp.Compile(parTreePattern)

	// ProcessBar 显示进度条
	ProcessBar *progressbar.ProgressBar

	// AntlrAnalysisPool 解析协程池
	AntlrAnalysisPool, _ = ants.NewPoolWithFunc(runtime.NumCPU()/4, func(commitDiff interface{}) {
		antlr.AnalyzeCommitDiff(commitDiff.(config.DiffParsedType))
		// 指示已经处理的diff数量
		processDiffs++
		log.SetOutput(logging.LogFile)
		log.Println("Diff No.", processDiffs, " From", commitDiff.(config.DiffParsedType).CommitHash, " Sent Into Channel.")
		// if processDiffs%50 == 0 {
		// runtime.GC()
		// }
	})

	// processDiffs 已处理的commit数
	processDiffs int

	// 编译正则
	patDiff, _     = regexp.Compile(parDiffPattern)
	patDiffPart, _ = regexp.Compile(parDiffPartPattern)
	markCompile, _ = regexp.Compile(markPattern)
)

// GetCommitInfo 获取commit信息
//
//	@param line commitInfo行
//	@return config.CommitInfoType 返回结构体
//	@author KevinMatt 2021-08-10 01:04:21
//	@function_mark PASS
func GetCommitInfo(line string) config.CommitInfoType {
	infoList := strings.Split(line, ",")
	var tempCommitInfo = config.CommitInfoType{
		CommitHash:     crypto.Base64Encrypt(infoList[0]),
		CommitterEmail: crypto.Base64Encrypt(infoList[1]),
		CommitTime:     crypto.Base64Encrypt(util.ToIso8601(strings.Split(infoList[len(infoList)-1][4:], " "))),
	}
	// 赋值commitAuthor(考虑多个Author的可能)
	for index := 2; index < len(infoList)-1; index++ {
		tempCommitInfo.CommitAuthor += infoList[index]
		if index != len(infoList)-2 {
			tempCommitInfo.CommitAuthor = util.ConCatStrings(tempCommitInfo.CommitAuthor, ",")
		}
	}
	tempCommitInfo.CommitAuthor = crypto.Base64Encrypt(tempCommitInfo.CommitAuthor)
	return tempCommitInfo
}

// MatchCommit 主体过程，最后直接生成结果集，位置在SourceCode下(此部分可做商榷)
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
			ParseDiff(fullCommit, commitInfo)
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
