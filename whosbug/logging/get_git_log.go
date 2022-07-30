package logging

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/util"

	"github.com/pkg/errors"
)

// GetLogInfo
//	@Description: 获取所有的git commit记录和所有的commit+diff，并返回存储的文件目录
//	@return string 所有diff信息的目录
//	@return string 所有commit信息的目录
//	@author KevinMatt 2021-07-29 17:25:39
//	@function_mark PASS
func GetGitLogInfo() (string, string) {
	// 切换到仓库目录
	err := os.Chdir(config.WhosbugConfig.ProjectUrl)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	util.GLogger.Infof("Work Path In %v", config.WorkPath)

	config.LocalHashLatest = ExecCommandOutput("git", "rev-parse", "HEAD")
	config.LocalHashLatest = config.LocalHashLatest[0 : len(config.LocalHashLatest)-1]
	cloudHashLatest, err := util.GetLatestRelease(config.WhosbugConfig.ProjectId)
	if err != nil {
		util.GLogger.Error(util.ErrorMessage(errors.WithStack(err)))
	}
	util.GLogger.Info("Head Got!")
	config.LatestCommitHash = cloudHashLatest
	if cloudHashLatest == config.LocalHashLatest {
		util.GLogger.Info("The server commit list is up-to-date.")
		os.Exit(0)
	} else {
		if cloudHashLatest == "" {
			util.GLogger.Info("Start Getting log")
			err := ExecRedirectToFile("", "git", "log", "--pretty=format:%H,%ce,%cn,%cd", "-n 10000", fmt.Sprint("--output=", config.WorkPath, "/commitInfo.out"))
			if err != nil {
				util.GLogger.Error(util.ErrorStack(err))
			}
			err = ExecRedirectToFile("", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw", "-n 10000", fmt.Sprint("--output=", config.WorkPath, "/allDiffs.out"))
			if err != nil {
				util.GLogger.Error(util.ErrorStack(err))
			}
		} else {
			err := ExecRedirectToFile("", "git", "log", "--pretty=format:%H,%ce,%cn,%cd", "-n 10000", fmt.Sprintf("%s...%s", config.LocalHashLatest, cloudHashLatest), fmt.Sprint("--output=", config.WorkPath, "/commitInfo.out"))
			if err != nil {
				util.GLogger.Error(util.ErrorStack(err))
			}
			err = ExecRedirectToFile("", "git", "log", "--full-diff", "-p", "-U10000", "-n 10000", "--pretty=raw", fmt.Sprintf("%s...%s", config.LocalHashLatest, cloudHashLatest), fmt.Sprint("--output=", config.WorkPath, "/allDiffs.out"))
			if err != nil {
				util.GLogger.Error(util.ErrorStack(err))
			}
		}
	}
	return util.ConCatStrings(config.WorkPath, "/allDiffs.out"), util.ConCatStrings(config.WorkPath, "/commitInfo.out")
}

// ExecCommandOutput
//	@Description: 执行命令并获取输出
//	@param command 命令
//	@param args 命令参数
//	@return string 命令输出
//	@author KevinMatt 2021-08-07 14:44:17
//	@function_mark PASS
func ExecCommandOutput(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	log.SetOutput(LogFile)
	log.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		util.GLogger.Error(err.Error())
		log.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		util.GLogger.Error(err.Error())
		log.Println(err)
	}
	return out.String()
}

// ExecRedirectToFile
//	@Description: 执行命令并将输出流重定向到目标文件中
//	@param fileName 目标文件目录
//	@param command 执行的指令头
//	@param args 执行指令的参数
//	@author KevinMatt 2021-07-29 17:31:00
//	@function_mark PASS
func ExecRedirectToFile(fileName string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	log.SetOutput(LogFile)
	log.Println("Cmd", cmd.Args)
	if fileName != "" {
		fd, _ := os.OpenFile(config.WorkPath+"/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
		defer fd.Close()
		cmd.Stdout = fd
		cmd.Stderr = fd
	}
	err := cmd.Start()
	if err != nil {
		return errors.Wrap(err, "Start cmd Fails.")
	}
	err = cmd.Wait()
	if err != nil {
		return errors.Wrap(err, "cmd Wait Fails.")
	}
	return err
}
