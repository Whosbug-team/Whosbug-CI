package utility

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* GetLatestRelease
/* @Description:
 * @param projectId 项目ID
 * @return string Release信息
 * @return error
 * @author KevinMatt 2021-08-03 18:12:18
 * @function_mark PASS
*/
func GetLatestRelease(projectId string) (string, error) {
	urlReq := ConCatStrings(_HOST, "/whosbug/releases/last/")
	method := "POST"

	pid := base64.StdEncoding.EncodeToString([]byte(Encrypt(projectId, _SECRET, projectId)))
	data := []byte("{\"pid\":\"" + pid + "\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer(data))
	if err != nil {
		err = fmt.Errorf("GetLatestRelease->Sending NewRequest: %w", err)
		return "", err
	}

	token, err := GenToken()
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", ConCatStrings("Token ", token))
	req.Header.Set("Content-Type", "application/Json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			//fmt.Println(string(body))
			return "", err
		}
		fmt.Println(string(body))
		commitHash := Json.Get(body, "commit_hash").ToString()
		commitHashByte, err := base64.StdEncoding.DecodeString(commitHash)
		return _decrypt(projectId, _SECRET, string(commitHashByte)), nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return "", err
		}
		fmt.Println(string(body))
		//TODO 改成相应的业务异常类型
		if res.StatusCode == 404 {
			return "", errors.New("The Project Not Found.")
		}
		return "", errors.New(string(body))
	}
}
