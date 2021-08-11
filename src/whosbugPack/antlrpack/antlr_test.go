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
	for _, item := range rest.AstInfoList.Classes {
		fmt.Println("StartLine: ", item.StartLine, "\tEndLine: ", item.EndLine, "\tClassName: ", item.ClassName, "\tImplements: ", item.MasterObject)
	}
	for _, item := range rest.AstInfoList.Methods {
		fmt.Println("Methods: ", item.CallMethods)
	}
	for _, item := range rest.AstInfoList.Fields {
		fmt.Println("Fields: ", item)
	}
}
