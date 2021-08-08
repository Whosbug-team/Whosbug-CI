package whosbugPack

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"whosbugPack/commit_diffpack"
	"whosbugPack/global_type"
	"whosbugPack/logpack"
	"whosbugPack/uploadpack"
)

/* init
/* @Description: 自动初始化配置
 * @author KevinMatt 2021-07-29 20:18:18
 * @function_mark PASS
*/
func init() {
	// 获得密钥
	global_type.Secret = os.Getenv("WHOSBUG_SECRET")
	//if secret == "" {
	//	secret = "defaultsecret"
	//}
	// 工作目录存档
	global_type.WorkPath, _ = os.Getwd()
	file, err := os.Open("src/input.json")
	if err != nil {
		log.Println(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&global_type.Config)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Get input.json succeed!")
	}
	fmt.Println("Version:\t", global_type.Config.ReleaseVersion, "\nProjectId:\t", global_type.Config.ProjectId, "\nBranchName:\t", global_type.Config.BranchName)

	global_type.ObjectChan = make(chan global_type.ObjectInfoType, 100000)
	err = os.Remove("allDiffs.out")
	if err != nil {
		log.Println(err)
	}
	err = os.Remove("commitInfo.out")
	if err != nil {
		log.Println(err)
	}
	//开启处理object上传的协程
	for i := 0; i < 5; i++ {
		go uploadpack.ProcessObjectUpload()
	}
}

// Analysis
/* @Description: 暴露给外部的函数，作为程序入口
 * @author KevinMatt 2021-07-29 17:51:28
 * @function_mark PASS
 */
func Analysis() {
	t := time.Now()
	// 获取git log命令得到的commit列表和完整的commit-diff信息存储的文件目录
	diffPath, commitPath := logpack.GetLogInfo()
	fmt.Println("Get log cost: ", time.Since(t))
	commit_diffpack.MatchCommit(diffPath, commitPath)
	// 等待关闭pool和channel
	for {
		time.Sleep(time.Millisecond * 500)
		if commit_diffpack.Pool.Running() == 0 {
			fmt.Println("Close Pool!")
			commit_diffpack.Pool.Release()
			close(global_type.ObjectChan)
			break
		}
	}
	fmt.Println("Analyse cost: ", time.Since(t))
	// 等待上传协程的结束
	uploadpack.UploadWaitGroup.Wait()
	fmt.Println("Total cost: ", time.Since(t))
}
