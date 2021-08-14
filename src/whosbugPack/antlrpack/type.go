package antlrpack

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type JavaTreeShapeListener struct {
	Infos AnalysisInfoType
}

type KotlinTreeShapeListener struct {
	MethodInfo MethodInfoType
	Infos      AnalysisInfoType
}

type CppTreeShapeListener struct{
	Infos AnalysisInfoType
}

type GoTreeShapeListener struct{
	Infos AnalysisInfoType
}

type JSTreeShapeListener struct {
	ObjectInfo ObjectInfoType
	ClassInfo  classInfoType
	Infos      AnalysisInfoType
}

func (k KotlinTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
	panic("implement me")
}

func (k KotlinTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
	panic("implement me")
}

func (k KotlinTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}

func (k KotlinTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	panic("implement me")
}


type CallMethodType struct {
	StartLine int
	Id        string
}

type AnalysisInfoType struct {
	CallMethods []CallMethodType
	AstInfoList astInfoType
}

type astInfoType struct {
	Classes []classInfoType
	Methods []MethodInfoType
	Objects []ObjectInfoType
}

type classInfoType struct {
	StartLine int
	EndLine   int
	ClassName string
	MasterObject masterObjectInfoType
	Extends      string
}

type MethodInfoType struct {
	StartLine    int
	EndLine      int
	MethodName   string
	MasterObject masterObjectInfoType
	CallMethods  []string
	Params       []string
}

type masterObjectInfoType struct {
	ObjectName string
	StartLine  int
}

type ObjectInfoType struct {
	StartLine   int
	EndLine     int
	ObjectName  string
	ObjFuncName []string
}