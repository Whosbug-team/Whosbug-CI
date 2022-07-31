package util

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"git.woa.com/bkdevops/whosbug/config"

	"github.com/pkg/errors"
)

// GetLatestRelease 获取最新的release
//  @param projectID string
//  @return string Release信息
//  @return error
//  @author: Kevineluo 2022-07-31 01:03:27
func GetLatestRelease(projectID string) (string, error) {
	urlReq := ConCatStrings(config.WhosbugConfig.WebServerHost, "/whosbug/releases/last/")
	method := "POST"

	pid := base64.StdEncoding.EncodeToString([]byte(Encrypt(projectID, config.WhosbugConfig.CryptoKey, projectID)))
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
	req.Header.Add("Token", token)
	req.Header.Set("Content-Type", "application/json")

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
		commitHash := json.Get(body, "last_commit_hash").ToString()
		commitHashByte, err := base64.StdEncoding.DecodeString(commitHash)
		return Decrypt(projectID, config.WhosbugConfig.CryptoKey, string(commitHashByte)), nil
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.WithStack(err)
		}
		if res.StatusCode == 404 {
			return "", errors.New("The Project Not Found. Get all commit to Initialize")
		}
		return "", errors.New(string(body))
	}
}
