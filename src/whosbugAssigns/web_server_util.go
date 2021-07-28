package whosbugAssigns

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
	"net/http"
	"net/url"
	"strings"
)

const _HOST = "http://127.0.0.1:8081"
const _SECRET = "456"
const _USERNAME = "user"
const _PASSWORD = "pwd"

//为了让代码跑起来加的
var encrypt = _encrypt
var decrypt = _decrypt
var secret = ""

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
	errorHandler(err, "getLatestRelease")
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(body))
			return "", err
		}
		return string(body), nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		fmt.Println(string(body))
		if res.StatusCode == 404 {
			return "", errors.New("The Project Not Found.")
		}
		return "", errors.New(string(body))
	}
}

func postObjects(projectId string, releaseVersion string, commitHash string, objects []ObjectInfoType) error {
	//TODO 待验证能否正确发送
	//生成待发送的数据
	//为了方便,创建一个更简单的加密函数
	tempEncrypt := func(text string) string {
		return _encrypt(projectId, _SECRET, text)
	}
	text := "123"
	fmt.Println(tempEncrypt(text))
	fmt.Println(_encrypt(projectId, _SECRET, text))

	projectData := fmt.Sprintf("{\"pid\":\"%s\"}", tempEncrypt(projectId))
	releaseData := fmt.Sprintf("{\"release\":\"%s\", \"commit_hash\":\"%s\"}", tempEncrypt(releaseVersion), tempEncrypt(commitHash))

	const objectFormatStr = "{\"owner\":\"%s\", \"file_path\":\"%s\", \"parent_name\":\"%s\", \"parent_hash\":\"%s\", \"name\":\"%s\", \"hash\":\"%s\", \"old_name\":\"%s\", \"commit_time\":\"%s\"}"
	//形如:
	//{
	//    "owner": "%s",
	//    "file_path": "%s",
	//    "parent_name": "%s",
	//    "parent_hash": "%s",
	//    "name": "%s",
	//    "hash": "%s",
	//    "old_name": "%s",
	//    "commit_time": "%s"
	//}
	var objectsStrForPost []string
	for _, object := range objects {
		objectStr := fmt.Sprintf(objectFormatStr,
			tempEncrypt(object.Owner), tempEncrypt(object.FilePath),
			tempEncrypt(object.ParName), tempEncrypt(object.ParHash),
			tempEncrypt(object.Name), tempEncrypt(object.Hash),
			tempEncrypt(object.OldName), tempEncrypt(object.CommitTime))
		objectsStrForPost = append(objectsStrForPost, objectStr)
	}
	objectsStr := strings.Join(objectsStrForPost, ",")
	objectsData := "[" + objectsStr + "]"

	dataForPost := fmt.Sprintf(`{"project":%s,"release":%s,"objects":%s}`,projectData, releaseData, objectsData)

	//准备发送
	urlReq := _HOST + "/whosbug/commits/diffs/"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer([]byte(dataForPost)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	token, err := _genToken()
	errorHandler(err, "postObjects")
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == 201 {
		return nil
	} else {
		fmt.Println(res.StatusCode)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(string(body))
		return errors.New(string(body))
	}
}

/* _hashCode64
/* @Description: 返回sha256编码的拼接字符串
 * @param projectId 项目ID
 * @param objectName
 * @param filePath 文件目录
 * @return string 返回编码字符串
 * @author KevinMatt 2021-07-26 20:49:17
 * @function_mark
*/
func _hashCode64(projectId, objectName, filePath []byte) (text [32]byte) {
	text = sha256.Sum256(append(append(projectId, objectName...), filePath...))
	return
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
	var dest = plainText
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream([]byte(dest), []byte(plainText))
	return dest
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
	var dest = plainText
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream([]byte(dest), []byte(plainText))
	return dest
}
