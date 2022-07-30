package util

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"git.woa.com/bkdevops/whosbug/config"

	"github.com/pkg/errors"
)

// GetLatestRelease
//	@Description:
//	@param projectId 项目ID
//	@return string Release信息
//	@return error
//	@author KevinMatt 2021-08-10 13:02:27
//	@function_mark PASS
func GetLatestRelease(projectId string) (string, error) {
	urlReq := ConCatStrings(config.WhosbugConfig.WebServerHost, "/whosbug/releases/last/")
	method := "POST"

	pid := base64.StdEncoding.EncodeToString([]byte(Encrypt(projectId, config.WhosbugConfig.CryptoKey, projectId)))
	data := []byte("{\"pid\":\"" + pid + "\"}")

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer(data))

	if err != nil {
		return "", errors.Wrapf(err, "GetLatestRelease->Sending NewRequest")
	}

	token, err := GenToken(config.DefaultExpire)
	if err != nil {
		return "", errors.WithStack(err)
	}
	req.Header.Add("TOKEN", token)
	req.Header.Set("Content-Type", "application/Json")

	res, err := client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		commitHash := Json.Get(body, "last_commit_hash").ToString()
		commitHashByte, err := base64.StdEncoding.DecodeString(commitHash)
		return Decrypt(projectId, config.WhosbugConfig.CryptoKey, string(commitHashByte)), nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		if res.StatusCode == 404 {
			return "", errors.New("The Project Not Found. Get all commit to Initialize.")
		}
		return "", errors.New(string(body))
	}
}
