// Generated from KotlinParser.g4 by ANTLR 4.7.

package antlr // KotlinParser

import (
	kotlin "git.woa.com/bkdevops/whosbug/antlr/kotlinLib"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ kotlin.KotlinParserListener = &KotlinTreeShapeListener{}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterFunctionDeclaration(ctx *kotlin.FunctionDeclarationContext) {
	if ctx.Identifier() == nil {
		return
	}
	methodInfo := MethodInfoType{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		MethodName: findKotlinDeclarationChain(ctx) + ctx.Identifier().GetText(),
		Parameters: "",
	}
	s.AstInfoList.Methods = append(s.AstInfoList.Methods, methodInfo)
}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterClassDeclaration(ctx *kotlin.ClassDeclarationContext) {
	classInfo := ClassInfoType{
		StartLine: ctx.GetStart().GetLine(),
		EndLine:   ctx.GetStop().GetLine(),
		ClassName: findKotlinDeclarationChain(ctx) + ctx.SimpleIdentifier().GetText(),
		Extends:   findKotlinClassExtends(ctx),
	}
	s.AstInfoList.Classes = append(s.AstInfoList.Classes, classInfo)
}

func findKotlinClassExtends(ctx *kotlin.ClassDeclarationContext) (extends string) {
	if ctx.DelegationSpecifiers() != nil {
		tempCtx := ctx.DelegationSpecifiers().(*kotlin.DelegationSpecifiersContext)
		if tempCtx.AllDelegationSpecifier() != nil {
			if tempCtx.DelegationSpecifier(0).(*kotlin.DelegationSpecifierContext).ConstructorInvocation() != nil {
				tempCtx2 := tempCtx.DelegationSpecifier(0).(*kotlin.DelegationSpecifierContext).ConstructorInvocation()
				extends = tempCtx2.(*kotlin.ConstructorInvocationContext).UserType().GetText()
			}
		}
	}
	return
}

func findKotlinDeclarationChain(ctx antlr.ParseTree) (chainName string) {
	currentContext := ctx.GetParent()
	for {
		if _, ok := currentContext.(*kotlin.ClassDeclarationContext); ok {
			chainName = currentContext.(*kotlin.ClassDeclarationContext).SimpleIdentifier().GetText() + "." + chainName
		}
		if _, ok := currentContext.(*kotlin.FunctionDeclarationContext); ok {
			chainName = currentContext.(*kotlin.FunctionDeclarationContext).Identifier().GetText() + "." + chainName
		}
		currentContext = currentContext.GetParent()
		if currentContext == nil {
			break
		}
	}
	return
}

// ?Antlr-GoStdKotLibFuncs
func (s *KotlinTreeShapeListener) ExitPostfixUnaryOperation(ctx *kotlin.PostfixUnaryOperationContext) {
}

// VisitTerminal is called when a terminal node is visited.
func (s *KotlinTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *KotlinTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *KotlinTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *KotlinTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKotlinFile is called when production kotlinFile is entered.
func (s *KotlinTreeShapeListener) EnterKotlinFile(ctx *kotlin.KotlinFileContext) {}

// ExitKotlinFile is called when production kotlinFile is exited.
func (s *KotlinTreeShapeListener) ExitKotlinFile(ctx *kotlin.KotlinFileContext) {}

// EnterScript is called when production script is entered.
func (s *KotlinTreeShapeListener) EnterScript(ctx *kotlin.ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *KotlinTreeShapeListener) ExitScript(ctx *kotlin.ScriptContext) {}

// EnterPreamble is called when production preamble is entered.
func (s *KotlinTreeShapeListener) EnterPreamble(ctx *kotlin.PreambleContext) {}

// ExitPreamble is called when production preamble is exited.
func (s *KotlinTreeShapeListener) ExitPreamble(ctx *kotlin.PreambleContext) {}

// EnterFileAnnotations is called when production fileAnnotations is entered.
func (s *KotlinTreeShapeListener) EnterFileAnnotations(ctx *kotlin.FileAnnotationsContext) {}

// ExitFileAnnotations is called when production fileAnnotations is exited.
func (s *KotlinTreeShapeListener) ExitFileAnnotations(ctx *kotlin.FileAnnotationsContext) {}

// EnterFileAnnotation is called when production fileAnnotation is entered.
func (s *KotlinTreeShapeListener) EnterFileAnnotation(ctx *kotlin.FileAnnotationContext) {}

// ExitFileAnnotation is called when production fileAnnotation is exited.
func (s *KotlinTreeShapeListener) ExitFileAnnotation(ctx *kotlin.FileAnnotationContext) {}

// EnterPackageHeader is called when production packageHeader is entered.
func (s *KotlinTreeShapeListener) EnterPackageHeader(ctx *kotlin.PackageHeaderContext) {}

// ExitPackageHeader is called when production packageHeader is exited.
func (s *KotlinTreeShapeListener) ExitPackageHeader(ctx *kotlin.PackageHeaderContext) {}

// EnterImportList is called when production importList is entered.
func (s *KotlinTreeShapeListener) EnterImportList(ctx *kotlin.ImportListContext) {}

// ExitImportList is called when production importList is exited.
func (s *KotlinTreeShapeListener) ExitImportList(ctx *kotlin.ImportListContext) {}

// EnterImportHeader is called when production importHeader is entered.
func (s *KotlinTreeShapeListener) EnterImportHeader(ctx *kotlin.ImportHeaderContext) {}

// ExitImportHeader is called when production importHeader is exited.
func (s *KotlinTreeShapeListener) ExitImportHeader(ctx *kotlin.ImportHeaderContext) {}

// EnterImportAlias is called when production importAlias is entered.
func (s *KotlinTreeShapeListener) EnterImportAlias(ctx *kotlin.ImportAliasContext) {}

// ExitImportAlias is called when production importAlias is exited.
func (s *KotlinTreeShapeListener) ExitImportAlias(ctx *kotlin.ImportAliasContext) {}

// EnterTopLevelObject is called when production topLevelObject is entered.
func (s *KotlinTreeShapeListener) EnterTopLevelObject(ctx *kotlin.TopLevelObjectContext) {}

// ExitTopLevelObject is called when production topLevelObject is exited.
func (s *KotlinTreeShapeListener) ExitTopLevelObject(ctx *kotlin.TopLevelObjectContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitFunctionDeclaration(ctx *kotlin.FunctionDeclarationContext) {}

// EnterPrimaryConstructor is called when production primaryConstructor is entered.
func (s *KotlinTreeShapeListener) EnterPrimaryConstructor(ctx *kotlin.PrimaryConstructorContext) {}

// ExitPrimaryConstructor is called when production primaryConstructor is exited.
func (s *KotlinTreeShapeListener) ExitPrimaryConstructor(ctx *kotlin.PrimaryConstructorContext) {}

// EnterClassParameters is called when production classParameters is entered.
func (s *KotlinTreeShapeListener) EnterClassParameters(ctx *kotlin.ClassParametersContext) {}

// ExitClassParameters is called when production classParameters is exited.
func (s *KotlinTreeShapeListener) ExitClassParameters(ctx *kotlin.ClassParametersContext) {}

// EnterClassParameter is called when production classParameter is entered.
func (s *KotlinTreeShapeListener) EnterClassParameter(ctx *kotlin.ClassParameterContext) {}

// ExitClassParameter is called when production classParameter is exited.
func (s *KotlinTreeShapeListener) ExitClassParameter(ctx *kotlin.ClassParameterContext) {}

// EnterDelegationSpecifiers is called when production delegationSpecifiers is entered.
func (s *KotlinTreeShapeListener) EnterDelegationSpecifiers(ctx *kotlin.DelegationSpecifiersContext) {
}

// ExitDelegationSpecifiers is called when production delegationSpecifiers is exited.
func (s *KotlinTreeShapeListener) ExitDelegationSpecifiers(ctx *kotlin.DelegationSpecifiersContext) {}

// EnterDelegationSpecifier is called when production delegationSpecifier is entered.
func (s *KotlinTreeShapeListener) EnterDelegationSpecifier(ctx *kotlin.DelegationSpecifierContext) {}

// ExitDelegationSpecifier is called when production delegationSpecifier is exited.
func (s *KotlinTreeShapeListener) ExitDelegationSpecifier(ctx *kotlin.DelegationSpecifierContext) {}

// EnterConstructorInvocation is called when production constructorInvocation is entered.
func (s *KotlinTreeShapeListener) EnterConstructorInvocation(ctx *kotlin.ConstructorInvocationContext) {
}

// ExitConstructorInvocation is called when production constructorInvocation is exited.
func (s *KotlinTreeShapeListener) ExitConstructorInvocation(ctx *kotlin.ConstructorInvocationContext) {
}

// EnterExplicitDelegation is called when production explicitDelegation is entered.
func (s *KotlinTreeShapeListener) EnterExplicitDelegation(ctx *kotlin.ExplicitDelegationContext) {}

// ExitExplicitDelegation is called when production explicitDelegation is exited.
func (s *KotlinTreeShapeListener) ExitExplicitDelegation(ctx *kotlin.ExplicitDelegationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *KotlinTreeShapeListener) EnterClassBody(ctx *kotlin.ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *KotlinTreeShapeListener) ExitClassBody(ctx *kotlin.ClassBodyContext) {}

// EnterClassMemberDeclaration is called when production classMemberDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterClassMemberDeclaration(ctx *kotlin.ClassMemberDeclarationContext) {
}

// ExitClassMemberDeclaration is called when production classMemberDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitClassMemberDeclaration(ctx *kotlin.ClassMemberDeclarationContext) {
}

// EnterAnonymousInitializer is called when production anonymousInitializer is entered.
func (s *KotlinTreeShapeListener) EnterAnonymousInitializer(ctx *kotlin.AnonymousInitializerContext) {
}

// ExitAnonymousInitializer is called when production anonymousInitializer is exited.
func (s *KotlinTreeShapeListener) ExitAnonymousInitializer(ctx *kotlin.AnonymousInitializerContext) {}

// EnterSecondaryConstructor is called when production secondaryConstructor is entered.
func (s *KotlinTreeShapeListener) EnterSecondaryConstructor(ctx *kotlin.SecondaryConstructorContext) {
}

// ExitSecondaryConstructor is called when production secondaryConstructor is exited.
func (s *KotlinTreeShapeListener) ExitSecondaryConstructor(ctx *kotlin.SecondaryConstructorContext) {}

// EnterConstructorDelegationCall is called when production constructorDelegationCall is entered.
func (s *KotlinTreeShapeListener) EnterConstructorDelegationCall(ctx *kotlin.ConstructorDelegationCallContext) {
}

// ExitConstructorDelegationCall is called when production constructorDelegationCall is exited.
func (s *KotlinTreeShapeListener) ExitConstructorDelegationCall(ctx *kotlin.ConstructorDelegationCallContext) {
}

// EnterEnumClassBody is called when production enumClassBody is entered.
func (s *KotlinTreeShapeListener) EnterEnumClassBody(ctx *kotlin.EnumClassBodyContext) {}

// ExitEnumClassBody is called when production enumClassBody is exited.
func (s *KotlinTreeShapeListener) ExitEnumClassBody(ctx *kotlin.EnumClassBodyContext) {}

// EnterEnumEntries is called when production enumEntries is entered.
func (s *KotlinTreeShapeListener) EnterEnumEntries(ctx *kotlin.EnumEntriesContext) {}

// ExitEnumEntries is called when production enumEntries is exited.
func (s *KotlinTreeShapeListener) ExitEnumEntries(ctx *kotlin.EnumEntriesContext) {}

// EnterEnumEntry is called when production enumEntry is entered.
func (s *KotlinTreeShapeListener) EnterEnumEntry(ctx *kotlin.EnumEntryContext) {}

// ExitEnumEntry is called when production enumEntry is exited.
func (s *KotlinTreeShapeListener) ExitEnumEntry(ctx *kotlin.EnumEntryContext) {}

// EnterFunctionValueParameters is called when production functionValueParameters is entered.
func (s *KotlinTreeShapeListener) EnterFunctionValueParameters(ctx *kotlin.FunctionValueParametersContext) {
	//fmt.Printf("!!!EnterFunctionValueParameters:%s,counts:%d\n",ctx.GetText(),ctx.GetChildCount())
}

// ExitFunctionValueParameters is called when production functionValueParameters is exited.
func (s *KotlinTreeShapeListener) ExitFunctionValueParameters(ctx *kotlin.FunctionValueParametersContext) {
}

// ExitFunctionValueParameter is called when production functionValueParameter is exited.
func (s *KotlinTreeShapeListener) ExitFunctionValueParameter(ctx *kotlin.FunctionValueParameterContext) {
}

// EnterParameter is called when production parameter is entered.
func (s *KotlinTreeShapeListener) EnterParameter(ctx *kotlin.ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *KotlinTreeShapeListener) ExitParameter(ctx *kotlin.ParameterContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *KotlinTreeShapeListener) EnterFunctionBody(ctx *kotlin.FunctionBodyContext) {}

// EnterObjectDeclaration is called when production objectDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterObjectDeclaration(ctx *kotlin.ObjectDeclarationContext) {}

// ExitInfixFunctionCall is called when production infixFunctionCall is exited.
func (s *KotlinTreeShapeListener) ExitInfixFunctionCall(ctx *kotlin.InfixFunctionCallContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *KotlinTreeShapeListener) ExitFunctionBody(ctx *kotlin.FunctionBodyContext) {}

// ExitObjectDeclaration is called when production objectDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitObjectDeclaration(ctx *kotlin.ObjectDeclarationContext) {}

// EnterCompanionObject is called when production companionObject is entered.
func (s *KotlinTreeShapeListener) EnterCompanionObject(ctx *kotlin.CompanionObjectContext) {}

// ExitCompanionObject is called when production companionObject is exited.
func (s *KotlinTreeShapeListener) ExitCompanionObject(ctx *kotlin.CompanionObjectContext) {}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterPropertyDeclaration(ctx *kotlin.PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitPropertyDeclaration(ctx *kotlin.PropertyDeclarationContext) {}

// EnterMultiVariableDeclaration is called when production multiVariableDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterMultiVariableDeclaration(ctx *kotlin.MultiVariableDeclarationContext) {
}

// ExitMultiVariableDeclaration is called when production multiVariableDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitMultiVariableDeclaration(ctx *kotlin.MultiVariableDeclarationContext) {
}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterVariableDeclaration(ctx *kotlin.VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitVariableDeclaration(ctx *kotlin.VariableDeclarationContext) {}

// EnterGetter is called when production getter is entered.
func (s *KotlinTreeShapeListener) EnterGetter(ctx *kotlin.GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *KotlinTreeShapeListener) ExitGetter(ctx *kotlin.GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *KotlinTreeShapeListener) EnterSetter(ctx *kotlin.SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *KotlinTreeShapeListener) ExitSetter(ctx *kotlin.SetterContext) {}

// EnterTypeAlias is called when production typeAlias is entered.
func (s *KotlinTreeShapeListener) EnterTypeAlias(ctx *kotlin.TypeAliasContext) {}

// ExitTypeAlias is called when production typeAlias is exited.
func (s *KotlinTreeShapeListener) ExitTypeAlias(ctx *kotlin.TypeAliasContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *KotlinTreeShapeListener) EnterTypeParameters(ctx *kotlin.TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *KotlinTreeShapeListener) ExitTypeParameters(ctx *kotlin.TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *KotlinTreeShapeListener) EnterTypeParameter(ctx *kotlin.TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *KotlinTreeShapeListener) ExitTypeParameter(ctx *kotlin.TypeParameterContext) {}

// EnterType is called when production type is entered.
func (s *KotlinTreeShapeListener) EnterType(ctx *kotlin.TypeContext) {}

// ExitType is called when production type is exited.
func (s *KotlinTreeShapeListener) ExitType(ctx *kotlin.TypeContext) {}

// EnterTypeModifierList is called when production typeModifierList is entered.
func (s *KotlinTreeShapeListener) EnterTypeModifierList(ctx *kotlin.TypeModifierListContext) {}

// ExitTypeModifierList is called when production typeModifierList is exited.
func (s *KotlinTreeShapeListener) ExitTypeModifierList(ctx *kotlin.TypeModifierListContext) {}

// EnterParenthesizedType is called when production parenthesizedType is entered.
func (s *KotlinTreeShapeListener) EnterParenthesizedType(ctx *kotlin.ParenthesizedTypeContext) {}

// ExitParenthesizedType is called when production parenthesizedType is exited.
func (s *KotlinTreeShapeListener) ExitParenthesizedType(ctx *kotlin.ParenthesizedTypeContext) {}

// EnterNullableType is called when production nullableType is entered.
func (s *KotlinTreeShapeListener) EnterNullableType(ctx *kotlin.NullableTypeContext) {}

// ExitNullableType is called when production nullableType is exited.
func (s *KotlinTreeShapeListener) ExitNullableType(ctx *kotlin.NullableTypeContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *KotlinTreeShapeListener) EnterTypeReference(ctx *kotlin.TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *KotlinTreeShapeListener) ExitTypeReference(ctx *kotlin.TypeReferenceContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *KotlinTreeShapeListener) EnterFunctionType(ctx *kotlin.FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *KotlinTreeShapeListener) ExitFunctionType(ctx *kotlin.FunctionTypeContext) {}

// EnterFunctionTypeReceiver is called when production functionTypeReceiver is entered.
func (s *KotlinTreeShapeListener) EnterFunctionTypeReceiver(ctx *kotlin.FunctionTypeReceiverContext) {
}

// ExitFunctionTypeReceiver is called when production functionTypeReceiver is exited.
func (s *KotlinTreeShapeListener) ExitFunctionTypeReceiver(ctx *kotlin.FunctionTypeReceiverContext) {}

// EnterUserType is called when production userType is entered.
func (s *KotlinTreeShapeListener) EnterUserType(ctx *kotlin.UserTypeContext) {}

// ExitUserType is called when production userType is exited.
func (s *KotlinTreeShapeListener) ExitUserType(ctx *kotlin.UserTypeContext) {}

// EnterSimpleUserType is called when production simpleUserType is entered.
func (s *KotlinTreeShapeListener) EnterSimpleUserType(ctx *kotlin.SimpleUserTypeContext) {}

// ExitSimpleUserType is called when production simpleUserType is exited.
func (s *KotlinTreeShapeListener) ExitSimpleUserType(ctx *kotlin.SimpleUserTypeContext) {}

// EnterFunctionTypeParameters is called when production functionTypeParameters is entered.
func (s *KotlinTreeShapeListener) EnterFunctionTypeParameters(ctx *kotlin.FunctionTypeParametersContext) {
}

// ExitFunctionTypeParameters is called when production functionTypeParameters is exited.
func (s *KotlinTreeShapeListener) ExitFunctionTypeParameters(ctx *kotlin.FunctionTypeParametersContext) {
}

// EnterTypeConstraints is called when production typeConstraints is entered.
func (s *KotlinTreeShapeListener) EnterTypeConstraints(ctx *kotlin.TypeConstraintsContext) {}

// ExitTypeConstraints is called when production typeConstraints is exited.
func (s *KotlinTreeShapeListener) ExitTypeConstraints(ctx *kotlin.TypeConstraintsContext) {}

// EnterTypeConstraint is called when production typeConstraint is entered.
func (s *KotlinTreeShapeListener) EnterTypeConstraint(ctx *kotlin.TypeConstraintContext) {}

// ExitTypeConstraint is called when production typeConstraint is exited.
func (s *KotlinTreeShapeListener) ExitTypeConstraint(ctx *kotlin.TypeConstraintContext) {}

// EnterBlock is called when production block is entered.
func (s *KotlinTreeShapeListener) EnterBlock(ctx *kotlin.BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *KotlinTreeShapeListener) ExitBlock(ctx *kotlin.BlockContext) {}

// EnterStatements is called when production statements is entered.
func (s *KotlinTreeShapeListener) EnterStatements(ctx *kotlin.StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *KotlinTreeShapeListener) ExitStatements(ctx *kotlin.StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *KotlinTreeShapeListener) EnterStatement(ctx *kotlin.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *KotlinTreeShapeListener) ExitStatement(ctx *kotlin.StatementContext) {}

// EnterBlockLevelExpression is called when production blockLevelExpression is entered.
func (s *KotlinTreeShapeListener) EnterBlockLevelExpression(ctx *kotlin.BlockLevelExpressionContext) {
}

// ExitBlockLevelExpression is called when production blockLevelExpression is exited.
func (s *KotlinTreeShapeListener) ExitBlockLevelExpression(ctx *kotlin.BlockLevelExpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *KotlinTreeShapeListener) EnterDeclaration(ctx *kotlin.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *KotlinTreeShapeListener) ExitDeclaration(ctx *kotlin.DeclarationContext) {}

// EnterExpression is called when production expression is entered.
func (s *KotlinTreeShapeListener) EnterExpression(ctx *kotlin.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *KotlinTreeShapeListener) ExitExpression(ctx *kotlin.ExpressionContext) {}

// EnterDisjunction is called when production disjunction is entered.
func (s *KotlinTreeShapeListener) EnterDisjunction(ctx *kotlin.DisjunctionContext) {}

// ExitDisjunction is called when production disjunction is exited.
func (s *KotlinTreeShapeListener) ExitDisjunction(ctx *kotlin.DisjunctionContext) {}

// EnterConjunction is called when production conjunction is entered.
func (s *KotlinTreeShapeListener) EnterConjunction(ctx *kotlin.ConjunctionContext) {}

// ExitConjunction is called when production conjunction is exited.
func (s *KotlinTreeShapeListener) ExitConjunction(ctx *kotlin.ConjunctionContext) {}

// EnterEqualityComparison is called when production equalityComparison is entered.
func (s *KotlinTreeShapeListener) EnterEqualityComparison(ctx *kotlin.EqualityComparisonContext) {}

// ExitEqualityComparison is called when production equalityComparison is exited.
func (s *KotlinTreeShapeListener) ExitEqualityComparison(ctx *kotlin.EqualityComparisonContext) {}

// EnterComparison is called when production comparison is entered.
func (s *KotlinTreeShapeListener) EnterComparison(ctx *kotlin.ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *KotlinTreeShapeListener) ExitComparison(ctx *kotlin.ComparisonContext) {}

// EnterNamedInfix is called when production namedInfix is entered.
func (s *KotlinTreeShapeListener) EnterNamedInfix(ctx *kotlin.NamedInfixContext) {}

// ExitNamedInfix is called when production namedInfix is exited.
func (s *KotlinTreeShapeListener) ExitNamedInfix(ctx *kotlin.NamedInfixContext) {}

// EnterElvisExpression is called when production elvisExpression is entered.
func (s *KotlinTreeShapeListener) EnterElvisExpression(ctx *kotlin.ElvisExpressionContext) {}

// ExitElvisExpression is called when production elvisExpression is exited.
func (s *KotlinTreeShapeListener) ExitElvisExpression(ctx *kotlin.ElvisExpressionContext) {}

// EnterInfixFunctionCall is called when production infixFunctionCall is entered.
func (s *KotlinTreeShapeListener) EnterInfixFunctionCall(ctx *kotlin.InfixFunctionCallContext) {
}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *KotlinTreeShapeListener) EnterRangeExpression(ctx *kotlin.RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *KotlinTreeShapeListener) ExitRangeExpression(ctx *kotlin.RangeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *KotlinTreeShapeListener) EnterAdditiveExpression(ctx *kotlin.AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *KotlinTreeShapeListener) ExitAdditiveExpression(ctx *kotlin.AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *KotlinTreeShapeListener) EnterMultiplicativeExpression(ctx *kotlin.MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *KotlinTreeShapeListener) ExitMultiplicativeExpression(ctx *kotlin.MultiplicativeExpressionContext) {
}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitClassDeclaration(ctx *kotlin.ClassDeclarationContext) {}

// EnterTypeRHS is called when production typeRHS is entered.
func (s *KotlinTreeShapeListener) EnterTypeRHS(ctx *kotlin.TypeRHSContext) {}

// ExitTypeRHS is called when production typeRHS is exited.
func (s *KotlinTreeShapeListener) ExitTypeRHS(ctx *kotlin.TypeRHSContext) {}

// EnterPrefixUnaryExpression is called when production prefixUnaryExpression is entered.
func (s *KotlinTreeShapeListener) EnterPrefixUnaryExpression(ctx *kotlin.PrefixUnaryExpressionContext) {
}

// ExitPrefixUnaryExpression is called when production prefixUnaryExpression is exited.
func (s *KotlinTreeShapeListener) ExitPrefixUnaryExpression(ctx *kotlin.PrefixUnaryExpressionContext) {
}

// EnterPostfixUnaryExpression is called when production postfixUnaryExpression is entered.
func (s *KotlinTreeShapeListener) EnterPostfixUnaryExpression(ctx *kotlin.PostfixUnaryExpressionContext) {
}

// ExitPostfixUnaryExpression is called when production postfixUnaryExpression is exited.
func (s *KotlinTreeShapeListener) ExitPostfixUnaryExpression(ctx *kotlin.PostfixUnaryExpressionContext) {
}

// EnterAtomicExpression is called when production atomicExpression is entered.
func (s *KotlinTreeShapeListener) EnterAtomicExpression(ctx *kotlin.AtomicExpressionContext) {}

// ExitAtomicExpression is called when production atomicExpression is exited.
func (s *KotlinTreeShapeListener) ExitAtomicExpression(ctx *kotlin.AtomicExpressionContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *KotlinTreeShapeListener) EnterParenthesizedExpression(ctx *kotlin.ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *KotlinTreeShapeListener) ExitParenthesizedExpression(ctx *kotlin.ParenthesizedExpressionContext) {
}

// EnterCallSuffix is called when production callSuffix is entered.
func (s *KotlinTreeShapeListener) EnterCallSuffix(ctx *kotlin.CallSuffixContext) {}

// ExitCallSuffix is called when production callSuffix is exited.
func (s *KotlinTreeShapeListener) ExitCallSuffix(ctx *kotlin.CallSuffixContext) {}

// EnterAnnotatedLambda is called when production annotatedLambda is entered.
func (s *KotlinTreeShapeListener) EnterAnnotatedLambda(ctx *kotlin.AnnotatedLambdaContext) {}

// ExitAnnotatedLambda is called when production annotatedLambda is exited.
func (s *KotlinTreeShapeListener) ExitAnnotatedLambda(ctx *kotlin.AnnotatedLambdaContext) {}

// EnterArrayAccess is called when production arrayAccess is entered.
func (s *KotlinTreeShapeListener) EnterArrayAccess(ctx *kotlin.ArrayAccessContext) {}

// ExitArrayAccess is called when production arrayAccess is exited.
func (s *KotlinTreeShapeListener) ExitArrayAccess(ctx *kotlin.ArrayAccessContext) {}

// EnterValueArguments is called when production valueArguments is entered.
func (s *KotlinTreeShapeListener) EnterValueArguments(ctx *kotlin.ValueArgumentsContext) {}

// ExitValueArguments is called when production valueArguments is exited.
func (s *KotlinTreeShapeListener) ExitValueArguments(ctx *kotlin.ValueArgumentsContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *KotlinTreeShapeListener) EnterTypeArguments(ctx *kotlin.TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *KotlinTreeShapeListener) ExitTypeArguments(ctx *kotlin.TypeArgumentsContext) {}

// EnterTypeProjection is called when production typeProjection is entered.
func (s *KotlinTreeShapeListener) EnterTypeProjection(ctx *kotlin.TypeProjectionContext) {}

// ExitTypeProjection is called when production typeProjection is exited.
func (s *KotlinTreeShapeListener) ExitTypeProjection(ctx *kotlin.TypeProjectionContext) {}

// EnterTypeProjectionModifierList is called when production typeProjectionModifierList is entered.
func (s *KotlinTreeShapeListener) EnterTypeProjectionModifierList(ctx *kotlin.TypeProjectionModifierListContext) {
}

// ExitTypeProjectionModifierList is called when production typeProjectionModifierList is exited.
func (s *KotlinTreeShapeListener) ExitTypeProjectionModifierList(ctx *kotlin.TypeProjectionModifierListContext) {
}

// EnterValueArgument is called when production valueArgument is entered.
func (s *KotlinTreeShapeListener) EnterValueArgument(ctx *kotlin.ValueArgumentContext) {}

// ExitValueArgument is called when production valueArgument is exited.
func (s *KotlinTreeShapeListener) ExitValueArgument(ctx *kotlin.ValueArgumentContext) {}

// EnterLiteralConstant is called when production literalConstant is entered.
func (s *KotlinTreeShapeListener) EnterLiteralConstant(ctx *kotlin.LiteralConstantContext) {}

// ExitLiteralConstant is called when production literalConstant is exited.
func (s *KotlinTreeShapeListener) ExitLiteralConstant(ctx *kotlin.LiteralConstantContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *KotlinTreeShapeListener) EnterStringLiteral(ctx *kotlin.StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *KotlinTreeShapeListener) ExitStringLiteral(ctx *kotlin.StringLiteralContext) {}

// EnterLineStringLiteral is called when production lineStringLiteral is entered.
func (s *KotlinTreeShapeListener) EnterLineStringLiteral(ctx *kotlin.LineStringLiteralContext) {}

// ExitLineStringLiteral is called when production lineStringLiteral is exited.
func (s *KotlinTreeShapeListener) ExitLineStringLiteral(ctx *kotlin.LineStringLiteralContext) {}

// EnterMultiLineStringLiteral is called when production multiLineStringLiteral is entered.
func (s *KotlinTreeShapeListener) EnterMultiLineStringLiteral(ctx *kotlin.MultiLineStringLiteralContext) {
}

// ExitMultiLineStringLiteral is called when production multiLineStringLiteral is exited.
func (s *KotlinTreeShapeListener) ExitMultiLineStringLiteral(ctx *kotlin.MultiLineStringLiteralContext) {
}

// EnterLineStringContent is called when production lineStringContent is entered.
func (s *KotlinTreeShapeListener) EnterLineStringContent(ctx *kotlin.LineStringContentContext) {}

// ExitLineStringContent is called when production lineStringContent is exited.
func (s *KotlinTreeShapeListener) ExitLineStringContent(ctx *kotlin.LineStringContentContext) {}

// EnterLineStringExpression is called when production lineStringExpression is entered.
func (s *KotlinTreeShapeListener) EnterLineStringExpression(ctx *kotlin.LineStringExpressionContext) {
}

// ExitLineStringExpression is called when production lineStringExpression is exited.
func (s *KotlinTreeShapeListener) ExitLineStringExpression(ctx *kotlin.LineStringExpressionContext) {}

// EnterMultiLineStringContent is called when production multiLineStringContent is entered.
func (s *KotlinTreeShapeListener) EnterMultiLineStringContent(ctx *kotlin.MultiLineStringContentContext) {
}

// ExitMultiLineStringContent is called when production multiLineStringContent is exited.
func (s *KotlinTreeShapeListener) ExitMultiLineStringContent(ctx *kotlin.MultiLineStringContentContext) {
}

// EnterMultiLineStringExpression is called when production multiLineStringExpression is entered.
func (s *KotlinTreeShapeListener) EnterMultiLineStringExpression(ctx *kotlin.MultiLineStringExpressionContext) {
}

// ExitMultiLineStringExpression is called when production multiLineStringExpression is exited.
func (s *KotlinTreeShapeListener) ExitMultiLineStringExpression(ctx *kotlin.MultiLineStringExpressionContext) {
}

// EnterFunctionLiteral is called when production functionLiteral is entered.
func (s *KotlinTreeShapeListener) EnterFunctionLiteral(ctx *kotlin.FunctionLiteralContext) {}

// ExitFunctionLiteral is called when production functionLiteral is exited.
func (s *KotlinTreeShapeListener) ExitFunctionLiteral(ctx *kotlin.FunctionLiteralContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *KotlinTreeShapeListener) EnterLambdaParameters(ctx *kotlin.LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *KotlinTreeShapeListener) ExitLambdaParameters(ctx *kotlin.LambdaParametersContext) {}

// EnterLambdaParameter is called when production lambdaParameter is entered.
func (s *KotlinTreeShapeListener) EnterLambdaParameter(ctx *kotlin.LambdaParameterContext) {}

// ExitLambdaParameter is called when production lambdaParameter is exited.
func (s *KotlinTreeShapeListener) ExitLambdaParameter(ctx *kotlin.LambdaParameterContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *KotlinTreeShapeListener) EnterObjectLiteral(ctx *kotlin.ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *KotlinTreeShapeListener) ExitObjectLiteral(ctx *kotlin.ObjectLiteralContext) {}

// EnterCollectionLiteral is called when production collectionLiteral is entered.
func (s *KotlinTreeShapeListener) EnterCollectionLiteral(ctx *kotlin.CollectionLiteralContext) {}

// ExitCollectionLiteral is called when production collectionLiteral is exited.
func (s *KotlinTreeShapeListener) ExitCollectionLiteral(ctx *kotlin.CollectionLiteralContext) {}

// EnterThisExpression is called when production thisExpression is entered.
func (s *KotlinTreeShapeListener) EnterThisExpression(ctx *kotlin.ThisExpressionContext) {}

// ExitThisExpression is called when production thisExpression is exited.
func (s *KotlinTreeShapeListener) ExitThisExpression(ctx *kotlin.ThisExpressionContext) {}

// EnterSuperExpression is called when production superExpression is entered.
func (s *KotlinTreeShapeListener) EnterSuperExpression(ctx *kotlin.SuperExpressionContext) {}

// ExitSuperExpression is called when production superExpression is exited.
func (s *KotlinTreeShapeListener) ExitSuperExpression(ctx *kotlin.SuperExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *KotlinTreeShapeListener) EnterConditionalExpression(ctx *kotlin.ConditionalExpressionContext) {
}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *KotlinTreeShapeListener) ExitConditionalExpression(ctx *kotlin.ConditionalExpressionContext) {
}

// EnterIfExpression is called when production ifExpression is entered.
func (s *KotlinTreeShapeListener) EnterIfExpression(ctx *kotlin.IfExpressionContext) {}

// ExitIfExpression is called when production ifExpression is exited.
func (s *KotlinTreeShapeListener) ExitIfExpression(ctx *kotlin.IfExpressionContext) {}

// EnterControlStructureBody is called when production controlStructureBody is entered.
func (s *KotlinTreeShapeListener) EnterControlStructureBody(ctx *kotlin.ControlStructureBodyContext) {
}

// ExitControlStructureBody is called when production controlStructureBody is exited.
func (s *KotlinTreeShapeListener) ExitControlStructureBody(ctx *kotlin.ControlStructureBodyContext) {}

// EnterWhenExpression is called when production whenExpression is entered.
func (s *KotlinTreeShapeListener) EnterWhenExpression(ctx *kotlin.WhenExpressionContext) {}

// ExitWhenExpression is called when production whenExpression is exited.
func (s *KotlinTreeShapeListener) ExitWhenExpression(ctx *kotlin.WhenExpressionContext) {}

// EnterWhenEntry is called when production whenEntry is entered.
func (s *KotlinTreeShapeListener) EnterWhenEntry(ctx *kotlin.WhenEntryContext) {}

// ExitWhenEntry is called when production whenEntry is exited.
func (s *KotlinTreeShapeListener) ExitWhenEntry(ctx *kotlin.WhenEntryContext) {}

// EnterWhenCondition is called when production whenCondition is entered.
func (s *KotlinTreeShapeListener) EnterWhenCondition(ctx *kotlin.WhenConditionContext) {}

// ExitWhenCondition is called when production whenCondition is exited.
func (s *KotlinTreeShapeListener) ExitWhenCondition(ctx *kotlin.WhenConditionContext) {}

// EnterRangeTest is called when production rangeTest is entered.
func (s *KotlinTreeShapeListener) EnterRangeTest(ctx *kotlin.RangeTestContext) {}

// ExitRangeTest is called when production rangeTest is exited.
func (s *KotlinTreeShapeListener) ExitRangeTest(ctx *kotlin.RangeTestContext) {}

// EnterTypeTest is called when production typeTest is entered.
func (s *KotlinTreeShapeListener) EnterTypeTest(ctx *kotlin.TypeTestContext) {}

// ExitTypeTest is called when production typeTest is exited.
func (s *KotlinTreeShapeListener) ExitTypeTest(ctx *kotlin.TypeTestContext) {}

// EnterTryExpression is called when production tryExpression is entered.
func (s *KotlinTreeShapeListener) EnterTryExpression(ctx *kotlin.TryExpressionContext) {}

// ExitTryExpression is called when production tryExpression is exited.
func (s *KotlinTreeShapeListener) ExitTryExpression(ctx *kotlin.TryExpressionContext) {}

// EnterCatchBlock is called when production catchBlock is entered.
func (s *KotlinTreeShapeListener) EnterCatchBlock(ctx *kotlin.CatchBlockContext) {}

// ExitCatchBlock is called when production catchBlock is exited.
func (s *KotlinTreeShapeListener) ExitCatchBlock(ctx *kotlin.CatchBlockContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *KotlinTreeShapeListener) EnterFinallyBlock(ctx *kotlin.FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *KotlinTreeShapeListener) ExitFinallyBlock(ctx *kotlin.FinallyBlockContext) {}

// EnterLoopExpression is called when production loopExpression is entered.
func (s *KotlinTreeShapeListener) EnterLoopExpression(ctx *kotlin.LoopExpressionContext) {}

// ExitLoopExpression is called when production loopExpression is exited.
func (s *KotlinTreeShapeListener) ExitLoopExpression(ctx *kotlin.LoopExpressionContext) {}

// EnterForExpression is called when production forExpression is entered.
func (s *KotlinTreeShapeListener) EnterForExpression(ctx *kotlin.ForExpressionContext) {}

// ExitForExpression is called when production forExpression is exited.
func (s *KotlinTreeShapeListener) ExitForExpression(ctx *kotlin.ForExpressionContext) {}

// EnterWhileExpression is called when production whileExpression is entered.
func (s *KotlinTreeShapeListener) EnterWhileExpression(ctx *kotlin.WhileExpressionContext) {}

// ExitWhileExpression is called when production whileExpression is exited.
func (s *KotlinTreeShapeListener) ExitWhileExpression(ctx *kotlin.WhileExpressionContext) {}

// EnterDoWhileExpression is called when production doWhileExpression is entered.
func (s *KotlinTreeShapeListener) EnterDoWhileExpression(ctx *kotlin.DoWhileExpressionContext) {}

// ExitDoWhileExpression is called when production doWhileExpression is exited.
func (s *KotlinTreeShapeListener) ExitDoWhileExpression(ctx *kotlin.DoWhileExpressionContext) {}

// EnterJumpExpression is called when production jumpExpression is entered.
func (s *KotlinTreeShapeListener) EnterJumpExpression(ctx *kotlin.JumpExpressionContext) {}

// ExitJumpExpression is called when production jumpExpression is exited.
func (s *KotlinTreeShapeListener) ExitJumpExpression(ctx *kotlin.JumpExpressionContext) {}

// EnterCallableReference is called when production callableReference is entered.
func (s *KotlinTreeShapeListener) EnterCallableReference(ctx *kotlin.CallableReferenceContext) {}

// ExitCallableReference is called when production callableReference is exited.
func (s *KotlinTreeShapeListener) ExitCallableReference(ctx *kotlin.CallableReferenceContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *KotlinTreeShapeListener) EnterAssignmentOperator(ctx *kotlin.AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *KotlinTreeShapeListener) ExitAssignmentOperator(ctx *kotlin.AssignmentOperatorContext) {}

// EnterEqualityOperation is called when production equalityOperation is entered.
func (s *KotlinTreeShapeListener) EnterEqualityOperation(ctx *kotlin.EqualityOperationContext) {}

// ExitEqualityOperation is called when production equalityOperation is exited.
func (s *KotlinTreeShapeListener) ExitEqualityOperation(ctx *kotlin.EqualityOperationContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *KotlinTreeShapeListener) EnterComparisonOperator(ctx *kotlin.ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *KotlinTreeShapeListener) ExitComparisonOperator(ctx *kotlin.ComparisonOperatorContext) {}

// EnterInOperator is called when production inOperator is entered.
func (s *KotlinTreeShapeListener) EnterInOperator(ctx *kotlin.InOperatorContext) {}

// ExitInOperator is called when production inOperator is exited.
func (s *KotlinTreeShapeListener) ExitInOperator(ctx *kotlin.InOperatorContext) {}

// EnterIsOperator is called when production isOperator is entered.
func (s *KotlinTreeShapeListener) EnterIsOperator(ctx *kotlin.IsOperatorContext) {}

// ExitIsOperator is called when production isOperator is exited.
func (s *KotlinTreeShapeListener) ExitIsOperator(ctx *kotlin.IsOperatorContext) {}

// EnterAdditiveOperator is called when production additiveOperator is entered.
func (s *KotlinTreeShapeListener) EnterAdditiveOperator(ctx *kotlin.AdditiveOperatorContext) {}

// ExitAdditiveOperator is called when production additiveOperator is exited.
func (s *KotlinTreeShapeListener) ExitAdditiveOperator(ctx *kotlin.AdditiveOperatorContext) {}

// EnterMultiplicativeOperation is called when production multiplicativeOperation is entered.
func (s *KotlinTreeShapeListener) EnterMultiplicativeOperation(ctx *kotlin.MultiplicativeOperationContext) {
}

// ExitMultiplicativeOperation is called when production multiplicativeOperation is exited.
func (s *KotlinTreeShapeListener) ExitMultiplicativeOperation(ctx *kotlin.MultiplicativeOperationContext) {
}

// EnterTypeOperation is called when production typeOperation is entered.
func (s *KotlinTreeShapeListener) EnterTypeOperation(ctx *kotlin.TypeOperationContext) {}

// ExitTypeOperation is called when production typeOperation is exited.
func (s *KotlinTreeShapeListener) ExitTypeOperation(ctx *kotlin.TypeOperationContext) {}

// EnterPrefixUnaryOperation is called when production prefixUnaryOperation is entered.
func (s *KotlinTreeShapeListener) EnterPrefixUnaryOperation(ctx *kotlin.PrefixUnaryOperationContext) {
}

// ExitPrefixUnaryOperation is called when production prefixUnaryOperation is exited.
func (s *KotlinTreeShapeListener) ExitPrefixUnaryOperation(ctx *kotlin.PrefixUnaryOperationContext) {
}

// EnterPostfixUnaryOperation is called when production postfixUnaryOperation is entered.
func (s *KotlinTreeShapeListener) EnterPostfixUnaryOperation(ctx *kotlin.PostfixUnaryOperationContext) {
}

// EnterMemberAccessOperator is called when production memberAccessOperator is entered.
func (s *KotlinTreeShapeListener) EnterMemberAccessOperator(ctx *kotlin.MemberAccessOperatorContext) {
}

// ExitMemberAccessOperator is called when production memberAccessOperator is exited.
func (s *KotlinTreeShapeListener) ExitMemberAccessOperator(ctx *kotlin.MemberAccessOperatorContext) {}

// EnterModifierList is called when production modifierList is entered.
func (s *KotlinTreeShapeListener) EnterModifierList(ctx *kotlin.ModifierListContext) {}

// ExitModifierList is called when production modifierList is exited.
func (s *KotlinTreeShapeListener) ExitModifierList(ctx *kotlin.ModifierListContext) {}

// EnterModifier is called when production modifier is entered.
func (s *KotlinTreeShapeListener) EnterModifier(ctx *kotlin.ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *KotlinTreeShapeListener) ExitModifier(ctx *kotlin.ModifierContext) {}

// EnterClassModifier is called when production classModifier is entered.
func (s *KotlinTreeShapeListener) EnterClassModifier(ctx *kotlin.ClassModifierContext) {}

// ExitClassModifier is called when production classModifier is exited.
func (s *KotlinTreeShapeListener) ExitClassModifier(ctx *kotlin.ClassModifierContext) {}

// EnterMemberModifier is called when production memberModifier is entered.
func (s *KotlinTreeShapeListener) EnterMemberModifier(ctx *kotlin.MemberModifierContext) {}

// ExitMemberModifier is called when production memberModifier is exited.
func (s *KotlinTreeShapeListener) ExitMemberModifier(ctx *kotlin.MemberModifierContext) {}

// EnterVisibilityModifier is called when production visibilityModifier is entered.
func (s *KotlinTreeShapeListener) EnterVisibilityModifier(ctx *kotlin.VisibilityModifierContext) {}

// ExitVisibilityModifier is called when production visibilityModifier is exited.
func (s *KotlinTreeShapeListener) ExitVisibilityModifier(ctx *kotlin.VisibilityModifierContext) {}

// EnterVarianceAnnotation is called when production varianceAnnotation is entered.
func (s *KotlinTreeShapeListener) EnterVarianceAnnotation(ctx *kotlin.VarianceAnnotationContext) {}

// ExitVarianceAnnotation is called when production varianceAnnotation is exited.
func (s *KotlinTreeShapeListener) ExitVarianceAnnotation(ctx *kotlin.VarianceAnnotationContext) {}

// EnterFunctionModifier is called when production functionModifier is entered.
func (s *KotlinTreeShapeListener) EnterFunctionModifier(ctx *kotlin.FunctionModifierContext) {}

// ExitFunctionModifier is called when production functionModifier is exited.
func (s *KotlinTreeShapeListener) ExitFunctionModifier(ctx *kotlin.FunctionModifierContext) {}

// EnterPropertyModifier is called when production propertyModifier is entered.
func (s *KotlinTreeShapeListener) EnterPropertyModifier(ctx *kotlin.PropertyModifierContext) {}

// ExitPropertyModifier is called when production propertyModifier is exited.
func (s *KotlinTreeShapeListener) ExitPropertyModifier(ctx *kotlin.PropertyModifierContext) {}

// EnterInheritanceModifier is called when production inheritanceModifier is entered.
func (s *KotlinTreeShapeListener) EnterInheritanceModifier(ctx *kotlin.InheritanceModifierContext) {}

// ExitInheritanceModifier is called when production inheritanceModifier is exited.
func (s *KotlinTreeShapeListener) ExitInheritanceModifier(ctx *kotlin.InheritanceModifierContext) {}

// EnterParameterModifier is called when production parameterModifier is entered.
func (s *KotlinTreeShapeListener) EnterParameterModifier(ctx *kotlin.ParameterModifierContext) {}

// ExitParameterModifier is called when production parameterModifier is exited.
func (s *KotlinTreeShapeListener) ExitParameterModifier(ctx *kotlin.ParameterModifierContext) {}

// EnterTypeParameterModifier is called when production typeParameterModifier is entered.
func (s *KotlinTreeShapeListener) EnterTypeParameterModifier(ctx *kotlin.TypeParameterModifierContext) {
}

// ExitTypeParameterModifier is called when production typeParameterModifier is exited.
func (s *KotlinTreeShapeListener) ExitTypeParameterModifier(ctx *kotlin.TypeParameterModifierContext) {
}

// EnterLabelDefinition is called when production labelDefinition is entered.
func (s *KotlinTreeShapeListener) EnterLabelDefinition(ctx *kotlin.LabelDefinitionContext) {}

// ExitLabelDefinition is called when production labelDefinition is exited.
func (s *KotlinTreeShapeListener) ExitLabelDefinition(ctx *kotlin.LabelDefinitionContext) {}

// EnterAnnotations is called when production annotations is entered.
func (s *KotlinTreeShapeListener) EnterAnnotations(ctx *kotlin.AnnotationsContext) {}

// ExitAnnotations is called when production annotations is exited.
func (s *KotlinTreeShapeListener) ExitAnnotations(ctx *kotlin.AnnotationsContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *KotlinTreeShapeListener) EnterAnnotation(ctx *kotlin.AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *KotlinTreeShapeListener) ExitAnnotation(ctx *kotlin.AnnotationContext) {}

// EnterAnnotationList is called when production annotationList is entered.
func (s *KotlinTreeShapeListener) EnterAnnotationList(ctx *kotlin.AnnotationListContext) {}

// ExitAnnotationList is called when production annotationList is exited.
func (s *KotlinTreeShapeListener) ExitAnnotationList(ctx *kotlin.AnnotationListContext) {}

// EnterAnnotationUseSiteTarget is called when production annotationUseSiteTarget is entered.
func (s *KotlinTreeShapeListener) EnterAnnotationUseSiteTarget(ctx *kotlin.AnnotationUseSiteTargetContext) {
}

// ExitAnnotationUseSiteTarget is called when production annotationUseSiteTarget is exited.
func (s *KotlinTreeShapeListener) ExitAnnotationUseSiteTarget(ctx *kotlin.AnnotationUseSiteTargetContext) {
}

// EnterUnescapedAnnotation is called when production unescapedAnnotation is entered.
func (s *KotlinTreeShapeListener) EnterUnescapedAnnotation(ctx *kotlin.UnescapedAnnotationContext) {}

// ExitUnescapedAnnotation is called when production unescapedAnnotation is exited.
func (s *KotlinTreeShapeListener) ExitUnescapedAnnotation(ctx *kotlin.UnescapedAnnotationContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *KotlinTreeShapeListener) EnterIdentifier(ctx *kotlin.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *KotlinTreeShapeListener) ExitIdentifier(ctx *kotlin.IdentifierContext) {}

// EnterSimpleIdentifier is called when production simpleIdentifier is entered.
func (s *KotlinTreeShapeListener) EnterSimpleIdentifier(ctx *kotlin.SimpleIdentifierContext) {}

// ExitSimpleIdentifier is called when production simpleIdentifier is exited.
func (s *KotlinTreeShapeListener) ExitSimpleIdentifier(ctx *kotlin.SimpleIdentifierContext) {}

// EnterSemi is called when production semi is entered.
func (s *KotlinTreeShapeListener) EnterSemi(ctx *kotlin.SemiContext) {}

// ExitSemi is called when production semi is exited.
func (s *KotlinTreeShapeListener) ExitSemi(ctx *kotlin.SemiContext) {}

// EnterAnysemi is called when production anysemi is entered.
func (s *KotlinTreeShapeListener) EnterAnysemi(ctx *kotlin.AnysemiContext) {}

// ExitAnysemi is called when production anysemi is exited.
func (s *KotlinTreeShapeListener) ExitAnysemi(ctx *kotlin.AnysemiContext) {}

func (s *KotlinTreeShapeListener) EnterFunctionValueParameter(c *kotlin.FunctionValueParameterContext) {
}
