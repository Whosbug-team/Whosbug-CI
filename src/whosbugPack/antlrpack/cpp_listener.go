package antlrpack

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wxnacy/wgo/arrays"
	"strings"
	cpp "whosbugPack/antlrpack/cpp_lib"
)

// EnterClassSpecifier is called when production classSpecifier is entered.
func (s *CppTreeShapeListener) EnterClassSpecifier(ctx *cpp.ClassSpecifierContext) {
	var classInfo  = classInfoType{
		StartLine: ctx.GetStart().GetLine(),
		EndLine: ctx.GetStop().GetLine(),
		MasterObject: masterObjectInfoType{},
	}
	name := ctx.GetChild(0).(*cpp.ClassHeadContext)
	if name.ClassHeadName() != nil{
		classInfo.ClassName = name.ClassHeadName().GetText()
	}
	s.Infos.AstInfoList.Classes = append(s.Infos.AstInfoList.Classes, classInfo)
}
// EnterNoPointerDeclarator is called when production noPointerDeclarator is entered.
func (s *CppTreeShapeListener) EnterNoPointerDeclarator(ctx *cpp.NoPointerDeclaratorContext) {}

// ExitNoPointerDeclarator is called when production noPointerDeclarator is exited.
func (s *CppTreeShapeListener) ExitNoPointerDeclarator(ctx *cpp.NoPointerDeclaratorContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *CppTreeShapeListener) EnterFunctionDefinition(ctx *cpp.FunctionDefinitionContext) {
	s.Type = "function"
}
// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *CppTreeShapeListener) ExitFunctionDefinition(ctx *cpp.FunctionDefinitionContext) {
	var methodInfo = MethodInfoType{
		StartLine:    ctx.GetStart().GetLine(),
		EndLine:      ctx.GetStop().GetLine(),
	}

	name := strings.Split(ctx.Declarator().GetText(), "(")
	FuncAndMaster := strings.Split(name[0],"::")
	len := len(FuncAndMaster)
	methodInfo.MethodName = FuncAndMaster[len - 1]
	if len == 1{	//该方法不是成员方法
		methodInfo.MasterObject = masterObjectInfoType{}
	}else{
		methodInfo.MasterObject = masterObjectInfoType{
			ObjectName: strings.Join(FuncAndMaster[:len],"."),
			StartLine:  0,
		}
	}

	methodInfo.CallMethods = s.findMethodCall()

	s.Infos.AstInfoList.Methods = append(s.Infos.AstInfoList.Methods, methodInfo)
	s.Infos.CallMethods = []CallMethodType{}
	s.Type = ""
}

func (s *CppTreeShapeListener) findMethodCall() []string{
	var structMethods []string
	for index := range s.Infos.CallMethods {
		structMethods = append(structMethods,s.Infos.CallMethods[index].Id)
	}
	return structMethods
}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *CppTreeShapeListener) EnterExpressionStatement(ctx *cpp.ExpressionStatementContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *CppTreeShapeListener) EnterPostfixExpression(ctx *cpp.PostfixExpressionContext) {
	children := ctx.GetChildren()
	_, parentOk :=ctx.GetParent().(*cpp.PostfixExpressionContext)
	if len(children) == 3 && parentOk && children[1].(antlr.ParseTree).GetText() == "."{
		if _, ok := ctx.GetChild(2).(*cpp.IdExpressionContext); ok {
			methodCalled := children[2].(antlr.ParseTree).GetText()
			classBelong := children[0].(antlr.ParseTree).GetText()
			var flag bool = false
			for _, member := range s.Declaration {
				if member.Value == classBelong{
					classBelong = member.Name
					flag = true
					break
				}
			}
			if flag == true{
				var callMethod = CallMethodType{
					StartLine: ctx.GetStart().GetLine(),
					Id:        classBelong+"."+methodCalled,
				}
				s.Infos.CallMethods = append(s.Infos.CallMethods, callMethod)
			}
		}
	}
}

func (s *CppTreeShapeListener) EnterSimpleDeclaration(ctx *cpp.SimpleDeclarationContext) {
	if s.Type == "function"{
		if ctx.DeclSpecifierSeq() != nil{
			def := arrays.ContainsString([]string{"int","double","bool","byte","float64","float32","auto"},ctx.DeclSpecifierSeq().GetText())
			if def < 0{
				var member MemberType
				member.Name = ctx.DeclSpecifierSeq().GetText()
				list := strings.Split(ctx.InitDeclaratorList().GetText(),",")
				for i := 0; i < len(list); i++ {
					member.Value = list[i]
					s.Declaration = append(s.Declaration, member)
				}
			}
		}
		//fmt.Printf("---Declaration:%+v\n",s.Declaration)
	}

}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *CppTreeShapeListener) ExitExpressionStatement(ctx *cpp.ExpressionStatementContext) {}



// VisitTerminal is called when a terminal node is visited.
func (s *CppTreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *CppTreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *CppTreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *CppTreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *CppTreeShapeListener) EnterTranslationUnit(ctx *cpp.TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *CppTreeShapeListener) ExitTranslationUnit(ctx *cpp.TranslationUnitContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *CppTreeShapeListener) EnterPrimaryExpression(ctx *cpp.PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *CppTreeShapeListener) ExitPrimaryExpression(ctx *cpp.PrimaryExpressionContext) {}

// EnterIdExpression is called when production idExpression is entered.
func (s *CppTreeShapeListener) EnterIdExpression(ctx *cpp.IdExpressionContext) {}

// ExitIdExpression is called when production idExpression is exited.
func (s *CppTreeShapeListener) ExitIdExpression(ctx *cpp.IdExpressionContext) {}

// EnterUnqualifiedId is called when production unqualifiedId is entered.
func (s *CppTreeShapeListener) EnterUnqualifiedId(ctx *cpp.UnqualifiedIdContext) {}

// ExitUnqualifiedId is called when production unqualifiedId is exited.
func (s *CppTreeShapeListener) ExitUnqualifiedId(ctx *cpp.UnqualifiedIdContext) {}

// EnterQualifiedId is called when production qualifiedId is entered.
func (s *CppTreeShapeListener) EnterQualifiedId(ctx *cpp.QualifiedIdContext) {}

// ExitQualifiedId is called when production qualifiedId is exited.
func (s *CppTreeShapeListener) ExitQualifiedId(ctx *cpp.QualifiedIdContext) {}

// EnterNestedNameSpecifier is called when production nestedNameSpecifier is entered.
func (s *CppTreeShapeListener) EnterNestedNameSpecifier(ctx *cpp.NestedNameSpecifierContext) {}

// ExitNestedNameSpecifier is called when production nestedNameSpecifier is exited.
func (s *CppTreeShapeListener) ExitNestedNameSpecifier(ctx *cpp.NestedNameSpecifierContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *CppTreeShapeListener) EnterLambdaExpression(ctx *cpp.LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *CppTreeShapeListener) ExitLambdaExpression(ctx *cpp.LambdaExpressionContext) {}

// EnterLambdaIntroducer is called when production lambdaIntroducer is entered.
func (s *CppTreeShapeListener) EnterLambdaIntroducer(ctx *cpp.LambdaIntroducerContext) {}

// ExitLambdaIntroducer is called when production lambdaIntroducer is exited.
func (s *CppTreeShapeListener) ExitLambdaIntroducer(ctx *cpp.LambdaIntroducerContext) {}

// EnterLambdaCapture is called when production lambdaCapture is entered.
func (s *CppTreeShapeListener) EnterLambdaCapture(ctx *cpp.LambdaCaptureContext) {}

// ExitLambdaCapture is called when production lambdaCapture is exited.
func (s *CppTreeShapeListener) ExitLambdaCapture(ctx *cpp.LambdaCaptureContext) {}

// EnterCaptureDefault is called when production captureDefault is entered.
func (s *CppTreeShapeListener) EnterCaptureDefault(ctx *cpp.CaptureDefaultContext) {}

// ExitCaptureDefault is called when production captureDefault is exited.
func (s *CppTreeShapeListener) ExitCaptureDefault(ctx *cpp.CaptureDefaultContext) {}

// EnterCaptureList is called when production captureList is entered.
func (s *CppTreeShapeListener) EnterCaptureList(ctx *cpp.CaptureListContext) {}

// ExitCaptureList is called when production captureList is exited.
func (s *CppTreeShapeListener) ExitCaptureList(ctx *cpp.CaptureListContext) {}

// EnterCapture is called when production capture is entered.
func (s *CppTreeShapeListener) EnterCapture(ctx *cpp.CaptureContext) {}

// ExitCapture is called when production capture is exited.
func (s *CppTreeShapeListener) ExitCapture(ctx *cpp.CaptureContext) {}

// EnterSimpleCapture is called when production simpleCapture is entered.
func (s *CppTreeShapeListener) EnterSimpleCapture(ctx *cpp.SimpleCaptureContext) {}

// ExitSimpleCapture is called when production simpleCapture is exited.
func (s *CppTreeShapeListener) ExitSimpleCapture(ctx *cpp.SimpleCaptureContext) {}

// EnterInitcapture is called when production initcapture is entered.
func (s *CppTreeShapeListener) EnterInitcapture(ctx *cpp.InitcaptureContext) {}

// ExitInitcapture is called when production initcapture is exited.
func (s *CppTreeShapeListener) ExitInitcapture(ctx *cpp.InitcaptureContext) {}

// EnterLambdaDeclarator is called when production lambdaDeclarator is entered.
func (s *CppTreeShapeListener) EnterLambdaDeclarator(ctx *cpp.LambdaDeclaratorContext) {}

// ExitLambdaDeclarator is called when production lambdaDeclarator is exited.
func (s *CppTreeShapeListener) ExitLambdaDeclarator(ctx *cpp.LambdaDeclaratorContext) {}



// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *CppTreeShapeListener) ExitPostfixExpression(ctx *cpp.PostfixExpressionContext) {
}

// EnterTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is entered.
func (s *CppTreeShapeListener) EnterTypeIdOfTheTypeId(ctx *cpp.TypeIdOfTheTypeIdContext) {}

// ExitTypeIdOfTheTypeId is called when production typeIdOfTheTypeId is exited.
func (s *CppTreeShapeListener) ExitTypeIdOfTheTypeId(ctx *cpp.TypeIdOfTheTypeIdContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *CppTreeShapeListener) EnterExpressionList(ctx *cpp.ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *CppTreeShapeListener) ExitExpressionList(ctx *cpp.ExpressionListContext) {}

// EnterPseudoDestructorName is called when production pseudoDestructorName is entered.
func (s *CppTreeShapeListener) EnterPseudoDestructorName(ctx *cpp.PseudoDestructorNameContext) {}

// ExitPseudoDestructorName is called when production pseudoDestructorName is exited.
func (s *CppTreeShapeListener) ExitPseudoDestructorName(ctx *cpp.PseudoDestructorNameContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *CppTreeShapeListener) EnterUnaryExpression(ctx *cpp.UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *CppTreeShapeListener) ExitUnaryExpression(ctx *cpp.UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *CppTreeShapeListener) EnterUnaryOperator(ctx *cpp.UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *CppTreeShapeListener) ExitUnaryOperator(ctx *cpp.UnaryOperatorContext) {}

// EnterNewExpression is called when production newExpression is entered.
func (s *CppTreeShapeListener) EnterNewExpression(ctx *cpp.NewExpressionContext) {}

// ExitNewExpression is called when production newExpression is exited.
func (s *CppTreeShapeListener) ExitNewExpression(ctx *cpp.NewExpressionContext) {}

// EnterNewPlacement is called when production newPlacement is entered.
func (s *CppTreeShapeListener) EnterNewPlacement(ctx *cpp.NewPlacementContext) {}

// ExitNewPlacement is called when production newPlacement is exited.
func (s *CppTreeShapeListener) ExitNewPlacement(ctx *cpp.NewPlacementContext) {}

// EnterNewTypeId is called when production newTypeId is entered.
func (s *CppTreeShapeListener) EnterNewTypeId(ctx *cpp.NewTypeIdContext) {}

// ExitNewTypeId is called when production newTypeId is exited.
func (s *CppTreeShapeListener) ExitNewTypeId(ctx *cpp.NewTypeIdContext) {}

// EnterNewDeclarator is called when production newDeclarator is entered.
func (s *CppTreeShapeListener) EnterNewDeclarator(ctx *cpp.NewDeclaratorContext) {}

// ExitNewDeclarator is called when production newDeclarator is exited.
func (s *CppTreeShapeListener) ExitNewDeclarator(ctx *cpp.NewDeclaratorContext) {}

// EnterNoPointerNewDeclarator is called when production noPointerNewDeclarator is entered.
func (s *CppTreeShapeListener) EnterNoPointerNewDeclarator(ctx *cpp.NoPointerNewDeclaratorContext) {}

// ExitNoPointerNewDeclarator is called when production noPointerNewDeclarator is exited.
func (s *CppTreeShapeListener) ExitNoPointerNewDeclarator(ctx *cpp.NoPointerNewDeclaratorContext) {}

// EnterNewInitializer is called when production newInitializer is entered.
func (s *CppTreeShapeListener) EnterNewInitializer(ctx *cpp.NewInitializerContext) {}

// ExitNewInitializer is called when production newInitializer is exited.
func (s *CppTreeShapeListener) ExitNewInitializer(ctx *cpp.NewInitializerContext) {}

// EnterDeleteExpression is called when production deleteExpression is entered.
func (s *CppTreeShapeListener) EnterDeleteExpression(ctx *cpp.DeleteExpressionContext) {}

// ExitDeleteExpression is called when production deleteExpression is exited.
func (s *CppTreeShapeListener) ExitDeleteExpression(ctx *cpp.DeleteExpressionContext) {}

// EnterNoExceptExpression is called when production noExceptExpression is entered.
func (s *CppTreeShapeListener) EnterNoExceptExpression(ctx *cpp.NoExceptExpressionContext) {}

// ExitNoExceptExpression is called when production noExceptExpression is exited.
func (s *CppTreeShapeListener) ExitNoExceptExpression(ctx *cpp.NoExceptExpressionContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *CppTreeShapeListener) EnterCastExpression(ctx *cpp.CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *CppTreeShapeListener) ExitCastExpression(ctx *cpp.CastExpressionContext) {}

// EnterPointerMemberExpression is called when production pointerMemberExpression is entered.
func (s *CppTreeShapeListener) EnterPointerMemberExpression(ctx *cpp.PointerMemberExpressionContext) {}

// ExitPointerMemberExpression is called when production pointerMemberExpression is exited.
func (s *CppTreeShapeListener) ExitPointerMemberExpression(ctx *cpp.PointerMemberExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *CppTreeShapeListener) EnterMultiplicativeExpression(ctx *cpp.MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *CppTreeShapeListener) ExitMultiplicativeExpression(ctx *cpp.MultiplicativeExpressionContext) {
}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *CppTreeShapeListener) EnterAdditiveExpression(ctx *cpp.AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *CppTreeShapeListener) ExitAdditiveExpression(ctx *cpp.AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *CppTreeShapeListener) EnterShiftExpression(ctx *cpp.ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *CppTreeShapeListener) ExitShiftExpression(ctx *cpp.ShiftExpressionContext) {}

// EnterShiftOperator is called when production shiftOperator is entered.
func (s *CppTreeShapeListener) EnterShiftOperator(ctx *cpp.ShiftOperatorContext) {}

// ExitShiftOperator is called when production shiftOperator is exited.
func (s *CppTreeShapeListener) ExitShiftOperator(ctx *cpp.ShiftOperatorContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *CppTreeShapeListener) EnterRelationalExpression(ctx *cpp.RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *CppTreeShapeListener) ExitRelationalExpression(ctx *cpp.RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *CppTreeShapeListener) EnterEqualityExpression(ctx *cpp.EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *CppTreeShapeListener) ExitEqualityExpression(ctx *cpp.EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *CppTreeShapeListener) EnterAndExpression(ctx *cpp.AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *CppTreeShapeListener) ExitAndExpression(ctx *cpp.AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *CppTreeShapeListener) EnterExclusiveOrExpression(ctx *cpp.ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *CppTreeShapeListener) ExitExclusiveOrExpression(ctx *cpp.ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *CppTreeShapeListener) EnterInclusiveOrExpression(ctx *cpp.InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *CppTreeShapeListener) ExitInclusiveOrExpression(ctx *cpp.InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *CppTreeShapeListener) EnterLogicalAndExpression(ctx *cpp.LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *CppTreeShapeListener) ExitLogicalAndExpression(ctx *cpp.LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *CppTreeShapeListener) EnterLogicalOrExpression(ctx *cpp.LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *CppTreeShapeListener) ExitLogicalOrExpression(ctx *cpp.LogicalOrExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *CppTreeShapeListener) EnterConditionalExpression(ctx *cpp.ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *CppTreeShapeListener) ExitConditionalExpression(ctx *cpp.ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *CppTreeShapeListener) EnterAssignmentExpression(ctx *cpp.AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *CppTreeShapeListener) ExitAssignmentExpression(ctx *cpp.AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *CppTreeShapeListener) EnterAssignmentOperator(ctx *cpp.AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *CppTreeShapeListener) ExitAssignmentOperator(ctx *cpp.AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *CppTreeShapeListener) EnterExpression(ctx *cpp.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *CppTreeShapeListener) ExitExpression(ctx *cpp.ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *CppTreeShapeListener) EnterConstantExpression(ctx *cpp.ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *CppTreeShapeListener) ExitConstantExpression(ctx *cpp.ConstantExpressionContext) {}

// EnterStatement is called when production statement is entered.
func (s *CppTreeShapeListener) EnterStatement(ctx *cpp.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *CppTreeShapeListener) ExitStatement(ctx *cpp.StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *CppTreeShapeListener) EnterLabeledStatement(ctx *cpp.LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *CppTreeShapeListener) ExitLabeledStatement(ctx *cpp.LabeledStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *CppTreeShapeListener) EnterCompoundStatement(ctx *cpp.CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *CppTreeShapeListener) ExitCompoundStatement(ctx *cpp.CompoundStatementContext) {}

// EnterStatementSeq is called when production statementSeq is entered.
func (s *CppTreeShapeListener) EnterStatementSeq(ctx *cpp.StatementSeqContext) {}

// ExitStatementSeq is called when production statementSeq is exited.
func (s *CppTreeShapeListener) ExitStatementSeq(ctx *cpp.StatementSeqContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *CppTreeShapeListener) EnterSelectionStatement(ctx *cpp.SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *CppTreeShapeListener) ExitSelectionStatement(ctx *cpp.SelectionStatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *CppTreeShapeListener) EnterCondition(ctx *cpp.ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *CppTreeShapeListener) ExitCondition(ctx *cpp.ConditionContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *CppTreeShapeListener) EnterIterationStatement(ctx *cpp.IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *CppTreeShapeListener) ExitIterationStatement(ctx *cpp.IterationStatementContext) {}

// EnterForInitStatement is called when production forInitStatement is entered.
func (s *CppTreeShapeListener) EnterForInitStatement(ctx *cpp.ForInitStatementContext) {}

// ExitForInitStatement is called when production forInitStatement is exited.
func (s *CppTreeShapeListener) ExitForInitStatement(ctx *cpp.ForInitStatementContext) {}

// EnterForRangeDeclaration is called when production forRangeDeclaration is entered.
func (s *CppTreeShapeListener) EnterForRangeDeclaration(ctx *cpp.ForRangeDeclarationContext) {}

// ExitForRangeDeclaration is called when production forRangeDeclaration is exited.
func (s *CppTreeShapeListener) ExitForRangeDeclaration(ctx *cpp.ForRangeDeclarationContext) {}

// EnterForRangeInitializer is called when production forRangeInitializer is entered.
func (s *CppTreeShapeListener) EnterForRangeInitializer(ctx *cpp.ForRangeInitializerContext) {}

// ExitForRangeInitializer is called when production forRangeInitializer is exited.
func (s *CppTreeShapeListener) ExitForRangeInitializer(ctx *cpp.ForRangeInitializerContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *CppTreeShapeListener) EnterJumpStatement(ctx *cpp.JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *CppTreeShapeListener) ExitJumpStatement(ctx *cpp.JumpStatementContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *CppTreeShapeListener) EnterDeclarationStatement(ctx *cpp.DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *CppTreeShapeListener) ExitDeclarationStatement(ctx *cpp.DeclarationStatementContext) {}

// EnterDeclarationseq is called when production declarationseq is entered.
func (s *CppTreeShapeListener) EnterDeclarationseq(ctx *cpp.DeclarationseqContext) {}

// ExitDeclarationseq is called when production declarationseq is exited.
func (s *CppTreeShapeListener) ExitDeclarationseq(ctx *cpp.DeclarationseqContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *CppTreeShapeListener) EnterDeclaration(ctx *cpp.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *CppTreeShapeListener) ExitDeclaration(ctx *cpp.DeclarationContext) {}

// EnterBlockDeclaration is called when production blockDeclaration is entered.
func (s *CppTreeShapeListener) EnterBlockDeclaration(ctx *cpp.BlockDeclarationContext) {}

// ExitBlockDeclaration is called when production blockDeclaration is exited.
func (s *CppTreeShapeListener) ExitBlockDeclaration(ctx *cpp.BlockDeclarationContext) {}

// EnterAliasDeclaration is called when production aliasDeclaration is entered.
func (s *CppTreeShapeListener) EnterAliasDeclaration(ctx *cpp.AliasDeclarationContext) {}

// ExitAliasDeclaration is called when production aliasDeclaration is exited.
func (s *CppTreeShapeListener) ExitAliasDeclaration(ctx *cpp.AliasDeclarationContext) {}

//// EnterSimpleDeclaration is called when production simpleDeclaration is entered.
//func (s *CppTreeShapeListener) EnterSimpleDeclaration(ctx *cpp.SimpleDeclarationContext) {}

// ExitSimpleDeclaration is called when production simpleDeclaration is exited.
func (s *CppTreeShapeListener) ExitSimpleDeclaration(ctx *cpp.SimpleDeclarationContext) {}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *CppTreeShapeListener) EnterStaticAssertDeclaration(ctx *cpp.StaticAssertDeclarationContext) {}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *CppTreeShapeListener) ExitStaticAssertDeclaration(ctx *cpp.StaticAssertDeclarationContext) {}

// EnterEmptyDeclaration is called when production emptyDeclaration is entered.
func (s *CppTreeShapeListener) EnterEmptyDeclaration(ctx *cpp.EmptyDeclarationContext) {}

// ExitEmptyDeclaration is called when production emptyDeclaration is exited.
func (s *CppTreeShapeListener) ExitEmptyDeclaration(ctx *cpp.EmptyDeclarationContext) {}

// EnterAttributeDeclaration is called when production attributeDeclaration is entered.
func (s *CppTreeShapeListener) EnterAttributeDeclaration(ctx *cpp.AttributeDeclarationContext) {}

// ExitAttributeDeclaration is called when production attributeDeclaration is exited.
func (s *CppTreeShapeListener) ExitAttributeDeclaration(ctx *cpp.AttributeDeclarationContext) {}

// EnterDeclSpecifier is called when production declSpecifier is entered.
func (s *CppTreeShapeListener) EnterDeclSpecifier(ctx *cpp.DeclSpecifierContext) {}

// ExitDeclSpecifier is called when production declSpecifier is exited.
func (s *CppTreeShapeListener) ExitDeclSpecifier(ctx *cpp.DeclSpecifierContext) {}

// EnterDeclSpecifierSeq is called when production declSpecifierSeq is entered.
func (s *CppTreeShapeListener) EnterDeclSpecifierSeq(ctx *cpp.DeclSpecifierSeqContext) {}

// ExitDeclSpecifierSeq is called when production declSpecifierSeq is exited.
func (s *CppTreeShapeListener) ExitDeclSpecifierSeq(ctx *cpp.DeclSpecifierSeqContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *CppTreeShapeListener) EnterStorageClassSpecifier(ctx *cpp.StorageClassSpecifierContext) {}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *CppTreeShapeListener) ExitStorageClassSpecifier(ctx *cpp.StorageClassSpecifierContext) {}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *CppTreeShapeListener) EnterFunctionSpecifier(ctx *cpp.FunctionSpecifierContext) {}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *CppTreeShapeListener) ExitFunctionSpecifier(ctx *cpp.FunctionSpecifierContext) {

}

// EnterTypedefName is called when production typedefName is entered.
func (s *CppTreeShapeListener) EnterTypedefName(ctx *cpp.TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *CppTreeShapeListener) ExitTypedefName(ctx *cpp.TypedefNameContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *CppTreeShapeListener) EnterTypeSpecifier(ctx *cpp.TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *CppTreeShapeListener) ExitTypeSpecifier(ctx *cpp.TypeSpecifierContext) {}

// EnterTrailingTypeSpecifier is called when production trailingTypeSpecifier is entered.
func (s *CppTreeShapeListener) EnterTrailingTypeSpecifier(ctx *cpp.TrailingTypeSpecifierContext) {}

// ExitTrailingTypeSpecifier is called when production trailingTypeSpecifier is exited.
func (s *CppTreeShapeListener) ExitTrailingTypeSpecifier(ctx *cpp.TrailingTypeSpecifierContext) {}

// EnterTypeSpecifierSeq is called when production typeSpecifierSeq is entered.
func (s *CppTreeShapeListener) EnterTypeSpecifierSeq(ctx *cpp.TypeSpecifierSeqContext) {}

// ExitTypeSpecifierSeq is called when production typeSpecifierSeq is exited.
func (s *CppTreeShapeListener) ExitTypeSpecifierSeq(ctx *cpp.TypeSpecifierSeqContext) {}

// EnterTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is entered.
func (s *CppTreeShapeListener) EnterTrailingTypeSpecifierSeq(ctx *cpp.TrailingTypeSpecifierSeqContext) {
}

// ExitTrailingTypeSpecifierSeq is called when production trailingTypeSpecifierSeq is exited.
func (s *CppTreeShapeListener) ExitTrailingTypeSpecifierSeq(ctx *cpp.TrailingTypeSpecifierSeqContext) {
}

// EnterSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is entered.
func (s *CppTreeShapeListener) EnterSimpleTypeLengthModifier(ctx *cpp.SimpleTypeLengthModifierContext) {
}

// ExitSimpleTypeLengthModifier is called when production simpleTypeLengthModifier is exited.
func (s *CppTreeShapeListener) ExitSimpleTypeLengthModifier(ctx *cpp.SimpleTypeLengthModifierContext) {
}

// EnterSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is entered.
func (s *CppTreeShapeListener) EnterSimpleTypeSignednessModifier(ctx *cpp.SimpleTypeSignednessModifierContext) {
}

// ExitSimpleTypeSignednessModifier is called when production simpleTypeSignednessModifier is exited.
func (s *CppTreeShapeListener) ExitSimpleTypeSignednessModifier(ctx *cpp.SimpleTypeSignednessModifierContext) {
}

// EnterSimpleTypeSpecifier is called when production simpleTypeSpecifier is entered.
func (s *CppTreeShapeListener) EnterSimpleTypeSpecifier(ctx *cpp.SimpleTypeSpecifierContext) {}

// ExitSimpleTypeSpecifier is called when production simpleTypeSpecifier is exited.
func (s *CppTreeShapeListener) ExitSimpleTypeSpecifier(ctx *cpp.SimpleTypeSpecifierContext) {}

// EnterTheTypeName is called when production theTypeName is entered.
func (s *CppTreeShapeListener) EnterTheTypeName(ctx *cpp.TheTypeNameContext) {}

// ExitTheTypeName is called when production theTypeName is exited.
func (s *CppTreeShapeListener) ExitTheTypeName(ctx *cpp.TheTypeNameContext) {}

// EnterDecltypeSpecifier is called when production decltypeSpecifier is entered.
func (s *CppTreeShapeListener) EnterDecltypeSpecifier(ctx *cpp.DecltypeSpecifierContext) {}

// ExitDecltypeSpecifier is called when production decltypeSpecifier is exited.
func (s *CppTreeShapeListener) ExitDecltypeSpecifier(ctx *cpp.DecltypeSpecifierContext) {}

// EnterElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is entered.
func (s *CppTreeShapeListener) EnterElaboratedTypeSpecifier(ctx *cpp.ElaboratedTypeSpecifierContext) {}

// ExitElaboratedTypeSpecifier is called when production elaboratedTypeSpecifier is exited.
func (s *CppTreeShapeListener) ExitElaboratedTypeSpecifier(ctx *cpp.ElaboratedTypeSpecifierContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *CppTreeShapeListener) EnterEnumName(ctx *cpp.EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *CppTreeShapeListener) ExitEnumName(ctx *cpp.EnumNameContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *CppTreeShapeListener) EnterEnumSpecifier(ctx *cpp.EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *CppTreeShapeListener) ExitEnumSpecifier(ctx *cpp.EnumSpecifierContext) {}

// EnterEnumHead is called when production enumHead is entered.
func (s *CppTreeShapeListener) EnterEnumHead(ctx *cpp.EnumHeadContext) {}

// ExitEnumHead is called when production enumHead is exited.
func (s *CppTreeShapeListener) ExitEnumHead(ctx *cpp.EnumHeadContext) {}

// EnterOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is entered.
func (s *CppTreeShapeListener) EnterOpaqueEnumDeclaration(ctx *cpp.OpaqueEnumDeclarationContext) {}

// ExitOpaqueEnumDeclaration is called when production opaqueEnumDeclaration is exited.
func (s *CppTreeShapeListener) ExitOpaqueEnumDeclaration(ctx *cpp.OpaqueEnumDeclarationContext) {}

// EnterEnumkey is called when production enumkey is entered.
func (s *CppTreeShapeListener) EnterEnumkey(ctx *cpp.EnumkeyContext) {}

// ExitEnumkey is called when production enumkey is exited.
func (s *CppTreeShapeListener) ExitEnumkey(ctx *cpp.EnumkeyContext) {}

// EnterEnumbase is called when production enumbase is entered.
func (s *CppTreeShapeListener) EnterEnumbase(ctx *cpp.EnumbaseContext) {}

// ExitEnumbase is called when production enumbase is exited.
func (s *CppTreeShapeListener) ExitEnumbase(ctx *cpp.EnumbaseContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *CppTreeShapeListener) EnterEnumeratorList(ctx *cpp.EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *CppTreeShapeListener) ExitEnumeratorList(ctx *cpp.EnumeratorListContext) {}

// EnterEnumeratorDefinition is called when production enumeratorDefinition is entered.
func (s *CppTreeShapeListener) EnterEnumeratorDefinition(ctx *cpp.EnumeratorDefinitionContext) {}

// ExitEnumeratorDefinition is called when production enumeratorDefinition is exited.
func (s *CppTreeShapeListener) ExitEnumeratorDefinition(ctx *cpp.EnumeratorDefinitionContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *CppTreeShapeListener) EnterEnumerator(ctx *cpp.EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *CppTreeShapeListener) ExitEnumerator(ctx *cpp.EnumeratorContext) {}

// EnterNamespaceName is called when production namespaceName is entered.
func (s *CppTreeShapeListener) EnterNamespaceName(ctx *cpp.NamespaceNameContext) {}

// ExitNamespaceName is called when production namespaceName is exited.
func (s *CppTreeShapeListener) ExitNamespaceName(ctx *cpp.NamespaceNameContext) {}

// EnterOriginalNamespaceName is called when production originalNamespaceName is entered.
func (s *CppTreeShapeListener) EnterOriginalNamespaceName(ctx *cpp.OriginalNamespaceNameContext) {}

// ExitOriginalNamespaceName is called when production originalNamespaceName is exited.
func (s *CppTreeShapeListener) ExitOriginalNamespaceName(ctx *cpp.OriginalNamespaceNameContext) {}

// EnterNamespaceDefinition is called when production namespaceDefinition is entered.
func (s *CppTreeShapeListener) EnterNamespaceDefinition(ctx *cpp.NamespaceDefinitionContext) {}

// ExitNamespaceDefinition is called when production namespaceDefinition is exited.
func (s *CppTreeShapeListener) ExitNamespaceDefinition(ctx *cpp.NamespaceDefinitionContext) {}

// EnterNamespaceAlias is called when production namespaceAlias is entered.
func (s *CppTreeShapeListener) EnterNamespaceAlias(ctx *cpp.NamespaceAliasContext) {}

// ExitNamespaceAlias is called when production namespaceAlias is exited.
func (s *CppTreeShapeListener) ExitNamespaceAlias(ctx *cpp.NamespaceAliasContext) {}

// EnterNamespaceAliasDefinition is called when production namespaceAliasDefinition is entered.
func (s *CppTreeShapeListener) EnterNamespaceAliasDefinition(ctx *cpp.NamespaceAliasDefinitionContext) {
}

// ExitNamespaceAliasDefinition is called when production namespaceAliasDefinition is exited.
func (s *CppTreeShapeListener) ExitNamespaceAliasDefinition(ctx *cpp.NamespaceAliasDefinitionContext) {
}

// EnterQualifiednamespacespecifier is called when production qualifiednamespacespecifier is entered.
func (s *CppTreeShapeListener) EnterQualifiednamespacespecifier(ctx *cpp.QualifiednamespacespecifierContext) {
}

// ExitQualifiednamespacespecifier is called when production qualifiednamespacespecifier is exited.
func (s *CppTreeShapeListener) ExitQualifiednamespacespecifier(ctx *cpp.QualifiednamespacespecifierContext) {
}

// EnterUsingDeclaration is called when production usingDeclaration is entered.
func (s *CppTreeShapeListener) EnterUsingDeclaration(ctx *cpp.UsingDeclarationContext) {}

// ExitUsingDeclaration is called when production usingDeclaration is exited.
func (s *CppTreeShapeListener) ExitUsingDeclaration(ctx *cpp.UsingDeclarationContext) {}

// EnterUsingDirective is called when production usingDirective is entered.
func (s *CppTreeShapeListener) EnterUsingDirective(ctx *cpp.UsingDirectiveContext) {}

// ExitUsingDirective is called when production usingDirective is exited.
func (s *CppTreeShapeListener) ExitUsingDirective(ctx *cpp.UsingDirectiveContext) {}

// EnterAsmDefinition is called when production asmDefinition is entered.
func (s *CppTreeShapeListener) EnterAsmDefinition(ctx *cpp.AsmDefinitionContext) {}

// ExitAsmDefinition is called when production asmDefinition is exited.
func (s *CppTreeShapeListener) ExitAsmDefinition(ctx *cpp.AsmDefinitionContext) {}

// EnterLinkageSpecification is called when production linkageSpecification is entered.
func (s *CppTreeShapeListener) EnterLinkageSpecification(ctx *cpp.LinkageSpecificationContext) {}

// ExitLinkageSpecification is called when production linkageSpecification is exited.
func (s *CppTreeShapeListener) ExitLinkageSpecification(ctx *cpp.LinkageSpecificationContext) {}

// EnterAttributeSpecifierSeq is called when production attributeSpecifierSeq is entered.
func (s *CppTreeShapeListener) EnterAttributeSpecifierSeq(ctx *cpp.AttributeSpecifierSeqContext) {}

// ExitAttributeSpecifierSeq is called when production attributeSpecifierSeq is exited.
func (s *CppTreeShapeListener) ExitAttributeSpecifierSeq(ctx *cpp.AttributeSpecifierSeqContext) {}

// EnterAttributeSpecifier is called when production attributeSpecifier is entered.
func (s *CppTreeShapeListener) EnterAttributeSpecifier(ctx *cpp.AttributeSpecifierContext) {}

// ExitAttributeSpecifier is called when production attributeSpecifier is exited.
func (s *CppTreeShapeListener) ExitAttributeSpecifier(ctx *cpp.AttributeSpecifierContext) {}

// EnterAlignmentspecifier is called when production alignmentspecifier is entered.
func (s *CppTreeShapeListener) EnterAlignmentspecifier(ctx *cpp.AlignmentspecifierContext) {}

// ExitAlignmentspecifier is called when production alignmentspecifier is exited.
func (s *CppTreeShapeListener) ExitAlignmentspecifier(ctx *cpp.AlignmentspecifierContext) {}

// EnterAttributeList is called when production attributeList is entered.
func (s *CppTreeShapeListener) EnterAttributeList(ctx *cpp.AttributeListContext) {}

// ExitAttributeList is called when production attributeList is exited.
func (s *CppTreeShapeListener) ExitAttributeList(ctx *cpp.AttributeListContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *CppTreeShapeListener) EnterAttribute(ctx *cpp.AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *CppTreeShapeListener) ExitAttribute(ctx *cpp.AttributeContext) {}

// EnterAttributeNamespace is called when production attributeNamespace is entered.
func (s *CppTreeShapeListener) EnterAttributeNamespace(ctx *cpp.AttributeNamespaceContext) {}

// ExitAttributeNamespace is called when production attributeNamespace is exited.
func (s *CppTreeShapeListener) ExitAttributeNamespace(ctx *cpp.AttributeNamespaceContext) {}

// EnterAttributeArgumentClause is called when production attributeArgumentClause is entered.
func (s *CppTreeShapeListener) EnterAttributeArgumentClause(ctx *cpp.AttributeArgumentClauseContext) {}

// ExitAttributeArgumentClause is called when production attributeArgumentClause is exited.
func (s *CppTreeShapeListener) ExitAttributeArgumentClause(ctx *cpp.AttributeArgumentClauseContext) {}

// EnterBalancedTokenSeq is called when production balancedTokenSeq is entered.
func (s *CppTreeShapeListener) EnterBalancedTokenSeq(ctx *cpp.BalancedTokenSeqContext) {}

// ExitBalancedTokenSeq is called when production balancedTokenSeq is exited.
func (s *CppTreeShapeListener) ExitBalancedTokenSeq(ctx *cpp.BalancedTokenSeqContext) {}

// EnterBalancedtoken is called when production balancedtoken is entered.
func (s *CppTreeShapeListener) EnterBalancedtoken(ctx *cpp.BalancedtokenContext) {}

// ExitBalancedtoken is called when production balancedtoken is exited.
func (s *CppTreeShapeListener) ExitBalancedtoken(ctx *cpp.BalancedtokenContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *CppTreeShapeListener) EnterInitDeclaratorList(ctx *cpp.InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *CppTreeShapeListener) ExitInitDeclaratorList(ctx *cpp.InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *CppTreeShapeListener) EnterInitDeclarator(ctx *cpp.InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *CppTreeShapeListener) ExitInitDeclarator(ctx *cpp.InitDeclaratorContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *CppTreeShapeListener) EnterDeclarator(ctx *cpp.DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *CppTreeShapeListener) ExitDeclarator(ctx *cpp.DeclaratorContext) {}

// EnterPointerDeclarator is called when production pointerDeclarator is entered.
func (s *CppTreeShapeListener) EnterPointerDeclarator(ctx *cpp.PointerDeclaratorContext) {}

// ExitPointerDeclarator is called when production pointerDeclarator is exited.
func (s *CppTreeShapeListener) ExitPointerDeclarator(ctx *cpp.PointerDeclaratorContext) {}

// EnterParametersAndQualifiers is called when production parametersAndQualifiers is entered.
func (s *CppTreeShapeListener) EnterParametersAndQualifiers(ctx *cpp.ParametersAndQualifiersContext) {}

// ExitParametersAndQualifiers is called when production parametersAndQualifiers is exited.
func (s *CppTreeShapeListener) ExitParametersAndQualifiers(ctx *cpp.ParametersAndQualifiersContext) {}

// EnterTrailingReturnType is called when production trailingReturnType is entered.
func (s *CppTreeShapeListener) EnterTrailingReturnType(ctx *cpp.TrailingReturnTypeContext) {}

// ExitTrailingReturnType is called when production trailingReturnType is exited.
func (s *CppTreeShapeListener) ExitTrailingReturnType(ctx *cpp.TrailingReturnTypeContext) {}

// EnterPointerOperator is called when production pointerOperator is entered.
func (s *CppTreeShapeListener) EnterPointerOperator(ctx *cpp.PointerOperatorContext) {}

// ExitPointerOperator is called when production pointerOperator is exited.
func (s *CppTreeShapeListener) ExitPointerOperator(ctx *cpp.PointerOperatorContext) {}

// EnterCvqualifierseq is called when production cvqualifierseq is entered.
func (s *CppTreeShapeListener) EnterCvqualifierseq(ctx *cpp.CvqualifierseqContext) {}

// ExitCvqualifierseq is called when production cvqualifierseq is exited.
func (s *CppTreeShapeListener) ExitCvqualifierseq(ctx *cpp.CvqualifierseqContext) {}

// EnterCvQualifier is called when production cvQualifier is entered.
func (s *CppTreeShapeListener) EnterCvQualifier(ctx *cpp.CvQualifierContext) {}

// ExitCvQualifier is called when production cvQualifier is exited.
func (s *CppTreeShapeListener) ExitCvQualifier(ctx *cpp.CvQualifierContext) {}

// EnterRefqualifier is called when production refqualifier is entered.
func (s *CppTreeShapeListener) EnterRefqualifier(ctx *cpp.RefqualifierContext) {}

// ExitRefqualifier is called when production refqualifier is exited.
func (s *CppTreeShapeListener) ExitRefqualifier(ctx *cpp.RefqualifierContext) {}

// EnterDeclaratorid is called when production declaratorid is entered.
func (s *CppTreeShapeListener) EnterDeclaratorid(ctx *cpp.DeclaratoridContext) {}

// ExitDeclaratorid is called when production declaratorid is exited.
func (s *CppTreeShapeListener) ExitDeclaratorid(ctx *cpp.DeclaratoridContext) {}

// EnterTheTypeId is called when production theTypeId is entered.
func (s *CppTreeShapeListener) EnterTheTypeId(ctx *cpp.TheTypeIdContext) {}

// ExitTheTypeId is called when production theTypeId is exited.
func (s *CppTreeShapeListener) ExitTheTypeId(ctx *cpp.TheTypeIdContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *CppTreeShapeListener) EnterAbstractDeclarator(ctx *cpp.AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *CppTreeShapeListener) ExitAbstractDeclarator(ctx *cpp.AbstractDeclaratorContext) {}

// EnterPointerAbstractDeclarator is called when production pointerAbstractDeclarator is entered.
func (s *CppTreeShapeListener) EnterPointerAbstractDeclarator(ctx *cpp.PointerAbstractDeclaratorContext) {
}

// ExitPointerAbstractDeclarator is called when production pointerAbstractDeclarator is exited.
func (s *CppTreeShapeListener) ExitPointerAbstractDeclarator(ctx *cpp.PointerAbstractDeclaratorContext) {
}

// EnterNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is entered.
func (s *CppTreeShapeListener) EnterNoPointerAbstractDeclarator(ctx *cpp.NoPointerAbstractDeclaratorContext) {
}

// ExitNoPointerAbstractDeclarator is called when production noPointerAbstractDeclarator is exited.
func (s *CppTreeShapeListener) ExitNoPointerAbstractDeclarator(ctx *cpp.NoPointerAbstractDeclaratorContext) {
}

// EnterAbstractPackDeclarator is called when production abstractPackDeclarator is entered.
func (s *CppTreeShapeListener) EnterAbstractPackDeclarator(ctx *cpp.AbstractPackDeclaratorContext) {}

// ExitAbstractPackDeclarator is called when production abstractPackDeclarator is exited.
func (s *CppTreeShapeListener) ExitAbstractPackDeclarator(ctx *cpp.AbstractPackDeclaratorContext) {}

// EnterNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is entered.
func (s *CppTreeShapeListener) EnterNoPointerAbstractPackDeclarator(ctx *cpp.NoPointerAbstractPackDeclaratorContext) {
}

// ExitNoPointerAbstractPackDeclarator is called when production noPointerAbstractPackDeclarator is exited.
func (s *CppTreeShapeListener) ExitNoPointerAbstractPackDeclarator(ctx *cpp.NoPointerAbstractPackDeclaratorContext) {
}

// EnterParameterDeclarationClause is called when production parameterDeclarationClause is entered.
func (s *CppTreeShapeListener) EnterParameterDeclarationClause(ctx *cpp.ParameterDeclarationClauseContext) {
}

// ExitParameterDeclarationClause is called when production parameterDeclarationClause is exited.
func (s *CppTreeShapeListener) ExitParameterDeclarationClause(ctx *cpp.ParameterDeclarationClauseContext) {
}

// EnterParameterDeclarationList is called when production parameterDeclarationList is entered.
func (s *CppTreeShapeListener) EnterParameterDeclarationList(ctx *cpp.ParameterDeclarationListContext) {
}

// ExitParameterDeclarationList is called when production parameterDeclarationList is exited.
func (s *CppTreeShapeListener) ExitParameterDeclarationList(ctx *cpp.ParameterDeclarationListContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *CppTreeShapeListener) EnterParameterDeclaration(ctx *cpp.ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *CppTreeShapeListener) ExitParameterDeclaration(ctx *cpp.ParameterDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *CppTreeShapeListener) EnterFunctionBody(ctx *cpp.FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *CppTreeShapeListener) ExitFunctionBody(ctx *cpp.FunctionBodyContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *CppTreeShapeListener) EnterInitializer(ctx *cpp.InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *CppTreeShapeListener) ExitInitializer(ctx *cpp.InitializerContext) {}

// EnterBraceOrEqualInitializer is called when production braceOrEqualInitializer is entered.
func (s *CppTreeShapeListener) EnterBraceOrEqualInitializer(ctx *cpp.BraceOrEqualInitializerContext) {}

// ExitBraceOrEqualInitializer is called when production braceOrEqualInitializer is exited.
func (s *CppTreeShapeListener) ExitBraceOrEqualInitializer(ctx *cpp.BraceOrEqualInitializerContext) {}

// EnterInitializerClause is called when production initializerClause is entered.
func (s *CppTreeShapeListener) EnterInitializerClause(ctx *cpp.InitializerClauseContext) {}

// ExitInitializerClause is called when production initializerClause is exited.
func (s *CppTreeShapeListener) ExitInitializerClause(ctx *cpp.InitializerClauseContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *CppTreeShapeListener) EnterInitializerList(ctx *cpp.InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *CppTreeShapeListener) ExitInitializerList(ctx *cpp.InitializerListContext) {}

// EnterBracedInitList is called when production bracedInitList is entered.
func (s *CppTreeShapeListener) EnterBracedInitList(ctx *cpp.BracedInitListContext) {}

// ExitBracedInitList is called when production bracedInitList is exited.
func (s *CppTreeShapeListener) ExitBracedInitList(ctx *cpp.BracedInitListContext) {}

// EnterClassName is called when production className is entered.
func (s *CppTreeShapeListener) EnterClassName(ctx *cpp.ClassNameContext) {}

// ExitClassName is called when production className is exited.
func (s *CppTreeShapeListener) ExitClassName(ctx *cpp.ClassNameContext) {}

// ExitClassSpecifier is called when production classSpecifier is exited.
func (s *CppTreeShapeListener) ExitClassSpecifier(ctx *cpp.ClassSpecifierContext) {}

// EnterClassHead is called when production classHead is entered.
func (s *CppTreeShapeListener) EnterClassHead(ctx *cpp.ClassHeadContext) {
}

// ExitClassHead is called when production classHead is exited.
func (s *CppTreeShapeListener) ExitClassHead(ctx *cpp.ClassHeadContext) {}

// EnterClassHeadName is called when production classHeadName is entered.
func (s *CppTreeShapeListener) EnterClassHeadName(ctx *cpp.ClassHeadNameContext) {}

// ExitClassHeadName is called when production classHeadName is exited.
func (s *CppTreeShapeListener) ExitClassHeadName(ctx *cpp.ClassHeadNameContext) {}

// EnterClassVirtSpecifier is called when production classVirtSpecifier is entered.
func (s *CppTreeShapeListener) EnterClassVirtSpecifier(ctx *cpp.ClassVirtSpecifierContext) {}

// ExitClassVirtSpecifier is called when production classVirtSpecifier is exited.
func (s *CppTreeShapeListener) ExitClassVirtSpecifier(ctx *cpp.ClassVirtSpecifierContext) {}

// EnterClassKey is called when production classKey is entered.
func (s *CppTreeShapeListener) EnterClassKey(ctx *cpp.ClassKeyContext) {}

// ExitClassKey is called when production classKey is exited.
func (s *CppTreeShapeListener) ExitClassKey(ctx *cpp.ClassKeyContext) {}

// EnterMemberSpecification is called when production memberSpecification is entered.
func (s *CppTreeShapeListener) EnterMemberSpecification(ctx *cpp.MemberSpecificationContext) {}

// ExitMemberSpecification is called when production memberSpecification is exited.
func (s *CppTreeShapeListener) ExitMemberSpecification(ctx *cpp.MemberSpecificationContext) {}

// EnterMemberdeclaration is called when production memberdeclaration is entered.
func (s *CppTreeShapeListener) EnterMemberdeclaration(ctx *cpp.MemberdeclarationContext) {}

// ExitMemberdeclaration is called when production memberdeclaration is exited.
func (s *CppTreeShapeListener) ExitMemberdeclaration(ctx *cpp.MemberdeclarationContext) {}

// EnterMemberDeclaratorList is called when production memberDeclaratorList is entered.
func (s *CppTreeShapeListener) EnterMemberDeclaratorList(ctx *cpp.MemberDeclaratorListContext) {}

// ExitMemberDeclaratorList is called when production memberDeclaratorList is exited.
func (s *CppTreeShapeListener) ExitMemberDeclaratorList(ctx *cpp.MemberDeclaratorListContext) {}

// EnterMemberDeclarator is called when production memberDeclarator is entered.
func (s *CppTreeShapeListener) EnterMemberDeclarator(ctx *cpp.MemberDeclaratorContext) {}

// ExitMemberDeclarator is called when production memberDeclarator is exited.
func (s *CppTreeShapeListener) ExitMemberDeclarator(ctx *cpp.MemberDeclaratorContext) {}

// EnterVirtualSpecifierSeq is called when production virtualSpecifierSeq is entered.
func (s *CppTreeShapeListener) EnterVirtualSpecifierSeq(ctx *cpp.VirtualSpecifierSeqContext) {}

// ExitVirtualSpecifierSeq is called when production virtualSpecifierSeq is exited.
func (s *CppTreeShapeListener) ExitVirtualSpecifierSeq(ctx *cpp.VirtualSpecifierSeqContext) {}

// EnterVirtualSpecifier is called when production virtualSpecifier is entered.
func (s *CppTreeShapeListener) EnterVirtualSpecifier(ctx *cpp.VirtualSpecifierContext) {}

// ExitVirtualSpecifier is called when production virtualSpecifier is exited.
func (s *CppTreeShapeListener) ExitVirtualSpecifier(ctx *cpp.VirtualSpecifierContext) {}

// EnterPureSpecifier is called when production pureSpecifier is entered.
func (s *CppTreeShapeListener) EnterPureSpecifier(ctx *cpp.PureSpecifierContext) {}

// ExitPureSpecifier is called when production pureSpecifier is exited.
func (s *CppTreeShapeListener) ExitPureSpecifier(ctx *cpp.PureSpecifierContext) {}

// EnterBaseClause is called when production baseClause is entered.
func (s *CppTreeShapeListener) EnterBaseClause(ctx *cpp.BaseClauseContext) {}

// ExitBaseClause is called when production baseClause is exited.
func (s *CppTreeShapeListener) ExitBaseClause(ctx *cpp.BaseClauseContext) {}

// EnterBaseSpecifierList is called when production baseSpecifierList is entered.
func (s *CppTreeShapeListener) EnterBaseSpecifierList(ctx *cpp.BaseSpecifierListContext) {}

// ExitBaseSpecifierList is called when production baseSpecifierList is exited.
func (s *CppTreeShapeListener) ExitBaseSpecifierList(ctx *cpp.BaseSpecifierListContext) {}

// EnterBaseSpecifier is called when production baseSpecifier is entered.
func (s *CppTreeShapeListener) EnterBaseSpecifier(ctx *cpp.BaseSpecifierContext) {}

// ExitBaseSpecifier is called when production baseSpecifier is exited.
func (s *CppTreeShapeListener) ExitBaseSpecifier(ctx *cpp.BaseSpecifierContext) {}

// EnterClassOrDeclType is called when production classOrDeclType is entered.
func (s *CppTreeShapeListener) EnterClassOrDeclType(ctx *cpp.ClassOrDeclTypeContext) {}

// ExitClassOrDeclType is called when production classOrDeclType is exited.
func (s *CppTreeShapeListener) ExitClassOrDeclType(ctx *cpp.ClassOrDeclTypeContext) {}

// EnterBaseTypeSpecifier is called when production baseTypeSpecifier is entered.
func (s *CppTreeShapeListener) EnterBaseTypeSpecifier(ctx *cpp.BaseTypeSpecifierContext) {}

// ExitBaseTypeSpecifier is called when production baseTypeSpecifier is exited.
func (s *CppTreeShapeListener) ExitBaseTypeSpecifier(ctx *cpp.BaseTypeSpecifierContext) {}

// EnterAccessSpecifier is called when production accessSpecifier is entered.
func (s *CppTreeShapeListener) EnterAccessSpecifier(ctx *cpp.AccessSpecifierContext) {}

// ExitAccessSpecifier is called when production accessSpecifier is exited.
func (s *CppTreeShapeListener) ExitAccessSpecifier(ctx *cpp.AccessSpecifierContext) {}

// EnterConversionFunctionId is called when production conversionFunctionId is entered.
func (s *CppTreeShapeListener) EnterConversionFunctionId(ctx *cpp.ConversionFunctionIdContext) {}

// ExitConversionFunctionId is called when production conversionFunctionId is exited.
func (s *CppTreeShapeListener) ExitConversionFunctionId(ctx *cpp.ConversionFunctionIdContext) {}

// EnterConversionTypeId is called when production conversionTypeId is entered.
func (s *CppTreeShapeListener) EnterConversionTypeId(ctx *cpp.ConversionTypeIdContext) {}

// ExitConversionTypeId is called when production conversionTypeId is exited.
func (s *CppTreeShapeListener) ExitConversionTypeId(ctx *cpp.ConversionTypeIdContext) {}

// EnterConversionDeclarator is called when production conversionDeclarator is entered.
func (s *CppTreeShapeListener) EnterConversionDeclarator(ctx *cpp.ConversionDeclaratorContext) {}

// ExitConversionDeclarator is called when production conversionDeclarator is exited.
func (s *CppTreeShapeListener) ExitConversionDeclarator(ctx *cpp.ConversionDeclaratorContext) {}

// EnterConstructorInitializer is called when production constructorInitializer is entered.
func (s *CppTreeShapeListener) EnterConstructorInitializer(ctx *cpp.ConstructorInitializerContext) {}

// ExitConstructorInitializer is called when production constructorInitializer is exited.
func (s *CppTreeShapeListener) ExitConstructorInitializer(ctx *cpp.ConstructorInitializerContext) {}

// EnterMemInitializerList is called when production memInitializerList is entered.
func (s *CppTreeShapeListener) EnterMemInitializerList(ctx *cpp.MemInitializerListContext) {}

// ExitMemInitializerList is called when production memInitializerList is exited.
func (s *CppTreeShapeListener) ExitMemInitializerList(ctx *cpp.MemInitializerListContext) {}

// EnterMemInitializer is called when production memInitializer is entered.
func (s *CppTreeShapeListener) EnterMemInitializer(ctx *cpp.MemInitializerContext) {}

// ExitMemInitializer is called when production memInitializer is exited.
func (s *CppTreeShapeListener) ExitMemInitializer(ctx *cpp.MemInitializerContext) {}

// EnterMeminitializerid is called when production meminitializerid is entered.
func (s *CppTreeShapeListener) EnterMeminitializerid(ctx *cpp.MeminitializeridContext) {}

// ExitMeminitializerid is called when production meminitializerid is exited.
func (s *CppTreeShapeListener) ExitMeminitializerid(ctx *cpp.MeminitializeridContext) {}

// EnterOperatorFunctionId is called when production operatorFunctionId is entered.
func (s *CppTreeShapeListener) EnterOperatorFunctionId(ctx *cpp.OperatorFunctionIdContext) {}

// ExitOperatorFunctionId is called when production operatorFunctionId is exited.
func (s *CppTreeShapeListener) ExitOperatorFunctionId(ctx *cpp.OperatorFunctionIdContext) {}

// EnterLiteralOperatorId is called when production literalOperatorId is entered.
func (s *CppTreeShapeListener) EnterLiteralOperatorId(ctx *cpp.LiteralOperatorIdContext) {}

// ExitLiteralOperatorId is called when production literalOperatorId is exited.
func (s *CppTreeShapeListener) ExitLiteralOperatorId(ctx *cpp.LiteralOperatorIdContext) {}

// EnterTemplateDeclaration is called when production templateDeclaration is entered.
func (s *CppTreeShapeListener) EnterTemplateDeclaration(ctx *cpp.TemplateDeclarationContext) {}

// ExitTemplateDeclaration is called when production templateDeclaration is exited.
func (s *CppTreeShapeListener) ExitTemplateDeclaration(ctx *cpp.TemplateDeclarationContext) {}

// EnterTemplateparameterList is called when production templateparameterList is entered.
func (s *CppTreeShapeListener) EnterTemplateparameterList(ctx *cpp.TemplateparameterListContext) {}

// ExitTemplateparameterList is called when production templateparameterList is exited.
func (s *CppTreeShapeListener) ExitTemplateparameterList(ctx *cpp.TemplateparameterListContext) {}

// EnterTemplateParameter is called when production templateParameter is entered.
func (s *CppTreeShapeListener) EnterTemplateParameter(ctx *cpp.TemplateParameterContext) {}

// ExitTemplateParameter is called when production templateParameter is exited.
func (s *CppTreeShapeListener) ExitTemplateParameter(ctx *cpp.TemplateParameterContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *CppTreeShapeListener) EnterTypeParameter(ctx *cpp.TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *CppTreeShapeListener) ExitTypeParameter(ctx *cpp.TypeParameterContext) {}

// EnterSimpleTemplateId is called when production simpleTemplateId is entered.
func (s *CppTreeShapeListener) EnterSimpleTemplateId(ctx *cpp.SimpleTemplateIdContext) {}

// ExitSimpleTemplateId is called when production simpleTemplateId is exited.
func (s *CppTreeShapeListener) ExitSimpleTemplateId(ctx *cpp.SimpleTemplateIdContext) {}

// EnterTemplateId is called when production templateId is entered.
func (s *CppTreeShapeListener) EnterTemplateId(ctx *cpp.TemplateIdContext) {}

// ExitTemplateId is called when production templateId is exited.
func (s *CppTreeShapeListener) ExitTemplateId(ctx *cpp.TemplateIdContext) {}

// EnterTemplateName is called when production templateName is entered.
func (s *CppTreeShapeListener) EnterTemplateName(ctx *cpp.TemplateNameContext) {}

// ExitTemplateName is called when production templateName is exited.
func (s *CppTreeShapeListener) ExitTemplateName(ctx *cpp.TemplateNameContext) {}

// EnterTemplateArgumentList is called when production templateArgumentList is entered.
func (s *CppTreeShapeListener) EnterTemplateArgumentList(ctx *cpp.TemplateArgumentListContext) {}

// ExitTemplateArgumentList is called when production templateArgumentList is exited.
func (s *CppTreeShapeListener) ExitTemplateArgumentList(ctx *cpp.TemplateArgumentListContext) {}

// EnterTemplateArgument is called when production templateArgument is entered.
func (s *CppTreeShapeListener) EnterTemplateArgument(ctx *cpp.TemplateArgumentContext) {}

// ExitTemplateArgument is called when production templateArgument is exited.
func (s *CppTreeShapeListener) ExitTemplateArgument(ctx *cpp.TemplateArgumentContext) {}

// EnterTypeNameSpecifier is called when production typeNameSpecifier is entered.
func (s *CppTreeShapeListener) EnterTypeNameSpecifier(ctx *cpp.TypeNameSpecifierContext) {}

// ExitTypeNameSpecifier is called when production typeNameSpecifier is exited.
func (s *CppTreeShapeListener) ExitTypeNameSpecifier(ctx *cpp.TypeNameSpecifierContext) {}

// EnterExplicitInstantiation is called when production explicitInstantiation is entered.
func (s *CppTreeShapeListener) EnterExplicitInstantiation(ctx *cpp.ExplicitInstantiationContext) {}

// ExitExplicitInstantiation is called when production explicitInstantiation is exited.
func (s *CppTreeShapeListener) ExitExplicitInstantiation(ctx *cpp.ExplicitInstantiationContext) {}

// EnterExplicitSpecialization is called when production explicitSpecialization is entered.
func (s *CppTreeShapeListener) EnterExplicitSpecialization(ctx *cpp.ExplicitSpecializationContext) {}

// ExitExplicitSpecialization is called when production explicitSpecialization is exited.
func (s *CppTreeShapeListener) ExitExplicitSpecialization(ctx *cpp.ExplicitSpecializationContext) {}

// EnterTryBlock is called when production tryBlock is entered.
func (s *CppTreeShapeListener) EnterTryBlock(ctx *cpp.TryBlockContext) {}

// ExitTryBlock is called when production tryBlock is exited.
func (s *CppTreeShapeListener) ExitTryBlock(ctx *cpp.TryBlockContext) {}

// EnterFunctionTryBlock is called when production functionTryBlock is entered.
func (s *CppTreeShapeListener) EnterFunctionTryBlock(ctx *cpp.FunctionTryBlockContext) {}

// ExitFunctionTryBlock is called when production functionTryBlock is exited.
func (s *CppTreeShapeListener) ExitFunctionTryBlock(ctx *cpp.FunctionTryBlockContext) {}

// EnterHandlerSeq is called when production handlerSeq is entered.
func (s *CppTreeShapeListener) EnterHandlerSeq(ctx *cpp.HandlerSeqContext) {}

// ExitHandlerSeq is called when production handlerSeq is exited.
func (s *CppTreeShapeListener) ExitHandlerSeq(ctx *cpp.HandlerSeqContext) {}

// EnterHandler is called when production handler is entered.
func (s *CppTreeShapeListener) EnterHandler(ctx *cpp.HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *CppTreeShapeListener) ExitHandler(ctx *cpp.HandlerContext) {}

// EnterExceptionDeclaration is called when production exceptionDeclaration is entered.
func (s *CppTreeShapeListener) EnterExceptionDeclaration(ctx *cpp.ExceptionDeclarationContext) {}

// ExitExceptionDeclaration is called when production exceptionDeclaration is exited.
func (s *CppTreeShapeListener) ExitExceptionDeclaration(ctx *cpp.ExceptionDeclarationContext) {}

// EnterThrowExpression is called when production throwExpression is entered.
func (s *CppTreeShapeListener) EnterThrowExpression(ctx *cpp.ThrowExpressionContext) {}

// ExitThrowExpression is called when production throwExpression is exited.
func (s *CppTreeShapeListener) ExitThrowExpression(ctx *cpp.ThrowExpressionContext) {}

// EnterExceptionSpecification is called when production exceptionSpecification is entered.
func (s *CppTreeShapeListener) EnterExceptionSpecification(ctx *cpp.ExceptionSpecificationContext) {}

// ExitExceptionSpecification is called when production exceptionSpecification is exited.
func (s *CppTreeShapeListener) ExitExceptionSpecification(ctx *cpp.ExceptionSpecificationContext) {}

// EnterDynamicExceptionSpecification is called when production dynamicExceptionSpecification is entered.
func (s *CppTreeShapeListener) EnterDynamicExceptionSpecification(ctx *cpp.DynamicExceptionSpecificationContext) {
}

// ExitDynamicExceptionSpecification is called when production dynamicExceptionSpecification is exited.
func (s *CppTreeShapeListener) ExitDynamicExceptionSpecification(ctx *cpp.DynamicExceptionSpecificationContext) {
}

// EnterTypeIdList is called when production typeIdList is entered.
func (s *CppTreeShapeListener) EnterTypeIdList(ctx *cpp.TypeIdListContext) {}

// ExitTypeIdList is called when production typeIdList is exited.
func (s *CppTreeShapeListener) ExitTypeIdList(ctx *cpp.TypeIdListContext) {}

// EnterNoeExceptSpecification is called when production noeExceptSpecification is entered.
func (s *CppTreeShapeListener) EnterNoeExceptSpecification(ctx *cpp.NoeExceptSpecificationContext) {}

// ExitNoeExceptSpecification is called when production noeExceptSpecification is exited.
func (s *CppTreeShapeListener) ExitNoeExceptSpecification(ctx *cpp.NoeExceptSpecificationContext) {}

// EnterTheOperator is called when production theOperator is entered.
func (s *CppTreeShapeListener) EnterTheOperator(ctx *cpp.TheOperatorContext) {}

// ExitTheOperator is called when production theOperator is exited.
func (s *CppTreeShapeListener) ExitTheOperator(ctx *cpp.TheOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *CppTreeShapeListener) EnterLiteral(ctx *cpp.LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *CppTreeShapeListener) ExitLiteral(ctx *cpp.LiteralContext) {}

