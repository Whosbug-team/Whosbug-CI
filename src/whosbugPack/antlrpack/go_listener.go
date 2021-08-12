package antlrpack

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	golang "whosbugPack/antlrpack/go_lib"
	"strings"
	"fmt"
)


type typeInfoType struct {
	Name string
	StartLine int
	EndLine int
}

//type AnalysisInfoType struct {
//	CallMethods []CallMethodType
//	AstInfoList astInfoType
//}

type astInfoType struct {
	Package   string
	Import    []string
	Funcs     []funcInfoType
	Const     []memberInfoType
	Var       []memberInfoType    //局部变量和全局变量还没去区分
	Type      []memberInfoType    //类型定义
	Struct    []structInfoType    //结构体定义
	Interface []interfaceInfoType //接口定义
	Classes   interface{}
}

type interfaceInfoType struct{
	typeInfoType
	InterFuncs []funcInfoType
	//Inhert []interfaceInfoType	"继承“还没法解析
}

type structInfoType struct {
	typeInfoType
	Memebers   []memberInfoType
	StructFuncs      []funcInfoType
	//Inhert 	[]structInfoType   “继承”还没法解析
}

type memberInfoType struct {
	Name string
	Value string
}

type funcInfoType struct {
	StartLine int
	EndLine   int
	FuncName  string
	Params    []memberInfoType
	ReturnType string
	//Depth        int
}

// EnterPackageClause is called when production packageClause is entered.
func (s *GoTreeShapeListener) EnterPackageClause(ctx *golang.PackageClauseContext) {
	packages := ctx.GetChild(1).(antlr.ParseTree).GetText()
	s.Infos.AstInfoList.Package = packages
	fmt.Printf("Package: %s \n",packages)
}

func (s *GoTreeShapeListener) EnterImportPath(ctx *golang.ImportPathContext) {
	importPtah := ctx.GetText()
	fmt.Printf("Imporots_path:%s \n",importPtah)
	s.Infos.AstInfoList.Import = append(s.Infos.AstInfoList.Import, importPtah)
}

func (s *GoTreeShapeListener) EnterConstSpec(ctx *golang.ConstSpecContext) {
	var memberInfo memberInfoType
	constName := strings.Split(ctx.GetChild(0).(antlr.ParseTree).GetText(),",")
	constValue := strings.Split(ctx.GetChild(2).(antlr.ParseTree).GetText(),",")
	for i := 0;i < len(constName);i++{
		memberInfo.Name = constName[i]
		memberInfo.Value = constValue[i]
		s.Infos.AstInfoList.Const = append(s.Infos.AstInfoList.Const, memberInfo)
		fmt.Printf("Const:%+v \n",memberInfo)
	}

}
//解析的是a = b中的a 或 a int中的a
func (s *GoTreeShapeListener) EnterIdentifierList(ctx *golang.IdentifierListContext) {
	s.memberInfo = memberInfoType{}
	switch s.InType{
	case "struct":
		s.memberInfo.Name = ctx.GetText()
		//fmt.Printf("INSTRUCT_Key:%s \n",s.memberInfo.Name)
	case "interface":
	default:

	}
}
//TODO 只能解析Struct的Type,但是不能解析内容，需要到Expression和Identifier解析
//这里有很多需要补充的，只能针对类型定义与结构体定义的区分，有如接口定义等还需要更为细致的划分
//对于EnterTypeSpec是先于所有Type子解析方法（EnterStructType等），那么进入EnterStructType后再置s.InType为struct
func (s *GoTreeShapeListener) EnterTypeSpec(ctx *golang.TypeSpecContext) {
	s.typeInfo.Name = ctx.GetChild(0).(antlr.ParseTree).GetText()
	s.typeInfo.StartLine = ctx.GetStart().GetLine()
	s.typeInfo.EndLine = ctx.GetStop().GetLine()
}

func (s *GoTreeShapeListener) ExitTypeSpec(ctx *golang.TypeSpecContext) {
	s.InType = ""
	//全局变量清空，以免第一次的解析结果留存到其他结构里面去
	s.structInfo = structInfoType{}
	s.interfaceInfo = interfaceInfoType{}
	s.memberInfo = memberInfoType{}
}
//只能解析非struct方法
func (s *GoTreeShapeListener) EnterFunctionDecl(ctx *golang.FunctionDeclContext) {
	var funcInfo funcInfoType
	funcInfo.FuncName = ctx.GetChild(1).(antlr.ParseTree).GetText()
	funcInfo.StartLine = ctx.GetStart().GetLine()
	funcInfo.EndLine = ctx.GetStop().GetLine()
	funcSignature := ctx.GetChild(2)		//函数签名
	if funcSignature.GetChildCount() == 2{ //该函数有返回类型
		funcInfo.ReturnType = funcSignature.GetChild(1).(antlr.ParseTree).GetText()
	}
	paramContext := funcSignature.GetChild(0)
	if 2 < paramContext.GetChildCount() {
		for i := 1; i < paramContext.GetChildCount(); {
			paramKeys := strings.Split(paramContext.GetChild(i).GetChild(0).(antlr.ParseTree).GetText(),",")
			paramType := paramContext.GetChild(i).GetChild(1).(antlr.ParseTree).GetText()
			for j := 0; j < len(paramKeys); j++ {
				funcInfo.Params = append(funcInfo.Params,
					memberInfoType{Name: paramKeys[j],Value: paramType})
			}
			i += 2
		}
	}
	s.Infos.AstInfoList.Funcs = append(s.Infos.AstInfoList.Funcs, funcInfo)
	fmt.Printf("Funcs：%+v已添加!!!!\n",funcInfo)
}
//struct方法,但是没办法分出对应的struct
func (s *GoTreeShapeListener) EnterMethodDecl(ctx *golang.MethodDeclContext) {
	var funcInfo funcInfoType

	//fmt.Println(ctx.GetChild(1).GetChild(0).GetChild(1).GetChildCount())
	//for i := 0; i < ctx.GetChildCount(); i++ {
	//	stmt := ctx.GetChild(i).(antlr.ParseTree).GetText()
	//	fmt.Printf("=====%d:%s=====\n",i,stmt)
	//}
	funcInfo.FuncName = ctx.GetChild(2).(antlr.ParseTree).GetText()
	funcInfo.StartLine = ctx.GetStart().GetLine()
	funcInfo.EndLine = ctx.GetStop().GetLine()
	funcSignature := ctx.GetChild(3)		//函数签名
	if funcSignature.GetChildCount() == 2{ //该函数有返回类型
		funcInfo.ReturnType = funcSignature.GetChild(1).(antlr.ParseTree).GetText()
	}
	paramContext := funcSignature.GetChild(0)
	if 2 < paramContext.GetChildCount() {
		for i := 1; i < paramContext.GetChildCount(); {
			paramKeys := strings.Split(paramContext.GetChild(i).GetChild(0).(antlr.ParseTree).GetText(),",")
			paramType := paramContext.GetChild(i).GetChild(1).(antlr.ParseTree).GetText()
			for j := 0; j < len(paramKeys); j++ {
				funcInfo.Params = append(funcInfo.Params,
					memberInfoType{Name: paramKeys[j],Value: paramType})
			}
			i += 2
		}
	}
	structName := ctx.GetChild(1).GetChild(0).GetChild(1)
	var struct_belong string
	if structName.GetChildCount() == 2{ //func (c *Cube) Area(...),有设置别名
		struct_belong = strings.Trim(structName.GetChild(1).(antlr.ParseTree).GetText(),"*")
	}else {
		struct_belong = strings.Trim(structName.GetChild(0).(antlr.ParseTree).GetText(),"*")
	}
	Structs := s.Infos.AstInfoList.Struct
	for i := 0; i < len(Structs); i++ {
		if Structs[i].Name == struct_belong {
			s.Infos.AstInfoList.Struct[i].StructFuncs = append(
				s.Infos.AstInfoList.Struct[i].StructFuncs, funcInfo)
		}
	}
	fmt.Printf("StructFuncs：%+v,belongs:%s已添加!!!!\n",funcInfo,struct_belong)

}
// EnterVarSpec is called when production varSpec is entered.
func (s *GoTreeShapeListener) EnterVarSpec(ctx *golang.VarSpecContext) {
	s.memberInfo = memberInfoType{}
	VarName := strings.Split(ctx.GetChild(0).(antlr.ParseTree).GetText(),",")
	VarType := ctx.GetChild(1).(antlr.ParseTree).GetText()
	for i := 0; i < len(VarName); i++ {
		s.memberInfo.Name = VarName[i]
		s.memberInfo.Value = VarType
		s.Infos.AstInfoList.Var = append(s.Infos.AstInfoList.Var, s.memberInfo)
	}
	//fmt.Printf("====VarName:%s,VarType:%s===\n",VarName,VarType)
}
// Struct成员变量类型，但是*[]int会重复三次，结果分别为*[]int,[]int,int，只取首个结果
func (s *GoTreeShapeListener) EnterType_(ctx *golang.Type_Context) {
	switch s.InType{
	case "struct":
		if "" != s.memberInfo.Name{
			//取首个结果匹配
			s.memberInfo.Value = ctx.GetText()
			//fmt.Printf("INSTRUCT_Type:%s \n",s.memberInfo.Value)
			s.structInfo.Memebers = append(s.structInfo.Memebers, s.memberInfo)
			fmt.Printf("memberInfo:%+v已添加！！！\n",s.memberInfo)
			s.memberInfo = memberInfoType{}
		}//否则就是多余的解析结果，退出的时候给memberInfo清空了
	case "interface":
		//stmt := ctx.GetText()
		//fmt.Printf("====Interface:%s====\n",stmt)
	default:

	}
}
func (s *GoTreeShapeListener) EnterInterfaceType(ctx *golang.InterfaceTypeContext) {
	s.InType = "interface"
	fmt.Printf("EnterInterface:%s \n",s.typeInfo.Name)
}
// ExitInterfaceType is called when production interfaceType is exited.
func (s *GoTreeShapeListener) ExitInterfaceType(ctx *golang.InterfaceTypeContext) {
	s.interfaceInfo.Name = s.typeInfo.Name
	s.interfaceInfo.StartLine = s.typeInfo.StartLine
	s.interfaceInfo.EndLine = s.typeInfo.EndLine

	s.Infos.AstInfoList.Interface = append(s.Infos.AstInfoList.Interface, s.interfaceInfo)
	fmt.Printf("ExitInterface:%s \n",s.typeInfo.Name)
	//s.interfaceInfo = s.interfaceInfoType{}
}
//专门生成interface中的函数声明解析
func (s *GoTreeShapeListener) EnterMethodSpec(ctx *golang.MethodSpecContext) {
	if s.InType == "interface"{
		var funcInfo funcInfoType
		funcInfo.StartLine = ctx.GetStart().GetLine()
		funcInfo.EndLine = ctx.GetStop().GetLine()
		funcInfo.FuncName = ctx.GetChild(0).(antlr.ParseTree).GetText()
		funcInfo.ReturnType = ctx.GetChild(2).(antlr.ParseTree).GetText()
		paramCount := ctx.GetChild(1).GetChildCount()
		if 2 < paramCount {
			for i := 1; i < paramCount; {
				param := ctx.GetChild(1).GetChild(i)
				paramName := strings.Split(param.GetChild(0).(antlr.ParseTree).GetText(), ",")
				paramType := param.GetChild(1).(antlr.ParseTree).GetText()
				for j := 0; j < len(paramName); j++ {
					funcInfo.Params = append(funcInfo.Params,
						memberInfoType{Name: paramName[j], Value: paramType})
				}
				i += 2
			}
		}
		s.interfaceInfo.InterFuncs = append(s.interfaceInfo.InterFuncs, funcInfo)
	}
}

//这个无法解析出Struct的名字！！解析名字，起始行的任务交给EnterTypeSpec
func (s *GoTreeShapeListener) EnterStructType(ctx *golang.StructTypeContext) {
	s.InType = "struct"
	fmt.Printf("EnterStruct:%s \n",s.typeInfo.Name)
	//进入Struct了，解析Struct内容的任务交给EnterIdentifierList
}
//退出Struct，把s.structInfo包装好然后上传
//Function在FunctionDecl包装,Member在EnterType_里包装
func (s *GoTreeShapeListener) ExitStructType(ctx *golang.StructTypeContext) {
	s.structInfo.Name = s.typeInfo.Name
	s.structInfo.StartLine = s.typeInfo.StartLine
	s.structInfo.EndLine = s.typeInfo.EndLine

	s.Infos.AstInfoList.Struct = append(s.Infos.AstInfoList.Struct, s.structInfo)
	fmt.Printf("ExitStruct:%s \n",s.typeInfo.Name)
	//structInfo = structInfoType{}
}

func (s *GoTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
}

func (s *GoTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (s *GoTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (s *GoTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}