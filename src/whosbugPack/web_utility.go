package whosbugPack

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const _HOST = "http://127.0.0.1:8081"
const _SECRET = ""
const _USERNAME = "user"
const _PASSWORD = "pwd"

func _genToken() (string, error) {
	// 拼接字符串
	var builder strings.Builder
	builder.WriteString(_HOST)
	builder.WriteString("/api-token-auth/")
	urls := builder.String()

	res, _ := http.PostForm(urls, url.Values{"username": []string{_USERNAME}, "password": []string{_PASSWORD}})
	defer res.Body.Close()
	if res.StatusCode == 200 {
		resBody, _ := ioutil.ReadAll(res.Body)
		tokenGot := strings.Split(string(resBody), "\"")[3]
		return tokenGot, nil
	} else {
		resBody, _ := ioutil.ReadAll(res.Body)
		println(string(resBody))
		return "", errors.New(string(resBody))
	}
}

/* getLatestRelease
/* @Description:
 * @param projectId 项目ID
 * @return string Release信息
 * @return error
 * @author KevinMatt 2021-08-03 18:12:18
 * @function_mark PASS
*/
func getLatestRelease(projectId string) (string, error) {
	urlReq := conCatStrings(_HOST, "/whosbug/releases/last/")
	method := "POST"

	pid := base64.StdEncoding.EncodeToString([]byte(_encrypt(projectId, _SECRET, projectId)))
	data := []byte("{\"pid\":\"" + pid + "\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer(data))
	if err != nil {
		err = fmt.Errorf("GetLatestRelease->Sending NewRequest: %w", err)
		return "", err
	}

	token, err := _genToken()
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", conCatStrings("Token ", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			//fmt.Println(string(body))
			return "", err
		}
		fmt.Println(string(body))
		commitHash := json.Get(body, "commit_hash").ToString()
		commitHashByte, err := base64.StdEncoding.DecodeString(commitHash)
		return _decrypt(projectId, _SECRET, string(commitHashByte)), nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return "", err
		}
		fmt.Println(string(body))
		//TODO 改成相应的业务异常类型
		if res.StatusCode == 404 {
			return "", errors.New("The Project Not Found.")
		}
		return "", errors.New(string(body))
	}
}

//协程里缓存队列的长度
const _objectBufferQueueLength = 10000

// 处理上传的协程
func processObjectUpload() {
	wg.Add(1)
	// object缓冲队列，满的时候再统一上传
	var objects []objectInfoType
	// 在objectChan关闭且objectChan为空后会自然退出
	for object := range ObjectChan {
		if object == (objectInfoType{}) {
			continue
		}
		if len(objects) > 0 && object == objects[len(objects)-1] {
			continue
		}
		if len(objects) < _objectBufferQueueLength/5 {
			objects = append(objects, object)
		} else {
			objects = append(objects, object)
			_processUpload(objects)
			objects = nil
		}
	}
	//自然退出后，缓冲队列可能还有残留
	_processUpload(objects)
	wg.Done()
	fmt.Println("Sending Finished")
}

func _processUpload(objects []objectInfoType) {
	err := postObjects(localHashLatest, objects)
	sendCount++
	if len(objects) > 0 {
		fmt.Println("Sent count: ", objects[0].Hash, sendCount)
	}
	if err != nil {
		//log.Println(err)
		return
	}
}

/* postObjects
/* @Description:
 * @param projectId
 * @param releaseVersion
 * @param commitHash
 * @param objects
 * @return error
 * @author KevinMatt 2021-08-03 17:22:13
 * @function_mark PASS
*/
func postObjects(commitHash string, objects []objectInfoType) error {
	token, err := _genToken()
	if err != nil {
		log.Println(err)
		return err
	}
	tempEncrypt := func(text string) string {
		return base64.StdEncoding.EncodeToString([]byte(_encrypt(config.ProjectId, _SECRET, text)))
	}
	// 使用sync池并回收变量
	dataForPost := postDataPool.Get().(*postData)
	defer postDataPool.Put(dataForPost)

	dataForPost.Project.Pid = tempEncrypt(config.ProjectId)
	dataForPost.Release.Release = tempEncrypt(config.ReleaseVersion)
	dataForPost.Release.CommitHash = tempEncrypt(commitHash)
	for index, _ := range objects {
		//var objectForAppend objectInfoType
		objectForAppend := objectInfoPool.Get().(*objectInfoType)
		objectForAppend.Owner = tempEncrypt(objects[index].Owner)
		objectForAppend.FilePath = tempEncrypt(objects[index].FilePath)
		objectForAppend.ParentName = tempEncrypt(objects[index].ParentName)
		objectForAppend.ParentHash = tempEncrypt(objects[index].ParentHash)
		objectForAppend.Name = tempEncrypt(objects[index].Name)
		objectForAppend.Hash = tempEncrypt(objects[index].Hash)
		objectForAppend.OldName = tempEncrypt(objects[index].OldName)
		objectForAppend.CommitTime = objects[index].CommitTime
		dataForPost.Objects = append(dataForPost.Objects, *objectForAppend)
		objectInfoPool.Put(objectForAppend)
	}

	data, err := json.MarshalToString(&dataForPost)
	if err != nil {
		log.Println(err)
	}
	//准备发送
	urlReq := conCatStrings(_HOST, "/whosbug/commits/diffs/")
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer([]byte(data)))

	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Authorization", conCatStrings("Token ", token))
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

// _generateKIV
/* @Description: 		生成AES-CFB需要的Key和IV
 * @param projectId 	项目ID
 * @param key 			加密密钥
 * @return []byte 		K密钥
 * @return []byte 		IV偏移密钥
 * @author KevinMatt 2021-07-25 20:07:20
 * @function_mark PASS
 */
func _generateKIV(projectId, key []byte) ([]byte, []byte) {
	hK := hmac.New(sha256.New, key)
	hIV := hmac.New(md5.New, key)
	hK.Write(projectId)
	hIV.Write(projectId)
	return hK.Sum(nil), hIV.Sum(nil)
}

// _encrypt
/* @Description: 		AES-CFB加密
 * @param projectId 	项目ID
 * @param Dest 			输出的加密后字符串
 * @param key 			加密密钥
 * @param plainText 	需要加密的文本
 * @return error 		错误抛出
 * @author KevinMatt 2021-07-25 13:34:09
 * @function_mark PASS
 */
func _encrypt(projectId, key, plainText string) string {
	K, IV := _generateKIV([]byte(projectId), []byte(key))
	aesBlockEncryptor, err := aes.NewCipher(K)
	if err != nil {
		fmt.Println(err)
	}
	var dest = []byte(plainText)
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream(dest, []byte(plainText))
	return string(dest)
}

// _decrypt
/* @Description: 		AES-CFB解密
 * @param projectId 	项目ID
 * @param Dest 			解密完成的字符串
 * @param key 			解密密钥
 * @param plainText 	需要解密的文本
 * @return error 		错误抛出
 * @author KevinMatt 2021-07-25 13:35:15
 * @function_mark PASS
 */
func _decrypt(projectId, key, plainText string) string {
	K, IV := _generateKIV([]byte(projectId), []byte(key))
	aesBlockDescriptor, err := aes.NewCipher(K)
	if err != nil {
		fmt.Println(err)
	}
	var dest = []byte(plainText)
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream(dest, []byte(plainText))
	return string(dest)
}
