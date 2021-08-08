package logpack

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"whosbugPack/global_type"
	"whosbugPack/utility"
)

/* GetLogInfo
/* @Description: 获取所有的git commit记录和所有的commit+diff，并返回存储的文件目录
 * @return string 所有diff信息的目录
 * @return string 所有commit信息的目录
 * @author KevinMatt 2021-07-29 17:25:39
 * @function_mark PASS
*/
func GetLogInfo() (string, string) {
	// 切换到仓库目录
	err := os.Chdir(global_type.Config.RepoPath)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Work Path Change to: ", global_type.Config.RepoPath)

	global_type.LocalHashLatest = execCommandOutput("git", "rev-parse", "HEAD")

	cloudHashLatest, err := utility.GetLatestRelease(global_type.Config.ProjectId)
	if err != nil {
		log.Println(err)
	}
	if cloudHashLatest != global_type.LocalHashLatest {
		if cloudHashLatest == "" {
			execRedirectToFile("commitInfo.out", "git", "log", "--pretty=format:%H,%ce,%cn,%cd")
			execRedirectToFile("allDiffs.out", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw")
		} else {
			execRedirectToFile("commitInfo.out", "git", "log", "--pretty=format:%H,%ce,%cn,%cd", fmt.Sprintf("%s...%s", global_type.LocalHashLatest, cloudHashLatest))
			execRedirectToFile("allDiffs.out", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw", fmt.Sprintf("%s...%s", global_type.LocalHashLatest, cloudHashLatest))
		}
	}
	return utility.ConCatStrings(global_type.WorkPath, "\\allDiffs.out"), utility.ConCatStrings(global_type.WorkPath, "\\commitInfo.out")
}

/* execCommandOutput
/* @Description: 执行命令并获取输出
 * @param command 命令
 * @param args 命令参数
 * @return string 命令输出
 * @author KevinMatt 2021-08-07 14:44:17
 * @function_mark PASS
*/
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
	fd, _ := os.OpenFile(global_type.WorkPath+"\\"+fileName, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	cmd.Stdout = fd
	cmd.Stderr = fd
	err := cmd.Start()
	if err != nil {
		log.Println(err.Error())
	}
	err = cmd.Wait()
	_ = fd.Close()
}
