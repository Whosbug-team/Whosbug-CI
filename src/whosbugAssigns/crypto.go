package whosbugAssigns

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
)

/** generateKIV
 * @Description: 生成AES-CFB需要的Key和IV
 * @param projectId 项目ID
 * @param key 加密密钥
 * @return []byte K密钥
 * @return []byte IV偏移密钥
 * @author KevinMatt 2021-07-22 13:23:16
 * @function_mark PASS
 */
func generateKIV(projectId, key []byte) ([]byte, []byte) {
	hK := hmac.New(sha256.New, key)
	hIV := hmac.New(md5.New, key)
	hK.Write(projectId)
	hIV.Write(projectId)
	return hK.Sum(nil), hIV.Sum(nil)
}

/** encrypt
 * @Description: AES-CFB加密
 * @param projectId 项目ID
 * @param Dest 输出的加密后字符串
 * @param key 加密密钥
 * @param plainText 需要加密的文本
 * @return error 错误抛出
 * @author KevinMatt 2021-07-22 13:23:24
 * @function_mark PASS
 */
func encrypt(projectId, Dest, key, plainText []byte) error {
	K, IV := generateKIV(projectId, key)
	aesBlockEncryptor, err := aes.NewCipher(K)
	if err != nil {
		return err
	}
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream(Dest, plainText)
	return nil
}

/** decrypt
 * @Description: AES-CFB解密
 * @param projectId 项目ID
 * @param Dest 解密完成的字符串
 * @param key 解密密钥
 * @param plainText 需要解密的文本
 * @return error 错误抛出
 * @author KevinMatt 2021-07-22 13:23:30
 * @function_mark PASS
 */
func decrypt(projectId, Dest, key, plainText []byte) error {
	K, IV := generateKIV(projectId, key)
	aesBlockDescriptor, err := aes.NewCipher(K)
	if err != nil {
		return err
	}
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream(Dest, plainText)
	return nil
}
