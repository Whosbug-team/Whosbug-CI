package antlr

import (
	"sync"

	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr/js"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// JSAstParser implement AstParser for javascript language
//
//	@author kevineluo
//	@update 2023-02-28 01:15:09
type JSAstParser struct {
	astInfo AstInfo
}

var _ js.JavaScriptParserListener = &JSAstParser{}

var (
	javascriptLexerPool = &sync.Pool{New: func() any {
		return js.NewJavaScriptLexer(nil)
	}}
	javascriptParserPool = &sync.Pool{New: func() any {
		return js.NewJavaScriptParser(nil)
	}}
	newJavaScriptAstParserPool = &sync.Pool{New: func() any {
		return new(JSAstParser)
	}}
)

// AstParse main parse process for javascript language
//
//	@receiver s *JSAstParser
//	@param diffText string
//	@return AstInfo
//	@author kevineluo
//	@update 2023-02-28 01:17:45
func (s *JSAstParser) AstParse(input string) AstInfo {
	//	截取目标文本的输入流
	inputStream := antlr.NewInputStream(input)
	//	初始化lexer
	lexer := javascriptLexerPool.Get().(*js.JavaScriptLexer)
	defer javaLexerPool.Put(lexer)
	lexer.SetInputStream(inputStream)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := javascriptParserPool.Get().(*js.JavaScriptParser)
	p.RemoveErrorListeners()
	defer javascriptParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.Program()
	//	创建listener
	listener := newJavaScriptAstParserPool.Get().(*JSAstParser)
	defer newJavaScriptAstParserPool.Put(listener)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.astInfo
}

func getExtendIdentifier(ctx *js.ClassDeclarationContext) (extend string) {
	if ctx.ClassTail() != nil {
		tempCtx := ctx.ClassTail().(*js.ClassTailContext).SingleExpression()
		for {
			if tempCtx == nil {
				return
			}
			if _, ok := tempCtx.(*js.ArgumentsExpressionContext); ok {
				tempCtx = tempCtx.(*js.ArgumentsExpressionContext).SingleExpression()
				continue
			}
			if _, ok := tempCtx.(*js.IdentifierExpressionContext); ok {
				extend = tempCtx.(*js.IdentifierExpressionContext).GetText()
				return
			}
			return
		}
	}
	return
}

func findJsDeclChain(ctx antlr.ParseTree) (chain string) {
	tempCtx := ctx.GetParent()
	for {
		if _, ok := tempCtx.(*js.ClassDeclarationContext); ok {
			chain = tempCtx.(*js.ClassDeclarationContext).Identifier().GetText() + "." + chain
		}
		if _, ok := tempCtx.(*js.FunctionDeclarationContext); ok {
			chain = tempCtx.(*js.FunctionDeclarationContext).Identifier().GetText() + "." + chain
		}
		tempCtx = tempCtx.GetParent()
		if tempCtx == nil {
			return
		}
	}
}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *JSAstParser) EnterClassDeclaration(ctx *js.ClassDeclarationContext) {
	classInfo := Class{
		StartLine: ctx.GetStart().GetLine(),
		EndLine:   ctx.GetStop().GetLine(),
		Name:      findJsDeclChain(ctx) + ctx.Identifier().GetText(),
		Extends:   getExtendIdentifier(ctx),
	}
	s.astInfo.Classes = append(s.astInfo.Classes, classInfo)
}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *JSAstParser) ExitFunctionDeclaration(ctx *js.FunctionDeclarationContext) {
	methodInfo := Method{
		StartLine: ctx.GetStart().GetLine(),
		EndLine:   ctx.GetStop().GetLine(),
		Name:      findJsDeclChain(ctx) + ctx.Identifier().GetText(),
	}
	if ctx.FormalParameterList() != nil {
		methodInfo.Parameters = ctx.FormalParameterList().GetText()
	}
	s.astInfo.Methods = append(s.astInfo.Methods, methodInfo)
}

// VisitTerminal is called when a terminal node is visited.
func (s *JSAstParser) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *JSAstParser) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *JSAstParser) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *JSAstParser) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (s *JSAstParser) EnterFunctionExpression(ctx *js.FunctionExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *JSAstParser) ExitArgumentsExpression(ctx *js.ArgumentsExpressionContext) {}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *JSAstParser) EnterObjectLiteralExpression(ctx *js.ObjectLiteralExpressionContext) {
}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *JSAstParser) ExitClassDeclaration(ctx *js.ClassDeclarationContext) {}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *JSAstParser) ExitVariableDeclarationList(ctx *js.VariableDeclarationListContext) {
}

func (s *JSAstParser) EnterNewExpression(ctx *js.NewExpressionContext) {}
