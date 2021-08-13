package antlrpack

import (
	"sync"
	golang "whosbugPack/antlrpack/go_lib"
	javaparser "whosbugPack/antlrpack/java_lib"
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

	newJavaTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(JavaTreeShapeListener)
	}}
	newKotlinTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(KotlinTreeShapeListener)
	}}
	newGoTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(GoTreeShapeListener)
	}}
)
