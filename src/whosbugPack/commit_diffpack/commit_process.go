package commit_diffpack

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"whosbugPack/global_type"
	"whosbugPack/uploadpack"
	"whosbugPack/utility"
)

// MatchCommit
//	@Description: 主体过程，最后直接生成结果集，位置在SourceCode下(此部分可做商榷)
//	@param diffPath diff-commit文件目录
//	@param commitPath commit-info文件目录
//	@author KevinMatt 2021-07-29 17:37:10
//	@function_mark PASS
func MatchCommit(diffPath, commitPath string) {

	err := uploadpack.PostCommitsInfo(commitPath)
	if err != nil {
		log.Println(utility.ErrorStack(err))
	}

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

			var commitInfo global_type.CommitInfoType

			commitInfo = utility.GetCommitInfo(string(commitLine))
			// 获取一次完整的commit，使用循环交错读取的方法避免跳过commit
			fullCommit, err := getFullCommit(patCommit, lineReaderDiff)
			if err != nil {
				fmt.Println(utility.ErrorStack(err))
			}

			// 获取单次commit中的每一次diff，并处理diff，送进协程
			ParseDiff(fullCommit, commitInfo)
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
