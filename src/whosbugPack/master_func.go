package whosbugPack

import (
	"bufio"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/panjf2000/ants"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// json 替换原始json库
var json = jsoniter.ConfigCompatibleWithStandardLibrary

var pool, _ = ants.NewPoolWithFunc(6, func(commitDiff interface{}) {
	AnalyzeCommitDiff(commitDiff.(diffParsedType))
})

var wg sync.WaitGroup

/* init
/* @Description: 自动初始化配置
 * @author KevinMatt 2021-07-29 20:18:18
 * @function_mark PASS
*/
func init() {
	// 获得密钥
	secret = os.Getenv("WHOSBUG_SECRET")
	if secret == "" {
		secret = "defaultsecret"
	}
	// 工作目录存档
	workPath, _ = os.Getwd()
	file, err := os.Open("src/input.json")
	if err != nil {
		log.Println(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Get input.json succeed!")
	}
	fmt.Println("Version:\t", config.ReleaseVersion, "\nProjectId:\t", config.ProjectId, "\nBranchName:\t", config.BranchName)

	ObjectChan = make(chan objectInfoType, 1000)
	err = os.Remove("allDiffs.out")
	if err != nil {
		log.Println(err)
	}
	err = os.Remove("commitInfo.out")
	if err != nil {
		log.Println(err)
	}
	//开启处理object上传的协程
	for i := 0; i < 5; i++ {
		go processObjectUpload()
	}
}

// Analysis
/* @Description: 暴露给外部的函数，作为程序入口
 * @author KevinMatt 2021-07-29 17:51:28
 * @function_mark PASS
 */
func Analysis() {
	//defer pool.Release()
	t := time.Now()
	// 获取git log命令得到的commit列表和完整的commit-diff信息存储的文件目录
	diffPath, commitPath := getLogInfo()
	fmt.Println("Get log cost: ", time.Since(t))
	matchCommit(diffPath, commitPath)
	for {
		if pool.Running() == 0 {
			fmt.Println("Close pool!")
			pool.Release()
			close(ObjectChan)
			break
		}
		time.Sleep(time.Second)
	}
	fmt.Println("Total cost: ", time.Since(t))
	// 等待上传协程的结束
	wg.Wait()
}

/* matchCommit
/* @Description: 主体过程，最后直接生成结果集，位置在SourceCode下(此部分可做商榷)
 * @param diffPath diff-commit文件目录
 * @param commitPath commit-info文件目录
 * @author KevinMatt 2021-07-29 17:37:10
 * @function_mark PASS
*/
func matchCommit(diffPath, commitPath string) {
	processCommits := 0
	commitFd, err := os.Open(commitPath)
	if err != nil {
		log.Println(err)
	}
	diffFd, err := os.Open(diffPath)
	if err != nil {
		log.Println(err)
	}
	lineReaderCommit := bufio.NewReader(commitFd)
	lineReaderDiff := bufio.NewReader(diffFd)
	for {
		line, _, err := lineReaderDiff.ReadLine()
		if err == io.EOF {
			break
		}

		// 匹配tree行
		res := patTree.FindString(string(line))
		if res != "" {
			// 匹配到一个commit的tree行，从commit info读一行
			commitLine, _, err := lineReaderCommit.ReadLine()
			if err == io.EOF {
				break
			}
			var commitInfo commitInfoType
			infoList := strings.Split(string(commitLine), ",")

			// 填充commitInfo结构体内的各项信息
			for index := 2; index < len(infoList)-1; index++ {
				commitInfo.committerName += infoList[index]
				if index != len(infoList)-2 {
					commitInfo.committerName += ","
				}
			}
			commitInfo.commitHash, commitInfo.committerEmail, commitInfo.commitTime = infoList[0], infoList[1], toIso8601(strings.Split(infoList[len(infoList)-1][4:], " "))

			// 获取一次完整的commit，使用循环交错读取的方法避免跳过commit
			fullCommit := getFullCommit(patCommit, lineReaderDiff)

			// 获取单次commit中的每一次diff，并处理diff，送进协程
			parseDiffToFile(fullCommit, commitInfo)

			// 指示已经处理的commit数量
			processCommits++
			fmt.Println("Commit No.", processCommits, " ", commitInfo.commitHash, " Sent Into Channel.")
		}
		// 强制触发GC,避免短解析作业在golang自动gc触发的两分钟阈值内大量堆积
		runtime.GC()
	}
	err = commitFd.Close()
	if err != nil {
		log.Println(err)
	}
	err = diffFd.Close()
	if err != nil {
		log.Println(err)
	}
}
