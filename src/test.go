package main

import (
	"encoding/json"
	"fmt"
	"os"
	"whosbugAssigns"
)

type input_json struct {
	ProjectId       string   `json:"__PROJRCT_ID"`
	ReleaseVersion  string   `json:"__RELEASE_VERSION"`
	ProjectUrl      string   `json:"__PROJECT_URL"`
	BranchName      string   `json:"__BRANCH_NAME"`
	LanguageSupport []string `json:"__LAN_SUPPORT"`
	WebServerHost   string   `json:"__WEB_SRV_HOST""`
}
type innerConfig struct {
	userId string
	secret string
}

var innerConf innerConfig
var config input_json

func main() {
	fmt.Println("Start!")
	GetInputConfig()
	projectId := config.ProjectId
	branchName := config.BranchName
	repoPath := config.ProjectUrl
	resCommits := whosbugAssigns.AnalysisTest(repoPath, branchName, projectId)
	result(resCommits, projectId, "1.0.0")
	fmt.Println("Whosbug analysis done")
}
func GetInputConfig() {
	file, err := os.Open("src/input.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Get input.json succeed!")
	}
	fmt.Println("Version:\t", config.ReleaseVersion, "\nProjectId:\t", config.ProjectId, "\nBranchName:\t", config.BranchName)
}

//func TestParseCommit() []whosbugAssigns.CommitParsedType {
//for _, temp := range TestParseCommit() {
////	fmt.Println(temp.CommitTime)
////	fmt.Println(temp.Commit)
////}
//	return whosbugAssigns.TestParseCommit()
//}

//func TestGetDiff() map[string]string {
//	return whosbugAssigns.GetDiffTest("C:\\Users\\KevinMatt\\Desktop\\java-test\\", "master", "whosbug_test_1")
//}

//func TestParseDiff() {
//	data, err := ioutil.ReadFile("C:\\Users\\KevinMatt\\Desktop\\whosbug-Golang\\logs.diff")
//	//data := "diff --git a/src/main/java/com/info_interface/demo/DemoApplication.java b/src/main/java/com/info_interface/demo/DemoApplication.java\nindex 0b29ecf..1cb1764 100644\n--- a/src/main/java/com/info_interface/demo/DemoApplication.java\n+++ b/src/main/java/com/info_interface/demo/DemoApplication.java\n@@ -1, 16 + 1, 16 @@"
//	//patDiff, err := regexp.Compile(`(diff\ \-\-git\ a/(.+)\ b/.+\n)`)
//	dataString := string(data)
//	//info := patDiff.FindStringSubmatch(dataString)
//	//fmt.Print(dataString)
//	if err != nil {
//		fmt.Println("dadad")
//	}
//	//fmt.Println(info)
//	//fmt.Println(info[2])
//	whosbugAssigns.ParseDiff(dataString)
//}
//func TestCrypto() {
//	var test string = "abcdefg"
//	var key string = "1234567890123456"
//	var projectId string = "sadasd"
//	recv := make([]byte, len(test))
//	fmt.Println(test)
//	whosbugAssigns.En([]byte(projectId), recv, []byte(key), []byte(test))
//	fmt.Printf("%s\n", recv)
//	whosbugAssigns.De([]byte(projectId), recv, []byte(key), recv)
//	fmt.Printf("%s\n", recv)
//}
