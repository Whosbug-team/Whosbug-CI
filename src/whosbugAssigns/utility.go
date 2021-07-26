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

// getKeysAndValues
/* @Description:
 * @param m
 * @return []int
 * @return []map[string]string
 * @author KevinMatt 2021-07-26 17:03:05
 * @function_mark
 */
func getKeysAndValues(m map[int]map[string]string) ([]int, []map[string]string) {
	keys := make([]int, 0, len(m))
	values := make([]map[string]string, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// Result
/* @Description:
 * @param resCommits
 * @param projectId
 * @param releaseVersion
 * @author KevinMatt 2021-07-26 17:03:02
 * @function_mark
 */
func Result(resCommits []CommitParsedType, projectId string, releaseVersion string) {
	fmt.Println("Pid: ", projectId)
	fmt.Println("Release: ", releaseVersion)
	latestCommitHash := resCommits[0].Commit

	project := map[string]string{
		"pid": encrypt(projectId, secret, projectId),
	}

	release := map[string]string{
		"release":     encrypt(projectId, secret, releaseVersion),
		"commit_hash": encrypt(projectId, secret, latestCommitHash),
	}

	objects := make([]map[string]string, 0)
	for _, commits := range resCommits {
		owner := fmt.Sprintf("%s-%s", commits.CommitterInfo.Name, commits.CommitterInfo.Email)
		for _, diffFile := range commits.CommitDiffs {
			filePath := path.Base(diffFile.DiffFilePath)
			for _, value := range diffFile.DiffContent {
				tempMap := map[string]string{
					"owner":       encrypt(projectId, secret, owner),
					"file_path":   encrypt(projectId, secret, filePath),
					"parent_name": encrypt(projectId, secret, value["parent_name"]),
					"parent_hash": encrypt(projectId, secret, value["parent_hash"]),
					"name":        encrypt(projectId, secret, value["name"]),
					"hash":        encrypt(projectId, secret, value["hash"]),
					"old_name":    "",
					"commit_time": encrypt(projectId, secret, value["commit_time"]),
				}
				objects = append(objects, tempMap)
			}
		}
	}
	res := map[string]interface{}{
		"objects": objects,
		"release": release,
		"project": project,
	}

	fd, err := os.OpenFile("res.json", os.O_RDWR|os.O_CREATE, 0755)
	errorHandler(err)
	jsonInfo, err := json.Marshal(res)
	_, _ = fd.WriteString(string(jsonInfo))

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
