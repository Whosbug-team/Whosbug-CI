package whosbugAssigns

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

func parseDiff(data string) []DiffParsedType {
	t := time.Now()
	patDiff, _ := regexp.Compile(`(diff\ \-\-git\ a/(.*)\ b/.+)`)
	patDiffPart, _ := regexp.Compile(`(@@\ .*?\ @@)`)
	rawDiffs := patDiff.FindAllStringSubmatch(data, -1)
	diffParsedList := make([]DiffParsedType, 0)
	indexList := patDiff.FindAllStringIndex(data, -1)

	for index, rawDiff := range rawDiffs {
		leftDiffIndex := indexList[index][0]
		var diffPartsContent string
		var rightDiffIndex int
		if index == len(rawDiffs)-1 {
			diffPartsContent = data[leftDiffIndex:]
		} else {
			rightDiffIndex = (indexList[index+1])[0]
			diffPartsContent = data[leftDiffIndex:rightDiffIndex]
		}

		diffHeadMatch := patDiffPart.FindAllString(diffPartsContent, -1)
		if diffHeadMatch == nil {
			continue
		}

		rightDiffHeadIndex := patDiffPart.FindStringIndex(diffPartsContent)[1]

		tempFileContent := diffPartsContent[rightDiffHeadIndex:]

		lines := (strings.Split(tempFileContent[0:], "\n"))[1:]

		var changeLineNumbers []ChangeLineNumberType
		changeLineNumbers = findAllChangedLineNumbers(lines)
		lines = replaceLines(lines)
		sourceCode := strings.Join(lines, "")
		fileName := path.Base(rawDiff[2])

		if lanFilter(fileName) {
			commitDicName := data[7:17]
			diffFilePath := fmt.Sprintf("SourceCode/%s/%s", commitDicName, fileName)

			if _, err := os.Stat(path.Dir(diffFilePath)); os.IsNotExist(err) {
				err = os.MkdirAll(path.Dir(diffFilePath), os.ModePerm)
				errorHandler(err, "mkdir ", diffFilePath)
			}

			fd, err := os.OpenFile(diffFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			errorHandler(err, "open ", diffFilePath)
			_, err = fd.WriteString(sourceCode)
			errorHandler(err, "write ", sourceCode)
			err = fd.Close()
			errorHandler(err, "close fd ", sourceCode)
			var diffParsed DiffParsedType
			diffParsed.DiffFile = rawDiff[2]
			diffParsed.DiffFilePath = diffFilePath
			diffParsed.ChangeLineNumbers = append(diffParsed.ChangeLineNumbers, changeLineNumbers...)
			diffParsedList = append(diffParsedList, diffParsed)
		} else {
			continue
		}
	}
	fmt.Println("parsDiff cost ", time.Since(t))
	return diffParsedList
}

/* getLogInfo
/* @Description: 获取release的diff信息
 * @param repoPath 仓库目录/url
 * @param branchName 分支名
 * @param projectId 项目id
 * @return ReleaseDiffType 返回releaseDiff结构体
 * @author KevinMatt 2021-07-26 20:48:22
 * @function_mark PASS
*/
func getLogInfo(repoPath, branchName, projectId string) ReleaseDiffType {
	t := time.Now()
	originPath, err := os.Getwd()
	errorHandler(err, "Get present path ")
	err = os.Chdir(Config.ProjectUrl)
	errorHandler(err, "Cd to ", Config.ProjectUrl)
	fmt.Println("Work path changed to:", Config.ProjectUrl)

	newReleaseCommitHash := execCommandOutput("git", "rev-parse", "HEAD")

	originHash := encrypt(projectId, secret, projectId)
	errorHandler(err, "encrypt Error")
	getLatestRelease(originHash)

	lastReleaseCommitHash := decrypt(projectId, secret, originHash)
	if lastReleaseCommitHash == originHash {
		lastReleaseCommitHash = originHash
	}
	// 执行命令获得输出目录
	getDiffInfo(lastReleaseCommitHash, newReleaseCommitHash)
	var releaseDiff ReleaseDiffType
	releaseDiff.CommitInfoPath, releaseDiff.DiffPath, releaseDiff.BranchName, releaseDiff.HeadCommitInfo = workPath+"\\commit-res", workPath+"\\full-res", branchName, newReleaseCommitHash
	// 返回原工作目录
	err = os.Chdir(originPath)
	fmt.Println("Work path changed back to:", originPath)
	errorHandler(err, "Cd to ", originPath)
	fmt.Println("getLogInfo cost ", time.Since(t))
	return releaseDiff
}

/* getDiffInfo
/* @Description: 根据命令行获取diff信息和commit信息
 * @param lastReleaseCommitHash
 * @param newReleaseCommitHash
 * @return diffInfo
 * @return commitInfo
 * @author KevinMatt 2021-07-26 20:44:07
 * @function_mark
*/
func getDiffInfo(lastReleaseCommitHash, newReleaseCommitHash string) {

	if lastReleaseCommitHash == "" {
		execCommandOutput1("full-res", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
		execCommandOutput1("commit-res", "git", "log", "--pretty=format:%H,%ce,%cn,%cd", fmt.Sprintf("%s..%s", lastReleaseCommitHash, newReleaseCommitHash))
	} else {
		t := time.Now()
		execCommandOutput1("full-res", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw")
		execCommandOutput1("commit-res", "git", "log", "--pretty=format:%H,%ce,%cn,%cd")
		fmt.Println("last release's Commit hash: ", lastReleaseCommitHash)
		fmt.Println("ExecLog cost ", time.Since(t))
	}
	fmt.Println("new release's Commit hash: ", newReleaseCommitHash)

}

//func ReadFile(filePath string, handle func(string)) error {
//	f, err := os.Open(filePath)
//	defer f.Close()
//	if err != nil {
//		return err
//	}
//	buf := bufio.NewReader(f)
//
//	for {
//		line, err := buf.ReadLine("\n")
//		line = strings.TrimSpace(line)
//		handle(line)
//		if err != nil {
//			if err == io.EOF {
//				return nil
//			}
//			return err
//		}
//		return nil
//	}
//}
