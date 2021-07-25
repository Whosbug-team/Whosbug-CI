package whosbugAssigns

import (
	javaparser "anrlr4_ast/java"
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
func antlrAnalysis(targetFilePath string, langMode string) javaparser.AnalysisInfoType {
	var result javaparser.AnalysisInfoType
	//TODO antlr4 GOLANG解析的使用方法
	switch langMode {
	case "java":
		result = executeJava(targetFilePath)
		javaparser.Infos.SetEmpty()
	default:
		break
	}
	return result
}

type TreeShapeListener struct {
	*javaparser.BaseJavaParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

//func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText())
//}

/** executeJava
 * @Description: 执行Java Antlr语法解析
 * @param targetFilePath
 * @return string
 * @return error
 * @author KevinMatt 2021-07-24 17:51:25
 * @function_mark
 */
func executeJava(targetFilePath string) javaparser.AnalysisInfoType {
	input, err := antlr.NewFileStream(targetFilePath)
	if err != nil {
		errorHandler(err)
	}
	lexer := javaparser.NewJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := javaparser.NewJavaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return javaparser.Infos
}
