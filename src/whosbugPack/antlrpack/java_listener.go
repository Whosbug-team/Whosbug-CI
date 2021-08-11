package antlrpack

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	javaparser "whosbugPack/antlrpack/java_lib"
	"whosbugPack/utility"
)

type MethodInfoType struct {
	StartLine    int
	EndLine      int
	MethodName   string
	MasterObject masterObjectInfoType
	CallMethods  []string
}

type masterObjectInfoType struct {
	ObjectName string
	StartLine  int
}

type classInfoType struct {
	StartLine    int
	EndLine      int
	ClassName    string
	MasterObject masterObjectInfoType
}

type astInfoType struct {
	Classes []classInfoType
	Methods []MethodInfoType
}
type CallMethodType struct {
	StartLine int
	Id        string
}
type AnalysisInfoType struct {
	CallMethods []CallMethodType
	AstInfoList astInfoType
}

// ExitMethodDeclaration
//	@Description: 匹配到方法结束时被调用
//	@receiver s
//	@param ctx
//	@author KevinMatt 2021-07-23 23:14:09
//	@function_mark PASS
func (s *TreeShapeListener) ExitMethodDeclaration(ctx *javaparser.MethodDeclarationContext) {
	var methodInfo MethodInfoType
	if ctx.GetChildCount() >= 2 {
		MethodName := ctx.GetChild(1).(antlr.ParseTree).GetText()
		methodInfo = MethodInfoType{
			StartLine:    ctx.GetStart().GetLine(),
			EndLine:      ctx.GetStop().GetLine(),
			MethodName:   MethodName,
			MasterObject: findMasterObjectClass(ctx),
		}
		resIndex := s.FindMethodCallIndex(methodInfo.StartLine, methodInfo.EndLine)
		if resIndex != nil {
			methodInfo.CallMethods = resIndex
		}
		s.Infos.AstInfoList.Methods = append(s.Infos.AstInfoList.Methods, methodInfo)
	}
}

func (s *TreeShapeListener) FindMethodCallIndex(targetStart, targetEnd int) []string {
	var resIndex []string
	for index := range s.Infos.CallMethods {
		if s.Infos.CallMethods[index].StartLine <= targetEnd && s.Infos.CallMethods[index].StartLine >= targetStart {
			resIndex = append(resIndex, s.Infos.CallMethods[index].Id)
		}
	}
	return resIndex
}

// EnterMethodCall
//	@Description: 匹配调用方法行并获取起始行列号
//	@receiver s
//	@param ctx
//	@author KevinMatt 2021-07-23 23:22:56
//	@function_mark PASS
func (s *TreeShapeListener) EnterMethodCall(ctx *javaparser.MethodCallContext) {
	if ctx.GetParent() != nil {
		newMasterObject := findMasterObjectClass(ctx)
		var insertTemp = CallMethodType{
			StartLine: ctx.GetStart().GetLine(),
			Id:        utility.ConCatStrings(newMasterObject.ObjectName, ".", ctx.GetParent().(antlr.ParseTree).GetText()),
		}
		s.Infos.CallMethods = append(s.Infos.CallMethods, insertTemp)
	}
}

// EnterClassDeclaration
//	@Description: 类对象匹配
//	@receiver s
//	@param ctx
//	@author KevinMatt 2021-07-24 11:44:50
//	@function_mark PASS
func (s *TreeShapeListener) EnterClassDeclaration(ctx *javaparser.ClassDeclarationContext) {
	var classInfo classInfoType
	childCount := ctx.GetChildCount()

	if childCount == 6 {
		className := ctx.GetChild(1).(antlr.ParseTree).GetText()
		classInfo.ClassName = className
	} else if childCount == 5 {
		className := ctx.GetChild(1).(antlr.ParseTree).GetText()
		classInfo.ClassName = className
	} else if childCount == 4 {
		// Generic classes: class AnnoyName<T>
		// 此处没有解析尖括号内的内容，其内如有继承关系，将一起连接被打印
		className := ctx.GetChild(1).(antlr.ParseTree).GetText()
		for index := 0; index < ctx.GetChild(2).GetChildCount(); index++ {
			if index%2 == 0 {
				className += "" + ctx.GetChild(2).GetChild(index).(antlr.ParseTree).GetText()
			} else {
				className += " " + ctx.GetChild(2).GetChild(index).(antlr.ParseTree).GetText()
			}
		}
		classInfo.ClassName = className
	} else if childCount == 3 {
		className := ctx.GetChild(1).(antlr.ParseTree).GetText()
		classInfo.ClassName = className
	}

	classInfo.StartLine = ctx.GetStart().GetLine()
	if ctx.ClassBody() != nil {
		classInfo.EndLine = ctx.ClassBody().GetStop().GetLine()
	}
	classInfo.MasterObject = findMasterObjectClass(ctx)
	s.Infos.AstInfoList.Classes = append(s.Infos.AstInfoList.Classes, classInfo)
}

// findMasterObjectClass
//	@Description: 找到主类实体
//	@param ctx
//	@param classInfo
//	@author KevinMatt 2021-07-24 11:45:14
//	@function_mark PASS
func findMasterObjectClass(ctx antlr.ParseTree) masterObjectInfoType {
	temp := ctx.GetParent()
	if temp == nil {
		return masterObjectInfoType{}
	}
	var masterObject masterObjectInfoType
	for {
		if _, ok := temp.(*javaparser.ClassDeclarationContext); ok {
			masterObject.ObjectName = temp.GetChild(1).GetText()
			masterObject.StartLine = temp.GetChild(temp.GetChildCount() - 1).(*javaparser.ClassBodyContext).GetStart().GetLine()
			return masterObject
		}
		temp = temp.GetParent()
		if temp == nil {
			return masterObjectInfoType{}
		}
	}
}

// VisitTerminal is called when a terminal node is visited.
func (s *TreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *TreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}
