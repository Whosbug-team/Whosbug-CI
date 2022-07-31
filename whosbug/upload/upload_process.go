package upload

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"

	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/crypto"
	"git.woa.com/bkdevops/whosbug/git"
	"git.woa.com/bkdevops/whosbug/logging"
	"git.woa.com/bkdevops/whosbug/util"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//协程里缓存队列的长度
const ObjectBufferQueueLength = 1000

// ProcessObjectUpload
//	@Description: 处理上传的协程
//	@author KevinMatt 2021-08-10 01:50:05
//	@function_mark PASS
func ProcessObjectUpload() {
	UploadWaitGroup.Add(1)
	// object缓冲队列，满的时候再统一上传
	var objects []config.ObjectInfoType
	// 在objectChan关闭且objectChan为空后会自然退出
	for object := range config.ObjectChan {
		if len(objects) > 0 && object.Equals(objects[len(objects)-1]) {
			// 队列中没有新增，暂停100ms后维持循环
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if len(objects) < ObjectBufferQueueLength/5 {
			objects = append(objects, object)
		} else {
			objects = append(objects, object)
			processUpload(objects)
			objects = nil
		}
	}
	//自然退出后，缓冲队列可能还有残留
	if objects != nil {
		processUpload(objects)
	}
	UploadWaitGroup.Done()
	log.SetOutput(logging.LogFile)
	log.Println("Sending Finished")
}

// processUpload
//	@Description:
//	@param objects
//	@author KevinMatt 2021-08-07 16:22:27
//	@function_mark
func processUpload(objects []config.ObjectInfoType) {
	err := PostObjects(objects)
	sendCount++
	if len(objects) > 0 {
		log.SetOutput(logging.LogFile)
		log.Println("Sent count: ", objects[0].CommitHash, sendCount)
	}
	if err != nil {
		log.Println(err)
		return
	}
}

// PostObjects
//	@Description:
//	@param objects objects切片
//	@return error 返回错误信息
//	@author KevinMatt 2021-08-10 13:02:37
//	@function_mark PASS
func PostObjects(objects []config.ObjectInfoType) error {
	token, err := crypto.GenToken()
	if err != nil {
		log.Println(err)
		return err
	}

	// 使用sync池并回收变量
	dataForPost := postDataPool.Get().(*postData)
	defer postDataPool.Put(dataForPost)
	dataForPost.PostCommitInfo = postProjectInfo
	//dataForPost.Project.Pid = crypto.Base64Encrypt(config.Config.ProjectId)
	//dataForPost.Release.Release = crypto.Base64Encrypt(config.Config.ReleaseVersion)
	//dataForPost.Release.CommitHash = crypto.Base64Encrypt(config.LatestCommitHash)
	dataForPost.Objects = objects

	data, err := json.MarshalToString(&dataForPost)
	if err != nil {
		log.Println(err)
	}
	util.WriteInfoFile("/data/workspace/whosbugGolang/diffs.json", data)

	//准备发送
	urlReq := util.ConCatStrings(config.WhosbugConfig.WebServerHost, "/whosbug/commits/diffs/")
	method := "POST"

	err = crypto.ReqWithToken(token, urlReq, method, data)
	if err != nil {
		log.Println(util.ErrorMessage(err))
	}
	return err
}

// PostCommitsInfo 发送结束信息
//	@param commitPath commit文件的目录
//	@return error 返回错误
//	@author KevinMatt 2021-08-10 01:06:05
//	@function_mark PASS
func PostCommitsInfo(commitPath string) error {
	InitTheProjectStruct()
	commitFd, err := os.Open(commitPath)
	if err != nil {
		return errors.Wrap(err, "Open commitPath to Post FIN fails:")
	}
	lineReaderCommit := bufio.NewReader(commitFd)
	var FinMessage = postCommits{
		PostCommitInfo: postProjectInfo,
	}
	for {
		line, _, err := lineReaderCommit.ReadLine()
		if err == io.EOF {
			break
		}
		CommitInfo := git.GetCommitInfo(string(line))
		FinMessage.Commit = append(FinMessage.Commit, CommitInfo)
	}
	commitFd.Close()
	data, _ := json.MarshalToString(&FinMessage)
	util.WriteInfoFile("/data/workspace/whosbugGolang/commits.json", data)
	token, _ := crypto.GenToken()
	url := config.WhosbugConfig.WebServerHost + "/whosbug/commits/commits-info/"
	err = crypto.ReqWithToken(token, url, "POST", data)
	if err != nil {
		log.Println(util.ErrorMessage(err))
	}
	return err
}

// PostReleaseInfo 发送Release信息
//	@return error 错误信息
//	@author KevinMatt 2021-08-10 12:29:35
//	@function_mark PASS
func PostReleaseInfo(address string) error {
	url := config.WhosbugConfig.WebServerHost + address
	if isInitial {
		InitTheProjectStruct()
	}

	data, err := json.MarshalToString(postProjectInfo)
	if err != nil {
		return errors.Wrap(err, "json MarshalToString Fail")
	}
	token, err := crypto.GenToken()
	if err != nil {
		return errors.Wrap(err, "GenToken Fail")
	}
	err = crypto.ReqWithToken(token, url, "POST", data)
	if err != nil {
		return err
	}
	return nil
}

// InitTheProjectStruct 初始化globalType
//	@author KevinMatt 2021-08-10 12:40:28
//	@function_mark PASS
func InitTheProjectStruct() {
	postProjectInfo.Project.Pid = crypto.Base64Encrypt(config.WhosbugConfig.ProjectID)
	postProjectInfo.Release.Release = crypto.Base64Encrypt(config.WhosbugConfig.ReleaseVersion)
	postProjectInfo.Release.CommitHash = crypto.Base64Encrypt(config.LocalHashLatest)
	isInitial = false
}
