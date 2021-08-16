//// Generated from KotlinParser.g4 by ANTLR 4.7.
//
package parser // KotlinParser
//
//import "github.com/antlr/antlr4/runtime/Go/antlr"
//
//// KotlinParserListener is a complete listener for a parse tree produced by KotlinParser.
//type KotlinParserListener_ interface {
//	antlr.ParseTreeListener
//
//	// EnterKotlinFile is called when entering the kotlinFile production.
//	EnterKotlinFile(c *KotlinFileContext)
//
//	// EnterScript is called when entering the script production.
//	EnterScript(c *ScriptContext)
//
//	// EnterFileAnnotation is called when entering the fileAnnotation production.
//	EnterFileAnnotation(c *FileAnnotationContext)
//
//	// EnterPackageHeader is called when entering the packageHeader production.
//	EnterPackageHeader(c *PackageHeaderContext)
//
//	// EnterImportList is called when entering the importList production.
//	EnterImportList(c *ImportListContext)
//
//	// EnterImportHeader is called when entering the importHeader production.
//	EnterImportHeader(c *ImportHeaderContext)
//
//	// EnterImportAlias is called when entering the importAlias production.
//	EnterImportAlias(c *ImportAliasContext)
//
//	// EnterTopLevelObject is called when entering the topLevelObject production.
//	EnterTopLevelObject(c *TopLevelObjectContext)
//
//	// EnterClassDeclaration is called when entering the classDeclaration production.
//	EnterClassDeclaration(c *ClassDeclarationContext)
//
//	// EnterPrimaryConstructor is called when entering the primaryConstructor production.
//	EnterPrimaryConstructor(c *PrimaryConstructorContext)
//
//	// EnterClassParameters is called when entering the classParameters production.
//	EnterClassParameters(c *ClassParametersContext)
//
//	// EnterClassParameter is called when entering the classParameter production.
//	EnterClassParameter(c *ClassParameterContext)
//
//	// EnterDelegationSpecifiers is called when entering the delegationSpecifiers production.
//	EnterDelegationSpecifiers(c *DelegationSpecifiersContext)
//
//	// EnterAnnotatedDelegationSpecifier is called when entering the annotatedDelegationSpecifier production.
//	EnterAnnotatedDelegationSpecifier(c *AnnotatedDelegationSpecifierContext)
//
//	// EnterDelegationSpecifier is called when entering the delegationSpecifier production.
//	EnterDelegationSpecifier(c *DelegationSpecifierContext)
//
//	// EnterConstructorInvocation is called when entering the constructorInvocation production.
//	EnterConstructorInvocation(c *ConstructorInvocationContext)
//
//	// EnterExplicitDelegation is called when entering the explicitDelegation production.
//	EnterExplicitDelegation(c *ExplicitDelegationContext)
//
//	// EnterClassBody is called when entering the classBody production.
//	EnterClassBody(c *ClassBodyContext)
//
//	// EnterClassMemberDeclarations is called when entering the classMemberDeclarations production.
//	EnterClassMemberDeclarations(c *ClassMemberDeclarationsContext)
//
//	// EnterClassMemberDeclaration is called when entering the classMemberDeclaration production.
//	EnterClassMemberDeclaration(c *ClassMemberDeclarationContext)
//
//	// EnterAnonymousInitializer is called when entering the anonymousInitializer production.
//	EnterAnonymousInitializer(c *AnonymousInitializerContext)
//
//	// EnterSecondaryConstructor is called when entering the secondaryConstructor production.
//	EnterSecondaryConstructor(c *SecondaryConstructorContext)
//
//	// EnterConstructorDelegationCall is called when entering the constructorDelegationCall production.
//	EnterConstructorDelegationCall(c *ConstructorDelegationCallContext)
//
//	// EnterEnumClassBody is called when entering the enumClassBody production.
//	EnterEnumClassBody(c *EnumClassBodyContext)
//
//	// EnterEnumEntries is called when entering the enumEntries production.
//	EnterEnumEntries(c *EnumEntriesContext)
//
//	// EnterEnumEntry is called when entering the enumEntry production.
//	EnterEnumEntry(c *EnumEntryContext)
//
//	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
//	EnterFunctionDeclaration(c *FunctionDeclarationContext)
//
//	// EnterFunctionValueParameters is called when entering the functionValueParameters production.
//	EnterFunctionValueParameters(c *FunctionValueParametersContext)
//
//	// EnterFunctionValueParameter is called when entering the functionValueParameter production.
//	EnterFunctionValueParameter(c *FunctionValueParameterContext)
//
//	// EnterParameter is called when entering the parameter production.
//	EnterParameter(c *ParameterContext)
//
//	// EnterSetterParameter is called when entering the setterParameter production.
//	EnterSetterParameter(c *SetterParameterContext)
//
//	// EnterFunctionBody is called when entering the functionBody production.
//	EnterFunctionBody(c *FunctionBodyContext)
//
//	// EnterObjectDeclaration is called when entering the objectDeclaration production.
//	EnterObjectDeclaration(c *ObjectDeclarationContext)
//
//	// EnterCompanionObject is called when entering the companionObject production.
//	EnterCompanionObject(c *CompanionObjectContext)
//
//	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
//	EnterPropertyDeclaration(c *PropertyDeclarationContext)
//
//	// EnterMultiVariableDeclaration is called when entering the multiVariableDeclaration production.
//	EnterMultiVariableDeclaration(c *MultiVariableDeclarationContext)
//
//	// EnterVariableDeclaration is called when entering the variableDeclaration production.
//	EnterVariableDeclaration(c *VariableDeclarationContext)
//
//	// EnterPropertyDelegate is called when entering the propertyDelegate production.
//	EnterPropertyDelegate(c *PropertyDelegateContext)
//
//	// EnterGetter is called when entering the getter production.
//	EnterGetter(c *GetterContext)
//
//	// EnterSetter is called when entering the setter production.
//	EnterSetter(c *SetterContext)
//
//	// EnterTypeAlias is called when entering the typeAlias production.
//	EnterTypeAlias(c *TypeAliasContext)
//
//	// EnterTypeParameters is called when entering the typeParameters production.
//	EnterTypeParameters(c *TypeParametersContext)
//
//	// EnterTypeParameter is called when entering the typeParameter production.
//	EnterTypeParameter(c *TypeParameterContext)
//
//	// EnterTypeParameterModifiers is called when entering the typeParameterModifiers production.
//	EnterTypeParameterModifiers(c *TypeParameterModifiersContext)
//
//	// EnterTypeParameterModifier is called when entering the typeParameterModifier production.
//	EnterTypeParameterModifier(c *TypeParameterModifierContext)
//
//	// EnterType_ is called when entering the type_ production.
//	EnterType_(c *Type_Context)
//
//	// EnterTypeModifiers is called when entering the typeModifiers production.
//	EnterTypeModifiers(c *TypeModifiersContext)
//
//	// EnterTypeModifier is called when entering the typeModifier production.
//	EnterTypeModifier(c *TypeModifierContext)
//
//	// EnterParenthesizedType is called when entering the parenthesizedType production.
//	EnterParenthesizedType(c *ParenthesizedTypeContext)
//
//	// EnterNullableType is called when entering the nullableType production.
//	EnterNullableType(c *NullableTypeContext)
//
//	// EnterTypeReference is called when entering the typeReference production.
//	EnterTypeReference(c *TypeReferenceContext)
//
//	// EnterFunctionType is called when entering the functionType production.
//	EnterFunctionType(c *FunctionTypeContext)
//
//	// EnterReceiverType is called when entering the receiverType production.
//	EnterReceiverType(c *ReceiverTypeContext)
//
//	// EnterUserType is called when entering the userType production.
//	EnterUserType(c *UserTypeContext)
//
//	// EnterParenthesizedUserType is called when entering the parenthesizedUserType production.
//	EnterParenthesizedUserType(c *ParenthesizedUserTypeContext)
//
//	// EnterSimpleUserType is called when entering the simpleUserType production.
//	EnterSimpleUserType(c *SimpleUserTypeContext)
//
//	// EnterFunctionTypeParameters is called when entering the functionTypeParameters production.
//	EnterFunctionTypeParameters(c *FunctionTypeParametersContext)
//
//	// EnterTypeConstraints is called when entering the typeConstraints production.
//	EnterTypeConstraints(c *TypeConstraintsContext)
//
//	// EnterTypeConstraint is called when entering the typeConstraint production.
//	EnterTypeConstraint(c *TypeConstraintContext)
//
//	// EnterBlock is called when entering the block production.
//	EnterBlock(c *BlockContext)
//
//	// EnterStatements is called when entering the statements production.
//	EnterStatements(c *StatementsContext)
//
//	// EnterStatement is called when entering the statement production.
//	EnterStatement(c *StatementContext)
//
//	// EnterDeclaration is called when entering the declaration production.
//	EnterDeclaration(c *DeclarationContext)
//
//	// EnterAssignment is called when entering the assignment production.
//	EnterAssignment(c *AssignmentContext)
//
//	// EnterExpression is called when entering the expression production.
//	EnterExpression(c *ExpressionContext)
//
//	// EnterDisjunction is called when entering the disjunction production.
//	EnterDisjunction(c *DisjunctionContext)
//
//	// EnterConjunction is called when entering the conjunction production.
//	EnterConjunction(c *ConjunctionContext)
//
//	// EnterEquality is called when entering the equality production.
//	EnterEquality(c *EqualityContext)
//
//	// EnterComparison is called when entering the comparison production.
//	EnterComparison(c *ComparisonContext)
//
//	// EnterInfixOperation is called when entering the infixOperation production.
//	EnterInfixOperation(c *InfixOperationContext)
//
//	// EnterElvisExpression is called when entering the elvisExpression production.
//	EnterElvisExpression(c *ElvisExpressionContext)
//
//	// EnterInfixFunctionCall is called when entering the infixFunctionCall production.
//	EnterInfixFunctionCall(c *InfixFunctionCallContext)
//
//	// EnterRangeExpression is called when entering the rangeExpression production.
//	EnterRangeExpression(c *RangeExpressionContext)
//
//	// EnterAdditiveExpression is called when entering the additiveExpression production.
//	EnterAdditiveExpression(c *AdditiveExpressionContext)
//
//	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
//	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)
//
//	// EnterAsExpression is called when entering the asExpression production.
//	EnterAsExpression(c *AsExpressionContext)
//
//	// EnterPrefixUnaryExpression is called when entering the prefixUnaryExpression production.
//	EnterPrefixUnaryExpression(c *PrefixUnaryExpressionContext)
//
//	// EnterUnaryPrefix is called when entering the unaryPrefix production.
//	EnterUnaryPrefix(c *UnaryPrefixContext)
//
//	// EnterPostfixUnaryExpression is called when entering the postfixUnaryExpression production.
//	EnterPostfixUnaryExpression(c *PostfixUnaryExpressionContext)
//
//	// EnterPostfixUnarySuffix is called when entering the postfixUnarySuffix production.
//	EnterPostfixUnarySuffix(c *PostfixUnarySuffixContext)
//
//	// EnterDirectlyAssignableExpression is called when entering the directlyAssignableExpression production.
//	EnterDirectlyAssignableExpression(c *DirectlyAssignableExpressionContext)
//
//	// EnterAssignableExpression is called when entering the assignableExpression production.
//	EnterAssignableExpression(c *AssignableExpressionContext)
//
//	// EnterAssignableSuffix is called when entering the assignableSuffix production.
//	EnterAssignableSuffix(c *AssignableSuffixContext)
//
//	// EnterIndexingSuffix is called when entering the indexingSuffix production.
//	EnterIndexingSuffix(c *IndexingSuffixContext)
//
//	// EnterNavigationSuffix is called when entering the navigationSuffix production.
//	EnterNavigationSuffix(c *NavigationSuffixContext)
//
//	// EnterCallSuffix is called when entering the callSuffix production.
//	EnterCallSuffix(c *CallSuffixContext)
//
//	// EnterAnnotatedLambda is called when entering the annotatedLambda production.
//	EnterAnnotatedLambda(c *AnnotatedLambdaContext)
//
//	// EnterValueArguments is called when entering the valueArguments production.
//	EnterValueArguments(c *ValueArgumentsContext)
//
//	// EnterTypeArguments is called when entering the typeArguments production.
//	EnterTypeArguments(c *TypeArgumentsContext)
//
//	// EnterTypeProjection is called when entering the typeProjection production.
//	EnterTypeProjection(c *TypeProjectionContext)
//
//	// EnterTypeProjectionModifiers is called when entering the typeProjectionModifiers production.
//	EnterTypeProjectionModifiers(c *TypeProjectionModifiersContext)
//
//	// EnterTypeProjectionModifier is called when entering the typeProjectionModifier production.
//	EnterTypeProjectionModifier(c *TypeProjectionModifierContext)
//
//	// EnterValueArgument is called when entering the valueArgument production.
//	EnterValueArgument(c *ValueArgumentContext)
//
//	// EnterPrimaryExpression is called when entering the primaryExpression production.
//	EnterPrimaryExpression(c *PrimaryExpressionContext)
//
//	// EnterParenthesizedExpression is called when entering the parenthesizedExpression production.
//	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)
//
//	// EnterCollectionLiteral is called when entering the collectionLiteral production.
//	EnterCollectionLiteral(c *CollectionLiteralContext)
//
//	// EnterLiteralConstant is called when entering the literalConstant production.
//	EnterLiteralConstant(c *LiteralConstantContext)
//
//	// EnterStringLiteral is called when entering the stringLiteral production.
//	EnterStringLiteral(c *StringLiteralContext)
//
//	// EnterLineStringLiteral is called when entering the lineStringLiteral production.
//	EnterLineStringLiteral(c *LineStringLiteralContext)
//
//	// EnterMultiLineStringLiteral is called when entering the multiLineStringLiteral production.
//	EnterMultiLineStringLiteral(c *MultiLineStringLiteralContext)
//
//	// EnterLineStringContent is called when entering the lineStringContent production.
//	EnterLineStringContent(c *LineStringContentContext)
//
//	// EnterLineStringExpression is called when entering the lineStringExpression production.
//	EnterLineStringExpression(c *LineStringExpressionContext)
//
//	// EnterMultiLineStringContent is called when entering the multiLineStringContent production.
//	EnterMultiLineStringContent(c *MultiLineStringContentContext)
//
//	// EnterMultiLineStringExpression is called when entering the multiLineStringExpression production.
//	EnterMultiLineStringExpression(c *MultiLineStringExpressionContext)
//
//	// EnterLambdaLiteral is called when entering the lambdaLiteral production.
//	EnterLambdaLiteral(c *LambdaLiteralContext)
//
//	// EnterLambdaParameters is called when entering the lambdaParameters production.
//	EnterLambdaParameters(c *LambdaParametersContext)
//
//	// EnterLambdaParameter is called when entering the lambdaParameter production.
//	EnterLambdaParameter(c *LambdaParameterContext)
//
//	// EnterAnonymousFunction is called when entering the anonymousFunction production.
//	EnterAnonymousFunction(c *AnonymousFunctionContext)
//
//	// EnterFunctionLiteral is called when entering the functionLiteral production.
//	EnterFunctionLiteral(c *FunctionLiteralContext)
//
//	// EnterObjectLiteral is called when entering the objectLiteral production.
//	EnterObjectLiteral(c *ObjectLiteralContext)
//
//	// EnterThisExpression is called when entering the thisExpression production.
//	EnterThisExpression(c *ThisExpressionContext)
//
//	// EnterSuperExpression is called when entering the superExpression production.
//	EnterSuperExpression(c *SuperExpressionContext)
//
//	// EnterControlStructureBody is called when entering the controlStructureBody production.
//	EnterControlStructureBody(c *ControlStructureBodyContext)
//
//	// EnterIfExpression is called when entering the ifExpression production.
//	EnterIfExpression(c *IfExpressionContext)
//
//	// EnterWhenExpression is called when entering the whenExpression production.
//	EnterWhenExpression(c *WhenExpressionContext)
//
//	// EnterWhenEntry is called when entering the whenEntry production.
//	EnterWhenEntry(c *WhenEntryContext)
//
//	// EnterWhenCondition is called when entering the whenCondition production.
//	EnterWhenCondition(c *WhenConditionContext)
//
//	// EnterRangeTest is called when entering the rangeTest production.
//	EnterRangeTest(c *RangeTestContext)
//
//	// EnterTypeTest is called when entering the typeTest production.
//	EnterTypeTest(c *TypeTestContext)
//
//	// EnterTryExpression is called when entering the tryExpression production.
//	EnterTryExpression(c *TryExpressionContext)
//
//	// EnterCatchBlock is called when entering the catchBlock production.
//	EnterCatchBlock(c *CatchBlockContext)
//
//	// EnterFinallyBlock is called when entering the finallyBlock production.
//	EnterFinallyBlock(c *FinallyBlockContext)
//
//	// EnterLoopStatement is called when entering the loopStatement production.
//	EnterLoopStatement(c *LoopStatementContext)
//
//	// EnterForStatement is called when entering the forStatement production.
//	EnterForStatement(c *ForStatementContext)
//
//	// EnterWhileStatement is called when entering the whileStatement production.
//	EnterWhileStatement(c *WhileStatementContext)
//
//	// EnterDoWhileStatement is called when entering the doWhileStatement production.
//	EnterDoWhileStatement(c *DoWhileStatementContext)
//
//	// EnterJumpExpression is called when entering the jumpExpression production.
//	EnterJumpExpression(c *JumpExpressionContext)
//
//	// EnterCallableReference is called when entering the callableReference production.
//	EnterCallableReference(c *CallableReferenceContext)
//
//	// EnterAssignmentAndOperator is called when entering the assignmentAndOperator production.
//	EnterAssignmentAndOperator(c *AssignmentAndOperatorContext)
//
//	// EnterEqualityOperator is called when entering the equalityOperator production.
//	EnterEqualityOperator(c *EqualityOperatorContext)
//
//	// EnterComparisonOperator is called when entering the comparisonOperator production.
//	EnterComparisonOperator(c *ComparisonOperatorContext)
//
//	// EnterInOperator is called when entering the inOperator production.
//	EnterInOperator(c *InOperatorContext)
//
//	// EnterIsOperator is called when entering the isOperator production.
//	EnterIsOperator(c *IsOperatorContext)
//
//	// EnterAdditiveOperator is called when entering the additiveOperator production.
//	EnterAdditiveOperator(c *AdditiveOperatorContext)
//
//	// EnterMultiplicativeOperator is called when entering the multiplicativeOperator production.
//	EnterMultiplicativeOperator(c *MultiplicativeOperatorContext)
//
//	// EnterAsOperator is called when entering the asOperator production.
//	EnterAsOperator(c *AsOperatorContext)
//
//	// EnterPrefixUnaryOperator is called when entering the prefixUnaryOperator production.
//	EnterPrefixUnaryOperator(c *PrefixUnaryOperatorContext)
//
//	// EnterPostfixUnaryOperator is called when entering the postfixUnaryOperator production.
//	EnterPostfixUnaryOperator(c *PostfixUnaryOperatorContext)
//
//	// EnterMemberAccessOperator is called when entering the memberAccessOperator production.
//	EnterMemberAccessOperator(c *MemberAccessOperatorContext)
//
//	// EnterModifiers is called when entering the modifiers production.
//	EnterModifiers(c *ModifiersContext)
//
//	// EnterModifier is called when entering the modifier production.
//	EnterModifier(c *ModifierContext)
//
//	// EnterClassModifier is called when entering the classModifier production.
//	EnterClassModifier(c *ClassModifierContext)
//
//	// EnterMemberModifier is called when entering the memberModifier production.
//	EnterMemberModifier(c *MemberModifierContext)
//
//	// EnterVisibilityModifier is called when entering the visibilityModifier production.
//	EnterVisibilityModifier(c *VisibilityModifierContext)
//
//	// EnterVarianceModifier is called when entering the varianceModifier production.
//	EnterVarianceModifier(c *VarianceModifierContext)
//
//	// EnterFunctionModifier is called when entering the functionModifier production.
//	EnterFunctionModifier(c *FunctionModifierContext)
//
//	// EnterPropertyModifier is called when entering the propertyModifier production.
//	EnterPropertyModifier(c *PropertyModifierContext)
//
//	// EnterInheritanceModifier is called when entering the inheritanceModifier production.
//	EnterInheritanceModifier(c *InheritanceModifierContext)
//
//	// EnterParameterModifier is called when entering the parameterModifier production.
//	EnterParameterModifier(c *ParameterModifierContext)
//
//	// EnterReificationModifier is called when entering the reificationModifier production.
//	EnterReificationModifier(c *ReificationModifierContext)
//
//	// EnterPlatformModifier is called when entering the platformModifier production.
//	EnterPlatformModifier(c *PlatformModifierContext)
//
//	// EnterLabel is called when entering the label production.
//	EnterLabel(c *LabelContext)
//
//	// EnterAnnotation is called when entering the annotation production.
//	EnterAnnotation(c *AnnotationContext)
//
//	// EnterSingleAnnotation is called when entering the singleAnnotation production.
//	EnterSingleAnnotation(c *SingleAnnotationContext)
//
//	// EnterMultiAnnotation is called when entering the multiAnnotation production.
//	EnterMultiAnnotation(c *MultiAnnotationContext)
//
//	// EnterAnnotationUseSiteTarget is called when entering the annotationUseSiteTarget production.
//	EnterAnnotationUseSiteTarget(c *AnnotationUseSiteTargetContext)
//
//	// EnterUnescapedAnnotation is called when entering the unescapedAnnotation production.
//	EnterUnescapedAnnotation(c *UnescapedAnnotationContext)
//
//	// EnterSimpleIdentifier is called when entering the simpleIdentifier production.
//	EnterSimpleIdentifier(c *SimpleIdentifierContext)
//
//	// EnterIdentifier is called when entering the identifier production.
//	EnterIdentifier(c *IdentifierContext)
//
//	// EnterShebangLine is called when entering the shebangLine production.
//	EnterShebangLine(c *ShebangLineContext)
//
//	// EnterQuest is called when entering the quest production.
//	EnterQuest(c *QuestContext)
//
//	// EnterElvis is called when entering the elvis production.
//	EnterElvis(c *ElvisContext)
//
//	// EnterSafeNav is called when entering the safeNav production.
//	EnterSafeNav(c *SafeNavContext)
//
//	// EnterExcl is called when entering the excl production.
//	EnterExcl(c *ExclContext)
//
//	// EnterSemi is called when entering the semi production.
//	EnterSemi(c *SemiContext)
//
//	// EnterSemis is called when entering the semis production.
//	EnterSemis(c *SemisContext)
//
//	// ExitKotlinFile is called when exiting the kotlinFile production.
//	ExitKotlinFile(c *KotlinFileContext)
//
//	// ExitScript is called when exiting the script production.
//	ExitScript(c *ScriptContext)
//
//	// ExitFileAnnotation is called when exiting the fileAnnotation production.
//	ExitFileAnnotation(c *FileAnnotationContext)
//
//	// ExitPackageHeader is called when exiting the packageHeader production.
//	ExitPackageHeader(c *PackageHeaderContext)
//
//	// ExitImportList is called when exiting the importList production.
//	ExitImportList(c *ImportListContext)
//
//	// ExitImportHeader is called when exiting the importHeader production.
//	ExitImportHeader(c *ImportHeaderContext)
//
//	// ExitImportAlias is called when exiting the importAlias production.
//	ExitImportAlias(c *ImportAliasContext)
//
//	// ExitTopLevelObject is called when exiting the topLevelObject production.
//	ExitTopLevelObject(c *TopLevelObjectContext)
//
//	// ExitClassDeclaration is called when exiting the classDeclaration production.
//	ExitClassDeclaration(c *ClassDeclarationContext)
//
//	// ExitPrimaryConstructor is called when exiting the primaryConstructor production.
//	ExitPrimaryConstructor(c *PrimaryConstructorContext)
//
//	// ExitClassParameters is called when exiting the classParameters production.
//	ExitClassParameters(c *ClassParametersContext)
//
//	// ExitClassParameter is called when exiting the classParameter production.
//	ExitClassParameter(c *ClassParameterContext)
//
//	// ExitDelegationSpecifiers is called when exiting the delegationSpecifiers production.
//	ExitDelegationSpecifiers(c *DelegationSpecifiersContext)
//
//	// ExitAnnotatedDelegationSpecifier is called when exiting the annotatedDelegationSpecifier production.
//	ExitAnnotatedDelegationSpecifier(c *AnnotatedDelegationSpecifierContext)
//
//	// ExitDelegationSpecifier is called when exiting the delegationSpecifier production.
//	ExitDelegationSpecifier(c *DelegationSpecifierContext)
//
//	// ExitConstructorInvocation is called when exiting the constructorInvocation production.
//	ExitConstructorInvocation(c *ConstructorInvocationContext)
//
//	// ExitExplicitDelegation is called when exiting the explicitDelegation production.
//	ExitExplicitDelegation(c *ExplicitDelegationContext)
//
//	// ExitClassBody is called when exiting the classBody production.
//	ExitClassBody(c *ClassBodyContext)
//
//	// ExitClassMemberDeclarations is called when exiting the classMemberDeclarations production.
//	ExitClassMemberDeclarations(c *ClassMemberDeclarationsContext)
//
//	// ExitClassMemberDeclaration is called when exiting the classMemberDeclaration production.
//	ExitClassMemberDeclaration(c *ClassMemberDeclarationContext)
//
//	// ExitAnonymousInitializer is called when exiting the anonymousInitializer production.
//	ExitAnonymousInitializer(c *AnonymousInitializerContext)
//
//	// ExitSecondaryConstructor is called when exiting the secondaryConstructor production.
//	ExitSecondaryConstructor(c *SecondaryConstructorContext)
//
//	// ExitConstructorDelegationCall is called when exiting the constructorDelegationCall production.
//	ExitConstructorDelegationCall(c *ConstructorDelegationCallContext)
//
//	// ExitEnumClassBody is called when exiting the enumClassBody production.
//	ExitEnumClassBody(c *EnumClassBodyContext)
//
//	// ExitEnumEntries is called when exiting the enumEntries production.
//	ExitEnumEntries(c *EnumEntriesContext)
//
//	// ExitEnumEntry is called when exiting the enumEntry production.
//	ExitEnumEntry(c *EnumEntryContext)
//
//	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
//	ExitFunctionDeclaration(c *FunctionDeclarationContext)
//
//	// ExitFunctionValueParameters is called when exiting the functionValueParameters production.
//	ExitFunctionValueParameters(c *FunctionValueParametersContext)
//
//	// ExitFunctionValueParameter is called when exiting the functionValueParameter production.
//	ExitFunctionValueParameter(c *FunctionValueParameterContext)
//
//	// ExitParameter is called when exiting the parameter production.
//	ExitParameter(c *ParameterContext)
//
//	// ExitSetterParameter is called when exiting the setterParameter production.
//	ExitSetterParameter(c *SetterParameterContext)
//
//	// ExitFunctionBody is called when exiting the functionBody production.
//	ExitFunctionBody(c *FunctionBodyContext)
//
//	// ExitObjectDeclaration is called when exiting the objectDeclaration production.
//	ExitObjectDeclaration(c *ObjectDeclarationContext)
//
//	// ExitCompanionObject is called when exiting the companionObject production.
//	ExitCompanionObject(c *CompanionObjectContext)
//
//	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
//	ExitPropertyDeclaration(c *PropertyDeclarationContext)
//
//	// ExitMultiVariableDeclaration is called when exiting the multiVariableDeclaration production.
//	ExitMultiVariableDeclaration(c *MultiVariableDeclarationContext)
//
//	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
//	ExitVariableDeclaration(c *VariableDeclarationContext)
//
//	// ExitPropertyDelegate is called when exiting the propertyDelegate production.
//	ExitPropertyDelegate(c *PropertyDelegateContext)
//
//	// ExitGetter is called when exiting the getter production.
//	ExitGetter(c *GetterContext)
//
//	// ExitSetter is called when exiting the setter production.
//	ExitSetter(c *SetterContext)
//
//	// ExitTypeAlias is called when exiting the typeAlias production.
//	ExitTypeAlias(c *TypeAliasContext)
//
//	// ExitTypeParameters is called when exiting the typeParameters production.
//	ExitTypeParameters(c *TypeParametersContext)
//
//	// ExitTypeParameter is called when exiting the typeParameter production.
//	ExitTypeParameter(c *TypeParameterContext)
//
//	// ExitTypeParameterModifiers is called when exiting the typeParameterModifiers production.
//	ExitTypeParameterModifiers(c *TypeParameterModifiersContext)
//
//	// ExitTypeParameterModifier is called when exiting the typeParameterModifier production.
//	ExitTypeParameterModifier(c *TypeParameterModifierContext)
//
//	// ExitType_ is called when exiting the type_ production.
//	ExitType_(c *Type_Context)
//
//	// ExitTypeModifiers is called when exiting the typeModifiers production.
//	ExitTypeModifiers(c *TypeModifiersContext)
//
//	// ExitTypeModifier is called when exiting the typeModifier production.
//	ExitTypeModifier(c *TypeModifierContext)
//
//	// ExitParenthesizedType is called when exiting the parenthesizedType production.
//	ExitParenthesizedType(c *ParenthesizedTypeContext)
//
//	// ExitNullableType is called when exiting the nullableType production.
//	ExitNullableType(c *NullableTypeContext)
//
//	// ExitTypeReference is called when exiting the typeReference production.
//	ExitTypeReference(c *TypeReferenceContext)
//
//	// ExitFunctionType is called when exiting the functionType production.
//	ExitFunctionType(c *FunctionTypeContext)
//
//	// ExitReceiverType is called when exiting the receiverType production.
//	ExitReceiverType(c *ReceiverTypeContext)
//
//	// ExitUserType is called when exiting the userType production.
//	ExitUserType(c *UserTypeContext)
//
//	// ExitParenthesizedUserType is called when exiting the parenthesizedUserType production.
//	ExitParenthesizedUserType(c *ParenthesizedUserTypeContext)
//
//	// ExitSimpleUserType is called when exiting the simpleUserType production.
//	ExitSimpleUserType(c *SimpleUserTypeContext)
//
//	// ExitFunctionTypeParameters is called when exiting the functionTypeParameters production.
//	ExitFunctionTypeParameters(c *FunctionTypeParametersContext)
//
//	// ExitTypeConstraints is called when exiting the typeConstraints production.
//	ExitTypeConstraints(c *TypeConstraintsContext)
//
//	// ExitTypeConstraint is called when exiting the typeConstraint production.
//	ExitTypeConstraint(c *TypeConstraintContext)
//
//	// ExitBlock is called when exiting the block production.
//	ExitBlock(c *BlockContext)
//
//	// ExitStatements is called when exiting the statements production.
//	ExitStatements(c *StatementsContext)
//
//	// ExitStatement is called when exiting the statement production.
//	ExitStatement(c *StatementContext)
//
//	// ExitDeclaration is called when exiting the declaration production.
//	ExitDeclaration(c *DeclarationContext)
//
//	// ExitAssignment is called when exiting the assignment production.
//	ExitAssignment(c *AssignmentContext)
//
//	// ExitExpression is called when exiting the expression production.
//	ExitExpression(c *ExpressionContext)
//
//	// ExitDisjunction is called when exiting the disjunction production.
//	ExitDisjunction(c *DisjunctionContext)
//
//	// ExitConjunction is called when exiting the conjunction production.
//	ExitConjunction(c *ConjunctionContext)
//
//	// ExitEquality is called when exiting the equality production.
//	ExitEquality(c *EqualityContext)
//
//	// ExitComparison is called when exiting the comparison production.
//	ExitComparison(c *ComparisonContext)
//
//	// ExitInfixOperation is called when exiting the infixOperation production.
//	ExitInfixOperation(c *InfixOperationContext)
//
//	// ExitElvisExpression is called when exiting the elvisExpression production.
//	ExitElvisExpression(c *ElvisExpressionContext)
//
//	// ExitInfixFunctionCall is called when exiting the infixFunctionCall production.
//	ExitInfixFunctionCall(c *InfixFunctionCallContext)
//
//	// ExitRangeExpression is called when exiting the rangeExpression production.
//	ExitRangeExpression(c *RangeExpressionContext)
//
//	// ExitAdditiveExpression is called when exiting the additiveExpression production.
//	ExitAdditiveExpression(c *AdditiveExpressionContext)
//
//	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
//	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)
//
//	// ExitAsExpression is called when exiting the asExpression production.
//	ExitAsExpression(c *AsExpressionContext)
//
//	// ExitPrefixUnaryExpression is called when exiting the prefixUnaryExpression production.
//	ExitPrefixUnaryExpression(c *PrefixUnaryExpressionContext)
//
//	// ExitUnaryPrefix is called when exiting the unaryPrefix production.
//	ExitUnaryPrefix(c *UnaryPrefixContext)
//
//	// ExitPostfixUnaryExpression is called when exiting the postfixUnaryExpression production.
//	ExitPostfixUnaryExpression(c *PostfixUnaryExpressionContext)
//
//	// ExitPostfixUnarySuffix is called when exiting the postfixUnarySuffix production.
//	ExitPostfixUnarySuffix(c *PostfixUnarySuffixContext)
//
//	// ExitDirectlyAssignableExpression is called when exiting the directlyAssignableExpression production.
//	ExitDirectlyAssignableExpression(c *DirectlyAssignableExpressionContext)
//
//	// ExitAssignableExpression is called when exiting the assignableExpression production.
//	ExitAssignableExpression(c *AssignableExpressionContext)
//
//	// ExitAssignableSuffix is called when exiting the assignableSuffix production.
//	ExitAssignableSuffix(c *AssignableSuffixContext)
//
//	// ExitIndexingSuffix is called when exiting the indexingSuffix production.
//	ExitIndexingSuffix(c *IndexingSuffixContext)
//
//	// ExitNavigationSuffix is called when exiting the navigationSuffix production.
//	ExitNavigationSuffix(c *NavigationSuffixContext)
//
//	// ExitCallSuffix is called when exiting the callSuffix production.
//	ExitCallSuffix(c *CallSuffixContext)
//
//	// ExitAnnotatedLambda is called when exiting the annotatedLambda production.
//	ExitAnnotatedLambda(c *AnnotatedLambdaContext)
//
//	// ExitValueArguments is called when exiting the valueArguments production.
//	ExitValueArguments(c *ValueArgumentsContext)
//
//	// ExitTypeArguments is called when exiting the typeArguments production.
//	ExitTypeArguments(c *TypeArgumentsContext)
//
//	// ExitTypeProjection is called when exiting the typeProjection production.
//	ExitTypeProjection(c *TypeProjectionContext)
//
//	// ExitTypeProjectionModifiers is called when exiting the typeProjectionModifiers production.
//	ExitTypeProjectionModifiers(c *TypeProjectionModifiersContext)
//
//	// ExitTypeProjectionModifier is called when exiting the typeProjectionModifier production.
//	ExitTypeProjectionModifier(c *TypeProjectionModifierContext)
//
//	// ExitValueArgument is called when exiting the valueArgument production.
//	ExitValueArgument(c *ValueArgumentContext)
//
//	// ExitPrimaryExpression is called when exiting the primaryExpression production.
//	ExitPrimaryExpression(c *PrimaryExpressionContext)
//
//	// ExitParenthesizedExpression is called when exiting the parenthesizedExpression production.
//	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)
//
//	// ExitCollectionLiteral is called when exiting the collectionLiteral production.
//	ExitCollectionLiteral(c *CollectionLiteralContext)
//
//	// ExitLiteralConstant is called when exiting the literalConstant production.
//	ExitLiteralConstant(c *LiteralConstantContext)
//
//	// ExitStringLiteral is called when exiting the stringLiteral production.
//	ExitStringLiteral(c *StringLiteralContext)
//
//	// ExitLineStringLiteral is called when exiting the lineStringLiteral production.
//	ExitLineStringLiteral(c *LineStringLiteralContext)
//
//	// ExitMultiLineStringLiteral is called when exiting the multiLineStringLiteral production.
//	ExitMultiLineStringLiteral(c *MultiLineStringLiteralContext)
//
//	// ExitLineStringContent is called when exiting the lineStringContent production.
//	ExitLineStringContent(c *LineStringContentContext)
//
//	// ExitLineStringExpression is called when exiting the lineStringExpression production.
//	ExitLineStringExpression(c *LineStringExpressionContext)
//
//	// ExitMultiLineStringContent is called when exiting the multiLineStringContent production.
//	ExitMultiLineStringContent(c *MultiLineStringContentContext)
//
//	// ExitMultiLineStringExpression is called when exiting the multiLineStringExpression production.
//	ExitMultiLineStringExpression(c *MultiLineStringExpressionContext)
//
//	// ExitLambdaLiteral is called when exiting the lambdaLiteral production.
//	ExitLambdaLiteral(c *LambdaLiteralContext)
//
//	// ExitLambdaParameters is called when exiting the lambdaParameters production.
//	ExitLambdaParameters(c *LambdaParametersContext)
//
//	// ExitLambdaParameter is called when exiting the lambdaParameter production.
//	ExitLambdaParameter(c *LambdaParameterContext)
//
//	// ExitAnonymousFunction is called when exiting the anonymousFunction production.
//	ExitAnonymousFunction(c *AnonymousFunctionContext)
//
//	// ExitFunctionLiteral is called when exiting the functionLiteral production.
//	ExitFunctionLiteral(c *FunctionLiteralContext)
//
//	// ExitObjectLiteral is called when exiting the objectLiteral production.
//	ExitObjectLiteral(c *ObjectLiteralContext)
//
//	// ExitThisExpression is called when exiting the thisExpression production.
//	ExitThisExpression(c *ThisExpressionContext)
//
//	// ExitSuperExpression is called when exiting the superExpression production.
//	ExitSuperExpression(c *SuperExpressionContext)
//
//	// ExitControlStructureBody is called when exiting the controlStructureBody production.
//	ExitControlStructureBody(c *ControlStructureBodyContext)
//
//	// ExitIfExpression is called when exiting the ifExpression production.
//	ExitIfExpression(c *IfExpressionContext)
//
//	// ExitWhenExpression is called when exiting the whenExpression production.
//	ExitWhenExpression(c *WhenExpressionContext)
//
//	// ExitWhenEntry is called when exiting the whenEntry production.
//	ExitWhenEntry(c *WhenEntryContext)
//
//	// ExitWhenCondition is called when exiting the whenCondition production.
//	ExitWhenCondition(c *WhenConditionContext)
//
//	// ExitRangeTest is called when exiting the rangeTest production.
//	ExitRangeTest(c *RangeTestContext)
//
//	// ExitTypeTest is called when exiting the typeTest production.
//	ExitTypeTest(c *TypeTestContext)
//
//	// ExitTryExpression is called when exiting the tryExpression production.
//	ExitTryExpression(c *TryExpressionContext)
//
//	// ExitCatchBlock is called when exiting the catchBlock production.
//	ExitCatchBlock(c *CatchBlockContext)
//
//	// ExitFinallyBlock is called when exiting the finallyBlock production.
//	ExitFinallyBlock(c *FinallyBlockContext)
//
//	// ExitLoopStatement is called when exiting the loopStatement production.
//	ExitLoopStatement(c *LoopStatementContext)
//
//	// ExitForStatement is called when exiting the forStatement production.
//	ExitForStatement(c *ForStatementContext)
//
//	// ExitWhileStatement is called when exiting the whileStatement production.
//	ExitWhileStatement(c *WhileStatementContext)
//
//	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
//	ExitDoWhileStatement(c *DoWhileStatementContext)
//
//	// ExitJumpExpression is called when exiting the jumpExpression production.
//	ExitJumpExpression(c *JumpExpressionContext)
//
//	// ExitCallableReference is called when exiting the callableReference production.
//	ExitCallableReference(c *CallableReferenceContext)
//
//	// ExitAssignmentAndOperator is called when exiting the assignmentAndOperator production.
//	ExitAssignmentAndOperator(c *AssignmentAndOperatorContext)
//
//	// ExitEqualityOperator is called when exiting the equalityOperator production.
//	ExitEqualityOperator(c *EqualityOperatorContext)
//
//	// ExitComparisonOperator is called when exiting the comparisonOperator production.
//	ExitComparisonOperator(c *ComparisonOperatorContext)
//
//	// ExitInOperator is called when exiting the inOperator production.
//	ExitInOperator(c *InOperatorContext)
//
//	// ExitIsOperator is called when exiting the isOperator production.
//	ExitIsOperator(c *IsOperatorContext)
//
//	// ExitAdditiveOperator is called when exiting the additiveOperator production.
//	ExitAdditiveOperator(c *AdditiveOperatorContext)
//
//	// ExitMultiplicativeOperator is called when exiting the multiplicativeOperator production.
//	ExitMultiplicativeOperator(c *MultiplicativeOperatorContext)
//
//	// ExitAsOperator is called when exiting the asOperator production.
//	ExitAsOperator(c *AsOperatorContext)
//
//	// ExitPrefixUnaryOperator is called when exiting the prefixUnaryOperator production.
//	ExitPrefixUnaryOperator(c *PrefixUnaryOperatorContext)
//
//	// ExitPostfixUnaryOperator is called when exiting the postfixUnaryOperator production.
//	ExitPostfixUnaryOperator(c *PostfixUnaryOperatorContext)
//
//	// ExitMemberAccessOperator is called when exiting the memberAccessOperator production.
//	ExitMemberAccessOperator(c *MemberAccessOperatorContext)
//
//	// ExitModifiers is called when exiting the modifiers production.
//	ExitModifiers(c *ModifiersContext)
//
//	// ExitModifier is called when exiting the modifier production.
//	ExitModifier(c *ModifierContext)
//
//	// ExitClassModifier is called when exiting the classModifier production.
//	ExitClassModifier(c *ClassModifierContext)
//
//	// ExitMemberModifier is called when exiting the memberModifier production.
//	ExitMemberModifier(c *MemberModifierContext)
//
//	// ExitVisibilityModifier is called when exiting the visibilityModifier production.
//	ExitVisibilityModifier(c *VisibilityModifierContext)
//
//	// ExitVarianceModifier is called when exiting the varianceModifier production.
//	ExitVarianceModifier(c *VarianceModifierContext)
//
//	// ExitFunctionModifier is called when exiting the functionModifier production.
//	ExitFunctionModifier(c *FunctionModifierContext)
//
//	// ExitPropertyModifier is called when exiting the propertyModifier production.
//	ExitPropertyModifier(c *PropertyModifierContext)
//
//	// ExitInheritanceModifier is called when exiting the inheritanceModifier production.
//	ExitInheritanceModifier(c *InheritanceModifierContext)
//
//	// ExitParameterModifier is called when exiting the parameterModifier production.
//	ExitParameterModifier(c *ParameterModifierContext)
//
//	// ExitReificationModifier is called when exiting the reificationModifier production.
//	ExitReificationModifier(c *ReificationModifierContext)
//
//	// ExitPlatformModifier is called when exiting the platformModifier production.
//	ExitPlatformModifier(c *PlatformModifierContext)
//
//	// ExitLabel is called when exiting the label production.
//	ExitLabel(c *LabelContext)
//
//	// ExitAnnotation is called when exiting the annotation production.
//	ExitAnnotation(c *AnnotationContext)
//
//	// ExitSingleAnnotation is called when exiting the singleAnnotation production.
//	ExitSingleAnnotation(c *SingleAnnotationContext)
//
//	// ExitMultiAnnotation is called when exiting the multiAnnotation production.
//	ExitMultiAnnotation(c *MultiAnnotationContext)
//
//	// ExitAnnotationUseSiteTarget is called when exiting the annotationUseSiteTarget production.
//	ExitAnnotationUseSiteTarget(c *AnnotationUseSiteTargetContext)
//
//	// ExitUnescapedAnnotation is called when exiting the unescapedAnnotation production.
//	ExitUnescapedAnnotation(c *UnescapedAnnotationContext)
//
//	// ExitSimpleIdentifier is called when exiting the simpleIdentifier production.
//	ExitSimpleIdentifier(c *SimpleIdentifierContext)
//
//	// ExitIdentifier is called when exiting the identifier production.
//	ExitIdentifier(c *IdentifierContext)
//
//	// ExitShebangLine is called when exiting the shebangLine production.
//	ExitShebangLine(c *ShebangLineContext)
//
//	// ExitQuest is called when exiting the quest production.
//	ExitQuest(c *QuestContext)
//
//	// ExitElvis is called when exiting the elvis production.
//	ExitElvis(c *ElvisContext)
//
//	// ExitSafeNav is called when exiting the safeNav production.
//	ExitSafeNav(c *SafeNavContext)
//
//	// ExitExcl is called when exiting the excl production.
//	ExitExcl(c *ExclContext)
//
//	// ExitSemi is called when exiting the semi production.
//	ExitSemi(c *SemiContext)
//
//	// ExitSemis is called when exiting the semis production.
//	ExitSemis(c *SemisContext)
//}
