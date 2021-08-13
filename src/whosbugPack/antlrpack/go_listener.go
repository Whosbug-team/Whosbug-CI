package antlrpack

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
	golang "whosbugPack/antlrpack/go_lib"
	"whosbugPack/utility"
)

//只能解析非struct方法
func (s *GoTreeShapeListener) EnterFunctionDecl(ctx *golang.FunctionDeclContext) {
	var funcInfo MethodInfoType
	funcInfo.MethodName = ctx.GetChild(1).(antlr.ParseTree).GetText()
	funcInfo.StartLine = ctx.GetStart().GetLine()
	funcInfo.EndLine = ctx.GetStop().GetLine()
	funcInfo.MasterObject = masterObjectInfoType{}
	funcInfo.CallMethods = s.findMethodCall("")

	s.Infos.CallMethods = []CallMethodType{}
}
//struct方法,但是没办法分出对应的struct
func (s *GoTreeShapeListener) EnterMethodDecl(ctx *golang.MethodDeclContext) {
	var funcInfo MethodInfoType

	funcInfo.MethodName = ctx.GetChild(2).(antlr.ParseTree).GetText()
	funcInfo.StartLine = ctx.GetStart().GetLine()
	funcInfo.EndLine = ctx.GetStop().GetLine()

	structName := ctx.GetChild(1).GetChild(0).GetChild(1)
	var struct_belong string
	if structName.GetChildCount() == 2{ //func (c *Cube) Area(...),有设置别名
		struct_belong = strings.Trim(structName.GetChild(1).(antlr.ParseTree).GetText(),"*")
	}else {
		struct_belong = strings.Trim(structName.GetChild(0).(antlr.ParseTree).GetText(),"*")
	}
	Structs := s.Infos.AstInfoList.Classes
	for i := 0; i < len(Structs); i++ {
		if Structs[i].ClassName == struct_belong {
			var masterObject = masterObjectInfoType{
				StartLine: s.Infos.AstInfoList.Classes[i].StartLine,
				ObjectName: struct_belong,
			}
			funcInfo.MasterObject = masterObject
			break
		}
	}
	funcInfo.CallMethods = s.findMethodCall(struct_belong)

	fmt.Printf("StructFuncs：%+v,belongs:%s已添加!!!!\n",funcInfo,struct_belong)

}

func (s *GoTreeShapeListener) EnterExpressionStmt(ctx *golang.ExpressionStmtContext) {
	var callMethod = CallMethodType{
		StartLine: ctx.GetStart().GetLine(),
		Id:        ctx.GetText(),
	}
	s.Infos.CallMethods = append(s.Infos.CallMethods, callMethod)
}

func (s *GoTreeShapeListener) findMethodCall(struct_belong string) []string{
	var struct_methods []string
	for index := range s.Infos.CallMethods {
		struct_methods = append(struct_methods,utility.ConCatStrings(struct_belong, ".", s.Infos.CallMethods[index].Id))
	}
	return struct_methods
}

//退出Struct，把s.structInfo包装好然后上传
//Function在FunctionDecl包装,Member在EnterType_里包装
func (s *GoTreeShapeListener) EnterStructType(ctx *golang.StructTypeContext) {
	var structInfo = classInfoType{
		StartLine:    ctx.GetStart().GetLine(),
		EndLine:      ctx.GetStop().GetLine(),
		ClassName:    ctx.GetChild(0).(antlr.ParseTree).GetText(),
		MasterObject: masterObjectInfoType{},
	}
	s.Infos.AstInfoList.Classes = append(s.Infos.AstInfoList.Classes, structInfo)
}

func (s *GoTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
}

func (s *GoTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (s *GoTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (s *GoTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}