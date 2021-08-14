// Generated from KotlinParser.g4 by ANTLR 4.7.

package antlrpack // KotlinParser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	kotlin "whosbugPack/antlrpack/kotlin_lib"
	"whosbugPack/utility"
)

var _ kotlin.KotlinParserListener = &KotlinTreeShapeListener{}

// ExitInfixFunctionCall is called when production infixFunctionCall is exited.
// 	@Description: 寻找方法调用
// 	@receiver s
// 	@param ctx
// 	@author KevinMatt 2021-08-12 23:41:08
// 	@function_mark
// ExitInfixFunctionCall is called when production infixFunctionCall is exited.
func (s *KotlinTreeShapeListener) ExitInfixFunctionCall(ctx *kotlin.InfixFunctionCallContext) {
	temp := ctx.GetChildren()
	for _, item := range temp {
		if _, ok := item.(*kotlin.SimpleIdentifierContext); ok {
			for index, item := range temp {
				fmt.Println(index, "", item.(antlr.ParseTree).GetText())
			}
			var methodCall CallMethodType
			methodCall.Id = item.(*kotlin.SimpleIdentifierContext).GetText()
			methodCall.StartLine = item.(*kotlin.SimpleIdentifierContext).GetStart().GetLine()
			s.MethodInfo.CallMethods = append(s.MethodInfo.CallMethods, methodCall.Id)
			return
		}
	}
}

// ExitFunctionBody is called when production functionBody is exited.
func (s *KotlinTreeShapeListener) ExitFunctionBody(ctx *kotlin.FunctionBodyContext) {
	if ctx.GetChildCount() == 1 { //正常函数声明
		Methods := ctx.GetChild(0).GetChild(1)
		for i := 0; i < Methods.GetChildCount(); i++ {
			method := Methods.GetChild(i).(antlr.ParseTree).GetText()
			if method != "\n" {
				s.MethodInfo.CallMethods = append(s.MethodInfo.CallMethods, method)
			}
		}
	} else if ctx.GetChildCount() == 2 { //单表示式函数
		s.MethodInfo.CallMethods = append(s.MethodInfo.CallMethods, ctx.GetChild(1).(antlr.ParseTree).GetText())
	} else {
		fmt.Println("Invaild Function")
	}
	s.Infos.AstInfoList.Methods = append(s.Infos.AstInfoList.Methods, s.MethodInfo)
}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterFunctionDeclaration(ctx *kotlin.FunctionDeclarationContext) {
	temp := ctx.GetChildren()

	s.MethodInfo.StartLine = ctx.GetStart().GetLine()
	for _, item := range temp {
		if _, ok := item.(*kotlin.SimpleIdentifierContext); ok {
			s.MethodInfo.MethodName = item.(*kotlin.SimpleIdentifierContext).GetText()
		}
	}
	s.MethodInfo.MasterObject = findKotlinMasterObject(ctx)
	utility.ForDebug()
}

// FindKotlinFuncCallIndex
// 	@Description:
// 	@receiver s
// 	@param targetStart
// 	@param targetEnd
// 	@return []string
// 	@author KevinMatt 2021-08-12 23:52:53
// 	@function_mark
func (s *KotlinTreeShapeListener) FindKotlinFuncCallIndex(targetStart, targetEnd int) []string {
	var resIndex []string
	for index := range s.Infos.CallMethods {
		if s.Infos.CallMethods[index].StartLine <= targetEnd && s.Infos.CallMethods[index].StartLine >= targetStart {
			resIndex = append(resIndex, s.Infos.CallMethods[index].Id)
		}
	}
	return resIndex
}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitFunctionDeclaration(ctx *kotlin.FunctionDeclarationContext) {
	// 方法信息置空
	s.MethodInfo = MethodInfoType{}
	temp := ctx.GetChildren()
	for _, item := range temp {
		if _, ok := item.(*kotlin.SimpleIdentifierContext); ok {
			s.MethodInfo.MethodName = item.(*kotlin.SimpleIdentifierContext).GetText()
			s.MethodInfo.StartLine = item.(*kotlin.SimpleIdentifierContext).GetStart().GetLine()
		}
	}
	utility.ForDebug()
}

// findKotlinMasterObject
// 	@Description: 寻找父类的算法
// 	@param ctx
// 	@return masterObjectInfoType
// 	@author KevinMatt 2021-08-11 18:59:06
// 	@function_mark PASS
func findKotlinMasterObject(ctx antlr.ParseTree) masterObjectInfoType {
	temp := ctx.GetParent().GetParent()
	if temp == nil {
		return masterObjectInfoType{}
	}
	var masterObject masterObjectInfoType
	for {
		if _, ok := temp.(*kotlin.DeclarationContext); ok {
			classDeclare := temp.GetChild(0)
			for index := range classDeclare.GetChildren() {
				if _, ok := classDeclare.GetChildren()[index].(*kotlin.SimpleIdentifierContext); ok {
					objectName := classDeclare.GetChildren()[index].(*kotlin.SimpleIdentifierContext).GetText()
					masterObject.ObjectName = objectName
					masterObject.StartLine = classDeclare.GetChildren()[index].(*kotlin.SimpleIdentifierContext).GetStart().GetLine()
					break
				}
			}
			return masterObject
		}
		temp = temp.GetParent()
		if temp == nil {
			return masterObjectInfoType{}
		}
	}
}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *KotlinTreeShapeListener) ExitClassDeclaration(ctx *kotlin.ClassDeclarationContext) {
	if ctx.GetChildCount() < 1 {
		return
	}
	var classInfo classInfoType
	classInfo.StartLine = ctx.GetStart().GetLine()
	classInfo.EndLine = ctx.GetStop().GetLine()
	temp := ctx.GetChildren()
	for index := range temp {
		if _, ok := temp[index].(*kotlin.SimpleIdentifierContext); ok {
			temp1 := temp[index].(*kotlin.SimpleIdentifierContext).GetText()
			classInfo.ClassName = temp1
			break
		}
	}
	classInfo.MasterObject = findKotlinMasterObject(ctx)
	s.Infos.AstInfoList.Classes = append(s.Infos.AstInfoList.Classes, classInfo)
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

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterClassDeclaration(ctx *kotlin.ClassDeclarationContext) {}

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

// EnterAnnotatedDelegationSpecifier is called when production annotatedDelegationSpecifier is entered.
func (s *KotlinTreeShapeListener) EnterAnnotatedDelegationSpecifier(ctx *kotlin.AnnotatedDelegationSpecifierContext) {
}

// ExitAnnotatedDelegationSpecifier is called when production annotatedDelegationSpecifier is exited.
func (s *KotlinTreeShapeListener) ExitAnnotatedDelegationSpecifier(ctx *kotlin.AnnotatedDelegationSpecifierContext) {
}

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

// EnterClassMemberDeclarations is called when production classMemberDeclarations is entered.
func (s *KotlinTreeShapeListener) EnterClassMemberDeclarations(ctx *kotlin.ClassMemberDeclarationsContext) {
}

// ExitClassMemberDeclarations is called when production classMemberDeclarations is exited.
func (s *KotlinTreeShapeListener) ExitClassMemberDeclarations(ctx *kotlin.ClassMemberDeclarationsContext) {
}

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
}

// ExitFunctionValueParameters is called when production functionValueParameters is exited.
func (s *KotlinTreeShapeListener) ExitFunctionValueParameters(ctx *kotlin.FunctionValueParametersContext) {
}

// EnterFunctionValueParameter is called when production functionValueParameter is entered.
func (s *KotlinTreeShapeListener) EnterFunctionValueParameter(ctx *kotlin.FunctionValueParameterContext) {
}

// ExitFunctionValueParameter is called when production functionValueParameter is exited.
func (s *KotlinTreeShapeListener) ExitFunctionValueParameter(ctx *kotlin.FunctionValueParameterContext) {
}

// EnterParameter is called when production parameter is entered.
func (s *KotlinTreeShapeListener) EnterParameter(ctx *kotlin.ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *KotlinTreeShapeListener) ExitParameter(ctx *kotlin.ParameterContext) {}

// EnterSetterParameter is called when production setterParameter is entered.
func (s *KotlinTreeShapeListener) EnterSetterParameter(ctx *kotlin.SetterParameterContext) {}

// ExitSetterParameter is called when production setterParameter is exited.
func (s *KotlinTreeShapeListener) ExitSetterParameter(ctx *kotlin.SetterParameterContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *KotlinTreeShapeListener) EnterFunctionBody(ctx *kotlin.FunctionBodyContext) {}

// EnterObjectDeclaration is called when production objectDeclaration is entered.
func (s *KotlinTreeShapeListener) EnterObjectDeclaration(ctx *kotlin.ObjectDeclarationContext) {}

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

// EnterPropertyDelegate is called when production propertyDelegate is entered.
func (s *KotlinTreeShapeListener) EnterPropertyDelegate(ctx *kotlin.PropertyDelegateContext) {}

// ExitPropertyDelegate is called when production propertyDelegate is exited.
func (s *KotlinTreeShapeListener) ExitPropertyDelegate(ctx *kotlin.PropertyDelegateContext) {}

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

// EnterTypeParameterModifiers is called when production typeParameterModifiers is entered.
func (s *KotlinTreeShapeListener) EnterTypeParameterModifiers(ctx *kotlin.TypeParameterModifiersContext) {
}

// ExitTypeParameterModifiers is called when production typeParameterModifiers is exited.
func (s *KotlinTreeShapeListener) ExitTypeParameterModifiers(ctx *kotlin.TypeParameterModifiersContext) {
}

// EnterTypeParameterModifier is called when production typeParameterModifier is entered.
func (s *KotlinTreeShapeListener) EnterTypeParameterModifier(ctx *kotlin.TypeParameterModifierContext) {
}

// ExitTypeParameterModifier is called when production typeParameterModifier is exited.
func (s *KotlinTreeShapeListener) ExitTypeParameterModifier(ctx *kotlin.TypeParameterModifierContext) {
}

// EnterType_ is called when production type_ is entered.
func (s *KotlinTreeShapeListener) EnterType_(ctx *kotlin.Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *KotlinTreeShapeListener) ExitType_(ctx *kotlin.Type_Context) {}

// EnterTypeModifiers is called when production typeModifiers is entered.
func (s *KotlinTreeShapeListener) EnterTypeModifiers(ctx *kotlin.TypeModifiersContext) {}

// ExitTypeModifiers is called when production typeModifiers is exited.
func (s *KotlinTreeShapeListener) ExitTypeModifiers(ctx *kotlin.TypeModifiersContext) {}

// EnterTypeModifier is called when production typeModifier is entered.
func (s *KotlinTreeShapeListener) EnterTypeModifier(ctx *kotlin.TypeModifierContext) {}

// ExitTypeModifier is called when production typeModifier is exited.
func (s *KotlinTreeShapeListener) ExitTypeModifier(ctx *kotlin.TypeModifierContext) {}

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

// EnterReceiverType is called when production receiverType is entered.
func (s *KotlinTreeShapeListener) EnterReceiverType(ctx *kotlin.ReceiverTypeContext) {}

// ExitReceiverType is called when production receiverType is exited.
func (s *KotlinTreeShapeListener) ExitReceiverType(ctx *kotlin.ReceiverTypeContext) {}

// EnterUserType is called when production userType is entered.
func (s *KotlinTreeShapeListener) EnterUserType(ctx *kotlin.UserTypeContext) {}

// ExitUserType is called when production userType is exited.
func (s *KotlinTreeShapeListener) ExitUserType(ctx *kotlin.UserTypeContext) {}

// EnterParenthesizedUserType is called when production parenthesizedUserType is entered.
func (s *KotlinTreeShapeListener) EnterParenthesizedUserType(ctx *kotlin.ParenthesizedUserTypeContext) {
}

// ExitParenthesizedUserType is called when production parenthesizedUserType is exited.
func (s *KotlinTreeShapeListener) ExitParenthesizedUserType(ctx *kotlin.ParenthesizedUserTypeContext) {
}

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

// EnterDeclaration is called when production declaration is entered.
func (s *KotlinTreeShapeListener) EnterDeclaration(ctx *kotlin.DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *KotlinTreeShapeListener) ExitDeclaration(ctx *kotlin.DeclarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *KotlinTreeShapeListener) EnterAssignment(ctx *kotlin.AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *KotlinTreeShapeListener) ExitAssignment(ctx *kotlin.AssignmentContext) {}

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

// EnterEquality is called when production equality is entered.
func (s *KotlinTreeShapeListener) EnterEquality(ctx *kotlin.EqualityContext) {}

// ExitEquality is called when production equality is exited.
func (s *KotlinTreeShapeListener) ExitEquality(ctx *kotlin.EqualityContext) {}

// EnterComparison is called when production comparison is entered.
func (s *KotlinTreeShapeListener) EnterComparison(ctx *kotlin.ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *KotlinTreeShapeListener) ExitComparison(ctx *kotlin.ComparisonContext) {}

// EnterInfixOperation is called when production infixOperation is entered.
func (s *KotlinTreeShapeListener) EnterInfixOperation(ctx *kotlin.InfixOperationContext) {}

// ExitInfixOperation is called when production infixOperation is exited.
func (s *KotlinTreeShapeListener) ExitInfixOperation(ctx *kotlin.InfixOperationContext) {}

// EnterElvisExpression is called when production elvisExpression is entered.
func (s *KotlinTreeShapeListener) EnterElvisExpression(ctx *kotlin.ElvisExpressionContext) {}

// ExitElvisExpression is called when production elvisExpression is exited.
func (s *KotlinTreeShapeListener) ExitElvisExpression(ctx *kotlin.ElvisExpressionContext) {}

// EnterInfixFunctionCall is called when production infixFunctionCall is entered.
func (s *KotlinTreeShapeListener) EnterInfixFunctionCall(ctx *kotlin.InfixFunctionCallContext) {}

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

// EnterAsExpression is called when production asExpression is entered.
func (s *KotlinTreeShapeListener) EnterAsExpression(ctx *kotlin.AsExpressionContext) {}

// ExitAsExpression is called when production asExpression is exited.
func (s *KotlinTreeShapeListener) ExitAsExpression(ctx *kotlin.AsExpressionContext) {}

// EnterPrefixUnaryExpression is called when production prefixUnaryExpression is entered.
func (s *KotlinTreeShapeListener) EnterPrefixUnaryExpression(ctx *kotlin.PrefixUnaryExpressionContext) {
}

// ExitPrefixUnaryExpression is called when production prefixUnaryExpression is exited.
func (s *KotlinTreeShapeListener) ExitPrefixUnaryExpression(ctx *kotlin.PrefixUnaryExpressionContext) {
}

// EnterUnaryPrefix is called when production unaryPrefix is entered.
func (s *KotlinTreeShapeListener) EnterUnaryPrefix(ctx *kotlin.UnaryPrefixContext) {}

// ExitUnaryPrefix is called when production unaryPrefix is exited.
func (s *KotlinTreeShapeListener) ExitUnaryPrefix(ctx *kotlin.UnaryPrefixContext) {}

// EnterPostfixUnaryExpression is called when production postfixUnaryExpression is entered.
func (s *KotlinTreeShapeListener) EnterPostfixUnaryExpression(ctx *kotlin.PostfixUnaryExpressionContext) {
}

// ExitPostfixUnaryExpression is called when production postfixUnaryExpression is exited.
func (s *KotlinTreeShapeListener) ExitPostfixUnaryExpression(ctx *kotlin.PostfixUnaryExpressionContext) {
}

// EnterPostfixUnarySuffix is called when production postfixUnarySuffix is entered.
func (s *KotlinTreeShapeListener) EnterPostfixUnarySuffix(ctx *kotlin.PostfixUnarySuffixContext) {}

// ExitPostfixUnarySuffix is called when production postfixUnarySuffix is exited.
func (s *KotlinTreeShapeListener) ExitPostfixUnarySuffix(ctx *kotlin.PostfixUnarySuffixContext) {}

// EnterDirectlyAssignableExpression is called when production directlyAssignableExpression is entered.
func (s *KotlinTreeShapeListener) EnterDirectlyAssignableExpression(ctx *kotlin.DirectlyAssignableExpressionContext) {
}

// ExitDirectlyAssignableExpression is called when production directlyAssignableExpression is exited.
func (s *KotlinTreeShapeListener) ExitDirectlyAssignableExpression(ctx *kotlin.DirectlyAssignableExpressionContext) {
}

// EnterAssignableExpression is called when production assignableExpression is entered.
func (s *KotlinTreeShapeListener) EnterAssignableExpression(ctx *kotlin.AssignableExpressionContext) {
}

// ExitAssignableExpression is called when production assignableExpression is exited.
func (s *KotlinTreeShapeListener) ExitAssignableExpression(ctx *kotlin.AssignableExpressionContext) {}

// EnterAssignableSuffix is called when production assignableSuffix is entered.
func (s *KotlinTreeShapeListener) EnterAssignableSuffix(ctx *kotlin.AssignableSuffixContext) {}

// ExitAssignableSuffix is called when production assignableSuffix is exited.
func (s *KotlinTreeShapeListener) ExitAssignableSuffix(ctx *kotlin.AssignableSuffixContext) {}

// EnterIndexingSuffix is called when production indexingSuffix is entered.
func (s *KotlinTreeShapeListener) EnterIndexingSuffix(ctx *kotlin.IndexingSuffixContext) {}

// ExitIndexingSuffix is called when production indexingSuffix is exited.
func (s *KotlinTreeShapeListener) ExitIndexingSuffix(ctx *kotlin.IndexingSuffixContext) {}

// EnterNavigationSuffix is called when production navigationSuffix is entered.
func (s *KotlinTreeShapeListener) EnterNavigationSuffix(ctx *kotlin.NavigationSuffixContext) {}

// ExitNavigationSuffix is called when production navigationSuffix is exited.
func (s *KotlinTreeShapeListener) ExitNavigationSuffix(ctx *kotlin.NavigationSuffixContext) {}

// EnterCallSuffix is called when production callSuffix is entered.
func (s *KotlinTreeShapeListener) EnterCallSuffix(ctx *kotlin.CallSuffixContext) {}

// ExitCallSuffix is called when production callSuffix is exited.
func (s *KotlinTreeShapeListener) ExitCallSuffix(ctx *kotlin.CallSuffixContext) {}

// EnterAnnotatedLambda is called when production annotatedLambda is entered.
func (s *KotlinTreeShapeListener) EnterAnnotatedLambda(ctx *kotlin.AnnotatedLambdaContext) {}

// ExitAnnotatedLambda is called when production annotatedLambda is exited.
func (s *KotlinTreeShapeListener) ExitAnnotatedLambda(ctx *kotlin.AnnotatedLambdaContext) {}

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

// EnterTypeProjectionModifiers is called when production typeProjectionModifiers is entered.
func (s *KotlinTreeShapeListener) EnterTypeProjectionModifiers(ctx *kotlin.TypeProjectionModifiersContext) {
}

// ExitTypeProjectionModifiers is called when production typeProjectionModifiers is exited.
func (s *KotlinTreeShapeListener) ExitTypeProjectionModifiers(ctx *kotlin.TypeProjectionModifiersContext) {
}

// EnterTypeProjectionModifier is called when production typeProjectionModifier is entered.
func (s *KotlinTreeShapeListener) EnterTypeProjectionModifier(ctx *kotlin.TypeProjectionModifierContext) {
}

// ExitTypeProjectionModifier is called when production typeProjectionModifier is exited.
func (s *KotlinTreeShapeListener) ExitTypeProjectionModifier(ctx *kotlin.TypeProjectionModifierContext) {
}

// EnterValueArgument is called when production valueArgument is entered.
func (s *KotlinTreeShapeListener) EnterValueArgument(ctx *kotlin.ValueArgumentContext) {}

// ExitValueArgument is called when production valueArgument is exited.
func (s *KotlinTreeShapeListener) ExitValueArgument(ctx *kotlin.ValueArgumentContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *KotlinTreeShapeListener) EnterPrimaryExpression(ctx *kotlin.PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *KotlinTreeShapeListener) ExitPrimaryExpression(ctx *kotlin.PrimaryExpressionContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *KotlinTreeShapeListener) EnterParenthesizedExpression(ctx *kotlin.ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *KotlinTreeShapeListener) ExitParenthesizedExpression(ctx *kotlin.ParenthesizedExpressionContext) {
}

// EnterCollectionLiteral is called when production collectionLiteral is entered.
func (s *KotlinTreeShapeListener) EnterCollectionLiteral(ctx *kotlin.CollectionLiteralContext) {}

// ExitCollectionLiteral is called when production collectionLiteral is exited.
func (s *KotlinTreeShapeListener) ExitCollectionLiteral(ctx *kotlin.CollectionLiteralContext) {}

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

// EnterLambdaLiteral is called when production lambdaLiteral is entered.
func (s *KotlinTreeShapeListener) EnterLambdaLiteral(ctx *kotlin.LambdaLiteralContext) {}

// ExitLambdaLiteral is called when production lambdaLiteral is exited.
func (s *KotlinTreeShapeListener) ExitLambdaLiteral(ctx *kotlin.LambdaLiteralContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *KotlinTreeShapeListener) EnterLambdaParameters(ctx *kotlin.LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *KotlinTreeShapeListener) ExitLambdaParameters(ctx *kotlin.LambdaParametersContext) {}

// EnterLambdaParameter is called when production lambdaParameter is entered.
func (s *KotlinTreeShapeListener) EnterLambdaParameter(ctx *kotlin.LambdaParameterContext) {}

// ExitLambdaParameter is called when production lambdaParameter is exited.
func (s *KotlinTreeShapeListener) ExitLambdaParameter(ctx *kotlin.LambdaParameterContext) {}

// EnterAnonymousFunction is called when production anonymousFunction is entered.
func (s *KotlinTreeShapeListener) EnterAnonymousFunction(ctx *kotlin.AnonymousFunctionContext) {}

// ExitAnonymousFunction is called when production anonymousFunction is exited.
func (s *KotlinTreeShapeListener) ExitAnonymousFunction(ctx *kotlin.AnonymousFunctionContext) {}

// EnterFunctionLiteral is called when production functionLiteral is entered.
func (s *KotlinTreeShapeListener) EnterFunctionLiteral(ctx *kotlin.FunctionLiteralContext) {}

// ExitFunctionLiteral is called when production functionLiteral is exited.
func (s *KotlinTreeShapeListener) ExitFunctionLiteral(ctx *kotlin.FunctionLiteralContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *KotlinTreeShapeListener) EnterObjectLiteral(ctx *kotlin.ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *KotlinTreeShapeListener) ExitObjectLiteral(ctx *kotlin.ObjectLiteralContext) {}

// EnterThisExpression is called when production thisExpression is entered.
func (s *KotlinTreeShapeListener) EnterThisExpression(ctx *kotlin.ThisExpressionContext) {}

// ExitThisExpression is called when production thisExpression is exited.
func (s *KotlinTreeShapeListener) ExitThisExpression(ctx *kotlin.ThisExpressionContext) {}

// EnterSuperExpression is called when production superExpression is entered.
func (s *KotlinTreeShapeListener) EnterSuperExpression(ctx *kotlin.SuperExpressionContext) {}

// ExitSuperExpression is called when production superExpression is exited.
func (s *KotlinTreeShapeListener) ExitSuperExpression(ctx *kotlin.SuperExpressionContext) {}

// EnterControlStructureBody is called when production controlStructureBody is entered.
func (s *KotlinTreeShapeListener) EnterControlStructureBody(ctx *kotlin.ControlStructureBodyContext) {
}

// ExitControlStructureBody is called when production controlStructureBody is exited.
func (s *KotlinTreeShapeListener) ExitControlStructureBody(ctx *kotlin.ControlStructureBodyContext) {}

// EnterIfExpression is called when production ifExpression is entered.
func (s *KotlinTreeShapeListener) EnterIfExpression(ctx *kotlin.IfExpressionContext) {}

// ExitIfExpression is called when production ifExpression is exited.
func (s *KotlinTreeShapeListener) ExitIfExpression(ctx *kotlin.IfExpressionContext) {}

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

// EnterLoopStatement is called when production loopStatement is entered.
func (s *KotlinTreeShapeListener) EnterLoopStatement(ctx *kotlin.LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *KotlinTreeShapeListener) ExitLoopStatement(ctx *kotlin.LoopStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *KotlinTreeShapeListener) EnterForStatement(ctx *kotlin.ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *KotlinTreeShapeListener) ExitForStatement(ctx *kotlin.ForStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *KotlinTreeShapeListener) EnterWhileStatement(ctx *kotlin.WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *KotlinTreeShapeListener) ExitWhileStatement(ctx *kotlin.WhileStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *KotlinTreeShapeListener) EnterDoWhileStatement(ctx *kotlin.DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *KotlinTreeShapeListener) ExitDoWhileStatement(ctx *kotlin.DoWhileStatementContext) {}

// EnterJumpExpression is called when production jumpExpression is entered.
func (s *KotlinTreeShapeListener) EnterJumpExpression(ctx *kotlin.JumpExpressionContext) {}

// ExitJumpExpression is called when production jumpExpression is exited.
func (s *KotlinTreeShapeListener) ExitJumpExpression(ctx *kotlin.JumpExpressionContext) {}

// EnterCallableReference is called when production callableReference is entered.
func (s *KotlinTreeShapeListener) EnterCallableReference(ctx *kotlin.CallableReferenceContext) {}

// ExitCallableReference is called when production callableReference is exited.
func (s *KotlinTreeShapeListener) ExitCallableReference(ctx *kotlin.CallableReferenceContext) {}

// EnterAssignmentAndOperator is called when production assignmentAndOperator is entered.
func (s *KotlinTreeShapeListener) EnterAssignmentAndOperator(ctx *kotlin.AssignmentAndOperatorContext) {
}

// ExitAssignmentAndOperator is called when production assignmentAndOperator is exited.
func (s *KotlinTreeShapeListener) ExitAssignmentAndOperator(ctx *kotlin.AssignmentAndOperatorContext) {
}

// EnterEqualityOperator is called when production equalityOperator is entered.
func (s *KotlinTreeShapeListener) EnterEqualityOperator(ctx *kotlin.EqualityOperatorContext) {}

// ExitEqualityOperator is called when production equalityOperator is exited.
func (s *KotlinTreeShapeListener) ExitEqualityOperator(ctx *kotlin.EqualityOperatorContext) {}

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

// EnterMultiplicativeOperator is called when production multiplicativeOperator is entered.
func (s *KotlinTreeShapeListener) EnterMultiplicativeOperator(ctx *kotlin.MultiplicativeOperatorContext) {
}

// ExitMultiplicativeOperator is called when production multiplicativeOperator is exited.
func (s *KotlinTreeShapeListener) ExitMultiplicativeOperator(ctx *kotlin.MultiplicativeOperatorContext) {
}

// EnterAsOperator is called when production asOperator is entered.
func (s *KotlinTreeShapeListener) EnterAsOperator(ctx *kotlin.AsOperatorContext) {}

// ExitAsOperator is called when production asOperator is exited.
func (s *KotlinTreeShapeListener) ExitAsOperator(ctx *kotlin.AsOperatorContext) {}

// EnterPrefixUnaryOperator is called when production prefixUnaryOperator is entered.
func (s *KotlinTreeShapeListener) EnterPrefixUnaryOperator(ctx *kotlin.PrefixUnaryOperatorContext) {}

// ExitPrefixUnaryOperator is called when production prefixUnaryOperator is exited.
func (s *KotlinTreeShapeListener) ExitPrefixUnaryOperator(ctx *kotlin.PrefixUnaryOperatorContext) {}

// EnterPostfixUnaryOperator is called when production postfixUnaryOperator is entered.
func (s *KotlinTreeShapeListener) EnterPostfixUnaryOperator(ctx *kotlin.PostfixUnaryOperatorContext) {
}

// ExitPostfixUnaryOperator is called when production postfixUnaryOperator is exited.
func (s *KotlinTreeShapeListener) ExitPostfixUnaryOperator(ctx *kotlin.PostfixUnaryOperatorContext) {}

// EnterMemberAccessOperator is called when production memberAccessOperator is entered.
func (s *KotlinTreeShapeListener) EnterMemberAccessOperator(ctx *kotlin.MemberAccessOperatorContext) {
}

// ExitMemberAccessOperator is called when production memberAccessOperator is exited.
func (s *KotlinTreeShapeListener) ExitMemberAccessOperator(ctx *kotlin.MemberAccessOperatorContext) {}

// EnterModifiers is called when production modifiers is entered.
func (s *KotlinTreeShapeListener) EnterModifiers(ctx *kotlin.ModifiersContext) {}

// ExitModifiers is called when production modifiers is exited.
func (s *KotlinTreeShapeListener) ExitModifiers(ctx *kotlin.ModifiersContext) {}

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

// EnterVarianceModifier is called when production varianceModifier is entered.
func (s *KotlinTreeShapeListener) EnterVarianceModifier(ctx *kotlin.VarianceModifierContext) {}

// ExitVarianceModifier is called when production varianceModifier is exited.
func (s *KotlinTreeShapeListener) ExitVarianceModifier(ctx *kotlin.VarianceModifierContext) {}

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

// EnterReificationModifier is called when production reificationModifier is entered.
func (s *KotlinTreeShapeListener) EnterReificationModifier(ctx *kotlin.ReificationModifierContext) {}

// ExitReificationModifier is called when production reificationModifier is exited.
func (s *KotlinTreeShapeListener) ExitReificationModifier(ctx *kotlin.ReificationModifierContext) {}

// EnterPlatformModifier is called when production platformModifier is entered.
func (s *KotlinTreeShapeListener) EnterPlatformModifier(ctx *kotlin.PlatformModifierContext) {}

// ExitPlatformModifier is called when production platformModifier is exited.
func (s *KotlinTreeShapeListener) ExitPlatformModifier(ctx *kotlin.PlatformModifierContext) {}

// EnterLabel is called when production label is entered.
func (s *KotlinTreeShapeListener) EnterLabel(ctx *kotlin.LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *KotlinTreeShapeListener) ExitLabel(ctx *kotlin.LabelContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *KotlinTreeShapeListener) EnterAnnotation(ctx *kotlin.AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *KotlinTreeShapeListener) ExitAnnotation(ctx *kotlin.AnnotationContext) {}

// EnterSingleAnnotation is called when production singleAnnotation is entered.
func (s *KotlinTreeShapeListener) EnterSingleAnnotation(ctx *kotlin.SingleAnnotationContext) {}

// ExitSingleAnnotation is called when production singleAnnotation is exited.
func (s *KotlinTreeShapeListener) ExitSingleAnnotation(ctx *kotlin.SingleAnnotationContext) {}

// EnterMultiAnnotation is called when production multiAnnotation is entered.
func (s *KotlinTreeShapeListener) EnterMultiAnnotation(ctx *kotlin.MultiAnnotationContext) {}

// ExitMultiAnnotation is called when production multiAnnotation is exited.
func (s *KotlinTreeShapeListener) ExitMultiAnnotation(ctx *kotlin.MultiAnnotationContext) {}

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

// EnterSimpleIdentifier is called when production simpleIdentifier is entered.
func (s *KotlinTreeShapeListener) EnterSimpleIdentifier(ctx *kotlin.SimpleIdentifierContext) {}

// ExitSimpleIdentifier is called when production simpleIdentifier is exited.
func (s *KotlinTreeShapeListener) ExitSimpleIdentifier(ctx *kotlin.SimpleIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *KotlinTreeShapeListener) EnterIdentifier(ctx *kotlin.IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *KotlinTreeShapeListener) ExitIdentifier(ctx *kotlin.IdentifierContext) {}

// EnterShebangLine is called when production shebangLine is entered.
func (s *KotlinTreeShapeListener) EnterShebangLine(ctx *kotlin.ShebangLineContext) {}

// ExitShebangLine is called when production shebangLine is exited.
func (s *KotlinTreeShapeListener) ExitShebangLine(ctx *kotlin.ShebangLineContext) {}

// EnterQuest is called when production quest is entered.
func (s *KotlinTreeShapeListener) EnterQuest(ctx *kotlin.QuestContext) {}

// ExitQuest is called when production quest is exited.
func (s *KotlinTreeShapeListener) ExitQuest(ctx *kotlin.QuestContext) {}

// EnterElvis is called when production elvis is entered.
func (s *KotlinTreeShapeListener) EnterElvis(ctx *kotlin.ElvisContext) {}

// ExitElvis is called when production elvis is exited.
func (s *KotlinTreeShapeListener) ExitElvis(ctx *kotlin.ElvisContext) {}

// EnterSafeNav is called when production safeNav is entered.
func (s *KotlinTreeShapeListener) EnterSafeNav(ctx *kotlin.SafeNavContext) {}

// ExitSafeNav is called when production safeNav is exited.
func (s *KotlinTreeShapeListener) ExitSafeNav(ctx *kotlin.SafeNavContext) {}

// EnterExcl is called when production excl is entered.
func (s *KotlinTreeShapeListener) EnterExcl(ctx *kotlin.ExclContext) {}

// ExitExcl is called when production excl is exited.
func (s *KotlinTreeShapeListener) ExitExcl(ctx *kotlin.ExclContext) {}

// EnterSemi is called when production semi is entered.
func (s *KotlinTreeShapeListener) EnterSemi(ctx *kotlin.SemiContext) {}

// ExitSemi is called when production semi is exited.
func (s *KotlinTreeShapeListener) ExitSemi(ctx *kotlin.SemiContext) {}

// EnterSemis is called when production semis is entered.
func (s *KotlinTreeShapeListener) EnterSemis(ctx *kotlin.SemisContext) {}

// ExitSemis is called when production semis is exited.
func (s *KotlinTreeShapeListener) ExitSemis(ctx *kotlin.SemisContext) {}
