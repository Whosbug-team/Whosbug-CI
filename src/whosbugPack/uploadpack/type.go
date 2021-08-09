package uploadpack

import (
	"whosbugPack/global_type"
)

// postData 存储要发送的json信息
type postData struct {
	Objects []global_type.ObjectInfoType `json:"objects"`
	Project struct {
		Pid string `json:"pid"`
	} `json:"project"`
	Release struct {
		Release    string `json:"release"`
		CommitHash string `json:"commit_hash"`
	} `json:"release"`
}

// postCommits 发送的commit信息
type postCommits struct {
	Project struct {
		Pid string `json:"pid"`
	} `json:"project"`
	Release struct {
		Release    string `json:"release"`
		CommitHash string `json:"commit_hash"`
	} `json:"release"`
	Commit []global_type.CommitInfoType `json:"commits"`
}

// postFin 发送完成的指示信息
type postFin struct {
	Project struct {
		Pid string `json:"pid"`
	} `json:"project"`
	Release struct {
		Release    string `json:"release"`
		CommitHash string `json:"commit_hash"`
	} `json:"release"`
}
