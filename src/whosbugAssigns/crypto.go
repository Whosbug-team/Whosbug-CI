package whosbugAssigns

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

// generateKIV
/* @Description: 		生成AES-CFB需要的Key和IV
 * @param projectId 	项目ID
 * @param key 			加密密钥
 * @return []byte 		K密钥
 * @return []byte 		IV偏移密钥
 * @author KevinMatt 2021-07-25 20:07:20
 * @function_mark PASS
 */
func generateKIV(projectId, key []byte) ([]byte, []byte) {
	hK := hmac.New(sha256.New, key)
	hIV := hmac.New(md5.New, key)
	hK.Write(projectId)
	hIV.Write(projectId)
	return hK.Sum(nil), hIV.Sum(nil)
}

// encrypt
/* @Description: 		AES-CFB加密
 * @param projectId 	项目ID
 * @param Dest 			输出的加密后字符串
 * @param key 			加密密钥
 * @param plainText 	需要加密的文本
 * @return error 		错误抛出
 * @author KevinMatt 2021-07-25 13:34:09
 * @function_mark PASS
 */
func encrypt(projectId, key, plainText string) string {
	K, IV := generateKIV([]byte(projectId), []byte(key))
	aesBlockEncryptor, err := aes.NewCipher(K)
	if err != nil {
		fmt.Println(err)
	}
	var dest = plainText
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream([]byte(dest), []byte(plainText))
	return dest
}

// decrypt
/* @Description: 		AES-CFB解密
 * @param projectId 	项目ID
 * @param Dest 			解密完成的字符串
 * @param key 			解密密钥
 * @param plainText 	需要解密的文本
 * @return error 		错误抛出
 * @author KevinMatt 2021-07-25 13:35:15
 * @function_mark PASS
 */
func decrypt(projectId, key, plainText string) string {
	K, IV := generateKIV([]byte(projectId), []byte(key))
	aesBlockDescriptor, err := aes.NewCipher(K)
	if err != nil {
		fmt.Println(err)
	}
	var dest = plainText
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream([]byte(dest), []byte(plainText))
	return dest
}
