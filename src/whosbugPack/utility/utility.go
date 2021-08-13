package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"whosbugPack/global_type"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

// 月份转换Map
var monthCorrespond = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

// ToIso8601
//	@Description: 时间戳转换
//	@param timeList
//	@return string
//	@author KevinMatt 2021-07-25 13:42:29
//	@function_mark PASS
func ToIso8601(timeList []string) string {
	temp := fmt.Sprintf("%s-%s-%sT%s+%s:%s", timeList[3], monthCorrespond[timeList[0]], timeList[1], timeList[2], timeList[4][1:3], timeList[4][3:])
	return temp
}

// GenerateKIV
//	@Description: 		生成AES-CFB需要的Key和IV
//	@param projectId 	项目ID
//	@param key 			加密密钥
//	@return []byte 		K密钥
//	@return []byte 		IV偏移密钥
//	@author KevinMatt 2021-07-25 20:07:20
//	@function_mark PASS
func GenerateKIV(projectId, key []byte) ([]byte, []byte) {
	hK := hmac.New(sha256.New, key)
	hIV := hmac.New(md5.New, key)
	hK.Write(projectId)
	hIV.Write(projectId)
	return hK.Sum(nil), hIV.Sum(nil)
}

// Encrypt
//	@Description: 		AES-CFB加密
//	@param projectId 	项目ID
//	@param Dest 			输出的加密后字符串
//	@param key 			加密密钥
//	@param plainText 	需要加密的文本
//	@return error 		错误抛出
//	@author KevinMatt 2021-07-25 13:34:09
//	@function_mark PASS
func Encrypt(projectId, key, plainText string) string {
	K, IV := GenerateKIV([]byte(projectId), []byte(key))
	aesBlockEncryptor, err := aes.NewCipher(K)
	if err != nil {
		fmt.Println(err)
	}
	var dest = []byte(plainText)
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream(dest, []byte(plainText))
	return string(dest)
}

// Decrypt
//	@Description: 		AES-CFB解密
//	@param projectId 	项目ID
//	@param Dest 			解密完成的字符串
//	@param key 			解密密钥
//	@param plainText 	需要解密的文本
//	@return error 		错误抛出
//	@author KevinMatt 2021-07-25 13:35:15
//	@function_mark PASS
func Decrypt(projectId, key, plainText string) string {
	K, IV := GenerateKIV([]byte(projectId), []byte(key))
	aesBlockDescriptor, err := aes.NewCipher(K)
	if err != nil {
		fmt.Println(err)
	}
	var dest = []byte(plainText)
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream(dest, []byte(plainText))
	return string(dest)
}

// GenToken
//	@Description: 生成Token
//	@return string Token字符串
//	@return error 错误信息
//	@author KevinMatt 2021-08-08 21:35:26
//	@function_mark PASS
func GenToken() (string, error) {
	// 拼接字符串
	var builder strings.Builder
	builder.WriteString(global_type.Config.WebServerHost)
	builder.WriteString("/api-token-auth/")
	urls := builder.String()

	res, err := http.PostForm(urls, url.Values{"username": []string{global_type.Config.WebServerUserName}, "password": []string{global_type.Config.WebserverKey}})
	if err != nil {
		fmt.Printf("%s", ErrorMessage(errors.Wrapf(err, "Genarate Key Failure. Check the username&password or the status of the server.\n")))
		os.Exit(1)
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Println(errors.WithMessage(err, "Res Body Close Fails!"))
		}
	}()
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

// ConCatStrings
//	@Description: 字符串高效拼接
//	@param stringList
//	@return string
//	@author KevinMatt 2021-08-05 20:03:50
//	@function_mark PASS
func ConCatStrings(stringList ...string) string {
	var builder strings.Builder
	for index := range stringList {
		builder.WriteString(stringList[index])
	}
	return builder.String()
}

// ErrorMessage
//	@Description: 只打印错误信息，不打印堆栈
//	@param err
//	@return string
//	@author KevinMatt 2021-08-08 16:14:42
//	@function_mark PASS
func ErrorMessage(err error) string {
	return err.Error()
}

// ErrorStack
//	@Description: 打印含堆栈的错误信息
//	@param err 错误
//	@return string 字符串
//	@author KevinMatt 2021-08-08 16:13:58
//	@function_mark PASS
func ErrorStack(err error) string {
	errMsg := fmt.Sprintf("%+v", err)
	return CleanPath(errMsg)
}

// CleanPath
//	@Description: 信息脱敏
//	@param s 传入信息
//	@return string 返回脱敏字符串
//	@author KevinMatt 2021-08-08 16:03:40
//	@function_mark PASS
func CleanPath(s string) string {
	return strings.ReplaceAll(s, strings.ReplaceAll(global_type.WorkPath, "\\", "/")+"/", "")
}

// var Base64Encrypt = func
//	@Description: 为原始的加密内容添加Base64编码
//	@param text 要加密的文本
//	@return string 加密的文本
//	@author KevinMatt 2021-08-10 01:07:41
//	@function_mark PASS
var Base64Encrypt = func(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(Encrypt(global_type.Config.ProjectId, global_type.Config.CryptoKey, text)))
}

// ForDebug
//	@Description: 断点小帮手
//	@param any
//	@author KevinMatt 2021-08-10 01:32:22
//	@function_mark PASS
func ForDebug(any ...interface{}) {
	return
}

//	GetCommitInfo
//	@Description: 获取commit信息
//	@param line commitInfo行
//	@return global_type.CommitInfoType 返回结构体
//	@author KevinMatt 2021-08-10 01:04:21
//	@function_mark PASS
func GetCommitInfo(line string) global_type.CommitInfoType {
	infoList := strings.Split(line, ",")
	var tempCommitInfo global_type.CommitInfoType
	tempCommitInfo = global_type.CommitInfoType{
		CommitHash:     infoList[0],
		CommitterEmail: infoList[1],
		CommitTime:     ToIso8601(strings.Split(infoList[len(infoList)-1][4:], " ")),
	}
	// 赋值commitAuthor(考虑多个Author的可能)
	for index := 2; index < len(infoList)-1; index++ {
		tempCommitInfo.CommitAuthor += infoList[index]
		if index != len(infoList)-2 {
			tempCommitInfo.CommitAuthor = ConCatStrings(tempCommitInfo.CommitAuthor, ",")
		}
	}
	return tempCommitInfo
}
