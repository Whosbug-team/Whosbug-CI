package antlrpack

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"whosbugPack/utility"
)

func TestExecuteJava(t *testing.T) {
	input, _ := os.Open("D:\\Desktop\\Whosbug_antlr_go\\antlr4_java\\testfiles\\test.java")
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
func TestExecuteCpp(t *testing.T) {
	input, _ := os.Open("C:\\Users\\Sirius\\Desktop\\test.cpp")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteCpp(string(text))
	for _, item := range rest.AstInfoList.Classes {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ClassName, "\tMasterObject: ", item.MasterObject)
	}
	for _, item := range rest.AstInfoList.Methods {
		if item.CallMethods != nil {
			fmt.Printf("Methods:%s,Calling:%s\n ", item.MethodName, item.CallMethods)
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
			fmt.Printf("Methods:%s,Calling:%s\n ", item.MethodName, item.CallMethods)
		}
	}
}

func TestExecuteKotlin(t *testing.T) {
	input, _ := os.Open("D:\\Desktop\\Whosbug_antlr_go\\antlr4_kotlin\\testfiles\\test.kt")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteKotlin(string(text))
	utility.ForDebug(rest)
	for _, item := range rest.AstInfoList.Classes {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ClassName, "\tMasterObject: ", item.MasterObject)
	}
	for _, item := range rest.AstInfoList.Methods {
		if item.CallMethods != nil {
			fmt.Println("Methods: ", item.MethodName, "Calling: ", item.CallMethods)
		}
	}
}

func TestExecuteJavaScript(t *testing.T) {
	input, _ := os.Open("D:\\Desktop\\Whosbug_antlr_go\\antlr4_js\\testfiles\\test.js")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteJavaScript(string(text))
	utility.ForDebug(rest)
	for _, item := range rest.AstInfoList.Classes {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ClassName)
	}
	for _, item := range rest.AstInfoList.Objects {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ObjectName)
	}

	var temp CallMethodType
	for _, item := range rest.CallMethods {
		if item != temp {
			fmt.Println("Methods: ", item.Id)
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
