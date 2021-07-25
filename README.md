# Whosbug-Golang

# 函数调用关系

## `main()`

```go
func main() {
fmt.Println("Start!")
whosbugAssigns.GetInputConfig()
projectId := whosbugAssigns.Config.ProjectId
branchName := whosbugAssigns.Config.BranchName
repoPath := whosbugAssigns.Config.ProjectUrl
resCommits := whosbugAssigns.Analysis(repoPath, branchName, projectId)
whosbugAssigns.Result(resCommits, projectId, "1.0.0")
fmt.Println("Whosbug analysis done")
}
```

调用关系：

`main–>whosbugAssigns.GetInputConfig()–> whosbugAssigns.AnalysisTest()–>whosbugAssigns.result()`

<img src="https://kevinmatt-1303917904.cos.ap-chengdu.myqcloud.com/temp/image-20210725142816533.png" alt="image-20210725142816533" style="zoom:50%;" />

### `GetInputConfig()`

```go
func GetInputConfig() {
file, err := os.Open("src/input.json")
if err != nil {
fmt.Println(err.Error())
}

decoder := json.NewDecoder(file)
err = decoder.Decode(&config)
if err != nil {
fmt.Println(err.Error())
} else {
fmt.Println("Get input.json succeed!")
}
fmt.Println("Version:\t", config.ReleaseVersion, "\nProjectId:\t", config.ProjectId, "\nBranchName:\t", config.BranchName)
}
// config: type input_json
// type input_json struct {
//	  ProjectId       string   `json:"__PROJRCT_ID"`
//	  ReleaseVersion  string   `json:"__RELEASE_VERSION"`
//	  ProjectUrl      string   `json:"__PROJECT_URL"`
//	  BranchName      string   `json:"__BRANCH_NAME"`
//	  LanguageSupport []string `json:"__LAN_SUPPORT"`
//	  WebServerHost   string   `json:"__WEB_SRV_HOST""`
//  }
```

读取src/input.json内的内容，存储到结构体中

### `Analysis()`

作为整体分析函数逻辑的主体部分，大部分调用从这里出发

```go
// Analysis
/* @Description: 分析调用主逻辑函数
 * @param repoPath 仓库地址/url
 * @param branchName 分支名
 * @param projectId 项目id
 * @return []CommitParsedType 返回解析后的commit信息
 * @author KevinMatt 2021-07-25 13:08:45
 * @function_mark PASS
 */
func Analysis(repoPath, branchName, projectId string) []CommitParsedType {
releaseDiff := getDiff(repoPath, branchName, projectId)
commits := parseCommit(releaseDiff.Diff, strings.Split(releaseDiff.CommitInfo, "\n"))
var parsedCommits []CommitParsedType
for index := range commits {
commit := commits[index]
commitId := commit.Commit
var diffPark string
if index == len(commits)-1 {
diffPark = releaseDiff.Diff[commit.CommitLeftIndex:]
} else {
nextCommitLeftIndex := commits[index+1].CommitLeftIndex
diffPark = releaseDiff.Diff[commit.CommitLeftIndex:nextCommitLeftIndex]
}
commitDiffs := parseDiff(diffPark)
commit = analyzeCommitDiff(projectId, commitDiffs, commitId, commit)
parsedCommits = append(parsedCommits, commit)
}
return parsedCommits
}
```

调用关系图：

<img src="https://kevinmatt-1303917904.cos.ap-chengdu.myqcloud.com/temp/image-20210725142839010.png" alt="image-20210725142839010" style="zoom: 50%;" />

#### `getDiff()`

该函数通过执行`git log --full-diff -p -U1000 --pretty=raw`获取输出的Diff信息

```go
// getDiff
/* @Description: 获取release的diff信息
 * @param repoPath 仓库目录/url
 * @param branchName 分支名
 * @param projectId 项目id
 * @return ReleaseDiffType 返回releaseDiff结构体
 * @author KevinMatt 2021-07-25 13:12:07
 * @function_mark PASS
 */
func getDiff(repoPath, branchName, projectId string) ReleaseDiffType {

secret := os.Getenv("WHOSBUG_SECRET")
originPath, err := os.Getwd()
errorHandler(err)
err = os.Chdir(repoPath)
errorHandler(err)
fmt.Println("Work path changed to:", repoPath)

newReleaseCommitHash := execCommandOutput("git", "rev-parse", "HEAD")

originHash := make([]byte, len(projectId))
err = encrypt([]byte(projectId), originHash, []byte(secret), []byte(projectId))
errorHandler(err)
getLatestRelease(string(originHash))
lastReleaseCommitHash := make([]byte, len(originHash))

err = decrypt([]byte(projectId), lastReleaseCommitHash, []byte(secret), originHash)
if string(lastReleaseCommitHash) != string(originHash) {
lastReleaseCommitHash = nil
}
errorHandler(err)
fmt.Println("last release's Commit hash: ", string(lastReleaseCommitHash))
fmt.Println("new release's Commit hash: ", newReleaseCommitHash)

var diff, commitInfo string
if string(lastReleaseCommitHash) != "" {
diff = execCommandOutput("git", "log", "--full-diff", "-p", "-U1000", "--pretty=raw", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
commitInfo = execCommandOutput("git", "log", "--pretty=format:%H,%ce,%cn,%cd", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
} else {
diff = execCommandOutput("git", "log", "--full-diff", "-p", "-U1000", "--pretty=raw")
commitInfo = execCommandOutput("git", "log", "--pretty=format:%H,%ce,%cn,%cd")
}
var releaseDiff ReleaseDiffType
releaseDiff.CommitInfo = commitInfo
releaseDiff.Diff = diff
releaseDiff.BranchName = branchName
releaseDiff.HeadCommitInfo = newReleaseCommitHash

// 返回原工作目录
err = os.Chdir(originPath)
fmt.Println("Work path changed back to:", originPath)
errorHandler(err)

return releaseDiff
}
```

<img src="https://kevinmatt-1303917904.cos.ap-chengdu.myqcloud.com/temp/image-20210725142908729.png" alt="image-20210725142908729" style="zoom:50%;" />

##### `execCommandOutput()`

```go
// execCommandOutput
/* @Description: 执行命令并获取命令标准输出
 * @param command 命令程序
 * @param args 命令参数(切片)
 * @return string 标准输出内容
 * @author KevinMatt 2021-07-25 13:16:22
 * @function_mark PASS
 */
func execCommandOutput(command string, args ...string) string {
cmd := exec.Command(command, args...)
output := bytes.Buffer{}
cmd.Stdout = &output
err := cmd.Run()
errorHandler(err, "exec command error:")
return output.String()
}
```

##### `encrypt()`

```go
// encrypt
/* @Description: AES-CFB加密
 * @param projectId 项目ID
 * @param Dest 输出的加密后字符串(输出参数)
 * @param key 加密密钥
 * @param plainText 需要加密的文本
 * @return error 错误抛出
 * @author KevinMatt 2021-07-25 13:34:09
 * @function_mark PASS
 */
func encrypt(projectId, Dest, key, plainText []byte) error {
K, IV := generateKIV(projectId, key)
aesBlockEncryptor, err := aes.NewCipher(K)
if err != nil {
return err
}
aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
aesEncryptor.XORKeyStream(Dest, plainText)
return nil
}
```

##### `decrypt()`

```go
// decrypt
/* @Description: AES-CFB解密
 * @param projectId 项目ID
 * @param Dest 解密完成的字符串
 * @param key 解密密钥
 * @param plainText 需要解密的文本
 * @return error 错误抛出
 * @author KevinMatt 2021-07-25 13:35:15
 * @function_mark PASS
 */
func decrypt(projectId, Dest, key, plainText []byte) error {
K, IV := generateKIV(projectId, key)
aesBlockDescriptor, err := aes.NewCipher(K)
if err != nil {
return err
}
aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
aesDescriptor.XORKeyStream(Dest, plainText)
return nil
}
```

#### `parseCommit()`

通过正则匹配表达式`(commit\ ([a-f0-9]{40}))`匹配每一次commit的信息，存入`parsedCommits`中返回

parsedCommit的类型：

```go
type CommitParsedType struct {
CommitLeftIndex int
Commit          string
CommitTime      string
CommitterInfo   CommitterInfoType
CommitDiffs     []DiffParsedType
}
```

```go
// parseCommit
/* @Description: 解析commit信息
 * @param data 传入数据的diff部分(git log元数据)
 * @param commitInfos  log元数据分片
 * @return []CommitParsedType
 * @author KevinMatt 2021-07-25 13:36:35
 * @function_mark PASS
 */
func parseCommit(data string, commitInfos []string) []CommitParsedType {
patCommit, err := regexp.Compile(`(commit\ ([a-f0-9]{40}))`)
errorHandler(err)
rawCommits := patCommit.FindAllStringSubmatch(data, -1)
var parsedCommits []CommitParsedType
for index, commitInfoLine := 0, commitInfos[0]; index < len(rawCommits) && index < len(commitInfos); index++ {
commitInfoLine = commitInfos[index]
infoList := strings.Split(commitInfoLine, ",")
timeList := strings.Split(infoList[3][4:], " ")
var parsedCommit CommitParsedType
parsedCommit.CommitLeftIndex = patCommit.FindAllStringSubmatchIndex(data, -1)[index][0]
parsedCommit.Commit = infoList[0]
parsedCommit.CommitTime = toIso8601(timeList)
parsedCommit.CommitterInfo.Name = infoList[2]
parsedCommit.CommitterInfo.Email = infoList[1]
parsedCommits = append(parsedCommits, parsedCommit)
}
return parsedCommits
}
```

##### `toIso8601()`

```go
// toIso8601
/* @Description: 时间戳转换
 * @param timeList
 * @return string
 * @author KevinMatt 2021-07-25 13:42:29
 * @function_mark PASS
 */
func toIso8601(timeList []string) string {
return fmt.Sprintf("%s-%s-%sT%s%s:%s", timeList[3], month_correspond[timeList[0]], timeList[1], timeList[2], timeList[4][3:], timeList[4][3:])
}
```

#### `parseDiff()`

通过正则表达式`(diff\ \-\-git\ a/(.*)\ b/.+)`, `(@@\ .*?\ @@)`在diff内容中匹配变动文件名、行数变化

```go
// parseDiff
/* @Description: 将git log的信息的diff部分分解提取
 * @param data
 * @return []DiffParsedType
 * @author KevinMatt 2021-07-25 13:43:55
 * @function_mark PASS
 */
func parseDiff(data string) []DiffParsedType {
patDiff, err := regexp.Compile(`(diff\ \-\-git\ a/(.*)\ b/.+)`)
errorHandler(err)
patDiffPart, err := regexp.Compile(`(@@\ .*?\ @@)`)
errorHandler(err)
rawDiffs := patDiff.FindAllStringSubmatch(data, -1)
diffParsed := make([]DiffParsedType, 0)

for index, rawCommit := range rawDiffs {
parts := rawCommit[2]
leftDiffIndex := patDiff.FindAllStringIndex(data, -1)[index][0]
var diffPartsContent string
var rightDiffIndex int
if index == len(rawDiffs)-1 {
diffPartsContent = data[leftDiffIndex:]
} else {
rightDiffIndex = (patDiff.FindAllStringIndex(data, -1)[index+1])[0]
diffPartsContent = data[leftDiffIndex:rightDiffIndex]
}
diffHeadMatch := patDiffPart.FindAllString(diffPartsContent, -1)

if diffHeadMatch == nil {
continue
}
rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)[1]
tempFileContent := diffPartsContent[rightDiffHeadIndex:]
lines := (strings.SplitAfter(tempFileContent[0:], "\n"))[1:]
var changeLineNumbers []ChangeLineNumberType
changeLineNumbers = findAllChangedLineNumbers(lines)
lines = replaceLines(lines)
sourceCode := strings.Join(lines, "")
fileName := path.Base(parts)

if lanFilter(fileName) {
commitDicName := data[7:17]
diffFilePath := fmt.Sprintf("SourceCode/%s/%s", commitDicName, fileName)

if _, err := os.Stat(path.Dir(diffFilePath)); os.IsNotExist(err) {
err := os.MkdirAll(path.Dir(diffFilePath), os.ModePerm)
errorHandler(err)
}
fd, err := os.OpenFile(diffFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
errorHandler(err)
_, err = fd.WriteString(sourceCode)
errorHandler(err)
err = fd.Close()
errorHandler(err)
var diffSingle DiffParsedType
diffSingle.DiffFile = parts
diffSingle.DiffFilePath = diffFilePath
diffSingle.ChangeLineNumbers = append(diffSingle.ChangeLineNumbers, changeLineNumbers...)
diffParsed = append(diffParsed, diffSingle)
} else {
continue
}
}
return diffParsed
}
```

##### `findAllChangedLineNumbers()`

匹配所有变更行并存储其变更类型到结构体中：

```go
type ChangeLineNumberType struct {
LineNumber int
ChangeType string
}
```

```go
// findAllChangedLineNumbers
/* @Description: 匹配所有改变行(以+/-开头的行)的行号
 * @param lines 传入diff中的所有代码行(完整文件代码行)
 * @return []ChangeLineNumberType 返回存储所有变更行信息的切片
 * @author KevinMatt 2021-07-25 13:47:42
 * @function_mark PASS
 */
func findAllChangedLineNumbers(lines []string) []ChangeLineNumberType {
markCompile, err := regexp.Compile(`^[\+\-]`)
errorHandler(err)
changeLineNumbers := make([]ChangeLineNumberType, 0)
lineNumber := 0
for _, line := range lines {
lineNumber++
res := markCompile.FindString(line)
if res != "" {
var tempStruct ChangeLineNumberType
tempStruct.LineNumber = lineNumber
tempStruct.ChangeType = string(line[0])
changeLineNumbers = append(changeLineNumbers, tempStruct)
}
}
return changeLineNumbers
}
```

##### `replaceLines()`

```go
// replaceLines
/* @Description: 清除+/-符号并移除-行和No newline提示
 * @param lines 传入行集合
 * @return []string
 * @author KevinMatt 2021-07-25 13:52:57
 * @function_mark PASS
 */
func replaceLines(lines []string) []string {
for index := 0; index < len(lines); index++ {
if len(lines[index]) > 1 {
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
```

#### `analyzeCommitDiff()`

```go
// analyzeCommitDiff
/* @Description: 分析commitDiff
 * @param projectId 项目ID
 * @param commitDiffs commitDiff切片
 * @param commitId CommitHash
 * @param commit 解析后的commit信息
 * @return CommitParsedType
 * @author KevinMatt 2021-07-25 13:54:04
 * @function_mark PASS
 */
func analyzeCommitDiff(projectId string, commitDiffs []DiffParsedType, commitId string, commit CommitParsedType) CommitParsedType {
for index := 0; index < len(commitDiffs); index++ {
commitDiff := commitDiffs[index]
commitDiff.Commit = commitId
// 处理后的源码路径
tempFile := commitDiff.DiffFilePath
// diff的原始路径
filePath := commitDiff.DiffFile
antlrAnalyzeRes := antlrAnalysis(tempFile, "java")

changeLineNumbers := commitDiff.ChangeLineNumbers
objects := make(map[int]map[string]string)
for _, changeLineNumber := range changeLineNumbers {
objects = addObjectFromChangeLineNumber(projectId, filePath, objects, changeLineNumber, antlrAnalyzeRes)
}
commitDiff.DiffContent = objects
commit.CommitDiffs = append(commit.CommitDiffs, commitDiff)
}
return commit
}
```

##### `antlrAnalysis()`

```go
// antlrAnalysis
/* @Description: 执行antlr分析入口函数
 * @param targetFilePath 目标代码目录
 * @param langMode 解析语言模式
 * @return javaparser.AnalysisInfoType
 * @author KevinMatt 2021-07-25 13:56:08
 * @function_mark PASS
 */
func antlrAnalysis(targetFilePath string, langMode string) javaparser.AnalysisInfoType {
var result javaparser.AnalysisInfoType
// TODO 目前只有Java的支持
switch langMode {
case "java":
result = executeJava(targetFilePath)
javaparser.Infos.SetEmpty()
default:
break
}
return result
}
```

###### `excuteJava()`

```go
// executeJava
/* @Description: 执行Java Antlr语法解析
 * @param targetFilePath 解析目标目录
 * @return javaparser.AnalysisInfoType
 * @author KevinMatt 2021-07-25 14:00:10
 * @function_mark PASS
 */
func executeJava(targetFilePath string) javaparser.AnalysisInfoType {
input, err := antlr.NewFileStream(targetFilePath)
if err != nil {
errorHandler(err)
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
```

##### `addObjectFromChangeLineNumber()`

```go
type AnalysisInfoType struct {
CallMethods []string
AstInfoList astInfoType
}
```

```go
// addObjectFromChangeLineNumber
/* @Description:  存储分析得到的方法改变信息(基于行号索引)
 * @param projectId 项目ID
 * @param filePath	文件目录
 * @param objects 	传入的空参数
 * @param changeLineNumber 改变行行号
 * @param antlrAnalyzeRes  分析结果
 * @return map[int]map[string]string
 * @author KevinMatt 2021-07-25 14:03:36
 * @function_mark PASS
 */
func addObjectFromChangeLineNumber(projectId string, filePath string, objects map[int]map[string]string, changeLineNumber ChangeLineNumberType, antlrAnalyzeRes javaParser.AnalysisInfoType) map[int]map[string]string {
changeMethod := findChangedMethod(changeLineNumber, antlrAnalyzeRes)
if len(objects) > 0 {
if _, ok := objects[changeMethod.StartLine]; ok {
return objects
}
}
childHashCode := hashCode64(projectId, changeMethod.MethodName, filePath)
parent := changeMethod.MasterObject
objects[changeMethod.StartLine] = make(map[string]string)
objects[changeMethod.StartLine] = map[string]string{
"Name":        changeMethod.MethodName,
"hash":        childHashCode,
"parent_name": parent.ObjectName,
"parent_hash": hashCode64(projectId, parent.ObjectName, filePath),
}
return objects
}
```

###### `findChangedMethod()`

```go
// findChangedMethod
/* @Description: 			寻找变更的方法
 * @param changeLineNumber 	变更行信息
 * @param antlrAnalyzeRes 	分析结果
 * @return javaParser.MethodInfoType
 * @author KevinMatt 2021-07-25 14:11:45
 * @function_mark
 */
func findChangedMethod(changeLineNumber ChangeLineNumberType, antlrAnalyzeRes javaParser.AnalysisInfoType) javaParser.MethodInfoType {
var changeMethodInfo javaParser.MethodInfoType
startLineNumbers := make([]int, 0)
for _, part := range antlrAnalyzeRes.AstInfoList.Methods {
startLineNumbers = append(startLineNumbers, part.StartLine)
}
resIndex := searchInsert(startLineNumbers, changeLineNumber.LineNumber)
if resIndex > -1 {
changeMethodInfo = antlrAnalyzeRes.AstInfoList.Methods[resIndex]
}
return changeMethodInfo
}
```

findIntervalIndex

```go
// findIntervalIndex
/* @Description: 	寻找插入空隙
 * @param nums		切片lineNumbers
 * @param target	目标lineNumber
 * @return int 		空隙位置
 * @author KevinMatt 2021-07-25 14:17:52
 * @function_mark 	PASS
 */
func findIntervalIndex(nums []int, target int) int {
if nums == nil {
return -1
}
if len(nums) >= 2 && target > nums[1] {
return -1
}
if target < nums[0] {
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
```

###### `hashCode64`

```go
// hashCode64
/* @Description: 返回sha256编码的拼接字符串
 * @param projectId 项目ID
 * @param objectName
 * @param filePath 文件目录
 * @return string 返回编码字符串
 * @author KevinMatt 2021-07-25 14:20:09
 * @function_mark
 */
func hashCode64(projectId string, objectName string, filePath string) string {
text := projectId + objectName + filePath
return string(sha256.New().Sum([]byte(text)))
}
```

