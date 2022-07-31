package antlr

import (
	javaparser "git.woa.com/bkdevops/whosbug/antlr/javaLib"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// EnterClassDeclaration 获取类声明/定义链
//  @receiver s
//  @param ctx
func (s *JavaTreeShapeListener) EnterClassDeclaration(ctx *javaparser.ClassDeclarationContext) {
	classInfo := ClassInfoType{
		ClassName: findJavaDeclarationChain(ctx) + ctx.IDENTIFIER().GetText(),
		StartLine: ctx.ClassBody().GetStart().GetLine(),
		EndLine:   ctx.ClassBody().GetStop().GetLine(),
		Extends:   findExtendsClass(ctx),
	}
	s.AstInfoList.Classes = append(s.AstInfoList.Classes, classInfo)
}

// ExitMethodDeclaration 获取方法声明/定义链
//  @receiver s
//  @param ctx
func (s *JavaTreeShapeListener) ExitMethodDeclaration(ctx *javaparser.MethodDeclarationContext) {
	methodInfo := MethodInfoType{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		MethodName: findJavaDeclarationChain(ctx) + ctx.IDENTIFIER().GetText(),
		Parameters: getParamsOfMethod(ctx),
	}
	s.AstInfoList.Methods = append(s.AstInfoList.Methods, methodInfo)
}

func findExtendsClass(ctx *javaparser.ClassDeclarationContext) (extendsClass string) {
	if ctx.EXTENDS() != nil {
		extendsClass = ctx.TypeType().GetText()
	}
	return
}

func getParamsOfMethod(ctx *javaparser.MethodDeclarationContext) (params string) {
	if ctx.FormalParameters() != nil {
		if ctx.FormalParameters().(*javaparser.FormalParametersContext).FormalParameterList() != nil {
			for index, item := range ctx.FormalParameters().(*javaparser.FormalParametersContext).FormalParameterList().(*javaparser.FormalParameterListContext).AllFormalParameter() {
				params += item.(*javaparser.FormalParameterContext).TypeType().GetText() + " "
				params += item.(*javaparser.FormalParameterContext).VariableDeclaratorId().GetText()
				if index != len(ctx.FormalParameters().(*javaparser.FormalParametersContext).FormalParameterList().(*javaparser.FormalParameterListContext).AllFormalParameter())-1 {
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
		if _, ok := currentContext.(*javaparser.ClassDeclarationContext); ok {
			chainName = currentContext.(*javaparser.ClassDeclarationContext).IDENTIFIER().GetText() + "." + chainName
		}
		if _, ok := currentContext.(*javaparser.MethodDeclarationContext); ok {
			chainName = currentContext.(*javaparser.MethodDeclarationContext).IDENTIFIER().GetText() + "." + chainName
		}
		currentContext = currentContext.GetParent()
		if currentContext == nil {
			break
		}
	}
	return
}

// EnterLocalVariableDeclaration
//  @receiver s
//  @param ctx
//  @return {}
func (s *JavaTreeShapeListener) EnterLocalVariableDeclaration(ctx *javaparser.LocalVariableDeclarationContext) {
}

// VisitTerminal is called when a terminal node is visited.
func (s *JavaTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *JavaTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *JavaTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *JavaTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (s *JavaTreeShapeListener) EnterMethodCall(ctx *javaparser.MethodCallContext) {}

func (s *JavaTreeShapeListener) EnterVariableDeclarator(ctx *javaparser.VariableDeclaratorContext) {}
