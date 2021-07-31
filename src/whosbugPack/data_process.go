package whosbugPack

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

/* getLogInfo
/* @Description: 获取所有的git commit记录和所有的commit+diff，并返回存储的文件目录
 * @return string 所有diff信息的目录
 * @return string 所有commit信息的目录
 * @author KevinMatt 2021-07-29 17:25:39
 * @function_mark PASS
*/
func getLogInfo() (string, string) {
	// 切换到仓库目录
	err := os.Chdir(config.RepoPath)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Work Path Change to: ", config.RepoPath)

	localHashLatest = execCommandOutput("git", "rev-parse", "HEAD")
	// TODO 获得服务器的最新commitHash，此处主要为了验证程序主体功能，暂时没有处理

	cloudHashLatest := ""
	if cloudHashLatest != localHashLatest {
		if cloudHashLatest == "" {
			execRedirectToFile("commitInfo.out", "git", "log", "--pretty=format:%H,%ce,%cn,%cd")
			execRedirectToFile("allDiffs.out", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw")
		} else {
			execRedirectToFile("commitInfo.out", "git", "log", "--pretty=format:%H,%ce,%cn,%cd", fmt.Sprintf("%s..%s", localHashLatest, cloudHashLatest))
			execRedirectToFile("allDiffs.out", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw", fmt.Sprintf("%s..%s", localHashLatest, cloudHashLatest))
		}
	}
	return workPath + "/allDiffs.out", workPath + "/commitInfo.out"
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

/* parseDiffToFile
/* @Description: 将commit内的diff解析后存入SourceCode中
 * @param data 传入的fullCommit字符串
 * @param commitHash 本次commit的Hash
 * @author KevinMatt 2021-07-29 22:54:33
 * @function_mark PASS
*/
func parseDiffToFile(data, commitHash string) {
	// 编译正则
	patDiff, _ := regexp.Compile(parDiffPattern)
	patDiffPart, _ := regexp.Compile(parDiffPartPattern)

	// 匹配所有diffs及子匹配->匹配去除a/ or b/的纯目录
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)

	// 匹配diff行的index列表
	indexList := patDiff.FindAllStringIndex(data, -1)

	// 遍历所有diff
	for index, rawDiff := range rawDiffs {
		// 如果非匹配的语言文件，直接跳过
		if !lanFilter(path.Base(rawDiff[2])) {
			continue
		} else {
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

			// 传入行的切片，寻找所有变动行
			changeLineNumbers := findAllChangedLineNumbers(lines)

			// 替换 +/-行，删除-行内容，切片传递，无需返回值
			replaceLines(lines)

			// 填入到结构体中，准备送入协程
			var diffParsed diffParsedType
			diffParsed.diffText = strings.Join(lines, "\n")
			diffParsed.diffFileName = rawDiff[2]
			diffParsed.changeLineNumbers = append(diffParsed.changeLineNumbers, changeLineNumbers...)
			diffParsed.commitHash = commitHash

			// 得到单个diff后直接送入analyze进行分析
			// TODO 从此可以开始使用goroutine
			fmt.Println("pool running: ", pool.Running())
			err := pool.Invoke(diffParsed)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

/* replaceLines
/* @Description: 替换处理传入的行
 * @param lines 传入的行切片
 * @author KevinMatt 2021-07-29 19:07:41
 * @function_mark PASS
*/
func replaceLines(lines []string) {
	for index := range lines {
		if len(lines[index]) >= 1 {
			if string(lines[index][0]) == "+" {
				lines[index] = "" + lines[index][1:]
				//strings.Replace(lines[index], string(lines[index][0]), "", 1)
			} else if string(lines[index][0]) == "-" || lines[index] == "\\ No newline at end of file" {
				lines[index] = ""
			} else {
				lines[index] = "" + lines[index][1:]
			}
		}
	}
}
