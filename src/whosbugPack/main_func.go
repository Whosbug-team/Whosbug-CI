package whosbugPack

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"runtime"
	"time"
	"whosbugPack/commit_diffpack"
	"whosbugPack/global_type"
	"whosbugPack/logpack"
	"whosbugPack/uploadpack"
	"whosbugPack/utility"
)

// init
//	@Description: 自动初始化
//	@author KevinMatt 2021-07-29 20:18:18
//	@function_mark PASS
//
func init() {
	// 获得密钥
	global_type.Config.CryptoKey = os.Getenv("WHOSBUG_SECRET")
	// 工作目录存档
	global_type.WorkPath, _ = os.Getwd()
	file, err := os.Open("src/input.json")
	if err != nil {
		fmt.Println(utility.ErrorMessage(err))
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&global_type.Config)
	if err != nil {
		log.Println(utility.ErrorMessage(err))
	} else {
		fmt.Println("Get input-config succeed!")
	}

	// 打印插件版本信息
	fmt.Println("Version:\t", global_type.Config.ReleaseVersion, "\nProjectId:\t", global_type.Config.ProjectId, "\nBranchName:\t", global_type.Config.BranchName)

	global_type.ObjectChan = make(chan global_type.ObjectInfoType, 100000)

	_, err = os.Stat("allDiffs.out")
	if !os.IsNotExist(err) {
		err = os.Remove("allDiffs.out")
		if err != nil {
			log.Println(utility.ErrorMessage(errors.WithStack(err)))
		}
	}

	_, err = os.Stat("commitInfo.out")
	if !os.IsNotExist(err) {
		err = os.Remove("commitInfo.out")
		if err != nil {
			log.Println(utility.ErrorMessage(errors.WithStack(err)))
		}
	}

	//开启处理object上传的协程
	for i := 0; i < 5; i++ {
		go uploadpack.ProcessObjectUpload()
	}
}

// Analysis
//	@Description: 暴露给外部的函数，作为程序入口
//	@author KevinMatt 2021-07-29 17:51:28
//	@function_mark PASS
func Analysis() {
	t := time.Now()

	// 获取git log命令得到的commit列表和完整的commit-diff信息存储的文件目录
	diffPath, commitPath := logpack.GetLogInfo()

	// 指示Webservice创建新的release
	err := uploadpack.PostReleaseInfo("/whosbug/create-project-release")
	if err != nil {
		log.Println(utility.ErrorStack(err))
	}

	fmt.Println("Get log cost: ", time.Since(t))
	commit_diffpack.MatchCommit(diffPath, commitPath)

	// 等待关闭pool和channel
	for {
		time.Sleep(time.Second / 10)
		if commit_diffpack.Pool.Running() == 0 {
			fmt.Println("Analyse cost: ", time.Since(t))
			fmt.Println("Routines pool closed.")
			commit_diffpack.Pool.Release()
			close(global_type.ObjectChan)
			break
		}
	}
	// 等待上传协程的结束
	uploadpack.UploadWaitGroup.Wait()

	// 回收所有内存，准备转入完成上传的FIN通知
	runtime.GC()

	// 通知Webservice上传结束
	err = uploadpack.PostReleaseInfo("/whosbug/commits/upload-done")
	if err != nil {
		log.Println(utility.ErrorStack(err))
	}
	fmt.Println("Total cost: ", time.Since(t))
}
