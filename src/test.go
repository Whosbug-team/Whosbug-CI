package main

import (
	"fmt"
	"io/ioutil"
	"whosbugAssigns"
)

func main() {
	data, err := ioutil.ReadFile("C:\\Users\\KevinMatt\\Desktop\\whosbug-Golang\\logs.diff")
	//data := "diff --git a/src/main/java/com/info_interface/demo/DemoApplication.java b/src/main/java/com/info_interface/demo/DemoApplication.java\nindex 0b29ecf..1cb1764 100644\n--- a/src/main/java/com/info_interface/demo/DemoApplication.java\n+++ b/src/main/java/com/info_interface/demo/DemoApplication.java\n@@ -1, 16 + 1, 16 @@"
	//patDiff, err := regexp.Compile(`(diff\ \-\-git\ a/(.+)\ b/.+\n)`)
	dataString := string(data)
	//info := patDiff.FindStringSubmatch(dataString)
	//fmt.Print(dataString)
	if err != nil {
		fmt.Println("dadad")
	}
	//fmt.Println(info)
	//fmt.Println(info[2])
	whosbugAssigns.ParseDiff(dataString)
}
