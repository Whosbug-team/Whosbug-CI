package antlrpack

///* hashCode64
///* @Description: 返回sha256编码的拼接32位byte编码值
// * @param projectId 项目ID
// * @param objectName
// * @param filePath 文件目录
// * @return string 返回32位byte编码字符串
// * @author KevinMatt 2021-07-26 20:49:17
// * @function_mark PASS
//*/
//func hashCode64(projectId, objectName, filePath []byte) (text [32]byte) {
//	text = sha256.Sum256(append(append(projectId, objectName...), filePath...))
//	return
//}

// RemoveRep
// 	@Description: 切片去重
// 	@param s
// 	@return []string
// 	@author KevinMatt 2021-08-14 15:14:28
// 	@function_mark
func RemoveRep(s []string) []string {
	var result []string
	m := make(map[string]bool)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}
