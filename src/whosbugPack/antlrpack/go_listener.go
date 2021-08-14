package antlrpack

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
	golang "whosbugPack/antlrpack/go_lib"
	"whosbugPack/utility"
)

//只能解析非struct方法
func (s *GoTreeShapeListener) ExitFunctionDecl(ctx *golang.FunctionDeclContext) {
	var funcInfo MethodInfoType
	funcInfo.MethodName = ctx.GetChild(1).(antlr.ParseTree).GetText()
	funcInfo.StartLine = ctx.GetStart().GetLine()
	funcInfo.EndLine = ctx.GetStop().GetLine()
	funcInfo.MasterObject = masterObjectInfoType{}
	funcInfo.CallMethods = s.findMethodCall("")

	s.Infos.AstInfoList.Methods = append(s.Infos.AstInfoList.Methods, funcInfo)
	s.Infos.CallMethods = []CallMethodType{}
}
//struct方法,但是没办法分出对应的struct
func (s *GoTreeShapeListener) ExitMethodDecl(ctx *golang.MethodDeclContext) {
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
	var masterObject = masterObjectInfoType{
		StartLine: 0,
		ObjectName: struct_belong,
	}
	funcInfo.MasterObject = masterObject
	funcInfo.CallMethods = s.findMethodCall(struct_belong)

	s.Infos.AstInfoList.Methods = append(s.Infos.AstInfoList.Methods, funcInfo)
	s.Infos.CallMethods = []CallMethodType{}
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

func (s *GoTreeShapeListener) EnterStructType(ctx *golang.StructTypeContext) {

	var structInfo = classInfoType{
		StartLine:    ctx.GetStart().GetLine(),
		EndLine:      ctx.GetStop().GetLine(),
		MasterObject: findGoMasterObjectClass(ctx),
	}
	structInfo.ClassName = structInfo.MasterObject.ObjectName
	s.Infos.AstInfoList.Classes = append(s.Infos.AstInfoList.Classes, structInfo)
}

func findGoMasterObjectClass(ctx antlr.ParseTree) masterObjectInfoType {
	temp := ctx.GetParent()
	if temp == nil {
		return masterObjectInfoType{}
	}
	var masterObject masterObjectInfoType
	for {
		if _, ok := temp.(*golang.TypeSpecContext); ok {
			masterObject.ObjectName = temp.GetChild(0).(antlr.ParseTree).GetText()
			return masterObject
		}
		temp = temp.GetParent()
		if temp == nil {
			return masterObjectInfoType{}
		}
	}
}

func (s *GoTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
}

func (s *GoTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (s *GoTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (s *GoTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (s *GoTreeShapeListener) EnterSlice(c *golang.Slice_Context) {}

func (s *GoTreeShapeListener) ExitSlice(c *golang.Slice_Context) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *GoTreeShapeListener) EnterSourceFile(ctx *golang.SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *GoTreeShapeListener) ExitSourceFile(ctx *golang.SourceFileContext) {}
// EnterPackageClause is called when production packageClause is entered.
func (s *GoTreeShapeListener) EnterPackageClause(ctx *golang.PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *GoTreeShapeListener) ExitPackageClause(ctx *golang.PackageClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *GoTreeShapeListener) EnterImportDecl(ctx *golang.ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *GoTreeShapeListener) ExitImportDecl(ctx *golang.ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *GoTreeShapeListener) EnterImportSpec(ctx *golang.ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *GoTreeShapeListener) ExitImportSpec(ctx *golang.ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *GoTreeShapeListener) EnterImportPath(ctx *golang.ImportPathContext) {
}

// ExitImportPath is called when production importPath is exited.
func (s *GoTreeShapeListener) ExitImportPath(ctx *golang.ImportPathContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *GoTreeShapeListener) EnterDeclaration(ctx *golang.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *GoTreeShapeListener) ExitDeclaration(ctx *golang.DeclarationContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *GoTreeShapeListener) EnterConstDecl(ctx *golang.ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *GoTreeShapeListener) ExitConstDecl(ctx *golang.ConstDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *GoTreeShapeListener) EnterConstSpec(ctx *golang.ConstSpecContext) {
}
func (s *GoTreeShapeListener) ExitConstSpec(ctx *golang.ConstSpecContext) {}

// EnterIdentifierList is called when production identifierList is entered.
//解析的是a = b中的a 或 a int中的a
func (s *GoTreeShapeListener) EnterIdentifierList(ctx *golang.IdentifierListContext) {
}

// ExitIdentifierList is called when production identifierList is exited.
func (s *GoTreeShapeListener) ExitIdentifierList(ctx *golang.IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
///解析的是a = b中的b
func (s *GoTreeShapeListener) EnterExpressionList(ctx *golang.ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *GoTreeShapeListener) ExitExpressionList(ctx *golang.ExpressionListContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *GoTreeShapeListener) EnterTypeDecl(ctx *golang.TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *GoTreeShapeListener) ExitTypeDecl(ctx *golang.TypeDeclContext) {}

func (s *GoTreeShapeListener) EnterTypeSpec(ctx *golang.TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *GoTreeShapeListener) ExitTypeSpec(ctx *golang.TypeSpecContext) {}

func (s *GoTreeShapeListener) EnterFunctionDecl(ctx *golang.FunctionDeclContext) {}
// ExitMethodDecl is called when production methodDecl is exited.
func (s *GoTreeShapeListener) EnterMethodDecl(ctx *golang.MethodDeclContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *GoTreeShapeListener) EnterReceiver(ctx *golang.ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *GoTreeShapeListener) ExitReceiver(ctx *golang.ReceiverContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *GoTreeShapeListener) EnterVarDecl(ctx *golang.VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *GoTreeShapeListener) ExitVarDecl(ctx *golang.VarDeclContext) {}

// EnterVarSpec is called when production varSpec is entered.
func (s *GoTreeShapeListener) EnterVarSpec(ctx *golang.VarSpecContext) {}

// ExitVarSpec is called when production varSpec is exited.
func (s *GoTreeShapeListener) ExitVarSpec(ctx *golang.VarSpecContext) {}

// EnterBlock is called when production block is entered.
func (s *GoTreeShapeListener) EnterBlock(ctx *golang.BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *GoTreeShapeListener) ExitBlock(ctx *golang.BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *GoTreeShapeListener) EnterStatementList(ctx *golang.StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *GoTreeShapeListener) ExitStatementList(ctx *golang.StatementListContext) {}

// EnterStatement is called when production statement is entered.

//对函数调用，赋值语句，声明语句有效
func (s *GoTreeShapeListener) EnterStatement(ctx *golang.StatementContext) {
	//fmt.Printf("EnterStatement:%s \n",ctx.GetText())
}

// ExitStatement is called when production statement is exited.
func (s *GoTreeShapeListener) ExitStatement(ctx *golang.StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.

//只对函数调用，赋值语句有效，声明语句无效
func (s *GoTreeShapeListener) EnterSimpleStmt(ctx *golang.SimpleStmtContext) {
	//fmt.Printf("EnterSimpleStmt:%s \n",ctx.GetText())
}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *GoTreeShapeListener) ExitSimpleStmt(ctx *golang.SimpleStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *GoTreeShapeListener) ExitExpressionStmt(ctx *golang.ExpressionStmtContext) {}

// EnterSendStmt is called when production sendStmt is entered.
func (s *GoTreeShapeListener) EnterSendStmt(ctx *golang.SendStmtContext) {}

// ExitSendStmt is called when production sendStmt is exited.
func (s *GoTreeShapeListener) ExitSendStmt(ctx *golang.SendStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *GoTreeShapeListener) EnterIncDecStmt(ctx *golang.IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *GoTreeShapeListener) ExitIncDecStmt(ctx *golang.IncDecStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *GoTreeShapeListener) EnterAssignment(ctx *golang.AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *GoTreeShapeListener) ExitAssignment(ctx *golang.AssignmentContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *GoTreeShapeListener) EnterAssign_op(ctx *golang.Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *GoTreeShapeListener) ExitAssign_op(ctx *golang.Assign_opContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *GoTreeShapeListener) EnterShortVarDecl(ctx *golang.ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *GoTreeShapeListener) ExitShortVarDecl(ctx *golang.ShortVarDeclContext) {}

// EnterEmptyStmt is called when production emptyStmt is entered.
func (s *GoTreeShapeListener) EnterEmptyStmt(ctx *golang.EmptyStmtContext) {}

// ExitEmptyStmt is called when production emptyStmt is exited.
func (s *GoTreeShapeListener) ExitEmptyStmt(ctx *golang.EmptyStmtContext) {}

// EnterLabeledStmt is called when production labeledStmt is entered.
func (s *GoTreeShapeListener) EnterLabeledStmt(ctx *golang.LabeledStmtContext) {}

// ExitLabeledStmt is called when production labeledStmt is exited.
func (s *GoTreeShapeListener) ExitLabeledStmt(ctx *golang.LabeledStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *GoTreeShapeListener) EnterReturnStmt(ctx *golang.ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *GoTreeShapeListener) ExitReturnStmt(ctx *golang.ReturnStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *GoTreeShapeListener) EnterBreakStmt(ctx *golang.BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *GoTreeShapeListener) ExitBreakStmt(ctx *golang.BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *GoTreeShapeListener) EnterContinueStmt(ctx *golang.ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *GoTreeShapeListener) ExitContinueStmt(ctx *golang.ContinueStmtContext) {}

// EnterGotoStmt is called when production gotoStmt is entered.
func (s *GoTreeShapeListener) EnterGotoStmt(ctx *golang.GotoStmtContext) {}

// ExitGotoStmt is called when production gotoStmt is exited.
func (s *GoTreeShapeListener) ExitGotoStmt(ctx *golang.GotoStmtContext) {}

// EnterFallthroughStmt is called when production fallthroughStmt is entered.
func (s *GoTreeShapeListener) EnterFallthroughStmt(ctx *golang.FallthroughStmtContext) {}

// ExitFallthroughStmt is called when production fallthroughStmt is exited.
func (s *GoTreeShapeListener) ExitFallthroughStmt(ctx *golang.FallthroughStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *GoTreeShapeListener) EnterDeferStmt(ctx *golang.DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *GoTreeShapeListener) ExitDeferStmt(ctx *golang.DeferStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *GoTreeShapeListener) EnterIfStmt(ctx *golang.IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *GoTreeShapeListener) ExitIfStmt(ctx *golang.IfStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *GoTreeShapeListener) EnterSwitchStmt(ctx *golang.SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *GoTreeShapeListener) ExitSwitchStmt(ctx *golang.SwitchStmtContext) {}

// EnterExprSwitchStmt is called when production exprSwitchStmt is entered.
func (s *GoTreeShapeListener) EnterExprSwitchStmt(ctx *golang.ExprSwitchStmtContext) {}

// ExitExprSwitchStmt is called when production exprSwitchStmt is exited.
func (s *GoTreeShapeListener) ExitExprSwitchStmt(ctx *golang.ExprSwitchStmtContext) {}

// EnterExprCaseClause is called when production exprCaseClause is entered.
func (s *GoTreeShapeListener) EnterExprCaseClause(ctx *golang.ExprCaseClauseContext) {}

// ExitExprCaseClause is called when production exprCaseClause is exited.
func (s *GoTreeShapeListener) ExitExprCaseClause(ctx *golang.ExprCaseClauseContext) {}

// EnterExprSwitchCase is called when production exprSwitchCase is entered.
func (s *GoTreeShapeListener) EnterExprSwitchCase(ctx *golang.ExprSwitchCaseContext) {}

// ExitExprSwitchCase is called when production exprSwitchCase is exited.
func (s *GoTreeShapeListener) ExitExprSwitchCase(ctx *golang.ExprSwitchCaseContext) {}

// EnterTypeSwitchStmt is called when production typeSwitchStmt is entered.
func (s *GoTreeShapeListener) EnterTypeSwitchStmt(ctx *golang.TypeSwitchStmtContext) {}

// ExitTypeSwitchStmt is called when production typeSwitchStmt is exited.
func (s *GoTreeShapeListener) ExitTypeSwitchStmt(ctx *golang.TypeSwitchStmtContext) {}

// EnterTypeSwitchGuard is called when production typeSwitchGuard is entered.
func (s *GoTreeShapeListener) EnterTypeSwitchGuard(ctx *golang.TypeSwitchGuardContext) {}

// ExitTypeSwitchGuard is called when production typeSwitchGuard is exited.
func (s *GoTreeShapeListener) ExitTypeSwitchGuard(ctx *golang.TypeSwitchGuardContext) {}

// EnterTypeCaseClause is called when production typeCaseClause is entered.
func (s *GoTreeShapeListener) EnterTypeCaseClause(ctx *golang.TypeCaseClauseContext) {}

// ExitTypeCaseClause is called when production typeCaseClause is exited.
func (s *GoTreeShapeListener) ExitTypeCaseClause(ctx *golang.TypeCaseClauseContext) {}

// EnterTypeSwitchCase is called when production typeSwitchCase is entered.
func (s *GoTreeShapeListener) EnterTypeSwitchCase(ctx *golang.TypeSwitchCaseContext) {}

// ExitTypeSwitchCase is called when production typeSwitchCase is exited.
func (s *GoTreeShapeListener) ExitTypeSwitchCase(ctx *golang.TypeSwitchCaseContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *GoTreeShapeListener) EnterTypeList(ctx *golang.TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *GoTreeShapeListener) ExitTypeList(ctx *golang.TypeListContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *GoTreeShapeListener) EnterSelectStmt(ctx *golang.SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *GoTreeShapeListener) ExitSelectStmt(ctx *golang.SelectStmtContext) {}

// EnterCommClause is called when production commClause is entered.
func (s *GoTreeShapeListener) EnterCommClause(ctx *golang.CommClauseContext) {}

// ExitCommClause is called when production commClause is exited.
func (s *GoTreeShapeListener) ExitCommClause(ctx *golang.CommClauseContext) {}

// EnterCommCase is called when production commCase is entered.
func (s *GoTreeShapeListener) EnterCommCase(ctx *golang.CommCaseContext) {}

// ExitCommCase is called when production commCase is exited.
func (s *GoTreeShapeListener) ExitCommCase(ctx *golang.CommCaseContext) {}

// EnterRecvStmt is called when production recvStmt is entered.
func (s *GoTreeShapeListener) EnterRecvStmt(ctx *golang.RecvStmtContext) {}

// ExitRecvStmt is called when production recvStmt is exited.
func (s *GoTreeShapeListener) ExitRecvStmt(ctx *golang.RecvStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *GoTreeShapeListener) EnterForStmt(ctx *golang.ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *GoTreeShapeListener) ExitForStmt(ctx *golang.ForStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *GoTreeShapeListener) EnterForClause(ctx *golang.ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *GoTreeShapeListener) ExitForClause(ctx *golang.ForClauseContext) {}

// EnterRangeClause is called when production rangeClause is entered.
func (s *GoTreeShapeListener) EnterRangeClause(ctx *golang.RangeClauseContext) {}

// ExitRangeClause is called when production rangeClause is exited.
func (s *GoTreeShapeListener) ExitRangeClause(ctx *golang.RangeClauseContext) {}

// EnterGoStmt is called when production goStmt is entered.
func (s *GoTreeShapeListener) EnterGoStmt(ctx *golang.GoStmtContext) {}

// ExitGoStmt is called when production goStmt is exited.
func (s *GoTreeShapeListener) ExitGoStmt(ctx *golang.GoStmtContext) {}
func (s *GoTreeShapeListener) EnterType_(ctx *golang.Type_Context) {
}

// ExitType_ is called when production type_ is exited.
func (s *GoTreeShapeListener) ExitType_(ctx *golang.Type_Context) {}

// EnterTypeName is called when production typeName is entered.
func (s *GoTreeShapeListener) EnterTypeName(ctx *golang.TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *GoTreeShapeListener) ExitTypeName(ctx *golang.TypeNameContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *GoTreeShapeListener) EnterTypeLit(ctx *golang.TypeLitContext) {
}

// ExitTypeLit is called when production typeLit is exited.
func (s *GoTreeShapeListener) ExitTypeLit(ctx *golang.TypeLitContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *GoTreeShapeListener) EnterArrayType(ctx *golang.ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *GoTreeShapeListener) ExitArrayType(ctx *golang.ArrayTypeContext) {}

// EnterArrayLength is called when production arrayLength is entered.
func (s *GoTreeShapeListener) EnterArrayLength(ctx *golang.ArrayLengthContext) {}

// ExitArrayLength is called when production arrayLength is exited.
func (s *GoTreeShapeListener) ExitArrayLength(ctx *golang.ArrayLengthContext) {}

// EnterElementType is called when production elementType is entered.
func (s *GoTreeShapeListener) EnterElementType(ctx *golang.ElementTypeContext) {
}

// ExitElementType is called when production elementType is exited.
func (s *GoTreeShapeListener) ExitElementType(ctx *golang.ElementTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *GoTreeShapeListener) EnterPointerType(ctx *golang.PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *GoTreeShapeListener) ExitPointerType(ctx *golang.PointerTypeContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *GoTreeShapeListener) EnterInterfaceType(ctx *golang.InterfaceTypeContext) {
}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *GoTreeShapeListener) ExitInterfaceType(ctx *golang.InterfaceTypeContext) {
}

// EnterSliceType is called when production sliceType is entered.
func (s *GoTreeShapeListener) EnterSliceType(ctx *golang.SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *GoTreeShapeListener) ExitSliceType(ctx *golang.SliceTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *GoTreeShapeListener) EnterMapType(ctx *golang.MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *GoTreeShapeListener) ExitMapType(ctx *golang.MapTypeContext) {}

// EnterChannelType is called when production channelType is entered.
func (s *GoTreeShapeListener) EnterChannelType(ctx *golang.ChannelTypeContext) {}

// ExitChannelType is called when production channelType is exited.
func (s *GoTreeShapeListener) ExitChannelType(ctx *golang.ChannelTypeContext) {}

func (s *GoTreeShapeListener) EnterMethodSpec(ctx *golang.MethodSpecContext) {
}

// ExitMethodSpec is called when production methodSpec is exited.
func (s *GoTreeShapeListener) ExitMethodSpec(ctx *golang.MethodSpecContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *GoTreeShapeListener) EnterFunctionType(ctx *golang.FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *GoTreeShapeListener) ExitFunctionType(ctx *golang.FunctionTypeContext) {}

// EnterSignature is called when production signature is entered.
func (s *GoTreeShapeListener) EnterSignature(ctx *golang.SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *GoTreeShapeListener) ExitSignature(ctx *golang.SignatureContext) {}

// EnterResult is called when production result is entered.
func (s *GoTreeShapeListener) EnterResult(ctx *golang.ResultContext) {}

// ExitResult is called when production result is exited.
func (s *GoTreeShapeListener) ExitResult(ctx *golang.ResultContext) {}

// EnterParameters is called when production parameters is entered.
func (s *GoTreeShapeListener) EnterParameters(ctx *golang.ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *GoTreeShapeListener) ExitParameters(ctx *golang.ParametersContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *GoTreeShapeListener) EnterParameterDecl(ctx *golang.ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *GoTreeShapeListener) ExitParameterDecl(ctx *golang.ParameterDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *GoTreeShapeListener) EnterExpression(ctx *golang.ExpressionContext) {
}

// ExitExpression is called when production expression is exited.
func (s *GoTreeShapeListener) ExitExpression(ctx *golang.ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *GoTreeShapeListener) EnterPrimaryExpr(ctx *golang.PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *GoTreeShapeListener) ExitPrimaryExpr(ctx *golang.PrimaryExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *GoTreeShapeListener) EnterUnaryExpr(ctx *golang.UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *GoTreeShapeListener) ExitUnaryExpr(ctx *golang.UnaryExprContext) {}

// EnterConversion is called when production conversion is entered.
func (s *GoTreeShapeListener) EnterConversion(ctx *golang.ConversionContext) {}

// ExitConversion is called when production conversion is exited.
func (s *GoTreeShapeListener) ExitConversion(ctx *golang.ConversionContext) {}

// EnterOperand is called when production operand is entered.
func (s *GoTreeShapeListener) EnterOperand(ctx *golang.OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *GoTreeShapeListener) ExitOperand(ctx *golang.OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *GoTreeShapeListener) EnterLiteral(ctx *golang.LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *GoTreeShapeListener) ExitLiteral(ctx *golang.LiteralContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *GoTreeShapeListener) EnterBasicLit(ctx *golang.BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *GoTreeShapeListener) ExitBasicLit(ctx *golang.BasicLitContext) {}

// EnterInteger is called when production integer is entered.
func (s *GoTreeShapeListener) EnterInteger(ctx *golang.IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *GoTreeShapeListener) ExitInteger(ctx *golang.IntegerContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *GoTreeShapeListener) EnterOperandName(ctx *golang.OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *GoTreeShapeListener) ExitOperandName(ctx *golang.OperandNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *GoTreeShapeListener) EnterQualifiedIdent(ctx *golang.QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *GoTreeShapeListener) ExitQualifiedIdent(ctx *golang.QualifiedIdentContext) {}

// EnterCompositeLit is called when production compositeLit is entered.
func (s *GoTreeShapeListener) EnterCompositeLit(ctx *golang.CompositeLitContext) {}

// ExitCompositeLit is called when production compositeLit is exited.
func (s *GoTreeShapeListener) ExitCompositeLit(ctx *golang.CompositeLitContext) {}

// EnterLiteralType is called when production literalType is entered.
func (s *GoTreeShapeListener) EnterLiteralType(ctx *golang.LiteralTypeContext) {
}

// ExitLiteralType is called when production literalType is exited.
func (s *GoTreeShapeListener) ExitLiteralType(ctx *golang.LiteralTypeContext) {}

// EnterLiteralValue is called when production literalValue is entered.

func (s *GoTreeShapeListener) EnterLiteralValue(ctx *golang.LiteralValueContext) {
}

// ExitLiteralValue is called when production literalValue is exited.
func (s *GoTreeShapeListener) ExitLiteralValue(ctx *golang.LiteralValueContext) {}

// EnterElementList is called when production elementList is entered.
func (s *GoTreeShapeListener) EnterElementList(ctx *golang.ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *GoTreeShapeListener) ExitElementList(ctx *golang.ElementListContext) {}

// EnterKeyedElement is called when production keyedElement is entered.
func (s *GoTreeShapeListener) EnterKeyedElement(ctx *golang.KeyedElementContext) {}

// ExitKeyedElement is called when production keyedElement is exited.
func (s *GoTreeShapeListener) ExitKeyedElement(ctx *golang.KeyedElementContext) {}

// EnterKey is called when production key is entered.
func (s *GoTreeShapeListener) EnterKey(ctx *golang.KeyContext) {}

// ExitKey is called when production key is exited.
func (s *GoTreeShapeListener) ExitKey(ctx *golang.KeyContext) {}

// EnterElement is called when production element is entered.
func (s *GoTreeShapeListener) EnterElement(ctx *golang.ElementContext) {}

// ExitElement is called when production element is exited.
func (s *GoTreeShapeListener) ExitElement(ctx *golang.ElementContext) {}

func (s *GoTreeShapeListener) ExitStructType(ctx *golang.StructTypeContext) {
}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *GoTreeShapeListener) EnterFieldDecl(ctx *golang.FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *GoTreeShapeListener) ExitFieldDecl(ctx *golang.FieldDeclContext) {}

// EnterString_ is called when production string_ is entered.
func (s *GoTreeShapeListener) EnterString_(ctx *golang.String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *GoTreeShapeListener) ExitString_(ctx *golang.String_Context) {}

// EnterEmbeddedField is called when production embeddedField is entered.
func (s *GoTreeShapeListener) EnterEmbeddedField(ctx *golang.EmbeddedFieldContext) {}

// ExitEmbeddedField is called when production embeddedField is exited.
func (s *GoTreeShapeListener) ExitEmbeddedField(ctx *golang.EmbeddedFieldContext) {}

// EnterFunctionLit is called when production functionLit is entered.
func (s *GoTreeShapeListener) EnterFunctionLit(ctx *golang.FunctionLitContext) {
}

// ExitFunctionLit is called when production functionLit is exited.
func (s *GoTreeShapeListener) ExitFunctionLit(ctx *golang.FunctionLitContext) {}

// EnterIndex is called when production index is entered.
func (s *GoTreeShapeListener) EnterIndex(ctx *golang.IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *GoTreeShapeListener) ExitIndex(ctx *golang.IndexContext) {}

// EnterSlice_ is called when production slice_ is entered.
func (s *GoTreeShapeListener) EnterSlice_(ctx *golang.Slice_Context) {}

// ExitSlice_ is called when production slice_ is exited.
func (s *GoTreeShapeListener) ExitSlice_(ctx *golang.Slice_Context) {}

// EnterTypeAssertion is called when production typeAssertion is entered.
func (s *GoTreeShapeListener) EnterTypeAssertion(ctx *golang.TypeAssertionContext) {
}

// ExitTypeAssertion is called when production typeAssertion is exited.
func (s *GoTreeShapeListener) ExitTypeAssertion(ctx *golang.TypeAssertionContext) {}

// EnterArguments is called when production arguments is entered.
func (s *GoTreeShapeListener) EnterArguments(ctx *golang.ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *GoTreeShapeListener) ExitArguments(ctx *golang.ArgumentsContext) {}

// EnterMethodExpr is called when production methodExpr is entered.
func (s *GoTreeShapeListener) EnterMethodExpr(ctx *golang.MethodExprContext) {}

// ExitMethodExpr is called when production methodExpr is exited.
func (s *GoTreeShapeListener) ExitMethodExpr(ctx *golang.MethodExprContext) {}

// EnterReceiverType is called when production receiverType is entered.
func (s *GoTreeShapeListener) EnterReceiverType(ctx *golang.ReceiverTypeContext) {}

// ExitReceiverType is called when production receiverType is exited.
func (s *GoTreeShapeListener) ExitReceiverType(ctx *golang.ReceiverTypeContext) {}

// EnterEos is called when production eos is entered.
func (s *GoTreeShapeListener) EnterEos(ctx *golang.EosContext) {}

// ExitEos is called when production eos is exited.
func (s *GoTreeShapeListener) ExitEos(ctx *golang.EosContext) {}


