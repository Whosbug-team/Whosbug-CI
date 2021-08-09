package utility

import (
	"bytes"
	"encoding/base64"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"whosbugPack/global_type"
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
	urlReq := ConCatStrings(global_type.Config.WebServerHost, "/whosbug/releases/last/")
	method := "POST"

	pid := base64.StdEncoding.EncodeToString([]byte(Encrypt(projectId, global_type.Config.CryptoKey, projectId)))
	data := []byte("{\"pid\":\"" + pid + "\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer(data))

	if err != nil {
		return "", errors.Wrapf(err, "GetLatestRelease->Sending NewRequest")
	}

	token, err := GenToken()
	if err != nil {
		return "", errors.WithStack(err)
	}
	req.Header.Add("Authorization", ConCatStrings("Token ", token))
	req.Header.Set("Content-Type", "application/Json")

	res, err := client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Println(ErrorStack(errors.WithStack(err)))
			os.Exit(1)
		}
	}()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		commitHash := Json.Get(body, "commit_hash").ToString()
		commitHashByte, err := base64.StdEncoding.DecodeString(commitHash)
		return Decrypt(projectId, global_type.Config.CryptoKey, string(commitHashByte)), nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		//fmt.Println(string(body))
		if res.StatusCode == 404 {
			return "", errors.New("The Project Not Found. Get all commit to Initialize.")
		}
		return "", errors.New(string(body))
	}
}
