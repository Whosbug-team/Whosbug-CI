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
	errorHandler(err, "getLatestRelease2")
	req.Header.Add("Authorization", "Token " + token)
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

func post_objects(projectId string, releaseVersion string, objects []ObjectInfoType) error {
	//TODO not finished
	return nil
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
