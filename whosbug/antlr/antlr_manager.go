package antlr

import (
	"sync"

	c "git.woa.com/bkdevops/whosbug/antlr/cLib"
	cpp "git.woa.com/bkdevops/whosbug/antlr/cppLib"
	golang "git.woa.com/bkdevops/whosbug/antlr/goLib"
	java "git.woa.com/bkdevops/whosbug/antlr/javaLib"
	js "git.woa.com/bkdevops/whosbug/antlr/jsLib"
	kt "git.woa.com/bkdevops/whosbug/antlr/kotlinLib"
)

var (
	cLexerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return c.NewCLexer(nil)
	}}
	cParserPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return c.NewCParser(nil)
	}}

	javaLexerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return java.NewJavaLexer(nil)
	}}
	javaParserPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return java.NewJavaParser(nil)
	}}

	goLexerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return golang.NewGoLexer(nil)
	}}
	goParserPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return golang.NewGoParser(nil)
	}}

	kotlinLexerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return kt.NewKotlinLexer(nil)
	}}
	kotlinParserPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return kt.NewKotlinParser(nil)
	}}

	cppLexerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return cpp.NewCPP14Lexer(nil)
	}}
	cppParserPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return cpp.NewCPP14Parser(nil)
	}}

	javascriptLexerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return js.NewJavaScriptLexer(nil)
	}}
	javascriptParserPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return js.NewJavaScriptParser(nil)
	}}

	newCTreeShapeListenerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return new(CTreeShapeListener)
	}}
	newJavaTreeShapeListenerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return new(JavaTreeShapeListener)
	}}
	newKotlinTreeShapeListenerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return new(KotlinTreeShapeListener)
	}}
	newGoTreeShapeListenerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return new(GoTreeShapeListener)
	}}
	newCppTreeShapeListenerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return new(CppTreeShapeListener)
	}}
	newJavaScriptTreeShapeListenerPool *sync.Pool = &sync.Pool{New: func() interface{} {
		return new(JSTreeShapeListener)
	}}
)
