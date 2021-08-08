package uploadpack

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"whosbugPack/global_type"
	"whosbugPack/logpack"
	"whosbugPack/utility"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//协程里缓存队列的长度
const _objectBufferQueueLength = 10000

func ProcessLargeObjectUpload() {
	fmt.Println("Sending Large start")
	UploadWaitGroup.Add(1)
	// object缓冲队列，满的时候再统一上传
	var objects []global_type.ObjectInfoType
	// 在objectChan关闭且objectChan为空后会自然退出
	for object := range global_type.ObjectChanLarge {
		if object == (global_type.ObjectInfoType{}) {
			continue
		}
		if len(objects) > 0 && object == objects[len(objects)-1] {
			continue
		}
		if len(objects) < 100000 {
			objects = append(objects, object)
		} else {
			objects = append(objects, object)
			_processUpload(objects)
			objects = nil
		}
	}
	UploadWaitGroup.Done()
	fmt.Println("Sending Large Object Finished")
}

// 处理上传的协程
func ProcessObjectUpload() {
	UploadWaitGroup.Add(1)
	// object缓冲队列，满的时候再统一上传
	var objects []global_type.ObjectInfoType
	// 在objectChan关闭且objectChan为空后会自然退出
	for object := range global_type.ObjectChan {
		if object == (global_type.ObjectInfoType{}) {
			continue
		}
		if len(objects) > 0 && object == objects[len(objects)-1] {
			continue
		}
		if len(objects) < _objectBufferQueueLength {
			objects = append(objects, object)
		} else {
			objects = append(objects, object)
			_processUpload(objects)
			objects = nil
		}
	}
	//自然退出后，缓冲队列可能还有残留
	_processUpload(objects)
	UploadWaitGroup.Done()
	log.SetOutput(logpack.LogFile)
	log.Println("Sending Finished")
}

/* _processUpload
/* @Description:
 * @param objects
 * @author KevinMatt 2021-08-07 16:22:27
 * @function_mark
*/
func _processUpload(objects []global_type.ObjectInfoType) {
	err := PostObjects(global_type.LocalHashLatest, objects)
	sendCount++
	if len(objects) > 0 {
		log.SetOutput(logpack.LogFile)
		log.Println("Sent count: ", objects[0].Hash, sendCount)
	}
	if err != nil {
		//log.Println(err)
		return
	}
}

/* PostObjects
/* @Description:
 * @param projectId
 * @param releaseVersion
 * @param commitHash
 * @param objects
 * @return error
 * @author KevinMatt 2021-08-03 17:22:13
 * @function_mark PASS
*/
func PostObjects(commitHash string, objects []global_type.ObjectInfoType) error {
	token, err := utility.GenToken()
	if err != nil {
		log.Println(err)
		return err
	}
	tempEncrypt := func(text string) string {
		return base64.StdEncoding.EncodeToString([]byte(utility.Encrypt(global_type.Config.ProjectId, global_type.Config.CryptoKey, text)))
	}

	// 使用sync池并回收变量
	dataForPost := postDataPool.Get().(*postData)
	defer postDataPool.Put(dataForPost)

	dataForPost.Project.Pid = tempEncrypt(global_type.Config.ProjectId)
	dataForPost.Release.Release = tempEncrypt(global_type.Config.ReleaseVersion)
	dataForPost.Release.CommitHash = tempEncrypt(commitHash)
	dataForPost.Objects = objects

	data, err := json.MarshalToString(&dataForPost)
	if err != nil {
		log.Println(err)
	}
	//准备发送
	urlReq := utility.ConCatStrings(global_type.Config.WebServerHost, "/whosbug/commits/diffs/")
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer([]byte(data)))

	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Authorization", utility.ConCatStrings("Token ", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()
	if res.StatusCode == 201 {
		return nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return err
		}
		return errors.New(string(body))
	}
}
