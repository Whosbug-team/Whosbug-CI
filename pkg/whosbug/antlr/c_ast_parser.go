package antlr

import (
	"strings"
	"sync"

	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr/c"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CAstParser implement AstParser for c language
//
//	@author kevineluo
//	@update 2023-02-28 12:44:18
type CAstParser struct {
	astInfo AstInfo
}

var _ c.CListener = &CAstParser{}

var (
	cLexerPool = &sync.Pool{New: func() any {
		return c.NewCLexer(nil)
	}}
	cParserPool = &sync.Pool{New: func() any {
		return c.NewCParser(nil)
	}}
	newAstParserPool = &sync.Pool{New: func() any {
		return new(CAstParser)
	}}
)

// AstParse main parse process for c language
//
//	@param text string
//	@return AstInfo
//	@author kevineluo
//	@update 2023-02-28 12:55:05
func (s *CAstParser) AstParse(input string) AstInfo {
	//	截取目标文本的输入流
	inputStream := antlr.NewInputStream(input)
	//	初始化 lexer
	lexer := cLexerPool.Get().(*c.CLexer)
	defer cLexerPool.Put(lexer)
	lexer.SetInputStream(inputStream)
	lexer.RemoveErrorListeners()
	//	初始化 Token 流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化 Parser
	p := cParserPool.Get().(*c.CParser)
	defer cParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用 SLL 两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	p.RemoveErrorListeners()
	//	解析模式 -> 每个编译单位
	tree := p.TranslationUnit()
	//	创建 listener
	listener := newAstParserPool.Get().(*CAstParser)
	defer newAstParserPool.Put(listener)
	//	初始化置空
	listener.astInfo = AstInfo{}
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.astInfo
}

// @Description: 获取 C 语言函数的方法名
// @param *c.FunctionDefinitionContext c 语言函数的上下文
// @return methodName 函数的名字
// @author Psy 2022-08-15 14:14:22
// @function_mark PASS
func matchCMethodName(ctx *c.FunctionDefinitionContext) (methodName string) {
	if ctx.Declarator() != nil {
		functionDefineStr := ctx.Declarator().GetText()
		rightIndex := strings.Index(functionDefineStr, "(")
		if rightIndex != -1 {
			spIndex := strings.Index(functionDefineStr, "&")
			if spIndex == -1 {
				spIndex = strings.Index(functionDefineStr, "*")
			}
			if spIndex != -1 && spIndex < rightIndex {
				methodName = functionDefineStr[spIndex+1 : rightIndex]
			} else {
				methodName = functionDefineStr[:rightIndex]
			}
		}
	}
	return
}

// @Description: 获取 C 语言函数的参数
// @param *c.FunctionDefinitionContext c 语言函数的上下文
// @return params 函数的参数
// @author Psy 2022-08-15 14:17:21
// @function_mark PASS
func matchCMethodParams(ctx *c.FunctionDefinitionContext) (params string) {
	if ctx.Declarator() != nil {
		functionDefineStr := ctx.Declarator().GetText()
		leftIndex := strings.Index(functionDefineStr, "(")
		rightIndex := strings.Index(functionDefineStr, ")")
		if leftIndex != -1 && rightIndex != -1 {
			params = functionDefineStr[strings.Index(functionDefineStr, "(")+1 : strings.Index(functionDefineStr, ")")]
		}
	}
	return
}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *CAstParser) ExitFunctionDefinition(ctx *c.FunctionDefinitionContext) {
	if ctx.Declarator() == nil {
		return
	}
	methodInfo := Method{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		Name:       matchCMethodName(ctx),
		Parameters: matchCMethodParams(ctx),
	}
	s.astInfo.Methods = append(s.astInfo.Methods, methodInfo)
}

// EnterStructOrUnionSpecifier is called when production structOrUnionSpecifier is entered.
func (s *CAstParser) EnterStructOrUnionSpecifier(ctx *c.StructOrUnionSpecifierContext) {}

// ExitStructOrUnionSpecifier is called when production structOrUnionSpecifier is exited.
func (s *CAstParser) ExitStructOrUnionSpecifier(ctx *c.StructOrUnionSpecifierContext) {}

// VisitTerminal is called when a terminal node is visited.
func (s *CAstParser) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *CAstParser) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *CAstParser) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *CAstParser) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *CAstParser) EnterPrimaryExpression(ctx *c.PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *CAstParser) ExitPrimaryExpression(ctx *c.PrimaryExpressionContext) {}

// EnterGenericSelection is called when production genericSelection is entered.
func (s *CAstParser) EnterGenericSelection(ctx *c.GenericSelectionContext) {}

// ExitGenericSelection is called when production genericSelection is exited.
func (s *CAstParser) ExitGenericSelection(ctx *c.GenericSelectionContext) {}

// EnterGenericAssocList is called when production genericAssocList is entered.
func (s *CAstParser) EnterGenericAssocList(ctx *c.GenericAssocListContext) {}

// ExitGenericAssocList is called when production genericAssocList is exited.
func (s *CAstParser) ExitGenericAssocList(ctx *c.GenericAssocListContext) {}

// EnterGenericAssociation is called when production genericAssociation is entered.
func (s *CAstParser) EnterGenericAssociation(ctx *c.GenericAssociationContext) {}

// ExitGenericAssociation is called when production genericAssociation is exited.
func (s *CAstParser) ExitGenericAssociation(ctx *c.GenericAssociationContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *CAstParser) EnterPostfixExpression(ctx *c.PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *CAstParser) ExitPostfixExpression(ctx *c.PostfixExpressionContext) {}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *CAstParser) EnterArgumentExpressionList(ctx *c.ArgumentExpressionListContext) {}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *CAstParser) ExitArgumentExpressionList(ctx *c.ArgumentExpressionListContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *CAstParser) EnterUnaryExpression(ctx *c.UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *CAstParser) ExitUnaryExpression(ctx *c.UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *CAstParser) EnterUnaryOperator(ctx *c.UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *CAstParser) ExitUnaryOperator(ctx *c.UnaryOperatorContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *CAstParser) EnterCastExpression(ctx *c.CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *CAstParser) ExitCastExpression(ctx *c.CastExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *CAstParser) EnterMultiplicativeExpression(ctx *c.MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *CAstParser) ExitMultiplicativeExpression(ctx *c.MultiplicativeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *CAstParser) EnterAdditiveExpression(ctx *c.AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *CAstParser) ExitAdditiveExpression(ctx *c.AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *CAstParser) EnterShiftExpression(ctx *c.ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *CAstParser) ExitShiftExpression(ctx *c.ShiftExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *CAstParser) EnterRelationalExpression(ctx *c.RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *CAstParser) ExitRelationalExpression(ctx *c.RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *CAstParser) EnterEqualityExpression(ctx *c.EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *CAstParser) ExitEqualityExpression(ctx *c.EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *CAstParser) EnterAndExpression(ctx *c.AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *CAstParser) ExitAndExpression(ctx *c.AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *CAstParser) EnterExclusiveOrExpression(ctx *c.ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *CAstParser) ExitExclusiveOrExpression(ctx *c.ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *CAstParser) EnterInclusiveOrExpression(ctx *c.InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *CAstParser) ExitInclusiveOrExpression(ctx *c.InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *CAstParser) EnterLogicalAndExpression(ctx *c.LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *CAstParser) ExitLogicalAndExpression(ctx *c.LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *CAstParser) EnterLogicalOrExpression(ctx *c.LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *CAstParser) ExitLogicalOrExpression(ctx *c.LogicalOrExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *CAstParser) EnterConditionalExpression(ctx *c.ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *CAstParser) ExitConditionalExpression(ctx *c.ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *CAstParser) EnterAssignmentExpression(ctx *c.AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *CAstParser) ExitAssignmentExpression(ctx *c.AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *CAstParser) EnterAssignmentOperator(ctx *c.AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *CAstParser) ExitAssignmentOperator(ctx *c.AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *CAstParser) EnterExpression(ctx *c.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *CAstParser) ExitExpression(ctx *c.ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *CAstParser) EnterConstantExpression(ctx *c.ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *CAstParser) ExitConstantExpression(ctx *c.ConstantExpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *CAstParser) EnterDeclaration(ctx *c.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *CAstParser) ExitDeclaration(ctx *c.DeclarationContext) {}

// EnterDeclarationSpecifiers is called when production declarationSpecifiers is entered.
func (s *CAstParser) EnterDeclarationSpecifiers(ctx *c.DeclarationSpecifiersContext) {}

// ExitDeclarationSpecifiers is called when production declarationSpecifiers is exited.
func (s *CAstParser) ExitDeclarationSpecifiers(ctx *c.DeclarationSpecifiersContext) {}

// EnterDeclarationSpecifiers2 is called when production declarationSpecifiers2 is entered.
func (s *CAstParser) EnterDeclarationSpecifiers2(ctx *c.DeclarationSpecifiers2Context) {}

// ExitDeclarationSpecifiers2 is called when production declarationSpecifiers2 is exited.
func (s *CAstParser) ExitDeclarationSpecifiers2(ctx *c.DeclarationSpecifiers2Context) {}

// EnterDeclarationSpecifier is called when production declarationSpecifier is entered.
func (s *CAstParser) EnterDeclarationSpecifier(ctx *c.DeclarationSpecifierContext) {}

// ExitDeclarationSpecifier is called when production declarationSpecifier is exited.
func (s *CAstParser) ExitDeclarationSpecifier(ctx *c.DeclarationSpecifierContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *CAstParser) EnterInitDeclaratorList(ctx *c.InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *CAstParser) ExitInitDeclaratorList(ctx *c.InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *CAstParser) EnterInitDeclarator(ctx *c.InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *CAstParser) ExitInitDeclarator(ctx *c.InitDeclaratorContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *CAstParser) EnterStorageClassSpecifier(ctx *c.StorageClassSpecifierContext) {}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *CAstParser) ExitStorageClassSpecifier(ctx *c.StorageClassSpecifierContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *CAstParser) EnterTypeSpecifier(ctx *c.TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *CAstParser) ExitTypeSpecifier(ctx *c.TypeSpecifierContext) {}

// EnterStructOrUnion is called when production structOrUnion is entered.
func (s *CAstParser) EnterStructOrUnion(ctx *c.StructOrUnionContext) {}

// ExitStructOrUnion is called when production structOrUnion is exited.
func (s *CAstParser) ExitStructOrUnion(ctx *c.StructOrUnionContext) {}

// EnterStructDeclarationList is called when production structDeclarationList is entered.
func (s *CAstParser) EnterStructDeclarationList(ctx *c.StructDeclarationListContext) {}

// ExitStructDeclarationList is called when production structDeclarationList is exited.
func (s *CAstParser) ExitStructDeclarationList(ctx *c.StructDeclarationListContext) {}

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *CAstParser) EnterStructDeclaration(ctx *c.StructDeclarationContext) {}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *CAstParser) ExitStructDeclaration(ctx *c.StructDeclarationContext) {}

// EnterSpecifierQualifierList is called when production specifierQualifierList is entered.
func (s *CAstParser) EnterSpecifierQualifierList(ctx *c.SpecifierQualifierListContext) {}

// ExitSpecifierQualifierList is called when production specifierQualifierList is exited.
func (s *CAstParser) ExitSpecifierQualifierList(ctx *c.SpecifierQualifierListContext) {}

// EnterStructDeclaratorList is called when production structDeclaratorList is entered.
func (s *CAstParser) EnterStructDeclaratorList(ctx *c.StructDeclaratorListContext) {}

// ExitStructDeclaratorList is called when production structDeclaratorList is exited.
func (s *CAstParser) ExitStructDeclaratorList(ctx *c.StructDeclaratorListContext) {}

// EnterStructDeclarator is called when production structDeclarator is entered.
func (s *CAstParser) EnterStructDeclarator(ctx *c.StructDeclaratorContext) {}

// ExitStructDeclarator is called when production structDeclarator is exited.
func (s *CAstParser) ExitStructDeclarator(ctx *c.StructDeclaratorContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *CAstParser) EnterEnumSpecifier(ctx *c.EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *CAstParser) ExitEnumSpecifier(ctx *c.EnumSpecifierContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *CAstParser) EnterEnumeratorList(ctx *c.EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *CAstParser) ExitEnumeratorList(ctx *c.EnumeratorListContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *CAstParser) EnterEnumerator(ctx *c.EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *CAstParser) ExitEnumerator(ctx *c.EnumeratorContext) {}

// EnterEnumerationConstant is called when production enumerationConstant is entered.
func (s *CAstParser) EnterEnumerationConstant(ctx *c.EnumerationConstantContext) {}

// ExitEnumerationConstant is called when production enumerationConstant is exited.
func (s *CAstParser) ExitEnumerationConstant(ctx *c.EnumerationConstantContext) {}

// EnterAtomicTypeSpecifier is called when production atomicTypeSpecifier is entered.
func (s *CAstParser) EnterAtomicTypeSpecifier(ctx *c.AtomicTypeSpecifierContext) {}

// ExitAtomicTypeSpecifier is called when production atomicTypeSpecifier is exited.
func (s *CAstParser) ExitAtomicTypeSpecifier(ctx *c.AtomicTypeSpecifierContext) {}

// EnterTypeQualifier is called when production typeQualifier is entered.
func (s *CAstParser) EnterTypeQualifier(ctx *c.TypeQualifierContext) {}

// ExitTypeQualifier is called when production typeQualifier is exited.
func (s *CAstParser) ExitTypeQualifier(ctx *c.TypeQualifierContext) {}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *CAstParser) EnterFunctionSpecifier(ctx *c.FunctionSpecifierContext) {}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *CAstParser) ExitFunctionSpecifier(ctx *c.FunctionSpecifierContext) {}

// EnterAlignmentSpecifier is called when production alignmentSpecifier is entered.
func (s *CAstParser) EnterAlignmentSpecifier(ctx *c.AlignmentSpecifierContext) {}

// ExitAlignmentSpecifier is called when production alignmentSpecifier is exited.
func (s *CAstParser) ExitAlignmentSpecifier(ctx *c.AlignmentSpecifierContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *CAstParser) EnterDeclarator(ctx *c.DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *CAstParser) ExitDeclarator(ctx *c.DeclaratorContext) {}

// EnterDirectDeclarator is called when production directDeclarator is entered.
func (s *CAstParser) EnterDirectDeclarator(ctx *c.DirectDeclaratorContext) {}

// ExitDirectDeclarator is called when production directDeclarator is exited.
func (s *CAstParser) ExitDirectDeclarator(ctx *c.DirectDeclaratorContext) {}

// EnterVcSpecificModifer is called when production vcSpecificModifer is entered.
func (s *CAstParser) EnterVcSpecificModifer(ctx *c.VcSpecificModiferContext) {}

// ExitVcSpecificModifer is called when production vcSpecificModifer is exited.
func (s *CAstParser) ExitVcSpecificModifer(ctx *c.VcSpecificModiferContext) {}

// EnterGccDeclaratorExtension is called when production gccDeclaratorExtension is entered.
func (s *CAstParser) EnterGccDeclaratorExtension(ctx *c.GccDeclaratorExtensionContext) {}

// ExitGccDeclaratorExtension is called when production gccDeclaratorExtension is exited.
func (s *CAstParser) ExitGccDeclaratorExtension(ctx *c.GccDeclaratorExtensionContext) {}

// EnterGccAttributeSpecifier is called when production gccAttributeSpecifier is entered.
func (s *CAstParser) EnterGccAttributeSpecifier(ctx *c.GccAttributeSpecifierContext) {}

// ExitGccAttributeSpecifier is called when production gccAttributeSpecifier is exited.
func (s *CAstParser) ExitGccAttributeSpecifier(ctx *c.GccAttributeSpecifierContext) {}

// EnterGccAttributeList is called when production gccAttributeList is entered.
func (s *CAstParser) EnterGccAttributeList(ctx *c.GccAttributeListContext) {}

// ExitGccAttributeList is called when production gccAttributeList is exited.
func (s *CAstParser) ExitGccAttributeList(ctx *c.GccAttributeListContext) {}

// EnterGccAttribute is called when production gccAttribute is entered.
func (s *CAstParser) EnterGccAttribute(ctx *c.GccAttributeContext) {}

// ExitGccAttribute is called when production gccAttribute is exited.
func (s *CAstParser) ExitGccAttribute(ctx *c.GccAttributeContext) {}

// EnterNestedParenthesesBlock is called when production nestedParenthesesBlock is entered.
func (s *CAstParser) EnterNestedParenthesesBlock(ctx *c.NestedParenthesesBlockContext) {}

// ExitNestedParenthesesBlock is called when production nestedParenthesesBlock is exited.
func (s *CAstParser) ExitNestedParenthesesBlock(ctx *c.NestedParenthesesBlockContext) {}

// EnterPointer is called when production pointer is entered.
func (s *CAstParser) EnterPointer(ctx *c.PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *CAstParser) ExitPointer(ctx *c.PointerContext) {}

// EnterTypeQualifierList is called when production typeQualifierList is entered.
func (s *CAstParser) EnterTypeQualifierList(ctx *c.TypeQualifierListContext) {}

// ExitTypeQualifierList is called when production typeQualifierList is exited.
func (s *CAstParser) ExitTypeQualifierList(ctx *c.TypeQualifierListContext) {}

// EnterParameterTypeList is called when production parameterTypeList is entered.
func (s *CAstParser) EnterParameterTypeList(ctx *c.ParameterTypeListContext) {}

// ExitParameterTypeList is called when production parameterTypeList is exited.
func (s *CAstParser) ExitParameterTypeList(ctx *c.ParameterTypeListContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *CAstParser) EnterParameterList(ctx *c.ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *CAstParser) ExitParameterList(ctx *c.ParameterListContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *CAstParser) EnterParameterDeclaration(ctx *c.ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *CAstParser) ExitParameterDeclaration(ctx *c.ParameterDeclarationContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *CAstParser) EnterIdentifierList(ctx *c.IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *CAstParser) ExitIdentifierList(ctx *c.IdentifierListContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *CAstParser) EnterTypeName(ctx *c.TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *CAstParser) ExitTypeName(ctx *c.TypeNameContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *CAstParser) EnterAbstractDeclarator(ctx *c.AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *CAstParser) ExitAbstractDeclarator(ctx *c.AbstractDeclaratorContext) {}

// EnterDirectAbstractDeclarator is called when production directAbstractDeclarator is entered.
func (s *CAstParser) EnterDirectAbstractDeclarator(ctx *c.DirectAbstractDeclaratorContext) {}

// ExitDirectAbstractDeclarator is called when production directAbstractDeclarator is exited.
func (s *CAstParser) ExitDirectAbstractDeclarator(ctx *c.DirectAbstractDeclaratorContext) {}

// EnterTypedefName is called when production typedefName is entered.
func (s *CAstParser) EnterTypedefName(ctx *c.TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *CAstParser) ExitTypedefName(ctx *c.TypedefNameContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *CAstParser) EnterInitializer(ctx *c.InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *CAstParser) ExitInitializer(ctx *c.InitializerContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *CAstParser) EnterInitializerList(ctx *c.InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *CAstParser) ExitInitializerList(ctx *c.InitializerListContext) {}

// EnterDesignation is called when production designation is entered.
func (s *CAstParser) EnterDesignation(ctx *c.DesignationContext) {}

// ExitDesignation is called when production designation is exited.
func (s *CAstParser) ExitDesignation(ctx *c.DesignationContext) {}

// EnterDesignatorList is called when production designatorList is entered.
func (s *CAstParser) EnterDesignatorList(ctx *c.DesignatorListContext) {}

// ExitDesignatorList is called when production designatorList is exited.
func (s *CAstParser) ExitDesignatorList(ctx *c.DesignatorListContext) {}

// EnterDesignator is called when production designator is entered.
func (s *CAstParser) EnterDesignator(ctx *c.DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *CAstParser) ExitDesignator(ctx *c.DesignatorContext) {}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *CAstParser) EnterStaticAssertDeclaration(ctx *c.StaticAssertDeclarationContext) {}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *CAstParser) ExitStaticAssertDeclaration(ctx *c.StaticAssertDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *CAstParser) EnterStatement(ctx *c.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *CAstParser) ExitStatement(ctx *c.StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *CAstParser) EnterLabeledStatement(ctx *c.LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *CAstParser) ExitLabeledStatement(ctx *c.LabeledStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *CAstParser) EnterCompoundStatement(ctx *c.CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *CAstParser) ExitCompoundStatement(ctx *c.CompoundStatementContext) {}

// EnterBlockItemList is called when production blockItemList is entered.
func (s *CAstParser) EnterBlockItemList(ctx *c.BlockItemListContext) {}

// ExitBlockItemList is called when production blockItemList is exited.
func (s *CAstParser) ExitBlockItemList(ctx *c.BlockItemListContext) {}

// EnterBlockItem is called when production blockItem is entered.
func (s *CAstParser) EnterBlockItem(ctx *c.BlockItemContext) {}

// ExitBlockItem is called when production blockItem is exited.
func (s *CAstParser) ExitBlockItem(ctx *c.BlockItemContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *CAstParser) EnterExpressionStatement(ctx *c.ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *CAstParser) ExitExpressionStatement(ctx *c.ExpressionStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *CAstParser) EnterSelectionStatement(ctx *c.SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *CAstParser) ExitSelectionStatement(ctx *c.SelectionStatementContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *CAstParser) EnterIterationStatement(ctx *c.IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *CAstParser) ExitIterationStatement(ctx *c.IterationStatementContext) {}

// EnterForCondition is called when production forCondition is entered.
func (s *CAstParser) EnterForCondition(ctx *c.ForConditionContext) {}

// ExitForCondition is called when production forCondition is exited.
func (s *CAstParser) ExitForCondition(ctx *c.ForConditionContext) {}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *CAstParser) EnterForDeclaration(ctx *c.ForDeclarationContext) {}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *CAstParser) ExitForDeclaration(ctx *c.ForDeclarationContext) {}

// EnterForExpression is called when production forExpression is entered.
func (s *CAstParser) EnterForExpression(ctx *c.ForExpressionContext) {}

// ExitForExpression is called when production forExpression is exited.
func (s *CAstParser) ExitForExpression(ctx *c.ForExpressionContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *CAstParser) EnterJumpStatement(ctx *c.JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *CAstParser) ExitJumpStatement(ctx *c.JumpStatementContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *CAstParser) EnterCompilationUnit(ctx *c.CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *CAstParser) ExitCompilationUnit(ctx *c.CompilationUnitContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *CAstParser) EnterTranslationUnit(ctx *c.TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *CAstParser) ExitTranslationUnit(ctx *c.TranslationUnitContext) {}

// EnterExternalDeclaration is called when production externalDeclaration is entered.
func (s *CAstParser) EnterExternalDeclaration(ctx *c.ExternalDeclarationContext) {}

// ExitExternalDeclaration is called when production externalDeclaration is exited.
func (s *CAstParser) ExitExternalDeclaration(ctx *c.ExternalDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *CAstParser) EnterFunctionDefinition(ctx *c.FunctionDefinitionContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *CAstParser) EnterDeclarationList(ctx *c.DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *CAstParser) ExitDeclarationList(ctx *c.DeclarationListContext) {}
