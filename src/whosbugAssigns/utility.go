package whosbugAssigns

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
)

var supportLans = []string{".java"}

/** execCommandOutput
 * @Description: 执行终端命令
 * @param command 命令程序
 * @param args 命令参数
 * @return string 标准输出内容
 * @author KevinMatt 2021-07-22 16:49:44
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

/** lanFilter
 * @Description: 语言过滤器，确定目标文件是否为支持的语言
 * @param fileName 文件名
 * @return bool 是否支持
 * @author KevinMatt 2021-07-22 16:48:53
 * @function_mark PASS
 */
func lanFilter(fileName string) bool {
	for _, supportLan := range supportLans {
		if path.Ext(fileName) == supportLan {
			return true
		}
	}
	return false
}

func result(resCommits []CommitParsedType, projectId string, releaseVersion string) {
	fmt.Println("projectId: ", projectId)
	fmt.Println("releaseVersion: ", releaseVersion)
	latestCommitHash := resCommits[0].Commit
	type projectInfo struct {
		projectId string
	}
	var project projectInfo
	var encryptDest []byte
	err := encrypt([]byte(projectId), encryptDest, []byte(secret), []byte(projectId))
	errorHandler(err)
	project.projectId = string(encryptDest)
	// TODO 完成解析结果输出内容功能
}
