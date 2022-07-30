package commit

import (
	"log"
	"regexp"
	"runtime"

	"git.woa.com/bkdevops/whosbug/antlr"
	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/logging"
	"github.com/panjf2000/ants"
	"github.com/schollz/progressbar/v3"
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
const parDiffPartPattern = `(@@\ (.*?)\ @@)`

// markPattern 匹配+/-变动行
const markPattern = `^[\+\-]`

// ProcessBar
var ProcessBar *progressbar.ProgressBar

// 解析协程池
var AntlrAnalysisPool, _ = ants.NewPoolWithFunc(runtime.NumCPU()/4, func(commitDiff interface{}) {
	antlr.AnalyzeCommitDiff(commitDiff.(config.DiffParsedType))
	// 指示已经处理的diff数量
	processDiffs++
	log.SetOutput(logging.LogFile)
	log.Println("Diff No.", processDiffs, " From", commitDiff.(config.DiffParsedType).CommitHash, " Sent Into Channel.")
	// if processDiffs%50 == 0 {
	// runtime.GC()
	// }
})

// processDiffs 已处理的commit数
var processDiffs int

// 编译正则
var patDiff, _ = regexp.Compile(parDiffPattern)
var patDiffPart, _ = regexp.Compile(parDiffPartPattern)
var markCompile, _ = regexp.Compile(markPattern)
