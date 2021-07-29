package whosbugPack

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

// 月份转换Map
var monthCorrespond = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

func execCommandOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Println(err)
	}
	return out.String()
}

/* execRedirectToFile
/* @Description: 执行命令并将输出流重定向到目标文件中
 * @param fileName 目标文件目录
 * @param command 执行的指令头
 * @param args 执行指令的参数
 * @author KevinMatt 2021-07-29 17:31:00
 * @function_mark PASS
*/
func execRedirectToFile(fileName string, command string, args ...string) {
	cmd := exec.Command(command, args...)
	fmt.Println("Cmd", cmd.Args)
	fd, _ := os.OpenFile(workPath+"\\"+fileName, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	cmd.Stdout = fd
	cmd.Stderr = fd
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	err = cmd.Wait()
	_ = fd.Close()
}

// toIso8601
/* @Description: 时间戳转换
 * @param timeList
 * @return string
 * @author KevinMatt 2021-07-25 13:42:29
 * @function_mark PASS
 */
func toIso8601(timeList []string) string {
	return fmt.Sprintf("%s-%s-%sT%s%s:%s", timeList[3], monthCorrespond[timeList[0]], timeList[1], timeList[2], timeList[4][3:], timeList[4][3:])
}

/* lanFilter
/* @Description: 语言过滤器，确定目标文件是否为支持的语言
 * @param fileName 文件名
 * @return bool 是否支持语言
 * @author KevinMatt 2021-07-26 20:48:57
 * @function_mark PASS
*/
func lanFilter(fileName string) bool {
	for index := range supportLans {
		if path.Ext(fileName) == supportLans[index] {
			return true
		}
	}
	return false
}

/* hashCode64
/* @Description: 返回sha256编码的拼接32位byte编码值
 * @param projectId 项目ID
 * @param objectName
 * @param filePath 文件目录
 * @return string 返回32位byte编码字符串
 * @author KevinMatt 2021-07-26 20:49:17
 * @function_mark PASS
*/
func hashCode64(projectId, objectName, filePath []byte) (text [32]byte) {
	text = sha256.Sum256(append(append(projectId, objectName...), filePath...))
	return
}

// forDebug 为了在debug时方便的装入未使用的参数避免出现编译错误的工具
func forDebug(a ...interface{}) {

}
