package antlr

import (
	"regexp"
	"strings"

	"git.woa.com/bkdevops/whosbug-ci/internal/util"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/crypto"
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

	// 编译正则
	patDiff, _     = regexp.Compile(parDiffPattern)
	patDiffPart, _ = regexp.Compile(parDiffPartPattern)
	markCompile, _ = regexp.Compile(markPattern)
)

// CommitInfoType 存储每一次commit的信息
type CommitInfoType struct {
	CommitHash     string `json:"hash"`
	CommitterEmail string `json:"email"`
	CommitAuthor   string `json:"author"`
	CommitTime     string `json:"time"`
}

// GetCommitInfo 获取commit信息
//
//	@param line commitInfo行
//	@return config.CommitInfoType 返回结构体
//	@author KevinMatt 2021-08-10 01:04:21
//	@function_mark PASS
func GetCommitInfo(line string) CommitInfoType {
	infoList := strings.Split(line, ",")
	var tempCommitInfo = CommitInfoType{
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
