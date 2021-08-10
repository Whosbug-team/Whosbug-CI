package uploadpack

import (
	"whosbugPack/global_type"
)

// postData 存储要发送的json信息
type postData struct {
	PostCommitInfo
	Objects []global_type.ObjectInfoType `json:"objects"`
}

// postCommits 发送的commit信息
type postCommits struct {
	PostCommitInfo
	Commit []global_type.CommitInfoType `json:"commits"`
}

// PostCommitInfo 发送完成的指示信息
type PostCommitInfo struct {
	Project struct {
		Pid string `json:"pid"`
	} `json:"project"`
	Release struct {
		Release    string `json:"release"`
		CommitHash string `json:"commit_hash"`
	} `json:"release"`
}
