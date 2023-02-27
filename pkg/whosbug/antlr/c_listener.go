package antlr // C

import (
	"strings"

	c "git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr/cLib"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CTreeShapeListener struct {
	AstInfoList AstResType
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
func (s *CTreeShapeListener) ExitFunctionDefinition(ctx *c.FunctionDefinitionContext) {
	if ctx.Declarator() == nil {
		return
	}
	methodInfo := MethodInfoType{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		MethodName: matchCMethodName(ctx),
		Parameters: matchCMethodParams(ctx),
	}
	s.AstInfoList.Methods = append(s.AstInfoList.Methods, methodInfo)
}

// EnterStructOrUnionSpecifier is called when production structOrUnionSpecifier is entered.
func (s *CTreeShapeListener) EnterStructOrUnionSpecifier(ctx *c.StructOrUnionSpecifierContext) {}

// ExitStructOrUnionSpecifier is called when production structOrUnionSpecifier is exited.
func (s *CTreeShapeListener) ExitStructOrUnionSpecifier(ctx *c.StructOrUnionSpecifierContext) {}

// VisitTerminal is called when a terminal node is visited.
func (s *CTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *CTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *CTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *CTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *CTreeShapeListener) EnterPrimaryExpression(ctx *c.PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *CTreeShapeListener) ExitPrimaryExpression(ctx *c.PrimaryExpressionContext) {}

// EnterGenericSelection is called when production genericSelection is entered.
func (s *CTreeShapeListener) EnterGenericSelection(ctx *c.GenericSelectionContext) {}

// ExitGenericSelection is called when production genericSelection is exited.
func (s *CTreeShapeListener) ExitGenericSelection(ctx *c.GenericSelectionContext) {}

// EnterGenericAssocList is called when production genericAssocList is entered.
func (s *CTreeShapeListener) EnterGenericAssocList(ctx *c.GenericAssocListContext) {}

// ExitGenericAssocList is called when production genericAssocList is exited.
func (s *CTreeShapeListener) ExitGenericAssocList(ctx *c.GenericAssocListContext) {}

// EnterGenericAssociation is called when production genericAssociation is entered.
func (s *CTreeShapeListener) EnterGenericAssociation(ctx *c.GenericAssociationContext) {}

// ExitGenericAssociation is called when production genericAssociation is exited.
func (s *CTreeShapeListener) ExitGenericAssociation(ctx *c.GenericAssociationContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *CTreeShapeListener) EnterPostfixExpression(ctx *c.PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *CTreeShapeListener) ExitPostfixExpression(ctx *c.PostfixExpressionContext) {}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *CTreeShapeListener) EnterArgumentExpressionList(ctx *c.ArgumentExpressionListContext) {}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *CTreeShapeListener) ExitArgumentExpressionList(ctx *c.ArgumentExpressionListContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *CTreeShapeListener) EnterUnaryExpression(ctx *c.UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *CTreeShapeListener) ExitUnaryExpression(ctx *c.UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *CTreeShapeListener) EnterUnaryOperator(ctx *c.UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *CTreeShapeListener) ExitUnaryOperator(ctx *c.UnaryOperatorContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *CTreeShapeListener) EnterCastExpression(ctx *c.CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *CTreeShapeListener) ExitCastExpression(ctx *c.CastExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *CTreeShapeListener) EnterMultiplicativeExpression(ctx *c.MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *CTreeShapeListener) ExitMultiplicativeExpression(ctx *c.MultiplicativeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *CTreeShapeListener) EnterAdditiveExpression(ctx *c.AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *CTreeShapeListener) ExitAdditiveExpression(ctx *c.AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *CTreeShapeListener) EnterShiftExpression(ctx *c.ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *CTreeShapeListener) ExitShiftExpression(ctx *c.ShiftExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *CTreeShapeListener) EnterRelationalExpression(ctx *c.RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *CTreeShapeListener) ExitRelationalExpression(ctx *c.RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *CTreeShapeListener) EnterEqualityExpression(ctx *c.EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *CTreeShapeListener) ExitEqualityExpression(ctx *c.EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *CTreeShapeListener) EnterAndExpression(ctx *c.AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *CTreeShapeListener) ExitAndExpression(ctx *c.AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *CTreeShapeListener) EnterExclusiveOrExpression(ctx *c.ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *CTreeShapeListener) ExitExclusiveOrExpression(ctx *c.ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *CTreeShapeListener) EnterInclusiveOrExpression(ctx *c.InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *CTreeShapeListener) ExitInclusiveOrExpression(ctx *c.InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *CTreeShapeListener) EnterLogicalAndExpression(ctx *c.LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *CTreeShapeListener) ExitLogicalAndExpression(ctx *c.LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *CTreeShapeListener) EnterLogicalOrExpression(ctx *c.LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *CTreeShapeListener) ExitLogicalOrExpression(ctx *c.LogicalOrExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *CTreeShapeListener) EnterConditionalExpression(ctx *c.ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *CTreeShapeListener) ExitConditionalExpression(ctx *c.ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *CTreeShapeListener) EnterAssignmentExpression(ctx *c.AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *CTreeShapeListener) ExitAssignmentExpression(ctx *c.AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *CTreeShapeListener) EnterAssignmentOperator(ctx *c.AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *CTreeShapeListener) ExitAssignmentOperator(ctx *c.AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *CTreeShapeListener) EnterExpression(ctx *c.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *CTreeShapeListener) ExitExpression(ctx *c.ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *CTreeShapeListener) EnterConstantExpression(ctx *c.ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *CTreeShapeListener) ExitConstantExpression(ctx *c.ConstantExpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *CTreeShapeListener) EnterDeclaration(ctx *c.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *CTreeShapeListener) ExitDeclaration(ctx *c.DeclarationContext) {}

// EnterDeclarationSpecifiers is called when production declarationSpecifiers is entered.
func (s *CTreeShapeListener) EnterDeclarationSpecifiers(ctx *c.DeclarationSpecifiersContext) {}

// ExitDeclarationSpecifiers is called when production declarationSpecifiers is exited.
func (s *CTreeShapeListener) ExitDeclarationSpecifiers(ctx *c.DeclarationSpecifiersContext) {}

// EnterDeclarationSpecifiers2 is called when production declarationSpecifiers2 is entered.
func (s *CTreeShapeListener) EnterDeclarationSpecifiers2(ctx *c.DeclarationSpecifiers2Context) {}

// ExitDeclarationSpecifiers2 is called when production declarationSpecifiers2 is exited.
func (s *CTreeShapeListener) ExitDeclarationSpecifiers2(ctx *c.DeclarationSpecifiers2Context) {}

// EnterDeclarationSpecifier is called when production declarationSpecifier is entered.
func (s *CTreeShapeListener) EnterDeclarationSpecifier(ctx *c.DeclarationSpecifierContext) {}

// ExitDeclarationSpecifier is called when production declarationSpecifier is exited.
func (s *CTreeShapeListener) ExitDeclarationSpecifier(ctx *c.DeclarationSpecifierContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *CTreeShapeListener) EnterInitDeclaratorList(ctx *c.InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *CTreeShapeListener) ExitInitDeclaratorList(ctx *c.InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *CTreeShapeListener) EnterInitDeclarator(ctx *c.InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *CTreeShapeListener) ExitInitDeclarator(ctx *c.InitDeclaratorContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *CTreeShapeListener) EnterStorageClassSpecifier(ctx *c.StorageClassSpecifierContext) {}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *CTreeShapeListener) ExitStorageClassSpecifier(ctx *c.StorageClassSpecifierContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *CTreeShapeListener) EnterTypeSpecifier(ctx *c.TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *CTreeShapeListener) ExitTypeSpecifier(ctx *c.TypeSpecifierContext) {}

// EnterStructOrUnion is called when production structOrUnion is entered.
func (s *CTreeShapeListener) EnterStructOrUnion(ctx *c.StructOrUnionContext) {}

// ExitStructOrUnion is called when production structOrUnion is exited.
func (s *CTreeShapeListener) ExitStructOrUnion(ctx *c.StructOrUnionContext) {}

// EnterStructDeclarationList is called when production structDeclarationList is entered.
func (s *CTreeShapeListener) EnterStructDeclarationList(ctx *c.StructDeclarationListContext) {}

// ExitStructDeclarationList is called when production structDeclarationList is exited.
func (s *CTreeShapeListener) ExitStructDeclarationList(ctx *c.StructDeclarationListContext) {}

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *CTreeShapeListener) EnterStructDeclaration(ctx *c.StructDeclarationContext) {}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *CTreeShapeListener) ExitStructDeclaration(ctx *c.StructDeclarationContext) {}

// EnterSpecifierQualifierList is called when production specifierQualifierList is entered.
func (s *CTreeShapeListener) EnterSpecifierQualifierList(ctx *c.SpecifierQualifierListContext) {}

// ExitSpecifierQualifierList is called when production specifierQualifierList is exited.
func (s *CTreeShapeListener) ExitSpecifierQualifierList(ctx *c.SpecifierQualifierListContext) {}

// EnterStructDeclaratorList is called when production structDeclaratorList is entered.
func (s *CTreeShapeListener) EnterStructDeclaratorList(ctx *c.StructDeclaratorListContext) {}

// ExitStructDeclaratorList is called when production structDeclaratorList is exited.
func (s *CTreeShapeListener) ExitStructDeclaratorList(ctx *c.StructDeclaratorListContext) {}

// EnterStructDeclarator is called when production structDeclarator is entered.
func (s *CTreeShapeListener) EnterStructDeclarator(ctx *c.StructDeclaratorContext) {}

// ExitStructDeclarator is called when production structDeclarator is exited.
func (s *CTreeShapeListener) ExitStructDeclarator(ctx *c.StructDeclaratorContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *CTreeShapeListener) EnterEnumSpecifier(ctx *c.EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *CTreeShapeListener) ExitEnumSpecifier(ctx *c.EnumSpecifierContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *CTreeShapeListener) EnterEnumeratorList(ctx *c.EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *CTreeShapeListener) ExitEnumeratorList(ctx *c.EnumeratorListContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *CTreeShapeListener) EnterEnumerator(ctx *c.EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *CTreeShapeListener) ExitEnumerator(ctx *c.EnumeratorContext) {}

// EnterEnumerationConstant is called when production enumerationConstant is entered.
func (s *CTreeShapeListener) EnterEnumerationConstant(ctx *c.EnumerationConstantContext) {}

// ExitEnumerationConstant is called when production enumerationConstant is exited.
func (s *CTreeShapeListener) ExitEnumerationConstant(ctx *c.EnumerationConstantContext) {}

// EnterAtomicTypeSpecifier is called when production atomicTypeSpecifier is entered.
func (s *CTreeShapeListener) EnterAtomicTypeSpecifier(ctx *c.AtomicTypeSpecifierContext) {}

// ExitAtomicTypeSpecifier is called when production atomicTypeSpecifier is exited.
func (s *CTreeShapeListener) ExitAtomicTypeSpecifier(ctx *c.AtomicTypeSpecifierContext) {}

// EnterTypeQualifier is called when production typeQualifier is entered.
func (s *CTreeShapeListener) EnterTypeQualifier(ctx *c.TypeQualifierContext) {}

// ExitTypeQualifier is called when production typeQualifier is exited.
func (s *CTreeShapeListener) ExitTypeQualifier(ctx *c.TypeQualifierContext) {}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *CTreeShapeListener) EnterFunctionSpecifier(ctx *c.FunctionSpecifierContext) {}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *CTreeShapeListener) ExitFunctionSpecifier(ctx *c.FunctionSpecifierContext) {}

// EnterAlignmentSpecifier is called when production alignmentSpecifier is entered.
func (s *CTreeShapeListener) EnterAlignmentSpecifier(ctx *c.AlignmentSpecifierContext) {}

// ExitAlignmentSpecifier is called when production alignmentSpecifier is exited.
func (s *CTreeShapeListener) ExitAlignmentSpecifier(ctx *c.AlignmentSpecifierContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *CTreeShapeListener) EnterDeclarator(ctx *c.DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *CTreeShapeListener) ExitDeclarator(ctx *c.DeclaratorContext) {}

// EnterDirectDeclarator is called when production directDeclarator is entered.
func (s *CTreeShapeListener) EnterDirectDeclarator(ctx *c.DirectDeclaratorContext) {}

// ExitDirectDeclarator is called when production directDeclarator is exited.
func (s *CTreeShapeListener) ExitDirectDeclarator(ctx *c.DirectDeclaratorContext) {}

// EnterVcSpecificModifer is called when production vcSpecificModifer is entered.
func (s *CTreeShapeListener) EnterVcSpecificModifer(ctx *c.VcSpecificModiferContext) {}

// ExitVcSpecificModifer is called when production vcSpecificModifer is exited.
func (s *CTreeShapeListener) ExitVcSpecificModifer(ctx *c.VcSpecificModiferContext) {}

// EnterGccDeclaratorExtension is called when production gccDeclaratorExtension is entered.
func (s *CTreeShapeListener) EnterGccDeclaratorExtension(ctx *c.GccDeclaratorExtensionContext) {}

// ExitGccDeclaratorExtension is called when production gccDeclaratorExtension is exited.
func (s *CTreeShapeListener) ExitGccDeclaratorExtension(ctx *c.GccDeclaratorExtensionContext) {}

// EnterGccAttributeSpecifier is called when production gccAttributeSpecifier is entered.
func (s *CTreeShapeListener) EnterGccAttributeSpecifier(ctx *c.GccAttributeSpecifierContext) {}

// ExitGccAttributeSpecifier is called when production gccAttributeSpecifier is exited.
func (s *CTreeShapeListener) ExitGccAttributeSpecifier(ctx *c.GccAttributeSpecifierContext) {}

// EnterGccAttributeList is called when production gccAttributeList is entered.
func (s *CTreeShapeListener) EnterGccAttributeList(ctx *c.GccAttributeListContext) {}

// ExitGccAttributeList is called when production gccAttributeList is exited.
func (s *CTreeShapeListener) ExitGccAttributeList(ctx *c.GccAttributeListContext) {}

// EnterGccAttribute is called when production gccAttribute is entered.
func (s *CTreeShapeListener) EnterGccAttribute(ctx *c.GccAttributeContext) {}

// ExitGccAttribute is called when production gccAttribute is exited.
func (s *CTreeShapeListener) ExitGccAttribute(ctx *c.GccAttributeContext) {}

// EnterNestedParenthesesBlock is called when production nestedParenthesesBlock is entered.
func (s *CTreeShapeListener) EnterNestedParenthesesBlock(ctx *c.NestedParenthesesBlockContext) {}

// ExitNestedParenthesesBlock is called when production nestedParenthesesBlock is exited.
func (s *CTreeShapeListener) ExitNestedParenthesesBlock(ctx *c.NestedParenthesesBlockContext) {}

// EnterPointer is called when production pointer is entered.
func (s *CTreeShapeListener) EnterPointer(ctx *c.PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *CTreeShapeListener) ExitPointer(ctx *c.PointerContext) {}

// EnterTypeQualifierList is called when production typeQualifierList is entered.
func (s *CTreeShapeListener) EnterTypeQualifierList(ctx *c.TypeQualifierListContext) {}

// ExitTypeQualifierList is called when production typeQualifierList is exited.
func (s *CTreeShapeListener) ExitTypeQualifierList(ctx *c.TypeQualifierListContext) {}

// EnterParameterTypeList is called when production parameterTypeList is entered.
func (s *CTreeShapeListener) EnterParameterTypeList(ctx *c.ParameterTypeListContext) {}

// ExitParameterTypeList is called when production parameterTypeList is exited.
func (s *CTreeShapeListener) ExitParameterTypeList(ctx *c.ParameterTypeListContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *CTreeShapeListener) EnterParameterList(ctx *c.ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *CTreeShapeListener) ExitParameterList(ctx *c.ParameterListContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *CTreeShapeListener) EnterParameterDeclaration(ctx *c.ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *CTreeShapeListener) ExitParameterDeclaration(ctx *c.ParameterDeclarationContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *CTreeShapeListener) EnterIdentifierList(ctx *c.IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *CTreeShapeListener) ExitIdentifierList(ctx *c.IdentifierListContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *CTreeShapeListener) EnterTypeName(ctx *c.TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *CTreeShapeListener) ExitTypeName(ctx *c.TypeNameContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *CTreeShapeListener) EnterAbstractDeclarator(ctx *c.AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *CTreeShapeListener) ExitAbstractDeclarator(ctx *c.AbstractDeclaratorContext) {}

// EnterDirectAbstractDeclarator is called when production directAbstractDeclarator is entered.
func (s *CTreeShapeListener) EnterDirectAbstractDeclarator(ctx *c.DirectAbstractDeclaratorContext) {}

// ExitDirectAbstractDeclarator is called when production directAbstractDeclarator is exited.
func (s *CTreeShapeListener) ExitDirectAbstractDeclarator(ctx *c.DirectAbstractDeclaratorContext) {}

// EnterTypedefName is called when production typedefName is entered.
func (s *CTreeShapeListener) EnterTypedefName(ctx *c.TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *CTreeShapeListener) ExitTypedefName(ctx *c.TypedefNameContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *CTreeShapeListener) EnterInitializer(ctx *c.InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *CTreeShapeListener) ExitInitializer(ctx *c.InitializerContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *CTreeShapeListener) EnterInitializerList(ctx *c.InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *CTreeShapeListener) ExitInitializerList(ctx *c.InitializerListContext) {}

// EnterDesignation is called when production designation is entered.
func (s *CTreeShapeListener) EnterDesignation(ctx *c.DesignationContext) {}

// ExitDesignation is called when production designation is exited.
func (s *CTreeShapeListener) ExitDesignation(ctx *c.DesignationContext) {}

// EnterDesignatorList is called when production designatorList is entered.
func (s *CTreeShapeListener) EnterDesignatorList(ctx *c.DesignatorListContext) {}

// ExitDesignatorList is called when production designatorList is exited.
func (s *CTreeShapeListener) ExitDesignatorList(ctx *c.DesignatorListContext) {}

// EnterDesignator is called when production designator is entered.
func (s *CTreeShapeListener) EnterDesignator(ctx *c.DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *CTreeShapeListener) ExitDesignator(ctx *c.DesignatorContext) {}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *CTreeShapeListener) EnterStaticAssertDeclaration(ctx *c.StaticAssertDeclarationContext) {}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *CTreeShapeListener) ExitStaticAssertDeclaration(ctx *c.StaticAssertDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *CTreeShapeListener) EnterStatement(ctx *c.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *CTreeShapeListener) ExitStatement(ctx *c.StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *CTreeShapeListener) EnterLabeledStatement(ctx *c.LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *CTreeShapeListener) ExitLabeledStatement(ctx *c.LabeledStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *CTreeShapeListener) EnterCompoundStatement(ctx *c.CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *CTreeShapeListener) ExitCompoundStatement(ctx *c.CompoundStatementContext) {}

// EnterBlockItemList is called when production blockItemList is entered.
func (s *CTreeShapeListener) EnterBlockItemList(ctx *c.BlockItemListContext) {}

// ExitBlockItemList is called when production blockItemList is exited.
func (s *CTreeShapeListener) ExitBlockItemList(ctx *c.BlockItemListContext) {}

// EnterBlockItem is called when production blockItem is entered.
func (s *CTreeShapeListener) EnterBlockItem(ctx *c.BlockItemContext) {}

// ExitBlockItem is called when production blockItem is exited.
func (s *CTreeShapeListener) ExitBlockItem(ctx *c.BlockItemContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *CTreeShapeListener) EnterExpressionStatement(ctx *c.ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *CTreeShapeListener) ExitExpressionStatement(ctx *c.ExpressionStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *CTreeShapeListener) EnterSelectionStatement(ctx *c.SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *CTreeShapeListener) ExitSelectionStatement(ctx *c.SelectionStatementContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *CTreeShapeListener) EnterIterationStatement(ctx *c.IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *CTreeShapeListener) ExitIterationStatement(ctx *c.IterationStatementContext) {}

// EnterForCondition is called when production forCondition is entered.
func (s *CTreeShapeListener) EnterForCondition(ctx *c.ForConditionContext) {}

// ExitForCondition is called when production forCondition is exited.
func (s *CTreeShapeListener) ExitForCondition(ctx *c.ForConditionContext) {}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *CTreeShapeListener) EnterForDeclaration(ctx *c.ForDeclarationContext) {}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *CTreeShapeListener) ExitForDeclaration(ctx *c.ForDeclarationContext) {}

// EnterForExpression is called when production forExpression is entered.
func (s *CTreeShapeListener) EnterForExpression(ctx *c.ForExpressionContext) {}

// ExitForExpression is called when production forExpression is exited.
func (s *CTreeShapeListener) ExitForExpression(ctx *c.ForExpressionContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *CTreeShapeListener) EnterJumpStatement(ctx *c.JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *CTreeShapeListener) ExitJumpStatement(ctx *c.JumpStatementContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *CTreeShapeListener) EnterCompilationUnit(ctx *c.CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *CTreeShapeListener) ExitCompilationUnit(ctx *c.CompilationUnitContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *CTreeShapeListener) EnterTranslationUnit(ctx *c.TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *CTreeShapeListener) ExitTranslationUnit(ctx *c.TranslationUnitContext) {}

// EnterExternalDeclaration is called when production externalDeclaration is entered.
func (s *CTreeShapeListener) EnterExternalDeclaration(ctx *c.ExternalDeclarationContext) {}

// ExitExternalDeclaration is called when production externalDeclaration is exited.
func (s *CTreeShapeListener) ExitExternalDeclaration(ctx *c.ExternalDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *CTreeShapeListener) EnterFunctionDefinition(ctx *c.FunctionDefinitionContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *CTreeShapeListener) EnterDeclarationList(ctx *c.DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *CTreeShapeListener) ExitDeclarationList(ctx *c.DeclarationListContext) {}
