// Code generated from JavaParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package javaParser // JavaParser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type paramInfoType struct {
	paramType string
	paramName string
}

type methodInfoType struct {
	StartLine    int
	EndLine      int
	ReturnType   string
	MethodName   string
	Params       []paramInfoType
	MasterObject string
	Depth        int
	CallMethods  []string
}
type classInfoType struct {
	StartLine    int
	EndLine      int
	ClassName    string
	Extends      string
	Implements   []string
	MasterObject string
	Depth        int
}
type astInfoType struct {
	PackageName string
	Classes     []classInfoType
	Imports     []string
	Fileds      []string
	Methods     []methodInfoType
}

type analysisInfoType struct {
	CallMethods []string
	AstInfoList astInfoType
}

var Infos analysisInfoType

/** ExitMethodDeclaration
 * @Description: 匹配到方法结束时被调用
 * @receiver s
 * @param ctx
 * @author KevinMatt 2021-07-23 23:14:09
 * @function_mark PASS
 */
func (s *BaseJavaParserListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {
	var methodInfo methodInfoType
	MethodName := ctx.GetChild(1)
	ReturnType := ctx.GetChild(0).(*TypeTypeOrVoidContext).GetText()
	Params := getParams(ctx)
	//fmt.Println(Params[0].paramName, Params[0].paramType)
	methodInfo.ReturnType = ReturnType
	methodInfo.StartLine = ctx.GetStart().GetLine()
	methodInfo.EndLine = ctx.GetStop().GetLine()
	methodInfo.MethodName = fmt.Sprintf("%s", MethodName)
	methodInfo.MasterObject = ""
	methodInfo.Params = append(methodInfo.Params, Params...)
	//TODO 解决语法树深度的高效计算方法
	Infos.AstInfoList.Methods = append(Infos.AstInfoList.Methods, methodInfo)
}

/** getParams
 * @Description: 获取参数名&参数类型结构体的切片
 * @param ctx *MethodDeclarationContext
 * @return []paramInfoType 返回追加后的切片
 * @author KevinMatt 2021-07-23 22:55:22
 * @function_mark
 */
func getParams(ctx *MethodDeclarationContext) []paramInfoType {
	var paramsCount int
	var paramInfo paramInfoType
	var result []paramInfoType
	if ctx.GetChildCount() == 3 {
		paramsCount = ctx.GetChild(1).GetChildCount()
		if paramsCount == 1 {
			paramType := ctx.GetChild(1).GetChild(0).GetChild(0).(*FormalParameterContext).GetText()
			paramName := ctx.GetChild(1).GetChild(0).GetChild(1).(*FormalParameterContext).GetText()
			paramInfo.paramName = paramName
			paramInfo.paramType = paramType
			result = append(result, paramInfo)
		} else if paramsCount > 1 {
			for index := 0; index < paramsCount; index++ {
				if index%2 == 0 {
					paramType := ctx.GetChild(1).GetChild(index).GetChild(0).(*FormalParameterContext).GetText()
					paramName := ctx.GetChild(1).GetChild(index).GetChild(1).(*FormalParameterContext).GetText()
					paramInfo.paramName = paramName
					paramInfo.paramType = paramType
					result = append(result, paramInfo)
				}
			}
		}
	}
	return result
}

/** EnterMethodCall
 * @Description: 匹配调用方法行并获取起始行列号
 * @receiver s
 * @param ctx
 * @author KevinMatt 2021-07-23 23:22:56
 * @function_mark PASS
 */
func (s *BaseJavaParserListener) EnterMethodCall(ctx *MethodCallContext) {
	lineNumber := ctx.GetStart().GetLine()
	columnNumber := ctx.GetStart().GetColumn()
	Infos.CallMethods = append(Infos.CallMethods, fmt.Sprintf("%s %s %s", strconv.Itoa(lineNumber), strconv.Itoa(columnNumber), ctx.GetParent().(*ExpressionContext).GetText()))
}

// BaseJavaParserListener is a complete listener for a parse tree produced by JavaParser.
type BaseJavaParserListener struct{}

var _ JavaParserListener = &BaseJavaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJavaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJavaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJavaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJavaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseJavaParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseJavaParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPackageDeclaration is called when production packageDeclaration is entered.
// 获得读取到的package声明语句
func (s *BaseJavaParserListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	//fmt.Println(ctx.GetText())
	//InfoStruct.analysisInfoType["ast_info"].(map[string]string)["PackageName"] = ctx.QualifiedName().GetText()
}

// ExitPackageDeclaration is called when production packageDeclaration is exited.
func (s *BaseJavaParserListener) ExitPackageDeclaration(ctx *PackageDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
// 获得读取到的import声明语句
func (s *BaseJavaParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	//fmt.Println(ctx.GetText())
	//importMatch := ctx.QualifiedName().GetText()
	//InfoStruct.analysisInfoType["ast_info"].(map[string][]string)["Imports"] = append(InfoStruct.analysisInfoType["ast_info"].(map[string][]string)["Imports"], importMatch)
}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseJavaParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseJavaParserListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseJavaParserListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseJavaParserListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseJavaParserListener) ExitModifier(ctx *ModifierContext) {}

// EnterClassOrInterfaceModifier is called when production classOrInterfaceModifier is entered.
func (s *BaseJavaParserListener) EnterClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {
}

// ExitClassOrInterfaceModifier is called when production classOrInterfaceModifier is exited.
func (s *BaseJavaParserListener) ExitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BaseJavaParserListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BaseJavaParserListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseJavaParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	var classInfo classInfoType
	classInfo.StartLine = ctx.GetStart().GetLine()
	if ctx.ClassBody() != nil {
		classInfo.EndLine = ctx.ClassBody().GetStop().GetLine()
	}
	childCount := ctx.GetChildCount()

	if childCount == 7 {
		// class Foo extends Bar implements Hoge
		// c1 = ctx.getChild(0)  # ---> class
		className := ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
		extendsClassName := ctx.GetChild(3).GetChild(0).(*antlr.TerminalNodeImpl).GetText()
		classInfo.ClassName = className
		classInfo.Extends = extendsClassName
		classInfo.Implements = findImplements(ctx.GetChild(5).(*antlr.TerminalNodeImpl))
		classInfo.MasterObject = findMasterObject()
	}
}

func findImplements(ctx *antlr.TerminalNodeImpl) []string {
	implementsCount := ctx.GetChildCount()
	var implements []string
	if implementsCount == 1 {
		implementClass := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
		implements = append(implements, implementClass)
	} else if implementsCount > 1 {
		for index := 0; index < implementsCount; index++ {
			if index%2 == 0 {
				implementClass := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
				implements = append(implements, implementClass)
			}
		}
	}
	return implements
}
func findMasterObject(ctx *ClassDeclarationContext, classInfo classInfoType) {
	// TODO 重构实现
	//parCtx := ctx.GetParent()
	//masterCtx := parCtx.GetParent().GetParent().GetParent()

}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseJavaParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseJavaParserListener) EnterTypeParameters(ctx *TypeParametersContext) {
}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseJavaParserListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseJavaParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseJavaParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeBound is called when production typeBound is entered.
func (s *BaseJavaParserListener) EnterTypeBound(ctx *TypeBoundContext) {}

// ExitTypeBound is called when production typeBound is exited.
func (s *BaseJavaParserListener) ExitTypeBound(ctx *TypeBoundContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseJavaParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseJavaParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumConstants is called when production enumConstants is entered.
func (s *BaseJavaParserListener) EnterEnumConstants(ctx *EnumConstantsContext) {}

// ExitEnumConstants is called when production enumConstants is exited.
func (s *BaseJavaParserListener) ExitEnumConstants(ctx *EnumConstantsContext) {}

// EnterEnumConstant is called when production enumConstant is entered.
func (s *BaseJavaParserListener) EnterEnumConstant(ctx *EnumConstantContext) {}

// ExitEnumConstant is called when production enumConstant is exited.
func (s *BaseJavaParserListener) ExitEnumConstant(ctx *EnumConstantContext) {}

// EnterEnumBodyDeclarations is called when production enumBodyDeclarations is entered.
func (s *BaseJavaParserListener) EnterEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// ExitEnumBodyDeclarations is called when production enumBodyDeclarations is exited.
func (s *BaseJavaParserListener) ExitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseJavaParserListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseJavaParserListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseJavaParserListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseJavaParserListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterInterfaceBody is called when production interfaceBody is entered.
func (s *BaseJavaParserListener) EnterInterfaceBody(ctx *InterfaceBodyContext) {}

// ExitInterfaceBody is called when production interfaceBody is exited.
func (s *BaseJavaParserListener) ExitInterfaceBody(ctx *InterfaceBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseJavaParserListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseJavaParserListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseJavaParserListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseJavaParserListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseJavaParserListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	//fmt.Println("Method Declaration:", ctx.GetText())
}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseJavaParserListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseJavaParserListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterTypeTypeOrVoid is called when production typeTypeOrVoid is entered.
func (s *BaseJavaParserListener) EnterTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {
	//fmt.Printf("Type: %s \n", ctx.GetText())
}

// ExitTypeTypeOrVoid is called when production typeTypeOrVoid is exited.
func (s *BaseJavaParserListener) ExitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// EnterGenericMethodDeclaration is called when production genericMethodDeclaration is entered.
func (s *BaseJavaParserListener) EnterGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) {
}

// ExitGenericMethodDeclaration is called when production genericMethodDeclaration is exited.
func (s *BaseJavaParserListener) ExitGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) {}

// EnterGenericConstructorDeclaration is called when production genericConstructorDeclaration is entered.
func (s *BaseJavaParserListener) EnterGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) {
}

// ExitGenericConstructorDeclaration is called when production genericConstructorDeclaration is exited.
func (s *BaseJavaParserListener) ExitGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) {
}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BaseJavaParserListener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BaseJavaParserListener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseJavaParserListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseJavaParserListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is entered.
func (s *BaseJavaParserListener) EnterInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {
}

// ExitInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is exited.
func (s *BaseJavaParserListener) ExitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {}

// EnterInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is entered.
func (s *BaseJavaParserListener) EnterInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {
}

// ExitInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is exited.
func (s *BaseJavaParserListener) ExitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {
}

// EnterConstDeclaration is called when production constDeclaration is entered.
func (s *BaseJavaParserListener) EnterConstDeclaration(ctx *ConstDeclarationContext) {}

// ExitConstDeclaration is called when production constDeclaration is exited.
func (s *BaseJavaParserListener) ExitConstDeclaration(ctx *ConstDeclarationContext) {}

// EnterConstantDeclarator is called when production constantDeclarator is entered.
func (s *BaseJavaParserListener) EnterConstantDeclarator(ctx *ConstantDeclaratorContext) {}

// ExitConstantDeclarator is called when production constantDeclarator is exited.
func (s *BaseJavaParserListener) ExitConstantDeclarator(ctx *ConstantDeclaratorContext) {}

// EnterInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is entered.
func (s *BaseJavaParserListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
}

// ExitInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is exited.
func (s *BaseJavaParserListener) ExitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
}

// EnterInterfaceMethodModifier is called when production interfaceMethodModifier is entered.
func (s *BaseJavaParserListener) EnterInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) {}

// ExitInterfaceMethodModifier is called when production interfaceMethodModifier is exited.
func (s *BaseJavaParserListener) ExitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) {}

// EnterGenericInterfaceMethodDeclaration is called when production genericInterfaceMethodDeclaration is entered.
func (s *BaseJavaParserListener) EnterGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) {
}

// ExitGenericInterfaceMethodDeclaration is called when production genericInterfaceMethodDeclaration is exited.
func (s *BaseJavaParserListener) ExitGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) {
}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseJavaParserListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseJavaParserListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseJavaParserListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseJavaParserListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseJavaParserListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseJavaParserListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseJavaParserListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseJavaParserListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseJavaParserListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseJavaParserListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseJavaParserListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseJavaParserListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseJavaParserListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseJavaParserListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterQualifiedNameList is called when production qualifiedNameList is entered.
func (s *BaseJavaParserListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {}

// ExitQualifiedNameList is called when production qualifiedNameList is exited.
func (s *BaseJavaParserListener) ExitQualifiedNameList(ctx *QualifiedNameListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseJavaParserListener) EnterFormalParameters(ctx *FormalParametersContext) {
	//fmt.Println("params:", ctx.GetText())
}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseJavaParserListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseJavaParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseJavaParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseJavaParserListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseJavaParserListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BaseJavaParserListener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BaseJavaParserListener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseJavaParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseJavaParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJavaParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJavaParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseJavaParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseJavaParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseJavaParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseJavaParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterAltAnnotationQualifiedName is called when production altAnnotationQualifiedName is entered.
func (s *BaseJavaParserListener) EnterAltAnnotationQualifiedName(ctx *AltAnnotationQualifiedNameContext) {
}

// ExitAltAnnotationQualifiedName is called when production altAnnotationQualifiedName is exited.
func (s *BaseJavaParserListener) ExitAltAnnotationQualifiedName(ctx *AltAnnotationQualifiedNameContext) {
}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseJavaParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseJavaParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterElementValuePairs is called when production elementValuePairs is entered.
func (s *BaseJavaParserListener) EnterElementValuePairs(ctx *ElementValuePairsContext) {}

// ExitElementValuePairs is called when production elementValuePairs is exited.
func (s *BaseJavaParserListener) ExitElementValuePairs(ctx *ElementValuePairsContext) {}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseJavaParserListener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseJavaParserListener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterElementValue is called when production elementValue is entered.
func (s *BaseJavaParserListener) EnterElementValue(ctx *ElementValueContext) {}

// ExitElementValue is called when production elementValue is exited.
func (s *BaseJavaParserListener) ExitElementValue(ctx *ElementValueContext) {}

// EnterElementValueArrayInitializer is called when production elementValueArrayInitializer is entered.
func (s *BaseJavaParserListener) EnterElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// ExitElementValueArrayInitializer is called when production elementValueArrayInitializer is exited.
func (s *BaseJavaParserListener) ExitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// EnterAnnotationTypeDeclaration is called when production annotationTypeDeclaration is entered.
func (s *BaseJavaParserListener) EnterAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {
}

// ExitAnnotationTypeDeclaration is called when production annotationTypeDeclaration is exited.
func (s *BaseJavaParserListener) ExitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {
}

// EnterAnnotationTypeBody is called when production annotationTypeBody is entered.
func (s *BaseJavaParserListener) EnterAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// ExitAnnotationTypeBody is called when production annotationTypeBody is exited.
func (s *BaseJavaParserListener) ExitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// EnterAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is entered.
func (s *BaseJavaParserListener) EnterAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// ExitAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is exited.
func (s *BaseJavaParserListener) ExitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// EnterAnnotationTypeElementRest is called when production annotationTypeElementRest is entered.
func (s *BaseJavaParserListener) EnterAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) {
}

// ExitAnnotationTypeElementRest is called when production annotationTypeElementRest is exited.
func (s *BaseJavaParserListener) ExitAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) {
}

// EnterAnnotationMethodOrConstantRest is called when production annotationMethodOrConstantRest is entered.
func (s *BaseJavaParserListener) EnterAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) {
}

// ExitAnnotationMethodOrConstantRest is called when production annotationMethodOrConstantRest is exited.
func (s *BaseJavaParserListener) ExitAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) {
}

// EnterAnnotationMethodRest is called when production annotationMethodRest is entered.
func (s *BaseJavaParserListener) EnterAnnotationMethodRest(ctx *AnnotationMethodRestContext) {}

// ExitAnnotationMethodRest is called when production annotationMethodRest is exited.
func (s *BaseJavaParserListener) ExitAnnotationMethodRest(ctx *AnnotationMethodRestContext) {}

// EnterAnnotationConstantRest is called when production annotationConstantRest is entered.
func (s *BaseJavaParserListener) EnterAnnotationConstantRest(ctx *AnnotationConstantRestContext) {}

// ExitAnnotationConstantRest is called when production annotationConstantRest is exited.
func (s *BaseJavaParserListener) ExitAnnotationConstantRest(ctx *AnnotationConstantRestContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseJavaParserListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseJavaParserListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJavaParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJavaParserListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseJavaParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseJavaParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseJavaParserListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseJavaParserListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// EnterLocalTypeDeclaration is called when production localTypeDeclaration is entered.
func (s *BaseJavaParserListener) EnterLocalTypeDeclaration(ctx *LocalTypeDeclarationContext) {}

// ExitLocalTypeDeclaration is called when production localTypeDeclaration is exited.
func (s *BaseJavaParserListener) ExitLocalTypeDeclaration(ctx *LocalTypeDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJavaParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJavaParserListener) ExitStatement(ctx *StatementContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseJavaParserListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseJavaParserListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterCatchType is called when production catchType is entered.
func (s *BaseJavaParserListener) EnterCatchType(ctx *CatchTypeContext) {}

// ExitCatchType is called when production catchType is exited.
func (s *BaseJavaParserListener) ExitCatchType(ctx *CatchTypeContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *BaseJavaParserListener) EnterFinallyBlock(ctx *FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *BaseJavaParserListener) ExitFinallyBlock(ctx *FinallyBlockContext) {}

// EnterResourceSpecification is called when production resourceSpecification is entered.
func (s *BaseJavaParserListener) EnterResourceSpecification(ctx *ResourceSpecificationContext) {}

// ExitResourceSpecification is called when production resourceSpecification is exited.
func (s *BaseJavaParserListener) ExitResourceSpecification(ctx *ResourceSpecificationContext) {}

// EnterResources is called when production resources is entered.
func (s *BaseJavaParserListener) EnterResources(ctx *ResourcesContext) {}

// ExitResources is called when production resources is exited.
func (s *BaseJavaParserListener) ExitResources(ctx *ResourcesContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseJavaParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseJavaParserListener) ExitResource(ctx *ResourceContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseJavaParserListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseJavaParserListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseJavaParserListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseJavaParserListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseJavaParserListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseJavaParserListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseJavaParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseJavaParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterEnhancedForControl is called when production enhancedForControl is entered.
func (s *BaseJavaParserListener) EnterEnhancedForControl(ctx *EnhancedForControlContext) {}

// ExitEnhancedForControl is called when production enhancedForControl is exited.
func (s *BaseJavaParserListener) ExitEnhancedForControl(ctx *EnhancedForControlContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseJavaParserListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseJavaParserListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseJavaParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseJavaParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseJavaParserListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJavaParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJavaParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseJavaParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseJavaParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *BaseJavaParserListener) EnterLambdaParameters(ctx *LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *BaseJavaParserListener) ExitLambdaParameters(ctx *LambdaParametersContext) {}

// EnterLambdaBody is called when production lambdaBody is entered.
func (s *BaseJavaParserListener) EnterLambdaBody(ctx *LambdaBodyContext) {}

// ExitLambdaBody is called when production lambdaBody is exited.
func (s *BaseJavaParserListener) ExitLambdaBody(ctx *LambdaBodyContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseJavaParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseJavaParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterClassType is called when production classType is entered.
func (s *BaseJavaParserListener) EnterClassType(ctx *ClassTypeContext) {}

// ExitClassType is called when production classType is exited.
func (s *BaseJavaParserListener) ExitClassType(ctx *ClassTypeContext) {}

// EnterCreator is called when production creator is entered.
func (s *BaseJavaParserListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BaseJavaParserListener) ExitCreator(ctx *CreatorContext) {}

// EnterCreatedName is called when production createdName is entered.
func (s *BaseJavaParserListener) EnterCreatedName(ctx *CreatedNameContext) {}

// ExitCreatedName is called when production createdName is exited.
func (s *BaseJavaParserListener) ExitCreatedName(ctx *CreatedNameContext) {}

// EnterInnerCreator is called when production innerCreator is entered.
func (s *BaseJavaParserListener) EnterInnerCreator(ctx *InnerCreatorContext) {}

// ExitInnerCreator is called when production innerCreator is exited.
func (s *BaseJavaParserListener) ExitInnerCreator(ctx *InnerCreatorContext) {}

// EnterArrayCreatorRest is called when production arrayCreatorRest is entered.
func (s *BaseJavaParserListener) EnterArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// ExitArrayCreatorRest is called when production arrayCreatorRest is exited.
func (s *BaseJavaParserListener) ExitArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// EnterClassCreatorRest is called when production classCreatorRest is entered.
func (s *BaseJavaParserListener) EnterClassCreatorRest(ctx *ClassCreatorRestContext) {}

// ExitClassCreatorRest is called when production classCreatorRest is exited.
func (s *BaseJavaParserListener) ExitClassCreatorRest(ctx *ClassCreatorRestContext) {}

// EnterExplicitGenericInvocation is called when production explicitGenericInvocation is entered.
func (s *BaseJavaParserListener) EnterExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) {
}

// ExitExplicitGenericInvocation is called when production explicitGenericInvocation is exited.
func (s *BaseJavaParserListener) ExitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) {
}

// EnterTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is entered.
func (s *BaseJavaParserListener) EnterTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// ExitTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is exited.
func (s *BaseJavaParserListener) ExitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// EnterNonWildcardTypeArgumentsOrDiamond is called when production nonWildcardTypeArgumentsOrDiamond is entered.
func (s *BaseJavaParserListener) EnterNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) {
}

// ExitNonWildcardTypeArgumentsOrDiamond is called when production nonWildcardTypeArgumentsOrDiamond is exited.
func (s *BaseJavaParserListener) ExitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) {
}

// EnterNonWildcardTypeArguments is called when production nonWildcardTypeArguments is entered.
func (s *BaseJavaParserListener) EnterNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {
}

// ExitNonWildcardTypeArguments is called when production nonWildcardTypeArguments is exited.
func (s *BaseJavaParserListener) ExitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseJavaParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseJavaParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BaseJavaParserListener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BaseJavaParserListener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseJavaParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseJavaParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseJavaParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseJavaParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterSuperSuffix is called when production superSuffix is entered.
func (s *BaseJavaParserListener) EnterSuperSuffix(ctx *SuperSuffixContext) {}

// ExitSuperSuffix is called when production superSuffix is exited.
func (s *BaseJavaParserListener) ExitSuperSuffix(ctx *SuperSuffixContext) {}

// EnterExplicitGenericInvocationSuffix is called when production explicitGenericInvocationSuffix is entered.
func (s *BaseJavaParserListener) EnterExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) {
}

// ExitExplicitGenericInvocationSuffix is called when production explicitGenericInvocationSuffix is exited.
func (s *BaseJavaParserListener) ExitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) {
}

// EnterArguments is called when production arguments is entered.
func (s *BaseJavaParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseJavaParserListener) ExitArguments(ctx *ArgumentsContext) {}
