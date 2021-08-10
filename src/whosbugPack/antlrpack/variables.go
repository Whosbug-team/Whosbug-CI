package antlrpack

import (
	javaparser "antlr4_ast/java"
	"sync"
)

var (
	lexerPool = &sync.Pool{New: func() interface{} {
		return javaparser.NewJavaLexer(nil)
	}}
	parserPool = &sync.Pool{New: func() interface{} {
		return javaparser.NewJavaParser(nil)
	}}
	newTreeShapeListenerPool = &sync.Pool{New: func() interface{} {
		return new(TreeShapeListener)
	}}
)
