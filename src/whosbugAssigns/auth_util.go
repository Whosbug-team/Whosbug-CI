package whosbugAssigns

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const HOST = "http://127.0.0.1:8082"
const secret = "3E5D4C94-A9FE-4690-BEF4-76C40EAE44AB"
const userId = "qapm"

var tokenGot string

func genToken() {
	urls := HOST + "/api-token-auth/"
	res, _ := http.PostForm(urls, url.Values{"username": {innerConf.username}, "password": {innerConf.password}})
	if res.StatusCode == 200 {
		resBody, _ := ioutil.ReadAll(res.Body)
		tokenGot = strings.Split(string(resBody), "\"")[3]
		fmt.Println(tokenGot)
	}
}
func GenTokenTest() {
	getLatestRelease("whosbug")
}

/** getLatestRelease
 * @Description: 获得最新的Release信息
 * @param projectId 项目ID
 * @return string Release信息
 * @author KevinMatt 2021-07-22 16:50:26
 * @function_mark
 */
func getLatestRelease(projectId string) string {
	// TODO Not Functioning
	urls := HOST + "/release/last/"
	headers := make(map[string]string)
	headers["Authorization"] = "Token " + tokenGot
	data := make(map[string]string)
	var dest []byte
	err := encrypt([]byte(projectId), dest, []byte(secret), []byte(projectId))
	data["pid"] = string(dest)

	res, err := http.PostForm(urls, url.Values{"pid": {string(dest)}, "header": {"Token " + tokenGot}})
	errorHandler(err)
	if res.StatusCode == 200 {
		res, err := ioutil.ReadAll(res.Body)
		errorHandler(err)
		return string(res)
	} else {
		fmt.Println(res.Body)
		return ""
	}
}
