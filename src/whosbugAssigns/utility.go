package whosbugAssigns

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

// GetInputConfig
/* @Description: 获取Input.json的参数
 * @author KevinMatt 2021-07-26 20:49:50
 * @function_mark PASS
 */
func GetInputConfig() {
	t := time.Now()
	file, err := os.Open("src/input.json")
	errorHandler(err)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Get input.json succeed!")
	}
	fmt.Println("Version:\t", Config.ReleaseVersion, "\nProjectId:\t", Config.ProjectId, "\nBranchName:\t", Config.BranchName)
	fmt.Println("GetInputConfig time cost: ", time.Since(t))
}

/* execCommandOutput
/* @Description:
 * @param command 命令程序
 * @param args 命令参数
 * @return string 标准输出内容
 * @author KevinMatt 2021-07-26 20:48:55
 * @function_mark PASS
*/
func execCommandOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	errorHandler(err)
	err = cmd.Wait()
	return out.String()
}

func execCommandOutput1(fileName string, command string, args ...string) {
	cmd := exec.Command(command, args...)
	fmt.Println("Cmd", cmd.Args)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = os.Stderr

	fd, _ := os.OpenFile(workPath+"\\"+fileName, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	cmd.Stdout = fd
	cmd.Stderr = fd
	err := cmd.Start()
	errorHandler(err)
	err = cmd.Wait()
	_ = fd.Close()
}

// 测试用的文本文件11M大小
var m11 string = `G:\runtime\log\ccapi\11M.log`

// 测试用的文本文件400M大小
var m400 string = `G:\runtime\log\ccapi\400M.log`

// 读取文件的每一行
func readEachLineReader(filePath string) {
	start1 := time.Now()
	FileHandle, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()
	lineReader := bufio.NewReader(FileHandle)
	for {
		// 相同使用场景下可以采用的方法
		// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		// func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
		// func (b *Reader) ReadString(delim byte) (line string, err error)
		line, _, err := lineReader.ReadLine()
		if err == io.EOF {
			break
		}
		// 如下是某些业务逻辑操作
		// 如下代码打印每次读取的文件行内容
		fmt.Println(string(line))
	}
	fmt.Println("readEachLineReader spend : ", time.Now().Sub(start1))
}

//func execCommandOutput(command string, args ...string) string {
//	cmd := exec.Command(command, args...)
//	output := bytes.Buffer{}
//	cmd.Stdout = &output
//	err := cmd.Run()
//	errorHandler(err, "exec command ", command)
//	return output.String()
//}

/* lanFilter
/* @Description: 语言过滤器，确定目标文件是否为支持的语言
 * @param fileName 文件名
 * @return bool 是否支持
 * @author KevinMatt 2021-07-26 20:48:57
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

/* getKeysAndValues
/* @Description:
 * @param m
 * @return []int
 * @return []map[string]string
 * @author KevinMatt 2021-07-26 20:49:01
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
 * @author KevinMatt 2021-07-26 20:49:04
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
					"commit_time": commits.CommitTime,
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

	fd, err := os.OpenFile("res.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	errorHandler(err)
	jsonInfo, err := json.Marshal(res)
	_, _ = fd.WriteString(string(jsonInfo))
	err = fd.Close()
}

/* hashCode64
/* @Description: 返回sha256编码的拼接32位byte编码值
 * @param projectId 项目ID
 * @param objectName
 * @param filePath 文件目录
 * @return string 返回32位byte编码字符串
 * @author KevinMatt 2021-07-26 20:49:17
 * @function_mark
*/
func hashCode64(projectId, objectName, filePath []byte) (text [32]byte) {
	text = sha256.Sum256(append(append(projectId, objectName...), filePath...))
	return
}
