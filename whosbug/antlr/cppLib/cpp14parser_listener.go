// Code generated from .\CPP14Parser.g4 by ANTLR 4.8. DO NOT EDIT.

package cppLib // CPP14Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CPP14ParserListener is a complete listener for a parse tree produced by CPP14Parser.
type CPP14ParserListener interface {
	antlr.ParseTreeListener

	// EnterTranslationUnit is called when entering the translationUnit production.
	EnterTranslationUnit(c *TranslationUnitContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterIdExpression is called when entering the idExpression production.
	EnterIdExpression(c *IdExpressionContext)

	// EnterUnqualifiedId is called when entering the unqualifiedId production.
	EnterUnqualifiedId(c *UnqualifiedIdContext)

	// EnterQualifiedId is called when entering the qualifiedId production.
	EnterQualifiedId(c *QualifiedIdContext)

	// EnterNestedNameSpecifier is called when entering the nestedNameSpecifier production.
	EnterNestedNameSpecifier(c *NestedNameSpecifierContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterLambdaIntroducer is called when entering the lambdaIntroducer production.
	EnterLambdaIntroducer(c *LambdaIntroducerContext)

	// EnterLambdaCapture is called when entering the lambdaCapture production.
	EnterLambdaCapture(c *LambdaCaptureContext)

	// EnterCaptureDefault is called when entering the captureDefault production.
	EnterCaptureDefault(c *CaptureDefaultContext)

	// EnterCaptureList is called when entering the captureList production.
	EnterCaptureList(c *CaptureListContext)

	// EnterCapture is called when entering the capture production.
	EnterCapture(c *CaptureContext)

	// EnterSimpleCapture is called when entering the simpleCapture production.
	EnterSimpleCapture(c *SimpleCaptureContext)

	// EnterInitcapture is called when entering the initcapture production.
	EnterInitcapture(c *InitcaptureContext)

	// EnterLambdaDeclarator is called when entering the lambdaDeclarator production.
	EnterLambdaDeclarator(c *LambdaDeclaratorContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterTypeIdOfTheTypeId is called when entering the typeIdOfTheTypeId production.
	EnterTypeIdOfTheTypeId(c *TypeIdOfTheTypeIdContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPseudoDestructorName is called when entering the pseudoDestructorName production.
	EnterPseudoDestructorName(c *PseudoDestructorNameContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterNewExpression is called when entering the newExpression production.
	EnterNewExpression(c *NewExpressionContext)

	// EnterNewPlacement is called when entering the newPlacement production.
	EnterNewPlacement(c *NewPlacementContext)

	// EnterNewTypeId is called when entering the newTypeId production.
	EnterNewTypeId(c *NewTypeIdContext)

	// EnterNewDeclarator is called when entering the newDeclarator production.
	EnterNewDeclarator(c *NewDeclaratorContext)

	// EnterNoPointerNewDeclarator is called when entering the noPointerNewDeclarator production.
	EnterNoPointerNewDeclarator(c *NoPointerNewDeclaratorContext)

	// EnterNewInitializer is called when entering the newInitializer production.
	EnterNewInitializer(c *NewInitializerContext)

	// EnterDeleteExpression is called when entering the deleteExpression production.
	EnterDeleteExpression(c *DeleteExpressionContext)

	// EnterNoExceptExpression is called when entering the noExceptExpression production.
	EnterNoExceptExpression(c *NoExceptExpressionContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterPointerMemberExpression is called when entering the pointerMemberExpression production.
	EnterPointerMemberExpression(c *PointerMemberExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterShiftOperator is called when entering the shiftOperator production.
	EnterShiftOperator(c *ShiftOperatorContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterInclusiveOrExpression is called when entering the inclusiveOrExpression production.
	EnterInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabeledStatement is called when entering the labeledStatement production.
	EnterLabeledStatement(c *LabeledStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterStatementSeq is called when entering the statementSeq production.
	EnterStatementSeq(c *StatementSeqContext)

	// EnterSelectionStatement is called when entering the selectionStatement production.
	EnterSelectionStatement(c *SelectionStatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterForInitStatement is called when entering the forInitStatement production.
	EnterForInitStatement(c *ForInitStatementContext)

	// EnterForRangeDeclaration is called when entering the forRangeDeclaration production.
	EnterForRangeDeclaration(c *ForRangeDeclarationContext)

	// EnterForRangeInitializer is called when entering the forRangeInitializer production.
	EnterForRangeInitializer(c *ForRangeInitializerContext)

	// EnterJumpStatement is called when entering the jumpStatement production.
	EnterJumpStatement(c *JumpStatementContext)

	// EnterDeclarationStatement is called when entering the declarationStatement production.
	EnterDeclarationStatement(c *DeclarationStatementContext)

	// EnterDeclarationseq is called when entering the declarationseq production.
	EnterDeclarationseq(c *DeclarationseqContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterBlockDeclaration is called when entering the blockDeclaration production.
	EnterBlockDeclaration(c *BlockDeclarationContext)

	// EnterAliasDeclaration is called when entering the aliasDeclaration production.
	EnterAliasDeclaration(c *AliasDeclarationContext)

	// EnterSimpleDeclaration is called when entering the simpleDeclaration production.
	EnterSimpleDeclaration(c *SimpleDeclarationContext)

	// EnterStaticAssertDeclaration is called when entering the staticAssertDeclaration production.
	EnterStaticAssertDeclaration(c *StaticAssertDeclarationContext)

	// EnterEmptyDeclaration is called when entering the emptyDeclaration production.
	EnterEmptyDeclaration(c *EmptyDeclarationContext)

	// EnterAttributeDeclaration is called when entering the attributeDeclaration production.
	EnterAttributeDeclaration(c *AttributeDeclarationContext)

	// EnterDeclSpecifier is called when entering the declSpecifier production.
	EnterDeclSpecifier(c *DeclSpecifierContext)

	// EnterDeclSpecifierSeq is called when entering the declSpecifierSeq production.
	EnterDeclSpecifierSeq(c *DeclSpecifierSeqContext)

	// EnterStorageClassSpecifier is called when entering the storageClassSpecifier production.
	EnterStorageClassSpecifier(c *StorageClassSpecifierContext)

	// EnterFunctionSpecifier is called when entering the functionSpecifier production.
	EnterFunctionSpecifier(c *FunctionSpecifierContext)

	// EnterTypedefName is called when entering the typedefName production.
	EnterTypedefName(c *TypedefNameContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterTrailingTypeSpecifier is called when entering the trailingTypeSpecifier production.
	EnterTrailingTypeSpecifier(c *TrailingTypeSpecifierContext)

	// EnterTypeSpecifierSeq is called when entering the typeSpecifierSeq production.
	EnterTypeSpecifierSeq(c *TypeSpecifierSeqContext)

	// EnterTrailingTypeSpecifierSeq is called when entering the trailingTypeSpecifierSeq production.
	EnterTrailingTypeSpecifierSeq(c *TrailingTypeSpecifierSeqContext)

	// EnterSimpleTypeLengthModifier is called when entering the simpleTypeLengthModifier production.
	EnterSimpleTypeLengthModifier(c *SimpleTypeLengthModifierContext)

	// EnterSimpleTypeSignednessModifier is called when entering the simpleTypeSignednessModifier production.
	EnterSimpleTypeSignednessModifier(c *SimpleTypeSignednessModifierContext)

	// EnterSimpleTypeSpecifier is called when entering the simpleTypeSpecifier production.
	EnterSimpleTypeSpecifier(c *SimpleTypeSpecifierContext)

	// EnterTheTypeName is called when entering the theTypeName production.
	EnterTheTypeName(c *TheTypeNameContext)

	// EnterDecltypeSpecifier is called when entering the decltypeSpecifier production.
	EnterDecltypeSpecifier(c *DecltypeSpecifierContext)

	// EnterElaboratedTypeSpecifier is called when entering the elaboratedTypeSpecifier production.
	EnterElaboratedTypeSpecifier(c *ElaboratedTypeSpecifierContext)

	// EnterEnumName is called when entering the enumName production.
	EnterEnumName(c *EnumNameContext)

	// EnterEnumSpecifier is called when entering the enumSpecifier production.
	EnterEnumSpecifier(c *EnumSpecifierContext)

	// EnterEnumHead is called when entering the enumHead production.
	EnterEnumHead(c *EnumHeadContext)

	// EnterOpaqueEnumDeclaration is called when entering the opaqueEnumDeclaration production.
	EnterOpaqueEnumDeclaration(c *OpaqueEnumDeclarationContext)

	// EnterEnumkey is called when entering the enumkey production.
	EnterEnumkey(c *EnumkeyContext)

	// EnterEnumbase is called when entering the enumbase production.
	EnterEnumbase(c *EnumbaseContext)

	// EnterEnumeratorList is called when entering the enumeratorList production.
	EnterEnumeratorList(c *EnumeratorListContext)

	// EnterEnumeratorDefinition is called when entering the enumeratorDefinition production.
	EnterEnumeratorDefinition(c *EnumeratorDefinitionContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterNamespaceName is called when entering the namespaceName production.
	EnterNamespaceName(c *NamespaceNameContext)

	// EnterOriginalNamespaceName is called when entering the originalNamespaceName production.
	EnterOriginalNamespaceName(c *OriginalNamespaceNameContext)

	// EnterNamespaceDefinition is called when entering the namespaceDefinition production.
	EnterNamespaceDefinition(c *NamespaceDefinitionContext)

	// EnterNamespaceAlias is called when entering the namespaceAlias production.
	EnterNamespaceAlias(c *NamespaceAliasContext)

	// EnterNamespaceAliasDefinition is called when entering the namespaceAliasDefinition production.
	EnterNamespaceAliasDefinition(c *NamespaceAliasDefinitionContext)

	// EnterQualifiednamespacespecifier is called when entering the qualifiednamespacespecifier production.
	EnterQualifiednamespacespecifier(c *QualifiednamespacespecifierContext)

	// EnterUsingDeclaration is called when entering the usingDeclaration production.
	EnterUsingDeclaration(c *UsingDeclarationContext)

	// EnterUsingDirective is called when entering the usingDirective production.
	EnterUsingDirective(c *UsingDirectiveContext)

	// EnterAsmDefinition is called when entering the asmDefinition production.
	EnterAsmDefinition(c *AsmDefinitionContext)

	// EnterLinkageSpecification is called when entering the linkageSpecification production.
	EnterLinkageSpecification(c *LinkageSpecificationContext)

	// EnterAttributeSpecifierSeq is called when entering the attributeSpecifierSeq production.
	EnterAttributeSpecifierSeq(c *AttributeSpecifierSeqContext)

	// EnterAttributeSpecifier is called when entering the attributeSpecifier production.
	EnterAttributeSpecifier(c *AttributeSpecifierContext)

	// EnterAlignmentspecifier is called when entering the alignmentspecifier production.
	EnterAlignmentspecifier(c *AlignmentspecifierContext)

	// EnterAttributeList is called when entering the attributeList production.
	EnterAttributeList(c *AttributeListContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterAttributeNamespace is called when entering the attributeNamespace production.
	EnterAttributeNamespace(c *AttributeNamespaceContext)

	// EnterAttributeArgumentClause is called when entering the attributeArgumentClause production.
	EnterAttributeArgumentClause(c *AttributeArgumentClauseContext)

	// EnterBalancedTokenSeq is called when entering the balancedTokenSeq production.
	EnterBalancedTokenSeq(c *BalancedTokenSeqContext)

	// EnterBalancedtoken is called when entering the balancedtoken production.
	EnterBalancedtoken(c *BalancedtokenContext)

	// EnterInitDeclaratorList is called when entering the initDeclaratorList production.
	EnterInitDeclaratorList(c *InitDeclaratorListContext)

	// EnterInitDeclarator is called when entering the initDeclarator production.
	EnterInitDeclarator(c *InitDeclaratorContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterPointerDeclarator is called when entering the pointerDeclarator production.
	EnterPointerDeclarator(c *PointerDeclaratorContext)

	// EnterNoPointerDeclarator is called when entering the noPointerDeclarator production.
	EnterNoPointerDeclarator(c *NoPointerDeclaratorContext)

	// EnterParametersAndQualifiers is called when entering the parametersAndQualifiers production.
	EnterParametersAndQualifiers(c *ParametersAndQualifiersContext)

	// EnterTrailingReturnType is called when entering the trailingReturnType production.
	EnterTrailingReturnType(c *TrailingReturnTypeContext)

	// EnterPointerOperator is called when entering the pointerOperator production.
	EnterPointerOperator(c *PointerOperatorContext)

	// EnterCvqualifierseq is called when entering the cvqualifierseq production.
	EnterCvqualifierseq(c *CvqualifierseqContext)

	// EnterCvQualifier is called when entering the cvQualifier production.
	EnterCvQualifier(c *CvQualifierContext)

	// EnterRefqualifier is called when entering the refqualifier production.
	EnterRefqualifier(c *RefqualifierContext)

	// EnterDeclaratorid is called when entering the declaratorid production.
	EnterDeclaratorid(c *DeclaratoridContext)

	// EnterTheTypeId is called when entering the theTypeId production.
	EnterTheTypeId(c *TheTypeIdContext)

	// EnterAbstractDeclarator is called when entering the abstractDeclarator production.
	EnterAbstractDeclarator(c *AbstractDeclaratorContext)

	// EnterPointerAbstractDeclarator is called when entering the pointerAbstractDeclarator production.
	EnterPointerAbstractDeclarator(c *PointerAbstractDeclaratorContext)

	// EnterNoPointerAbstractDeclarator is called when entering the noPointerAbstractDeclarator production.
	EnterNoPointerAbstractDeclarator(c *NoPointerAbstractDeclaratorContext)

	// EnterAbstractPackDeclarator is called when entering the abstractPackDeclarator production.
	EnterAbstractPackDeclarator(c *AbstractPackDeclaratorContext)

	// EnterNoPointerAbstractPackDeclarator is called when entering the noPointerAbstractPackDeclarator production.
	EnterNoPointerAbstractPackDeclarator(c *NoPointerAbstractPackDeclaratorContext)

	// EnterParameterDeclarationClause is called when entering the parameterDeclarationClause production.
	EnterParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// EnterParameterDeclarationList is called when entering the parameterDeclarationList production.
	EnterParameterDeclarationList(c *ParameterDeclarationListContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterBraceOrEqualInitializer is called when entering the braceOrEqualInitializer production.
	EnterBraceOrEqualInitializer(c *BraceOrEqualInitializerContext)

	// EnterInitializerClause is called when entering the initializerClause production.
	EnterInitializerClause(c *InitializerClauseContext)

	// EnterInitializerList is called when entering the initializerList production.
	EnterInitializerList(c *InitializerListContext)

	// EnterBracedInitList is called when entering the bracedInitList production.
	EnterBracedInitList(c *BracedInitListContext)

	// EnterClassName is called when entering the className production.
	EnterClassName(c *ClassNameContext)

	// EnterClassSpecifier is called when entering the classSpecifier production.
	EnterClassSpecifier(c *ClassSpecifierContext)

	// EnterClassHead is called when entering the classHead production.
	EnterClassHead(c *ClassHeadContext)

	// EnterClassHeadName is called when entering the classHeadName production.
	EnterClassHeadName(c *ClassHeadNameContext)

	// EnterClassVirtSpecifier is called when entering the classVirtSpecifier production.
	EnterClassVirtSpecifier(c *ClassVirtSpecifierContext)

	// EnterClassKey is called when entering the classKey production.
	EnterClassKey(c *ClassKeyContext)

	// EnterMemberSpecification is called when entering the memberSpecification production.
	EnterMemberSpecification(c *MemberSpecificationContext)

	// EnterMemberdeclaration is called when entering the memberdeclaration production.
	EnterMemberdeclaration(c *MemberdeclarationContext)

	// EnterMemberDeclaratorList is called when entering the memberDeclaratorList production.
	EnterMemberDeclaratorList(c *MemberDeclaratorListContext)

	// EnterMemberDeclarator is called when entering the memberDeclarator production.
	EnterMemberDeclarator(c *MemberDeclaratorContext)

	// EnterVirtualSpecifierSeq is called when entering the virtualSpecifierSeq production.
	EnterVirtualSpecifierSeq(c *VirtualSpecifierSeqContext)

	// EnterVirtualSpecifier is called when entering the virtualSpecifier production.
	EnterVirtualSpecifier(c *VirtualSpecifierContext)

	// EnterPureSpecifier is called when entering the pureSpecifier production.
	EnterPureSpecifier(c *PureSpecifierContext)

	// EnterBaseClause is called when entering the baseClause production.
	EnterBaseClause(c *BaseClauseContext)

	// EnterBaseSpecifierList is called when entering the baseSpecifierList production.
	EnterBaseSpecifierList(c *BaseSpecifierListContext)

	// EnterBaseSpecifier is called when entering the baseSpecifier production.
	EnterBaseSpecifier(c *BaseSpecifierContext)

	// EnterClassOrDeclType is called when entering the classOrDeclType production.
	EnterClassOrDeclType(c *ClassOrDeclTypeContext)

	// EnterBaseTypeSpecifier is called when entering the baseTypeSpecifier production.
	EnterBaseTypeSpecifier(c *BaseTypeSpecifierContext)

	// EnterAccessSpecifier is called when entering the accessSpecifier production.
	EnterAccessSpecifier(c *AccessSpecifierContext)

	// EnterConversionFunctionId is called when entering the conversionFunctionId production.
	EnterConversionFunctionId(c *ConversionFunctionIdContext)

	// EnterConversionTypeId is called when entering the conversionTypeId production.
	EnterConversionTypeId(c *ConversionTypeIdContext)

	// EnterConversionDeclarator is called when entering the conversionDeclarator production.
	EnterConversionDeclarator(c *ConversionDeclaratorContext)

	// EnterConstructorInitializer is called when entering the constructorInitializer production.
	EnterConstructorInitializer(c *ConstructorInitializerContext)

	// EnterMemInitializerList is called when entering the memInitializerList production.
	EnterMemInitializerList(c *MemInitializerListContext)

	// EnterMemInitializer is called when entering the memInitializer production.
	EnterMemInitializer(c *MemInitializerContext)

	// EnterMeminitializerid is called when entering the meminitializerid production.
	EnterMeminitializerid(c *MeminitializeridContext)

	// EnterOperatorFunctionId is called when entering the operatorFunctionId production.
	EnterOperatorFunctionId(c *OperatorFunctionIdContext)

	// EnterLiteralOperatorId is called when entering the literalOperatorId production.
	EnterLiteralOperatorId(c *LiteralOperatorIdContext)

	// EnterTemplateDeclaration is called when entering the templateDeclaration production.
	EnterTemplateDeclaration(c *TemplateDeclarationContext)

	// EnterTemplateparameterList is called when entering the templateparameterList production.
	EnterTemplateparameterList(c *TemplateparameterListContext)

	// EnterTemplateParameter is called when entering the templateParameter production.
	EnterTemplateParameter(c *TemplateParameterContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterSimpleTemplateId is called when entering the simpleTemplateId production.
	EnterSimpleTemplateId(c *SimpleTemplateIdContext)

	// EnterTemplateId is called when entering the templateId production.
	EnterTemplateId(c *TemplateIdContext)

	// EnterTemplateName is called when entering the templateName production.
	EnterTemplateName(c *TemplateNameContext)

	// EnterTemplateArgumentList is called when entering the templateArgumentList production.
	EnterTemplateArgumentList(c *TemplateArgumentListContext)

	// EnterTemplateArgument is called when entering the templateArgument production.
	EnterTemplateArgument(c *TemplateArgumentContext)

	// EnterTypeNameSpecifier is called when entering the typeNameSpecifier production.
	EnterTypeNameSpecifier(c *TypeNameSpecifierContext)

	// EnterExplicitInstantiation is called when entering the explicitInstantiation production.
	EnterExplicitInstantiation(c *ExplicitInstantiationContext)

	// EnterExplicitSpecialization is called when entering the explicitSpecialization production.
	EnterExplicitSpecialization(c *ExplicitSpecializationContext)

	// EnterTryBlock is called when entering the tryBlock production.
	EnterTryBlock(c *TryBlockContext)

	// EnterFunctionTryBlock is called when entering the functionTryBlock production.
	EnterFunctionTryBlock(c *FunctionTryBlockContext)

	// EnterHandlerSeq is called when entering the handlerSeq production.
	EnterHandlerSeq(c *HandlerSeqContext)

	// EnterHandler is called when entering the handler production.
	EnterHandler(c *HandlerContext)

	// EnterExceptionDeclaration is called when entering the exceptionDeclaration production.
	EnterExceptionDeclaration(c *ExceptionDeclarationContext)

	// EnterThrowExpression is called when entering the throwExpression production.
	EnterThrowExpression(c *ThrowExpressionContext)

	// EnterExceptionSpecification is called when entering the exceptionSpecification production.
	EnterExceptionSpecification(c *ExceptionSpecificationContext)

	// EnterDynamicExceptionSpecification is called when entering the dynamicExceptionSpecification production.
	EnterDynamicExceptionSpecification(c *DynamicExceptionSpecificationContext)

	// EnterTypeIdList is called when entering the typeIdList production.
	EnterTypeIdList(c *TypeIdListContext)

	// EnterNoeExceptSpecification is called when entering the noeExceptSpecification production.
	EnterNoeExceptSpecification(c *NoeExceptSpecificationContext)

	// EnterTheOperator is called when entering the theOperator production.
	EnterTheOperator(c *TheOperatorContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitTranslationUnit is called when exiting the translationUnit production.
	ExitTranslationUnit(c *TranslationUnitContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitIdExpression is called when exiting the idExpression production.
	ExitIdExpression(c *IdExpressionContext)

	// ExitUnqualifiedId is called when exiting the unqualifiedId production.
	ExitUnqualifiedId(c *UnqualifiedIdContext)

	// ExitQualifiedId is called when exiting the qualifiedId production.
	ExitQualifiedId(c *QualifiedIdContext)

	// ExitNestedNameSpecifier is called when exiting the nestedNameSpecifier production.
	ExitNestedNameSpecifier(c *NestedNameSpecifierContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitLambdaIntroducer is called when exiting the lambdaIntroducer production.
	ExitLambdaIntroducer(c *LambdaIntroducerContext)

	// ExitLambdaCapture is called when exiting the lambdaCapture production.
	ExitLambdaCapture(c *LambdaCaptureContext)

	// ExitCaptureDefault is called when exiting the captureDefault production.
	ExitCaptureDefault(c *CaptureDefaultContext)

	// ExitCaptureList is called when exiting the captureList production.
	ExitCaptureList(c *CaptureListContext)

	// ExitCapture is called when exiting the capture production.
	ExitCapture(c *CaptureContext)

	// ExitSimpleCapture is called when exiting the simpleCapture production.
	ExitSimpleCapture(c *SimpleCaptureContext)

	// ExitInitcapture is called when exiting the initcapture production.
	ExitInitcapture(c *InitcaptureContext)

	// ExitLambdaDeclarator is called when exiting the lambdaDeclarator production.
	ExitLambdaDeclarator(c *LambdaDeclaratorContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitTypeIdOfTheTypeId is called when exiting the typeIdOfTheTypeId production.
	ExitTypeIdOfTheTypeId(c *TypeIdOfTheTypeIdContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPseudoDestructorName is called when exiting the pseudoDestructorName production.
	ExitPseudoDestructorName(c *PseudoDestructorNameContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitNewExpression is called when exiting the newExpression production.
	ExitNewExpression(c *NewExpressionContext)

	// ExitNewPlacement is called when exiting the newPlacement production.
	ExitNewPlacement(c *NewPlacementContext)

	// ExitNewTypeId is called when exiting the newTypeId production.
	ExitNewTypeId(c *NewTypeIdContext)

	// ExitNewDeclarator is called when exiting the newDeclarator production.
	ExitNewDeclarator(c *NewDeclaratorContext)

	// ExitNoPointerNewDeclarator is called when exiting the noPointerNewDeclarator production.
	ExitNoPointerNewDeclarator(c *NoPointerNewDeclaratorContext)

	// ExitNewInitializer is called when exiting the newInitializer production.
	ExitNewInitializer(c *NewInitializerContext)

	// ExitDeleteExpression is called when exiting the deleteExpression production.
	ExitDeleteExpression(c *DeleteExpressionContext)

	// ExitNoExceptExpression is called when exiting the noExceptExpression production.
	ExitNoExceptExpression(c *NoExceptExpressionContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitPointerMemberExpression is called when exiting the pointerMemberExpression production.
	ExitPointerMemberExpression(c *PointerMemberExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitShiftOperator is called when exiting the shiftOperator production.
	ExitShiftOperator(c *ShiftOperatorContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitInclusiveOrExpression is called when exiting the inclusiveOrExpression production.
	ExitInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabeledStatement is called when exiting the labeledStatement production.
	ExitLabeledStatement(c *LabeledStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitStatementSeq is called when exiting the statementSeq production.
	ExitStatementSeq(c *StatementSeqContext)

	// ExitSelectionStatement is called when exiting the selectionStatement production.
	ExitSelectionStatement(c *SelectionStatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitForInitStatement is called when exiting the forInitStatement production.
	ExitForInitStatement(c *ForInitStatementContext)

	// ExitForRangeDeclaration is called when exiting the forRangeDeclaration production.
	ExitForRangeDeclaration(c *ForRangeDeclarationContext)

	// ExitForRangeInitializer is called when exiting the forRangeInitializer production.
	ExitForRangeInitializer(c *ForRangeInitializerContext)

	// ExitJumpStatement is called when exiting the jumpStatement production.
	ExitJumpStatement(c *JumpStatementContext)

	// ExitDeclarationStatement is called when exiting the declarationStatement production.
	ExitDeclarationStatement(c *DeclarationStatementContext)

	// ExitDeclarationseq is called when exiting the declarationseq production.
	ExitDeclarationseq(c *DeclarationseqContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitBlockDeclaration is called when exiting the blockDeclaration production.
	ExitBlockDeclaration(c *BlockDeclarationContext)

	// ExitAliasDeclaration is called when exiting the aliasDeclaration production.
	ExitAliasDeclaration(c *AliasDeclarationContext)

	// ExitSimpleDeclaration is called when exiting the simpleDeclaration production.
	ExitSimpleDeclaration(c *SimpleDeclarationContext)

	// ExitStaticAssertDeclaration is called when exiting the staticAssertDeclaration production.
	ExitStaticAssertDeclaration(c *StaticAssertDeclarationContext)

	// ExitEmptyDeclaration is called when exiting the emptyDeclaration production.
	ExitEmptyDeclaration(c *EmptyDeclarationContext)

	// ExitAttributeDeclaration is called when exiting the attributeDeclaration production.
	ExitAttributeDeclaration(c *AttributeDeclarationContext)

	// ExitDeclSpecifier is called when exiting the declSpecifier production.
	ExitDeclSpecifier(c *DeclSpecifierContext)

	// ExitDeclSpecifierSeq is called when exiting the declSpecifierSeq production.
	ExitDeclSpecifierSeq(c *DeclSpecifierSeqContext)

	// ExitStorageClassSpecifier is called when exiting the storageClassSpecifier production.
	ExitStorageClassSpecifier(c *StorageClassSpecifierContext)

	// ExitFunctionSpecifier is called when exiting the functionSpecifier production.
	ExitFunctionSpecifier(c *FunctionSpecifierContext)

	// ExitTypedefName is called when exiting the typedefName production.
	ExitTypedefName(c *TypedefNameContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitTrailingTypeSpecifier is called when exiting the trailingTypeSpecifier production.
	ExitTrailingTypeSpecifier(c *TrailingTypeSpecifierContext)

	// ExitTypeSpecifierSeq is called when exiting the typeSpecifierSeq production.
	ExitTypeSpecifierSeq(c *TypeSpecifierSeqContext)

	// ExitTrailingTypeSpecifierSeq is called when exiting the trailingTypeSpecifierSeq production.
	ExitTrailingTypeSpecifierSeq(c *TrailingTypeSpecifierSeqContext)

	// ExitSimpleTypeLengthModifier is called when exiting the simpleTypeLengthModifier production.
	ExitSimpleTypeLengthModifier(c *SimpleTypeLengthModifierContext)

	// ExitSimpleTypeSignednessModifier is called when exiting the simpleTypeSignednessModifier production.
	ExitSimpleTypeSignednessModifier(c *SimpleTypeSignednessModifierContext)

	// ExitSimpleTypeSpecifier is called when exiting the simpleTypeSpecifier production.
	ExitSimpleTypeSpecifier(c *SimpleTypeSpecifierContext)

	// ExitTheTypeName is called when exiting the theTypeName production.
	ExitTheTypeName(c *TheTypeNameContext)

	// ExitDecltypeSpecifier is called when exiting the decltypeSpecifier production.
	ExitDecltypeSpecifier(c *DecltypeSpecifierContext)

	// ExitElaboratedTypeSpecifier is called when exiting the elaboratedTypeSpecifier production.
	ExitElaboratedTypeSpecifier(c *ElaboratedTypeSpecifierContext)

	// ExitEnumName is called when exiting the enumName production.
	ExitEnumName(c *EnumNameContext)

	// ExitEnumSpecifier is called when exiting the enumSpecifier production.
	ExitEnumSpecifier(c *EnumSpecifierContext)

	// ExitEnumHead is called when exiting the enumHead production.
	ExitEnumHead(c *EnumHeadContext)

	// ExitOpaqueEnumDeclaration is called when exiting the opaqueEnumDeclaration production.
	ExitOpaqueEnumDeclaration(c *OpaqueEnumDeclarationContext)

	// ExitEnumkey is called when exiting the enumkey production.
	ExitEnumkey(c *EnumkeyContext)

	// ExitEnumbase is called when exiting the enumbase production.
	ExitEnumbase(c *EnumbaseContext)

	// ExitEnumeratorList is called when exiting the enumeratorList production.
	ExitEnumeratorList(c *EnumeratorListContext)

	// ExitEnumeratorDefinition is called when exiting the enumeratorDefinition production.
	ExitEnumeratorDefinition(c *EnumeratorDefinitionContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitNamespaceName is called when exiting the namespaceName production.
	ExitNamespaceName(c *NamespaceNameContext)

	// ExitOriginalNamespaceName is called when exiting the originalNamespaceName production.
	ExitOriginalNamespaceName(c *OriginalNamespaceNameContext)

	// ExitNamespaceDefinition is called when exiting the namespaceDefinition production.
	ExitNamespaceDefinition(c *NamespaceDefinitionContext)

	// ExitNamespaceAlias is called when exiting the namespaceAlias production.
	ExitNamespaceAlias(c *NamespaceAliasContext)

	// ExitNamespaceAliasDefinition is called when exiting the namespaceAliasDefinition production.
	ExitNamespaceAliasDefinition(c *NamespaceAliasDefinitionContext)

	// ExitQualifiednamespacespecifier is called when exiting the qualifiednamespacespecifier production.
	ExitQualifiednamespacespecifier(c *QualifiednamespacespecifierContext)

	// ExitUsingDeclaration is called when exiting the usingDeclaration production.
	ExitUsingDeclaration(c *UsingDeclarationContext)

	// ExitUsingDirective is called when exiting the usingDirective production.
	ExitUsingDirective(c *UsingDirectiveContext)

	// ExitAsmDefinition is called when exiting the asmDefinition production.
	ExitAsmDefinition(c *AsmDefinitionContext)

	// ExitLinkageSpecification is called when exiting the linkageSpecification production.
	ExitLinkageSpecification(c *LinkageSpecificationContext)

	// ExitAttributeSpecifierSeq is called when exiting the attributeSpecifierSeq production.
	ExitAttributeSpecifierSeq(c *AttributeSpecifierSeqContext)

	// ExitAttributeSpecifier is called when exiting the attributeSpecifier production.
	ExitAttributeSpecifier(c *AttributeSpecifierContext)

	// ExitAlignmentspecifier is called when exiting the alignmentspecifier production.
	ExitAlignmentspecifier(c *AlignmentspecifierContext)

	// ExitAttributeList is called when exiting the attributeList production.
	ExitAttributeList(c *AttributeListContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitAttributeNamespace is called when exiting the attributeNamespace production.
	ExitAttributeNamespace(c *AttributeNamespaceContext)

	// ExitAttributeArgumentClause is called when exiting the attributeArgumentClause production.
	ExitAttributeArgumentClause(c *AttributeArgumentClauseContext)

	// ExitBalancedTokenSeq is called when exiting the balancedTokenSeq production.
	ExitBalancedTokenSeq(c *BalancedTokenSeqContext)

	// ExitBalancedtoken is called when exiting the balancedtoken production.
	ExitBalancedtoken(c *BalancedtokenContext)

	// ExitInitDeclaratorList is called when exiting the initDeclaratorList production.
	ExitInitDeclaratorList(c *InitDeclaratorListContext)

	// ExitInitDeclarator is called when exiting the initDeclarator production.
	ExitInitDeclarator(c *InitDeclaratorContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitPointerDeclarator is called when exiting the pointerDeclarator production.
	ExitPointerDeclarator(c *PointerDeclaratorContext)

	// ExitNoPointerDeclarator is called when exiting the noPointerDeclarator production.
	ExitNoPointerDeclarator(c *NoPointerDeclaratorContext)

	// ExitParametersAndQualifiers is called when exiting the parametersAndQualifiers production.
	ExitParametersAndQualifiers(c *ParametersAndQualifiersContext)

	// ExitTrailingReturnType is called when exiting the trailingReturnType production.
	ExitTrailingReturnType(c *TrailingReturnTypeContext)

	// ExitPointerOperator is called when exiting the pointerOperator production.
	ExitPointerOperator(c *PointerOperatorContext)

	// ExitCvqualifierseq is called when exiting the cvqualifierseq production.
	ExitCvqualifierseq(c *CvqualifierseqContext)

	// ExitCvQualifier is called when exiting the cvQualifier production.
	ExitCvQualifier(c *CvQualifierContext)

	// ExitRefqualifier is called when exiting the refqualifier production.
	ExitRefqualifier(c *RefqualifierContext)

	// ExitDeclaratorid is called when exiting the declaratorid production.
	ExitDeclaratorid(c *DeclaratoridContext)

	// ExitTheTypeId is called when exiting the theTypeId production.
	ExitTheTypeId(c *TheTypeIdContext)

	// ExitAbstractDeclarator is called when exiting the abstractDeclarator production.
	ExitAbstractDeclarator(c *AbstractDeclaratorContext)

	// ExitPointerAbstractDeclarator is called when exiting the pointerAbstractDeclarator production.
	ExitPointerAbstractDeclarator(c *PointerAbstractDeclaratorContext)

	// ExitNoPointerAbstractDeclarator is called when exiting the noPointerAbstractDeclarator production.
	ExitNoPointerAbstractDeclarator(c *NoPointerAbstractDeclaratorContext)

	// ExitAbstractPackDeclarator is called when exiting the abstractPackDeclarator production.
	ExitAbstractPackDeclarator(c *AbstractPackDeclaratorContext)

	// ExitNoPointerAbstractPackDeclarator is called when exiting the noPointerAbstractPackDeclarator production.
	ExitNoPointerAbstractPackDeclarator(c *NoPointerAbstractPackDeclaratorContext)

	// ExitParameterDeclarationClause is called when exiting the parameterDeclarationClause production.
	ExitParameterDeclarationClause(c *ParameterDeclarationClauseContext)

	// ExitParameterDeclarationList is called when exiting the parameterDeclarationList production.
	ExitParameterDeclarationList(c *ParameterDeclarationListContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitBraceOrEqualInitializer is called when exiting the braceOrEqualInitializer production.
	ExitBraceOrEqualInitializer(c *BraceOrEqualInitializerContext)

	// ExitInitializerClause is called when exiting the initializerClause production.
	ExitInitializerClause(c *InitializerClauseContext)

	// ExitInitializerList is called when exiting the initializerList production.
	ExitInitializerList(c *InitializerListContext)

	// ExitBracedInitList is called when exiting the bracedInitList production.
	ExitBracedInitList(c *BracedInitListContext)

	// ExitClassName is called when exiting the className production.
	ExitClassName(c *ClassNameContext)

	// ExitClassSpecifier is called when exiting the classSpecifier production.
	ExitClassSpecifier(c *ClassSpecifierContext)

	// ExitClassHead is called when exiting the classHead production.
	ExitClassHead(c *ClassHeadContext)

	// ExitClassHeadName is called when exiting the classHeadName production.
	ExitClassHeadName(c *ClassHeadNameContext)

	// ExitClassVirtSpecifier is called when exiting the classVirtSpecifier production.
	ExitClassVirtSpecifier(c *ClassVirtSpecifierContext)

	// ExitClassKey is called when exiting the classKey production.
	ExitClassKey(c *ClassKeyContext)

	// ExitMemberSpecification is called when exiting the memberSpecification production.
	ExitMemberSpecification(c *MemberSpecificationContext)

	// ExitMemberdeclaration is called when exiting the memberdeclaration production.
	ExitMemberdeclaration(c *MemberdeclarationContext)

	// ExitMemberDeclaratorList is called when exiting the memberDeclaratorList production.
	ExitMemberDeclaratorList(c *MemberDeclaratorListContext)

	// ExitMemberDeclarator is called when exiting the memberDeclarator production.
	ExitMemberDeclarator(c *MemberDeclaratorContext)

	// ExitVirtualSpecifierSeq is called when exiting the virtualSpecifierSeq production.
	ExitVirtualSpecifierSeq(c *VirtualSpecifierSeqContext)

	// ExitVirtualSpecifier is called when exiting the virtualSpecifier production.
	ExitVirtualSpecifier(c *VirtualSpecifierContext)

	// ExitPureSpecifier is called when exiting the pureSpecifier production.
	ExitPureSpecifier(c *PureSpecifierContext)

	// ExitBaseClause is called when exiting the baseClause production.
	ExitBaseClause(c *BaseClauseContext)

	// ExitBaseSpecifierList is called when exiting the baseSpecifierList production.
	ExitBaseSpecifierList(c *BaseSpecifierListContext)

	// ExitBaseSpecifier is called when exiting the baseSpecifier production.
	ExitBaseSpecifier(c *BaseSpecifierContext)

	// ExitClassOrDeclType is called when exiting the classOrDeclType production.
	ExitClassOrDeclType(c *ClassOrDeclTypeContext)

	// ExitBaseTypeSpecifier is called when exiting the baseTypeSpecifier production.
	ExitBaseTypeSpecifier(c *BaseTypeSpecifierContext)

	// ExitAccessSpecifier is called when exiting the accessSpecifier production.
	ExitAccessSpecifier(c *AccessSpecifierContext)

	// ExitConversionFunctionId is called when exiting the conversionFunctionId production.
	ExitConversionFunctionId(c *ConversionFunctionIdContext)

	// ExitConversionTypeId is called when exiting the conversionTypeId production.
	ExitConversionTypeId(c *ConversionTypeIdContext)

	// ExitConversionDeclarator is called when exiting the conversionDeclarator production.
	ExitConversionDeclarator(c *ConversionDeclaratorContext)

	// ExitConstructorInitializer is called when exiting the constructorInitializer production.
	ExitConstructorInitializer(c *ConstructorInitializerContext)

	// ExitMemInitializerList is called when exiting the memInitializerList production.
	ExitMemInitializerList(c *MemInitializerListContext)

	// ExitMemInitializer is called when exiting the memInitializer production.
	ExitMemInitializer(c *MemInitializerContext)

	// ExitMeminitializerid is called when exiting the meminitializerid production.
	ExitMeminitializerid(c *MeminitializeridContext)

	// ExitOperatorFunctionId is called when exiting the operatorFunctionId production.
	ExitOperatorFunctionId(c *OperatorFunctionIdContext)

	// ExitLiteralOperatorId is called when exiting the literalOperatorId production.
	ExitLiteralOperatorId(c *LiteralOperatorIdContext)

	// ExitTemplateDeclaration is called when exiting the templateDeclaration production.
	ExitTemplateDeclaration(c *TemplateDeclarationContext)

	// ExitTemplateparameterList is called when exiting the templateparameterList production.
	ExitTemplateparameterList(c *TemplateparameterListContext)

	// ExitTemplateParameter is called when exiting the templateParameter production.
	ExitTemplateParameter(c *TemplateParameterContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitSimpleTemplateId is called when exiting the simpleTemplateId production.
	ExitSimpleTemplateId(c *SimpleTemplateIdContext)

	// ExitTemplateId is called when exiting the templateId production.
	ExitTemplateId(c *TemplateIdContext)

	// ExitTemplateName is called when exiting the templateName production.
	ExitTemplateName(c *TemplateNameContext)

	// ExitTemplateArgumentList is called when exiting the templateArgumentList production.
	ExitTemplateArgumentList(c *TemplateArgumentListContext)

	// ExitTemplateArgument is called when exiting the templateArgument production.
	ExitTemplateArgument(c *TemplateArgumentContext)

	// ExitTypeNameSpecifier is called when exiting the typeNameSpecifier production.
	ExitTypeNameSpecifier(c *TypeNameSpecifierContext)

	// ExitExplicitInstantiation is called when exiting the explicitInstantiation production.
	ExitExplicitInstantiation(c *ExplicitInstantiationContext)

	// ExitExplicitSpecialization is called when exiting the explicitSpecialization production.
	ExitExplicitSpecialization(c *ExplicitSpecializationContext)

	// ExitTryBlock is called when exiting the tryBlock production.
	ExitTryBlock(c *TryBlockContext)

	// ExitFunctionTryBlock is called when exiting the functionTryBlock production.
	ExitFunctionTryBlock(c *FunctionTryBlockContext)

	// ExitHandlerSeq is called when exiting the handlerSeq production.
	ExitHandlerSeq(c *HandlerSeqContext)

	// ExitHandler is called when exiting the handler production.
	ExitHandler(c *HandlerContext)

	// ExitExceptionDeclaration is called when exiting the exceptionDeclaration production.
	ExitExceptionDeclaration(c *ExceptionDeclarationContext)

	// ExitThrowExpression is called when exiting the throwExpression production.
	ExitThrowExpression(c *ThrowExpressionContext)

	// ExitExceptionSpecification is called when exiting the exceptionSpecification production.
	ExitExceptionSpecification(c *ExceptionSpecificationContext)

	// ExitDynamicExceptionSpecification is called when exiting the dynamicExceptionSpecification production.
	ExitDynamicExceptionSpecification(c *DynamicExceptionSpecificationContext)

	// ExitTypeIdList is called when exiting the typeIdList production.
	ExitTypeIdList(c *TypeIdListContext)

	// ExitNoeExceptSpecification is called when exiting the noeExceptSpecification production.
	ExitNoeExceptSpecification(c *NoeExceptSpecificationContext)

	// ExitTheOperator is called when exiting the theOperator production.
	ExitTheOperator(c *TheOperatorContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
