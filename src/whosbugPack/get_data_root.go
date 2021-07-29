package whosbugPack

import (
	"fmt"
	"log"
	"os"
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

	localHashLatest := execCommandOutput("git", "rev-parse", "HEAD")
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
