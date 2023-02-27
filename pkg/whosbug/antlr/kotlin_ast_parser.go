// Generated from KotlinParser.g4 by ANTLR 4.7.

package antlr // KotlinParser

import (
	"sync"

	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr/kotlin"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// KotlinAstParser implement AstParser for kotlin language
//
//	@author kevineluo
//	@update 2023-02-28 01:23:41
type KotlinAstParser struct {
	astInfo AstInfo
}

var _ kotlin.KotlinParserListener = &KotlinAstParser{}

var (
	kotlinLexerPool = &sync.Pool{New: func() any {
		return kotlin.NewKotlinLexer(nil)
	}}
	kotlinParserPool = &sync.Pool{New: func() any {
		return kotlin.NewKotlinParser(nil)
	}}
	newKotlinAstParserPool = &sync.Pool{New: func() any {
		return new(KotlinAstParser)
	}}
)

// AstParse main parse process for kotlin language
//
//	@receiver s *KotlinAstParser
//	@param input string
//	@return AstInfo
//	@author kevineluo
//	@update 2023-02-28 01:28:03
func (s *KotlinAstParser) AstParse(input string) AstInfo {
	//	截取目标文本的输入流
	inputStream := antlr.NewInputStream(input)
	//	初始化lexer
	lexer := kotlinLexerPool.Get().(*kotlin.KotlinLexer)
	defer kotlinLexerPool.Put(lexer)
	lexer.SetInputStream(inputStream)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := kotlinParserPool.Get().(*kotlin.KotlinParser)
	p.RemoveErrorListeners()
	defer kotlinParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	//	启用SLL两阶段加速解析模式
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.KotlinFile()
	//	创建listener
	listener := newKotlinAstParserPool.Get().(*KotlinAstParser)
	defer newKotlinAstParserPool.Put(listener)
	// 初始化置空
	listener.astInfo = AstInfo{}
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.astInfo
}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *KotlinAstParser) EnterFunctionDeclaration(ctx *kotlin.FunctionDeclarationContext) {
	if ctx.Identifier() == nil {
		return
	}
	methodInfo := Method{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		Name:       findKotlinDeclarationChain(ctx) + ctx.Identifier().GetText(),
		Parameters: "",
	}
	s.astInfo.Methods = append(s.astInfo.Methods, methodInfo)
}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *KotlinAstParser) EnterClassDeclaration(ctx *kotlin.ClassDeclarationContext) {
	classInfo := Class{
		StartLine: ctx.GetStart().GetLine(),
		EndLine:   ctx.GetStop().GetLine(),
		Name:      findKotlinDeclarationChain(ctx) + ctx.SimpleIdentifier().GetText(),
		Extends:   findKotlinClassExtends(ctx),
	}
	s.astInfo.Classes = append(s.astInfo.Classes, classInfo)
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
func (s *KotlinAstParser) ExitPostfixUnaryOperation(ctx *kotlin.PostfixUnaryOperationContext) {
}

// VisitTerminal is called when a terminal node is visited.
func (s *KotlinAstParser) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *KotlinAstParser) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *KotlinAstParser) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *KotlinAstParser) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKotlinFile is called when production kotlinFile is entered.
func (s *KotlinAstParser) EnterKotlinFile(ctx *kotlin.KotlinFileContext) {}

// ExitKotlinFile is called when production kotlinFile is exited.
func (s *KotlinAstParser) ExitKotlinFile(ctx *kotlin.KotlinFileContext) {}

// EnterScript is called when production script is entered.
func (s *KotlinAstParser) EnterScript(ctx *kotlin.ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *KotlinAstParser) ExitScript(ctx *kotlin.ScriptContext) {}

// EnterPreamble is called when production preamble is entered.
func (s *KotlinAstParser) EnterPreamble(ctx *kotlin.PreambleContext) {}

// ExitPreamble is called when production preamble is exited.
func (s *KotlinAstParser) ExitPreamble(ctx *kotlin.PreambleContext) {}

// EnterFileAnnotations is called when production fileAnnotations is entered.
func (s *KotlinAstParser) EnterFileAnnotations(ctx *kotlin.FileAnnotationsContext) {}

// ExitFileAnnotations is called when production fileAnnotations is exited.
func (s *KotlinAstParser) ExitFileAnnotations(ctx *kotlin.FileAnnotationsContext) {}

// EnterFileAnnotation is called when production fileAnnotation is entered.
func (s *KotlinAstParser) EnterFileAnnotation(ctx *kotlin.FileAnnotationContext) {}

// ExitFileAnnotation is called when production fileAnnotation is exited.
func (s *KotlinAstParser) ExitFileAnnotation(ctx *kotlin.FileAnnotationContext) {}

// EnterPackageHeader is called when production packageHeader is entered.
func (s *KotlinAstParser) EnterPackageHeader(ctx *kotlin.PackageHeaderContext) {}

// ExitPackageHeader is called when production packageHeader is exited.
func (s *KotlinAstParser) ExitPackageHeader(ctx *kotlin.PackageHeaderContext) {}

// EnterImportList is called when production importList is entered.
func (s *KotlinAstParser) EnterImportList(ctx *kotlin.ImportListContext) {}

// ExitImportList is called when production importList is exited.
func (s *KotlinAstParser) ExitImportList(ctx *kotlin.ImportListContext) {}

// EnterImportHeader is called when production importHeader is entered.
func (s *KotlinAstParser) EnterImportHeader(ctx *kotlin.ImportHeaderContext) {}

// ExitImportHeader is called when production importHeader is exited.
func (s *KotlinAstParser) ExitImportHeader(ctx *kotlin.ImportHeaderContext) {}

// EnterImportAlias is called when production importAlias is entered.
func (s *KotlinAstParser) EnterImportAlias(ctx *kotlin.ImportAliasContext) {}

// ExitImportAlias is called when production importAlias is exited.
func (s *KotlinAstParser) ExitImportAlias(ctx *kotlin.ImportAliasContext) {}

// EnterTopLevelObject is called when production topLevelObject is entered.
func (s *KotlinAstParser) EnterTopLevelObject(ctx *kotlin.TopLevelObjectContext) {}

// ExitTopLevelObject is called when production topLevelObject is exited.
func (s *KotlinAstParser) ExitTopLevelObject(ctx *kotlin.TopLevelObjectContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *KotlinAstParser) ExitFunctionDeclaration(ctx *kotlin.FunctionDeclarationContext) {}

// EnterPrimaryConstructor is called when production primaryConstructor is entered.
func (s *KotlinAstParser) EnterPrimaryConstructor(ctx *kotlin.PrimaryConstructorContext) {}

// ExitPrimaryConstructor is called when production primaryConstructor is exited.
func (s *KotlinAstParser) ExitPrimaryConstructor(ctx *kotlin.PrimaryConstructorContext) {}

// EnterClassParameters is called when production classParameters is entered.
func (s *KotlinAstParser) EnterClassParameters(ctx *kotlin.ClassParametersContext) {}

// ExitClassParameters is called when production classParameters is exited.
func (s *KotlinAstParser) ExitClassParameters(ctx *kotlin.ClassParametersContext) {}

// EnterClassParameter is called when production classParameter is entered.
func (s *KotlinAstParser) EnterClassParameter(ctx *kotlin.ClassParameterContext) {}

// ExitClassParameter is called when production classParameter is exited.
func (s *KotlinAstParser) ExitClassParameter(ctx *kotlin.ClassParameterContext) {}

// EnterDelegationSpecifiers is called when production delegationSpecifiers is entered.
func (s *KotlinAstParser) EnterDelegationSpecifiers(ctx *kotlin.DelegationSpecifiersContext) {
}

// ExitDelegationSpecifiers is called when production delegationSpecifiers is exited.
func (s *KotlinAstParser) ExitDelegationSpecifiers(ctx *kotlin.DelegationSpecifiersContext) {}

// EnterDelegationSpecifier is called when production delegationSpecifier is entered.
func (s *KotlinAstParser) EnterDelegationSpecifier(ctx *kotlin.DelegationSpecifierContext) {}

// ExitDelegationSpecifier is called when production delegationSpecifier is exited.
func (s *KotlinAstParser) ExitDelegationSpecifier(ctx *kotlin.DelegationSpecifierContext) {}

// EnterConstructorInvocation is called when production constructorInvocation is entered.
func (s *KotlinAstParser) EnterConstructorInvocation(ctx *kotlin.ConstructorInvocationContext) {
}

// ExitConstructorInvocation is called when production constructorInvocation is exited.
func (s *KotlinAstParser) ExitConstructorInvocation(ctx *kotlin.ConstructorInvocationContext) {
}

// EnterExplicitDelegation is called when production explicitDelegation is entered.
func (s *KotlinAstParser) EnterExplicitDelegation(ctx *kotlin.ExplicitDelegationContext) {}

// ExitExplicitDelegation is called when production explicitDelegation is exited.
func (s *KotlinAstParser) ExitExplicitDelegation(ctx *kotlin.ExplicitDelegationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *KotlinAstParser) EnterClassBody(ctx *kotlin.ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *KotlinAstParser) ExitClassBody(ctx *kotlin.ClassBodyContext) {}

// EnterClassMemberDeclaration is called when production classMemberDeclaration is entered.
func (s *KotlinAstParser) EnterClassMemberDeclaration(ctx *kotlin.ClassMemberDeclarationContext) {
}

// ExitClassMemberDeclaration is called when production classMemberDeclaration is exited.
func (s *KotlinAstParser) ExitClassMemberDeclaration(ctx *kotlin.ClassMemberDeclarationContext) {
}

// EnterAnonymousInitializer is called when production anonymousInitializer is entered.
func (s *KotlinAstParser) EnterAnonymousInitializer(ctx *kotlin.AnonymousInitializerContext) {
}

// ExitAnonymousInitializer is called when production anonymousInitializer is exited.
func (s *KotlinAstParser) ExitAnonymousInitializer(ctx *kotlin.AnonymousInitializerContext) {}

// EnterSecondaryConstructor is called when production secondaryConstructor is entered.
func (s *KotlinAstParser) EnterSecondaryConstructor(ctx *kotlin.SecondaryConstructorContext) {
}

// ExitSecondaryConstructor is called when production secondaryConstructor is exited.
func (s *KotlinAstParser) ExitSecondaryConstructor(ctx *kotlin.SecondaryConstructorContext) {}

// EnterConstructorDelegationCall is called when production constructorDelegationCall is entered.
func (s *KotlinAstParser) EnterConstructorDelegationCall(ctx *kotlin.ConstructorDelegationCallContext) {
}

// ExitConstructorDelegationCall is called when production constructorDelegationCall is exited.
func (s *KotlinAstParser) ExitConstructorDelegationCall(ctx *kotlin.ConstructorDelegationCallContext) {
}

// EnterEnumClassBody is called when production enumClassBody is entered.
func (s *KotlinAstParser) EnterEnumClassBody(ctx *kotlin.EnumClassBodyContext) {}

// ExitEnumClassBody is called when production enumClassBody is exited.
func (s *KotlinAstParser) ExitEnumClassBody(ctx *kotlin.EnumClassBodyContext) {}

// EnterEnumEntries is called when production enumEntries is entered.
func (s *KotlinAstParser) EnterEnumEntries(ctx *kotlin.EnumEntriesContext) {}

// ExitEnumEntries is called when production enumEntries is exited.
func (s *KotlinAstParser) ExitEnumEntries(ctx *kotlin.EnumEntriesContext) {}

// EnterEnumEntry is called when production enumEntry is entered.
func (s *KotlinAstParser) EnterEnumEntry(ctx *kotlin.EnumEntryContext) {}

// ExitEnumEntry is called when production enumEntry is exited.
func (s *KotlinAstParser) ExitEnumEntry(ctx *kotlin.EnumEntryContext) {}

// EnterFunctionValueParameters is called when production functionValueParameters is entered.
func (s *KotlinAstParser) EnterFunctionValueParameters(ctx *kotlin.FunctionValueParametersContext) {
	//fmt.Printf("!!!EnterFunctionValueParameters:%s,counts:%d\n",ctx.GetText(),ctx.GetChildCount())
}

// ExitFunctionValueParameters is called when production functionValueParameters is exited.
func (s *KotlinAstParser) ExitFunctionValueParameters(ctx *kotlin.FunctionValueParametersContext) {
}

// ExitFunctionValueParameter is called when production functionValueParameter is exited.
func (s *KotlinAstParser) ExitFunctionValueParameter(ctx *kotlin.FunctionValueParameterContext) {
}

// EnterParameter is called when production parameter is entered.
func (s *KotlinAstParser) EnterParameter(ctx *kotlin.ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *KotlinAstParser) ExitParameter(ctx *kotlin.ParameterContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *KotlinAstParser) EnterFunctionBody(ctx *kotlin.FunctionBodyContext) {}

// EnterObjectDeclaration is called when production objectDeclaration is entered.
func (s *KotlinAstParser) EnterObjectDeclaration(ctx *kotlin.ObjectDeclarationContext) {}

// ExitInfixFunctionCall is called when production infixFunctionCall is exited.
func (s *KotlinAstParser) ExitInfixFunctionCall(ctx *kotlin.InfixFunctionCallContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *KotlinAstParser) ExitFunctionBody(ctx *kotlin.FunctionBodyContext) {}

// ExitObjectDeclaration is called when production objectDeclaration is exited.
func (s *KotlinAstParser) ExitObjectDeclaration(ctx *kotlin.ObjectDeclarationContext) {}

// EnterCompanionObject is called when production companionObject is entered.
func (s *KotlinAstParser) EnterCompanionObject(ctx *kotlin.CompanionObjectContext) {}

// ExitCompanionObject is called when production companionObject is exited.
func (s *KotlinAstParser) ExitCompanionObject(ctx *kotlin.CompanionObjectContext) {}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *KotlinAstParser) EnterPropertyDeclaration(ctx *kotlin.PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *KotlinAstParser) ExitPropertyDeclaration(ctx *kotlin.PropertyDeclarationContext) {}

// EnterMultiVariableDeclaration is called when production multiVariableDeclaration is entered.
func (s *KotlinAstParser) EnterMultiVariableDeclaration(ctx *kotlin.MultiVariableDeclarationContext) {
}

// ExitMultiVariableDeclaration is called when production multiVariableDeclaration is exited.
func (s *KotlinAstParser) ExitMultiVariableDeclaration(ctx *kotlin.MultiVariableDeclarationContext) {
}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *KotlinAstParser) EnterVariableDeclaration(ctx *kotlin.VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *KotlinAstParser) ExitVariableDeclaration(ctx *kotlin.VariableDeclarationContext) {}

// EnterGetter is called when production getter is entered.
func (s *KotlinAstParser) EnterGetter(ctx *kotlin.GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *KotlinAstParser) ExitGetter(ctx *kotlin.GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *KotlinAstParser) EnterSetter(ctx *kotlin.SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *KotlinAstParser) ExitSetter(ctx *kotlin.SetterContext) {}

// EnterTypeAlias is called when production typeAlias is entered.
func (s *KotlinAstParser) EnterTypeAlias(ctx *kotlin.TypeAliasContext) {}

// ExitTypeAlias is called when production typeAlias is exited.
func (s *KotlinAstParser) ExitTypeAlias(ctx *kotlin.TypeAliasContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *KotlinAstParser) EnterTypeParameters(ctx *kotlin.TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *KotlinAstParser) ExitTypeParameters(ctx *kotlin.TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *KotlinAstParser) EnterTypeParameter(ctx *kotlin.TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *KotlinAstParser) ExitTypeParameter(ctx *kotlin.TypeParameterContext) {}

// EnterType is called when production type is entered.
func (s *KotlinAstParser) EnterType(ctx *kotlin.TypeContext) {}

// ExitType is called when production type is exited.
func (s *KotlinAstParser) ExitType(ctx *kotlin.TypeContext) {}

// EnterTypeModifierList is called when production typeModifierList is entered.
func (s *KotlinAstParser) EnterTypeModifierList(ctx *kotlin.TypeModifierListContext) {}

// ExitTypeModifierList is called when production typeModifierList is exited.
func (s *KotlinAstParser) ExitTypeModifierList(ctx *kotlin.TypeModifierListContext) {}

// EnterParenthesizedType is called when production parenthesizedType is entered.
func (s *KotlinAstParser) EnterParenthesizedType(ctx *kotlin.ParenthesizedTypeContext) {}

// ExitParenthesizedType is called when production parenthesizedType is exited.
func (s *KotlinAstParser) ExitParenthesizedType(ctx *kotlin.ParenthesizedTypeContext) {}

// EnterNullableType is called when production nullableType is entered.
func (s *KotlinAstParser) EnterNullableType(ctx *kotlin.NullableTypeContext) {}

// ExitNullableType is called when production nullableType is exited.
func (s *KotlinAstParser) ExitNullableType(ctx *kotlin.NullableTypeContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *KotlinAstParser) EnterTypeReference(ctx *kotlin.TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *KotlinAstParser) ExitTypeReference(ctx *kotlin.TypeReferenceContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *KotlinAstParser) EnterFunctionType(ctx *kotlin.FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *KotlinAstParser) ExitFunctionType(ctx *kotlin.FunctionTypeContext) {}

// EnterFunctionTypeReceiver is called when production functionTypeReceiver is entered.
func (s *KotlinAstParser) EnterFunctionTypeReceiver(ctx *kotlin.FunctionTypeReceiverContext) {
}

// ExitFunctionTypeReceiver is called when production functionTypeReceiver is exited.
func (s *KotlinAstParser) ExitFunctionTypeReceiver(ctx *kotlin.FunctionTypeReceiverContext) {}

// EnterUserType is called when production userType is entered.
func (s *KotlinAstParser) EnterUserType(ctx *kotlin.UserTypeContext) {}

// ExitUserType is called when production userType is exited.
func (s *KotlinAstParser) ExitUserType(ctx *kotlin.UserTypeContext) {}

// EnterSimpleUserType is called when production simpleUserType is entered.
func (s *KotlinAstParser) EnterSimpleUserType(ctx *kotlin.SimpleUserTypeContext) {}

// ExitSimpleUserType is called when production simpleUserType is exited.
func (s *KotlinAstParser) ExitSimpleUserType(ctx *kotlin.SimpleUserTypeContext) {}

// EnterFunctionTypeParameters is called when production functionTypeParameters is entered.
func (s *KotlinAstParser) EnterFunctionTypeParameters(ctx *kotlin.FunctionTypeParametersContext) {
}

// ExitFunctionTypeParameters is called when production functionTypeParameters is exited.
func (s *KotlinAstParser) ExitFunctionTypeParameters(ctx *kotlin.FunctionTypeParametersContext) {
}

// EnterTypeConstraints is called when production typeConstraints is entered.
func (s *KotlinAstParser) EnterTypeConstraints(ctx *kotlin.TypeConstraintsContext) {}

// ExitTypeConstraints is called when production typeConstraints is exited.
func (s *KotlinAstParser) ExitTypeConstraints(ctx *kotlin.TypeConstraintsContext) {}

// EnterTypeConstraint is called when production typeConstraint is entered.
func (s *KotlinAstParser) EnterTypeConstraint(ctx *kotlin.TypeConstraintContext) {}

// ExitTypeConstraint is called when production typeConstraint is exited.
func (s *KotlinAstParser) ExitTypeConstraint(ctx *kotlin.TypeConstraintContext) {}

// EnterBlock is called when production block is entered.
func (s *KotlinAstParser) EnterBlock(ctx *kotlin.BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *KotlinAstParser) ExitBlock(ctx *kotlin.BlockContext) {}

// EnterStatements is called when production statements is entered.
func (s *KotlinAstParser) EnterStatements(ctx *kotlin.StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *KotlinAstParser) ExitStatements(ctx *kotlin.StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *KotlinAstParser) EnterStatement(ctx *kotlin.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *KotlinAstParser) ExitStatement(ctx *kotlin.StatementContext) {}

// EnterBlockLevelExpression is called when production blockLevelExpression is entered.
func (s *KotlinAstParser) EnterBlockLevelExpression(ctx *kotlin.BlockLevelExpressionContext) {
}

// ExitBlockLevelExpression is called when production blockLevelExpression is exited.
func (s *KotlinAstParser) ExitBlockLevelExpression(ctx *kotlin.BlockLevelExpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *KotlinAstParser) EnterDeclaration(ctx *kotlin.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *KotlinAstParser) ExitDeclaration(ctx *kotlin.DeclarationContext) {}

// EnterExpression is called when production expression is entered.
func (s *KotlinAstParser) EnterExpression(ctx *kotlin.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *KotlinAstParser) ExitExpression(ctx *kotlin.ExpressionContext) {}

// EnterDisjunction is called when production disjunction is entered.
func (s *KotlinAstParser) EnterDisjunction(ctx *kotlin.DisjunctionContext) {}

// ExitDisjunction is called when production disjunction is exited.
func (s *KotlinAstParser) ExitDisjunction(ctx *kotlin.DisjunctionContext) {}

// EnterConjunction is called when production conjunction is entered.
func (s *KotlinAstParser) EnterConjunction(ctx *kotlin.ConjunctionContext) {}

// ExitConjunction is called when production conjunction is exited.
func (s *KotlinAstParser) ExitConjunction(ctx *kotlin.ConjunctionContext) {}

// EnterEqualityComparison is called when production equalityComparison is entered.
func (s *KotlinAstParser) EnterEqualityComparison(ctx *kotlin.EqualityComparisonContext) {}

// ExitEqualityComparison is called when production equalityComparison is exited.
func (s *KotlinAstParser) ExitEqualityComparison(ctx *kotlin.EqualityComparisonContext) {}

// EnterComparison is called when production comparison is entered.
func (s *KotlinAstParser) EnterComparison(ctx *kotlin.ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *KotlinAstParser) ExitComparison(ctx *kotlin.ComparisonContext) {}

// EnterNamedInfix is called when production namedInfix is entered.
func (s *KotlinAstParser) EnterNamedInfix(ctx *kotlin.NamedInfixContext) {}

// ExitNamedInfix is called when production namedInfix is exited.
func (s *KotlinAstParser) ExitNamedInfix(ctx *kotlin.NamedInfixContext) {}

// EnterElvisExpression is called when production elvisExpression is entered.
func (s *KotlinAstParser) EnterElvisExpression(ctx *kotlin.ElvisExpressionContext) {}

// ExitElvisExpression is called when production elvisExpression is exited.
func (s *KotlinAstParser) ExitElvisExpression(ctx *kotlin.ElvisExpressionContext) {}

// EnterInfixFunctionCall is called when production infixFunctionCall is entered.
func (s *KotlinAstParser) EnterInfixFunctionCall(ctx *kotlin.InfixFunctionCallContext) {
}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *KotlinAstParser) EnterRangeExpression(ctx *kotlin.RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *KotlinAstParser) ExitRangeExpression(ctx *kotlin.RangeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *KotlinAstParser) EnterAdditiveExpression(ctx *kotlin.AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *KotlinAstParser) ExitAdditiveExpression(ctx *kotlin.AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *KotlinAstParser) EnterMultiplicativeExpression(ctx *kotlin.MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *KotlinAstParser) ExitMultiplicativeExpression(ctx *kotlin.MultiplicativeExpressionContext) {
}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *KotlinAstParser) ExitClassDeclaration(ctx *kotlin.ClassDeclarationContext) {}

// EnterTypeRHS is called when production typeRHS is entered.
func (s *KotlinAstParser) EnterTypeRHS(ctx *kotlin.TypeRHSContext) {}

// ExitTypeRHS is called when production typeRHS is exited.
func (s *KotlinAstParser) ExitTypeRHS(ctx *kotlin.TypeRHSContext) {}

// EnterPrefixUnaryExpression is called when production prefixUnaryExpression is entered.
func (s *KotlinAstParser) EnterPrefixUnaryExpression(ctx *kotlin.PrefixUnaryExpressionContext) {
}

// ExitPrefixUnaryExpression is called when production prefixUnaryExpression is exited.
func (s *KotlinAstParser) ExitPrefixUnaryExpression(ctx *kotlin.PrefixUnaryExpressionContext) {
}

// EnterPostfixUnaryExpression is called when production postfixUnaryExpression is entered.
func (s *KotlinAstParser) EnterPostfixUnaryExpression(ctx *kotlin.PostfixUnaryExpressionContext) {
}

// ExitPostfixUnaryExpression is called when production postfixUnaryExpression is exited.
func (s *KotlinAstParser) ExitPostfixUnaryExpression(ctx *kotlin.PostfixUnaryExpressionContext) {
}

// EnterAtomicExpression is called when production atomicExpression is entered.
func (s *KotlinAstParser) EnterAtomicExpression(ctx *kotlin.AtomicExpressionContext) {}

// ExitAtomicExpression is called when production atomicExpression is exited.
func (s *KotlinAstParser) ExitAtomicExpression(ctx *kotlin.AtomicExpressionContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *KotlinAstParser) EnterParenthesizedExpression(ctx *kotlin.ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *KotlinAstParser) ExitParenthesizedExpression(ctx *kotlin.ParenthesizedExpressionContext) {
}

// EnterCallSuffix is called when production callSuffix is entered.
func (s *KotlinAstParser) EnterCallSuffix(ctx *kotlin.CallSuffixContext) {}

// ExitCallSuffix is called when production callSuffix is exited.
func (s *KotlinAstParser) ExitCallSuffix(ctx *kotlin.CallSuffixContext) {}

// EnterAnnotatedLambda is called when production annotatedLambda is entered.
func (s *KotlinAstParser) EnterAnnotatedLambda(ctx *kotlin.AnnotatedLambdaContext) {}

// ExitAnnotatedLambda is called when production annotatedLambda is exited.
func (s *KotlinAstParser) ExitAnnotatedLambda(ctx *kotlin.AnnotatedLambdaContext) {}

// EnterArrayAccess is called when production arrayAccess is entered.
func (s *KotlinAstParser) EnterArrayAccess(ctx *kotlin.ArrayAccessContext) {}

// ExitArrayAccess is called when production arrayAccess is exited.
func (s *KotlinAstParser) ExitArrayAccess(ctx *kotlin.ArrayAccessContext) {}

// EnterValueArguments is called when production valueArguments is entered.
func (s *KotlinAstParser) EnterValueArguments(ctx *kotlin.ValueArgumentsContext) {}

// ExitValueArguments is called when production valueArguments is exited.
func (s *KotlinAstParser) ExitValueArguments(ctx *kotlin.ValueArgumentsContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *KotlinAstParser) EnterTypeArguments(ctx *kotlin.TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *KotlinAstParser) ExitTypeArguments(ctx *kotlin.TypeArgumentsContext) {}

// EnterTypeProjection is called when production typeProjection is entered.
func (s *KotlinAstParser) EnterTypeProjection(ctx *kotlin.TypeProjectionContext) {}

// ExitTypeProjection is called when production typeProjection is exited.
func (s *KotlinAstParser) ExitTypeProjection(ctx *kotlin.TypeProjectionContext) {}

// EnterTypeProjectionModifierList is called when production typeProjectionModifierList is entered.
func (s *KotlinAstParser) EnterTypeProjectionModifierList(ctx *kotlin.TypeProjectionModifierListContext) {
}

// ExitTypeProjectionModifierList is called when production typeProjectionModifierList is exited.
func (s *KotlinAstParser) ExitTypeProjectionModifierList(ctx *kotlin.TypeProjectionModifierListContext) {
}

// EnterValueArgument is called when production valueArgument is entered.
func (s *KotlinAstParser) EnterValueArgument(ctx *kotlin.ValueArgumentContext) {}

// ExitValueArgument is called when production valueArgument is exited.
func (s *KotlinAstParser) ExitValueArgument(ctx *kotlin.ValueArgumentContext) {}

// EnterLiteralConstant is called when production literalConstant is entered.
func (s *KotlinAstParser) EnterLiteralConstant(ctx *kotlin.LiteralConstantContext) {}

// ExitLiteralConstant is called when production literalConstant is exited.
func (s *KotlinAstParser) ExitLiteralConstant(ctx *kotlin.LiteralConstantContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *KotlinAstParser) EnterStringLiteral(ctx *kotlin.StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *KotlinAstParser) ExitStringLiteral(ctx *kotlin.StringLiteralContext) {}

// EnterLineStringLiteral is called when production lineStringLiteral is entered.
func (s *KotlinAstParser) EnterLineStringLiteral(ctx *kotlin.LineStringLiteralContext) {}

// ExitLineStringLiteral is called when production lineStringLiteral is exited.
func (s *KotlinAstParser) ExitLineStringLiteral(ctx *kotlin.LineStringLiteralContext) {}

// EnterMultiLineStringLiteral is called when production multiLineStringLiteral is entered.
func (s *KotlinAstParser) EnterMultiLineStringLiteral(ctx *kotlin.MultiLineStringLiteralContext) {
}

// ExitMultiLineStringLiteral is called when production multiLineStringLiteral is exited.
func (s *KotlinAstParser) ExitMultiLineStringLiteral(ctx *kotlin.MultiLineStringLiteralContext) {
}

// EnterLineStringContent is called when production lineStringContent is entered.
func (s *KotlinAstParser) EnterLineStringContent(ctx *kotlin.LineStringContentContext) {}

// ExitLineStringContent is called when production lineStringContent is exited.
func (s *KotlinAstParser) ExitLineStringContent(ctx *kotlin.LineStringContentContext) {}

// EnterLineStringExpression is called when production lineStringExpression is entered.
func (s *KotlinAstParser) EnterLineStringExpression(ctx *kotlin.LineStringExpressionContext) {
}

// ExitLineStringExpression is called when production lineStringExpression is exited.
func (s *KotlinAstParser) ExitLineStringExpression(ctx *kotlin.LineStringExpressionContext) {}

// EnterMultiLineStringContent is called when production multiLineStringContent is entered.
func (s *KotlinAstParser) EnterMultiLineStringContent(ctx *kotlin.MultiLineStringContentContext) {
}

// ExitMultiLineStringContent is called when production multiLineStringContent is exited.
func (s *KotlinAstParser) ExitMultiLineStringContent(ctx *kotlin.MultiLineStringContentContext) {
}

// EnterMultiLineStringExpression is called when production multiLineStringExpression is entered.
func (s *KotlinAstParser) EnterMultiLineStringExpression(ctx *kotlin.MultiLineStringExpressionContext) {
}

// ExitMultiLineStringExpression is called when production multiLineStringExpression is exited.
func (s *KotlinAstParser) ExitMultiLineStringExpression(ctx *kotlin.MultiLineStringExpressionContext) {
}

// EnterFunctionLiteral is called when production functionLiteral is entered.
func (s *KotlinAstParser) EnterFunctionLiteral(ctx *kotlin.FunctionLiteralContext) {}

// ExitFunctionLiteral is called when production functionLiteral is exited.
func (s *KotlinAstParser) ExitFunctionLiteral(ctx *kotlin.FunctionLiteralContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *KotlinAstParser) EnterLambdaParameters(ctx *kotlin.LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *KotlinAstParser) ExitLambdaParameters(ctx *kotlin.LambdaParametersContext) {}

// EnterLambdaParameter is called when production lambdaParameter is entered.
func (s *KotlinAstParser) EnterLambdaParameter(ctx *kotlin.LambdaParameterContext) {}

// ExitLambdaParameter is called when production lambdaParameter is exited.
func (s *KotlinAstParser) ExitLambdaParameter(ctx *kotlin.LambdaParameterContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *KotlinAstParser) EnterObjectLiteral(ctx *kotlin.ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *KotlinAstParser) ExitObjectLiteral(ctx *kotlin.ObjectLiteralContext) {}

// EnterCollectionLiteral is called when production collectionLiteral is entered.
func (s *KotlinAstParser) EnterCollectionLiteral(ctx *kotlin.CollectionLiteralContext) {}

// ExitCollectionLiteral is called when production collectionLiteral is exited.
func (s *KotlinAstParser) ExitCollectionLiteral(ctx *kotlin.CollectionLiteralContext) {}

// EnterThisExpression is called when production thisExpression is entered.
func (s *KotlinAstParser) EnterThisExpression(ctx *kotlin.ThisExpressionContext) {}

// ExitThisExpression is called when production thisExpression is exited.
func (s *KotlinAstParser) ExitThisExpression(ctx *kotlin.ThisExpressionContext) {}

// EnterSuperExpression is called when production superExpression is entered.
func (s *KotlinAstParser) EnterSuperExpression(ctx *kotlin.SuperExpressionContext) {}

// ExitSuperExpression is called when production superExpression is exited.
func (s *KotlinAstParser) ExitSuperExpression(ctx *kotlin.SuperExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *KotlinAstParser) EnterConditionalExpression(ctx *kotlin.ConditionalExpressionContext) {
}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *KotlinAstParser) ExitConditionalExpression(ctx *kotlin.ConditionalExpressionContext) {
}

// EnterIfExpression is called when production ifExpression is entered.
func (s *KotlinAstParser) EnterIfExpression(ctx *kotlin.IfExpressionContext) {}

// ExitIfExpression is called when production ifExpression is exited.
func (s *KotlinAstParser) ExitIfExpression(ctx *kotlin.IfExpressionContext) {}

// EnterControlStructureBody is called when production controlStructureBody is entered.
func (s *KotlinAstParser) EnterControlStructureBody(ctx *kotlin.ControlStructureBodyContext) {
}

// ExitControlStructureBody is called when production controlStructureBody is exited.
func (s *KotlinAstParser) ExitControlStructureBody(ctx *kotlin.ControlStructureBodyContext) {}

// EnterWhenExpression is called when production whenExpression is entered.
func (s *KotlinAstParser) EnterWhenExpression(ctx *kotlin.WhenExpressionContext) {}

// ExitWhenExpression is called when production whenExpression is exited.
func (s *KotlinAstParser) ExitWhenExpression(ctx *kotlin.WhenExpressionContext) {}

// EnterWhenEntry is called when production whenEntry is entered.
func (s *KotlinAstParser) EnterWhenEntry(ctx *kotlin.WhenEntryContext) {}

// ExitWhenEntry is called when production whenEntry is exited.
func (s *KotlinAstParser) ExitWhenEntry(ctx *kotlin.WhenEntryContext) {}

// EnterWhenCondition is called when production whenCondition is entered.
func (s *KotlinAstParser) EnterWhenCondition(ctx *kotlin.WhenConditionContext) {}

// ExitWhenCondition is called when production whenCondition is exited.
func (s *KotlinAstParser) ExitWhenCondition(ctx *kotlin.WhenConditionContext) {}

// EnterRangeTest is called when production rangeTest is entered.
func (s *KotlinAstParser) EnterRangeTest(ctx *kotlin.RangeTestContext) {}

// ExitRangeTest is called when production rangeTest is exited.
func (s *KotlinAstParser) ExitRangeTest(ctx *kotlin.RangeTestContext) {}

// EnterTypeTest is called when production typeTest is entered.
func (s *KotlinAstParser) EnterTypeTest(ctx *kotlin.TypeTestContext) {}

// ExitTypeTest is called when production typeTest is exited.
func (s *KotlinAstParser) ExitTypeTest(ctx *kotlin.TypeTestContext) {}

// EnterTryExpression is called when production tryExpression is entered.
func (s *KotlinAstParser) EnterTryExpression(ctx *kotlin.TryExpressionContext) {}

// ExitTryExpression is called when production tryExpression is exited.
func (s *KotlinAstParser) ExitTryExpression(ctx *kotlin.TryExpressionContext) {}

// EnterCatchBlock is called when production catchBlock is entered.
func (s *KotlinAstParser) EnterCatchBlock(ctx *kotlin.CatchBlockContext) {}

// ExitCatchBlock is called when production catchBlock is exited.
func (s *KotlinAstParser) ExitCatchBlock(ctx *kotlin.CatchBlockContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *KotlinAstParser) EnterFinallyBlock(ctx *kotlin.FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *KotlinAstParser) ExitFinallyBlock(ctx *kotlin.FinallyBlockContext) {}

// EnterLoopExpression is called when production loopExpression is entered.
func (s *KotlinAstParser) EnterLoopExpression(ctx *kotlin.LoopExpressionContext) {}

// ExitLoopExpression is called when production loopExpression is exited.
func (s *KotlinAstParser) ExitLoopExpression(ctx *kotlin.LoopExpressionContext) {}

// EnterForExpression is called when production forExpression is entered.
func (s *KotlinAstParser) EnterForExpression(ctx *kotlin.ForExpressionContext) {}

// ExitForExpression is called when production forExpression is exited.
func (s *KotlinAstParser) ExitForExpression(ctx *kotlin.ForExpressionContext) {}

// EnterWhileExpression is called when production whileExpression is entered.
func (s *KotlinAstParser) EnterWhileExpression(ctx *kotlin.WhileExpressionContext) {}

// ExitWhileExpression is called when production whileExpression is exited.
func (s *KotlinAstParser) ExitWhileExpression(ctx *kotlin.WhileExpressionContext) {}

// EnterDoWhileExpression is called when production doWhileExpression is entered.
func (s *KotlinAstParser) EnterDoWhileExpression(ctx *kotlin.DoWhileExpressionContext) {}

// ExitDoWhileExpression is called when production doWhileExpression is exited.
func (s *KotlinAstParser) ExitDoWhileExpression(ctx *kotlin.DoWhileExpressionContext) {}

// EnterJumpExpression is called when production jumpExpression is entered.
func (s *KotlinAstParser) EnterJumpExpression(ctx *kotlin.JumpExpressionContext) {}

// ExitJumpExpression is called when production jumpExpression is exited.
func (s *KotlinAstParser) ExitJumpExpression(ctx *kotlin.JumpExpressionContext) {}

// EnterCallableReference is called when production callableReference is entered.
func (s *KotlinAstParser) EnterCallableReference(ctx *kotlin.CallableReferenceContext) {}

// ExitCallableReference is called when production callableReference is exited.
func (s *KotlinAstParser) ExitCallableReference(ctx *kotlin.CallableReferenceContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *KotlinAstParser) EnterAssignmentOperator(ctx *kotlin.AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *KotlinAstParser) ExitAssignmentOperator(ctx *kotlin.AssignmentOperatorContext) {}

// EnterEqualityOperation is called when production equalityOperation is entered.
func (s *KotlinAstParser) EnterEqualityOperation(ctx *kotlin.EqualityOperationContext) {}

// ExitEqualityOperation is called when production equalityOperation is exited.
func (s *KotlinAstParser) ExitEqualityOperation(ctx *kotlin.EqualityOperationContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *KotlinAstParser) EnterComparisonOperator(ctx *kotlin.ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *KotlinAstParser) ExitComparisonOperator(ctx *kotlin.ComparisonOperatorContext) {}

// EnterInOperator is called when production inOperator is entered.
func (s *KotlinAstParser) EnterInOperator(ctx *kotlin.InOperatorContext) {}

// ExitInOperator is called when production inOperator is exited.
func (s *KotlinAstParser) ExitInOperator(ctx *kotlin.InOperatorContext) {}

// EnterIsOperator is called when production isOperator is entered.
func (s *KotlinAstParser) EnterIsOperator(ctx *kotlin.IsOperatorContext) {}

// ExitIsOperator is called when production isOperator is exited.
func (s *KotlinAstParser) ExitIsOperator(ctx *kotlin.IsOperatorContext) {}

// EnterAdditiveOperator is called when production additiveOperator is entered.
func (s *KotlinAstParser) EnterAdditiveOperator(ctx *kotlin.AdditiveOperatorContext) {}

// ExitAdditiveOperator is called when production additiveOperator is exited.
func (s *KotlinAstParser) ExitAdditiveOperator(ctx *kotlin.AdditiveOperatorContext) {}

// EnterMultiplicativeOperation is called when production multiplicativeOperation is entered.
func (s *KotlinAstParser) EnterMultiplicativeOperation(ctx *kotlin.MultiplicativeOperationContext) {
}

// ExitMultiplicativeOperation is called when production multiplicativeOperation is exited.
func (s *KotlinAstParser) ExitMultiplicativeOperation(ctx *kotlin.MultiplicativeOperationContext) {
}

// EnterTypeOperation is called when production typeOperation is entered.
func (s *KotlinAstParser) EnterTypeOperation(ctx *kotlin.TypeOperationContext) {}

// ExitTypeOperation is called when production typeOperation is exited.
func (s *KotlinAstParser) ExitTypeOperation(ctx *kotlin.TypeOperationContext) {}

// EnterPrefixUnaryOperation is called when production prefixUnaryOperation is entered.
func (s *KotlinAstParser) EnterPrefixUnaryOperation(ctx *kotlin.PrefixUnaryOperationContext) {
}

// ExitPrefixUnaryOperation is called when production prefixUnaryOperation is exited.
func (s *KotlinAstParser) ExitPrefixUnaryOperation(ctx *kotlin.PrefixUnaryOperationContext) {
}

// EnterPostfixUnaryOperation is called when production postfixUnaryOperation is entered.
func (s *KotlinAstParser) EnterPostfixUnaryOperation(ctx *kotlin.PostfixUnaryOperationContext) {
}

// EnterMemberAccessOperator is called when production memberAccessOperator is entered.
func (s *KotlinAstParser) EnterMemberAccessOperator(ctx *kotlin.MemberAccessOperatorContext) {
}

// ExitMemberAccessOperator is called when production memberAccessOperator is exited.
func (s *KotlinAstParser) ExitMemberAccessOperator(ctx *kotlin.MemberAccessOperatorContext) {}

// EnterModifierList is called when production modifierList is entered.
func (s *KotlinAstParser) EnterModifierList(ctx *kotlin.ModifierListContext) {}

// ExitModifierList is called when production modifierList is exited.
func (s *KotlinAstParser) ExitModifierList(ctx *kotlin.ModifierListContext) {}

// EnterModifier is called when production modifier is entered.
func (s *KotlinAstParser) EnterModifier(ctx *kotlin.ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *KotlinAstParser) ExitModifier(ctx *kotlin.ModifierContext) {}

// EnterClassModifier is called when production classModifier is entered.
func (s *KotlinAstParser) EnterClassModifier(ctx *kotlin.ClassModifierContext) {}

// ExitClassModifier is called when production classModifier is exited.
func (s *KotlinAstParser) ExitClassModifier(ctx *kotlin.ClassModifierContext) {}

// EnterMemberModifier is called when production memberModifier is entered.
func (s *KotlinAstParser) EnterMemberModifier(ctx *kotlin.MemberModifierContext) {}

// ExitMemberModifier is called when production memberModifier is exited.
func (s *KotlinAstParser) ExitMemberModifier(ctx *kotlin.MemberModifierContext) {}

// EnterVisibilityModifier is called when production visibilityModifier is entered.
func (s *KotlinAstParser) EnterVisibilityModifier(ctx *kotlin.VisibilityModifierContext) {}

// ExitVisibilityModifier is called when production visibilityModifier is exited.
func (s *KotlinAstParser) ExitVisibilityModifier(ctx *kotlin.VisibilityModifierContext) {}

// EnterVarianceAnnotation is called when production varianceAnnotation is entered.
func (s *KotlinAstParser) EnterVarianceAnnotation(ctx *kotlin.VarianceAnnotationContext) {}

// ExitVarianceAnnotation is called when production varianceAnnotation is exited.
func (s *KotlinAstParser) ExitVarianceAnnotation(ctx *kotlin.VarianceAnnotationContext) {}

// EnterFunctionModifier is called when production functionModifier is entered.
func (s *KotlinAstParser) EnterFunctionModifier(ctx *kotlin.FunctionModifierContext) {}

// ExitFunctionModifier is called when production functionModifier is exited.
func (s *KotlinAstParser) ExitFunctionModifier(ctx *kotlin.FunctionModifierContext) {}

// EnterPropertyModifier is called when production propertyModifier is entered.
func (s *KotlinAstParser) EnterPropertyModifier(ctx *kotlin.PropertyModifierContext) {}

// ExitPropertyModifier is called when production propertyModifier is exited.
func (s *KotlinAstParser) ExitPropertyModifier(ctx *kotlin.PropertyModifierContext) {}

// EnterInheritanceModifier is called when production inheritanceModifier is entered.
func (s *KotlinAstParser) EnterInheritanceModifier(ctx *kotlin.InheritanceModifierContext) {}

// ExitInheritanceModifier is called when production inheritanceModifier is exited.
func (s *KotlinAstParser) ExitInheritanceModifier(ctx *kotlin.InheritanceModifierContext) {}

// EnterParameterModifier is called when production parameterModifier is entered.
func (s *KotlinAstParser) EnterParameterModifier(ctx *kotlin.ParameterModifierContext) {}

// ExitParameterModifier is called when production parameterModifier is exited.
func (s *KotlinAstParser) ExitParameterModifier(ctx *kotlin.ParameterModifierContext) {}

// EnterTypeParameterModifier is called when production typeParameterModifier is entered.
func (s *KotlinAstParser) EnterTypeParameterModifier(ctx *kotlin.TypeParameterModifierContext) {
}

// ExitTypeParameterModifier is called when production typeParameterModifier is exited.
func (s *KotlinAstParser) ExitTypeParameterModifier(ctx *kotlin.TypeParameterModifierContext) {
}

// EnterLabelDefinition is called when production labelDefinition is entered.
func (s *KotlinAstParser) EnterLabelDefinition(ctx *kotlin.LabelDefinitionContext) {}

// ExitLabelDefinition is called when production labelDefinition is exited.
func (s *KotlinAstParser) ExitLabelDefinition(ctx *kotlin.LabelDefinitionContext) {}

// EnterAnnotations is called when production annotations is entered.
func (s *KotlinAstParser) EnterAnnotations(ctx *kotlin.AnnotationsContext) {}

// ExitAnnotations is called when production annotations is exited.
func (s *KotlinAstParser) ExitAnnotations(ctx *kotlin.AnnotationsContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *KotlinAstParser) EnterAnnotation(ctx *kotlin.AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *KotlinAstParser) ExitAnnotation(ctx *kotlin.AnnotationContext) {}

// EnterAnnotationList is called when production annotationList is entered.
func (s *KotlinAstParser) EnterAnnotationList(ctx *kotlin.AnnotationListContext) {}

// ExitAnnotationList is called when production annotationList is exited.
func (s *KotlinAstParser) ExitAnnotationList(ctx *kotlin.AnnotationListContext) {}

// EnterAnnotationUseSiteTarget is called when production annotationUseSiteTarget is entered.
func (s *KotlinAstParser) EnterAnnotationUseSiteTarget(ctx *kotlin.AnnotationUseSiteTargetContext) {
}

// ExitAnnotationUseSiteTarget is called when production annotationUseSiteTarget is exited.
func (s *KotlinAstParser) ExitAnnotationUseSiteTarget(ctx *kotlin.AnnotationUseSiteTargetContext) {
}

// EnterUnescapedAnnotation is called when production unescapedAnnotation is entered.
func (s *KotlinAstParser) EnterUnescapedAnnotation(ctx *kotlin.UnescapedAnnotationContext) {}

// ExitUnescapedAnnotation is called when production unescapedAnnotation is exited.
func (s *KotlinAstParser) ExitUnescapedAnnotation(ctx *kotlin.UnescapedAnnotationContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *KotlinAstParser) EnterIdentifier(ctx *kotlin.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *KotlinAstParser) ExitIdentifier(ctx *kotlin.IdentifierContext) {}

// EnterSimpleIdentifier is called when production simpleIdentifier is entered.
func (s *KotlinAstParser) EnterSimpleIdentifier(ctx *kotlin.SimpleIdentifierContext) {}

// ExitSimpleIdentifier is called when production simpleIdentifier is exited.
func (s *KotlinAstParser) ExitSimpleIdentifier(ctx *kotlin.SimpleIdentifierContext) {}

// EnterSemi is called when production semi is entered.
func (s *KotlinAstParser) EnterSemi(ctx *kotlin.SemiContext) {}

// ExitSemi is called when production semi is exited.
func (s *KotlinAstParser) ExitSemi(ctx *kotlin.SemiContext) {}

// EnterAnysemi is called when production anysemi is entered.
func (s *KotlinAstParser) EnterAnysemi(ctx *kotlin.AnysemiContext) {}

// ExitAnysemi is called when production anysemi is exited.
func (s *KotlinAstParser) ExitAnysemi(ctx *kotlin.AnysemiContext) {}

func (s *KotlinAstParser) EnterFunctionValueParameter(c *kotlin.FunctionValueParameterContext) {
}
