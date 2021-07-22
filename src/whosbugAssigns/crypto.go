package whosbugAssigns

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
)

/**
 * @Description: 生成AES-CFB需要的Key和IV
 * @param projectId 项目ID
 * @param key 加密密钥
 * @return []byte K密钥
 * @return []byte IV偏移密钥
 * @author KevinMatt (he_yuheng@163.com)
 */
func generateKIV(projectId, key []byte) ([]byte, []byte) {
	hK := hmac.New(sha256.New, key)
	hIV := hmac.New(md5.New, key)
	hK.Write(projectId)
	hIV.Write(projectId)
	return hK.Sum(nil), hIV.Sum(nil)
}

/**
 * @Description: AES-CFB加密
 * @param projectId 项目ID
 * @param Dest 输出的加密后字符串
 * @param key 加密密钥
 * @param plainText 需要加密的文本
 * @return error 错误抛出
 * @author KevinMatt (he_yuheng@163.com)
 */
func encrypt(projectId, Dest, key, plainText []byte) error {
	K, IV := generateKIV(projectId, key)
	aesBlockEncryptor, err := aes.NewCipher(K)
	if err != nil {
		return err
	}
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream(plainText, Dest)
	return nil
}

/**
 * @Description: AES-CFB解密
 * @param projectId 项目ID
 * @param Dest 解密完成的字符串
 * @param key 解密密钥
 * @param plainText 需要解密的文本
 * @return error 错误抛出
 * @author KevinMatt (he_yuheng@163.com)
 */
func decrypt(projectId, Dest, key, plainText []byte) error {
	K, IV := generateKIV(projectId, key)
	aesBlockDescriptor, err := aes.NewCipher(K)
	if err != nil {
		return err
	}
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream(plainText, Dest)
	return nil
}
