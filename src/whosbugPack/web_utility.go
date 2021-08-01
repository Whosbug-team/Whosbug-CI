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
	urls := _HOST + "/api-token-auth/"
	res, _ := http.PostForm(urls, url.Values{"username": []string{_USERNAME}, "password": []string{_PASSWORD}})
	if res.StatusCode == 200 {
		resBody, _ := ioutil.ReadAll(res.Body)
		tokenGot := strings.Split(string(resBody), "\"")[3]
		return tokenGot, nil
	} else {
		resBody, _ := ioutil.ReadAll(res.Body)
		println(string(resBody))
		return "", errors.New("got token failed")
	}
}

/** getLatestRelease
 * @Description: 获得最新的Release信息
 * @param projectId 项目ID
 * @return string Release信息
 * @author KevinMatt 2021-07-22 16:50:26
 * @function_mark PASS
 */
func getLatestRelease(projectId string) (string, error) {
	urlReq := _HOST + "/whosbug/releases/last/"
	method := "POST"

	pid := base64.StdEncoding.EncodeToString([]byte(_encrypt(projectId, _SECRET, projectId)))
	data := []byte("{\"pid\":\"" + pid + "\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	token, err := _genToken()
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", "Token "+token)
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
			fmt.Println(string(body))
			return "", err
		}
		return _decrypt(projectId, _SECRET, string(body)), nil
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
const _objectBufferQueueLength = 100

// 处理上传的协程
func processObjectUpload() {
	//object缓冲队列，满的时候再统一上传
	var objects []ObjectInfoType
	//在objectChan关闭且objectChan为空后会自然退出
	for object := range objectChan {
		if len(objects) < _objectBufferQueueLength {
			objects = append(objects, object)
		} else {
			_processUpload(objects)
		}
	}
	//自然退出后，缓冲队列可能还有残留
	_processUpload(objects)
}
func _processUpload(objects []ObjectInfoType) {
	projectId := config.ProjectId
	releaseVersion := config.ReleaseVersion
	//TODO 之后再测试对接
	err := postObjects(projectId, releaseVersion, localHashLatest, objects)
	if err != nil {
		log.Println(err)
		return
	}
}

/** getLatestRelease
 * @Description: 发送解析结果到server
 * @param projectId 项目ID
 * @param releaseVersion release版本
 * @param commitHash 本地最新commitHash
 * @param objects 待上传的object集
 * @author lxchx 2021-07-29 16:50:26
 * @function_mark PASS
 */
func postObjects(projectId string, releaseVersion string, commitHash string, objects []ObjectInfoType) error {
	//TODO 待验证能否正确发送
	//生成待发送的数据
	//为了方便,创建一个更简单的加密函数
	tempEncrypt := func(text string) string {
		return _encrypt(projectId, _SECRET, text)
	}

	projectData := fmt.Sprintf("{\"pid\":\"%s\"}", tempEncrypt(projectId))
	releaseData := fmt.Sprintf("{\"release\":\"%s\", \"commit_hash\":\"%s\"}", tempEncrypt(releaseVersion), tempEncrypt(commitHash))

	const objectFormatStr = "{\"owner\":\"%s\", \"file_path\":\"%s\", \"parent_name\":\"%s\", \"parent_hash\":\"%x\", \"name\":\"%s\", \"hash\":\"%x\", \"old_name\":\"%s\", \"commit_time\":\"%s\"}"
	//形如:
	//{
	//    "owner": "%s",
	//    "file_path": "%s",
	//    "parent_name": "%s",
	//    "parent_hash": "%x",
	//    "name": "%s",
	//    "hash": "%x",
	//    "old_name": "%s",
	//    "commit_time": "%s"
	//}
	var objectsStrForPost []string
	for _, object := range objects {
		objectStr := fmt.Sprintf(objectFormatStr,
			tempEncrypt(object["owner"]), tempEncrypt(object["file_path"]),
			tempEncrypt(object["parent_name"]), tempEncrypt(object["parent_hash"]),
			tempEncrypt(object["name"]), tempEncrypt(object["hash"]),
			tempEncrypt(object["old_name"]), object["commit_time"])
		objectsStrForPost = append(objectsStrForPost, objectStr)
	}
	objectsStr := strings.Join(objectsStrForPost, ",")
	objectsData := "[" + objectsStr + "]"

	dataForPost := fmt.Sprintf(`{"project":%s,"release":%s,"objects":%s}`, projectData, releaseData, objectsData)

	//准备发送
	urlReq := _HOST + "/whosbug/commits/diffs/"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer([]byte(dataForPost)))
	if err != nil {
		log.Println(err)
		return err
	}

	token, err := _genToken()
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Token "+token)
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
		fmt.Println(res.StatusCode)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(string(body))
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
