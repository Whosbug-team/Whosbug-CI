package antlr

import (
	"sync"

	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr/java"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// JavaAstParser implement AstParser for java language
//
//	@author kevineluo
//	@update 2023-02-28 01:09:55
type JavaAstParser struct {
	AstInfo AstInfo
}

var (
	_ java.JavaParserListener = &JavaAstParser{}
	_ AstParser               = &JavaAstParser{}
)

var (
	javaLexerPool = &sync.Pool{New: func() any {
		return java.NewJavaLexer(nil)
	}}
	javaParserPool = &sync.Pool{New: func() any {
		return java.NewJavaParser(nil)
	}}
	newJavaAstParserPool = &sync.Pool{New: func() any {
		return new(JavaAstParser)
	}}
)

// Init 初始化AstParser
//
//	@receiver s *CAstParser
//	@author kevineluo
//	@update 2023-02-28 03:13:43
func (s *JavaAstParser) Init() (err error) {
	s.AstInfo.Classes = make([]Class, 0)
	s.AstInfo.Methods = make([]Method, 0)
	return
}

// AstParse main parse process for java language
//
//	@receiver s *JavaAstParser
//	@param diffText string
//	@return AstInfo
//	@author kevineluo
//	@update 2023-02-28 01:11:10
func (s *JavaAstParser) AstParse(diffText string) AstInfo {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := javaLexerPool.Get().(*java.JavaLexer)
	defer javaLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := javaParserPool.Get().(*java.JavaParser)
	p.RemoveErrorListeners()
	defer javaParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.CompilationUnit()
	//	创建listener
	listener := newJavaAstParserPool.Get().(*JavaAstParser)
	defer newJavaAstParserPool.Put(listener)
	// 初始化置空
	listener.AstInfo = *new(AstInfo)
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.AstInfo
}

func findExtendsClass(ctx *java.ClassDeclarationContext) (extendsClass string) {
	if ctx.EXTENDS() != nil {
		extendsClass = ctx.TypeType().GetText()
	}
	return
}

func getParamsOfMethod(ctx *java.MethodDeclarationContext) (params string) {
	if ctx.FormalParameters() != nil {
		if ctx.FormalParameters().(*java.FormalParametersContext).FormalParameterList() != nil {
			for index, item := range ctx.FormalParameters().(*java.FormalParametersContext).FormalParameterList().(*java.FormalParameterListContext).AllFormalParameter() {
				params += item.(*java.FormalParameterContext).TypeType().GetText() + " "
				params += item.(*java.FormalParameterContext).VariableDeclaratorId().GetText()
				if index != len(ctx.FormalParameters().(*java.FormalParametersContext).FormalParameterList().(*java.FormalParameterListContext).AllFormalParameter())-1 {
					params += ", "
				}
			}
		}
	}
	return
}

func findJavaDeclarationChain(ctx antlr.ParseTree) (chainName string) {
	currentContext := ctx.GetParent()
	for {
		if _, ok := currentContext.(*java.ClassDeclarationContext); ok {
			chainName = currentContext.(*java.ClassDeclarationContext).IDENTIFIER().GetText() + "." + chainName
		}
		if _, ok := currentContext.(*java.MethodDeclarationContext); ok {
			chainName = currentContext.(*java.MethodDeclarationContext).IDENTIFIER().GetText() + "." + chainName
		}
		currentContext = currentContext.GetParent()
		if currentContext == nil {
			break
		}
	}
	return
}

// EnterClassDeclaration 获取类声明/定义链
//
//	@receiver s
//	@param ctx
func (s *JavaAstParser) EnterClassDeclaration(ctx *java.ClassDeclarationContext) {
	classInfo := Class{
		Name:      findJavaDeclarationChain(ctx) + ctx.IDENTIFIER().GetText(),
		StartLine: ctx.ClassBody().GetStart().GetLine(),
		EndLine:   ctx.ClassBody().GetStop().GetLine(),
		Extends:   findExtendsClass(ctx),
	}
	s.AstInfo.Classes = append(s.AstInfo.Classes, classInfo)
}

// ExitMethodDeclaration 获取方法声明/定义链
//
//	@receiver s
//	@param ctx
func (s *JavaAstParser) ExitMethodDeclaration(ctx *java.MethodDeclarationContext) {
	methodInfo := Method{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		Name:       findJavaDeclarationChain(ctx) + ctx.IDENTIFIER().GetText(),
		Parameters: getParamsOfMethod(ctx),
	}
	s.AstInfo.Methods = append(s.AstInfo.Methods, methodInfo)
}

// EnterLocalVariableDeclaration
func (s *JavaAstParser) EnterLocalVariableDeclaration(ctx *java.LocalVariableDeclarationContext) {
}

// VisitTerminal is called when a terminal node is visited.
func (s *JavaAstParser) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *JavaAstParser) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *JavaAstParser) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *JavaAstParser) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (s *JavaAstParser) EnterMethodCall(ctx *java.MethodCallContext) {}

func (s *JavaAstParser) EnterVariableDeclarator(ctx *java.VariableDeclaratorContext) {}
