package whosbugPack

import (
	javaparser "anrlr4_ast/java"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"time"
)

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

			// 获取单次commit中的每一次diff，并处理diff，存入SourceCode下对应commit文件夹(非支持的语言文件将会被过滤)
			commitDiffs := parseDiffToFile(fullCommit, commitInfo.commitHash)

			// 单次commit中没有含有语言文件改动的diff，直接跳过分析环节
			if commitDiffs != nil {
				analyzeCommitDiff("projectId", commitDiffs, commitInfo.commitHash)
				Result(commitDiffs, commitInfo.commitName, commitInfo.commitEmail, commitInfo.commitTime)
			}

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

/* getFullCommit
/* @Description: 交错读取commit-diff文件
 * @param patCommit 预编译的正则表达式
 * @param lineReaderDiff 全局共享fd
 * @return string 返回完整的commit串
 * @author KevinMatt 2021-07-29 17:52:58
 * @function_mark PASS
*/
func getFullCommit(patCommit *regexp.Regexp, lineReaderDiff *bufio.Reader) string {
	var lines []string
	//lines = make([]string, 500)
	for {
		line, _, err := lineReaderDiff.ReadLine()
		if err == io.EOF {
			break
		}
		// 匹配commit行，交错读取
		res := patCommit.FindString(string(line))
		if res != "" {
			break
		}
		lines = append(lines, string(line))
	}
	return strings.Join(lines, "\n")
}

func Result(resCommits []diffParsedType, committerName string, committerEmail string, commitTime string) {
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

/* parseDiffToFile
/* @Description: 将commit内的diff解析后存入SourceCode中
 * @param data 传入的fullCommit字符串
 * @param commitHash 本次commit的Hash
 * @return []diffParsedType 解析后的信息切片
 * @author KevinMatt 2021-07-29 17:58:43
 * @function_mark PASS
*/
func parseDiffToFile(data, commitHash string) []diffParsedType {
	// 编译正则
	patDiff, _ := regexp.Compile(parDiffPattern)
	patDiffPart, _ := regexp.Compile(parDiffPartPattern)

	// 匹配所有diffs及子匹配->匹配去除a/ or b/的纯目录
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)

	var diffParsedList []diffParsedType
	// 匹配diff行的index列表
	indexList := patDiff.FindAllStringIndex(data, -1)

	// 遍历所有diff
	for index, rawDiff := range rawDiffs {
		// 如果非匹配的语言文件，直接跳过
		if lanFilter(path.Base(rawDiff[2])) {
			// 获得左索引
			leftDiffIndex := indexList[index][0]

			var diffPartsContent string
			var rightDiffIndex int
			// 判断是否为最后一项diff，随后获取代码段
			if index == len(rawDiffs)-1 {
				diffPartsContent = data[leftDiffIndex:]
			} else {
				rightDiffIndex = (indexList[index+1])[0]
				diffPartsContent = data[leftDiffIndex:rightDiffIndex]
			}

			// 匹配@@行
			rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)

			// 无有效匹配直接跳过
			if rightDiffHeadIndex == nil {
				continue
			}

			// 获取所有行，并按"\n"切分，略去第一行(@@行)
			lines := (strings.Split(diffPartsContent[rightDiffHeadIndex[1]:][0:], "\n"))[1:]

			// 传入行切片，寻找所有变动行
			changeLineNumbers := findAllChangedLineNumbers(lines)

			lines = replaceLines(lines)
			sourceCode := strings.Join(lines, "\n")
			diffFilePath := fmt.Sprintf(workPath+"/SourceCode/%s/%s", commitHash[0:10], path.Base(rawDiff[2]))

			if _, err := os.Stat(path.Dir(diffFilePath)); os.IsNotExist(err) {
				err = os.MkdirAll(path.Dir(diffFilePath), os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}

			fd, err := os.OpenFile(diffFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
			_, err = fd.WriteString(sourceCode)
			if err != nil {
				log.Println(err)
			}
			err = fd.Close()
			if err != nil {
				log.Println(err)
			}
			var diffParsed diffParsedType
			diffParsed.diffFileName = rawDiff[2]
			diffParsed.diffFilePath = diffFilePath
			diffParsed.changeLineNumbers = append(diffParsed.changeLineNumbers, changeLineNumbers...)
			diffParsedList = append(diffParsedList, diffParsed)
		} else {
			continue
		}

	}

	return diffParsedList
}

func findAllChangedLineNumbers(lines []string) []changeLineType {
	markCompile, err := regexp.Compile(markPattern)
	if err != nil {
		log.Println(err)
	}
	var changeLineNumbers []changeLineType
	lineNumber := 0
	for index, line := range lines {
		lineNumber = index + 1
		res := markCompile.FindString(line)
		if res != "" {
			var tempStruct changeLineType
			tempStruct.lineNumber = lineNumber
			tempStruct.changeType = string(line[0])
			changeLineNumbers = append(changeLineNumbers, tempStruct)
		}
	}
	return changeLineNumbers
}

func replaceLines(lines []string) []string {
	for index := 0; index < len(lines); index++ {
		if len(lines[index]) >= 1 {
			if string(lines[index][0]) == "+" {
				lines[index] = "" + lines[index][1:]
				//strings.Replace(lines[index], string(lines[index][0]), "", 1)
			} else if string(lines[index][0]) == "-" || lines[index] == "\\ No newline at end of file\r\n" {
				lines[index] = ""
			} else {
				lines[index] = "" + lines[index][1:]
			}
		}
	}
	return lines
}

func analyzeCommitDiff(projectId string, CommitDiffs []diffParsedType, commitId string) {
	//t := time.Now()
	for index := range CommitDiffs {
		CommitDiffs[index].commitHash = commitId
		// 处理后的源码路径
		tempFile := CommitDiffs[index].diffFilePath
		// diff的原始路径
		filePath := CommitDiffs[index].diffFileName
		//t1 := time.Now()
		antlrAnalyzeRes := antlrAnalysis(tempFile, "java")
		//fmt.Println("antlrAnalysis cost", time.Since(t1))
		changeLineNumbers := CommitDiffs[index].changeLineNumbers
		objects := make(map[int]map[string]string)
		for _, changeLineNumber := range changeLineNumbers {
			objects = addObjectFromChangeLineNumber(projectId, filePath, objects, changeLineNumber, antlrAnalyzeRes)
		}
		CommitDiffs[index].diffContent = objects
	}
	//fmt.Println("analyzeCommitDiff cost ", commitId, time.Since(t))
}

func antlrAnalysis(targetFilePath string, langMode string) javaparser.AnalysisInfoType {
	var result javaparser.AnalysisInfoType
	switch langMode {
	case "java":
		result = executeJava(targetFilePath)
		javaparser.Infos.SetEmpty()
	default:
		break
	}
	return result
}

type TreeShapeListener struct {
	*javaparser.BaseJavaParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func executeJava(targetFilePath string) javaparser.AnalysisInfoType {
	input, err := antlr.NewFileStream(targetFilePath)
	if err != nil {
		log.Println(err)
	}
	lexer := javaparser.NewJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := javaparser.NewJavaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return javaparser.Infos
}

func addObjectFromChangeLineNumber(projectId string, filePath string, objects map[int]map[string]string, changeLineNumber changeLineType, antlrAnalyzeRes javaparser.AnalysisInfoType) map[int]map[string]string {
	changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
	if len(objects) > 0 {
		if _, ok := objects[changeMethod.StartLine]; ok {
			return objects
		}
	}
	childHashCode := fmt.Sprintf("%x", hashCode64([]byte(projectId), []byte(changeMethod.MethodName), []byte(filePath)))
	parent := changeMethod.MasterObject
	objects[changeMethod.StartLine] = make(map[string]string)
	objects[changeMethod.StartLine] = map[string]string{
		"name":        changeMethod.MethodName,
		"hash":        childHashCode,
		"parent_name": parent.ObjectName,
		"parent_hash": fmt.Sprintf("%x", hashCode64([]byte(projectId), []byte(parent.ObjectName), []byte(filePath))),
	}
	return objects
}

func findChangedMethod(changeLineNumber changeLineType, antlrAnalyzeRes javaparser.AnalysisInfoType) (changeMethodInfo javaparser.MethodInfoType) {
	var startLineNumbers []int
	for _, part := range antlrAnalyzeRes.AstInfoList.Methods {
		startLineNumbers = append(startLineNumbers, part.StartLine)
	}
	resIndex := findIntervalIndex(startLineNumbers, changeLineNumber.lineNumber)
	if resIndex > -1 {
		changeMethodInfo = antlrAnalyzeRes.AstInfoList.Methods[resIndex]
	}
	return
}

func findIntervalIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}
	for index := range nums {
		if target < nums[index] {
			return index - 1
		} else if target == nums[index] {
			return index
		}
	}
	return -1
}
