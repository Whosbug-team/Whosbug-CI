package whosbugAssigns

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
)

/** generateKIV
 * @Description:
 * @param projectId
 * @param secret
 * @author KevinMatt
 */
func generateKIV(projectId, key string) ([]byte, []byte) {
	hK := hmac.New(sha256.New, []byte(key))
	hIV := hmac.New(md5.New, []byte(key))
	hK.Write([]byte(projectId))
	hIV.Write([]byte(projectId))
	return hK.Sum(nil), hIV.Sum(nil)
}

/**
 * @Description: AES-CFB加密
 * @param src 传入的待加密字符串
 * @param Dest 输出的加密后字符串
 * @param key 加密密钥
 * @param plainText
 * @return error
 * @author KevinMatt
 */
func encrypt(projectId string, key string, plainText string) (string, error) {
	K, IV := generateKIV(projectId, key)
	aesBlockEncryptor, err := aes.NewCipher(K)
	if err != nil {
		return "", err
	}
	var Dest string
	aesEncryptor := cipher.NewCFBEncrypter(aesBlockEncryptor, IV)
	aesEncryptor.XORKeyStream([]byte(plainText), []byte(Dest))
	return Dest, nil
}

/**
 * @Description: AES-CFB解密
 * @param src 需要解密的字符串
 * @param Dest 解密完成的字符串
 * @param key 解密密钥
 * @return error
 * @author KevinMatt
 */
func decrypt(src string, key string, plainText string) (string, error) {
	K, IV := generateKIV(src, key)
	aesBlockDescriptor, err := aes.NewCipher(K)
	if err != nil {
		return "", err
	}
	var Dest string
	aesDescriptor := cipher.NewCFBDecrypter(aesBlockDescriptor, IV)
	aesDescriptor.XORKeyStream([]byte(plainText), []byte(Dest))
	return Dest, nil
}
