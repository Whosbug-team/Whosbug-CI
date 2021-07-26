package whosbugAssigns

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

const HOST = "http://127.0.0.1:8082"

var secret string

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
	genToken()
	getLatestRelease("whosbug_test_1")
}

/** getLatestRelease
 * @Description: 获得最新的Release信息
 * @param projectId 项目ID
 * @return string Release信息
 * @author KevinMatt 2021-07-22 16:50:26
 * @function_mark PASS
 */
func getLatestRelease(projectId string) string {
	// TODO Not Functioning
	urlReq := HOST + "/whosbug/releases/last/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	pid := base64.StdEncoding.EncodeToString([]byte(encrypt(projectId, secret, projectId)))
	_ = writer.WriteField("pid", pid)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Authorization", "Token 69b6ea89661c66b03d2b1e7181f2f74134e4af4e")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(body)
	} else {
		fmt.Println(res.StatusCode)
		return ""
	}
}

//func postDiffResult(diffRes) {
//	urlPost := HOST + "/whosbug/commits/diffs/"
//	method := "POST"
//
//	payload := &bytes.Buffer{}
//	writer := multipart.NewWriter(payload)
//
//
//}
