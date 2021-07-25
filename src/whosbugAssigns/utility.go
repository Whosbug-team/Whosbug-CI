package whosbugAssigns

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
)

// GetInputConfig
/* @Description:
 * @author KevinMatt 2021-07-25 18:08:42
 * @function_mark
 */
func GetInputConfig() {
	file, err := os.Open("src/input.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Get input.json succeed!")
	}
	fmt.Println("Version:\t", Config.ReleaseVersion, "\nProjectId:\t", Config.ProjectId, "\nBranchName:\t", Config.BranchName)
}

// execCommandOutput
/* @Description:
 * @param command 命令程序
 * @param args 命令参数
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

// lanFilter
/* @Description: 语言过滤器，确定目标文件是否为支持的语言
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

func Result(resCommits []CommitParsedType, projectId string, releaseVersion string) {
	//fmt.Println("projectId: ", projectId)
	//fmt.Println("releaseVersion: ", releaseVersion)
	//latestCommitHash := resCommits[0].Commit
	//type projectInfo struct {
	//	projectId string
	//}
	//var project projectInfo
	//var encryptDest []byte
	//err := encrypt([]byte(projectId), encryptDest, []byte(secret), []byte(projectId))
	//errorHandler(err)
	//project.projectId = string(encryptDest)
	fmt.Println("latestCommitHash")
	// TODO 完成解析结果输出内容功能
}

// hashCode64
/* @Description: 		返回sha256编码的拼接字符串
 * @param projectId 	项目ID
 * @param objectName
 * @param filePath 		文件目录
 * @return string 		返回编码字符串
 * @author KevinMatt 2021-07-25 14:20:09
 * @function_mark
 */
func hashCode64(projectId string, objectName string, filePath string) string {
	text := projectId + objectName + filePath
	return string(sha256.New().Sum([]byte(text)))
}
