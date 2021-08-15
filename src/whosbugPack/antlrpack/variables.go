package antlrpack

import (
	"sync"
	cpp "whosbugPack/antlrpack/cpp_lib"
	golang "whosbugPack/antlrpack/go_lib"
	javaparser "whosbugPack/antlrpack/java_lib"
	javascript "whosbugPack/antlrpack/js_lib"
	kotlin "whosbugPack/antlrpack/kotlin_lib"
)

var (
	javaLexerPool = &sync.Pool{New: func() interface{} {
		return javaparser.NewJavaLexer(nil)
	}}
	javaParserPool = &sync.Pool{New: func() interface{} {
		return javaparser.NewJavaParser(nil)
	}}

	goLexerPool = &sync.Pool{New: func() interface{} {
		return golang.NewGoLexer(nil)
	}}
	goParserPool = &sync.Pool{New: func() interface{} {
		return golang.NewGoParser(nil)
	}}

	kotlinLexerPool = &sync.Pool{New: func() interface{} {
		return kotlin.NewKotlinLexer(nil)
	}}
	kotlinParserPool = &sync.Pool{New: func() interface{} {
		return kotlin.NewKotlinParser(nil)
	}}

	cppLexerPool = &sync.Pool{New: func() interface{} {
		return cpp.NewCPP14Lexer(nil)
	}}
	cppParserPool = &sync.Pool{New: func() interface{} {
		return cpp.NewCPP14Parser(nil)
	}}

	javascriptLexerPool = &sync.Pool{New: func() interface{} {
		return javascript.NewJavaScriptLexer(nil)
	}}
	javascriptParserPool = &sync.Pool{New: func() interface{} {
		return javascript.NewJavaScriptParser(nil)
	}}

	newJavaTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(JavaTreeShapeListener)
	}}
	newKotlinTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(KotlinTreeShapeListener)
	}}
	newGoTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(GoTreeShapeListener)
	}}
	newCppTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(CppTreeShapeListener)
	}}
	newJavaScriptTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(JSTreeShapeListener)
	}}
)
