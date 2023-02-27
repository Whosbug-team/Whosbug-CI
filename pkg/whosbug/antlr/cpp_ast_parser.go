package antlr

import (
	"strings"
	"sync"

	"git.woa.com/bkdevops/whosbug-ci/internal/util"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/antlr/cpp"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CppAstParser implement AstParser for c++ language
//
//	@author kevineluo
//	@update 2023-02-28 12:57:34
type CppAstParser struct {
	AstInfo AstInfo
}

var (
	_ cpp.CPP14ParserListener = &CppAstParser{}
	_ AstParser               = &CppAstParser{}
)

var (
	cppLexerPool = &sync.Pool{New: func() any {
		return cpp.NewCPP14Lexer(nil)
	}}
	cppParserPool = &sync.Pool{New: func() any {
		return cpp.NewCPP14Parser(nil)
	}}
	newCppAstParserPool = &sync.Pool{New: func() any {
		return new(CppAstParser)
	}}
)

// Init 初始化AstParser
//
//	@receiver s *CAstParser
//	@author kevineluo
//	@update 2023-02-28 03:13:43
func (s *CppAstParser) Init() (err error) {
	s.AstInfo.Classes = make([]Class, 0)
	s.AstInfo.Methods = make([]Method, 0)
	return
}

// AstParse main parse process for c++ language
//
//	@receiver s *CppAstParser
//	@param diffText string
//	@return AstInfo
//	@author kevineluo
//	@update 2023-02-28 01:03:36
func (s *CppAstParser) AstParse(diffText string) AstInfo {
	//	截取目标文本的输入流
	input := antlr.NewInputStream(diffText)
	//	初始化lexer
	lexer := cppLexerPool.Get().(*cpp.CPP14Lexer)
	defer cppLexerPool.Put(lexer)
	lexer.SetInputStream(input)
	lexer.RemoveErrorListeners()
	//	初始化Token流
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//	初始化Parser
	p := cppParserPool.Get().(*cpp.CPP14Parser)
	p.RemoveErrorListeners()
	defer cppParserPool.Put(p)
	p.SetTokenStream(stream)
	//	构建语法解析树
	p.BuildParseTrees = true
	// // 启用SLL两阶段加速解析模式
	// p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	//	解析模式->每个编译单位
	tree := p.TranslationUnit()

	//	创建listener
	listener := newCppAstParserPool.Get().(*CppAstParser)
	defer newCppAstParserPool.Put(listener)
	// 初始化置空
	listener.AstInfo = AstInfo{}
	//	执行分析
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.AstInfo
}

// EnterClassSpecifier is called when production classSpecifier is entered.
func (s *CppAstParser) EnterClassSpecifier(ctx *cpp.ClassSpecifierContext) {
	if ctx.ClassHead() == nil {
		return
	}
	var classInfo = Class{
		StartLine: ctx.GetStart().GetLine(),
		EndLine:   ctx.GetStop().GetLine(),
		Name:      findCppDeclChain(ctx) + ctx.ClassHead().(*cpp.ClassHeadContext).ClassHeadName().GetText(),
		Extends:   getCppClassExtends(ctx),
	}
	util.ForDebug()
	s.AstInfo.Classes = append(s.AstInfo.Classes, classInfo)
}

func getCppClassExtends(ctx *cpp.ClassSpecifierContext) (extends string) {
	if ctx.ClassHead().(*cpp.ClassHeadContext).BaseClause() != nil {
		baseClause := ctx.ClassHead().(*cpp.ClassHeadContext).BaseClause().(*cpp.BaseClauseContext)
		if baseClause.BaseSpecifierList().(*cpp.BaseSpecifierListContext).AllBaseSpecifier() != nil {
			allBaseSpecifier := baseClause.BaseSpecifierList().(*cpp.BaseSpecifierListContext).AllBaseSpecifier()
			for index, item := range allBaseSpecifier {
				extends += item.(*cpp.BaseSpecifierContext).AccessSpecifier().GetText() + " "
				extends += item.(*cpp.BaseSpecifierContext).BaseTypeSpecifier().GetText()
				if index != len(allBaseSpecifier)-1 {
					extends += ", "
				}
			}
		}
	}
	return
}

func findCppDeclChain(ctx antlr.ParseTree) (chain string) {
	tempCtx := ctx.GetParent()
	for {
		if _, ok := tempCtx.(*cpp.ClassSpecifierContext); ok {
			chain = tempCtx.(*cpp.ClassSpecifierContext).ClassHead().(*cpp.ClassHeadContext).ClassHeadName().GetText() + "." + chain
		}
		if _, ok := tempCtx.(*cpp.FunctionDefinitionContext); ok {
			chain = matchMethodName(tempCtx.(*cpp.FunctionDefinitionContext)) + "." + chain
		}
		tempCtx = tempCtx.GetParent()
		if tempCtx == nil {
			return
		}
	}
}

func matchMethodName(ctx *cpp.FunctionDefinitionContext) (methodName string) {
	if ctx.Declarator() != nil {
		functionDefineStr := ctx.Declarator().GetText()
		rightIndex := strings.Index(functionDefineStr, "(")
		if rightIndex != -1 {
			spIndex := strings.Index(functionDefineStr, "&")
			if spIndex != -1 && spIndex < rightIndex {
				methodName = functionDefineStr[spIndex+1 : rightIndex]
			} else {
				methodName = functionDefineStr[:rightIndex]
			}
		}

	}
	return
}
func matchMethodParams(ctx *cpp.FunctionDefinitionContext) (params string) {
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
func (s *CppAstParser) ExitFunctionDefinition(ctx *cpp.FunctionDefinitionContext) {
	// ! c++无法准确解析到函数
	if ctx.Declarator() == nil {
		return
	}
	methodInfo := Method{
		StartLine:  ctx.GetStart().GetLine(),
		EndLine:    ctx.GetStop().GetLine(),
		Name:       findCppDeclChain(ctx) + matchMethodName(ctx),
		Parameters: matchMethodParams(ctx),
	}
	// TODO 没有对参数进行区分
	s.AstInfo.Methods = append(s.AstInfo.Methods, methodInfo)
}

// EnterNoPointerDeclarator is called when production noPointerDeclarator is entered.
func (s *CppAstParser) EnterNoPointerDeclarator(ctx *cpp.NoPointerDeclaratorContext) {}

// ExitNoPointerDeclarator is called when production noPointerDeclarator is exited.
func (s *CppAstParser) ExitNoPointerDeclarator(ctx *cpp.NoPointerDeclaratorContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *CppAstParser) EnterFunctionDefinition(ctx *cpp.FunctionDefinitionContext) {}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *CppAstParser) EnterFunctionSpecifier(ctx *cpp.FunctionSpecifierContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *CppAstParser) EnterExpressionStatement(ctx *cpp.ExpressionStatementContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *CppAstParser) EnterPostfixExpression(ctx *cpp.PostfixExpressionContext) {}

func (s *CppAstParser) EnterSimpleDeclaration(ctx *cpp.SimpleDeclarationContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *CppAstParser) ExitExpressionStatement(ctx *cpp.ExpressionStatementContext) {}

// VisitTerminal is called when a terminal node is visited.
func (s *CppAstParser) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *CppAstParser) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *CppAstParser) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *CppAstParser) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *CppAstParser) EnterTranslationUnit(ctx *cpp.TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *CppAstParser) ExitTranslationUnit(ctx *cpp.TranslationUnitContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *CppAstParser) EnterPrimaryExpression(ctx *cpp.PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *CppAstParser) ExitPrimaryExpression(ctx *cpp.PrimaryExpressionContext) {}

// EnterIdExpression is called when production idExpression is entered.
func (s *CppAstParser) EnterIdExpression(ctx *cpp.IdExpressionContext) {}

// ExitIdExpression is called when production idExpression is exited.
func (s *CppAstParser) ExitIdExpression(ctx *cpp.IdExpressionContext) {}

// EnterUnqualifiedId is called when production unqualifiedId is entered.
func (s *CppAstParser) EnterUnqualifiedId(ctx *cpp.UnqualifiedIdContext) {}

// ExitUnqualifiedId is called when production unqualifiedId is exited.
func (s *CppAstParser) ExitUnqualifiedId(ctx *cpp.UnqualifiedIdContext) {}

// EnterQualifiedId is called when production qualifiedId is entered.
func (s *CppAstParser) EnterQualifiedId(ctx *cpp.QualifiedIdContext) {}

// ExitQualifiedId is called when production qualifiedId is exited.
func (s *CppAstParser) ExitQualifiedId(ctx *cpp.QualifiedIdContext) {}

// EnterNestedNameSpecifier is called when production nestedNameSpecifier is entered.
func (s *CppAstParser) EnterNestedNameSpecifier(ctx *cpp.NestedNameSpecifierContext) {}

// ExitNestedNameSpecifier is called when production nestedNameSpecifier is exited.
func (s *CppAstParser) ExitNestedNameSpecifier(ctx *cpp.NestedNameSpecifierContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *CppAstParser) EnterLambdaExpression(ctx *cpp.LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *CppAstParser) ExitLambdaExpression(ctx *cpp.LambdaExpressionContext) {}

// EnterLambdaIntroducer is called when production lambdaIntroducer is entered.
func (s *CppAstParser) EnterLambdaIntroducer(ctx *cpp.LambdaIntroducerContext) {}

// ExitLambdaIntroducer is called when production lambdaIntroducer is exited.
func (s *CppAstParser) ExitLambdaIntroducer(ctx *cpp.LambdaIntroducerContext) {}

// EnterLambdaCapture is called when production lambdaCapture is entered.
func (s *CppAstParser) EnterLambdaCapture(ctx *cpp.LambdaCaptureContext) {}

// ExitLambdaCapture is called when production lambdaCapture is exited.
func (s *CppAstParser) ExitLambdaCapture(ctx *cpp.LambdaCaptureContext) {}

// EnterCaptureDefault is called when production captureDefault is entered.
func (s *CppAstParser) EnterCaptureDefault(ctx *cpp.CaptureDefaultContext) {}

// ExitCaptureDefault is called when production captureDefault is exited.
func (s *CppAstParser) ExitCaptureDefault(ctx *cpp.CaptureDefaultContext) {}

// EnterCaptureList is called when production captureList is entered.
func (s *CppAstParser) EnterCaptureList(ctx *cpp.CaptureListContext) {}

// ExitCaptureList is called when production captureList is exited.
func (s *CppAstParser) ExitCaptureList(ctx *cpp.CaptureListContext) {}

// EnterCapture is called when production capture is entered.
func (s *CppAstParser) EnterCapture(ctx *cpp.CaptureContext) {}

// ExitCapture is called when production capture is exited.
func (s *CppAstParser) ExitCapture(ctx *cpp.CaptureContext) {}

// EnterSimpleCapture is called when production simpleCapture is entered.
func (s *CppAstParser) EnterSimpleCapture(ctx *cpp.SimpleCaptureContext) {}

// ExitSimpleCapture is called when production simpleCapture is exited.
func (s *CppAstParser) ExitSimpleCapture(ctx *cpp.SimpleCaptureContext) {}

// EnterInitcapture is called when production initcapture is entered.
func (s *CppAstParser) EnterInitcapture(ctx *cpp.InitcaptureContext) {}

// ExitInitcapture is called when production initcapture is exited.
func (s *CppAstParser) ExitInitcapture(ctx *cpp.InitcaptureContext) {}

// EnterLambdaDeclarator is called when production lambdaDeclarator is entered.
func (s *CppAstParser) EnterLambdaDeclarator(ctx *cpp.LambdaDeclaratorContext) {}

// ExitLambdaDeclarator is called when production lambdaDeclarator is exited.
func (s *CppAstParser) ExitLambdaDeclarator(ctx *cpp.LambdaDeclaratorContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *CppAstParser) ExitPostfixExpression(ctx *cpp.PostfixExpressionContext) {
}

// EnterTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is entered.
func (s *CppAstParser) EnterTypeIdOfTheTypeId(ctx *cpp.TypeIdOfTheTypeIdContext) {}

// ExitTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is exited.
func (s *CppAstParser) ExitTypeIdOfTheTypeId(ctx *cpp.TypeIdOfTheTypeIdContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *CppAstParser) EnterExpressionList(ctx *cpp.ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *CppAstParser) ExitExpressionList(ctx *cpp.ExpressionListContext) {}

// EnterPseudoDestructorName is called when production pseudoDestructorName is entered.
func (s *CppAstParser) EnterPseudoDestructorName(ctx *cpp.PseudoDestructorNameContext) {}

// ExitPseudoDestructorName is called when production pseudoDestructorName is exited.
func (s *CppAstParser) ExitPseudoDestructorName(ctx *cpp.PseudoDestructorNameContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *CppAstParser) EnterUnaryExpression(ctx *cpp.UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *CppAstParser) ExitUnaryExpression(ctx *cpp.UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *CppAstParser) EnterUnaryOperator(ctx *cpp.UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *CppAstParser) ExitUnaryOperator(ctx *cpp.UnaryOperatorContext) {}

// EnterNewExpression is called when production newExpression is entered.
func (s *CppAstParser) EnterNewExpression(ctx *cpp.NewExpressionContext) {}

// ExitNewExpression is called when production newExpression is exited.
func (s *CppAstParser) ExitNewExpression(ctx *cpp.NewExpressionContext) {}

// EnterNewPlacement is called when production newPlacement is entered.
func (s *CppAstParser) EnterNewPlacement(ctx *cpp.NewPlacementContext) {}

// ExitNewPlacement is called when production newPlacement is exited.
func (s *CppAstParser) ExitNewPlacement(ctx *cpp.NewPlacementContext) {}

// EnterNewTypeId is called when production newTypeId is entered.
func (s *CppAstParser) EnterNewTypeId(ctx *cpp.NewTypeIdContext) {}

// ExitNewTypeId is called when production newTypeId is exited.
func (s *CppAstParser) ExitNewTypeId(ctx *cpp.NewTypeIdContext) {}

// EnterNewDeclarator is called when production newDeclarator is entered.
func (s *CppAstParser) EnterNewDeclarator(ctx *cpp.NewDeclaratorContext) {}

// ExitNewDeclarator is called when production newDeclarator is exited.
func (s *CppAstParser) ExitNewDeclarator(ctx *cpp.NewDeclaratorContext) {}

// EnterNoPointerNewDeclarator is called when production noPointerNewDeclarator is entered.
func (s *CppAstParser) EnterNoPointerNewDeclarator(ctx *cpp.NoPointerNewDeclaratorContext) {}

// ExitNoPointerNewDeclarator is called when production noPointerNewDeclarator is exited.
func (s *CppAstParser) ExitNoPointerNewDeclarator(ctx *cpp.NoPointerNewDeclaratorContext) {}

// EnterNewInitializer is called when production newInitializer is entered.
func (s *CppAstParser) EnterNewInitializer(ctx *cpp.NewInitializerContext) {}

// ExitNewInitializer is called when production newInitializer is exited.
func (s *CppAstParser) ExitNewInitializer(ctx *cpp.NewInitializerContext) {}

// EnterDeleteExpression is called when production deleteExpression is entered.
func (s *CppAstParser) EnterDeleteExpression(ctx *cpp.DeleteExpressionContext) {}

// ExitDeleteExpression is called when production deleteExpression is exited.
func (s *CppAstParser) ExitDeleteExpression(ctx *cpp.DeleteExpressionContext) {}

// EnterNoExceptExpression is called when production noExceptExpression is entered.
func (s *CppAstParser) EnterNoExceptExpression(ctx *cpp.NoExceptExpressionContext) {}

// ExitNoExceptExpression is called when production noExceptExpression is exited.
func (s *CppAstParser) ExitNoExceptExpression(ctx *cpp.NoExceptExpressionContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *CppAstParser) EnterCastExpression(ctx *cpp.CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *CppAstParser) ExitCastExpression(ctx *cpp.CastExpressionContext) {}

// EnterPointerMemberExpression is called when production pointerMemberExpression is entered.
func (s *CppAstParser) EnterPointerMemberExpression(ctx *cpp.PointerMemberExpressionContext) {
}

// ExitPointerMemberExpression is called when production pointerMemberExpression is exited.
func (s *CppAstParser) ExitPointerMemberExpression(ctx *cpp.PointerMemberExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *CppAstParser) EnterMultiplicativeExpression(ctx *cpp.MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *CppAstParser) ExitMultiplicativeExpression(ctx *cpp.MultiplicativeExpressionContext) {
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *CppAstParser) EnterAdditiveExpression(ctx *cpp.AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *CppAstParser) ExitAdditiveExpression(ctx *cpp.AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *CppAstParser) EnterShiftExpression(ctx *cpp.ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *CppAstParser) ExitShiftExpression(ctx *cpp.ShiftExpressionContext) {}

// EnterShiftOperator is called when production shiftOperator is entered.
func (s *CppAstParser) EnterShiftOperator(ctx *cpp.ShiftOperatorContext) {}

// ExitShiftOperator is called when production shiftOperator is exited.
func (s *CppAstParser) ExitShiftOperator(ctx *cpp.ShiftOperatorContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *CppAstParser) EnterRelationalExpression(ctx *cpp.RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *CppAstParser) ExitRelationalExpression(ctx *cpp.RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *CppAstParser) EnterEqualityExpression(ctx *cpp.EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *CppAstParser) ExitEqualityExpression(ctx *cpp.EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *CppAstParser) EnterAndExpression(ctx *cpp.AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *CppAstParser) ExitAndExpression(ctx *cpp.AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *CppAstParser) EnterExclusiveOrExpression(ctx *cpp.ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *CppAstParser) ExitExclusiveOrExpression(ctx *cpp.ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *CppAstParser) EnterInclusiveOrExpression(ctx *cpp.InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *CppAstParser) ExitInclusiveOrExpression(ctx *cpp.InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *CppAstParser) EnterLogicalAndExpression(ctx *cpp.LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *CppAstParser) ExitLogicalAndExpression(ctx *cpp.LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *CppAstParser) EnterLogicalOrExpression(ctx *cpp.LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *CppAstParser) ExitLogicalOrExpression(ctx *cpp.LogicalOrExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *CppAstParser) EnterConditionalExpression(ctx *cpp.ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *CppAstParser) ExitConditionalExpression(ctx *cpp.ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *CppAstParser) EnterAssignmentExpression(ctx *cpp.AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *CppAstParser) ExitAssignmentExpression(ctx *cpp.AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *CppAstParser) EnterAssignmentOperator(ctx *cpp.AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *CppAstParser) ExitAssignmentOperator(ctx *cpp.AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *CppAstParser) EnterExpression(ctx *cpp.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *CppAstParser) ExitExpression(ctx *cpp.ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *CppAstParser) EnterConstantExpression(ctx *cpp.ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *CppAstParser) ExitConstantExpression(ctx *cpp.ConstantExpressionContext) {}

// EnterStatement is called when production statement is entered.
func (s *CppAstParser) EnterStatement(ctx *cpp.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *CppAstParser) ExitStatement(ctx *cpp.StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *CppAstParser) EnterLabeledStatement(ctx *cpp.LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *CppAstParser) ExitLabeledStatement(ctx *cpp.LabeledStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *CppAstParser) EnterCompoundStatement(ctx *cpp.CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *CppAstParser) ExitCompoundStatement(ctx *cpp.CompoundStatementContext) {}

// EnterStatementSeq is called when production statementSeq is entered.
func (s *CppAstParser) EnterStatementSeq(ctx *cpp.StatementSeqContext) {}

// ExitStatementSeq is called when production statementSeq is exited.
func (s *CppAstParser) ExitStatementSeq(ctx *cpp.StatementSeqContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *CppAstParser) EnterSelectionStatement(ctx *cpp.SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *CppAstParser) ExitSelectionStatement(ctx *cpp.SelectionStatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *CppAstParser) EnterCondition(ctx *cpp.ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *CppAstParser) ExitCondition(ctx *cpp.ConditionContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *CppAstParser) EnterIterationStatement(ctx *cpp.IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *CppAstParser) ExitIterationStatement(ctx *cpp.IterationStatementContext) {}

// EnterForInitStatement is called when production forInitStatement is entered.
func (s *CppAstParser) EnterForInitStatement(ctx *cpp.ForInitStatementContext) {}

// ExitForInitStatement is called when production forInitStatement is exited.
func (s *CppAstParser) ExitForInitStatement(ctx *cpp.ForInitStatementContext) {}

// EnterForRangeDeclaration is called when production forRangeDeclaration is entered.
func (s *CppAstParser) EnterForRangeDeclaration(ctx *cpp.ForRangeDeclarationContext) {}

// ExitForRangeDeclaration is called when production forRangeDeclaration is exited.
func (s *CppAstParser) ExitForRangeDeclaration(ctx *cpp.ForRangeDeclarationContext) {}

// EnterForRangeInitializer is called when production forRangeInitializer is entered.
func (s *CppAstParser) EnterForRangeInitializer(ctx *cpp.ForRangeInitializerContext) {}

// ExitForRangeInitializer is called when production forRangeInitializer is exited.
func (s *CppAstParser) ExitForRangeInitializer(ctx *cpp.ForRangeInitializerContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *CppAstParser) EnterJumpStatement(ctx *cpp.JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *CppAstParser) ExitJumpStatement(ctx *cpp.JumpStatementContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *CppAstParser) EnterDeclarationStatement(ctx *cpp.DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *CppAstParser) ExitDeclarationStatement(ctx *cpp.DeclarationStatementContext) {}

// EnterDeclarationseq is called when production declarationseq is entered.
func (s *CppAstParser) EnterDeclarationseq(ctx *cpp.DeclarationseqContext) {}

// ExitDeclarationseq is called when production declarationseq is exited.
func (s *CppAstParser) ExitDeclarationseq(ctx *cpp.DeclarationseqContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *CppAstParser) EnterDeclaration(ctx *cpp.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *CppAstParser) ExitDeclaration(ctx *cpp.DeclarationContext) {}

// EnterBlockDeclaration is called when production blockDeclaration is entered.
func (s *CppAstParser) EnterBlockDeclaration(ctx *cpp.BlockDeclarationContext) {}

// ExitBlockDeclaration is called when production blockDeclaration is exited.
func (s *CppAstParser) ExitBlockDeclaration(ctx *cpp.BlockDeclarationContext) {}

// EnterAliasDeclaration is called when production aliasDeclaration is entered.
func (s *CppAstParser) EnterAliasDeclaration(ctx *cpp.AliasDeclarationContext) {}

// ExitAliasDeclaration is called when production aliasDeclaration is exited.
func (s *CppAstParser) ExitAliasDeclaration(ctx *cpp.AliasDeclarationContext) {}

//// EnterSimpleDeclaration is called when production simpleDeclaration is entered.
//func (s *CppTreeShapeListener) EnterSimpleDeclaration(ctx *cpp.SimpleDeclarationContext) {}

// ExitSimpleDeclaration is called when production simpleDeclaration is exited.
func (s *CppAstParser) ExitSimpleDeclaration(ctx *cpp.SimpleDeclarationContext) {}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *CppAstParser) EnterStaticAssertDeclaration(ctx *cpp.StaticAssertDeclarationContext) {
}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *CppAstParser) ExitStaticAssertDeclaration(ctx *cpp.StaticAssertDeclarationContext) {}

// EnterEmptyDeclaration is called when production emptyDeclaration is entered.
func (s *CppAstParser) EnterEmptyDeclaration(ctx *cpp.EmptyDeclarationContext) {}

// ExitEmptyDeclaration is called when production emptyDeclaration is exited.
func (s *CppAstParser) ExitEmptyDeclaration(ctx *cpp.EmptyDeclarationContext) {}

// EnterAttributeDeclaration is called when production attributeDeclaration is entered.
func (s *CppAstParser) EnterAttributeDeclaration(ctx *cpp.AttributeDeclarationContext) {}

// ExitAttributeDeclaration is called when production attributeDeclaration is exited.
func (s *CppAstParser) ExitAttributeDeclaration(ctx *cpp.AttributeDeclarationContext) {}

// EnterDeclSpecifier is called when production declSpecifier is entered.
func (s *CppAstParser) EnterDeclSpecifier(ctx *cpp.DeclSpecifierContext) {}

// ExitDeclSpecifier is called when production declSpecifier is exited.
func (s *CppAstParser) ExitDeclSpecifier(ctx *cpp.DeclSpecifierContext) {}

// EnterDeclSpecifierSeq is called when production declSpecifierSeq is entered.
func (s *CppAstParser) EnterDeclSpecifierSeq(ctx *cpp.DeclSpecifierSeqContext) {}

// ExitDeclSpecifierSeq is called when production declSpecifierSeq is exited.
func (s *CppAstParser) ExitDeclSpecifierSeq(ctx *cpp.DeclSpecifierSeqContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *CppAstParser) EnterStorageClassSpecifier(ctx *cpp.StorageClassSpecifierContext) {}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *CppAstParser) ExitStorageClassSpecifier(ctx *cpp.StorageClassSpecifierContext) {}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *CppAstParser) ExitFunctionSpecifier(ctx *cpp.FunctionSpecifierContext) {}

// EnterTypedefName is called when production typedefName is entered.
func (s *CppAstParser) EnterTypedefName(ctx *cpp.TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *CppAstParser) ExitTypedefName(ctx *cpp.TypedefNameContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *CppAstParser) EnterTypeSpecifier(ctx *cpp.TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *CppAstParser) ExitTypeSpecifier(ctx *cpp.TypeSpecifierContext) {}

// EnterTrailingTypeSpecifier is called when production trailingTypeSpecifier is entered.
func (s *CppAstParser) EnterTrailingTypeSpecifier(ctx *cpp.TrailingTypeSpecifierContext) {}

// ExitTrailingTypeSpecifier is called when production trailingTypeSpecifier is exited.
func (s *CppAstParser) ExitTrailingTypeSpecifier(ctx *cpp.TrailingTypeSpecifierContext) {}

// EnterTypeSpecifierSeq is called when production typeSpecifierSeq is entered.
func (s *CppAstParser) EnterTypeSpecifierSeq(ctx *cpp.TypeSpecifierSeqContext) {}

// ExitTypeSpecifierSeq is called when production typeSpecifierSeq is exited.
func (s *CppAstParser) ExitTypeSpecifierSeq(ctx *cpp.TypeSpecifierSeqContext) {}

// EnterTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is entered.
func (s *CppAstParser) EnterTrailingTypeSpecifierSeq(ctx *cpp.TrailingTypeSpecifierSeqContext) {
}

// ExitTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is exited.
func (s *CppAstParser) ExitTrailingTypeSpecifierSeq(ctx *cpp.TrailingTypeSpecifierSeqContext) {
}

// EnterSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is entered.
func (s *CppAstParser) EnterSimpleTypeLengthModifier(ctx *cpp.SimpleTypeLengthModifierContext) {
}

// ExitSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is exited.
func (s *CppAstParser) ExitSimpleTypeLengthModifier(ctx *cpp.SimpleTypeLengthModifierContext) {
}

// EnterSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is entered.
func (s *CppAstParser) EnterSimpleTypeSignednessModifier(ctx *cpp.SimpleTypeSignednessModifierContext) {
}

// ExitSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is exited.
func (s *CppAstParser) ExitSimpleTypeSignednessModifier(ctx *cpp.SimpleTypeSignednessModifierContext) {
}

// EnterSimpleTypeSpecifier is called when production simpleTypeSpecifier is entered.
func (s *CppAstParser) EnterSimpleTypeSpecifier(ctx *cpp.SimpleTypeSpecifierContext) {}

// ExitSimpleTypeSpecifier is called when production simpleTypeSpecifier is exited.
func (s *CppAstParser) ExitSimpleTypeSpecifier(ctx *cpp.SimpleTypeSpecifierContext) {}

// EnterTheTypeName is called when production theTypeName is entered.
func (s *CppAstParser) EnterTheTypeName(ctx *cpp.TheTypeNameContext) {}

// ExitTheTypeName is called when production theTypeName is exited.
func (s *CppAstParser) ExitTheTypeName(ctx *cpp.TheTypeNameContext) {}

// EnterDecltypeSpecifier is called when production decltypeSpecifier is entered.
func (s *CppAstParser) EnterDecltypeSpecifier(ctx *cpp.DecltypeSpecifierContext) {}

// ExitDecltypeSpecifier is called when production decltypeSpecifier is exited.
func (s *CppAstParser) ExitDecltypeSpecifier(ctx *cpp.DecltypeSpecifierContext) {}

// EnterElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is entered.
func (s *CppAstParser) EnterElaboratedTypeSpecifier(ctx *cpp.ElaboratedTypeSpecifierContext) {
}

// ExitElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is exited.
func (s *CppAstParser) ExitElaboratedTypeSpecifier(ctx *cpp.ElaboratedTypeSpecifierContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *CppAstParser) EnterEnumName(ctx *cpp.EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *CppAstParser) ExitEnumName(ctx *cpp.EnumNameContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *CppAstParser) EnterEnumSpecifier(ctx *cpp.EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *CppAstParser) ExitEnumSpecifier(ctx *cpp.EnumSpecifierContext) {}

// EnterEnumHead is called when production enumHead is entered.
func (s *CppAstParser) EnterEnumHead(ctx *cpp.EnumHeadContext) {}

// ExitEnumHead is called when production enumHead is exited.
func (s *CppAstParser) ExitEnumHead(ctx *cpp.EnumHeadContext) {}

// EnterOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is entered.
func (s *CppAstParser) EnterOpaqueEnumDeclaration(ctx *cpp.OpaqueEnumDeclarationContext) {}

// ExitOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is exited.
func (s *CppAstParser) ExitOpaqueEnumDeclaration(ctx *cpp.OpaqueEnumDeclarationContext) {}

// EnterEnumkey is called when production enumkey is entered.
func (s *CppAstParser) EnterEnumkey(ctx *cpp.EnumkeyContext) {}

// ExitEnumkey is called when production enumkey is exited.
func (s *CppAstParser) ExitEnumkey(ctx *cpp.EnumkeyContext) {}

// EnterEnumbase is called when production enumbase is entered.
func (s *CppAstParser) EnterEnumbase(ctx *cpp.EnumbaseContext) {}

// ExitEnumbase is called when production enumbase is exited.
func (s *CppAstParser) ExitEnumbase(ctx *cpp.EnumbaseContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *CppAstParser) EnterEnumeratorList(ctx *cpp.EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *CppAstParser) ExitEnumeratorList(ctx *cpp.EnumeratorListContext) {}

// EnterEnumeratorDefinition is called when production enumeratorDefinition is entered.
func (s *CppAstParser) EnterEnumeratorDefinition(ctx *cpp.EnumeratorDefinitionContext) {}

// ExitEnumeratorDefinition is called when production enumeratorDefinition is exited.
func (s *CppAstParser) ExitEnumeratorDefinition(ctx *cpp.EnumeratorDefinitionContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *CppAstParser) EnterEnumerator(ctx *cpp.EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *CppAstParser) ExitEnumerator(ctx *cpp.EnumeratorContext) {}

// EnterNamespaceName is called when production namespaceName is entered.
func (s *CppAstParser) EnterNamespaceName(ctx *cpp.NamespaceNameContext) {}

// ExitNamespaceName is called when production namespaceName is exited.
func (s *CppAstParser) ExitNamespaceName(ctx *cpp.NamespaceNameContext) {}

// EnterOriginalNamespaceName is called when production originalNamespaceName is entered.
func (s *CppAstParser) EnterOriginalNamespaceName(ctx *cpp.OriginalNamespaceNameContext) {}

// ExitOriginalNamespaceName is called when production originalNamespaceName is exited.
func (s *CppAstParser) ExitOriginalNamespaceName(ctx *cpp.OriginalNamespaceNameContext) {}

// EnterNamespaceDefinition is called when production namespaceDefinition is entered.
func (s *CppAstParser) EnterNamespaceDefinition(ctx *cpp.NamespaceDefinitionContext) {}

// ExitNamespaceDefinition is called when production namespaceDefinition is exited.
func (s *CppAstParser) ExitNamespaceDefinition(ctx *cpp.NamespaceDefinitionContext) {}

// EnterNamespaceAlias is called when production namespaceAlias is entered.
func (s *CppAstParser) EnterNamespaceAlias(ctx *cpp.NamespaceAliasContext) {}

// ExitNamespaceAlias is called when production namespaceAlias is exited.
func (s *CppAstParser) ExitNamespaceAlias(ctx *cpp.NamespaceAliasContext) {}

// EnterNamespaceAliasDefinition is called when production namespaceAliasDefinition is entered.
func (s *CppAstParser) EnterNamespaceAliasDefinition(ctx *cpp.NamespaceAliasDefinitionContext) {
}

// ExitNamespaceAliasDefinition is called when production namespaceAliasDefinition is exited.
func (s *CppAstParser) ExitNamespaceAliasDefinition(ctx *cpp.NamespaceAliasDefinitionContext) {
}

// EnterQualifiednamespacespecifier is called when production qualifiednamespacespecifier is entered.
func (s *CppAstParser) EnterQualifiednamespacespecifier(ctx *cpp.QualifiednamespacespecifierContext) {
}

// ExitQualifiednamespacespecifier is called when production qualifiednamespacespecifier is exited.
func (s *CppAstParser) ExitQualifiednamespacespecifier(ctx *cpp.QualifiednamespacespecifierContext) {
}

// EnterUsingDeclaration is called when production usingDeclaration is entered.
func (s *CppAstParser) EnterUsingDeclaration(ctx *cpp.UsingDeclarationContext) {}

// ExitUsingDeclaration is called when production usingDeclaration is exited.
func (s *CppAstParser) ExitUsingDeclaration(ctx *cpp.UsingDeclarationContext) {}

// EnterUsingDirective is called when production usingDirective is entered.
func (s *CppAstParser) EnterUsingDirective(ctx *cpp.UsingDirectiveContext) {}

// ExitUsingDirective is called when production usingDirective is exited.
func (s *CppAstParser) ExitUsingDirective(ctx *cpp.UsingDirectiveContext) {}

// EnterAsmDefinition is called when production asmDefinition is entered.
func (s *CppAstParser) EnterAsmDefinition(ctx *cpp.AsmDefinitionContext) {}

// ExitAsmDefinition is called when production asmDefinition is exited.
func (s *CppAstParser) ExitAsmDefinition(ctx *cpp.AsmDefinitionContext) {}

// EnterLinkageSpecification is called when production linkageSpecification is entered.
func (s *CppAstParser) EnterLinkageSpecification(ctx *cpp.LinkageSpecificationContext) {}

// ExitLinkageSpecification is called when production linkageSpecification is exited.
func (s *CppAstParser) ExitLinkageSpecification(ctx *cpp.LinkageSpecificationContext) {}

// EnterAttributeSpecifierSeq is called when production attributeSpecifierSeq is entered.
func (s *CppAstParser) EnterAttributeSpecifierSeq(ctx *cpp.AttributeSpecifierSeqContext) {}

// ExitAttributeSpecifierSeq is called when production attributeSpecifierSeq is exited.
func (s *CppAstParser) ExitAttributeSpecifierSeq(ctx *cpp.AttributeSpecifierSeqContext) {}

// EnterAttributeSpecifier is called when production attributeSpecifier is entered.
func (s *CppAstParser) EnterAttributeSpecifier(ctx *cpp.AttributeSpecifierContext) {}

// ExitAttributeSpecifier is called when production attributeSpecifier is exited.
func (s *CppAstParser) ExitAttributeSpecifier(ctx *cpp.AttributeSpecifierContext) {}

// EnterAlignmentspecifier is called when production alignmentspecifier is entered.
func (s *CppAstParser) EnterAlignmentspecifier(ctx *cpp.AlignmentspecifierContext) {}

// ExitAlignmentspecifier is called when production alignmentspecifier is exited.
func (s *CppAstParser) ExitAlignmentspecifier(ctx *cpp.AlignmentspecifierContext) {}

// EnterAttributeList is called when production attributeList is entered.
func (s *CppAstParser) EnterAttributeList(ctx *cpp.AttributeListContext) {}

// ExitAttributeList is called when production attributeList is exited.
func (s *CppAstParser) ExitAttributeList(ctx *cpp.AttributeListContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *CppAstParser) EnterAttribute(ctx *cpp.AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *CppAstParser) ExitAttribute(ctx *cpp.AttributeContext) {}

// EnterAttributeNamespace is called when production attributeNamespace is entered.
func (s *CppAstParser) EnterAttributeNamespace(ctx *cpp.AttributeNamespaceContext) {}

// ExitAttributeNamespace is called when production attributeNamespace is exited.
func (s *CppAstParser) ExitAttributeNamespace(ctx *cpp.AttributeNamespaceContext) {}

// EnterAttributeArgumentClause is called when production attributeArgumentClause is entered.
func (s *CppAstParser) EnterAttributeArgumentClause(ctx *cpp.AttributeArgumentClauseContext) {
}

// ExitAttributeArgumentClause is called when production attributeArgumentClause is exited.
func (s *CppAstParser) ExitAttributeArgumentClause(ctx *cpp.AttributeArgumentClauseContext) {}

// EnterBalancedTokenSeq is called when production balancedTokenSeq is entered.
func (s *CppAstParser) EnterBalancedTokenSeq(ctx *cpp.BalancedTokenSeqContext) {}

// ExitBalancedTokenSeq is called when production balancedTokenSeq is exited.
func (s *CppAstParser) ExitBalancedTokenSeq(ctx *cpp.BalancedTokenSeqContext) {}

// EnterBalancedtoken is called when production balancedtoken is entered.
func (s *CppAstParser) EnterBalancedtoken(ctx *cpp.BalancedtokenContext) {}

// ExitBalancedtoken is called when production balancedtoken is exited.
func (s *CppAstParser) ExitBalancedtoken(ctx *cpp.BalancedtokenContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *CppAstParser) EnterInitDeclaratorList(ctx *cpp.InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *CppAstParser) ExitInitDeclaratorList(ctx *cpp.InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *CppAstParser) EnterInitDeclarator(ctx *cpp.InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *CppAstParser) ExitInitDeclarator(ctx *cpp.InitDeclaratorContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *CppAstParser) EnterDeclarator(ctx *cpp.DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *CppAstParser) ExitDeclarator(ctx *cpp.DeclaratorContext) {}

// EnterPointerDeclarator is called when production pointerDeclarator is entered.
func (s *CppAstParser) EnterPointerDeclarator(ctx *cpp.PointerDeclaratorContext) {}

// ExitPointerDeclarator is called when production pointerDeclarator is exited.
func (s *CppAstParser) ExitPointerDeclarator(ctx *cpp.PointerDeclaratorContext) {}

// EnterParametersAndQualifiers is called when production parametersAndQualifiers is entered.
func (s *CppAstParser) EnterParametersAndQualifiers(ctx *cpp.ParametersAndQualifiersContext) {
}

// ExitParametersAndQualifiers is called when production parametersAndQualifiers is exited.
func (s *CppAstParser) ExitParametersAndQualifiers(ctx *cpp.ParametersAndQualifiersContext) {}

// EnterTrailingReturnType is called when production trailingReturnType is entered.
func (s *CppAstParser) EnterTrailingReturnType(ctx *cpp.TrailingReturnTypeContext) {}

// ExitTrailingReturnType is called when production trailingReturnType is exited.
func (s *CppAstParser) ExitTrailingReturnType(ctx *cpp.TrailingReturnTypeContext) {}

// EnterPointerOperator is called when production pointerOperator is entered.
func (s *CppAstParser) EnterPointerOperator(ctx *cpp.PointerOperatorContext) {}

// ExitPointerOperator is called when production pointerOperator is exited.
func (s *CppAstParser) ExitPointerOperator(ctx *cpp.PointerOperatorContext) {}

// EnterCvqualifierseq is called when production cvqualifierseq is entered.
func (s *CppAstParser) EnterCvqualifierseq(ctx *cpp.CvqualifierseqContext) {}

// ExitCvqualifierseq is called when production cvqualifierseq is exited.
func (s *CppAstParser) ExitCvqualifierseq(ctx *cpp.CvqualifierseqContext) {}

// EnterCvQualifier is called when production cvQualifier is entered.
func (s *CppAstParser) EnterCvQualifier(ctx *cpp.CvQualifierContext) {}

// ExitCvQualifier is called when production cvQualifier is exited.
func (s *CppAstParser) ExitCvQualifier(ctx *cpp.CvQualifierContext) {}

// EnterRefqualifier is called when production refqualifier is entered.
func (s *CppAstParser) EnterRefqualifier(ctx *cpp.RefqualifierContext) {}

// ExitRefqualifier is called when production refqualifier is exited.
func (s *CppAstParser) ExitRefqualifier(ctx *cpp.RefqualifierContext) {}

// EnterDeclaratorid is called when production declaratorid is entered.
func (s *CppAstParser) EnterDeclaratorid(ctx *cpp.DeclaratoridContext) {}

// ExitDeclaratorid is called when production declaratorid is exited.
func (s *CppAstParser) ExitDeclaratorid(ctx *cpp.DeclaratoridContext) {}

// EnterTheTypeId is called when production theTypeId is entered.
func (s *CppAstParser) EnterTheTypeId(ctx *cpp.TheTypeIdContext) {}

// ExitTheTypeId is called when production theTypeId is exited.
func (s *CppAstParser) ExitTheTypeId(ctx *cpp.TheTypeIdContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *CppAstParser) EnterAbstractDeclarator(ctx *cpp.AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *CppAstParser) ExitAbstractDeclarator(ctx *cpp.AbstractDeclaratorContext) {}

// EnterPointerAbstractDeclarator is called when production pointerAbstractDeclarator is entered.
func (s *CppAstParser) EnterPointerAbstractDeclarator(ctx *cpp.PointerAbstractDeclaratorContext) {
}

// ExitPointerAbstractDeclarator is called when production pointerAbstractDeclarator is exited.
func (s *CppAstParser) ExitPointerAbstractDeclarator(ctx *cpp.PointerAbstractDeclaratorContext) {
}

// EnterNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is entered.
func (s *CppAstParser) EnterNoPointerAbstractDeclarator(ctx *cpp.NoPointerAbstractDeclaratorContext) {
}

// ExitNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is exited.
func (s *CppAstParser) ExitNoPointerAbstractDeclarator(ctx *cpp.NoPointerAbstractDeclaratorContext) {
}

// EnterAbstractPackDeclarator is called when production abstractPackDeclarator is entered.
func (s *CppAstParser) EnterAbstractPackDeclarator(ctx *cpp.AbstractPackDeclaratorContext) {}

// ExitAbstractPackDeclarator is called when production abstractPackDeclarator is exited.
func (s *CppAstParser) ExitAbstractPackDeclarator(ctx *cpp.AbstractPackDeclaratorContext) {}

// EnterNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is entered.
func (s *CppAstParser) EnterNoPointerAbstractPackDeclarator(ctx *cpp.NoPointerAbstractPackDeclaratorContext) {
}

// ExitNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is exited.
func (s *CppAstParser) ExitNoPointerAbstractPackDeclarator(ctx *cpp.NoPointerAbstractPackDeclaratorContext) {
}

// EnterParameterDeclarationClause is called when production parameterDeclarationClause is entered.
func (s *CppAstParser) EnterParameterDeclarationClause(ctx *cpp.ParameterDeclarationClauseContext) {
}

// ExitParameterDeclarationClause is called when production parameterDeclarationClause is exited.
func (s *CppAstParser) ExitParameterDeclarationClause(ctx *cpp.ParameterDeclarationClauseContext) {
}

// EnterParameterDeclarationList is called when production parameterDeclarationList is entered.
func (s *CppAstParser) EnterParameterDeclarationList(ctx *cpp.ParameterDeclarationListContext) {
}

// ExitParameterDeclarationList is called when production parameterDeclarationList is exited.
func (s *CppAstParser) ExitParameterDeclarationList(ctx *cpp.ParameterDeclarationListContext) {
}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *CppAstParser) EnterParameterDeclaration(ctx *cpp.ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *CppAstParser) ExitParameterDeclaration(ctx *cpp.ParameterDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *CppAstParser) EnterFunctionBody(ctx *cpp.FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *CppAstParser) ExitFunctionBody(ctx *cpp.FunctionBodyContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *CppAstParser) EnterInitializer(ctx *cpp.InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *CppAstParser) ExitInitializer(ctx *cpp.InitializerContext) {}

// EnterBraceOrEqualInitializer is called when production braceOrEqualInitializer is entered.
func (s *CppAstParser) EnterBraceOrEqualInitializer(ctx *cpp.BraceOrEqualInitializerContext) {
}

// ExitBraceOrEqualInitializer is called when production braceOrEqualInitializer is exited.
func (s *CppAstParser) ExitBraceOrEqualInitializer(ctx *cpp.BraceOrEqualInitializerContext) {}

// EnterInitializerClause is called when production initializerClause is entered.
func (s *CppAstParser) EnterInitializerClause(ctx *cpp.InitializerClauseContext) {}

// ExitInitializerClause is called when production initializerClause is exited.
func (s *CppAstParser) ExitInitializerClause(ctx *cpp.InitializerClauseContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *CppAstParser) EnterInitializerList(ctx *cpp.InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *CppAstParser) ExitInitializerList(ctx *cpp.InitializerListContext) {}

// EnterBracedInitList is called when production bracedInitList is entered.
func (s *CppAstParser) EnterBracedInitList(ctx *cpp.BracedInitListContext) {}

// ExitBracedInitList is called when production bracedInitList is exited.
func (s *CppAstParser) ExitBracedInitList(ctx *cpp.BracedInitListContext) {}

// EnterClassName is called when production className is entered.
func (s *CppAstParser) EnterClassName(ctx *cpp.ClassNameContext) {
	var temp []string
	for _, item := range ctx.GetChildren() {
		temp = append(temp, item.(antlr.ParseTree).GetText())
	}
	util.ForDebug(temp)
}

// ExitClassName is called when production className is exited.
func (s *CppAstParser) ExitClassName(ctx *cpp.ClassNameContext) {}

// ExitClassSpecifier is called when production classSpecifier is exited.
func (s *CppAstParser) ExitClassSpecifier(ctx *cpp.ClassSpecifierContext) {}

// EnterClassHead is called when production classHead is entered.
func (s *CppAstParser) EnterClassHead(ctx *cpp.ClassHeadContext) {
	var temp []string
	for _, item := range ctx.GetChildren() {
		temp = append(temp, item.(antlr.ParseTree).GetText())
	}
	util.ForDebug(temp)
}

// ExitClassHead is called when production classHead is exited.
func (s *CppAstParser) ExitClassHead(ctx *cpp.ClassHeadContext) {}

// EnterClassHeadName is called when production classHeadName is entered.
func (s *CppAstParser) EnterClassHeadName(ctx *cpp.ClassHeadNameContext) {}

// ExitClassHeadName is called when production classHeadName is exited.
func (s *CppAstParser) ExitClassHeadName(ctx *cpp.ClassHeadNameContext) {}

// EnterClassVirtSpecifier is called when production classVirtSpecifier is entered.
func (s *CppAstParser) EnterClassVirtSpecifier(ctx *cpp.ClassVirtSpecifierContext) {}

// ExitClassVirtSpecifier is called when production classVirtSpecifier is exited.
func (s *CppAstParser) ExitClassVirtSpecifier(ctx *cpp.ClassVirtSpecifierContext) {}

// EnterClassKey is called when production classKey is entered.
func (s *CppAstParser) EnterClassKey(ctx *cpp.ClassKeyContext) {}

// ExitClassKey is called when production classKey is exited.
func (s *CppAstParser) ExitClassKey(ctx *cpp.ClassKeyContext) {}

// EnterMemberSpecification is called when production memberSpecification is entered.
func (s *CppAstParser) EnterMemberSpecification(ctx *cpp.MemberSpecificationContext) {}

// ExitMemberSpecification is called when production memberSpecification is exited.
func (s *CppAstParser) ExitMemberSpecification(ctx *cpp.MemberSpecificationContext) {}

// EnterMemberdeclaration is called when production memberdeclaration is entered.
func (s *CppAstParser) EnterMemberdeclaration(ctx *cpp.MemberdeclarationContext) {}

// ExitMemberdeclaration is called when production memberdeclaration is exited.
func (s *CppAstParser) ExitMemberdeclaration(ctx *cpp.MemberdeclarationContext) {}

// EnterMemberDeclaratorList is called when production memberDeclaratorList is entered.
func (s *CppAstParser) EnterMemberDeclaratorList(ctx *cpp.MemberDeclaratorListContext) {}

// ExitMemberDeclaratorList is called when production memberDeclaratorList is exited.
func (s *CppAstParser) ExitMemberDeclaratorList(ctx *cpp.MemberDeclaratorListContext) {}

// EnterMemberDeclarator is called when production memberDeclarator is entered.
func (s *CppAstParser) EnterMemberDeclarator(ctx *cpp.MemberDeclaratorContext) {}

// ExitMemberDeclarator is called when production memberDeclarator is exited.
func (s *CppAstParser) ExitMemberDeclarator(ctx *cpp.MemberDeclaratorContext) {}

// EnterVirtualSpecifierSeq is called when production virtualSpecifierSeq is entered.
func (s *CppAstParser) EnterVirtualSpecifierSeq(ctx *cpp.VirtualSpecifierSeqContext) {}

// ExitVirtualSpecifierSeq is called when production virtualSpecifierSeq is exited.
func (s *CppAstParser) ExitVirtualSpecifierSeq(ctx *cpp.VirtualSpecifierSeqContext) {}

// EnterVirtualSpecifier is called when production virtualSpecifier is entered.
func (s *CppAstParser) EnterVirtualSpecifier(ctx *cpp.VirtualSpecifierContext) {}

// ExitVirtualSpecifier is called when production virtualSpecifier is exited.
func (s *CppAstParser) ExitVirtualSpecifier(ctx *cpp.VirtualSpecifierContext) {}

// EnterPureSpecifier is called when production pureSpecifier is entered.
func (s *CppAstParser) EnterPureSpecifier(ctx *cpp.PureSpecifierContext) {}

// ExitPureSpecifier is called when production pureSpecifier is exited.
func (s *CppAstParser) ExitPureSpecifier(ctx *cpp.PureSpecifierContext) {}

// EnterBaseClause is called when production baseClause is entered.
func (s *CppAstParser) EnterBaseClause(ctx *cpp.BaseClauseContext) {}

// ExitBaseClause is called when production baseClause is exited.
func (s *CppAstParser) ExitBaseClause(ctx *cpp.BaseClauseContext) {}

// EnterBaseSpecifierList is called when production baseSpecifierList is entered.
func (s *CppAstParser) EnterBaseSpecifierList(ctx *cpp.BaseSpecifierListContext) {}

// ExitBaseSpecifierList is called when production baseSpecifierList is exited.
func (s *CppAstParser) ExitBaseSpecifierList(ctx *cpp.BaseSpecifierListContext) {}

// EnterBaseSpecifier is called when production baseSpecifier is entered.
func (s *CppAstParser) EnterBaseSpecifier(ctx *cpp.BaseSpecifierContext) {}

// ExitBaseSpecifier is called when production baseSpecifier is exited.
func (s *CppAstParser) ExitBaseSpecifier(ctx *cpp.BaseSpecifierContext) {}

// EnterClassOrDeclType is called when production classOrDeclType is entered.
func (s *CppAstParser) EnterClassOrDeclType(ctx *cpp.ClassOrDeclTypeContext) {}

// ExitClassOrDeclType is called when production classOrDeclType is exited.
func (s *CppAstParser) ExitClassOrDeclType(ctx *cpp.ClassOrDeclTypeContext) {}

// EnterBaseTypeSpecifier is called when production baseTypeSpecifier is entered.
func (s *CppAstParser) EnterBaseTypeSpecifier(ctx *cpp.BaseTypeSpecifierContext) {}

// ExitBaseTypeSpecifier is called when production baseTypeSpecifier is exited.
func (s *CppAstParser) ExitBaseTypeSpecifier(ctx *cpp.BaseTypeSpecifierContext) {}

// EnterAccessSpecifier is called when production accessSpecifier is entered.
func (s *CppAstParser) EnterAccessSpecifier(ctx *cpp.AccessSpecifierContext) {}

// ExitAccessSpecifier is called when production accessSpecifier is exited.
func (s *CppAstParser) ExitAccessSpecifier(ctx *cpp.AccessSpecifierContext) {}

// EnterConversionFunctionId is called when production conversionFunctionId is entered.
func (s *CppAstParser) EnterConversionFunctionId(ctx *cpp.ConversionFunctionIdContext) {}

// ExitConversionFunctionId is called when production conversionFunctionId is exited.
func (s *CppAstParser) ExitConversionFunctionId(ctx *cpp.ConversionFunctionIdContext) {}

// EnterConversionTypeId is called when production conversionTypeId is entered.
func (s *CppAstParser) EnterConversionTypeId(ctx *cpp.ConversionTypeIdContext) {}

// ExitConversionTypeId is called when production conversionTypeId is exited.
func (s *CppAstParser) ExitConversionTypeId(ctx *cpp.ConversionTypeIdContext) {}

// EnterConversionDeclarator is called when production conversionDeclarator is entered.
func (s *CppAstParser) EnterConversionDeclarator(ctx *cpp.ConversionDeclaratorContext) {}

// ExitConversionDeclarator is called when production conversionDeclarator is exited.
func (s *CppAstParser) ExitConversionDeclarator(ctx *cpp.ConversionDeclaratorContext) {}

// EnterConstructorInitializer is called when production constructorInitializer is entered.
func (s *CppAstParser) EnterConstructorInitializer(ctx *cpp.ConstructorInitializerContext) {}

// ExitConstructorInitializer is called when production constructorInitializer is exited.
func (s *CppAstParser) ExitConstructorInitializer(ctx *cpp.ConstructorInitializerContext) {}

// EnterMemInitializerList is called when production memInitializerList is entered.
func (s *CppAstParser) EnterMemInitializerList(ctx *cpp.MemInitializerListContext) {}

// ExitMemInitializerList is called when production memInitializerList is exited.
func (s *CppAstParser) ExitMemInitializerList(ctx *cpp.MemInitializerListContext) {}

// EnterMemInitializer is called when production memInitializer is entered.
func (s *CppAstParser) EnterMemInitializer(ctx *cpp.MemInitializerContext) {}

// ExitMemInitializer is called when production memInitializer is exited.
func (s *CppAstParser) ExitMemInitializer(ctx *cpp.MemInitializerContext) {}

// EnterMeminitializerid is called when production meminitializerid is entered.
func (s *CppAstParser) EnterMeminitializerid(ctx *cpp.MeminitializeridContext) {}

// ExitMeminitializerid is called when production meminitializerid is exited.
func (s *CppAstParser) ExitMeminitializerid(ctx *cpp.MeminitializeridContext) {}

// EnterOperatorFunctionId is called when production operatorFunctionId is entered.
func (s *CppAstParser) EnterOperatorFunctionId(ctx *cpp.OperatorFunctionIdContext) {}

// ExitOperatorFunctionId is called when production operatorFunctionId is exited.
func (s *CppAstParser) ExitOperatorFunctionId(ctx *cpp.OperatorFunctionIdContext) {}

// EnterLiteralOperatorId is called when production literalOperatorId is entered.
func (s *CppAstParser) EnterLiteralOperatorId(ctx *cpp.LiteralOperatorIdContext) {}

// ExitLiteralOperatorId is called when production literalOperatorId is exited.
func (s *CppAstParser) ExitLiteralOperatorId(ctx *cpp.LiteralOperatorIdContext) {}

// EnterTemplateDeclaration is called when production templateDeclaration is entered.
func (s *CppAstParser) EnterTemplateDeclaration(ctx *cpp.TemplateDeclarationContext) {}

// ExitTemplateDeclaration is called when production templateDeclaration is exited.
func (s *CppAstParser) ExitTemplateDeclaration(ctx *cpp.TemplateDeclarationContext) {}

// EnterTemplateparameterList is called when production templateparameterList is entered.
func (s *CppAstParser) EnterTemplateparameterList(ctx *cpp.TemplateparameterListContext) {}

// ExitTemplateparameterList is called when production templateparameterList is exited.
func (s *CppAstParser) ExitTemplateparameterList(ctx *cpp.TemplateparameterListContext) {}

// EnterTemplateParameter is called when production templateParameter is entered.
func (s *CppAstParser) EnterTemplateParameter(ctx *cpp.TemplateParameterContext) {}

// ExitTemplateParameter is called when production templateParameter is exited.
func (s *CppAstParser) ExitTemplateParameter(ctx *cpp.TemplateParameterContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *CppAstParser) EnterTypeParameter(ctx *cpp.TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *CppAstParser) ExitTypeParameter(ctx *cpp.TypeParameterContext) {}

// EnterSimpleTemplateId is called when production simpleTemplateId is entered.
func (s *CppAstParser) EnterSimpleTemplateId(ctx *cpp.SimpleTemplateIdContext) {}

// ExitSimpleTemplateId is called when production simpleTemplateId is exited.
func (s *CppAstParser) ExitSimpleTemplateId(ctx *cpp.SimpleTemplateIdContext) {}

// EnterTemplateId is called when production templateId is entered.
func (s *CppAstParser) EnterTemplateId(ctx *cpp.TemplateIdContext) {}

// ExitTemplateId is called when production templateId is exited.
func (s *CppAstParser) ExitTemplateId(ctx *cpp.TemplateIdContext) {}

// EnterTemplateName is called when production templateName is entered.
func (s *CppAstParser) EnterTemplateName(ctx *cpp.TemplateNameContext) {}

// ExitTemplateName is called when production templateName is exited.
func (s *CppAstParser) ExitTemplateName(ctx *cpp.TemplateNameContext) {}

// EnterTemplateArgumentList is called when production templateArgumentList is entered.
func (s *CppAstParser) EnterTemplateArgumentList(ctx *cpp.TemplateArgumentListContext) {}

// ExitTemplateArgumentList is called when production templateArgumentList is exited.
func (s *CppAstParser) ExitTemplateArgumentList(ctx *cpp.TemplateArgumentListContext) {}

// EnterTemplateArgument is called when production templateArgument is entered.
func (s *CppAstParser) EnterTemplateArgument(ctx *cpp.TemplateArgumentContext) {}

// ExitTemplateArgument is called when production templateArgument is exited.
func (s *CppAstParser) ExitTemplateArgument(ctx *cpp.TemplateArgumentContext) {}

// EnterTypeNameSpecifier is called when production typeNameSpecifier is entered.
func (s *CppAstParser) EnterTypeNameSpecifier(ctx *cpp.TypeNameSpecifierContext) {}

// ExitTypeNameSpecifier is called when production typeNameSpecifier is exited.
func (s *CppAstParser) ExitTypeNameSpecifier(ctx *cpp.TypeNameSpecifierContext) {}

// EnterExplicitInstantiation is called when production explicitInstantiation is entered.
func (s *CppAstParser) EnterExplicitInstantiation(ctx *cpp.ExplicitInstantiationContext) {}

// ExitExplicitInstantiation is called when production explicitInstantiation is exited.
func (s *CppAstParser) ExitExplicitInstantiation(ctx *cpp.ExplicitInstantiationContext) {}

// EnterExplicitSpecialization is called when production explicitSpecialization is entered.
func (s *CppAstParser) EnterExplicitSpecialization(ctx *cpp.ExplicitSpecializationContext) {}

// ExitExplicitSpecialization is called when production explicitSpecialization is exited.
func (s *CppAstParser) ExitExplicitSpecialization(ctx *cpp.ExplicitSpecializationContext) {}

// EnterTryBlock is called when production tryBlock is entered.
func (s *CppAstParser) EnterTryBlock(ctx *cpp.TryBlockContext) {}

// ExitTryBlock is called when production tryBlock is exited.
func (s *CppAstParser) ExitTryBlock(ctx *cpp.TryBlockContext) {}

// EnterFunctionTryBlock is called when production functionTryBlock is entered.
func (s *CppAstParser) EnterFunctionTryBlock(ctx *cpp.FunctionTryBlockContext) {}

// ExitFunctionTryBlock is called when production functionTryBlock is exited.
func (s *CppAstParser) ExitFunctionTryBlock(ctx *cpp.FunctionTryBlockContext) {}

// EnterHandlerSeq is called when production handlerSeq is entered.
func (s *CppAstParser) EnterHandlerSeq(ctx *cpp.HandlerSeqContext) {}

// ExitHandlerSeq is called when production handlerSeq is exited.
func (s *CppAstParser) ExitHandlerSeq(ctx *cpp.HandlerSeqContext) {}

// EnterHandler is called when production handler is entered.
func (s *CppAstParser) EnterHandler(ctx *cpp.HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *CppAstParser) ExitHandler(ctx *cpp.HandlerContext) {}

// EnterExceptionDeclaration is called when production exceptionDeclaration is entered.
func (s *CppAstParser) EnterExceptionDeclaration(ctx *cpp.ExceptionDeclarationContext) {}

// ExitExceptionDeclaration is called when production exceptionDeclaration is exited.
func (s *CppAstParser) ExitExceptionDeclaration(ctx *cpp.ExceptionDeclarationContext) {}

// EnterThrowExpression is called when production throwExpression is entered.
func (s *CppAstParser) EnterThrowExpression(ctx *cpp.ThrowExpressionContext) {}

// ExitThrowExpression is called when production throwExpression is exited.
func (s *CppAstParser) ExitThrowExpression(ctx *cpp.ThrowExpressionContext) {}

// EnterExceptionSpecification is called when production exceptionSpecification is entered.
func (s *CppAstParser) EnterExceptionSpecification(ctx *cpp.ExceptionSpecificationContext) {}

// ExitExceptionSpecification is called when production exceptionSpecification is exited.
func (s *CppAstParser) ExitExceptionSpecification(ctx *cpp.ExceptionSpecificationContext) {}

// EnterDynamicExceptionSpecification is called when production dynamicExceptionSpecification is entered.
func (s *CppAstParser) EnterDynamicExceptionSpecification(ctx *cpp.DynamicExceptionSpecificationContext) {
}

// ExitDynamicExceptionSpecification is called when production dynamicExceptionSpecification is exited.
func (s *CppAstParser) ExitDynamicExceptionSpecification(ctx *cpp.DynamicExceptionSpecificationContext) {
}

// EnterTypeIdList is called when production typeIdList is entered.
func (s *CppAstParser) EnterTypeIdList(ctx *cpp.TypeIdListContext) {}

// ExitTypeIdList is called when production typeIdList is exited.
func (s *CppAstParser) ExitTypeIdList(ctx *cpp.TypeIdListContext) {}

// EnterNoeExceptSpecification is called when production noeExceptSpecification is entered.
func (s *CppAstParser) EnterNoeExceptSpecification(ctx *cpp.NoeExceptSpecificationContext) {}

// ExitNoeExceptSpecification is called when production noeExceptSpecification is exited.
func (s *CppAstParser) ExitNoeExceptSpecification(ctx *cpp.NoeExceptSpecificationContext) {}

// EnterTheOperator is called when production theOperator is entered.
func (s *CppAstParser) EnterTheOperator(ctx *cpp.TheOperatorContext) {}

// ExitTheOperator is called when production theOperator is exited.
func (s *CppAstParser) ExitTheOperator(ctx *cpp.TheOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *CppAstParser) EnterLiteral(ctx *cpp.LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *CppAstParser) ExitLiteral(ctx *cpp.LiteralContext) {}
