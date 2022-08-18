package logging

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/crypto"
	"git.woa.com/bkdevops/whosbug/util"
	"git.woa.com/bkdevops/whosbug/zaplog"
	"github.com/go-git/go-git/v5"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// GetGitLogInfo 获取所有的git commit记录和所有的commit+diff，并返回存储的文件目录
//	@return string 所有diff信息的目录
//	@return string 所有commit信息的目录
//	@author KevinMatt 2021-07-29 17:25:39
//	@function_mark PASS
func GetGitLogInfo() (string, string) {
	// 使用 go-git 获取本地仓库的 git 提交信息
	r, err := git.PlainOpen(config.WhosbugConfig.ProjectURL)
	if err != nil {
		zaplog.Logger.Error(err.Error())
	}
	rHead, _ := r.Head()
	rHeadStr := rHead.String()
	rHeadIdx := strings.Index(rHeadStr, " ")
	config.LocalHashLatest = rHeadStr[:rHeadIdx]
	cloudHashLatest, err := GetLatestRelease(config.WhosbugConfig.ProjectID)
	if err != nil {
		if util.ErrorMessage(errors.WithStack(err)) == "404" {
			zaplog.Logger.Warn("The Project Not Found. Get all commit to Initialize")
		} else {
			zaplog.Logger.Error(util.ErrorMessage(errors.WithStack(err)))
		}
	}
	zaplog.Logger.Info("Head Got!")
	config.LatestCommitHash = cloudHashLatest
	if cloudHashLatest == config.LocalHashLatest {
		zaplog.Logger.Info("The server commit list is up-to-date.")
		os.Exit(0)
	} else {
		// 切换到仓库目录
		err := os.Chdir(config.WhosbugConfig.ProjectURL)
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		zaplog.Logger.Info("cd to work path", zaplog.String("workPath", config.WorkPath))
		if cloudHashLatest == "" {
			zaplog.Logger.Info("Start Getting log")
			err := ExecRedirectToFile("", "git", "log", "--pretty=format:%H,%ce,%cn,%cd", "-n 10000", fmt.Sprint("--output=", config.WorkPath, "/commitInfo.out"))
			if err != nil {
				zaplog.Logger.Error(util.ErrorStack(err))
			}
			err = ExecRedirectToFile("", "git", "log", "--full-diff", "-p", "-U10000", "--pretty=raw", "-n 10000", fmt.Sprint("--output=", config.WorkPath, "/allDiffs.out"))
			if err != nil {
				zaplog.Logger.Error(util.ErrorStack(err))
			}
		} else {
			err := ExecRedirectToFile("", "git", "log", "--pretty=format:%H,%ce,%cn,%cd", "-n 10000", fmt.Sprintf("%s...%s", config.LocalHashLatest, cloudHashLatest), fmt.Sprint("--output=", config.WorkPath, "/commitInfo.out"))
			if err != nil {
				zaplog.Logger.Error(util.ErrorStack(err))
			}
			err = ExecRedirectToFile("", "git", "log", "--full-diff", "-p", "-U10000", "-n 10000", "--pretty=raw", fmt.Sprintf("%s...%s", config.LocalHashLatest, cloudHashLatest), fmt.Sprint("--output=", config.WorkPath, "/allDiffs.out"))
			if err != nil {
				zaplog.Logger.Error(util.ErrorStack(err))
			}
		}
	}
	return util.ConCatStrings(config.WorkPath, "/allDiffs.out"), util.ConCatStrings(config.WorkPath, "/commitInfo.out")
}

// ExecCommandOutput 执行命令并获取输出
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
		zaplog.Logger.Error(err.Error())
		log.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		zaplog.Logger.Error(err.Error())
		log.Println(err)
	}
	return out.String()
}

// ExecRedirectToFile 执行命令并将输出流重定向到目标文件中
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

// GetLatestRelease 获取最新的release
//  @param projectID string
//  @return string Release信息
//  @return error
//  @author: Kevineluo 2022-07-31 01:03:27
func GetLatestRelease(projectID string) (string, error) {
	urlReq := util.ConCatStrings(config.WhosbugConfig.WebServerHost, "/v1/releases/last")
	method := "POST"

	pid := crypto.Base64Encrypt(projectID)
	data := []byte("{\"pid\":\"" + pid + "\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer(data))

	if err != nil {
		return "", errors.Wrapf(err, "GetLatestRelease->Sending NewRequest")
	}

	token, err := crypto.GenToken()
	if err != nil {
		return "", errors.WithStack(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		encryptedCommitHash := json.Get(body, "last_commit_hash").ToString()
		commitHash, err := crypto.Base64Decrypt(encryptedCommitHash)
		if err != nil {
			return "", errors.WithStack(err)
		}
		return commitHash, nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		if res.StatusCode == 404 {
			return "", errors.New("404")
		}
		return "", errors.New(string(body))
	}
}
