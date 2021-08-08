package antlrpack

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestExecuteJava(t *testing.T) {
	input, _ := os.Open("C:\\Users\\KevinMatt\\Desktop\\whosbug-Golang\\AllInOne7.java")
	text, _ := ioutil.ReadAll(input)
	rest := ExecuteJava(string(text))
	fmt.Println(rest.AstInfoList.PackageName)
	for _, item := range rest.AstInfoList.Imports {
		fmt.Println("Imports: ", item)
	}
	for _, item := range rest.AstInfoList.Classes {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ClassName, "\tExtends: ", item.Extends, "\tImplements: ", item.Implements, "\tMasterObject: ", item.MasterObject, "\tDepth: ", item.Depth)
	}
	for _, item := range rest.AstInfoList.Methods {
		fmt.Println("Methods: ", item)
	}
	for _, item := range rest.AstInfoList.Fields {
		fmt.Println("Fields: ", item)
	}

	//fmt.Println(rest.AstInfoList.Imports)
	//fmt.Println(rest.AstInfoList.Classes)
	//fmt.Println(rest.AstInfoList.Methods)
	//fmt.Println(rest.AstInfoList.Fields)
}
