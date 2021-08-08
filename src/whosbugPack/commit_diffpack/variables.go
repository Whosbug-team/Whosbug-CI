package commit_diffpack

import (
	"github.com/panjf2000/ants"
	"regexp"
	"whosbugPack/antlrpack"
	"whosbugPack/global_type"
)

// parCommitPattern 匹配commit行
const parCommitPattern = `(commit\ ([a-f0-9]{40}))`

// parTreePattern 匹配tree行，用于`交错匹配`
const parTreePattern = `(tree\ ([a-f0-9]{40}))`

// patCommit, patTree 预编译正则匹配
var patCommit, _ = regexp.Compile(parCommitPattern)
var patTree, _ = regexp.Compile(parTreePattern)

// parDiffPattern 匹配diff行，用于每一次commit信息的处理
const parDiffPattern = `(diff\ \-\-git\ a/(.*)\ b/.+)`

// parDiffPartPattern 匹配diff段的末行@@行，用于获取diff代码内容的起始位置
const parDiffPartPattern = `(@@\ .*?\ @@)`

// markPattern 匹配+/-变动行
const markPattern = `^[\+\-]`

// 解析协程池
var Pool, _ = ants.NewPoolWithFunc(6, func(commitDiff interface{}) {
	antlrpack.AnalyzeCommitDiff(commitDiff.(global_type.DiffParsedType))
})
