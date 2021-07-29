package whosbugPack

import (
	"bufio"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// json 替换原始json库
var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
/* @Description: 唯一暴露给外部的函数，作为程序入口
 * @author KevinMatt 2021-07-29 17:51:28
 * @function_mark PASS
 */
func Analysis() {
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
			commitInfo.commitHash, commitInfo.commitEmail, commitInfo.commitName, commitInfo.commitTime = infoList[0], infoList[1], infoList[2], toIso8601(strings.Split(infoList[3][4:], " "))

			// 获取一次完整的commit，使用双循环交错读取方法避免跳过commit
			fullCommit := getFullCommit(patCommit, lineReaderDiff)

			// 强制手动触发gc，及时释放getFullCommit的原始拷贝字符串
			runtime.GC()

			// 获取单次commit中的每一次diff，并处理diff，送进协程
			parseDiffToFile(fullCommit, commitInfo.commitHash)

			// 指示已经处理的commit数量
			processCommits++
			fmt.Println("Commit No.", processCommits, " ", commitInfo.commitHash, " done.")
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

/* resultToFile
/* @Description: 保留函数，暂时弃用
 * @param resCommits
 * @param committerName
 * @param committerEmail
 * @param commitTime
 * @author KevinMatt 2021-07-29 23:11:54
 * @function_mark
*/
func resultToFile(resCommits []diffParsedType, committerName string, committerEmail string, commitTime string) {
	latestCommitHash := resCommits[0].commitHash

	project := map[string]string{
		"pid": config.ProjectId,
	}

	release := map[string]string{
		"release":     fmt.Sprintf("%x", encrypt(config.ProjectId, secret, config.ReleaseVersion)),
		"commit_hash": fmt.Sprintf("%x", encrypt(config.ProjectId, secret, latestCommitHash)),
	}
	var objects []map[string]string
	owner := fmt.Sprintf("%s-%s", committerName, committerEmail)
	for _, diffFile := range resCommits {
		filePath := path.Base(diffFile.diffFilePath)
		for _, value := range diffFile.diffContent {
			if value["name"] == "" {
				continue
			}
			tempMap := map[string]string{
				"owner":       fmt.Sprintf("%x", encrypt(config.ProjectId, secret, owner)),
				"file_path":   fmt.Sprintf("%x", encrypt(config.ProjectId, secret, filePath)),
				"parent_name": fmt.Sprintf("%x", encrypt(config.ProjectId, secret, value["parent_name"])),
				"parent_hash": fmt.Sprintf("%x", encrypt(config.ProjectId, secret, value["parent_hash"])),
				"name":        fmt.Sprintf("%x", encrypt(config.ProjectId, secret, value["name"])),
				"hash":        fmt.Sprintf("%x", encrypt(config.ProjectId, secret, value["hash"])),
				"old_name":    "",
				"commit_time": commitTime,
			}
			objects = append(objects, tempMap)
		}
	}

	res := map[string]interface{}{
		"objects": objects,
		"release": release,
		"project": project,
	}
	if _, err := os.Stat(workPath + "/SourceCode/" + latestCommitHash[0:10] + "/"); os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(workPath+"/SourceCode/"+latestCommitHash[0:10]+"/"), os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	fd, err := os.OpenFile(workPath+"/SourceCode/"+latestCommitHash[0:10]+"/res.json", os.O_RDWR|os.O_CREATE|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	jsonInfo, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	_, _ = fd.WriteString(string(jsonInfo))
	err = fd.Close()
}
