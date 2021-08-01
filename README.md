# 函数调用关系分析

## 1.入口函数

`main()`，直接调用`WhosbugPack`下的导出方法：`Analysis()`

在此之前，编译器导入包阶段，`master_func.go`下的`init`将被首先执行：

```go
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

	objectChan = make(chan ObjectInfoType, 1000)
	//开启处理object上传的协程
	for i := 0; i < 1; i++ {
		go processObjectUpload()
	}

}
```

获得全局变量和`input.json`中的配置信息，存入全局变量结构体`config`内.

随后执行`Analysis()`:

```go
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
```

*时间统计后续可能移除

### 1.1 `getLogInfo()`

该函数主要依据`init()`获取的`config`信息，进入仓库目录执行`git log`命令并将输出重定向到原始工作目录下，随后返回重定向文件的路径：

```go
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
```

### 1.2 `matchCommit(diffPath, commitPath string)`

该函数为整个分析过程的主体流程函数，获得的`diff`和`commit`信息将会在这里经历整个数据流动的生命周期，除去内部对其他函数的调用，该函数主要的功能为与`getFullCommit`函数完成`commitDiff`内容的**交错读取**(规避重复读取的关键方法)：

```go
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
```

#### 1.2.1 `getFullCommit(patCommit *regexp.Regexp, lineReaderDiff *bufio.Reader)`

在该函数中，实现了**交错读取**的功能：

`git diff`的存储模式如下：

commit <commitHash>

tree <treeHash>

我们在主函数中逐行读取，并选择匹配**tree**行时转入此函数，此函数依据**commit**行作为匹配标准，因为二者共享同一文件描述符和同一Reader，从而避免了内外匹配标准一致导致的**commit**行丢失、跳行的问题。

```go
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
```

#### 1.2.2 `parseDiffToFile(data, commitHash string)`

此函数为`diff`处理的主流程函数，每一个`diff`的生命周期从此处开始，在此函数内结束。

此函数通过循环匹配将`diff`内容取出，随后送入协程池中进行`antlr`处理:

```go
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
			fmt.Println("pool running: ", pool.Running())
			err := pool.Invoke(diffParsed)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
```

--------

后续调用待完整完成······

