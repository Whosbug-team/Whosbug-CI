package antlrpack

import (
	javaparser "anrlr4_ast/java"
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

// processDiffs 已处理的commit数
var processDiffs = 0

// LargeObjects 大型objects值
const LargeObjects = 100000
