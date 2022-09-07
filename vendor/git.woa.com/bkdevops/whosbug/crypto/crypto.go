package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/util"
	"git.woa.com/bkdevops/whosbug/zaplog"
	"github.com/pkg/errors"
)

// GenerateKIV 生成AES-CFB需要的Key和IV
//	@param projectID 	项目ID
//	@param key 			加密密钥
//	@return []byte 		K密钥
//	@return []byte 		IV偏移密钥
//	@author KevinMatt 2021-07-25 20:07:20
//	@function_mark PASS
func GenerateKIV(projectID, key []byte) ([]byte, []byte) {
	hK := hmac.New(sha256.New, key)
	hIV := hmac.New(md5.New, key)
	hK.Write(projectID)
	hIV.Write(projectID)
	return hK.Sum(nil), hIV.Sum(nil)
}

// Encrypt AES-CFB加密
//	@param projectID 	项目ID
//	@param Dest 			输出的加密后字符串
//	@param key 			加密密钥
//	@param plainText 	需要加密的文本
//	@return error 		错误抛出
//	@author KevinMatt 2021-07-25 13:34:09
//	@function_mark PASS
func Encrypt(projectID, key, plainText string) string {
	K, IV := GenerateKIV([]byte(projectID), []byte(key))
	aesBlockEncryptor, err := aes.NewCipher(K)
	if err != nil {
		zaplog.Logger.Error(err.Error())
	}
	var dest = []byte(plainText)
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream(dest, []byte(plainText))
	return string(dest)
}

// Decrypt AES-CFB解密
//	@param projectID 	项目ID
//	@param Dest 			解密完成的字符串
//	@param key 			解密密钥
//	@param plainText 	需要解密的文本
//	@return error 		错误抛出
//	@author KevinMatt 2021-07-25 13:35:15
//	@function_mark PASS
func Decrypt(projectID, key, plainText string) string {
	K, IV := GenerateKIV([]byte(projectID), []byte(key))
	aesBlockDescriptor, err := aes.NewCipher(K)
	if err != nil {
		zaplog.Logger.Error(err.Error())
	}
	var dest = []byte(plainText)
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream(dest, []byte(plainText))
	return string(dest)
}

// Base64Encrypt 输出base64编码的加密内容
//	@Description:
//	@param text 要加密的文本
//	@return string 加密的文本
//	@author KevinMatt 2021-08-10 01:07:41
//	@function_mark PASS
func Base64Encrypt(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(Encrypt(config.WhosbugConfig.ProjectID, config.WhosbugConfig.CryptoKey, text)))
}

// Base64Decrypt 输出base64解码后的解密内容
//  @param text string
//  @return string
//  @return error
//  @author: Kevineluo 2022-07-31 07:29:37
func Base64Decrypt(text string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	return Decrypt(config.WhosbugConfig.ProjectID, config.WhosbugConfig.CryptoKey, string(decodedBytes)), nil
}

// GenToken 生成Token
//  @return string Token字符串
//  @return error 错误信息
//  @author: Kevineluo 2022-07-31 12:52:51
func GenToken() (string, error) {
	// 拼接字符串
	var builder strings.Builder
	builder.WriteString(config.WhosbugConfig.WebServerHost)
	builder.WriteString("/v1/token")
	urls := builder.String()

	res, err := http.PostForm(urls, url.Values{"username": []string{config.WhosbugConfig.WebServerUserName}, "password": []string{config.WhosbugConfig.WebServerKey}})
	if err != nil {
		fmt.Printf("%s", util.ErrorMessage(errors.Wrapf(err, "Generate Key Failure. Check the username&password or the status of the server")))
		os.Exit(1)
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Println(errors.WithMessage(err, "Res Body Close Fails"))
		}
	}()
	if res.StatusCode == 200 {
		resBody, _ := ioutil.ReadAll(res.Body)
		res := string(resBody)
		tokenGot := strings.Split(res, "\"")[3]
		return tokenGot, nil
	} else {
		resBody, _ := ioutil.ReadAll(res.Body)
		println(string(resBody))
		return "", errors.New(string(resBody))
	}
}

// ReqWithToken 发起http请求
//  @param token string
//  @param url string
//  @param method string
//  @param data string
//  @return error 返回错误信息
//  @author: Kevineluo 2022-07-31 12:57:45
func ReqWithToken(token, url, method, data string) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return errors.Wrapf(err, "Create Request with method: %s Fails \n With data: %s", method, data)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return errors.Wrapf(err, "Sending Request with method: %s Fails\n With data: %s", method, data)
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Println(errors.WithMessage(err, "Res Body Close Fails"))
		}
	}()

	if res.StatusCode == 201 || res.StatusCode == 200 {
		return nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		temp := string(body)
		util.ForDebug(temp)
		if err != nil {
			return errors.WithMessage(err, "Read Body Fail")
		}
		return errors.New(string(body))
	}
}
