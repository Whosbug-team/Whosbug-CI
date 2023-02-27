package antlr

import (
	"sync"

	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr/_go"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// GoAstParser implement AstParser for go language
//
//	@author kevineluo
//	@update 2023-02-28 01:06:15
type GoAstParser struct {
	astInfo AstInfo
}

var _ _go.GoParserListener = &GoAstParser{}

var (
	goLexerPool = &sync.Pool{New: func() any {
		return _go.NewGoLexer(nil)
	}}
	goParserPool = &sync.Pool{New: func() any {
		return _go.NewGoParser(nil)
	}}
	newGoAstParserPool = &sync.Pool{New: func() any {
		return new(GoAstParser)
	}}
)

// AstParse main parse process for go language
//
//	@receiver s *GoAstParser
//	@param diffText string
//	@return AstInfo
//	@author kevineluo
//	@update 2023-02-28 01:07:29
func (s *GoAstParser) AstParse(input string) AstInfo {
	//	截取目标文本的输入流
	inputStream := antlr.NewInputStream(input)
	//	初始化lexer
	lexer := goLexerPool.Get().(*_go.GoLexer)
	defer goLexerPool.Put(lexer)
	lexer.SetInputStream(inputStream)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := goParserPool.Get().(*_go.GoParser)
	defer goParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	p.RemoveErrorListeners()
	//	解析模式->每个编译单位
	tree := p.SourceFile()
	//	创建listener
	listener := newGoAstParserPool.Get().(*GoAstParser)
	defer newGoAstParserPool.Put(listener)
	// 初始化置空
	listener.astInfo = AstInfo{}
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.astInfo
}

// EnterFunctionDecl 解析常规函数定义
//
//	@receiver s
//	@param ctx
func (s *GoAstParser) EnterFunctionDecl(ctx *_go.FunctionDeclContext) {

	methodInfo := Method{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		Name:       ctx.IDENTIFIER().GetText(),
		Parameters: getFunctionAndMethodParams(ctx),
	}
	s.astInfo.Methods = append(s.astInfo.Methods, methodInfo)
}

func getFunctionAndMethodParams(ctx antlr.ParseTree) (params string) {
	if _, ok := ctx.(*_go.FunctionDeclContext); ok {
		if ctx.(*_go.FunctionDeclContext).Signature().(*_go.SignatureContext).Parameters() != nil {
			for index, item := range ctx.(*_go.FunctionDeclContext).Signature().(*_go.SignatureContext).Parameters().(*_go.ParametersContext).AllParameterDecl() {
				params += item.(*_go.ParameterDeclContext).IdentifierList().GetText() + " "
				params += item.(*_go.ParameterDeclContext).Type_().GetText()
				if index != len(ctx.(*_go.FunctionDeclContext).Signature().(*_go.SignatureContext).Parameters().(*_go.ParametersContext).AllParameterDecl())-1 {
					params += ", "
				}
			}
		}
	}
	if _, ok := ctx.(*_go.MethodDeclContext); ok {
		if ctx.(*_go.MethodDeclContext).Signature().(*_go.SignatureContext).Parameters() != nil {
			for index, item := range ctx.(*_go.MethodDeclContext).Signature().(*_go.SignatureContext).Parameters().(*_go.ParametersContext).AllParameterDecl() {
				params += item.(*_go.ParameterDeclContext).IdentifierList().GetText() + " "
				params += item.(*_go.ParameterDeclContext).Type_().GetText()
				if index != len(ctx.(*_go.MethodDeclContext).Signature().(*_go.SignatureContext).Parameters().(*_go.ParametersContext).AllParameterDecl())-1 {
					params += ", "
				}
			}
		}
	}
	return
}

// EnterMethodDecl ExitMethodDecl is called when production methodDecl is exited.
//
//	@receiver s
//	@param ctx
func (s *GoAstParser) EnterMethodDecl(ctx *_go.MethodDeclContext) {
	var methodInfo Method
	if temp := getRecvrTypes(ctx); len(temp) > 0 {
		methodInfo = Method{
			StartLine:  ctx.GetStart().GetLine(),
			EndLine:    ctx.GetStop().GetLine(),
			Name:       temp[0] + "." + ctx.IDENTIFIER().GetText(),
			Parameters: getFunctionAndMethodParams(ctx),
		}
		s.astInfo.Methods = append(s.astInfo.Methods, methodInfo)
	}
}

// getIdentAndType 获取方法实现的结构体类型
//
//	@param ctx
//	@return identifiers
//	@return types
func getRecvrTypes(ctx *_go.MethodDeclContext) (types []string) {
	temp := ctx.Receiver().(*_go.ReceiverContext).Parameters().(*_go.ParametersContext).AllParameterDecl()
	for _, item := range temp {
		types = append(types, item.(*_go.ParameterDeclContext).Type_().GetText())
	}
	return
}

// EnterTypeDecl 主要用于解析结构体
//
//	@receiver s
//	@param ctx
func (s *GoAstParser) EnterTypeDecl(ctx *_go.TypeDeclContext) {
	var structInfo = Class{
		Name:      ctx.TypeSpec(0).(*_go.TypeSpecContext).IDENTIFIER().GetText(),
		StartLine: ctx.GetStart().GetLine(),
		EndLine:   ctx.GetStop().GetLine(),
	}
	s.astInfo.Classes = append(s.astInfo.Classes, structInfo)
}

// ? GoParserStdMethods
func (s *GoAstParser) EnterStructType(ctx *_go.StructTypeContext) {}

func (s *GoAstParser) ExitFunctionDecl(ctx *_go.FunctionDeclContext) {}

func (s *GoAstParser) ExitMethodDecl(ctx *_go.MethodDeclContext) {
}

// EnterVarSpec is called when production varSpec is entered.
func (s *GoAstParser) EnterVarSpec(ctx *_go.VarSpecContext) {
}

// EnterImportDecl is called when production importDecl is entered.
func (s *GoAstParser) EnterImportDecl(ctx *_go.ImportDeclContext) {
}

// EnterCompositeLit is called when production compositeLit is entered.
//
//	cube1 := Cube{}
//	var cube2 = Cube{}
func (s *GoAstParser) EnterCompositeLit(ctx *_go.CompositeLitContext) {

}

func (s *GoAstParser) EnterExpressionStmt(ctx *_go.ExpressionStmtContext) {
}

func (s *GoAstParser) VisitTerminal(node antlr.TerminalNode) {
}

func (s *GoAstParser) VisitErrorNode(node antlr.ErrorNode) {
}

func (s *GoAstParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (s *GoAstParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

func (s *GoAstParser) EnterSlice(c *_go.Slice_Context) {}

func (s *GoAstParser) ExitSlice(c *_go.Slice_Context) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *GoAstParser) EnterSourceFile(ctx *_go.SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *GoAstParser) ExitSourceFile(ctx *_go.SourceFileContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *GoAstParser) EnterPackageClause(ctx *_go.PackageClauseContext) {
}

// ExitPackageClause is called when production packageClause is exited.
func (s *GoAstParser) ExitPackageClause(ctx *_go.PackageClauseContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *GoAstParser) ExitImportDecl(ctx *_go.ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *GoAstParser) EnterImportSpec(ctx *_go.ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *GoAstParser) ExitImportSpec(ctx *_go.ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *GoAstParser) EnterImportPath(ctx *_go.ImportPathContext) {
}

// ExitImportPath is called when production importPath is exited.
func (s *GoAstParser) ExitImportPath(ctx *_go.ImportPathContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *GoAstParser) EnterDeclaration(ctx *_go.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *GoAstParser) ExitDeclaration(ctx *_go.DeclarationContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *GoAstParser) EnterConstDecl(ctx *_go.ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *GoAstParser) ExitConstDecl(ctx *_go.ConstDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *GoAstParser) EnterConstSpec(ctx *_go.ConstSpecContext) {
}
func (s *GoAstParser) ExitConstSpec(ctx *_go.ConstSpecContext) {}

// EnterIdentifierList is called when production identifierList is entered.
// 解析的是a = b中的a 或 a int中的a
func (s *GoAstParser) EnterIdentifierList(ctx *_go.IdentifierListContext) {
}

// ExitIdentifierList is called when production identifierList is exited.
func (s *GoAstParser) ExitIdentifierList(ctx *_go.IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
// /解析的是a = b中的b
func (s *GoAstParser) EnterExpressionList(ctx *_go.ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *GoAstParser) ExitExpressionList(ctx *_go.ExpressionListContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *GoAstParser) ExitTypeDecl(ctx *_go.TypeDeclContext) {}

func (s *GoAstParser) EnterTypeSpec(ctx *_go.TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *GoAstParser) ExitTypeSpec(ctx *_go.TypeSpecContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *GoAstParser) EnterReceiver(ctx *_go.ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *GoAstParser) ExitReceiver(ctx *_go.ReceiverContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *GoAstParser) EnterVarDecl(ctx *_go.VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *GoAstParser) ExitVarDecl(ctx *_go.VarDeclContext) {}

// ExitVarSpec is called when production varSpec is exited.
func (s *GoAstParser) ExitVarSpec(ctx *_go.VarSpecContext) {}

// EnterBlock is called when production block is entered.
func (s *GoAstParser) EnterBlock(ctx *_go.BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *GoAstParser) ExitBlock(ctx *_go.BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *GoAstParser) EnterStatementList(ctx *_go.StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *GoAstParser) ExitStatementList(ctx *_go.StatementListContext) {}

// EnterStatement is called when production statement is entered.

// 对函数调用，赋值语句，声明语句有效
func (s *GoAstParser) EnterStatement(ctx *_go.StatementContext) {
	//fmt.Printf("EnterStatement:%s \n",ctx.GetText())
}

// ExitStatement is called when production statement is exited.
func (s *GoAstParser) ExitStatement(ctx *_go.StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.

// 只对函数调用，赋值语句有效，声明语句无效
func (s *GoAstParser) EnterSimpleStmt(ctx *_go.SimpleStmtContext) {
	//fmt.Printf("EnterSimpleStmt:%s \n",ctx.GetText())
}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *GoAstParser) ExitSimpleStmt(ctx *_go.SimpleStmtContext) {}

// ExitExpressionStmt is called when production expressionStmt is exited.
func (s *GoAstParser) ExitExpressionStmt(ctx *_go.ExpressionStmtContext) {}

// EnterSendStmt is called when production sendStmt is entered.
func (s *GoAstParser) EnterSendStmt(ctx *_go.SendStmtContext) {}

// ExitSendStmt is called when production sendStmt is exited.
func (s *GoAstParser) ExitSendStmt(ctx *_go.SendStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *GoAstParser) EnterIncDecStmt(ctx *_go.IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *GoAstParser) ExitIncDecStmt(ctx *_go.IncDecStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *GoAstParser) EnterAssignment(ctx *_go.AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *GoAstParser) ExitAssignment(ctx *_go.AssignmentContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *GoAstParser) EnterAssign_op(ctx *_go.Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *GoAstParser) ExitAssign_op(ctx *_go.Assign_opContext) {}

// EnterShortVarDecl is called when production shortVarDecl is entered.
func (s *GoAstParser) EnterShortVarDecl(ctx *_go.ShortVarDeclContext) {}

// ExitShortVarDecl is called when production shortVarDecl is exited.
func (s *GoAstParser) ExitShortVarDecl(ctx *_go.ShortVarDeclContext) {}

// EnterEmptyStmt is called when production emptyStmt is entered.
func (s *GoAstParser) EnterEmptyStmt(ctx *_go.EmptyStmtContext) {}

// ExitEmptyStmt is called when production emptyStmt is exited.
func (s *GoAstParser) ExitEmptyStmt(ctx *_go.EmptyStmtContext) {}

// EnterLabeledStmt is called when production labeledStmt is entered.
func (s *GoAstParser) EnterLabeledStmt(ctx *_go.LabeledStmtContext) {}

// ExitLabeledStmt is called when production labeledStmt is exited.
func (s *GoAstParser) ExitLabeledStmt(ctx *_go.LabeledStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *GoAstParser) EnterReturnStmt(ctx *_go.ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *GoAstParser) ExitReturnStmt(ctx *_go.ReturnStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *GoAstParser) EnterBreakStmt(ctx *_go.BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *GoAstParser) ExitBreakStmt(ctx *_go.BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *GoAstParser) EnterContinueStmt(ctx *_go.ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *GoAstParser) ExitContinueStmt(ctx *_go.ContinueStmtContext) {}

// EnterGotoStmt is called when production gotoStmt is entered.
func (s *GoAstParser) EnterGotoStmt(ctx *_go.GotoStmtContext) {}

// ExitGotoStmt is called when production gotoStmt is exited.
func (s *GoAstParser) ExitGotoStmt(ctx *_go.GotoStmtContext) {}

// EnterFallthroughStmt is called when production fallthroughStmt is entered.
func (s *GoAstParser) EnterFallthroughStmt(ctx *_go.FallthroughStmtContext) {}

// ExitFallthroughStmt is called when production fallthroughStmt is exited.
func (s *GoAstParser) ExitFallthroughStmt(ctx *_go.FallthroughStmtContext) {}

// EnterDeferStmt is called when production deferStmt is entered.
func (s *GoAstParser) EnterDeferStmt(ctx *_go.DeferStmtContext) {}

// ExitDeferStmt is called when production deferStmt is exited.
func (s *GoAstParser) ExitDeferStmt(ctx *_go.DeferStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *GoAstParser) EnterIfStmt(ctx *_go.IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *GoAstParser) ExitIfStmt(ctx *_go.IfStmtContext) {}

// EnterSwitchStmt is called when production switchStmt is entered.
func (s *GoAstParser) EnterSwitchStmt(ctx *_go.SwitchStmtContext) {}

// ExitSwitchStmt is called when production switchStmt is exited.
func (s *GoAstParser) ExitSwitchStmt(ctx *_go.SwitchStmtContext) {}

// EnterExprSwitchStmt is called when production exprSwitchStmt is entered.
func (s *GoAstParser) EnterExprSwitchStmt(ctx *_go.ExprSwitchStmtContext) {}

// ExitExprSwitchStmt is called when production exprSwitchStmt is exited.
func (s *GoAstParser) ExitExprSwitchStmt(ctx *_go.ExprSwitchStmtContext) {}

// EnterExprCaseClause is called when production exprCaseClause is entered.
func (s *GoAstParser) EnterExprCaseClause(ctx *_go.ExprCaseClauseContext) {}

// ExitExprCaseClause is called when production exprCaseClause is exited.
func (s *GoAstParser) ExitExprCaseClause(ctx *_go.ExprCaseClauseContext) {}

// EnterExprSwitchCase is called when production exprSwitchCase is entered.
func (s *GoAstParser) EnterExprSwitchCase(ctx *_go.ExprSwitchCaseContext) {}

// ExitExprSwitchCase is called when production exprSwitchCase is exited.
func (s *GoAstParser) ExitExprSwitchCase(ctx *_go.ExprSwitchCaseContext) {}

// EnterTypeSwitchStmt is called when production typeSwitchStmt is entered.
func (s *GoAstParser) EnterTypeSwitchStmt(ctx *_go.TypeSwitchStmtContext) {}

// ExitTypeSwitchStmt is called when production typeSwitchStmt is exited.
func (s *GoAstParser) ExitTypeSwitchStmt(ctx *_go.TypeSwitchStmtContext) {}

// EnterTypeSwitchGuard is called when production typeSwitchGuard is entered.
func (s *GoAstParser) EnterTypeSwitchGuard(ctx *_go.TypeSwitchGuardContext) {}

// ExitTypeSwitchGuard is called when production typeSwitchGuard is exited.
func (s *GoAstParser) ExitTypeSwitchGuard(ctx *_go.TypeSwitchGuardContext) {}

// EnterTypeCaseClause is called when production typeCaseClause is entered.
func (s *GoAstParser) EnterTypeCaseClause(ctx *_go.TypeCaseClauseContext) {}

// ExitTypeCaseClause is called when production typeCaseClause is exited.
func (s *GoAstParser) ExitTypeCaseClause(ctx *_go.TypeCaseClauseContext) {}

// EnterTypeSwitchCase is called when production typeSwitchCase is entered.
func (s *GoAstParser) EnterTypeSwitchCase(ctx *_go.TypeSwitchCaseContext) {}

// ExitTypeSwitchCase is called when production typeSwitchCase is exited.
func (s *GoAstParser) ExitTypeSwitchCase(ctx *_go.TypeSwitchCaseContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *GoAstParser) EnterTypeList(ctx *_go.TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *GoAstParser) ExitTypeList(ctx *_go.TypeListContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *GoAstParser) EnterSelectStmt(ctx *_go.SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *GoAstParser) ExitSelectStmt(ctx *_go.SelectStmtContext) {}

// EnterCommClause is called when production commClause is entered.
func (s *GoAstParser) EnterCommClause(ctx *_go.CommClauseContext) {}

// ExitCommClause is called when production commClause is exited.
func (s *GoAstParser) ExitCommClause(ctx *_go.CommClauseContext) {}

// EnterCommCase is called when production commCase is entered.
func (s *GoAstParser) EnterCommCase(ctx *_go.CommCaseContext) {}

// ExitCommCase is called when production commCase is exited.
func (s *GoAstParser) ExitCommCase(ctx *_go.CommCaseContext) {}

// EnterRecvStmt is called when production recvStmt is entered.
func (s *GoAstParser) EnterRecvStmt(ctx *_go.RecvStmtContext) {}

// ExitRecvStmt is called when production recvStmt is exited.
func (s *GoAstParser) ExitRecvStmt(ctx *_go.RecvStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *GoAstParser) EnterForStmt(ctx *_go.ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *GoAstParser) ExitForStmt(ctx *_go.ForStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *GoAstParser) EnterForClause(ctx *_go.ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *GoAstParser) ExitForClause(ctx *_go.ForClauseContext) {}

// EnterRangeClause is called when production rangeClause is entered.
func (s *GoAstParser) EnterRangeClause(ctx *_go.RangeClauseContext) {}

// ExitRangeClause is called when production rangeClause is exited.
func (s *GoAstParser) ExitRangeClause(ctx *_go.RangeClauseContext) {}

// EnterGoStmt is called when production goStmt is entered.
func (s *GoAstParser) EnterGoStmt(ctx *_go.GoStmtContext) {}

// ExitGoStmt is called when production goStmt is exited.
func (s *GoAstParser) ExitGoStmt(ctx *_go.GoStmtContext) {}
func (s *GoAstParser) EnterType_(ctx *_go.Type_Context) {
}

// ExitType_ is called when production type_ is exited.
func (s *GoAstParser) ExitType_(ctx *_go.Type_Context) {}

// EnterTypeName is called when production typeName is entered.
func (s *GoAstParser) EnterTypeName(ctx *_go.TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *GoAstParser) ExitTypeName(ctx *_go.TypeNameContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *GoAstParser) EnterTypeLit(ctx *_go.TypeLitContext) {
}

// ExitTypeLit is called when production typeLit is exited.
func (s *GoAstParser) ExitTypeLit(ctx *_go.TypeLitContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *GoAstParser) EnterArrayType(ctx *_go.ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *GoAstParser) ExitArrayType(ctx *_go.ArrayTypeContext) {}

// EnterArrayLength is called when production arrayLength is entered.
func (s *GoAstParser) EnterArrayLength(ctx *_go.ArrayLengthContext) {}

// ExitArrayLength is called when production arrayLength is exited.
func (s *GoAstParser) ExitArrayLength(ctx *_go.ArrayLengthContext) {}

// EnterElementType is called when production elementType is entered.
func (s *GoAstParser) EnterElementType(ctx *_go.ElementTypeContext) {
}

// ExitElementType is called when production elementType is exited.
func (s *GoAstParser) ExitElementType(ctx *_go.ElementTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *GoAstParser) EnterPointerType(ctx *_go.PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *GoAstParser) ExitPointerType(ctx *_go.PointerTypeContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *GoAstParser) EnterInterfaceType(ctx *_go.InterfaceTypeContext) {
}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *GoAstParser) ExitInterfaceType(ctx *_go.InterfaceTypeContext) {
}

// EnterSliceType is called when production sliceType is entered.
func (s *GoAstParser) EnterSliceType(ctx *_go.SliceTypeContext) {}

// ExitSliceType is called when production sliceType is exited.
func (s *GoAstParser) ExitSliceType(ctx *_go.SliceTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *GoAstParser) EnterMapType(ctx *_go.MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *GoAstParser) ExitMapType(ctx *_go.MapTypeContext) {}

// EnterChannelType is called when production channelType is entered.
func (s *GoAstParser) EnterChannelType(ctx *_go.ChannelTypeContext) {}

// ExitChannelType is called when production channelType is exited.
func (s *GoAstParser) ExitChannelType(ctx *_go.ChannelTypeContext) {}

func (s *GoAstParser) EnterMethodSpec(ctx *_go.MethodSpecContext) {
}

// ExitMethodSpec is called when production methodSpec is exited.
func (s *GoAstParser) ExitMethodSpec(ctx *_go.MethodSpecContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *GoAstParser) EnterFunctionType(ctx *_go.FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *GoAstParser) ExitFunctionType(ctx *_go.FunctionTypeContext) {}

// EnterSignature is called when production signature is entered.
func (s *GoAstParser) EnterSignature(ctx *_go.SignatureContext) {}

// ExitSignature is called when production signature is exited.
func (s *GoAstParser) ExitSignature(ctx *_go.SignatureContext) {}

// EnterResult is called when production result is entered.
func (s *GoAstParser) EnterResult(ctx *_go.ResultContext) {}

// ExitResult is called when production result is exited.
func (s *GoAstParser) ExitResult(ctx *_go.ResultContext) {}

// EnterParameters is called when production parameters is entered.
func (s *GoAstParser) EnterParameters(ctx *_go.ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *GoAstParser) ExitParameters(ctx *_go.ParametersContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *GoAstParser) EnterParameterDecl(ctx *_go.ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *GoAstParser) ExitParameterDecl(ctx *_go.ParameterDeclContext) {}

// EnterExpression is called when production expression is entered.
func (s *GoAstParser) EnterExpression(ctx *_go.ExpressionContext) {
}

// ExitExpression is called when production expression is exited.
func (s *GoAstParser) ExitExpression(ctx *_go.ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *GoAstParser) EnterPrimaryExpr(ctx *_go.PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *GoAstParser) ExitPrimaryExpr(ctx *_go.PrimaryExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *GoAstParser) EnterUnaryExpr(ctx *_go.UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *GoAstParser) ExitUnaryExpr(ctx *_go.UnaryExprContext) {}

// EnterConversion is called when production conversion is entered.
func (s *GoAstParser) EnterConversion(ctx *_go.ConversionContext) {}

// ExitConversion is called when production conversion is exited.
func (s *GoAstParser) ExitConversion(ctx *_go.ConversionContext) {}

// EnterOperand is called when production operand is entered.
func (s *GoAstParser) EnterOperand(ctx *_go.OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *GoAstParser) ExitOperand(ctx *_go.OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *GoAstParser) EnterLiteral(ctx *_go.LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *GoAstParser) ExitLiteral(ctx *_go.LiteralContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *GoAstParser) EnterBasicLit(ctx *_go.BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *GoAstParser) ExitBasicLit(ctx *_go.BasicLitContext) {}

// EnterInteger is called when production integer is entered.
func (s *GoAstParser) EnterInteger(ctx *_go.IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *GoAstParser) ExitInteger(ctx *_go.IntegerContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *GoAstParser) EnterOperandName(ctx *_go.OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *GoAstParser) ExitOperandName(ctx *_go.OperandNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *GoAstParser) EnterQualifiedIdent(ctx *_go.QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *GoAstParser) ExitQualifiedIdent(ctx *_go.QualifiedIdentContext) {}

// ExitCompositeLit is called when production compositeLit is exited.
func (s *GoAstParser) ExitCompositeLit(ctx *_go.CompositeLitContext) {}

// EnterLiteralType is called when production literalType is entered.
func (s *GoAstParser) EnterLiteralType(ctx *_go.LiteralTypeContext) {
}

// ExitLiteralType is called when production literalType is exited.
func (s *GoAstParser) ExitLiteralType(ctx *_go.LiteralTypeContext) {}

// EnterLiteralValue is called when production literalValue is entered.

func (s *GoAstParser) EnterLiteralValue(ctx *_go.LiteralValueContext) {
}

// ExitLiteralValue is called when production literalValue is exited.
func (s *GoAstParser) ExitLiteralValue(ctx *_go.LiteralValueContext) {}

// EnterElementList is called when production elementList is entered.
func (s *GoAstParser) EnterElementList(ctx *_go.ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *GoAstParser) ExitElementList(ctx *_go.ElementListContext) {}

// EnterKeyedElement is called when production keyedElement is entered.
func (s *GoAstParser) EnterKeyedElement(ctx *_go.KeyedElementContext) {}

// ExitKeyedElement is called when production keyedElement is exited.
func (s *GoAstParser) ExitKeyedElement(ctx *_go.KeyedElementContext) {}

// EnterKey is called when production key is entered.
func (s *GoAstParser) EnterKey(ctx *_go.KeyContext) {}

// ExitKey is called when production key is exited.
func (s *GoAstParser) ExitKey(ctx *_go.KeyContext) {}

// EnterElement is called when production element is entered.
func (s *GoAstParser) EnterElement(ctx *_go.ElementContext) {}

// ExitElement is called when production element is exited.
func (s *GoAstParser) ExitElement(ctx *_go.ElementContext) {}

func (s *GoAstParser) ExitStructType(ctx *_go.StructTypeContext) {
}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *GoAstParser) EnterFieldDecl(ctx *_go.FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *GoAstParser) ExitFieldDecl(ctx *_go.FieldDeclContext) {}

// EnterString_ is called when production string_ is entered.
func (s *GoAstParser) EnterString_(ctx *_go.String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *GoAstParser) ExitString_(ctx *_go.String_Context) {}

// EnterEmbeddedField is called when production embeddedField is entered.
func (s *GoAstParser) EnterEmbeddedField(ctx *_go.EmbeddedFieldContext) {}

// ExitEmbeddedField is called when production embeddedField is exited.
func (s *GoAstParser) ExitEmbeddedField(ctx *_go.EmbeddedFieldContext) {}

// EnterFunctionLit is called when production functionLit is entered.
func (s *GoAstParser) EnterFunctionLit(ctx *_go.FunctionLitContext) {
}

// ExitFunctionLit is called when production functionLit is exited.
func (s *GoAstParser) ExitFunctionLit(ctx *_go.FunctionLitContext) {}

// EnterIndex is called when production index is entered.
func (s *GoAstParser) EnterIndex(ctx *_go.IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *GoAstParser) ExitIndex(ctx *_go.IndexContext) {}

// EnterSlice_ is called when production slice_ is entered.
func (s *GoAstParser) EnterSlice_(ctx *_go.Slice_Context) {}

// ExitSlice_ is called when production slice_ is exited.
func (s *GoAstParser) ExitSlice_(ctx *_go.Slice_Context) {}

// EnterTypeAssertion is called when production typeAssertion is entered.
func (s *GoAstParser) EnterTypeAssertion(ctx *_go.TypeAssertionContext) {
}

// ExitTypeAssertion is called when production typeAssertion is exited.
func (s *GoAstParser) ExitTypeAssertion(ctx *_go.TypeAssertionContext) {}

// EnterArguments is called when production arguments is entered.
func (s *GoAstParser) EnterArguments(ctx *_go.ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *GoAstParser) ExitArguments(ctx *_go.ArgumentsContext) {}

// EnterMethodExpr is called when production methodExpr is entered.
func (s *GoAstParser) EnterMethodExpr(ctx *_go.MethodExprContext) {}

// ExitMethodExpr is called when production methodExpr is exited.
func (s *GoAstParser) ExitMethodExpr(ctx *_go.MethodExprContext) {}

// EnterReceiverType is called when production receiverType is entered.
func (s *GoAstParser) EnterReceiverType(ctx *_go.ReceiverTypeContext) {}

// ExitReceiverType is called when production receiverType is exited.
func (s *GoAstParser) ExitReceiverType(ctx *_go.ReceiverTypeContext) {}

// EnterEos is called when production eos is entered.
func (s *GoAstParser) EnterEos(ctx *_go.EosContext) {}

// ExitEos is called when production eos is exited.
func (s *GoAstParser) ExitEos(ctx *_go.EosContext) {}
