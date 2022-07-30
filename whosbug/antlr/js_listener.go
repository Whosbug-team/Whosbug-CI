package antlr

import (
	javascript "git.woa.com/bkdevops/whosbug/antlr/jsLib"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *JSTreeShapeListener) EnterClassDeclaration(ctx *javascript.ClassDeclarationContext) {
	classInfo := ClassInfoType{
		StartLine: ctx.GetStart().GetLine(),
		EndLine:   ctx.GetStop().GetLine(),
		ClassName: findJsDeclChain(ctx) + ctx.Identifier().GetText(),
		Extends:   getExtendIdentifier(ctx),
	}
	s.AstInfoList.Classes = append(s.AstInfoList.Classes, classInfo)
}

func getExtendIdentifier(ctx *javascript.ClassDeclarationContext) (extend string) {
	if ctx.ClassTail() != nil {
		tempCtx := ctx.ClassTail().(*javascript.ClassTailContext).SingleExpression()
		for {
			if tempCtx == nil {
				return
			}
			if _, ok := tempCtx.(*javascript.ArgumentsExpressionContext); ok {
				tempCtx = tempCtx.(*javascript.ArgumentsExpressionContext).SingleExpression()
				continue
			}
			if _, ok := tempCtx.(*javascript.IdentifierExpressionContext); ok {
				extend = tempCtx.(*javascript.IdentifierExpressionContext).GetText()
				return
			}
			return
		}
	}
	return
}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *JSTreeShapeListener) ExitFunctionDeclaration(ctx *javascript.FunctionDeclarationContext) {
	methodInfo := MethodInfoType{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		MethodName: findJsDeclChain(ctx) + ctx.Identifier().GetText(),
	}
	if ctx.FormalParameterList() != nil {
		methodInfo.Parameters = ctx.FormalParameterList().GetText()
	}
	s.AstInfoList.Methods = append(s.AstInfoList.Methods, methodInfo)
}

func findJsDeclChain(ctx antlr.ParseTree) (chain string) {
	tempCtx := ctx.GetParent()
	for {
		if _, ok := tempCtx.(*javascript.ClassDeclarationContext); ok {
			chain = tempCtx.(*javascript.ClassDeclarationContext).Identifier().GetText() + "." + chain
		}
		if _, ok := tempCtx.(*javascript.FunctionDeclarationContext); ok {
			chain = tempCtx.(*javascript.FunctionDeclarationContext).Identifier().GetText() + "." + chain
		}
		tempCtx = tempCtx.GetParent()
		if tempCtx == nil {
			return
		}
	}
}

// VisitTerminal is called when a terminal node is visited.
func (s *JSTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *JSTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *JSTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *JSTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (s *JSTreeShapeListener) EnterFunctionExpression(ctx *javascript.FunctionExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *JSTreeShapeListener) ExitArgumentsExpression(ctx *javascript.ArgumentsExpressionContext) {}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *JSTreeShapeListener) EnterObjectLiteralExpression(ctx *javascript.ObjectLiteralExpressionContext) {
}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *JSTreeShapeListener) ExitClassDeclaration(ctx *javascript.ClassDeclarationContext) {}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *JSTreeShapeListener) ExitVariableDeclarationList(ctx *javascript.VariableDeclarationListContext) {
}

func (s *JSTreeShapeListener) EnterNewExpression(ctx *javascript.NewExpressionContext) {}
