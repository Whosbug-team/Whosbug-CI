package whosbugPack

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type postData struct {
	Objects []struct {
		Owner      string `json:"owner"`
		FilePath   string `json:"file_path"`
		ParentName string `json:"parent_name"`
		ParentHash string `json:"parent_hash"`
		Name       string `json:"name"`
		Hash       string `json:"hash"`
		OldName    string `json:"old_name"`
		CommitTime string `json:"commit_time"`
	} `json:"objects"`
	Project struct {
		Pid string `json:"pid"`
	} `json:"project"`
	Release struct {
		Release    string `json:"release"`
		CommitHash string `json:"commit_hash"`
	} `json:"release"`
}

type objectForPost struct {
	Owner      string `json:"owner"`
	FilePath   string `json:"file_path"`
	ParentName string `json:"parent_name"`
	ParentHash string `json:"parent_hash"`
	Name       string `json:"name"`
	Hash       string `json:"hash"`
	OldName    string `json:"old_name"`
	CommitTime string `json:"commit_time"`
}

func postObjects(projectId string, releaseVersion string, commitHash string, objects []ObjectInfoType) error {
	tempEncrypt := func(text string) string {

		return base64.StdEncoding.EncodeToString([]byte(_encrypt(projectId, _SECRET, text)))
	}
	if tempEncrypt(projectId) == "" {
		return nil
	}
	var dataForPost postData
	dataForPost.Project.Pid = tempEncrypt(projectId)
	dataForPost.Release.Release = tempEncrypt(releaseVersion)
	dataForPost.Release.CommitHash = tempEncrypt(commitHash)
	for _, object := range objects {
		var objectForAppend objectForPost
		objectForAppend.Owner = tempEncrypt(object["owner"])
		objectForAppend.FilePath = tempEncrypt(object["file_path"])
		objectForAppend.ParentName = tempEncrypt(object["parent_name"])
		objectForAppend.ParentHash = tempEncrypt(object["parent_hash"])
		objectForAppend.Name = tempEncrypt(object["name"])
		objectForAppend.Hash = tempEncrypt(object["hash"])
		objectForAppend.OldName = tempEncrypt(object["old_name"])
		objectForAppend.CommitTime = object["commit_time"]
		dataForPost.Objects = append(dataForPost.Objects, objectForAppend)
	}
	data, err := json.MarshalToString(&dataForPost)
	if err != nil {
		log.Println(err)
	}
	//准备发送
	urlReq := _HOST + "/whosbug/commits/diffs/"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, urlReq, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Println(err)
		return err
	}

	token, err := _genToken()
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == 201 {
		return nil
	} else {
		fmt.Println(res.StatusCode)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return err
		}
		//fmt.Println(string(body))
		return errors.New(string(body))
	}
}
