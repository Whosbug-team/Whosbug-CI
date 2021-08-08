package uploadpack

import (
	"sync"
	"whosbugPack/global_type"
)

// postDataPool, objectInfoPool 上传数据变量同步池
var (
	postDataPool = &sync.Pool{New: func() interface{} {
		return new(postData)
	}}
	objectInfoPool = &sync.Pool{New: func() interface{} {
		return new(global_type.ObjectInfoType)
	}}
)

// sendCount 发送的协程计数
var sendCount int

const _HOST = "http://127.0.0.1:8081"
const _SECRET = ""
const _USERNAME = "user"
const _PASSWORD = "pwd"

var UploadWaitGroup sync.WaitGroup
