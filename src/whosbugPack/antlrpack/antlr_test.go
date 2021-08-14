package antlrpack

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"whosbugPack/utility"
)

func TestExecuteJava(t *testing.T) {
	input, _ := os.Open("C:\\Users\\KevinMatt\\Desktop\\whosbug-Golang\\AllInOne7.java")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteJava(string(text))
	for _, item := range rest.AstInfoList.Classes {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ClassName, "\tMasterObject: ", item.MasterObject)
	}
	for _, item := range rest.AstInfoList.Methods {
		if item.CallMethods != nil {
			fmt.Println("Methods: ", item.CallMethods)
		}
	}
}

func TestExecuteKotlin(t *testing.T) {
	input, _ := os.Open("C:\\Users\\KevinMatt\\Desktop\\whosbug-Golang\\test.kt")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteKotlin(string(text))
	utility.ForDebug(rest)
	for _, item := range rest.AstInfoList.Classes {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ClassName, "\tMasterObject: ", item.MasterObject)
	}
	for _, item := range rest.AstInfoList.Methods {
		if item.CallMethods != nil {
			fmt.Println("Methods: ", item.CallMethods)
		}
	}
}

func TestRemoveRep(t *testing.T) {
	input := []string{"a", "b", "a", "b"}
	res := RemoveRep(input)
	for _, item := range res {
		fmt.Println(item)
	}
}
