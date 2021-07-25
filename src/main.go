package main

import (
	"fmt"
	"whosbugAssigns"
)

func main() {
	fmt.Println("Start!")
	whosbugAssigns.GetInputConfig()
	projectId := whosbugAssigns.Config.ProjectId
	branchName := whosbugAssigns.Config.BranchName
	repoPath := whosbugAssigns.Config.ProjectUrl
	resCommits := whosbugAssigns.Analysis(repoPath, branchName, projectId)
	whosbugAssigns.Result(resCommits, projectId, "1.0.0")
	for _, resCommit := range resCommits {
		fmt.Println(resCommit.Commit, " ", resCommit.CommitDiffs[0].DiffContent[0]["Name"])
	}
	fmt.Println("Whosbug analysis done")
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
