package whosbug

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"git.woa.com/bkdevops/whosbug/commit"
	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/logging"
	"git.woa.com/bkdevops/whosbug/upload"
	"git.woa.com/bkdevops/whosbug/util"
	"github.com/schollz/progressbar/v3"
)

func init() {
	var (
		err error
	)

	config.WorkPath, _ = os.Getwd()

	config.ObjectChan = make(chan config.ObjectInfoType, upload.ObjectBufferQueueLength)

	_, err = os.Stat(config.WorkPath + "/allDiffs.out")
	if !os.IsNotExist(err) {
		err = os.Remove(config.WorkPath + "/allDiffs.out")
		if err != nil {
			util.GLogger.Error(err.Error())
		}
	}

	_, err = os.Stat(config.WorkPath + "/commitInfo.out")
	if !os.IsNotExist(err) {
		err = os.Remove(config.WorkPath + "/commitInfo.out")
		if err != nil {
			util.GLogger.Error(err.Error())
		}
	}

	//开启处理object上传的协程
	for i := 0; i < 1; i++ {
		go upload.ProcessObjectUpload()
	}
}

// Analysis whosbug分析入口
//  @param whosbugConfig *config.Config
//  @author: Kevineluo 2022-05-02 08:03:23
func Analysis(whosbugConfig *config.Config) {
	config.WhosbugConfig = *whosbugConfig
	t := time.Now()

	// 获取git log命令得到的commit列表和完整的commit-diff信息存储的文件目录
	diffPath, commitPath := logging.GetGitLogInfo()
	util.GLogger.Infof("diffPath: %s, commitPath: %s", diffPath, commitPath)
	commit.ProcessBar = progressbar.Default(util.GetLineCount(), "Progress")
	// 指示Webservice创建新的release
	err := upload.PostReleaseInfo("/whosbug/create-project-release/")
	if err != nil {
		util.GLogger.Error(err.Error())
		if strings.Contains(err.Error(), config.AlreadyExistsError) {
			util.GLogger.Info("The Release is already exists and has the same latest commit to your repo.")
			os.Exit(0)
		}
		if os.Getenv("IS_DEBUG") == "" {
			os.Exit(-1)
		}
	}

	util.GLogger.Infof("Get git log cost: %v", time.Since(t).String())
	commit.MatchCommit(diffPath, commitPath)

	// 等待关闭pool和channel
	for {
		time.Sleep(time.Second / 10)
		if commit.AntlrAnalysisPool.Running() == 0 {
			util.GLogger.Infof("Analyse cost: %v", time.Since(t).String())
			util.GLogger.Info("Routines pool closed.")
			commit.AntlrAnalysisPool.Release()
			close(config.ObjectChan)
			break
		}
	}
	// 等待上传协程的结束
	upload.UploadWaitGroup.Wait()

	// 回收所有内存，准备转入完成上传的FIN通知
	runtime.GC()

	// 通知Webservice上传结束
	err = upload.PostReleaseInfo("/whosbug/commits/upload-done/")
	if err != nil {
		util.GLogger.Error(util.ErrorStack(err))
	}
	err = upload.PostReleaseInfo("/whosbug/commits/delete_uncalculate/")
	if err != nil {
		util.GLogger.Error(util.ErrorStack(err))
	}
	util.GLogger.Infof("Total cost: %v", time.Since(t).String())

	fmt.Println("Your ProjectName is", whosbugConfig.ProjectId, "You'll need this to en/decrypt your data")
}
