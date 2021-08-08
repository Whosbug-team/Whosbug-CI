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
