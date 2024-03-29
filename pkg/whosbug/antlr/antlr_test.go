package antlr

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"git.woa.com/bkdevops/whosbug-ci/internal/util"
	"git.woa.com/bkdevops/whosbug-ci/internal/zaplog"
)

func TestExecuteJava(t *testing.T) {
	input, _ := os.Open("/root/whosbugGolang/AllInOne7.java")
	text, _ := ioutil.ReadAll(input)
	parser := javaParserPool.Get().(*JavaAstParser)
	rest := parser.AstParse(string(text))
	util.ForDebug(rest)
	// for _, item := range rest.AstInfoList.Classes {
	// 	zaplog.Logger.Info("StartLine: %v\tEndLine: %v\tClassName: %v\tMasterObject: %v", item.StartLine, item.EndLine, item.ClassName, item.MasterObject)
	// }
	// for _, item := range rest.AstInfoList.Methods {
	// 	if item.CallMethods != nil {
	// 		zaplog.Logger.Info("Methods: %v", item.CallMethods)
	// 	}
	// }
}
func TestExecuteCpp(t *testing.T) {
	files, _ := filepath.Glob("/root/example/CppPrimer/*/*.h")
	file2, _ := filepath.Glob("/root/example/CppPrimer/*/*.cpp")
	files = append(files, file2...)
	parser := cppParserPool.Get().(*CppAstParser)
	for _, f := range files {
		input, _ := os.Open(f)
		text, _ := ioutil.ReadAll(input)
		input.Close()
		rest := parser.AstParse(string(text))
		util.ForDebug(rest)
	}
	util.ForDebug()
	// for _, item := range rest.AstInfoList.Classes {
	// 	zaplog.Logger.Info("StartLine: %v\tEndLine: %v\tClassName: %v\tMasterObject: %v", item.StartLine, item.EndLine, item.ClassName, item.MasterObject)
	// }
	// for _, item := range rest.AstInfoList.Methods {
	// 	if item.CallMethods != nil {
	// 		zaplog.Logger.Info("Methods:%s,Calling:%s\n ", item.MethodName, item.CallMethods)
	// 	}
	// }
}

func TestExecuteGolang(t *testing.T) {
	input, _ := os.Open("/root/example/struct_promotion.go")
	text, _ := ioutil.ReadAll(input)
	parser := goParserPool.Get().(*GoAstParser)
	rest := parser.AstParse(string(text))
	util.ForDebug(rest)
	// for _, item := range rest.AstInfoList.Classes {
	// 	zaplog.Logger.Info("StartLine: %v\tEndLine: %v\tClassName: %v\tMasterObject: %v", item.StartLine, item.EndLine, item.ClassName, item.MasterObject)
	// }
	// for _, item := range rest.AstInfoList.Methods {
	// 	if item.CallMethods != nil {
	// 		zaplog.Logger.Info("Methods:%s,Calling:%s\n ", item.MethodName, item.CallMethods)
	// 	}
	// }
}

func TestExecuteKotlin(t *testing.T) {
	input, _ := os.Open("/root/example/Test.kt")
	text, _ := ioutil.ReadAll(input)
	parser := kotlinParserPool.Get().(*KotlinAstParser)
	rest := parser.AstParse(string(text))
	util.ForDebug(rest)
	// for _, item := range rest.AstInfoList.Classes {
	// 	zaplog.Logger.Info("StartLine: %v\tEndLine: %v\tClassName: %v\tMasterObject: %v", item.StartLine, item.EndLine, item.ClassName, item.MasterObject)
	// }
	// for _, item := range rest.AstInfoList.Methods {
	// 	if item.CallMethods != nil {
	// 		zaplog.Logger.Info("Methods: %v, Calling: %v", item.MethodName, item.CallMethods)
	// 	}
	// }
}

func TestExecuteJavaScript(t *testing.T) {
	files, _ := filepath.Glob("/root/example/js/*.js")
	for _, f := range files {
		input, _ := os.Open(f)
		text, _ := ioutil.ReadAll(input)
		input.Close()
		parser := javascriptParserPool.Get().(*JavaAstParser)
		rest := parser.AstParse(string(text))
		if rest.Classes == nil && rest.Methods == nil {
			continue
		}
		util.ForDebug(rest)
	}
	util.ForDebug()
}

func TestRemoveRep(t *testing.T) {
	input := []string{"a", "b", "a", "b"}
	res := RemoveRep(input)
	for _, item := range res {
		zaplog.Logger.Info(item)
	}
}
