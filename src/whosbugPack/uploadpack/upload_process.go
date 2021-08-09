package uploadpack

import (
	"bufio"
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"whosbugPack/commit_diffpack"
	"whosbugPack/global_type"
	"whosbugPack/logpack"
	"whosbugPack/utility"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//协程里缓存队列的长度
const _objectBufferQueueLength = 10000

//func ProcessLargeObjectUpload() {
//	fmt.Println("Sending Large start")
//	UploadWaitGroup.Add(1)
//	// object缓冲队列，满的时候再统一上传
//	var objects []global_type.ObjectInfoType
//	// 在objectChan关闭且objectChan为空后会自然退出
//	for object := range global_type.ObjectChanLarge {
//		if object.Equals(global_type.ObjectInfoType{}) {
//			continue
//		}
//		if len(objects) > 0 && object.Equals(objects[len(objects)-1]) {
//			continue
//		}
//		if len(objects) < 100000 {
//			objects = append(objects, object)
//		} else {
//			objects = append(objects, object)
//			processUpload(objects)
//			objects = nil
//		}
//	}
//	UploadWaitGroup.Done()
//	fmt.Println("Sending Large Object Finished")
//}
/* ProcessObjectUpload
/* @Description: 处理上传的协程
 * @author KevinMatt 2021-08-10 01:50:05
 * @function_mark PASS
*/
func ProcessObjectUpload() {
	UploadWaitGroup.Add(1)
	// object缓冲队列，满的时候再统一上传
	var objects []global_type.ObjectInfoType
	// 在objectChan关闭且objectChan为空后会自然退出
	for object := range global_type.ObjectChan {
		if object.Equals(global_type.ObjectInfoType{}) {
			continue
		}
		if len(objects) > 0 && object.Equals(objects[len(objects)-1]) {
			continue
		}
		if len(objects) < _objectBufferQueueLength {
			objects = append(objects, object)
		} else {
			objects = append(objects, object)
			processUpload(objects)
			objects = nil
		}
	}
	//自然退出后，缓冲队列可能还有残留
	processUpload(objects)
	UploadWaitGroup.Done()
	log.SetOutput(logpack.LogFile)
	log.Println("Sending Finished")
}

/* processUpload
/* @Description:
 * @param objects
 * @author KevinMatt 2021-08-07 16:22:27
 * @function_mark
*/
func processUpload(objects []global_type.ObjectInfoType) {
	err := PostObjects(objects)
	sendCount++
	if len(objects) > 0 {
		log.SetOutput(logpack.LogFile)
		log.Println("Sent count: ", objects[0].CommitHash, sendCount)
	}
	if err != nil {
		//log.Println(err)
		return
	}
}

/* PostObjects
/* @Description: 发送objects
 * @param objects objects切片
 * @return error 返回错误信息
 * @author KevinMatt 2021-08-10 01:07:02
 * @function_mark PASS
*/
func PostObjects(objects []global_type.ObjectInfoType) error {
	token, err := utility.GenToken()
	if err != nil {
		log.Println(err)
		return err
	}

	// 使用sync池并回收变量
	dataForPost := postDataPool.Get().(*postData)
	defer postDataPool.Put(dataForPost)

	dataForPost.Project.Pid = utility.Base64Encrypt(global_type.Config.ProjectId)
	dataForPost.Release.Release = utility.Base64Encrypt(global_type.Config.ReleaseVersion)
	dataForPost.Release.CommitHash = utility.Base64Encrypt(global_type.LatestCommitHash)
	dataForPost.Objects = objects

	data, err := json.MarshalToString(&dataForPost)

	if err != nil {
		log.Println(err)
	}

	//准备发送
	urlReq := utility.ConCatStrings(global_type.Config.WebServerHost, "/whosbug/commits/diffs/")
	method := "POST"

	err = ReqWithToken(token, urlReq, method, data)
	if err != nil {
		log.Println(utility.ErrorMessage(err))
	}
	return err
}

/* PostCommitsInfo
/* @Description: 发送结束信息
 * @param commitPath commit文件的目录
 * @return error 返回错误
 * @author KevinMatt 2021-08-10 01:06:05
 * @function_mark PASS
*/
func PostCommitsInfo(commitPath string) error {
	commitFd, err := os.Open(commitPath)
	if err != nil {
		return errors.Wrap(err, "Open commitPath to Post FIN fails:")
	}
	lineReaderCommit := bufio.NewReader(commitFd)
	var FinMessage postCommits
	FinMessage.Project.Pid = utility.Base64Encrypt(global_type.Config.ProjectId)
	FinMessage.Release.Release = utility.Base64Encrypt(global_type.Config.ReleaseVersion)
	FinMessage.Release.CommitHash = utility.Base64Encrypt(global_type.LocalHashLatest)
	for {
		line, _, err := lineReaderCommit.ReadLine()
		if err == io.EOF {
			break
		}
		CommitInfo := commit_diffpack.GetCommitInfo(string(line))
		FinMessage.Commit = append(FinMessage.Commit, CommitInfo)
	}

	data, err := json.MarshalToString(&FinMessage)
	token, err := utility.GenToken()
	url := global_type.Config.WebServerHost + "/whosbug/commits/commits-info/"
	err = ReqWithToken(token, url, "POST", data)
	if err != nil {
		log.Println(utility.ErrorMessage(err))
	}
	return err
}

/* ReqWithToken
/* @Description: 发起http请求
 * @param token 生成的token
 * @param url 请求的url
 * @param method 请求方法
 * @param data 请求带有的数据
 * @return error 返回错误信息
 * @author KevinMatt 2021-08-10 00:47:48
 * @function_mark PASS
*/
func ReqWithToken(token, url, method, data string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return errors.Wrapf(err, "Create Request with method: %s Fails \n With data: %s", method, data)
	}
	req.Header.Add("Authorization", utility.ConCatStrings("Token ", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return errors.Wrapf(err, "Sending Request with method: %s Fails\n With data: %s", method, data)
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Println(errors.WithMessage(err, "Res Body Close Fails!"))
		}
	}()

	if res.StatusCode == 201 {
		return nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.WithMessage(err, "Read Body Fail")
		}
		return errors.New(string(body))
	}
}

func PostFin() error {
	url := global_type.Config.WebServerHost + "/whosbug/commits/upload-done"
	var FinMessage postCommits
	FinMessage.Project.Pid = utility.Base64Encrypt(global_type.Config.ProjectId)
	FinMessage.Release.Release = utility.Base64Encrypt(global_type.Config.ReleaseVersion)
	FinMessage.Release.CommitHash = utility.Base64Encrypt(global_type.LocalHashLatest)
	data, err := json.MarshalToString(FinMessage)
	if err != nil {
		return errors.Wrap(err, "json MarshalToString Fail")
	}
	token, err := utility.GenToken()
	if err != nil {
		return errors.Wrap(err, "GenToken Fail")
	}
	err = ReqWithToken(token, url, "POST", data)
	if err != nil {
		return err
	}
	return nil
}
