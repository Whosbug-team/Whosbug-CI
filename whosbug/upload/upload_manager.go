package upload

import (
	"sync"
)

// postDataPool, objectInfoPool 上传数据变量同步池
var (
	postDataPool = &sync.Pool{New: func() interface{} {
		return new(postData)
	}}
)

// sendCount 发送的协程计数
var sendCount int

// UploadWaitGroup 上传协程池
var UploadWaitGroup sync.WaitGroup

var postProjectInfo PostCommitInfo

var isInitial = true
