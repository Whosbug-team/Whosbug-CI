package antlrpack

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
	javascript "whosbugPack/antlrpack/js_lib"
)

// VisitTerminal is called when a terminal node is visited.
func (s *JSTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *JSTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *JSTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *JSTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (s *JSTreeShapeListener) EnterFunctionExpression(ctx *javascript.FunctionExpressionContext) {
	if s.ObjectInfo.ObjectName != "" {
		s.ObjectInfo.ObjFuncName = append(s.ObjectInfo.ObjFuncName, ctx.GetParent().GetChild(0).(antlr.ParseTree).GetText())
	}
}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *JSTreeShapeListener) ExitArgumentsExpression(ctx *javascript.ArgumentsExpressionContext) {
	if ctx.Arguments().GetChild(ctx.Arguments().GetChildCount()-1).(antlr.ParseTree).GetText() == ")" {
		var callInfo CallMethodType
		call := strings.Split(ctx.GetChild(0).(antlr.ParseTree).GetText(), "(")[0]
		if call == "Function" || call == "constructor" || call == "super" {
			return
		}
		if (strings.Contains(call, "this.") || strings.Contains(call, "super.")) && s.ClassInfo.ClassName != "" {
			call = strings.Replace(call, "this", s.ClassInfo.ClassName, -1)
			call = strings.Replace(call, "super", s.ClassInfo.Extends, -1)
		}
		callInfo.StartLine = ctx.GetStart().GetLine()

		for _, temp := range s.Declaration {
			if strings.Split(ctx.GetChild(0).(antlr.ParseTree).GetText(), ".")[0] == temp.Name {
				callInfo.Id = strings.Replace(call, temp.Name, temp.Type, 1)
				break
			}
		}

		if call == "" {
			return
		}

		if callInfo.Id == "" {
			callInfo.Id = call
		}

		for _, temp := range s.Infos.CallMethods {
			if callInfo.Id == temp.Id {
				return
			}
		}

		for _, temp := range s.Infos.AstInfoList.Classes {
			if call == temp.ClassName {
				return
			}
		}
		s.Infos.CallMethods = append(s.Infos.CallMethods, callInfo)
		//s.Infos.CallMethods = RemoveRep(s.Infos.CallMethods)
		//fmt.Println(s.Infos.CallMethods)
	}
}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *JSTreeShapeListener) EnterObjectLiteralExpression(ctx *javascript.ObjectLiteralExpressionContext) {
	if ctx.GetParent().GetChildCount() == 3 && ctx.GetParent().GetParent().GetChild(0).(antlr.ParseTree).GetText() == "var" {
		s.ObjectInfo.ObjectName = ctx.GetParent().GetChild(0).(antlr.ParseTree).GetText()
	}
}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *JSTreeShapeListener) ExitClassDeclaration(ctx *javascript.ClassDeclarationContext) {
	var cal classInfoType
	s.ClassInfo = cal
}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *JSTreeShapeListener) EnterClassDeclaration(ctx *javascript.ClassDeclarationContext) {
	s.ClassInfo.StartLine = ctx.GetStart().GetLine()
	s.ClassInfo.EndLine = ctx.GetStop().GetLine()

	s.ClassInfo.ClassName = ctx.GetChild(1).(antlr.ParseTree).GetText()
	if ctx.ClassTail().GetChild(0).(antlr.ParseTree).GetText() == "extends" {
		s.ClassInfo.Extends = ctx.ClassTail().GetChild(1).(antlr.ParseTree).GetText()
	} else {
		s.ClassInfo.Extends = "None"
	}
	s.Infos.AstInfoList.Classes = append(s.Infos.AstInfoList.Classes, s.ClassInfo)
}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *JSTreeShapeListener) ExitFunctionDeclaration(ctx *javascript.FunctionDeclarationContext) {
	var funcInfo MethodInfoType
	funcInfo.StartLine = ctx.GetStart().GetLine()
	funcInfo.EndLine = ctx.GetStop().GetLine()
	funcInfo.MethodName = ctx.Identifier().GetText()
	if ctx.FormalParameterList() != nil {
		i := 0
		for i < ctx.FormalParameterList().GetChildCount() {
			funcInfo.Params = append(funcInfo.Params, ctx.FormalParameterList().GetChild(i).(antlr.ParseTree).GetText())
			i += 2
		}
	}
	resIndex := s.FindMethodCallIndex(funcInfo.StartLine, funcInfo.EndLine)
	if resIndex != nil {
		funcInfo.CallMethods = RemoveRep(resIndex)
	}
	s.Infos.AstInfoList.Methods = append(s.Infos.AstInfoList.Methods, funcInfo)
	//fmt.Println(s.Infos.AstInfoList.Methods)
}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *JSTreeShapeListener) ExitVariableDeclarationList(ctx *javascript.VariableDeclarationListContext) {
	if s.ObjectInfo.ObjectName != "" {
		s.ObjectInfo.StartLine = ctx.GetStart().GetLine()
		s.ObjectInfo.EndLine = ctx.GetStop().GetLine()
		s.Infos.AstInfoList.Objects = append(s.Infos.AstInfoList.Objects, s.ObjectInfo)
		//fmt.Println(s.Infos.AstInfoList.Objects)
		var obj ObjectInfoType
		s.ObjectInfo = obj
	}
}

func (s *JSTreeShapeListener) FindMethodCallIndex(targetStart, targetEnd int) []string {
	var resIndex []string
	for index := range s.Infos.CallMethods {
		if s.Infos.CallMethods[index].StartLine <= targetEnd && s.Infos.CallMethods[index].StartLine >= targetStart {
			resIndex = append(resIndex, s.Infos.CallMethods[index].Id)
		}
	}
	return resIndex
}

func (s *JSTreeShapeListener) EnterNewExpression(ctx *javascript.NewExpressionContext) {
	for _, temp := range s.Infos.AstInfoList.Classes {
		if strings.Split(ctx.GetChild(1).(antlr.ParseTree).GetText(), "(")[0] == temp.ClassName {
			var member MemberType
			member.Name = ctx.GetParent().GetChild(0).(antlr.ParseTree).GetText()
			member.Type = temp.ClassName
			s.Declaration = append(s.Declaration, member)
		}
	}
}
