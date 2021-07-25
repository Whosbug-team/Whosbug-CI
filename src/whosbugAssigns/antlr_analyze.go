package whosbugAssigns

import (
	javaparser "anrlr4_ast/java"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// antlrAnalysis
/* @Description: 执行antlr分析入口函数
 * @param targetFilePath 目标代码目录
 * @param langMode 解析语言模式
 * @return javaparser.AnalysisInfoType
 * @author KevinMatt 2021-07-25 13:56:08
 * @function_mark PASS
 */
func antlrAnalysis(targetFilePath string, langMode string) javaparser.AnalysisInfoType {
	var result javaparser.AnalysisInfoType
	// TODO 目前只有Java的支持
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

// executeJava
/* @Description: 执行Java Antlr语法解析
 * @param targetFilePath 解析目标目录
 * @return javaparser.AnalysisInfoType
 * @author KevinMatt 2021-07-25 14:00:10
 * @function_mark PASS
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
