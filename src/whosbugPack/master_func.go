package whosbugPack

import (
	"bufio"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/panjf2000/ants"
	"io"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// json 替换原始json库
var json = jsoniter.ConfigCompatibleWithStandardLibrary

var pool, _ = ants.NewPoolWithFunc(10, func(commitDiff interface{}) {
	AnalyzeCommitDiff(commitDiff.(diffParsedType))
	runtime.GC()
})

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
	fmt.Println("Total cost: ", time.Since(t))
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
	patCommit, _ := regexp.Compile(parCommitPattern)
	patTree, _ := regexp.Compile(parTreePattern)
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

			// 获取一次完整的commit，使用双循环交错读取方法避免跳过commit
			fullCommit := getFullCommit(patCommit, lineReaderDiff)

			// 强制手动触发gc，及时释放getFullCommit的原始拷贝字符串
			runtime.GC()

			// 获取单次commit中的每一次diff，并处理diff，送进协程
			parseDiffToFile(fullCommit, commitInfo.commitHash)

			// 指示已经处理的commit数量
			processCommits++
			fmt.Println("Commit No.", processCommits, " ", commitInfo.commitHash, " done.", "pool available", pool.Free())
		}
		// 强制手动触发GC,避免短解析作业在golang自动gc触发的两分钟阈值内大量堆积内存
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
