package whosbugAssigns

import "crypto/sha256"

// hashCode64
/* @Description: 返回sha256编码的拼接字符串
 * @param projectId 项目ID
 * @param objectName
 * @param filePath 文件目录
 * @return string 返回编码字符串
 * @author KevinMatt 2021-07-25 14:20:09
 * @function_mark
 */
func hashCode64(projectId string, objectName string, filePath string) string {
	text := projectId + objectName + filePath
	return string(sha256.New().Sum([]byte(text)))
}
