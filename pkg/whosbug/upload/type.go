// Package upload 解析结果上报相关 package
//
//	@update 2023-02-27 11:16:53
package upload

import (
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/analyze"
)

// postData 存储要发送的json信息
type postData struct {
	PostCommitInfo
	Objects []analyze.ObjectInfo `json:"objects"`
}

// postCommits 发送的commit信息
type postCommits struct {
	PostCommitInfo
	Commit []analyze.CommitInfoType `json:"commits"`
}

// PostCommitInfo 发送完成的指示信息
type PostCommitInfo struct {
	Project struct {
		Pid string `json:"pid"`
	} `json:"project"`
	Release struct {
		Release    string `json:"version"`
		CommitHash string `json:"last_commit_hash"`
	} `json:"release"`
}
