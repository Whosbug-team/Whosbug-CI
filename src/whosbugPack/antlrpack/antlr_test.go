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

func TestExecuteGolang(t *testing.T) {
	input, _ := os.Open("C:\\Users\\Sirius\\Desktop\\test1.go")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteGolang(string(text))
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

func TestExecuteJavaScript(t *testing.T) {
	input, _ := os.Open("D:\\Desktop\\Whosbug_antlr_go\\antlr4_js\\testfiles\\test.js")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteJavaScript(string(text))
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
