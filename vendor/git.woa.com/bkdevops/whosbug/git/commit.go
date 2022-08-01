package git

import (
	"strings"

	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/crypto"
	"git.woa.com/bkdevops/whosbug/util"
)

// GetCommitInfo 获取commit信息
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
