package whosbug

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	"git.woa.com/bkdevops/whosbug/commit"
	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/logging"
	"git.woa.com/bkdevops/whosbug/upload"
	"git.woa.com/bkdevops/whosbug/util"
	"git.woa.com/bkdevops/whosbug/zaplog"
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
			zaplog.Logger.Error(err.Error())
		}
	}

	_, err = os.Stat(config.WorkPath + "/commitInfo.out")
	if !os.IsNotExist(err) {
		err = os.Remove(config.WorkPath + "/commitInfo.out")
		if err != nil {
			zaplog.Logger.Error(err.Error())
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
	zaplog.Logger.Info("got git log info", zaplog.String("diffPath", diffPath), zaplog.String("commitPath", commitPath))
	commit.ProcessBar = progressbar.Default(util.GetLineCount(config.WorkPath+"/"+"commitInfo.out"), "Progress")
	// 指示Web-service创建新的release
	err := upload.PostReleaseInfo("/v1/create-project-release")
	if err != nil {
		zaplog.Logger.Error(err.Error())
		if errors.Is(err, config.AlreadyExistsError) {
			zaplog.Logger.Info("The Release is already exists and has the same latest commit to your repo.")
			os.Exit(0)
		}
		if os.Getenv("IS_DEBUG") == "" {
			os.Exit(-1)
		}
	}

	zaplog.Logger.Info("Get git log", zaplog.String("time", time.Since(t).String()))
	commit.MatchCommit(diffPath, commitPath)

	// 等待关闭pool和channel
	for {
		time.Sleep(time.Second / 10)
		if commit.AntlrAnalysisPool.Running() == 0 {
			zaplog.Logger.Info("AntlrAnalysisPool is empty", zaplog.String("cost", time.Since(t).String()))
			zaplog.Logger.Info("Routines pool closed.")
			commit.AntlrAnalysisPool.Release()
			close(config.ObjectChan)
			break
		}
	}
	// 等待上传协程的结束
	upload.UploadWaitGroup.Wait()

	// 回收所有内存，准备转入完成上传的FIN通知
	runtime.GC()

	// 通知Web-service上传结束
	err = upload.PostReleaseInfo("/v1/commits/upload-done")
	if err != nil {
		zaplog.Logger.Error(util.ErrorStack(err))
	}
	err = upload.PostReleaseInfo("/v1/commits/delete_uncalculate")
	if err != nil {
		zaplog.Logger.Error(util.ErrorStack(err))
	}
	zaplog.Logger.Info("Analysis all done", zaplog.String("cost", time.Since(t).String()))

	fmt.Println("Your ProjectName is", whosbugConfig.ProjectID, "You'll need this to en/decrypt your data")
}
