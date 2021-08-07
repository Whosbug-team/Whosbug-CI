package whosbugPack

import (
	"regexp"
	"sync"
)

// 全局变量，存储config信息
var config inputJson

// 全局变量，记录起始路径
var workPath string

// 全局变量，存储密钥
var secret string

// 全局变量，记录本地最后commit hash
var localHashLatest string

// 全局变量，object传输的通道
var ObjectChan chan objectInfoType

// 全局变量，大型object传输的通道
var ObjectChanLarge chan objectInfoType

// parCommitPattern 匹配commit行
const parCommitPattern = `(commit\ ([a-f0-9]{40}))`

// parTreePattern 匹配tree行，用于`交错匹配`
const parTreePattern = `(tree\ ([a-f0-9]{40}))`

// parDiffPattern 匹配diff行，用于每一次commit信息的处理
const parDiffPattern = `(diff\ \-\-git\ a/(.*)\ b/.+)`

// parDiffPartPattern 匹配diff段的末行@@行，用于获取diff代码内容的起始位置
const parDiffPartPattern = `(@@\ .*?\ @@)`

// markPattern 匹配+/-变动行
const markPattern = `^[\+\-]`

// supportLans 语言的支持
var supportLans = []string{".java"}

// patCommit, patTree 预编译正则匹配
var patCommit, _ = regexp.Compile(parCommitPattern)
var patTree, _ = regexp.Compile(parTreePattern)

// sendCount 上传计数器
var sendCount int

// processCommits 已处理的commit数
var processCommits = 0

// postDataPool, objectInfoPool 上传数据变量同步池
var (
	postDataPool = &sync.Pool{New: func() interface{} {
		return new(postData)
	}}
	objectInfoPool = &sync.Pool{New: func() interface{} {
		return new(objectInfoType)
	}}
)
