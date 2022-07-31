// Code generated from D:/Desktop/Whosbug_antlr_go/antlr4_kotlin/ast_kotlin\KotlinParser.g4 by ANTLR 4.9.1. DO NOT EDIT.
package kotlinLib // KotlinParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// KotlinParserListener is a complete listener for a parse tree produced by KotlinParser.
type KotlinParserListener interface {
	antlr.ParseTreeListener

	// EnterKotlinFile is called when entering the kotlinFile production.
	EnterKotlinFile(c *KotlinFileContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterPreamble is called when entering the preamble production.
	EnterPreamble(c *PreambleContext)

	// EnterFileAnnotations is called when entering the fileAnnotations production.
	EnterFileAnnotations(c *FileAnnotationsContext)

	// EnterFileAnnotation is called when entering the fileAnnotation production.
	EnterFileAnnotation(c *FileAnnotationContext)

	// EnterPackageHeader is called when entering the packageHeader production.
	EnterPackageHeader(c *PackageHeaderContext)

	// EnterImportList is called when entering the importList production.
	EnterImportList(c *ImportListContext)

	// EnterImportHeader is called when entering the importHeader production.
	EnterImportHeader(c *ImportHeaderContext)

	// EnterImportAlias is called when entering the importAlias production.
	EnterImportAlias(c *ImportAliasContext)

	// EnterTopLevelObject is called when entering the topLevelObject production.
	EnterTopLevelObject(c *TopLevelObjectContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterPrimaryConstructor is called when entering the primaryConstructor production.
	EnterPrimaryConstructor(c *PrimaryConstructorContext)

	// EnterClassParameters is called when entering the classParameters production.
	EnterClassParameters(c *ClassParametersContext)

	// EnterClassParameter is called when entering the classParameter production.
	EnterClassParameter(c *ClassParameterContext)

	// EnterDelegationSpecifiers is called when entering the delegationSpecifiers production.
	EnterDelegationSpecifiers(c *DelegationSpecifiersContext)

	// EnterDelegationSpecifier is called when entering the delegationSpecifier production.
	EnterDelegationSpecifier(c *DelegationSpecifierContext)

	// EnterConstructorInvocation is called when entering the constructorInvocation production.
	EnterConstructorInvocation(c *ConstructorInvocationContext)

	// EnterExplicitDelegation is called when entering the explicitDelegation production.
	EnterExplicitDelegation(c *ExplicitDelegationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassMemberDeclaration is called when entering the classMemberDeclaration production.
	EnterClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// EnterAnonymousInitializer is called when entering the anonymousInitializer production.
	EnterAnonymousInitializer(c *AnonymousInitializerContext)

	// EnterSecondaryConstructor is called when entering the secondaryConstructor production.
	EnterSecondaryConstructor(c *SecondaryConstructorContext)

	// EnterConstructorDelegationCall is called when entering the constructorDelegationCall production.
	EnterConstructorDelegationCall(c *ConstructorDelegationCallContext)

	// EnterEnumClassBody is called when entering the enumClassBody production.
	EnterEnumClassBody(c *EnumClassBodyContext)

	// EnterEnumEntries is called when entering the enumEntries production.
	EnterEnumEntries(c *EnumEntriesContext)

	// EnterEnumEntry is called when entering the enumEntry production.
	EnterEnumEntry(c *EnumEntryContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionValueParameters is called when entering the functionValueParameters production.
	EnterFunctionValueParameters(c *FunctionValueParametersContext)

	// EnterFunctionValueParameter is called when entering the functionValueParameter production.
	EnterFunctionValueParameter(c *FunctionValueParameterContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterObjectDeclaration is called when entering the objectDeclaration production.
	EnterObjectDeclaration(c *ObjectDeclarationContext)

	// EnterCompanionObject is called when entering the companionObject production.
	EnterCompanionObject(c *CompanionObjectContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterMultiVariableDeclaration is called when entering the multiVariableDeclaration production.
	EnterMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterGetter is called when entering the getter production.
	EnterGetter(c *GetterContext)

	// EnterSetter is called when entering the setter production.
	EnterSetter(c *SetterContext)

	// EnterTypeAlias is called when entering the typeAlias production.
	EnterTypeAlias(c *TypeAliasContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeModifierList is called when entering the typeModifierList production.
	EnterTypeModifierList(c *TypeModifierListContext)

	// EnterParenthesizedType is called when entering the parenthesizedType production.
	EnterParenthesizedType(c *ParenthesizedTypeContext)

	// EnterNullableType is called when entering the nullableType production.
	EnterNullableType(c *NullableTypeContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterFunctionTypeReceiver is called when entering the functionTypeReceiver production.
	EnterFunctionTypeReceiver(c *FunctionTypeReceiverContext)

	// EnterUserType is called when entering the userType production.
	EnterUserType(c *UserTypeContext)

	// EnterSimpleUserType is called when entering the simpleUserType production.
	EnterSimpleUserType(c *SimpleUserTypeContext)

	// EnterFunctionTypeParameters is called when entering the functionTypeParameters production.
	EnterFunctionTypeParameters(c *FunctionTypeParametersContext)

	// EnterTypeConstraints is called when entering the typeConstraints production.
	EnterTypeConstraints(c *TypeConstraintsContext)

	// EnterTypeConstraint is called when entering the typeConstraint production.
	EnterTypeConstraint(c *TypeConstraintContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlockLevelExpression is called when entering the blockLevelExpression production.
	EnterBlockLevelExpression(c *BlockLevelExpressionContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDisjunction is called when entering the disjunction production.
	EnterDisjunction(c *DisjunctionContext)

	// EnterConjunction is called when entering the conjunction production.
	EnterConjunction(c *ConjunctionContext)

	// EnterEqualityComparison is called when entering the equalityComparison production.
	EnterEqualityComparison(c *EqualityComparisonContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterNamedInfix is called when entering the namedInfix production.
	EnterNamedInfix(c *NamedInfixContext)

	// EnterElvisExpression is called when entering the elvisExpression production.
	EnterElvisExpression(c *ElvisExpressionContext)

	// EnterInfixFunctionCall is called when entering the infixFunctionCall production.
	EnterInfixFunctionCall(c *InfixFunctionCallContext)

	// EnterRangeExpression is called when entering the rangeExpression production.
	EnterRangeExpression(c *RangeExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterTypeRHS is called when entering the typeRHS production.
	EnterTypeRHS(c *TypeRHSContext)

	// EnterPrefixUnaryExpression is called when entering the prefixUnaryExpression production.
	EnterPrefixUnaryExpression(c *PrefixUnaryExpressionContext)

	// EnterPostfixUnaryExpression is called when entering the postfixUnaryExpression production.
	EnterPostfixUnaryExpression(c *PostfixUnaryExpressionContext)

	// EnterAtomicExpression is called when entering the atomicExpression production.
	EnterAtomicExpression(c *AtomicExpressionContext)

	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterCallSuffix is called when entering the callSuffix production.
	EnterCallSuffix(c *CallSuffixContext)

	// EnterAnnotatedLambda is called when entering the annotatedLambda production.
	EnterAnnotatedLambda(c *AnnotatedLambdaContext)

	// EnterArrayAccess is called when entering the arrayAccess production.
	EnterArrayAccess(c *ArrayAccessContext)

	// EnterValueArguments is called when entering the valueArguments production.
	EnterValueArguments(c *ValueArgumentsContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeProjection is called when entering the typeProjection production.
	EnterTypeProjection(c *TypeProjectionContext)

	// EnterTypeProjectionModifierList is called when entering the typeProjectionModifierList production.
	EnterTypeProjectionModifierList(c *TypeProjectionModifierListContext)

	// EnterValueArgument is called when entering the valueArgument production.
	EnterValueArgument(c *ValueArgumentContext)

	// EnterLiteralConstant is called when entering the literalConstant production.
	EnterLiteralConstant(c *LiteralConstantContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterLineStringLiteral is called when entering the lineStringLiteral production.
	EnterLineStringLiteral(c *LineStringLiteralContext)

	// EnterMultiLineStringLiteral is called when entering the multiLineStringLiteral production.
	EnterMultiLineStringLiteral(c *MultiLineStringLiteralContext)

	// EnterLineStringContent is called when entering the lineStringContent production.
	EnterLineStringContent(c *LineStringContentContext)

	// EnterLineStringExpression is called when entering the lineStringExpression production.
	EnterLineStringExpression(c *LineStringExpressionContext)

	// EnterMultiLineStringContent is called when entering the multiLineStringContent production.
	EnterMultiLineStringContent(c *MultiLineStringContentContext)

	// EnterMultiLineStringExpression is called when entering the multiLineStringExpression production.
	EnterMultiLineStringExpression(c *MultiLineStringExpressionContext)

	// EnterFunctionLiteral is called when entering the functionLiteral production.
	EnterFunctionLiteral(c *FunctionLiteralContext)

	// EnterLambdaParameters is called when entering the lambdaParameters production.
	EnterLambdaParameters(c *LambdaParametersContext)

	// EnterLambdaParameter is called when entering the lambdaParameter production.
	EnterLambdaParameter(c *LambdaParameterContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterCollectionLiteral is called when entering the collectionLiteral production.
	EnterCollectionLiteral(c *CollectionLiteralContext)

	// EnterThisExpression is called when entering the thisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterSuperExpression is called when entering the superExpression production.
	EnterSuperExpression(c *SuperExpressionContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterIfExpression is called when entering the ifExpression production.
	EnterIfExpression(c *IfExpressionContext)

	// EnterControlStructureBody is called when entering the controlStructureBody production.
	EnterControlStructureBody(c *ControlStructureBodyContext)

	// EnterWhenExpression is called when entering the whenExpression production.
	EnterWhenExpression(c *WhenExpressionContext)

	// EnterWhenEntry is called when entering the whenEntry production.
	EnterWhenEntry(c *WhenEntryContext)

	// EnterWhenCondition is called when entering the whenCondition production.
	EnterWhenCondition(c *WhenConditionContext)

	// EnterRangeTest is called when entering the rangeTest production.
	EnterRangeTest(c *RangeTestContext)

	// EnterTypeTest is called when entering the typeTest production.
	EnterTypeTest(c *TypeTestContext)

	// EnterTryExpression is called when entering the tryExpression production.
	EnterTryExpression(c *TryExpressionContext)

	// EnterCatchBlock is called when entering the catchBlock production.
	EnterCatchBlock(c *CatchBlockContext)

	// EnterFinallyBlock is called when entering the finallyBlock production.
	EnterFinallyBlock(c *FinallyBlockContext)

	// EnterLoopExpression is called when entering the loopExpression production.
	EnterLoopExpression(c *LoopExpressionContext)

	// EnterForExpression is called when entering the forExpression production.
	EnterForExpression(c *ForExpressionContext)

	// EnterWhileExpression is called when entering the whileExpression production.
	EnterWhileExpression(c *WhileExpressionContext)

	// EnterDoWhileExpression is called when entering the doWhileExpression production.
	EnterDoWhileExpression(c *DoWhileExpressionContext)

	// EnterJumpExpression is called when entering the jumpExpression production.
	EnterJumpExpression(c *JumpExpressionContext)

	// EnterCallableReference is called when entering the callableReference production.
	EnterCallableReference(c *CallableReferenceContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterEqualityOperation is called when entering the equalityOperation production.
	EnterEqualityOperation(c *EqualityOperationContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterInOperator is called when entering the inOperator production.
	EnterInOperator(c *InOperatorContext)

	// EnterIsOperator is called when entering the isOperator production.
	EnterIsOperator(c *IsOperatorContext)

	// EnterAdditiveOperator is called when entering the additiveOperator production.
	EnterAdditiveOperator(c *AdditiveOperatorContext)

	// EnterMultiplicativeOperation is called when entering the multiplicativeOperation production.
	EnterMultiplicativeOperation(c *MultiplicativeOperationContext)

	// EnterTypeOperation is called when entering the typeOperation production.
	EnterTypeOperation(c *TypeOperationContext)

	// EnterPrefixUnaryOperation is called when entering the prefixUnaryOperation production.
	EnterPrefixUnaryOperation(c *PrefixUnaryOperationContext)

	// EnterPostfixUnaryOperation is called when entering the postfixUnaryOperation production.
	EnterPostfixUnaryOperation(c *PostfixUnaryOperationContext)

	// EnterMemberAccessOperator is called when entering the memberAccessOperator production.
	EnterMemberAccessOperator(c *MemberAccessOperatorContext)

	// EnterModifierList is called when entering the modifierList production.
	EnterModifierList(c *ModifierListContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterClassModifier is called when entering the classModifier production.
	EnterClassModifier(c *ClassModifierContext)

	// EnterMemberModifier is called when entering the memberModifier production.
	EnterMemberModifier(c *MemberModifierContext)

	// EnterVisibilityModifier is called when entering the visibilityModifier production.
	EnterVisibilityModifier(c *VisibilityModifierContext)

	// EnterVarianceAnnotation is called when entering the varianceAnnotation production.
	EnterVarianceAnnotation(c *VarianceAnnotationContext)

	// EnterFunctionModifier is called when entering the functionModifier production.
	EnterFunctionModifier(c *FunctionModifierContext)

	// EnterPropertyModifier is called when entering the propertyModifier production.
	EnterPropertyModifier(c *PropertyModifierContext)

	// EnterInheritanceModifier is called when entering the inheritanceModifier production.
	EnterInheritanceModifier(c *InheritanceModifierContext)

	// EnterParameterModifier is called when entering the parameterModifier production.
	EnterParameterModifier(c *ParameterModifierContext)

	// EnterTypeParameterModifier is called when entering the typeParameterModifier production.
	EnterTypeParameterModifier(c *TypeParameterModifierContext)

	// EnterLabelDefinition is called when entering the labelDefinition production.
	EnterLabelDefinition(c *LabelDefinitionContext)

	// EnterAnnotations is called when entering the annotations production.
	EnterAnnotations(c *AnnotationsContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationList is called when entering the annotationList production.
	EnterAnnotationList(c *AnnotationListContext)

	// EnterAnnotationUseSiteTarget is called when entering the annotationUseSiteTarget production.
	EnterAnnotationUseSiteTarget(c *AnnotationUseSiteTargetContext)

	// EnterUnescapedAnnotation is called when entering the unescapedAnnotation production.
	EnterUnescapedAnnotation(c *UnescapedAnnotationContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterSimpleIdentifier is called when entering the simpleIdentifier production.
	EnterSimpleIdentifier(c *SimpleIdentifierContext)

	// EnterSemi is called when entering the semi production.
	EnterSemi(c *SemiContext)

	// EnterAnysemi is called when entering the anysemi production.
	EnterAnysemi(c *AnysemiContext)

	// ExitKotlinFile is called when exiting the kotlinFile production.
	ExitKotlinFile(c *KotlinFileContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitPreamble is called when exiting the preamble production.
	ExitPreamble(c *PreambleContext)

	// ExitFileAnnotations is called when exiting the fileAnnotations production.
	ExitFileAnnotations(c *FileAnnotationsContext)

	// ExitFileAnnotation is called when exiting the fileAnnotation production.
	ExitFileAnnotation(c *FileAnnotationContext)

	// ExitPackageHeader is called when exiting the packageHeader production.
	ExitPackageHeader(c *PackageHeaderContext)

	// ExitImportList is called when exiting the importList production.
	ExitImportList(c *ImportListContext)

	// ExitImportHeader is called when exiting the importHeader production.
	ExitImportHeader(c *ImportHeaderContext)

	// ExitImportAlias is called when exiting the importAlias production.
	ExitImportAlias(c *ImportAliasContext)

	// ExitTopLevelObject is called when exiting the topLevelObject production.
	ExitTopLevelObject(c *TopLevelObjectContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitPrimaryConstructor is called when exiting the primaryConstructor production.
	ExitPrimaryConstructor(c *PrimaryConstructorContext)

	// ExitClassParameters is called when exiting the classParameters production.
	ExitClassParameters(c *ClassParametersContext)

	// ExitClassParameter is called when exiting the classParameter production.
	ExitClassParameter(c *ClassParameterContext)

	// ExitDelegationSpecifiers is called when exiting the delegationSpecifiers production.
	ExitDelegationSpecifiers(c *DelegationSpecifiersContext)

	// ExitDelegationSpecifier is called when exiting the delegationSpecifier production.
	ExitDelegationSpecifier(c *DelegationSpecifierContext)

	// ExitConstructorInvocation is called when exiting the constructorInvocation production.
	ExitConstructorInvocation(c *ConstructorInvocationContext)

	// ExitExplicitDelegation is called when exiting the explicitDelegation production.
	ExitExplicitDelegation(c *ExplicitDelegationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassMemberDeclaration is called when exiting the classMemberDeclaration production.
	ExitClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// ExitAnonymousInitializer is called when exiting the anonymousInitializer production.
	ExitAnonymousInitializer(c *AnonymousInitializerContext)

	// ExitSecondaryConstructor is called when exiting the secondaryConstructor production.
	ExitSecondaryConstructor(c *SecondaryConstructorContext)

	// ExitConstructorDelegationCall is called when exiting the constructorDelegationCall production.
	ExitConstructorDelegationCall(c *ConstructorDelegationCallContext)

	// ExitEnumClassBody is called when exiting the enumClassBody production.
	ExitEnumClassBody(c *EnumClassBodyContext)

	// ExitEnumEntries is called when exiting the enumEntries production.
	ExitEnumEntries(c *EnumEntriesContext)

	// ExitEnumEntry is called when exiting the enumEntry production.
	ExitEnumEntry(c *EnumEntryContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionValueParameters is called when exiting the functionValueParameters production.
	ExitFunctionValueParameters(c *FunctionValueParametersContext)

	// ExitFunctionValueParameter is called when exiting the functionValueParameter production.
	ExitFunctionValueParameter(c *FunctionValueParameterContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitObjectDeclaration is called when exiting the objectDeclaration production.
	ExitObjectDeclaration(c *ObjectDeclarationContext)

	// ExitCompanionObject is called when exiting the companionObject production.
	ExitCompanionObject(c *CompanionObjectContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitMultiVariableDeclaration is called when exiting the multiVariableDeclaration production.
	ExitMultiVariableDeclaration(c *MultiVariableDeclarationContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitGetter is called when exiting the getter production.
	ExitGetter(c *GetterContext)

	// ExitSetter is called when exiting the setter production.
	ExitSetter(c *SetterContext)

	// ExitTypeAlias is called when exiting the typeAlias production.
	ExitTypeAlias(c *TypeAliasContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeModifierList is called when exiting the typeModifierList production.
	ExitTypeModifierList(c *TypeModifierListContext)

	// ExitParenthesizedType is called when exiting the parenthesizedType production.
	ExitParenthesizedType(c *ParenthesizedTypeContext)

	// ExitNullableType is called when exiting the nullableType production.
	ExitNullableType(c *NullableTypeContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitFunctionTypeReceiver is called when exiting the functionTypeReceiver production.
	ExitFunctionTypeReceiver(c *FunctionTypeReceiverContext)

	// ExitUserType is called when exiting the userType production.
	ExitUserType(c *UserTypeContext)

	// ExitSimpleUserType is called when exiting the simpleUserType production.
	ExitSimpleUserType(c *SimpleUserTypeContext)

	// ExitFunctionTypeParameters is called when exiting the functionTypeParameters production.
	ExitFunctionTypeParameters(c *FunctionTypeParametersContext)

	// ExitTypeConstraints is called when exiting the typeConstraints production.
	ExitTypeConstraints(c *TypeConstraintsContext)

	// ExitTypeConstraint is called when exiting the typeConstraint production.
	ExitTypeConstraint(c *TypeConstraintContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlockLevelExpression is called when exiting the blockLevelExpression production.
	ExitBlockLevelExpression(c *BlockLevelExpressionContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDisjunction is called when exiting the disjunction production.
	ExitDisjunction(c *DisjunctionContext)

	// ExitConjunction is called when exiting the conjunction production.
	ExitConjunction(c *ConjunctionContext)

	// ExitEqualityComparison is called when exiting the equalityComparison production.
	ExitEqualityComparison(c *EqualityComparisonContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitNamedInfix is called when exiting the namedInfix production.
	ExitNamedInfix(c *NamedInfixContext)

	// ExitElvisExpression is called when exiting the elvisExpression production.
	ExitElvisExpression(c *ElvisExpressionContext)

	// ExitInfixFunctionCall is called when exiting the infixFunctionCall production.
	ExitInfixFunctionCall(c *InfixFunctionCallContext)

	// ExitRangeExpression is called when exiting the rangeExpression production.
	ExitRangeExpression(c *RangeExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitTypeRHS is called when exiting the typeRHS production.
	ExitTypeRHS(c *TypeRHSContext)

	// ExitPrefixUnaryExpression is called when exiting the prefixUnaryExpression production.
	ExitPrefixUnaryExpression(c *PrefixUnaryExpressionContext)

	// ExitPostfixUnaryExpression is called when exiting the postfixUnaryExpression production.
	ExitPostfixUnaryExpression(c *PostfixUnaryExpressionContext)

	// ExitAtomicExpression is called when exiting the atomicExpression production.
	ExitAtomicExpression(c *AtomicExpressionContext)

	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitCallSuffix is called when exiting the callSuffix production.
	ExitCallSuffix(c *CallSuffixContext)

	// ExitAnnotatedLambda is called when exiting the annotatedLambda production.
	ExitAnnotatedLambda(c *AnnotatedLambdaContext)

	// ExitArrayAccess is called when exiting the arrayAccess production.
	ExitArrayAccess(c *ArrayAccessContext)

	// ExitValueArguments is called when exiting the valueArguments production.
	ExitValueArguments(c *ValueArgumentsContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeProjection is called when exiting the typeProjection production.
	ExitTypeProjection(c *TypeProjectionContext)

	// ExitTypeProjectionModifierList is called when exiting the typeProjectionModifierList production.
	ExitTypeProjectionModifierList(c *TypeProjectionModifierListContext)

	// ExitValueArgument is called when exiting the valueArgument production.
	ExitValueArgument(c *ValueArgumentContext)

	// ExitLiteralConstant is called when exiting the literalConstant production.
	ExitLiteralConstant(c *LiteralConstantContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitLineStringLiteral is called when exiting the lineStringLiteral production.
	ExitLineStringLiteral(c *LineStringLiteralContext)

	// ExitMultiLineStringLiteral is called when exiting the multiLineStringLiteral production.
	ExitMultiLineStringLiteral(c *MultiLineStringLiteralContext)

	// ExitLineStringContent is called when exiting the lineStringContent production.
	ExitLineStringContent(c *LineStringContentContext)

	// ExitLineStringExpression is called when exiting the lineStringExpression production.
	ExitLineStringExpression(c *LineStringExpressionContext)

	// ExitMultiLineStringContent is called when exiting the multiLineStringContent production.
	ExitMultiLineStringContent(c *MultiLineStringContentContext)

	// ExitMultiLineStringExpression is called when exiting the multiLineStringExpression production.
	ExitMultiLineStringExpression(c *MultiLineStringExpressionContext)

	// ExitFunctionLiteral is called when exiting the functionLiteral production.
	ExitFunctionLiteral(c *FunctionLiteralContext)

	// ExitLambdaParameters is called when exiting the lambdaParameters production.
	ExitLambdaParameters(c *LambdaParametersContext)

	// ExitLambdaParameter is called when exiting the lambdaParameter production.
	ExitLambdaParameter(c *LambdaParameterContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitCollectionLiteral is called when exiting the collectionLiteral production.
	ExitCollectionLiteral(c *CollectionLiteralContext)

	// ExitThisExpression is called when exiting the thisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitSuperExpression is called when exiting the superExpression production.
	ExitSuperExpression(c *SuperExpressionContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitIfExpression is called when exiting the ifExpression production.
	ExitIfExpression(c *IfExpressionContext)

	// ExitControlStructureBody is called when exiting the controlStructureBody production.
	ExitControlStructureBody(c *ControlStructureBodyContext)

	// ExitWhenExpression is called when exiting the whenExpression production.
	ExitWhenExpression(c *WhenExpressionContext)

	// ExitWhenEntry is called when exiting the whenEntry production.
	ExitWhenEntry(c *WhenEntryContext)

	// ExitWhenCondition is called when exiting the whenCondition production.
	ExitWhenCondition(c *WhenConditionContext)

	// ExitRangeTest is called when exiting the rangeTest production.
	ExitRangeTest(c *RangeTestContext)

	// ExitTypeTest is called when exiting the typeTest production.
	ExitTypeTest(c *TypeTestContext)

	// ExitTryExpression is called when exiting the tryExpression production.
	ExitTryExpression(c *TryExpressionContext)

	// ExitCatchBlock is called when exiting the catchBlock production.
	ExitCatchBlock(c *CatchBlockContext)

	// ExitFinallyBlock is called when exiting the finallyBlock production.
	ExitFinallyBlock(c *FinallyBlockContext)

	// ExitLoopExpression is called when exiting the loopExpression production.
	ExitLoopExpression(c *LoopExpressionContext)

	// ExitForExpression is called when exiting the forExpression production.
	ExitForExpression(c *ForExpressionContext)

	// ExitWhileExpression is called when exiting the whileExpression production.
	ExitWhileExpression(c *WhileExpressionContext)

	// ExitDoWhileExpression is called when exiting the doWhileExpression production.
	ExitDoWhileExpression(c *DoWhileExpressionContext)

	// ExitJumpExpression is called when exiting the jumpExpression production.
	ExitJumpExpression(c *JumpExpressionContext)

	// ExitCallableReference is called when exiting the callableReference production.
	ExitCallableReference(c *CallableReferenceContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitEqualityOperation is called when exiting the equalityOperation production.
	ExitEqualityOperation(c *EqualityOperationContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitInOperator is called when exiting the inOperator production.
	ExitInOperator(c *InOperatorContext)

	// ExitIsOperator is called when exiting the isOperator production.
	ExitIsOperator(c *IsOperatorContext)

	// ExitAdditiveOperator is called when exiting the additiveOperator production.
	ExitAdditiveOperator(c *AdditiveOperatorContext)

	// ExitMultiplicativeOperation is called when exiting the multiplicativeOperation production.
	ExitMultiplicativeOperation(c *MultiplicativeOperationContext)

	// ExitTypeOperation is called when exiting the typeOperation production.
	ExitTypeOperation(c *TypeOperationContext)

	// ExitPrefixUnaryOperation is called when exiting the prefixUnaryOperation production.
	ExitPrefixUnaryOperation(c *PrefixUnaryOperationContext)

	// ExitPostfixUnaryOperation is called when exiting the postfixUnaryOperation production.
	ExitPostfixUnaryOperation(c *PostfixUnaryOperationContext)

	// ExitMemberAccessOperator is called when exiting the memberAccessOperator production.
	ExitMemberAccessOperator(c *MemberAccessOperatorContext)

	// ExitModifierList is called when exiting the modifierList production.
	ExitModifierList(c *ModifierListContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitClassModifier is called when exiting the classModifier production.
	ExitClassModifier(c *ClassModifierContext)

	// ExitMemberModifier is called when exiting the memberModifier production.
	ExitMemberModifier(c *MemberModifierContext)

	// ExitVisibilityModifier is called when exiting the visibilityModifier production.
	ExitVisibilityModifier(c *VisibilityModifierContext)

	// ExitVarianceAnnotation is called when exiting the varianceAnnotation production.
	ExitVarianceAnnotation(c *VarianceAnnotationContext)

	// ExitFunctionModifier is called when exiting the functionModifier production.
	ExitFunctionModifier(c *FunctionModifierContext)

	// ExitPropertyModifier is called when exiting the propertyModifier production.
	ExitPropertyModifier(c *PropertyModifierContext)

	// ExitInheritanceModifier is called when exiting the inheritanceModifier production.
	ExitInheritanceModifier(c *InheritanceModifierContext)

	// ExitParameterModifier is called when exiting the parameterModifier production.
	ExitParameterModifier(c *ParameterModifierContext)

	// ExitTypeParameterModifier is called when exiting the typeParameterModifier production.
	ExitTypeParameterModifier(c *TypeParameterModifierContext)

	// ExitLabelDefinition is called when exiting the labelDefinition production.
	ExitLabelDefinition(c *LabelDefinitionContext)

	// ExitAnnotations is called when exiting the annotations production.
	ExitAnnotations(c *AnnotationsContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationList is called when exiting the annotationList production.
	ExitAnnotationList(c *AnnotationListContext)

	// ExitAnnotationUseSiteTarget is called when exiting the annotationUseSiteTarget production.
	ExitAnnotationUseSiteTarget(c *AnnotationUseSiteTargetContext)

	// ExitUnescapedAnnotation is called when exiting the unescapedAnnotation production.
	ExitUnescapedAnnotation(c *UnescapedAnnotationContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitSimpleIdentifier is called when exiting the simpleIdentifier production.
	ExitSimpleIdentifier(c *SimpleIdentifierContext)

	// ExitSemi is called when exiting the semi production.
	ExitSemi(c *SemiContext)

	// ExitAnysemi is called when exiting the anysemi production.
	ExitAnysemi(c *AnysemiContext)
}
