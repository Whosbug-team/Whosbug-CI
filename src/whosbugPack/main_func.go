package whosbugPack

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
	"time"
	"whosbugPack/commit_diffpack"
	"whosbugPack/global_type"
	"whosbugPack/logpack"
	"whosbugPack/uploadpack"
	"whosbugPack/util"
)

// init
//	@Description: 自动初始化
//	@author KevinMatt 2021-07-29 20:18:18
//	@function_mark PASS
//
func init() {
	var (
		err            error
		inputJsonBytes []byte
	)
	util.GLogger.Info("Initing whosbug...")
	// 工作目录存档
	global_type.WorkPath, _ = os.Getwd()
	inputJsonBytes, err = ioutil.ReadFile("./input.json")
	if err != nil {
		util.GLogger.Emergency(err.Error())
		os.Exit(-1)
	}

	err = json.Unmarshal(inputJsonBytes, &global_type.Config)
	if err != nil {
		util.GLogger.Emergency(err.Error())
		os.Exit(-1)
	} else {
		util.GLogger.Info("Get input-config succeed!")
	}

	// 打印插件版本信息
	util.GLogger.Infof("\nVersion:\t%s\nProjectId:\t%s\nBranchName:\t%s", global_type.Config.ReleaseVersion, global_type.Config.ProjectId, global_type.Config.BranchName)

	global_type.ObjectChan = make(chan global_type.ObjectInfoType, 10000)

	_, err = os.Stat(global_type.WorkPath + "/allDiffs.out")
	if !os.IsNotExist(err) {
		err = os.Remove(global_type.WorkPath + "/allDiffs.out")
		if err != nil {
			util.GLogger.Error(err.Error())
		}
	}

	_, err = os.Stat(global_type.WorkPath + "/commitInfo.out")
	if !os.IsNotExist(err) {
		err = os.Remove(global_type.WorkPath + "/commitInfo.out")
		if err != nil {
			util.GLogger.Error(err.Error())
		}
	}

	//开启处理object上传的协程
	for i := 0; i < 1; i++ {
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
	diffPath, commitPath := logpack.GetGitLogInfo()
	util.GLogger.Infof("diffPath: %s, commitPath: %s", diffPath, commitPath)
	// 指示Webservice创建新的release
	err := uploadpack.PostReleaseInfo("/whosbug/create-project-release/")
	if err != nil {
		util.GLogger.Error(err.Error())
	}

	util.GLogger.Infof("Get log cost: %d", time.Since(t))
	commit_diffpack.MatchCommit(diffPath, commitPath)

	// 等待关闭pool和channel
	for {
		time.Sleep(time.Second / 10)
		if commit_diffpack.Pool.Running() == 0 {
			util.GLogger.Infof("Analyse cost: %d", time.Since(t))
			util.GLogger.Info("Routines pool closed.")
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
	err = uploadpack.PostReleaseInfo("/whosbug/commits/upload-done/")
	if err != nil {
		util.GLogger.Error(util.ErrorStack(err))
	}
	util.GLogger.Infof("Total cost: %d", time.Since(t))
}
