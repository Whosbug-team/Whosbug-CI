package whosbugAssigns

import (
	parser "anrlr4_ast/java"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/** antlrAnalysis
 * @Description:
 * @param targetFilePath
 * @param langMode
 * @return string
 * @author KevinMatt 2021-07-22 21:45:01
 * @function_mark
 */
func antlrAnalysis(targetFilePath string, langMode string) string {
	//TODO antlr4 GOLANG解析的使用方法
	switch langMode {
	case "java":
		result, err := executeJava(targetFilePath)
		errorHandler(err)
		fmt.Println(result)

	}
	return "true"
}

type TreeShapeListener struct {
	*parser.BaseJavaParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
func executeJava(targetFilePath string) (string, error) {
	input, err := antlr.NewFileStream(targetFilePath)
	if err != nil {
		return "", err
	}
	lexer := parser.NewJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJavaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return "?", nil
}
