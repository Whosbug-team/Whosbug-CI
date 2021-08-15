// Code generated from JavaScriptParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JavaScriptParser

import (
	"fmt"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type JavaScriptParser struct {
	JavaScriptParserBase
}

// NewJavaScriptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *JavaScriptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJavaScriptParser(input antlr.TokenStream) *JavaScriptParser {
	this := new(JavaScriptParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "JavaScriptParser.g4"

	return this
}

// JavaScriptParser tokens.
const (
	JavaScriptParserEOF                        = antlr.TokenEOF
	JavaScriptParserHashBangLine               = 1
	JavaScriptParserMultiLineComment           = 2
	JavaScriptParserSingleLineComment          = 3
	JavaScriptParserRegularExpressionLiteral   = 4
	JavaScriptParserOpenBracket                = 5
	JavaScriptParserCloseBracket               = 6
	JavaScriptParserOpenParen                  = 7
	JavaScriptParserCloseParen                 = 8
	JavaScriptParserOpenBrace                  = 9
	JavaScriptParserCloseBrace                 = 10
	JavaScriptParserSemiColon                  = 11
	JavaScriptParserComma                      = 12
	JavaScriptParserAssign                     = 13
	JavaScriptParserQuestionMark               = 14
	JavaScriptParserColon                      = 15
	JavaScriptParserEllipsis                   = 16
	JavaScriptParserDot                        = 17
	JavaScriptParserPlusPlus                   = 18
	JavaScriptParserMinusMinus                 = 19
	JavaScriptParserPlus                       = 20
	JavaScriptParserMinus                      = 21
	JavaScriptParserBitNot                     = 22
	JavaScriptParserNot                        = 23
	JavaScriptParserMultiply                   = 24
	JavaScriptParserDivide                     = 25
	JavaScriptParserModulus                    = 26
	JavaScriptParserPower                      = 27
	JavaScriptParserNullCoalesce               = 28
	JavaScriptParserHashtag                    = 29
	JavaScriptParserRightShiftArithmetic       = 30
	JavaScriptParserLeftShiftArithmetic        = 31
	JavaScriptParserRightShiftLogical          = 32
	JavaScriptParserLessThan                   = 33
	JavaScriptParserMoreThan                   = 34
	JavaScriptParserLessThanEquals             = 35
	JavaScriptParserGreaterThanEquals          = 36
	JavaScriptParserEquals_                    = 37
	JavaScriptParserNotEquals                  = 38
	JavaScriptParserIdentityEquals             = 39
	JavaScriptParserIdentityNotEquals          = 40
	JavaScriptParserBitAnd                     = 41
	JavaScriptParserBitXOr                     = 42
	JavaScriptParserBitOr                      = 43
	JavaScriptParserAnd                        = 44
	JavaScriptParserOr                         = 45
	JavaScriptParserMultiplyAssign             = 46
	JavaScriptParserDivideAssign               = 47
	JavaScriptParserModulusAssign              = 48
	JavaScriptParserPlusAssign                 = 49
	JavaScriptParserMinusAssign                = 50
	JavaScriptParserLeftShiftArithmeticAssign  = 51
	JavaScriptParserRightShiftArithmeticAssign = 52
	JavaScriptParserRightShiftLogicalAssign    = 53
	JavaScriptParserBitAndAssign               = 54
	JavaScriptParserBitXorAssign               = 55
	JavaScriptParserBitOrAssign                = 56
	JavaScriptParserPowerAssign                = 57
	JavaScriptParserARROW                      = 58
	JavaScriptParserNullLiteral                = 59
	JavaScriptParserBooleanLiteral             = 60
	JavaScriptParserDecimalLiteral             = 61
	JavaScriptParserHexIntegerLiteral          = 62
	JavaScriptParserOctalIntegerLiteral        = 63
	JavaScriptParserOctalIntegerLiteral2       = 64
	JavaScriptParserBinaryIntegerLiteral       = 65
	JavaScriptParserBigHexIntegerLiteral       = 66
	JavaScriptParserBigOctalIntegerLiteral     = 67
	JavaScriptParserBigBinaryIntegerLiteral    = 68
	JavaScriptParserBigDecimalIntegerLiteral   = 69
	JavaScriptParserBreak                      = 70
	JavaScriptParserDo                         = 71
	JavaScriptParserInstanceof                 = 72
	JavaScriptParserTypeof                     = 73
	JavaScriptParserCase                       = 74
	JavaScriptParserElse                       = 75
	JavaScriptParserNew                        = 76
	JavaScriptParserVar                        = 77
	JavaScriptParserCatch                      = 78
	JavaScriptParserFinally                    = 79
	JavaScriptParserReturn                     = 80
	JavaScriptParserVoid                       = 81
	JavaScriptParserContinue                   = 82
	JavaScriptParserFor                        = 83
	JavaScriptParserSwitch                     = 84
	JavaScriptParserWhile                      = 85
	JavaScriptParserDebugger                   = 86
	JavaScriptParserFunction                   = 87
	JavaScriptParserThis                       = 88
	JavaScriptParserWith                       = 89
	JavaScriptParserDefault                    = 90
	JavaScriptParserIf                         = 91
	JavaScriptParserThrow                      = 92
	JavaScriptParserDelete                     = 93
	JavaScriptParserIn                         = 94
	JavaScriptParserTry                        = 95
	JavaScriptParserAs                         = 96
	JavaScriptParserFrom                       = 97
	JavaScriptParserClass                      = 98
	JavaScriptParserEnum                       = 99
	JavaScriptParserExtends                    = 100
	JavaScriptParserSuper                      = 101
	JavaScriptParserConst                      = 102
	JavaScriptParserExport                     = 103
	JavaScriptParserImport                     = 104
	JavaScriptParserAsync                      = 105
	JavaScriptParserAwait                      = 106
	JavaScriptParserImplements                 = 107
	JavaScriptParserStrictLet                  = 108
	JavaScriptParserNonStrictLet               = 109
	JavaScriptParserPrivate                    = 110
	JavaScriptParserPublic                     = 111
	JavaScriptParserInterface                  = 112
	JavaScriptParserPackage                    = 113
	JavaScriptParserProtected                  = 114
	JavaScriptParserStatic                     = 115
	JavaScriptParserYield                      = 116
	JavaScriptParserIdentifier                 = 117
	JavaScriptParserStringLiteral              = 118
	JavaScriptParserTemplateStringLiteral      = 119
	JavaScriptParserWhiteSpaces                = 120
	JavaScriptParserLineTerminator             = 121
	JavaScriptParserHtmlComment                = 122
	JavaScriptParserCDataComment               = 123
	JavaScriptParserUnexpectedCharacter        = 124
)

// JavaScriptParser rules.
const (
	JavaScriptParserRULE_program                 = 0
	JavaScriptParserRULE_sourceElement           = 1
	JavaScriptParserRULE_statement               = 2
	JavaScriptParserRULE_block                   = 3
	JavaScriptParserRULE_statementList           = 4
	JavaScriptParserRULE_importStatement         = 5
	JavaScriptParserRULE_importFromBlock         = 6
	JavaScriptParserRULE_moduleItems             = 7
	JavaScriptParserRULE_importDefault           = 8
	JavaScriptParserRULE_importNamespace         = 9
	JavaScriptParserRULE_importFrom              = 10
	JavaScriptParserRULE_aliasName               = 11
	JavaScriptParserRULE_exportStatement         = 12
	JavaScriptParserRULE_exportFromBlock         = 13
	JavaScriptParserRULE_declaration             = 14
	JavaScriptParserRULE_variableStatement       = 15
	JavaScriptParserRULE_variableDeclarationList = 16
	JavaScriptParserRULE_variableDeclaration     = 17
	JavaScriptParserRULE_emptyStatement          = 18
	JavaScriptParserRULE_expressionStatement     = 19
	JavaScriptParserRULE_ifStatement             = 20
	JavaScriptParserRULE_iterationStatement      = 21
	JavaScriptParserRULE_varModifier             = 22
	JavaScriptParserRULE_continueStatement       = 23
	JavaScriptParserRULE_breakStatement          = 24
	JavaScriptParserRULE_returnStatement         = 25
	JavaScriptParserRULE_yieldStatement          = 26
	JavaScriptParserRULE_withStatement           = 27
	JavaScriptParserRULE_switchStatement         = 28
	JavaScriptParserRULE_caseBlock               = 29
	JavaScriptParserRULE_caseClauses             = 30
	JavaScriptParserRULE_caseClause              = 31
	JavaScriptParserRULE_defaultClause           = 32
	JavaScriptParserRULE_labelledStatement       = 33
	JavaScriptParserRULE_throwStatement          = 34
	JavaScriptParserRULE_tryStatement            = 35
	JavaScriptParserRULE_catchProduction         = 36
	JavaScriptParserRULE_finallyProduction       = 37
	JavaScriptParserRULE_debuggerStatement       = 38
	JavaScriptParserRULE_functionDeclaration     = 39
	JavaScriptParserRULE_classDeclaration        = 40
	JavaScriptParserRULE_classTail               = 41
	JavaScriptParserRULE_classElement            = 42
	JavaScriptParserRULE_methodDefinition        = 43
	JavaScriptParserRULE_formalParameterList     = 44
	JavaScriptParserRULE_formalParameterArg      = 45
	JavaScriptParserRULE_lastFormalParameterArg  = 46
	JavaScriptParserRULE_functionBody            = 47
	JavaScriptParserRULE_sourceElements          = 48
	JavaScriptParserRULE_arrayLiteral            = 49
	JavaScriptParserRULE_elementList             = 50
	JavaScriptParserRULE_arrayElement            = 51
	JavaScriptParserRULE_propertyAssignment      = 52
	JavaScriptParserRULE_propertyName            = 53
	JavaScriptParserRULE_arguments               = 54
	JavaScriptParserRULE_argument                = 55
	JavaScriptParserRULE_expressionSequence      = 56
	JavaScriptParserRULE_singleExpression        = 57
	JavaScriptParserRULE_assignable              = 58
	JavaScriptParserRULE_objectLiteral           = 59
	JavaScriptParserRULE_anoymousFunction        = 60
	JavaScriptParserRULE_arrowFunctionParameters = 61
	JavaScriptParserRULE_arrowFunctionBody       = 62
	JavaScriptParserRULE_assignmentOperator      = 63
	JavaScriptParserRULE_literal                 = 64
	JavaScriptParserRULE_numericLiteral          = 65
	JavaScriptParserRULE_bigintLiteral           = 66
	JavaScriptParserRULE_getter                  = 67
	JavaScriptParserRULE_setter                  = 68
	JavaScriptParserRULE_identifierName          = 69
	JavaScriptParserRULE_identifier              = 70
	JavaScriptParserRULE_reservedWord            = 71
	JavaScriptParserRULE_keyword                 = 72
	JavaScriptParserRULE_let_                    = 73
	JavaScriptParserRULE_eos                     = 74
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEOF, 0)
}

func (s *ProgramContext) HashBangLine() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserHashBangLine, 0)
}

func (s *ProgramContext) SourceElements() ISourceElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceElementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *JavaScriptParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JavaScriptParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(150)
			p.Match(JavaScriptParserHashBangLine)
		}

	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(153)
			p.SourceElements()
		}

	}
	{
		p.SetState(156)
		p.Match(JavaScriptParserEOF)
	}

	return localctx
}

// ISourceElementContext is an interface to support dynamic dispatch.
type ISourceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementContext differentiates from other interfaces.
	IsSourceElementContext()
}

type SourceElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementContext() *SourceElementContext {
	var p = new(SourceElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_sourceElement
	return p
}

func (*SourceElementContext) IsSourceElementContext() {}

func NewSourceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementContext {
	var p = new(SourceElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_sourceElement

	return p
}

func (s *SourceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SourceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterSourceElement(s)
	}
}

func (s *SourceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitSourceElement(s)
	}
}

func (p *JavaScriptParser) SourceElement() (localctx ISourceElementContext) {
	localctx = NewSourceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JavaScriptParserRULE_sourceElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Statement()
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

// func NewEmptyStatementContext() *StatementContext {
// 	var p = new(StatementContext)
// 	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
// 	p.RuleIndex = JavaScriptParserRULE_statement
// 	return p
// }

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) VariableStatement() IVariableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableStatementContext)
}

func (s *StatementContext) ImportStatement() IImportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *StatementContext) ExportStatement() IExportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportStatementContext)
}

func (s *StatementContext) EmptyStatement() IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *StatementContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) IterationStatement() IIterationStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) YieldStatement() IYieldStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYieldStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYieldStatementContext)
}

func (s *StatementContext) WithStatement() IWithStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithStatementContext)
}

func (s *StatementContext) LabelledStatement() ILabelledStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelledStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelledStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) ThrowStatement() IThrowStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThrowStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThrowStatementContext)
}

func (s *StatementContext) TryStatement() ITryStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITryStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITryStatementContext)
}

func (s *StatementContext) DebuggerStatement() IDebuggerStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDebuggerStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDebuggerStatementContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *JavaScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JavaScriptParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.VariableStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.ImportStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(163)
			p.ExportStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(164)
			p.EmptyStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(165)
			p.ClassDeclaration()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(166)
			p.ExpressionStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(167)
			p.IfStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(168)
			p.IterationStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(169)
			p.ContinueStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(170)
			p.BreakStatement()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(171)
			p.ReturnStatement()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(172)
			p.YieldStatement()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(173)
			p.WithStatement()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(174)
			p.LabelledStatement()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(175)
			p.SwitchStatement()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(176)
			p.ThrowStatement()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(177)
			p.TryStatement()
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(178)
			p.DebuggerStatement()
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(179)
			p.FunctionDeclaration()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBrace, 0)
}

func (s *BlockContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBrace, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *JavaScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JavaScriptParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(JavaScriptParserOpenBrace)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(183)
			p.StatementList()
		}

	}
	{
		p.SetState(186)
		p.Match(JavaScriptParserCloseBrace)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *JavaScriptParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JavaScriptParserRULE_statementList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(188)
				p.Statement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) Import() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserImport, 0)
}

func (s *ImportStatementContext) ImportFromBlock() IImportFromBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportFromBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportFromBlockContext)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *JavaScriptParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JavaScriptParserRULE_importStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(JavaScriptParserImport)
	}
	{
		p.SetState(194)
		p.ImportFromBlock()
	}

	return localctx
}

// IImportFromBlockContext is an interface to support dynamic dispatch.
type IImportFromBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportFromBlockContext differentiates from other interfaces.
	IsImportFromBlockContext()
}

type ImportFromBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportFromBlockContext() *ImportFromBlockContext {
	var p = new(ImportFromBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_importFromBlock
	return p
}

func (*ImportFromBlockContext) IsImportFromBlockContext() {}

func NewImportFromBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportFromBlockContext {
	var p = new(ImportFromBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_importFromBlock

	return p
}

func (s *ImportFromBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportFromBlockContext) ImportFrom() IImportFromContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportFromContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportFromContext)
}

func (s *ImportFromBlockContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportFromBlockContext) ImportNamespace() IImportNamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportNamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportNamespaceContext)
}

func (s *ImportFromBlockContext) ModuleItems() IModuleItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleItemsContext)
}

func (s *ImportFromBlockContext) ImportDefault() IImportDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDefaultContext)
}

func (s *ImportFromBlockContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserStringLiteral, 0)
}

func (s *ImportFromBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportFromBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportFromBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterImportFromBlock(s)
	}
}

func (s *ImportFromBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitImportFromBlock(s)
	}
}

func (p *JavaScriptParser) ImportFromBlock() (localctx IImportFromBlockContext) {
	localctx = NewImportFromBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JavaScriptParserRULE_importFromBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(208)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserOpenBrace, JavaScriptParserMultiply, JavaScriptParserNullLiteral, JavaScriptParserBooleanLiteral, JavaScriptParserBreak, JavaScriptParserDo, JavaScriptParserInstanceof, JavaScriptParserTypeof, JavaScriptParserCase, JavaScriptParserElse, JavaScriptParserNew, JavaScriptParserVar, JavaScriptParserCatch, JavaScriptParserFinally, JavaScriptParserReturn, JavaScriptParserVoid, JavaScriptParserContinue, JavaScriptParserFor, JavaScriptParserSwitch, JavaScriptParserWhile, JavaScriptParserDebugger, JavaScriptParserFunction, JavaScriptParserThis, JavaScriptParserWith, JavaScriptParserDefault, JavaScriptParserIf, JavaScriptParserThrow, JavaScriptParserDelete, JavaScriptParserIn, JavaScriptParserTry, JavaScriptParserAs, JavaScriptParserFrom, JavaScriptParserClass, JavaScriptParserEnum, JavaScriptParserExtends, JavaScriptParserSuper, JavaScriptParserConst, JavaScriptParserExport, JavaScriptParserImport, JavaScriptParserAsync, JavaScriptParserAwait, JavaScriptParserImplements, JavaScriptParserStrictLet, JavaScriptParserNonStrictLet, JavaScriptParserPrivate, JavaScriptParserPublic, JavaScriptParserInterface, JavaScriptParserPackage, JavaScriptParserProtected, JavaScriptParserStatic, JavaScriptParserYield, JavaScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(197)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(196)
				p.ImportDefault()
			}

		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case JavaScriptParserMultiply, JavaScriptParserNullLiteral, JavaScriptParserBooleanLiteral, JavaScriptParserBreak, JavaScriptParserDo, JavaScriptParserInstanceof, JavaScriptParserTypeof, JavaScriptParserCase, JavaScriptParserElse, JavaScriptParserNew, JavaScriptParserVar, JavaScriptParserCatch, JavaScriptParserFinally, JavaScriptParserReturn, JavaScriptParserVoid, JavaScriptParserContinue, JavaScriptParserFor, JavaScriptParserSwitch, JavaScriptParserWhile, JavaScriptParserDebugger, JavaScriptParserFunction, JavaScriptParserThis, JavaScriptParserWith, JavaScriptParserDefault, JavaScriptParserIf, JavaScriptParserThrow, JavaScriptParserDelete, JavaScriptParserIn, JavaScriptParserTry, JavaScriptParserAs, JavaScriptParserFrom, JavaScriptParserClass, JavaScriptParserEnum, JavaScriptParserExtends, JavaScriptParserSuper, JavaScriptParserConst, JavaScriptParserExport, JavaScriptParserImport, JavaScriptParserAsync, JavaScriptParserAwait, JavaScriptParserImplements, JavaScriptParserStrictLet, JavaScriptParserNonStrictLet, JavaScriptParserPrivate, JavaScriptParserPublic, JavaScriptParserInterface, JavaScriptParserPackage, JavaScriptParserProtected, JavaScriptParserStatic, JavaScriptParserYield, JavaScriptParserIdentifier:
			{
				p.SetState(199)
				p.ImportNamespace()
			}

		case JavaScriptParserOpenBrace:
			{
				p.SetState(200)
				p.ModuleItems()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(203)
			p.ImportFrom()
		}
		{
			p.SetState(204)
			p.Eos()
		}

	case JavaScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(JavaScriptParserStringLiteral)
		}
		{
			p.SetState(207)
			p.Eos()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IModuleItemsContext is an interface to support dynamic dispatch.
type IModuleItemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleItemsContext differentiates from other interfaces.
	IsModuleItemsContext()
}

type ModuleItemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleItemsContext() *ModuleItemsContext {
	var p = new(ModuleItemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_moduleItems
	return p
}

func (*ModuleItemsContext) IsModuleItemsContext() {}

func NewModuleItemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleItemsContext {
	var p = new(ModuleItemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_moduleItems

	return p
}

func (s *ModuleItemsContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleItemsContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBrace, 0)
}

func (s *ModuleItemsContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBrace, 0)
}

func (s *ModuleItemsContext) AllAliasName() []IAliasNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAliasNameContext)(nil)).Elem())
	var tst = make([]IAliasNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAliasNameContext)
		}
	}

	return tst
}

func (s *ModuleItemsContext) AliasName(i int) IAliasNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAliasNameContext)
}

func (s *ModuleItemsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserComma)
}

func (s *ModuleItemsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, i)
}

func (s *ModuleItemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleItemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleItemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterModuleItems(s)
	}
}

func (s *ModuleItemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitModuleItems(s)
	}
}

func (p *JavaScriptParser) ModuleItems() (localctx IModuleItemsContext) {
	localctx = NewModuleItemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JavaScriptParserRULE_moduleItems)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(JavaScriptParserOpenBrace)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(211)
				p.AliasName()
			}
			{
				p.SetState(212)
				p.Match(JavaScriptParserComma)
			}

		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(JavaScriptParserNullLiteral-59))|(1<<(JavaScriptParserBooleanLiteral-59))|(1<<(JavaScriptParserBreak-59))|(1<<(JavaScriptParserDo-59))|(1<<(JavaScriptParserInstanceof-59))|(1<<(JavaScriptParserTypeof-59))|(1<<(JavaScriptParserCase-59))|(1<<(JavaScriptParserElse-59))|(1<<(JavaScriptParserNew-59))|(1<<(JavaScriptParserVar-59))|(1<<(JavaScriptParserCatch-59))|(1<<(JavaScriptParserFinally-59))|(1<<(JavaScriptParserReturn-59))|(1<<(JavaScriptParserVoid-59))|(1<<(JavaScriptParserContinue-59))|(1<<(JavaScriptParserFor-59))|(1<<(JavaScriptParserSwitch-59))|(1<<(JavaScriptParserWhile-59))|(1<<(JavaScriptParserDebugger-59))|(1<<(JavaScriptParserFunction-59))|(1<<(JavaScriptParserThis-59))|(1<<(JavaScriptParserWith-59))|(1<<(JavaScriptParserDefault-59)))) != 0) || (((_la-91)&-(0x1f+1)) == 0 && ((1<<uint((_la-91)))&((1<<(JavaScriptParserIf-91))|(1<<(JavaScriptParserThrow-91))|(1<<(JavaScriptParserDelete-91))|(1<<(JavaScriptParserIn-91))|(1<<(JavaScriptParserTry-91))|(1<<(JavaScriptParserAs-91))|(1<<(JavaScriptParserFrom-91))|(1<<(JavaScriptParserClass-91))|(1<<(JavaScriptParserEnum-91))|(1<<(JavaScriptParserExtends-91))|(1<<(JavaScriptParserSuper-91))|(1<<(JavaScriptParserConst-91))|(1<<(JavaScriptParserExport-91))|(1<<(JavaScriptParserImport-91))|(1<<(JavaScriptParserAsync-91))|(1<<(JavaScriptParserAwait-91))|(1<<(JavaScriptParserImplements-91))|(1<<(JavaScriptParserStrictLet-91))|(1<<(JavaScriptParserNonStrictLet-91))|(1<<(JavaScriptParserPrivate-91))|(1<<(JavaScriptParserPublic-91))|(1<<(JavaScriptParserInterface-91))|(1<<(JavaScriptParserPackage-91))|(1<<(JavaScriptParserProtected-91))|(1<<(JavaScriptParserStatic-91))|(1<<(JavaScriptParserYield-91))|(1<<(JavaScriptParserIdentifier-91)))) != 0) {
		{
			p.SetState(219)
			p.AliasName()
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserComma {
			{
				p.SetState(220)
				p.Match(JavaScriptParserComma)
			}

		}

	}
	{
		p.SetState(225)
		p.Match(JavaScriptParserCloseBrace)
	}

	return localctx
}

// IImportDefaultContext is an interface to support dynamic dispatch.
type IImportDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDefaultContext differentiates from other interfaces.
	IsImportDefaultContext()
}

type ImportDefaultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDefaultContext() *ImportDefaultContext {
	var p = new(ImportDefaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_importDefault
	return p
}

func (*ImportDefaultContext) IsImportDefaultContext() {}

func NewImportDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDefaultContext {
	var p = new(ImportDefaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_importDefault

	return p
}

func (s *ImportDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDefaultContext) AliasName() IAliasNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasNameContext)
}

func (s *ImportDefaultContext) Comma() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, 0)
}

func (s *ImportDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterImportDefault(s)
	}
}

func (s *ImportDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitImportDefault(s)
	}
}

func (p *JavaScriptParser) ImportDefault() (localctx IImportDefaultContext) {
	localctx = NewImportDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JavaScriptParserRULE_importDefault)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.AliasName()
	}
	{
		p.SetState(228)
		p.Match(JavaScriptParserComma)
	}

	return localctx
}

// IImportNamespaceContext is an interface to support dynamic dispatch.
type IImportNamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportNamespaceContext differentiates from other interfaces.
	IsImportNamespaceContext()
}

type ImportNamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportNamespaceContext() *ImportNamespaceContext {
	var p = new(ImportNamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_importNamespace
	return p
}

func (*ImportNamespaceContext) IsImportNamespaceContext() {}

func NewImportNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportNamespaceContext {
	var p = new(ImportNamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_importNamespace

	return p
}

func (s *ImportNamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportNamespaceContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMultiply, 0)
}

func (s *ImportNamespaceContext) AllIdentifierName() []IIdentifierNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem())
	var tst = make([]IIdentifierNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierNameContext)
		}
	}

	return tst
}

func (s *ImportNamespaceContext) IdentifierName(i int) IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *ImportNamespaceContext) As() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAs, 0)
}

func (s *ImportNamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportNamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportNamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterImportNamespace(s)
	}
}

func (s *ImportNamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitImportNamespace(s)
	}
}

func (p *JavaScriptParser) ImportNamespace() (localctx IImportNamespaceContext) {
	localctx = NewImportNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JavaScriptParserRULE_importNamespace)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserMultiply:
		{
			p.SetState(230)
			p.Match(JavaScriptParserMultiply)
		}

	case JavaScriptParserNullLiteral, JavaScriptParserBooleanLiteral, JavaScriptParserBreak, JavaScriptParserDo, JavaScriptParserInstanceof, JavaScriptParserTypeof, JavaScriptParserCase, JavaScriptParserElse, JavaScriptParserNew, JavaScriptParserVar, JavaScriptParserCatch, JavaScriptParserFinally, JavaScriptParserReturn, JavaScriptParserVoid, JavaScriptParserContinue, JavaScriptParserFor, JavaScriptParserSwitch, JavaScriptParserWhile, JavaScriptParserDebugger, JavaScriptParserFunction, JavaScriptParserThis, JavaScriptParserWith, JavaScriptParserDefault, JavaScriptParserIf, JavaScriptParserThrow, JavaScriptParserDelete, JavaScriptParserIn, JavaScriptParserTry, JavaScriptParserAs, JavaScriptParserFrom, JavaScriptParserClass, JavaScriptParserEnum, JavaScriptParserExtends, JavaScriptParserSuper, JavaScriptParserConst, JavaScriptParserExport, JavaScriptParserImport, JavaScriptParserAsync, JavaScriptParserAwait, JavaScriptParserImplements, JavaScriptParserStrictLet, JavaScriptParserNonStrictLet, JavaScriptParserPrivate, JavaScriptParserPublic, JavaScriptParserInterface, JavaScriptParserPackage, JavaScriptParserProtected, JavaScriptParserStatic, JavaScriptParserYield, JavaScriptParserIdentifier:
		{
			p.SetState(231)
			p.IdentifierName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserAs {
		{
			p.SetState(234)
			p.Match(JavaScriptParserAs)
		}
		{
			p.SetState(235)
			p.IdentifierName()
		}

	}

	return localctx
}

// IImportFromContext is an interface to support dynamic dispatch.
type IImportFromContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportFromContext differentiates from other interfaces.
	IsImportFromContext()
}

type ImportFromContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportFromContext() *ImportFromContext {
	var p = new(ImportFromContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_importFrom
	return p
}

func (*ImportFromContext) IsImportFromContext() {}

func NewImportFromContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportFromContext {
	var p = new(ImportFromContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_importFrom

	return p
}

func (s *ImportFromContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportFromContext) From() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFrom, 0)
}

func (s *ImportFromContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserStringLiteral, 0)
}

func (s *ImportFromContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportFromContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportFromContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterImportFrom(s)
	}
}

func (s *ImportFromContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitImportFrom(s)
	}
}

func (p *JavaScriptParser) ImportFrom() (localctx IImportFromContext) {
	localctx = NewImportFromContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JavaScriptParserRULE_importFrom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(JavaScriptParserFrom)
	}
	{
		p.SetState(239)
		p.Match(JavaScriptParserStringLiteral)
	}

	return localctx
}

// IAliasNameContext is an interface to support dynamic dispatch.
type IAliasNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasNameContext differentiates from other interfaces.
	IsAliasNameContext()
}

type AliasNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasNameContext() *AliasNameContext {
	var p = new(AliasNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_aliasName
	return p
}

func (*AliasNameContext) IsAliasNameContext() {}

func NewAliasNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasNameContext {
	var p = new(AliasNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_aliasName

	return p
}

func (s *AliasNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasNameContext) AllIdentifierName() []IIdentifierNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem())
	var tst = make([]IIdentifierNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierNameContext)
		}
	}

	return tst
}

func (s *AliasNameContext) IdentifierName(i int) IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *AliasNameContext) As() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAs, 0)
}

func (s *AliasNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAliasName(s)
	}
}

func (s *AliasNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAliasName(s)
	}
}

func (p *JavaScriptParser) AliasName() (localctx IAliasNameContext) {
	localctx = NewAliasNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JavaScriptParserRULE_aliasName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.IdentifierName()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserAs {
		{
			p.SetState(242)
			p.Match(JavaScriptParserAs)
		}
		{
			p.SetState(243)
			p.IdentifierName()
		}

	}

	return localctx
}

// IExportStatementContext is an interface to support dynamic dispatch.
type IExportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportStatementContext differentiates from other interfaces.
	IsExportStatementContext()
}

type ExportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportStatementContext() *ExportStatementContext {
	var p = new(ExportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_exportStatement
	return p
}

func (*ExportStatementContext) IsExportStatementContext() {}

func NewExportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportStatementContext {
	var p = new(ExportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_exportStatement

	return p
}

func (s *ExportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportStatementContext) CopyFrom(ctx *ExportStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExportDefaultDeclarationContext struct {
	*ExportStatementContext
}

func NewExportDefaultDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExportDefaultDeclarationContext {
	var p = new(ExportDefaultDeclarationContext)

	p.ExportStatementContext = NewEmptyExportStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExportStatementContext))

	return p
}

func (s *ExportDefaultDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportDefaultDeclarationContext) Export() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserExport, 0)
}

func (s *ExportDefaultDeclarationContext) Default() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDefault, 0)
}

func (s *ExportDefaultDeclarationContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ExportDefaultDeclarationContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ExportDefaultDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterExportDefaultDeclaration(s)
	}
}

func (s *ExportDefaultDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitExportDefaultDeclaration(s)
	}
}

type ExportDeclarationContext struct {
	*ExportStatementContext
}

func NewExportDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExportDeclarationContext {
	var p = new(ExportDeclarationContext)

	p.ExportStatementContext = NewEmptyExportStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExportStatementContext))

	return p
}

func (s *ExportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportDeclarationContext) Export() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserExport, 0)
}

func (s *ExportDeclarationContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ExportDeclarationContext) ExportFromBlock() IExportFromBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportFromBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportFromBlockContext)
}

func (s *ExportDeclarationContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ExportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterExportDeclaration(s)
	}
}

func (s *ExportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitExportDeclaration(s)
	}
}

func (p *JavaScriptParser) ExportStatement() (localctx IExportStatementContext) {
	localctx = NewExportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JavaScriptParserRULE_exportStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExportDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.Match(JavaScriptParserExport)
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(247)
				p.ExportFromBlock()
			}

		case 2:
			{
				p.SetState(248)
				p.Declaration()
			}

		}
		{
			p.SetState(251)
			p.Eos()
		}

	case 2:
		localctx = NewExportDefaultDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(JavaScriptParserExport)
		}
		{
			p.SetState(254)
			p.Match(JavaScriptParserDefault)
		}
		{
			p.SetState(255)
			p.singleExpression(0)
		}
		{
			p.SetState(256)
			p.Eos()
		}

	}

	return localctx
}

// IExportFromBlockContext is an interface to support dynamic dispatch.
type IExportFromBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportFromBlockContext differentiates from other interfaces.
	IsExportFromBlockContext()
}

type ExportFromBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportFromBlockContext() *ExportFromBlockContext {
	var p = new(ExportFromBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_exportFromBlock
	return p
}

func (*ExportFromBlockContext) IsExportFromBlockContext() {}

func NewExportFromBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportFromBlockContext {
	var p = new(ExportFromBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_exportFromBlock

	return p
}

func (s *ExportFromBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportFromBlockContext) ImportNamespace() IImportNamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportNamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportNamespaceContext)
}

func (s *ExportFromBlockContext) ImportFrom() IImportFromContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportFromContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportFromContext)
}

func (s *ExportFromBlockContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ExportFromBlockContext) ModuleItems() IModuleItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleItemsContext)
}

func (s *ExportFromBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportFromBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportFromBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterExportFromBlock(s)
	}
}

func (s *ExportFromBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitExportFromBlock(s)
	}
}

func (p *JavaScriptParser) ExportFromBlock() (localctx IExportFromBlockContext) {
	localctx = NewExportFromBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JavaScriptParserRULE_exportFromBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserMultiply, JavaScriptParserNullLiteral, JavaScriptParserBooleanLiteral, JavaScriptParserBreak, JavaScriptParserDo, JavaScriptParserInstanceof, JavaScriptParserTypeof, JavaScriptParserCase, JavaScriptParserElse, JavaScriptParserNew, JavaScriptParserVar, JavaScriptParserCatch, JavaScriptParserFinally, JavaScriptParserReturn, JavaScriptParserVoid, JavaScriptParserContinue, JavaScriptParserFor, JavaScriptParserSwitch, JavaScriptParserWhile, JavaScriptParserDebugger, JavaScriptParserFunction, JavaScriptParserThis, JavaScriptParserWith, JavaScriptParserDefault, JavaScriptParserIf, JavaScriptParserThrow, JavaScriptParserDelete, JavaScriptParserIn, JavaScriptParserTry, JavaScriptParserAs, JavaScriptParserFrom, JavaScriptParserClass, JavaScriptParserEnum, JavaScriptParserExtends, JavaScriptParserSuper, JavaScriptParserConst, JavaScriptParserExport, JavaScriptParserImport, JavaScriptParserAsync, JavaScriptParserAwait, JavaScriptParserImplements, JavaScriptParserStrictLet, JavaScriptParserNonStrictLet, JavaScriptParserPrivate, JavaScriptParserPublic, JavaScriptParserInterface, JavaScriptParserPackage, JavaScriptParserProtected, JavaScriptParserStatic, JavaScriptParserYield, JavaScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.ImportNamespace()
		}
		{
			p.SetState(261)
			p.ImportFrom()
		}
		{
			p.SetState(262)
			p.Eos()
		}

	case JavaScriptParserOpenBrace:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.ModuleItems()
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(265)
				p.ImportFrom()
			}

		}
		{
			p.SetState(268)
			p.Eos()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VariableStatement() IVariableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableStatementContext)
}

func (s *DeclarationContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *JavaScriptParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JavaScriptParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserVar, JavaScriptParserConst, JavaScriptParserStrictLet, JavaScriptParserNonStrictLet:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.VariableStatement()
		}

	case JavaScriptParserClass:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.ClassDeclaration()
		}

	case JavaScriptParserFunction, JavaScriptParserAsync:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(274)
			p.FunctionDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableStatementContext is an interface to support dynamic dispatch.
type IVariableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableStatementContext differentiates from other interfaces.
	IsVariableStatementContext()
}

type VariableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableStatementContext() *VariableStatementContext {
	var p = new(VariableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_variableStatement
	return p
}

func (*VariableStatementContext) IsVariableStatementContext() {}

func NewVariableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableStatementContext {
	var p = new(VariableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_variableStatement

	return p
}

func (s *VariableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *VariableStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *VariableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterVariableStatement(s)
	}
}

func (s *VariableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitVariableStatement(s)
	}
}

func (p *JavaScriptParser) VariableStatement() (localctx IVariableStatementContext) {
	localctx = NewVariableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JavaScriptParserRULE_variableStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.VariableDeclarationList()
	}
	{
		p.SetState(278)
		p.Eos()
	}

	return localctx
}

// IVariableDeclarationListContext is an interface to support dynamic dispatch.
type IVariableDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationListContext differentiates from other interfaces.
	IsVariableDeclarationListContext()
}

type VariableDeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationListContext() *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_variableDeclarationList
	return p
}

func (*VariableDeclarationListContext) IsVariableDeclarationListContext() {}

func NewVariableDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_variableDeclarationList

	return p
}

func (s *VariableDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationListContext) VarModifier() IVarModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarModifierContext)
}

func (s *VariableDeclarationListContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *VariableDeclarationListContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *VariableDeclarationListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserComma)
}

func (s *VariableDeclarationListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, i)
}

func (s *VariableDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterVariableDeclarationList(s)
	}
}

func (s *VariableDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitVariableDeclarationList(s)
	}
}

func (p *JavaScriptParser) VariableDeclarationList() (localctx IVariableDeclarationListContext) {
	localctx = NewVariableDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JavaScriptParserRULE_variableDeclarationList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.VarModifier()
	}
	{
		p.SetState(281)
		p.VariableDeclaration()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(282)
				p.Match(JavaScriptParserComma)
			}
			{
				p.SetState(283)
				p.VariableDeclaration()
			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Assignable() IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *VariableDeclarationContext) Assign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAssign, 0)
}

func (s *VariableDeclarationContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *JavaScriptParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JavaScriptParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Assignable()
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(290)
			p.Match(JavaScriptParserAssign)
		}
		{
			p.SetState(291)
			p.singleExpression(0)
		}

	}

	return localctx
}

// IEmptyStatementContext is an interface to support dynamic dispatch.
type IEmptyStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStatementContext differentiates from other interfaces.
	IsEmptyStatementContext()
}

type EmptyStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatementContext() *EmptyStatementContext {
	var p = new(EmptyStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_emptyStatement
	return p
}

func (*EmptyStatementContext) IsEmptyStatementContext() {}

func NewEmptyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatementContext {
	var p = new(EmptyStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_emptyStatement

	return p
}

func (s *EmptyStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSemiColon, 0)
}

func (s *EmptyStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterEmptyStatement(s)
	}
}

func (s *EmptyStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitEmptyStatement(s)
	}
}

func (p *JavaScriptParser) EmptyStatement() (localctx IEmptyStatementContext) {
	localctx = NewEmptyStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JavaScriptParserRULE_emptyStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(JavaScriptParserSemiColon)
	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ExpressionStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (p *JavaScriptParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JavaScriptParserRULE_expressionStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(296)

	if !(p.notOpenBraceAndNotFunction()) {
		panic(antlr.NewFailedPredicateException(p, "p.notOpenBraceAndNotFunction()", ""))
	}
	{
		p.SetState(297)
		p.ExpressionSequence()
	}
	{
		p.SetState(298)
		p.Eos()
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIf, 0)
}

func (s *IfStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *IfStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *IfStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserElse, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *JavaScriptParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JavaScriptParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(JavaScriptParserIf)
	}
	{
		p.SetState(301)
		p.Match(JavaScriptParserOpenParen)
	}
	{
		p.SetState(302)
		p.ExpressionSequence()
	}
	{
		p.SetState(303)
		p.Match(JavaScriptParserCloseParen)
	}
	{
		p.SetState(304)
		p.Statement()
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(305)
			p.Match(JavaScriptParserElse)
		}
		{
			p.SetState(306)
			p.Statement()
		}

	}

	return localctx
}

// IIterationStatementContext is an interface to support dynamic dispatch.
type IIterationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStatementContext differentiates from other interfaces.
	IsIterationStatementContext()
}

type IterationStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStatementContext() *IterationStatementContext {
	var p = new(IterationStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_iterationStatement
	return p
}

func (*IterationStatementContext) IsIterationStatementContext() {}

func NewIterationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStatementContext {
	var p = new(IterationStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_iterationStatement

	return p
}

func (s *IterationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStatementContext) CopyFrom(ctx *IterationStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IterationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DoStatementContext struct {
	*IterationStatementContext
}

func NewDoStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoStatementContext {
	var p = new(DoStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *DoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDo, 0)
}

func (s *DoStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoStatementContext) While() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserWhile, 0)
}

func (s *DoStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *DoStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *DoStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *DoStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DoStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterDoStatement(s)
	}
}

func (s *DoStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitDoStatement(s)
	}
}

type WhileStatementContext struct {
	*IterationStatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserWhile, 0)
}

func (s *WhileStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *WhileStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *WhileStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

type ForStatementContext struct {
	*IterationStatementContext
}

func NewForStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStatementContext {
	var p = new(ForStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFor, 0)
}

func (s *ForStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *ForStatementContext) AllSemiColon() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserSemiColon)
}

func (s *ForStatementContext) SemiColon(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSemiColon, i)
}

func (s *ForStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) AllExpressionSequence() []IExpressionSequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem())
	var tst = make([]IExpressionSequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionSequenceContext)
		}
	}

	return tst
}

func (s *ForStatementContext) ExpressionSequence(i int) IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitForStatement(s)
	}
}

type ForInStatementContext struct {
	*IterationStatementContext
}

func NewForInStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInStatementContext {
	var p = new(ForInStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForInStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStatementContext) For() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFor, 0)
}

func (s *ForInStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *ForInStatementContext) In() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIn, 0)
}

func (s *ForInStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForInStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *ForInStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForInStatementContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ForInStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *ForInStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterForInStatement(s)
	}
}

func (s *ForInStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitForInStatement(s)
	}
}

type ForOfStatementContext struct {
	*IterationStatementContext
}

func NewForOfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForOfStatementContext {
	var p = new(ForOfStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForOfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForOfStatementContext) For() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFor, 0)
}

func (s *ForOfStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *ForOfStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ForOfStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForOfStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *ForOfStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForOfStatementContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ForOfStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *ForOfStatementContext) Await() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAwait, 0)
}

func (s *ForOfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterForOfStatement(s)
	}
}

func (s *ForOfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitForOfStatement(s)
	}
}

func (p *JavaScriptParser) IterationStatement() (localctx IIterationStatementContext) {
	localctx = NewIterationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JavaScriptParserRULE_iterationStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDoStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(JavaScriptParserDo)
		}
		{
			p.SetState(310)
			p.Statement()
		}
		{
			p.SetState(311)
			p.Match(JavaScriptParserWhile)
		}
		{
			p.SetState(312)
			p.Match(JavaScriptParserOpenParen)
		}
		{
			p.SetState(313)
			p.ExpressionSequence()
		}
		{
			p.SetState(314)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(315)
			p.Eos()
		}

	case 2:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.Match(JavaScriptParserWhile)
		}
		{
			p.SetState(318)
			p.Match(JavaScriptParserOpenParen)
		}
		{
			p.SetState(319)
			p.ExpressionSequence()
		}
		{
			p.SetState(320)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(321)
			p.Statement()
		}

	case 3:
		localctx = NewForStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(323)
			p.Match(JavaScriptParserFor)
		}
		{
			p.SetState(324)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(325)
				p.ExpressionSequence()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(326)
				p.VariableDeclarationList()
			}

		}
		{
			p.SetState(329)
			p.Match(JavaScriptParserSemiColon)
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserRegularExpressionLiteral)|(1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenParen)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserPlusPlus)|(1<<JavaScriptParserMinusMinus)|(1<<JavaScriptParserPlus)|(1<<JavaScriptParserMinus)|(1<<JavaScriptParserBitNot)|(1<<JavaScriptParserNot))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(JavaScriptParserNullLiteral-59))|(1<<(JavaScriptParserBooleanLiteral-59))|(1<<(JavaScriptParserDecimalLiteral-59))|(1<<(JavaScriptParserHexIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral2-59))|(1<<(JavaScriptParserBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigHexIntegerLiteral-59))|(1<<(JavaScriptParserBigOctalIntegerLiteral-59))|(1<<(JavaScriptParserBigBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigDecimalIntegerLiteral-59))|(1<<(JavaScriptParserTypeof-59))|(1<<(JavaScriptParserNew-59))|(1<<(JavaScriptParserVoid-59))|(1<<(JavaScriptParserFunction-59))|(1<<(JavaScriptParserThis-59)))) != 0) || (((_la-93)&-(0x1f+1)) == 0 && ((1<<uint((_la-93)))&((1<<(JavaScriptParserDelete-93))|(1<<(JavaScriptParserClass-93))|(1<<(JavaScriptParserSuper-93))|(1<<(JavaScriptParserImport-93))|(1<<(JavaScriptParserAsync-93))|(1<<(JavaScriptParserAwait-93))|(1<<(JavaScriptParserNonStrictLet-93))|(1<<(JavaScriptParserYield-93))|(1<<(JavaScriptParserIdentifier-93))|(1<<(JavaScriptParserStringLiteral-93))|(1<<(JavaScriptParserTemplateStringLiteral-93)))) != 0) {
			{
				p.SetState(330)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(333)
			p.Match(JavaScriptParserSemiColon)
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserRegularExpressionLiteral)|(1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenParen)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserPlusPlus)|(1<<JavaScriptParserMinusMinus)|(1<<JavaScriptParserPlus)|(1<<JavaScriptParserMinus)|(1<<JavaScriptParserBitNot)|(1<<JavaScriptParserNot))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(JavaScriptParserNullLiteral-59))|(1<<(JavaScriptParserBooleanLiteral-59))|(1<<(JavaScriptParserDecimalLiteral-59))|(1<<(JavaScriptParserHexIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral2-59))|(1<<(JavaScriptParserBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigHexIntegerLiteral-59))|(1<<(JavaScriptParserBigOctalIntegerLiteral-59))|(1<<(JavaScriptParserBigBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigDecimalIntegerLiteral-59))|(1<<(JavaScriptParserTypeof-59))|(1<<(JavaScriptParserNew-59))|(1<<(JavaScriptParserVoid-59))|(1<<(JavaScriptParserFunction-59))|(1<<(JavaScriptParserThis-59)))) != 0) || (((_la-93)&-(0x1f+1)) == 0 && ((1<<uint((_la-93)))&((1<<(JavaScriptParserDelete-93))|(1<<(JavaScriptParserClass-93))|(1<<(JavaScriptParserSuper-93))|(1<<(JavaScriptParserImport-93))|(1<<(JavaScriptParserAsync-93))|(1<<(JavaScriptParserAwait-93))|(1<<(JavaScriptParserNonStrictLet-93))|(1<<(JavaScriptParserYield-93))|(1<<(JavaScriptParserIdentifier-93))|(1<<(JavaScriptParserStringLiteral-93))|(1<<(JavaScriptParserTemplateStringLiteral-93)))) != 0) {
			{
				p.SetState(334)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(337)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(338)
			p.Statement()
		}

	case 4:
		localctx = NewForInStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(339)
			p.Match(JavaScriptParserFor)
		}
		{
			p.SetState(340)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(341)
				p.singleExpression(0)
			}

		case 2:
			{
				p.SetState(342)
				p.VariableDeclarationList()
			}

		}
		{
			p.SetState(345)
			p.Match(JavaScriptParserIn)
		}
		{
			p.SetState(346)
			p.ExpressionSequence()
		}
		{
			p.SetState(347)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(348)
			p.Statement()
		}

	case 5:
		localctx = NewForOfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(350)
			p.Match(JavaScriptParserFor)
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserAwait {
			{
				p.SetState(351)
				p.Match(JavaScriptParserAwait)
			}

		}
		{
			p.SetState(354)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(355)
				p.singleExpression(0)
			}

		case 2:
			{
				p.SetState(356)
				p.VariableDeclarationList()
			}

		}
		{
			p.SetState(359)
			p.Identifier()
		}
		p.SetState(360)

		if !(p.p("of")) {
			panic(antlr.NewFailedPredicateException(p, "p.p(\"of\")", ""))
		}
		{
			p.SetState(361)
			p.ExpressionSequence()
		}
		{
			p.SetState(362)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(363)
			p.Statement()
		}

	}

	return localctx
}

// IVarModifierContext is an interface to support dynamic dispatch.
type IVarModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarModifierContext differentiates from other interfaces.
	IsVarModifierContext()
}

type VarModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarModifierContext() *VarModifierContext {
	var p = new(VarModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_varModifier
	return p
}

func (*VarModifierContext) IsVarModifierContext() {}

func NewVarModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarModifierContext {
	var p = new(VarModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_varModifier

	return p
}

func (s *VarModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VarModifierContext) Var() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserVar, 0)
}

func (s *VarModifierContext) Let_() ILet_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILet_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILet_Context)
}

func (s *VarModifierContext) Const() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserConst, 0)
}

func (s *VarModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterVarModifier(s)
	}
}

func (s *VarModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitVarModifier(s)
	}
}

func (p *JavaScriptParser) VarModifier() (localctx IVarModifierContext) {
	localctx = NewVarModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, JavaScriptParserRULE_varModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserVar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(JavaScriptParserVar)
		}

	case JavaScriptParserStrictLet, JavaScriptParserNonStrictLet:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.Let_()
		}

	case JavaScriptParserConst:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(369)
			p.Match(JavaScriptParserConst)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserContinue, 0)
}

func (s *ContinueStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ContinueStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (p *JavaScriptParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, JavaScriptParserRULE_continueStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(JavaScriptParserContinue)
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		p.SetState(373)

		if !(p.notLineTerminator()) {
			panic(antlr.NewFailedPredicateException(p, "p.notLineTerminator()", ""))
		}
		{
			p.SetState(374)
			p.Identifier()
		}

	}
	{
		p.SetState(377)
		p.Eos()
	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBreak, 0)
}

func (s *BreakStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *BreakStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (p *JavaScriptParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, JavaScriptParserRULE_breakStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.Match(JavaScriptParserBreak)
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		p.SetState(380)

		if !(p.notLineTerminator()) {
			panic(antlr.NewFailedPredicateException(p, "p.notLineTerminator()", ""))
		}
		{
			p.SetState(381)
			p.Identifier()
		}

	}
	{
		p.SetState(384)
		p.Eos()
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserReturn, 0)
}

func (s *ReturnStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ReturnStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *JavaScriptParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, JavaScriptParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(JavaScriptParserReturn)
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		p.SetState(387)

		if !(p.notLineTerminator()) {
			panic(antlr.NewFailedPredicateException(p, "p.notLineTerminator()", ""))
		}
		{
			p.SetState(388)
			p.ExpressionSequence()
		}

	}
	{
		p.SetState(391)
		p.Eos()
	}

	return localctx
}

// IYieldStatementContext is an interface to support dynamic dispatch.
type IYieldStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYieldStatementContext differentiates from other interfaces.
	IsYieldStatementContext()
}

type YieldStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYieldStatementContext() *YieldStatementContext {
	var p = new(YieldStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_yieldStatement
	return p
}

func (*YieldStatementContext) IsYieldStatementContext() {}

func NewYieldStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YieldStatementContext {
	var p = new(YieldStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_yieldStatement

	return p
}

func (s *YieldStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *YieldStatementContext) Yield() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserYield, 0)
}

func (s *YieldStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *YieldStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *YieldStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YieldStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YieldStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterYieldStatement(s)
	}
}

func (s *YieldStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitYieldStatement(s)
	}
}

func (p *JavaScriptParser) YieldStatement() (localctx IYieldStatementContext) {
	localctx = NewYieldStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, JavaScriptParserRULE_yieldStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(JavaScriptParserYield)
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		p.SetState(394)

		if !(p.notLineTerminator()) {
			panic(antlr.NewFailedPredicateException(p, "p.notLineTerminator()", ""))
		}
		{
			p.SetState(395)
			p.ExpressionSequence()
		}

	}
	{
		p.SetState(398)
		p.Eos()
	}

	return localctx
}

// IWithStatementContext is an interface to support dynamic dispatch.
type IWithStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithStatementContext differentiates from other interfaces.
	IsWithStatementContext()
}

type WithStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithStatementContext() *WithStatementContext {
	var p = new(WithStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_withStatement
	return p
}

func (*WithStatementContext) IsWithStatementContext() {}

func NewWithStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithStatementContext {
	var p = new(WithStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_withStatement

	return p
}

func (s *WithStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WithStatementContext) With() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserWith, 0)
}

func (s *WithStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *WithStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *WithStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *WithStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WithStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterWithStatement(s)
	}
}

func (s *WithStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitWithStatement(s)
	}
}

func (p *JavaScriptParser) WithStatement() (localctx IWithStatementContext) {
	localctx = NewWithStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, JavaScriptParserRULE_withStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(JavaScriptParserWith)
	}
	{
		p.SetState(401)
		p.Match(JavaScriptParserOpenParen)
	}
	{
		p.SetState(402)
		p.ExpressionSequence()
	}
	{
		p.SetState(403)
		p.Match(JavaScriptParserCloseParen)
	}
	{
		p.SetState(404)
		p.Statement()
	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSwitch, 0)
}

func (s *SwitchStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *SwitchStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *SwitchStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *SwitchStatementContext) CaseBlock() ICaseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseBlockContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *JavaScriptParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, JavaScriptParserRULE_switchStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(JavaScriptParserSwitch)
	}
	{
		p.SetState(407)
		p.Match(JavaScriptParserOpenParen)
	}
	{
		p.SetState(408)
		p.ExpressionSequence()
	}
	{
		p.SetState(409)
		p.Match(JavaScriptParserCloseParen)
	}
	{
		p.SetState(410)
		p.CaseBlock()
	}

	return localctx
}

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_caseBlock
	return p
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBrace, 0)
}

func (s *CaseBlockContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBrace, 0)
}

func (s *CaseBlockContext) AllCaseClauses() []ICaseClausesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseClausesContext)(nil)).Elem())
	var tst = make([]ICaseClausesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseClausesContext)
		}
	}

	return tst
}

func (s *CaseBlockContext) CaseClauses(i int) ICaseClausesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseClausesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseClausesContext)
}

func (s *CaseBlockContext) DefaultClause() IDefaultClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultClauseContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (p *JavaScriptParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, JavaScriptParserRULE_caseBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Match(JavaScriptParserOpenBrace)
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserCase {
		{
			p.SetState(413)
			p.CaseClauses()
		}

	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserDefault {
		{
			p.SetState(416)
			p.DefaultClause()
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserCase {
			{
				p.SetState(417)
				p.CaseClauses()
			}

		}

	}
	{
		p.SetState(422)
		p.Match(JavaScriptParserCloseBrace)
	}

	return localctx
}

// ICaseClausesContext is an interface to support dynamic dispatch.
type ICaseClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseClausesContext differentiates from other interfaces.
	IsCaseClausesContext()
}

type CaseClausesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClausesContext() *CaseClausesContext {
	var p = new(CaseClausesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_caseClauses
	return p
}

func (*CaseClausesContext) IsCaseClausesContext() {}

func NewCaseClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClausesContext {
	var p = new(CaseClausesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_caseClauses

	return p
}

func (s *CaseClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClausesContext) AllCaseClause() []ICaseClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseClauseContext)(nil)).Elem())
	var tst = make([]ICaseClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseClauseContext)
		}
	}

	return tst
}

func (s *CaseClausesContext) CaseClause(i int) ICaseClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseClauseContext)
}

func (s *CaseClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCaseClauses(s)
	}
}

func (s *CaseClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCaseClauses(s)
	}
}

func (p *JavaScriptParser) CaseClauses() (localctx ICaseClausesContext) {
	localctx = NewCaseClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, JavaScriptParserRULE_caseClauses)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JavaScriptParserCase {
		{
			p.SetState(424)
			p.CaseClause()
		}

		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICaseClauseContext is an interface to support dynamic dispatch.
type ICaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseClauseContext differentiates from other interfaces.
	IsCaseClauseContext()
}

type CaseClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClauseContext() *CaseClauseContext {
	var p = new(CaseClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_caseClause
	return p
}

func (*CaseClauseContext) IsCaseClauseContext() {}

func NewCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClauseContext {
	var p = new(CaseClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_caseClause

	return p
}

func (s *CaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClauseContext) Case() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCase, 0)
}

func (s *CaseClauseContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *CaseClauseContext) Colon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserColon, 0)
}

func (s *CaseClauseContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCaseClause(s)
	}
}

func (s *CaseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCaseClause(s)
	}
}

func (p *JavaScriptParser) CaseClause() (localctx ICaseClauseContext) {
	localctx = NewCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, JavaScriptParserRULE_caseClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.Match(JavaScriptParserCase)
	}
	{
		p.SetState(430)
		p.ExpressionSequence()
	}
	{
		p.SetState(431)
		p.Match(JavaScriptParserColon)
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(432)
			p.StatementList()
		}

	}

	return localctx
}

// IDefaultClauseContext is an interface to support dynamic dispatch.
type IDefaultClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultClauseContext differentiates from other interfaces.
	IsDefaultClauseContext()
}

type DefaultClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultClauseContext() *DefaultClauseContext {
	var p = new(DefaultClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_defaultClause
	return p
}

func (*DefaultClauseContext) IsDefaultClauseContext() {}

func NewDefaultClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultClauseContext {
	var p = new(DefaultClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_defaultClause

	return p
}

func (s *DefaultClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultClauseContext) Default() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDefault, 0)
}

func (s *DefaultClauseContext) Colon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserColon, 0)
}

func (s *DefaultClauseContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *DefaultClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterDefaultClause(s)
	}
}

func (s *DefaultClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitDefaultClause(s)
	}
}

func (p *JavaScriptParser) DefaultClause() (localctx IDefaultClauseContext) {
	localctx = NewDefaultClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, JavaScriptParserRULE_defaultClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		p.Match(JavaScriptParserDefault)
	}
	{
		p.SetState(436)
		p.Match(JavaScriptParserColon)
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(437)
			p.StatementList()
		}

	}

	return localctx
}

// ILabelledStatementContext is an interface to support dynamic dispatch.
type ILabelledStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelledStatementContext differentiates from other interfaces.
	IsLabelledStatementContext()
}

type LabelledStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelledStatementContext() *LabelledStatementContext {
	var p = new(LabelledStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_labelledStatement
	return p
}

func (*LabelledStatementContext) IsLabelledStatementContext() {}

func NewLabelledStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelledStatementContext {
	var p = new(LabelledStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_labelledStatement

	return p
}

func (s *LabelledStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelledStatementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LabelledStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserColon, 0)
}

func (s *LabelledStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LabelledStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelledStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelledStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterLabelledStatement(s)
	}
}

func (s *LabelledStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitLabelledStatement(s)
	}
}

func (p *JavaScriptParser) LabelledStatement() (localctx ILabelledStatementContext) {
	localctx = NewLabelledStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, JavaScriptParserRULE_labelledStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.Identifier()
	}
	{
		p.SetState(441)
		p.Match(JavaScriptParserColon)
	}
	{
		p.SetState(442)
		p.Statement()
	}

	return localctx
}

// IThrowStatementContext is an interface to support dynamic dispatch.
type IThrowStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThrowStatementContext differentiates from other interfaces.
	IsThrowStatementContext()
}

type ThrowStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrowStatementContext() *ThrowStatementContext {
	var p = new(ThrowStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_throwStatement
	return p
}

func (*ThrowStatementContext) IsThrowStatementContext() {}

func NewThrowStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowStatementContext {
	var p = new(ThrowStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_throwStatement

	return p
}

func (s *ThrowStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ThrowStatementContext) Throw() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserThrow, 0)
}

func (s *ThrowStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ThrowStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ThrowStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThrowStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterThrowStatement(s)
	}
}

func (s *ThrowStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitThrowStatement(s)
	}
}

func (p *JavaScriptParser) ThrowStatement() (localctx IThrowStatementContext) {
	localctx = NewThrowStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, JavaScriptParserRULE_throwStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(JavaScriptParserThrow)
	}
	p.SetState(445)

	if !(p.notLineTerminator()) {
		panic(antlr.NewFailedPredicateException(p, "p.notLineTerminator()", ""))
	}
	{
		p.SetState(446)
		p.ExpressionSequence()
	}
	{
		p.SetState(447)
		p.Eos()
	}

	return localctx
}

// ITryStatementContext is an interface to support dynamic dispatch.
type ITryStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTryStatementContext differentiates from other interfaces.
	IsTryStatementContext()
}

type TryStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryStatementContext() *TryStatementContext {
	var p = new(TryStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_tryStatement
	return p
}

func (*TryStatementContext) IsTryStatementContext() {}

func NewTryStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryStatementContext {
	var p = new(TryStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_tryStatement

	return p
}

func (s *TryStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TryStatementContext) Try() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTry, 0)
}

func (s *TryStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TryStatementContext) CatchProduction() ICatchProductionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICatchProductionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICatchProductionContext)
}

func (s *TryStatementContext) FinallyProduction() IFinallyProductionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFinallyProductionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFinallyProductionContext)
}

func (s *TryStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterTryStatement(s)
	}
}

func (s *TryStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitTryStatement(s)
	}
}

func (p *JavaScriptParser) TryStatement() (localctx ITryStatementContext) {
	localctx = NewTryStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, JavaScriptParserRULE_tryStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Match(JavaScriptParserTry)
	}
	{
		p.SetState(450)
		p.Block()
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserCatch:
		{
			p.SetState(451)
			p.CatchProduction()
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(452)
				p.FinallyProduction()
			}

		}

	case JavaScriptParserFinally:
		{
			p.SetState(455)
			p.FinallyProduction()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICatchProductionContext is an interface to support dynamic dispatch.
type ICatchProductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCatchProductionContext differentiates from other interfaces.
	IsCatchProductionContext()
}

type CatchProductionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatchProductionContext() *CatchProductionContext {
	var p = new(CatchProductionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_catchProduction
	return p
}

func (*CatchProductionContext) IsCatchProductionContext() {}

func NewCatchProductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchProductionContext {
	var p = new(CatchProductionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_catchProduction

	return p
}

func (s *CatchProductionContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchProductionContext) Catch() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCatch, 0)
}

func (s *CatchProductionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CatchProductionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *CatchProductionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *CatchProductionContext) Assignable() IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *CatchProductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchProductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchProductionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCatchProduction(s)
	}
}

func (s *CatchProductionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCatchProduction(s)
	}
}

func (p *JavaScriptParser) CatchProduction() (localctx ICatchProductionContext) {
	localctx = NewCatchProductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, JavaScriptParserRULE_catchProduction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.Match(JavaScriptParserCatch)
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserOpenParen {
		{
			p.SetState(459)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserOpenBracket || _la == JavaScriptParserOpenBrace || (((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
			{
				p.SetState(460)
				p.Assignable()
			}

		}
		{
			p.SetState(463)
			p.Match(JavaScriptParserCloseParen)
		}

	}
	{
		p.SetState(466)
		p.Block()
	}

	return localctx
}

// IFinallyProductionContext is an interface to support dynamic dispatch.
type IFinallyProductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFinallyProductionContext differentiates from other interfaces.
	IsFinallyProductionContext()
}

type FinallyProductionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinallyProductionContext() *FinallyProductionContext {
	var p = new(FinallyProductionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_finallyProduction
	return p
}

func (*FinallyProductionContext) IsFinallyProductionContext() {}

func NewFinallyProductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinallyProductionContext {
	var p = new(FinallyProductionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_finallyProduction

	return p
}

func (s *FinallyProductionContext) GetParser() antlr.Parser { return s.parser }

func (s *FinallyProductionContext) Finally() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFinally, 0)
}

func (s *FinallyProductionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FinallyProductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FinallyProductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FinallyProductionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFinallyProduction(s)
	}
}

func (s *FinallyProductionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFinallyProduction(s)
	}
}

func (p *JavaScriptParser) FinallyProduction() (localctx IFinallyProductionContext) {
	localctx = NewFinallyProductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, JavaScriptParserRULE_finallyProduction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.Match(JavaScriptParserFinally)
	}
	{
		p.SetState(469)
		p.Block()
	}

	return localctx
}

// IDebuggerStatementContext is an interface to support dynamic dispatch.
type IDebuggerStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDebuggerStatementContext differentiates from other interfaces.
	IsDebuggerStatementContext()
}

type DebuggerStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebuggerStatementContext() *DebuggerStatementContext {
	var p = new(DebuggerStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_debuggerStatement
	return p
}

func (*DebuggerStatementContext) IsDebuggerStatementContext() {}

func NewDebuggerStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebuggerStatementContext {
	var p = new(DebuggerStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_debuggerStatement

	return p
}

func (s *DebuggerStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DebuggerStatementContext) Debugger() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDebugger, 0)
}

func (s *DebuggerStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DebuggerStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DebuggerStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DebuggerStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterDebuggerStatement(s)
	}
}

func (s *DebuggerStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitDebuggerStatement(s)
	}
}

func (p *JavaScriptParser) DebuggerStatement() (localctx IDebuggerStatementContext) {
	localctx = NewDebuggerStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, JavaScriptParserRULE_debuggerStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Match(JavaScriptParserDebugger)
	}
	{
		p.SetState(472)
		p.Eos()
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Function() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFunction, 0)
}

func (s *FunctionDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionDeclarationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *FunctionDeclarationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) Async() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAsync, 0)
}

func (s *FunctionDeclarationContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMultiply, 0)
}

func (s *FunctionDeclarationContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *JavaScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, JavaScriptParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserAsync {
		{
			p.SetState(474)
			p.Match(JavaScriptParserAsync)
		}

	}
	{
		p.SetState(477)
		p.Match(JavaScriptParserFunction)
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserMultiply {
		{
			p.SetState(478)
			p.Match(JavaScriptParserMultiply)
		}

	}
	{
		p.SetState(481)
		p.Identifier()
	}
	{
		p.SetState(482)
		p.Match(JavaScriptParserOpenParen)
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis))) != 0) || (((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
		{
			p.SetState(483)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(486)
		p.Match(JavaScriptParserCloseParen)
	}
	{
		p.SetState(487)
		p.FunctionBody()
	}

	return localctx
}

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) Class() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserClass, 0)
}

func (s *ClassDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ClassDeclarationContext) ClassTail() IClassTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassTailContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (p *JavaScriptParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, JavaScriptParserRULE_classDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Match(JavaScriptParserClass)
	}
	{
		p.SetState(490)
		p.Identifier()
	}
	{
		p.SetState(491)
		p.ClassTail()
	}

	return localctx
}

// IClassTailContext is an interface to support dynamic dispatch.
type IClassTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassTailContext differentiates from other interfaces.
	IsClassTailContext()
}

type ClassTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassTailContext() *ClassTailContext {
	var p = new(ClassTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_classTail
	return p
}

func (*ClassTailContext) IsClassTailContext() {}

func NewClassTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassTailContext {
	var p = new(ClassTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_classTail

	return p
}

func (s *ClassTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassTailContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBrace, 0)
}

func (s *ClassTailContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBrace, 0)
}

func (s *ClassTailContext) Extends() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserExtends, 0)
}

func (s *ClassTailContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ClassTailContext) AllClassElement() []IClassElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassElementContext)(nil)).Elem())
	var tst = make([]IClassElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassElementContext)
		}
	}

	return tst
}

func (s *ClassTailContext) ClassElement(i int) IClassElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassElementContext)
}

func (s *ClassTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterClassTail(s)
	}
}

func (s *ClassTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitClassTail(s)
	}
}

func (p *JavaScriptParser) ClassTail() (localctx IClassTailContext) {
	localctx = NewClassTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, JavaScriptParserRULE_classTail)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserExtends {
		{
			p.SetState(493)
			p.Match(JavaScriptParserExtends)
		}
		{
			p.SetState(494)
			p.singleExpression(0)
		}

	}
	{
		p.SetState(497)
		p.Match(JavaScriptParserOpenBrace)
	}
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(498)
				p.ClassElement()
			}

		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}
	{
		p.SetState(504)
		p.Match(JavaScriptParserCloseBrace)
	}

	return localctx
}

// IClassElementContext is an interface to support dynamic dispatch.
type IClassElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassElementContext differentiates from other interfaces.
	IsClassElementContext()
}

type ClassElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassElementContext() *ClassElementContext {
	var p = new(ClassElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_classElement
	return p
}

func (*ClassElementContext) IsClassElementContext() {}

func NewClassElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassElementContext {
	var p = new(ClassElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_classElement

	return p
}

func (s *ClassElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassElementContext) MethodDefinition() IMethodDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodDefinitionContext)
}

func (s *ClassElementContext) Assignable() IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *ClassElementContext) Assign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAssign, 0)
}

func (s *ClassElementContext) ObjectLiteral() IObjectLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ClassElementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSemiColon, 0)
}

func (s *ClassElementContext) AllStatic() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserStatic)
}

func (s *ClassElementContext) Static(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserStatic, i)
}

func (s *ClassElementContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *ClassElementContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ClassElementContext) AllAsync() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserAsync)
}

func (s *ClassElementContext) Async(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAsync, i)
}

func (s *ClassElementContext) EmptyStatement() IEmptyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *ClassElementContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *ClassElementContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ClassElementContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserHashtag, 0)
}

func (s *ClassElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterClassElement(s)
	}
}

func (s *ClassElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitClassElement(s)
	}
}

func (p *JavaScriptParser) ClassElement() (localctx IClassElementContext) {
	localctx = NewClassElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, JavaScriptParserRULE_classElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(510)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(506)
						p.Match(JavaScriptParserStatic)
					}

				case 2:
					p.SetState(507)

					if !(p.n("static")) {
						panic(antlr.NewFailedPredicateException(p, "p.n(\"static\")", ""))
					}
					{
						p.SetState(508)
						p.Identifier()
					}

				case 3:
					{
						p.SetState(509)
						p.Match(JavaScriptParserAsync)
					}

				}

			}
			p.SetState(514)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
		}
		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(515)
				p.MethodDefinition()
			}

		case 2:
			{
				p.SetState(516)
				p.Assignable()
			}
			{
				p.SetState(517)
				p.Match(JavaScriptParserAssign)
			}
			{
				p.SetState(518)
				p.ObjectLiteral()
			}
			{
				p.SetState(519)
				p.Match(JavaScriptParserSemiColon)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(523)
			p.EmptyStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(525)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserHashtag {
			{
				p.SetState(524)
				p.Match(JavaScriptParserHashtag)
			}

		}
		{
			p.SetState(527)
			p.PropertyName()
		}
		{
			p.SetState(528)
			p.Match(JavaScriptParserAssign)
		}
		{
			p.SetState(529)
			p.singleExpression(0)
		}

	}

	return localctx
}

// IMethodDefinitionContext is an interface to support dynamic dispatch.
type IMethodDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodDefinitionContext differentiates from other interfaces.
	IsMethodDefinitionContext()
}

type MethodDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodDefinitionContext() *MethodDefinitionContext {
	var p = new(MethodDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_methodDefinition
	return p
}

func (*MethodDefinitionContext) IsMethodDefinitionContext() {}

func NewMethodDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDefinitionContext {
	var p = new(MethodDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_methodDefinition

	return p
}

func (s *MethodDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDefinitionContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *MethodDefinitionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *MethodDefinitionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *MethodDefinitionContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *MethodDefinitionContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMultiply, 0)
}

func (s *MethodDefinitionContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserHashtag, 0)
}

func (s *MethodDefinitionContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *MethodDefinitionContext) Getter() IGetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetterContext)
}

func (s *MethodDefinitionContext) Setter() ISetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetterContext)
}

func (s *MethodDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterMethodDefinition(s)
	}
}

func (s *MethodDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitMethodDefinition(s)
	}
}

func (p *JavaScriptParser) MethodDefinition() (localctx IMethodDefinitionContext) {
	localctx = NewMethodDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, JavaScriptParserRULE_methodDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(534)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserMultiply {
			{
				p.SetState(533)
				p.Match(JavaScriptParserMultiply)
			}

		}
		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserHashtag {
			{
				p.SetState(536)
				p.Match(JavaScriptParserHashtag)
			}

		}
		{
			p.SetState(539)
			p.PropertyName()
		}
		{
			p.SetState(540)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis))) != 0) || (((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
			{
				p.SetState(541)
				p.FormalParameterList()
			}

		}
		{
			p.SetState(544)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(545)
			p.FunctionBody()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(548)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(547)
				p.Match(JavaScriptParserMultiply)
			}

		}
		p.SetState(551)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(550)
				p.Match(JavaScriptParserHashtag)
			}

		}
		{
			p.SetState(553)
			p.Getter()
		}
		{
			p.SetState(554)
			p.Match(JavaScriptParserOpenParen)
		}
		{
			p.SetState(555)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(556)
			p.FunctionBody()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(559)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(558)
				p.Match(JavaScriptParserMultiply)
			}

		}
		p.SetState(562)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(561)
				p.Match(JavaScriptParserHashtag)
			}

		}
		{
			p.SetState(564)
			p.Setter()
		}
		{
			p.SetState(565)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(567)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis))) != 0) || (((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
			{
				p.SetState(566)
				p.FormalParameterList()
			}

		}
		{
			p.SetState(569)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(570)
			p.FunctionBody()
		}

	}

	return localctx
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_formalParameterList
	return p
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllFormalParameterArg() []IFormalParameterArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormalParameterArgContext)(nil)).Elem())
	var tst = make([]IFormalParameterArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormalParameterArgContext)
		}
	}

	return tst
}

func (s *FormalParameterListContext) FormalParameterArg(i int) IFormalParameterArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterArgContext)
}

func (s *FormalParameterListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserComma)
}

func (s *FormalParameterListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, i)
}

func (s *FormalParameterListContext) LastFormalParameterArg() ILastFormalParameterArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILastFormalParameterArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILastFormalParameterArgContext)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (p *JavaScriptParser) FormalParameterList() (localctx IFormalParameterListContext) {
	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, JavaScriptParserRULE_formalParameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(587)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserOpenBracket, JavaScriptParserOpenBrace, JavaScriptParserAsync, JavaScriptParserNonStrictLet, JavaScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(574)
			p.FormalParameterArg()
		}
		p.SetState(579)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(575)
					p.Match(JavaScriptParserComma)
				}
				{
					p.SetState(576)
					p.FormalParameterArg()
				}

			}
			p.SetState(581)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())
		}
		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserComma {
			{
				p.SetState(582)
				p.Match(JavaScriptParserComma)
			}
			{
				p.SetState(583)
				p.LastFormalParameterArg()
			}

		}

	case JavaScriptParserEllipsis:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(586)
			p.LastFormalParameterArg()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormalParameterArgContext is an interface to support dynamic dispatch.
type IFormalParameterArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterArgContext differentiates from other interfaces.
	IsFormalParameterArgContext()
}

type FormalParameterArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterArgContext() *FormalParameterArgContext {
	var p = new(FormalParameterArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_formalParameterArg
	return p
}

func (*FormalParameterArgContext) IsFormalParameterArgContext() {}

func NewFormalParameterArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterArgContext {
	var p = new(FormalParameterArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_formalParameterArg

	return p
}

func (s *FormalParameterArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterArgContext) Assignable() IAssignableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignableContext)
}

func (s *FormalParameterArgContext) Assign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAssign, 0)
}

func (s *FormalParameterArgContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *FormalParameterArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFormalParameterArg(s)
	}
}

func (s *FormalParameterArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFormalParameterArg(s)
	}
}

func (p *JavaScriptParser) FormalParameterArg() (localctx IFormalParameterArgContext) {
	localctx = NewFormalParameterArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, JavaScriptParserRULE_formalParameterArg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		p.Assignable()
	}
	p.SetState(592)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserAssign {
		{
			p.SetState(590)
			p.Match(JavaScriptParserAssign)
		}
		{
			p.SetState(591)
			p.singleExpression(0)
		}

	}

	return localctx
}

// ILastFormalParameterArgContext is an interface to support dynamic dispatch.
type ILastFormalParameterArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLastFormalParameterArgContext differentiates from other interfaces.
	IsLastFormalParameterArgContext()
}

type LastFormalParameterArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastFormalParameterArgContext() *LastFormalParameterArgContext {
	var p = new(LastFormalParameterArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_lastFormalParameterArg
	return p
}

func (*LastFormalParameterArgContext) IsLastFormalParameterArgContext() {}

func NewLastFormalParameterArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastFormalParameterArgContext {
	var p = new(LastFormalParameterArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_lastFormalParameterArg

	return p
}

func (s *LastFormalParameterArgContext) GetParser() antlr.Parser { return s.parser }

func (s *LastFormalParameterArgContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEllipsis, 0)
}

func (s *LastFormalParameterArgContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *LastFormalParameterArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastFormalParameterArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LastFormalParameterArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterLastFormalParameterArg(s)
	}
}

func (s *LastFormalParameterArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitLastFormalParameterArg(s)
	}
}

func (p *JavaScriptParser) LastFormalParameterArg() (localctx ILastFormalParameterArgContext) {
	localctx = NewLastFormalParameterArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, JavaScriptParserRULE_lastFormalParameterArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)
		p.Match(JavaScriptParserEllipsis)
	}
	{
		p.SetState(595)
		p.singleExpression(0)
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBrace, 0)
}

func (s *FunctionBodyContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBrace, 0)
}

func (s *FunctionBodyContext) SourceElements() ISourceElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceElementsContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (p *JavaScriptParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, JavaScriptParserRULE_functionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(597)
		p.Match(JavaScriptParserOpenBrace)
	}
	p.SetState(599)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(598)
			p.SourceElements()
		}

	}
	{
		p.SetState(601)
		p.Match(JavaScriptParserCloseBrace)
	}

	return localctx
}

// ISourceElementsContext is an interface to support dynamic dispatch.
type ISourceElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementsContext differentiates from other interfaces.
	IsSourceElementsContext()
}

type SourceElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementsContext() *SourceElementsContext {
	var p = new(SourceElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_sourceElements
	return p
}

func (*SourceElementsContext) IsSourceElementsContext() {}

func NewSourceElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementsContext {
	var p = new(SourceElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_sourceElements

	return p
}

func (s *SourceElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementsContext) AllSourceElement() []ISourceElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceElementContext)(nil)).Elem())
	var tst = make([]ISourceElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceElementContext)
		}
	}

	return tst
}

func (s *SourceElementsContext) SourceElement(i int) ISourceElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceElementContext)
}

func (s *SourceElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterSourceElements(s)
	}
}

func (s *SourceElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitSourceElements(s)
	}
}

func (p *JavaScriptParser) SourceElements() (localctx ISourceElementsContext) {
	localctx = NewSourceElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, JavaScriptParserRULE_sourceElements)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(604)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(603)
				p.SourceElement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(606)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBracket, 0)
}

func (s *ArrayLiteralContext) ElementList() IElementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementListContext)
}

func (s *ArrayLiteralContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBracket, 0)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *JavaScriptParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, JavaScriptParserRULE_arrayLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(608)
		p.Match(JavaScriptParserOpenBracket)
	}
	{
		p.SetState(609)
		p.ElementList()
	}
	{
		p.SetState(610)
		p.Match(JavaScriptParserCloseBracket)
	}

	return localctx
}

// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_elementList
	return p
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_elementList

	return p
}

func (s *ElementListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserComma)
}

func (s *ElementListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, i)
}

func (s *ElementListContext) AllArrayElement() []IArrayElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArrayElementContext)(nil)).Elem())
	var tst = make([]IArrayElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArrayElementContext)
		}
	}

	return tst
}

func (s *ElementListContext) ArrayElement(i int) IArrayElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArrayElementContext)
}

func (s *ElementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterElementList(s)
	}
}

func (s *ElementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitElementList(s)
	}
}

func (p *JavaScriptParser) ElementList() (localctx IElementListContext) {
	localctx = NewElementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, JavaScriptParserRULE_elementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(612)
				p.Match(JavaScriptParserComma)
			}

		}
		p.SetState(617)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext())
	}
	p.SetState(619)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserRegularExpressionLiteral)|(1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenParen)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis)|(1<<JavaScriptParserPlusPlus)|(1<<JavaScriptParserMinusMinus)|(1<<JavaScriptParserPlus)|(1<<JavaScriptParserMinus)|(1<<JavaScriptParserBitNot)|(1<<JavaScriptParserNot))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(JavaScriptParserNullLiteral-59))|(1<<(JavaScriptParserBooleanLiteral-59))|(1<<(JavaScriptParserDecimalLiteral-59))|(1<<(JavaScriptParserHexIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral2-59))|(1<<(JavaScriptParserBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigHexIntegerLiteral-59))|(1<<(JavaScriptParserBigOctalIntegerLiteral-59))|(1<<(JavaScriptParserBigBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigDecimalIntegerLiteral-59))|(1<<(JavaScriptParserTypeof-59))|(1<<(JavaScriptParserNew-59))|(1<<(JavaScriptParserVoid-59))|(1<<(JavaScriptParserFunction-59))|(1<<(JavaScriptParserThis-59)))) != 0) || (((_la-93)&-(0x1f+1)) == 0 && ((1<<uint((_la-93)))&((1<<(JavaScriptParserDelete-93))|(1<<(JavaScriptParserClass-93))|(1<<(JavaScriptParserSuper-93))|(1<<(JavaScriptParserImport-93))|(1<<(JavaScriptParserAsync-93))|(1<<(JavaScriptParserAwait-93))|(1<<(JavaScriptParserNonStrictLet-93))|(1<<(JavaScriptParserYield-93))|(1<<(JavaScriptParserIdentifier-93))|(1<<(JavaScriptParserStringLiteral-93))|(1<<(JavaScriptParserTemplateStringLiteral-93)))) != 0) {
		{
			p.SetState(618)
			p.ArrayElement()
		}

	}
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(622)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == JavaScriptParserComma {
				{
					p.SetState(621)
					p.Match(JavaScriptParserComma)
				}

				p.SetState(624)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(626)
				p.ArrayElement()
			}

		}
		p.SetState(631)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())
	}
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JavaScriptParserComma {
		{
			p.SetState(632)
			p.Match(JavaScriptParserComma)
		}

		p.SetState(637)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrayElementContext is an interface to support dynamic dispatch.
type IArrayElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayElementContext differentiates from other interfaces.
	IsArrayElementContext()
}

type ArrayElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayElementContext() *ArrayElementContext {
	var p = new(ArrayElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_arrayElement
	return p
}

func (*ArrayElementContext) IsArrayElementContext() {}

func NewArrayElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayElementContext {
	var p = new(ArrayElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_arrayElement

	return p
}

func (s *ArrayElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayElementContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArrayElementContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEllipsis, 0)
}

func (s *ArrayElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArrayElement(s)
	}
}

func (s *ArrayElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArrayElement(s)
	}
}

func (p *JavaScriptParser) ArrayElement() (localctx IArrayElementContext) {
	localctx = NewArrayElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, JavaScriptParserRULE_arrayElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(639)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserEllipsis {
		{
			p.SetState(638)
			p.Match(JavaScriptParserEllipsis)
		}

	}
	{
		p.SetState(641)
		p.singleExpression(0)
	}

	return localctx
}

// IPropertyAssignmentContext is an interface to support dynamic dispatch.
type IPropertyAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyAssignmentContext differentiates from other interfaces.
	IsPropertyAssignmentContext()
}

type PropertyAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAssignmentContext() *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_propertyAssignment
	return p
}

func (*PropertyAssignmentContext) IsPropertyAssignmentContext() {}

func NewPropertyAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_propertyAssignment

	return p
}

func (s *PropertyAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAssignmentContext) CopyFrom(ctx *PropertyAssignmentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PropertyAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropertyExpressionAssignmentContext struct {
	*PropertyAssignmentContext
}

func NewPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyExpressionAssignmentContext {
	var p = new(PropertyExpressionAssignmentContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyExpressionAssignmentContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyExpressionAssignmentContext) Colon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserColon, 0)
}

func (s *PropertyExpressionAssignmentContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PropertyExpressionAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPropertyExpressionAssignment(s)
	}
}

func (s *PropertyExpressionAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPropertyExpressionAssignment(s)
	}
}

type ComputedPropertyExpressionAssignmentContext struct {
	*PropertyAssignmentContext
}

func NewComputedPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedPropertyExpressionAssignmentContext {
	var p = new(ComputedPropertyExpressionAssignmentContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *ComputedPropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedPropertyExpressionAssignmentContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBracket, 0)
}

func (s *ComputedPropertyExpressionAssignmentContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ComputedPropertyExpressionAssignmentContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ComputedPropertyExpressionAssignmentContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBracket, 0)
}

func (s *ComputedPropertyExpressionAssignmentContext) Colon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserColon, 0)
}

func (s *ComputedPropertyExpressionAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterComputedPropertyExpressionAssignment(s)
	}
}

func (s *ComputedPropertyExpressionAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitComputedPropertyExpressionAssignment(s)
	}
}

type PropertyShorthandContext struct {
	*PropertyAssignmentContext
}

func NewPropertyShorthandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyShorthandContext {
	var p = new(PropertyShorthandContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyShorthandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyShorthandContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PropertyShorthandContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEllipsis, 0)
}

func (s *PropertyShorthandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPropertyShorthand(s)
	}
}

func (s *PropertyShorthandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPropertyShorthand(s)
	}
}

type PropertySetterContext struct {
	*PropertyAssignmentContext
}

func NewPropertySetterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertySetterContext {
	var p = new(PropertySetterContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertySetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertySetterContext) Setter() ISetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetterContext)
}

func (s *PropertySetterContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *PropertySetterContext) FormalParameterArg() IFormalParameterArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterArgContext)
}

func (s *PropertySetterContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *PropertySetterContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *PropertySetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPropertySetter(s)
	}
}

func (s *PropertySetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPropertySetter(s)
	}
}

type PropertyGetterContext struct {
	*PropertyAssignmentContext
}

func NewPropertyGetterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyGetterContext {
	var p = new(PropertyGetterContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyGetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyGetterContext) Getter() IGetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetterContext)
}

func (s *PropertyGetterContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *PropertyGetterContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *PropertyGetterContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *PropertyGetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPropertyGetter(s)
	}
}

func (s *PropertyGetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPropertyGetter(s)
	}
}

type FunctionPropertyContext struct {
	*PropertyAssignmentContext
}

func NewFunctionPropertyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionPropertyContext {
	var p = new(FunctionPropertyContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *FunctionPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionPropertyContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *FunctionPropertyContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *FunctionPropertyContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *FunctionPropertyContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionPropertyContext) Async() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAsync, 0)
}

func (s *FunctionPropertyContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMultiply, 0)
}

func (s *FunctionPropertyContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FunctionPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFunctionProperty(s)
	}
}

func (s *FunctionPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFunctionProperty(s)
	}
}

func (p *JavaScriptParser) PropertyAssignment() (localctx IPropertyAssignmentContext) {
	localctx = NewPropertyAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, JavaScriptParserRULE_propertyAssignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(682)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(643)
			p.PropertyName()
		}
		{
			p.SetState(644)
			p.Match(JavaScriptParserColon)
		}
		{
			p.SetState(645)
			p.singleExpression(0)
		}

	case 2:
		localctx = NewComputedPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(647)
			p.Match(JavaScriptParserOpenBracket)
		}
		{
			p.SetState(648)
			p.singleExpression(0)
		}
		{
			p.SetState(649)
			p.Match(JavaScriptParserCloseBracket)
		}
		{
			p.SetState(650)
			p.Match(JavaScriptParserColon)
		}
		{
			p.SetState(651)
			p.singleExpression(0)
		}

	case 3:
		localctx = NewFunctionPropertyContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(654)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 75, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(653)
				p.Match(JavaScriptParserAsync)
			}

		}
		p.SetState(657)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserMultiply {
			{
				p.SetState(656)
				p.Match(JavaScriptParserMultiply)
			}

		}
		{
			p.SetState(659)
			p.PropertyName()
		}
		{
			p.SetState(660)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(662)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis))) != 0) || (((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
			{
				p.SetState(661)
				p.FormalParameterList()
			}

		}
		{
			p.SetState(664)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(665)
			p.FunctionBody()
		}

	case 4:
		localctx = NewPropertyGetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(667)
			p.Getter()
		}
		{
			p.SetState(668)
			p.Match(JavaScriptParserOpenParen)
		}
		{
			p.SetState(669)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(670)
			p.FunctionBody()
		}

	case 5:
		localctx = NewPropertySetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(672)
			p.Setter()
		}
		{
			p.SetState(673)
			p.Match(JavaScriptParserOpenParen)
		}
		{
			p.SetState(674)
			p.FormalParameterArg()
		}
		{
			p.SetState(675)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(676)
			p.FunctionBody()
		}

	case 6:
		localctx = NewPropertyShorthandContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(679)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserEllipsis {
			{
				p.SetState(678)
				p.Match(JavaScriptParserEllipsis)
			}

		}
		{
			p.SetState(681)
			p.singleExpression(0)
		}

	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *PropertyNameContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserStringLiteral, 0)
}

func (s *PropertyNameContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *PropertyNameContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBracket, 0)
}

func (s *PropertyNameContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PropertyNameContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBracket, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *JavaScriptParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, JavaScriptParserRULE_propertyName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(691)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserNullLiteral, JavaScriptParserBooleanLiteral, JavaScriptParserBreak, JavaScriptParserDo, JavaScriptParserInstanceof, JavaScriptParserTypeof, JavaScriptParserCase, JavaScriptParserElse, JavaScriptParserNew, JavaScriptParserVar, JavaScriptParserCatch, JavaScriptParserFinally, JavaScriptParserReturn, JavaScriptParserVoid, JavaScriptParserContinue, JavaScriptParserFor, JavaScriptParserSwitch, JavaScriptParserWhile, JavaScriptParserDebugger, JavaScriptParserFunction, JavaScriptParserThis, JavaScriptParserWith, JavaScriptParserDefault, JavaScriptParserIf, JavaScriptParserThrow, JavaScriptParserDelete, JavaScriptParserIn, JavaScriptParserTry, JavaScriptParserAs, JavaScriptParserFrom, JavaScriptParserClass, JavaScriptParserEnum, JavaScriptParserExtends, JavaScriptParserSuper, JavaScriptParserConst, JavaScriptParserExport, JavaScriptParserImport, JavaScriptParserAsync, JavaScriptParserAwait, JavaScriptParserImplements, JavaScriptParserStrictLet, JavaScriptParserNonStrictLet, JavaScriptParserPrivate, JavaScriptParserPublic, JavaScriptParserInterface, JavaScriptParserPackage, JavaScriptParserProtected, JavaScriptParserStatic, JavaScriptParserYield, JavaScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(684)
			p.IdentifierName()
		}

	case JavaScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(685)
			p.Match(JavaScriptParserStringLiteral)
		}

	case JavaScriptParserDecimalLiteral, JavaScriptParserHexIntegerLiteral, JavaScriptParserOctalIntegerLiteral, JavaScriptParserOctalIntegerLiteral2, JavaScriptParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(686)
			p.NumericLiteral()
		}

	case JavaScriptParserOpenBracket:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(687)
			p.Match(JavaScriptParserOpenBracket)
		}
		{
			p.SetState(688)
			p.singleExpression(0)
		}
		{
			p.SetState(689)
			p.Match(JavaScriptParserCloseBracket)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *ArgumentsContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserComma)
}

func (s *ArgumentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *JavaScriptParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, JavaScriptParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(693)
		p.Match(JavaScriptParserOpenParen)
	}
	p.SetState(705)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserRegularExpressionLiteral)|(1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenParen)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis)|(1<<JavaScriptParserPlusPlus)|(1<<JavaScriptParserMinusMinus)|(1<<JavaScriptParserPlus)|(1<<JavaScriptParserMinus)|(1<<JavaScriptParserBitNot)|(1<<JavaScriptParserNot))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(JavaScriptParserNullLiteral-59))|(1<<(JavaScriptParserBooleanLiteral-59))|(1<<(JavaScriptParserDecimalLiteral-59))|(1<<(JavaScriptParserHexIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral-59))|(1<<(JavaScriptParserOctalIntegerLiteral2-59))|(1<<(JavaScriptParserBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigHexIntegerLiteral-59))|(1<<(JavaScriptParserBigOctalIntegerLiteral-59))|(1<<(JavaScriptParserBigBinaryIntegerLiteral-59))|(1<<(JavaScriptParserBigDecimalIntegerLiteral-59))|(1<<(JavaScriptParserTypeof-59))|(1<<(JavaScriptParserNew-59))|(1<<(JavaScriptParserVoid-59))|(1<<(JavaScriptParserFunction-59))|(1<<(JavaScriptParserThis-59)))) != 0) || (((_la-93)&-(0x1f+1)) == 0 && ((1<<uint((_la-93)))&((1<<(JavaScriptParserDelete-93))|(1<<(JavaScriptParserClass-93))|(1<<(JavaScriptParserSuper-93))|(1<<(JavaScriptParserImport-93))|(1<<(JavaScriptParserAsync-93))|(1<<(JavaScriptParserAwait-93))|(1<<(JavaScriptParserNonStrictLet-93))|(1<<(JavaScriptParserYield-93))|(1<<(JavaScriptParserIdentifier-93))|(1<<(JavaScriptParserStringLiteral-93))|(1<<(JavaScriptParserTemplateStringLiteral-93)))) != 0) {
		{
			p.SetState(694)
			p.Argument()
		}
		p.SetState(699)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(695)
					p.Match(JavaScriptParserComma)
				}
				{
					p.SetState(696)
					p.Argument()
				}

			}
			p.SetState(701)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext())
		}
		p.SetState(703)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserComma {
			{
				p.SetState(702)
				p.Match(JavaScriptParserComma)
			}

		}

	}
	{
		p.SetState(707)
		p.Match(JavaScriptParserCloseParen)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArgumentContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEllipsis, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *JavaScriptParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, JavaScriptParserRULE_argument)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(710)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserEllipsis {
		{
			p.SetState(709)
			p.Match(JavaScriptParserEllipsis)
		}

	}
	p.SetState(714)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 85, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(712)
			p.singleExpression(0)
		}

	case 2:
		{
			p.SetState(713)
			p.Identifier()
		}

	}

	return localctx
}

// IExpressionSequenceContext is an interface to support dynamic dispatch.
type IExpressionSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionSequenceContext differentiates from other interfaces.
	IsExpressionSequenceContext()
}

type ExpressionSequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSequenceContext() *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_expressionSequence
	return p
}

func (*ExpressionSequenceContext) IsExpressionSequenceContext() {}

func NewExpressionSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_expressionSequence

	return p
}

func (s *ExpressionSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSequenceContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionSequenceContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ExpressionSequenceContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserComma)
}

func (s *ExpressionSequenceContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, i)
}

func (s *ExpressionSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterExpressionSequence(s)
	}
}

func (s *ExpressionSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitExpressionSequence(s)
	}
}

func (p *JavaScriptParser) ExpressionSequence() (localctx IExpressionSequenceContext) {
	localctx = NewExpressionSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, JavaScriptParserRULE_expressionSequence)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(716)
		p.singleExpression(0)
	}
	p.SetState(721)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 86, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(717)
				p.Match(JavaScriptParserComma)
			}
			{
				p.SetState(718)
				p.singleExpression(0)
			}

		}
		p.SetState(723)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 86, p.GetParserRuleContext())
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) CopyFrom(ctx *SingleExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TemplateStringExpressionContext struct {
	*SingleExpressionContext
}

func NewTemplateStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TemplateStringExpressionContext {
	var p = new(TemplateStringExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TemplateStringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateStringExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *TemplateStringExpressionContext) TemplateStringLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTemplateStringLiteral, 0)
}

func (s *TemplateStringExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterTemplateStringExpression(s)
	}
}

func (s *TemplateStringExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitTemplateStringExpression(s)
	}
}

type TernaryExpressionContext struct {
	*SingleExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *TernaryExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserQuestionMark, 0)
}

func (s *TernaryExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserColon, 0)
}

func (s *TernaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitTernaryExpression(s)
	}
}

type LogicalAndExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *LogicalAndExpressionContext) And() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAnd, 0)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

type PowerExpressionContext struct {
	*SingleExpressionContext
}

func NewPowerExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *PowerExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PowerExpressionContext) Power() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPower, 0)
}

func (s *PowerExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPowerExpression(s)
	}
}

func (s *PowerExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPowerExpression(s)
	}
}

type PreIncrementExpressionContext struct {
	*SingleExpressionContext
}

func NewPreIncrementExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreIncrementExpressionContext {
	var p = new(PreIncrementExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PreIncrementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreIncrementExpressionContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPlusPlus, 0)
}

func (s *PreIncrementExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PreIncrementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPreIncrementExpression(s)
	}
}

func (s *PreIncrementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPreIncrementExpression(s)
	}
}

type ObjectLiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewObjectLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectLiteralExpressionContext {
	var p = new(ObjectLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ObjectLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralExpressionContext) ObjectLiteral() IObjectLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ObjectLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterObjectLiteralExpression(s)
	}
}

func (s *ObjectLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitObjectLiteralExpression(s)
	}
}

type MetaExpressionContext struct {
	*SingleExpressionContext
}

func NewMetaExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MetaExpressionContext {
	var p = new(MetaExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MetaExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaExpressionContext) New() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNew, 0)
}

func (s *MetaExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDot, 0)
}

func (s *MetaExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MetaExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterMetaExpression(s)
	}
}

func (s *MetaExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitMetaExpression(s)
	}
}

type InExpressionContext struct {
	*SingleExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *InExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIn, 0)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type LogicalOrExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *LogicalOrExpressionContext) Or() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOr, 0)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

type NotExpressionContext struct {
	*SingleExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) Not() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNot, 0)
}

func (s *NotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type PreDecreaseExpressionContext struct {
	*SingleExpressionContext
}

func NewPreDecreaseExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreDecreaseExpressionContext {
	var p = new(PreDecreaseExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PreDecreaseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreDecreaseExpressionContext) MinusMinus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMinusMinus, 0)
}

func (s *PreDecreaseExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PreDecreaseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPreDecreaseExpression(s)
	}
}

func (s *PreDecreaseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPreDecreaseExpression(s)
	}
}

type ArgumentsExpressionContext struct {
	*SingleExpressionContext
}

func NewArgumentsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsExpressionContext {
	var p = new(ArgumentsExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArgumentsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentsExpressionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ArgumentsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArgumentsExpression(s)
	}
}

func (s *ArgumentsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArgumentsExpression(s)
	}
}

type AwaitExpressionContext struct {
	*SingleExpressionContext
}

func NewAwaitExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AwaitExpressionContext {
	var p = new(AwaitExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AwaitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AwaitExpressionContext) Await() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAwait, 0)
}

func (s *AwaitExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AwaitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAwaitExpression(s)
	}
}

func (s *AwaitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAwaitExpression(s)
	}
}

type ThisExpressionContext struct {
	*SingleExpressionContext
}

func NewThisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisExpressionContext {
	var p = new(ThisExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ThisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisExpressionContext) This() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserThis, 0)
}

func (s *ThisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterThisExpression(s)
	}
}

func (s *ThisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitThisExpression(s)
	}
}

type FunctionExpressionContext struct {
	*SingleExpressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) AnoymousFunction() IAnoymousFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnoymousFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnoymousFunctionContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

type UnaryMinusExpressionContext struct {
	*SingleExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMinus, 0)
}

func (s *UnaryMinusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *UnaryMinusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterUnaryMinusExpression(s)
	}
}

func (s *UnaryMinusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitUnaryMinusExpression(s)
	}
}

type AssignmentExpressionContext struct {
	*SingleExpressionContext
}

func NewAssignmentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AssignmentExpressionContext) Assign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAssign, 0)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

type PostDecreaseExpressionContext struct {
	*SingleExpressionContext
}

func NewPostDecreaseExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostDecreaseExpressionContext {
	var p = new(PostDecreaseExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PostDecreaseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostDecreaseExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PostDecreaseExpressionContext) MinusMinus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMinusMinus, 0)
}

func (s *PostDecreaseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPostDecreaseExpression(s)
	}
}

func (s *PostDecreaseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPostDecreaseExpression(s)
	}
}

type TypeofExpressionContext struct {
	*SingleExpressionContext
}

func NewTypeofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeofExpressionContext {
	var p = new(TypeofExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TypeofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeofExpressionContext) Typeof() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTypeof, 0)
}

func (s *TypeofExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *TypeofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterTypeofExpression(s)
	}
}

func (s *TypeofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitTypeofExpression(s)
	}
}

type InstanceofExpressionContext struct {
	*SingleExpressionContext
}

func NewInstanceofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanceofExpressionContext {
	var p = new(InstanceofExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *InstanceofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceofExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *InstanceofExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InstanceofExpressionContext) Instanceof() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserInstanceof, 0)
}

func (s *InstanceofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterInstanceofExpression(s)
	}
}

func (s *InstanceofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitInstanceofExpression(s)
	}
}

type UnaryPlusExpressionContext struct {
	*SingleExpressionContext
}

func NewUnaryPlusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryPlusExpressionContext {
	var p = new(UnaryPlusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryPlusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPlusExpressionContext) Plus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPlus, 0)
}

func (s *UnaryPlusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *UnaryPlusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterUnaryPlusExpression(s)
	}
}

func (s *UnaryPlusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitUnaryPlusExpression(s)
	}
}

type DeleteExpressionContext struct {
	*SingleExpressionContext
}

func NewDeleteExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteExpressionContext {
	var p = new(DeleteExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *DeleteExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteExpressionContext) Delete() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDelete, 0)
}

func (s *DeleteExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *DeleteExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterDeleteExpression(s)
	}
}

func (s *DeleteExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitDeleteExpression(s)
	}
}

type ImportExpressionContext struct {
	*SingleExpressionContext
}

func NewImportExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportExpressionContext {
	var p = new(ImportExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ImportExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportExpressionContext) Import() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserImport, 0)
}

func (s *ImportExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *ImportExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ImportExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *ImportExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterImportExpression(s)
	}
}

func (s *ImportExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitImportExpression(s)
	}
}

type EqualityExpressionContext struct {
	*SingleExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *EqualityExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *EqualityExpressionContext) Equals_() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEquals_, 0)
}

func (s *EqualityExpressionContext) NotEquals() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNotEquals, 0)
}

func (s *EqualityExpressionContext) IdentityEquals() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIdentityEquals, 0)
}

func (s *EqualityExpressionContext) IdentityNotEquals() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIdentityNotEquals, 0)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

type BitXOrExpressionContext struct {
	*SingleExpressionContext
}

func NewBitXOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXOrExpressionContext {
	var p = new(BitXOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitXOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitXOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitXOrExpressionContext) BitXOr() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBitXOr, 0)
}

func (s *BitXOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBitXOrExpression(s)
	}
}

func (s *BitXOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBitXOrExpression(s)
	}
}

type SuperExpressionContext struct {
	*SingleExpressionContext
}

func NewSuperExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SuperExpressionContext {
	var p = new(SuperExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *SuperExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperExpressionContext) Super() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSuper, 0)
}

func (s *SuperExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterSuperExpression(s)
	}
}

func (s *SuperExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitSuperExpression(s)
	}
}

type MultiplicativeExpressionContext struct {
	*SingleExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MultiplicativeExpressionContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMultiply, 0)
}

func (s *MultiplicativeExpressionContext) Divide() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDivide, 0)
}

func (s *MultiplicativeExpressionContext) Modulus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserModulus, 0)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

type BitShiftExpressionContext struct {
	*SingleExpressionContext
}

func NewBitShiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitShiftExpressionContext {
	var p = new(BitShiftExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitShiftExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitShiftExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitShiftExpressionContext) LeftShiftArithmetic() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLeftShiftArithmetic, 0)
}

func (s *BitShiftExpressionContext) RightShiftArithmetic() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRightShiftArithmetic, 0)
}

func (s *BitShiftExpressionContext) RightShiftLogical() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRightShiftLogical, 0)
}

func (s *BitShiftExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBitShiftExpression(s)
	}
}

func (s *BitShiftExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBitShiftExpression(s)
	}
}

type ParenthesizedExpressionContext struct {
	*SingleExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *ParenthesizedExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ParenthesizedExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *ParenthesizedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitParenthesizedExpression(s)
	}
}

type AdditiveExpressionContext struct {
	*SingleExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AdditiveExpressionContext) Plus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPlus, 0)
}

func (s *AdditiveExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMinus, 0)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

type RelationalExpressionContext struct {
	*SingleExpressionContext
}

func NewRelationalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *RelationalExpressionContext) LessThan() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLessThan, 0)
}

func (s *RelationalExpressionContext) MoreThan() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMoreThan, 0)
}

func (s *RelationalExpressionContext) LessThanEquals() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLessThanEquals, 0)
}

func (s *RelationalExpressionContext) GreaterThanEquals() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserGreaterThanEquals, 0)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

type PostIncrementExpressionContext struct {
	*SingleExpressionContext
}

func NewPostIncrementExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostIncrementExpressionContext {
	var p = new(PostIncrementExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PostIncrementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostIncrementExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PostIncrementExpressionContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPlusPlus, 0)
}

func (s *PostIncrementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterPostIncrementExpression(s)
	}
}

func (s *PostIncrementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitPostIncrementExpression(s)
	}
}

type YieldExpressionContext struct {
	*SingleExpressionContext
}

func NewYieldExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *YieldExpressionContext {
	var p = new(YieldExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *YieldExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YieldExpressionContext) YieldStatement() IYieldStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYieldStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYieldStatementContext)
}

func (s *YieldExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterYieldExpression(s)
	}
}

func (s *YieldExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitYieldExpression(s)
	}
}

type BitNotExpressionContext struct {
	*SingleExpressionContext
}

func NewBitNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitNotExpressionContext {
	var p = new(BitNotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitNotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitNotExpressionContext) BitNot() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBitNot, 0)
}

func (s *BitNotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitNotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBitNotExpression(s)
	}
}

func (s *BitNotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBitNotExpression(s)
	}
}

type NewExpressionContext struct {
	*SingleExpressionContext
}

func NewNewExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewExpressionContext {
	var p = new(NewExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NewExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewExpressionContext) New() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNew, 0)
}

func (s *NewExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *NewExpressionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *NewExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterNewExpression(s)
	}
}

func (s *NewExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitNewExpression(s)
	}
}

type LiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

type ArrayLiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewArrayLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExpressionContext {
	var p = new(ArrayLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArrayLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArrayLiteralExpression(s)
	}
}

func (s *ArrayLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArrayLiteralExpression(s)
	}
}

type MemberDotExpressionContext struct {
	*SingleExpressionContext
}

func NewMemberDotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberDotExpressionContext {
	var p = new(MemberDotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MemberDotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MemberDotExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDot, 0)
}

func (s *MemberDotExpressionContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *MemberDotExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserQuestionMark, 0)
}

func (s *MemberDotExpressionContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserHashtag, 0)
}

func (s *MemberDotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterMemberDotExpression(s)
	}
}

func (s *MemberDotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitMemberDotExpression(s)
	}
}

type ClassExpressionContext struct {
	*SingleExpressionContext
}

func NewClassExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassExpressionContext {
	var p = new(ClassExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ClassExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassExpressionContext) Class() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserClass, 0)
}

func (s *ClassExpressionContext) ClassTail() IClassTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassTailContext)
}

func (s *ClassExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ClassExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterClassExpression(s)
	}
}

func (s *ClassExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitClassExpression(s)
	}
}

type MemberIndexExpressionContext struct {
	*SingleExpressionContext
}

func NewMemberIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberIndexExpressionContext {
	var p = new(MemberIndexExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MemberIndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberIndexExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MemberIndexExpressionContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBracket, 0)
}

func (s *MemberIndexExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *MemberIndexExpressionContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBracket, 0)
}

func (s *MemberIndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterMemberIndexExpression(s)
	}
}

func (s *MemberIndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitMemberIndexExpression(s)
	}
}

type IdentifierExpressionContext struct {
	*SingleExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

type BitAndExpressionContext struct {
	*SingleExpressionContext
}

func NewBitAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExpressionContext {
	var p = new(BitAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitAndExpressionContext) BitAnd() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBitAnd, 0)
}

func (s *BitAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBitAndExpression(s)
	}
}

func (s *BitAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBitAndExpression(s)
	}
}

type BitOrExpressionContext struct {
	*SingleExpressionContext
}

func NewBitOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExpressionContext {
	var p = new(BitOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitOrExpressionContext) BitOr() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBitOr, 0)
}

func (s *BitOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBitOrExpression(s)
	}
}

func (s *BitOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBitOrExpression(s)
	}
}

type AssignmentOperatorExpressionContext struct {
	*SingleExpressionContext
}

func NewAssignmentOperatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentOperatorExpressionContext {
	var p = new(AssignmentOperatorExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AssignmentOperatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentOperatorExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AssignmentOperatorExpressionContext) AssignmentOperator() IAssignmentOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentOperatorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAssignmentOperatorExpression(s)
	}
}

func (s *AssignmentOperatorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAssignmentOperatorExpression(s)
	}
}

type VoidExpressionContext struct {
	*SingleExpressionContext
}

func NewVoidExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidExpressionContext {
	var p = new(VoidExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *VoidExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidExpressionContext) Void() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserVoid, 0)
}

func (s *VoidExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *VoidExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterVoidExpression(s)
	}
}

func (s *VoidExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitVoidExpression(s)
	}
}

type CoalesceExpressionContext struct {
	*SingleExpressionContext
}

func NewCoalesceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoalesceExpressionContext {
	var p = new(CoalesceExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *CoalesceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoalesceExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *CoalesceExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *CoalesceExpressionContext) NullCoalesce() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNullCoalesce, 0)
}

func (s *CoalesceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterCoalesceExpression(s)
	}
}

func (s *CoalesceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitCoalesceExpression(s)
	}
}

func (p *JavaScriptParser) SingleExpression() (localctx ISingleExpressionContext) {
	return p.singleExpression(0)
}

func (p *JavaScriptParser) singleExpression(_p int) (localctx ISingleExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 114
	p.EnterRecursionRule(localctx, 114, JavaScriptParserRULE_singleExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(775)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 89, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(725)
			p.AnoymousFunction()
		}

	case 2:
		localctx = NewClassExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(726)
			p.Match(JavaScriptParserClass)
		}
		p.SetState(728)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0 {
			{
				p.SetState(727)
				p.Identifier()
			}

		}
		{
			p.SetState(730)
			p.ClassTail()
		}

	case 3:
		localctx = NewNewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(731)
			p.Match(JavaScriptParserNew)
		}
		{
			p.SetState(732)
			p.singleExpression(0)
		}
		p.SetState(734)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 88, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(733)
				p.Arguments()
			}

		}

	case 4:
		localctx = NewMetaExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(736)
			p.Match(JavaScriptParserNew)
		}
		{
			p.SetState(737)
			p.Match(JavaScriptParserDot)
		}
		{
			p.SetState(738)
			p.Identifier()
		}

	case 5:
		localctx = NewDeleteExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(739)
			p.Match(JavaScriptParserDelete)
		}
		{
			p.SetState(740)
			p.singleExpression(37)
		}

	case 6:
		localctx = NewVoidExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(741)
			p.Match(JavaScriptParserVoid)
		}
		{
			p.SetState(742)
			p.singleExpression(36)
		}

	case 7:
		localctx = NewTypeofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(743)
			p.Match(JavaScriptParserTypeof)
		}
		{
			p.SetState(744)
			p.singleExpression(35)
		}

	case 8:
		localctx = NewPreIncrementExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(745)
			p.Match(JavaScriptParserPlusPlus)
		}
		{
			p.SetState(746)
			p.singleExpression(34)
		}

	case 9:
		localctx = NewPreDecreaseExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(747)
			p.Match(JavaScriptParserMinusMinus)
		}
		{
			p.SetState(748)
			p.singleExpression(33)
		}

	case 10:
		localctx = NewUnaryPlusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(749)
			p.Match(JavaScriptParserPlus)
		}
		{
			p.SetState(750)
			p.singleExpression(32)
		}

	case 11:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(751)
			p.Match(JavaScriptParserMinus)
		}
		{
			p.SetState(752)
			p.singleExpression(31)
		}

	case 12:
		localctx = NewBitNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(753)
			p.Match(JavaScriptParserBitNot)
		}
		{
			p.SetState(754)
			p.singleExpression(30)
		}

	case 13:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(755)
			p.Match(JavaScriptParserNot)
		}
		{
			p.SetState(756)
			p.singleExpression(29)
		}

	case 14:
		localctx = NewAwaitExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(757)
			p.Match(JavaScriptParserAwait)
		}
		{
			p.SetState(758)
			p.singleExpression(28)
		}

	case 15:
		localctx = NewImportExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(759)
			p.Match(JavaScriptParserImport)
		}
		{
			p.SetState(760)
			p.Match(JavaScriptParserOpenParen)
		}
		{
			p.SetState(761)
			p.singleExpression(0)
		}
		{
			p.SetState(762)
			p.Match(JavaScriptParserCloseParen)
		}

	case 16:
		localctx = NewYieldExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(764)
			p.YieldStatement()
		}

	case 17:
		localctx = NewThisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(765)
			p.Match(JavaScriptParserThis)
		}

	case 18:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(766)
			p.Identifier()
		}

	case 19:
		localctx = NewSuperExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(767)
			p.Match(JavaScriptParserSuper)
		}

	case 20:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(768)
			p.Literal()
		}

	case 21:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(769)
			p.ArrayLiteral()
		}

	case 22:
		localctx = NewObjectLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(770)
			p.ObjectLiteral()
		}

	case 23:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(771)
			p.Match(JavaScriptParserOpenParen)
		}
		{
			p.SetState(772)
			p.ExpressionSequence()
		}
		{
			p.SetState(773)
			p.Match(JavaScriptParserCloseParen)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(858)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 93, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(856)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 92, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(777)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
				}
				{
					p.SetState(778)
					p.Match(JavaScriptParserPower)
				}
				{
					p.SetState(779)
					p.singleExpression(27)
				}

			case 2:
				localctx = NewMultiplicativeExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(780)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(781)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserMultiply)|(1<<JavaScriptParserDivide)|(1<<JavaScriptParserModulus))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(782)
					p.singleExpression(27)
				}

			case 3:
				localctx = NewAdditiveExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(783)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				{
					p.SetState(784)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JavaScriptParserPlus || _la == JavaScriptParserMinus) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(785)
					p.singleExpression(26)
				}

			case 4:
				localctx = NewCoalesceExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(786)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
				}
				{
					p.SetState(787)
					p.Match(JavaScriptParserNullCoalesce)
				}
				{
					p.SetState(788)
					p.singleExpression(25)
				}

			case 5:
				localctx = NewBitShiftExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(789)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				{
					p.SetState(790)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(JavaScriptParserRightShiftArithmetic-30))|(1<<(JavaScriptParserLeftShiftArithmetic-30))|(1<<(JavaScriptParserRightShiftLogical-30)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(791)
					p.singleExpression(24)
				}

			case 6:
				localctx = NewRelationalExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(792)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(793)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(JavaScriptParserLessThan-33))|(1<<(JavaScriptParserMoreThan-33))|(1<<(JavaScriptParserLessThanEquals-33))|(1<<(JavaScriptParserGreaterThanEquals-33)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(794)
					p.singleExpression(23)
				}

			case 7:
				localctx = NewInstanceofExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(795)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(796)
					p.Match(JavaScriptParserInstanceof)
				}
				{
					p.SetState(797)
					p.singleExpression(22)
				}

			case 8:
				localctx = NewInExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(798)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(799)
					p.Match(JavaScriptParserIn)
				}
				{
					p.SetState(800)
					p.singleExpression(21)
				}

			case 9:
				localctx = NewEqualityExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(801)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(802)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(JavaScriptParserEquals_-37))|(1<<(JavaScriptParserNotEquals-37))|(1<<(JavaScriptParserIdentityEquals-37))|(1<<(JavaScriptParserIdentityNotEquals-37)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(803)
					p.singleExpression(20)
				}

			case 10:
				localctx = NewBitAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(804)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(805)
					p.Match(JavaScriptParserBitAnd)
				}
				{
					p.SetState(806)
					p.singleExpression(19)
				}

			case 11:
				localctx = NewBitXOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(807)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(808)
					p.Match(JavaScriptParserBitXOr)
				}
				{
					p.SetState(809)
					p.singleExpression(18)
				}

			case 12:
				localctx = NewBitOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(810)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(811)
					p.Match(JavaScriptParserBitOr)
				}
				{
					p.SetState(812)
					p.singleExpression(17)
				}

			case 13:
				localctx = NewLogicalAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(813)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(814)
					p.Match(JavaScriptParserAnd)
				}
				{
					p.SetState(815)
					p.singleExpression(16)
				}

			case 14:
				localctx = NewLogicalOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(816)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(817)
					p.Match(JavaScriptParserOr)
				}
				{
					p.SetState(818)
					p.singleExpression(15)
				}

			case 15:
				localctx = NewTernaryExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(819)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(820)
					p.Match(JavaScriptParserQuestionMark)
				}
				{
					p.SetState(821)
					p.singleExpression(0)
				}
				{
					p.SetState(822)
					p.Match(JavaScriptParserColon)
				}
				{
					p.SetState(823)
					p.singleExpression(14)
				}

			case 16:
				localctx = NewAssignmentExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(825)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(826)
					p.Match(JavaScriptParserAssign)
				}
				{
					p.SetState(827)
					p.singleExpression(12)
				}

			case 17:
				localctx = NewAssignmentOperatorExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(828)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(829)
					p.AssignmentOperator()
				}
				{
					p.SetState(830)
					p.singleExpression(11)
				}

			case 18:
				localctx = NewMemberIndexExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(832)

				if !(p.Precpred(p.GetParserRuleContext(), 44)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 44)", ""))
				}
				{
					p.SetState(833)
					p.Match(JavaScriptParserOpenBracket)
				}
				{
					p.SetState(834)
					p.ExpressionSequence()
				}
				{
					p.SetState(835)
					p.Match(JavaScriptParserCloseBracket)
				}

			case 19:
				localctx = NewMemberDotExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(837)

				if !(p.Precpred(p.GetParserRuleContext(), 43)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 43)", ""))
				}
				p.SetState(839)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JavaScriptParserQuestionMark {
					{
						p.SetState(838)
						p.Match(JavaScriptParserQuestionMark)
					}

				}
				{
					p.SetState(841)
					p.Match(JavaScriptParserDot)
				}
				p.SetState(843)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JavaScriptParserHashtag {
					{
						p.SetState(842)
						p.Match(JavaScriptParserHashtag)
					}

				}
				{
					p.SetState(845)
					p.IdentifierName()
				}

			case 20:
				localctx = NewArgumentsExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(846)

				if !(p.Precpred(p.GetParserRuleContext(), 42)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 42)", ""))
				}
				{
					p.SetState(847)
					p.Arguments()
				}

			case 21:
				localctx = NewPostIncrementExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(848)

				if !(p.Precpred(p.GetParserRuleContext(), 39)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 39)", ""))
				}
				p.SetState(849)

				if !(p.notLineTerminator()) {
					panic(antlr.NewFailedPredicateException(p, "p.notLineTerminator()", ""))
				}
				{
					p.SetState(850)
					p.Match(JavaScriptParserPlusPlus)
				}

			case 22:
				localctx = NewPostDecreaseExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(851)

				if !(p.Precpred(p.GetParserRuleContext(), 38)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 38)", ""))
				}
				p.SetState(852)

				if !(p.notLineTerminator()) {
					panic(antlr.NewFailedPredicateException(p, "p.notLineTerminator()", ""))
				}
				{
					p.SetState(853)
					p.Match(JavaScriptParserMinusMinus)
				}

			case 23:
				localctx = NewTemplateStringExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JavaScriptParserRULE_singleExpression)
				p.SetState(854)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(855)
					p.Match(JavaScriptParserTemplateStringLiteral)
				}

			}

		}
		p.SetState(860)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 93, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignableContext is an interface to support dynamic dispatch.
type IAssignableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignableContext differentiates from other interfaces.
	IsAssignableContext()
}

type AssignableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignableContext() *AssignableContext {
	var p = new(AssignableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_assignable
	return p
}

func (*AssignableContext) IsAssignableContext() {}

func NewAssignableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignableContext {
	var p = new(AssignableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_assignable

	return p
}

func (s *AssignableContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignableContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignableContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *AssignableContext) ObjectLiteral() IObjectLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *AssignableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAssignable(s)
	}
}

func (s *AssignableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAssignable(s)
	}
}

func (p *JavaScriptParser) Assignable() (localctx IAssignableContext) {
	localctx = NewAssignableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, JavaScriptParserRULE_assignable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(864)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserAsync, JavaScriptParserNonStrictLet, JavaScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(861)
			p.Identifier()
		}

	case JavaScriptParserOpenBracket:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(862)
			p.ArrayLiteral()
		}

	case JavaScriptParserOpenBrace:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(863)
			p.ObjectLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_objectLiteral
	return p
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenBrace, 0)
}

func (s *ObjectLiteralContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseBrace, 0)
}

func (s *ObjectLiteralContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem())
	var tst = make([]IPropertyAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyAssignmentContext)
		}
	}

	return tst
}

func (s *ObjectLiteralContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyAssignmentContext)
}

func (s *ObjectLiteralContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(JavaScriptParserComma)
}

func (s *ObjectLiteralContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(JavaScriptParserComma, i)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (p *JavaScriptParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, JavaScriptParserRULE_objectLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(866)
		p.Match(JavaScriptParserOpenBrace)
	}
	p.SetState(875)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 96, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(867)
			p.PropertyAssignment()
		}
		p.SetState(872)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 95, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(868)
					p.Match(JavaScriptParserComma)
				}
				{
					p.SetState(869)
					p.PropertyAssignment()
				}

			}
			p.SetState(874)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 95, p.GetParserRuleContext())
		}

	}
	p.SetState(878)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavaScriptParserComma {
		{
			p.SetState(877)
			p.Match(JavaScriptParserComma)
		}

	}
	{
		p.SetState(880)
		p.Match(JavaScriptParserCloseBrace)
	}

	return localctx
}

// IAnoymousFunctionContext is an interface to support dynamic dispatch.
type IAnoymousFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnoymousFunctionContext differentiates from other interfaces.
	IsAnoymousFunctionContext()
}

type AnoymousFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnoymousFunctionContext() *AnoymousFunctionContext {
	var p = new(AnoymousFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_anoymousFunction
	return p
}

func (*AnoymousFunctionContext) IsAnoymousFunctionContext() {}

func NewAnoymousFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnoymousFunctionContext {
	var p = new(AnoymousFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_anoymousFunction

	return p
}

func (s *AnoymousFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AnoymousFunctionContext) CopyFrom(ctx *AnoymousFunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AnoymousFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnoymousFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AnoymousFunctionDeclContext struct {
	*AnoymousFunctionContext
}

func NewAnoymousFunctionDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnoymousFunctionDeclContext {
	var p = new(AnoymousFunctionDeclContext)

	p.AnoymousFunctionContext = NewEmptyAnoymousFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnoymousFunctionContext))

	return p
}

func (s *AnoymousFunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnoymousFunctionDeclContext) Function() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFunction, 0)
}

func (s *AnoymousFunctionDeclContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *AnoymousFunctionDeclContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *AnoymousFunctionDeclContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *AnoymousFunctionDeclContext) Async() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAsync, 0)
}

func (s *AnoymousFunctionDeclContext) Multiply() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMultiply, 0)
}

func (s *AnoymousFunctionDeclContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *AnoymousFunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAnoymousFunctionDecl(s)
	}
}

func (s *AnoymousFunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAnoymousFunctionDecl(s)
	}
}

type ArrowFunctionContext struct {
	*AnoymousFunctionContext
}

func NewArrowFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrowFunctionContext {
	var p = new(ArrowFunctionContext)

	p.AnoymousFunctionContext = NewEmptyAnoymousFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnoymousFunctionContext))

	return p
}

func (s *ArrowFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowFunctionContext) ArrowFunctionParameters() IArrowFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrowFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrowFunctionParametersContext)
}

func (s *ArrowFunctionContext) ARROW() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserARROW, 0)
}

func (s *ArrowFunctionContext) ArrowFunctionBody() IArrowFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrowFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrowFunctionBodyContext)
}

func (s *ArrowFunctionContext) Async() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAsync, 0)
}

func (s *ArrowFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArrowFunction(s)
	}
}

func (s *ArrowFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArrowFunction(s)
	}
}

type FunctionDeclContext struct {
	*AnoymousFunctionContext
}

func NewFunctionDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDeclContext {
	var p = new(FunctionDeclContext)

	p.AnoymousFunctionContext = NewEmptyAnoymousFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnoymousFunctionContext))

	return p
}

func (s *FunctionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *FunctionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterFunctionDecl(s)
	}
}

func (s *FunctionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitFunctionDecl(s)
	}
}

func (p *JavaScriptParser) AnoymousFunction() (localctx IAnoymousFunctionContext) {
	localctx = NewAnoymousFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, JavaScriptParserRULE_anoymousFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(903)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 102, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(882)
			p.FunctionDeclaration()
		}

	case 2:
		localctx = NewAnoymousFunctionDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(884)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserAsync {
			{
				p.SetState(883)
				p.Match(JavaScriptParserAsync)
			}

		}
		{
			p.SetState(886)
			p.Match(JavaScriptParserFunction)
		}
		p.SetState(888)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JavaScriptParserMultiply {
			{
				p.SetState(887)
				p.Match(JavaScriptParserMultiply)
			}

		}
		{
			p.SetState(890)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(892)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis))) != 0) || (((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
			{
				p.SetState(891)
				p.FormalParameterList()
			}

		}
		{
			p.SetState(894)
			p.Match(JavaScriptParserCloseParen)
		}
		{
			p.SetState(895)
			p.FunctionBody()
		}

	case 3:
		localctx = NewArrowFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(897)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 101, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(896)
				p.Match(JavaScriptParserAsync)
			}

		}
		{
			p.SetState(899)
			p.ArrowFunctionParameters()
		}
		{
			p.SetState(900)
			p.Match(JavaScriptParserARROW)
		}
		{
			p.SetState(901)
			p.ArrowFunctionBody()
		}

	}

	return localctx
}

// IArrowFunctionParametersContext is an interface to support dynamic dispatch.
type IArrowFunctionParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrowFunctionParametersContext differentiates from other interfaces.
	IsArrowFunctionParametersContext()
}

type ArrowFunctionParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrowFunctionParametersContext() *ArrowFunctionParametersContext {
	var p = new(ArrowFunctionParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_arrowFunctionParameters
	return p
}

func (*ArrowFunctionParametersContext) IsArrowFunctionParametersContext() {}

func NewArrowFunctionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowFunctionParametersContext {
	var p = new(ArrowFunctionParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_arrowFunctionParameters

	return p
}

func (s *ArrowFunctionParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowFunctionParametersContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArrowFunctionParametersContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOpenParen, 0)
}

func (s *ArrowFunctionParametersContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCloseParen, 0)
}

func (s *ArrowFunctionParametersContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *ArrowFunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowFunctionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrowFunctionParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArrowFunctionParameters(s)
	}
}

func (s *ArrowFunctionParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArrowFunctionParameters(s)
	}
}

func (p *JavaScriptParser) ArrowFunctionParameters() (localctx IArrowFunctionParametersContext) {
	localctx = NewArrowFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, JavaScriptParserRULE_arrowFunctionParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(911)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserAsync, JavaScriptParserNonStrictLet, JavaScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(905)
			p.Identifier()
		}

	case JavaScriptParserOpenParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(906)
			p.Match(JavaScriptParserOpenParen)
		}
		p.SetState(908)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavaScriptParserOpenBracket)|(1<<JavaScriptParserOpenBrace)|(1<<JavaScriptParserEllipsis))) != 0) || (((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
			{
				p.SetState(907)
				p.FormalParameterList()
			}

		}
		{
			p.SetState(910)
			p.Match(JavaScriptParserCloseParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrowFunctionBodyContext is an interface to support dynamic dispatch.
type IArrowFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrowFunctionBodyContext differentiates from other interfaces.
	IsArrowFunctionBodyContext()
}

type ArrowFunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrowFunctionBodyContext() *ArrowFunctionBodyContext {
	var p = new(ArrowFunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_arrowFunctionBody
	return p
}

func (*ArrowFunctionBodyContext) IsArrowFunctionBodyContext() {}

func NewArrowFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowFunctionBodyContext {
	var p = new(ArrowFunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_arrowFunctionBody

	return p
}

func (s *ArrowFunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowFunctionBodyContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArrowFunctionBodyContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *ArrowFunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowFunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrowFunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterArrowFunctionBody(s)
	}
}

func (s *ArrowFunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitArrowFunctionBody(s)
	}
}

func (p *JavaScriptParser) ArrowFunctionBody() (localctx IArrowFunctionBodyContext) {
	localctx = NewArrowFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, JavaScriptParserRULE_arrowFunctionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(915)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 105, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(913)
			p.singleExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(914)
			p.FunctionBody()
		}

	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) MultiplyAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMultiplyAssign, 0)
}

func (s *AssignmentOperatorContext) DivideAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDivideAssign, 0)
}

func (s *AssignmentOperatorContext) ModulusAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserModulusAssign, 0)
}

func (s *AssignmentOperatorContext) PlusAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPlusAssign, 0)
}

func (s *AssignmentOperatorContext) MinusAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserMinusAssign, 0)
}

func (s *AssignmentOperatorContext) LeftShiftArithmeticAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserLeftShiftArithmeticAssign, 0)
}

func (s *AssignmentOperatorContext) RightShiftArithmeticAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRightShiftArithmeticAssign, 0)
}

func (s *AssignmentOperatorContext) RightShiftLogicalAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRightShiftLogicalAssign, 0)
}

func (s *AssignmentOperatorContext) BitAndAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBitAndAssign, 0)
}

func (s *AssignmentOperatorContext) BitXorAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBitXorAssign, 0)
}

func (s *AssignmentOperatorContext) BitOrAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBitOrAssign, 0)
}

func (s *AssignmentOperatorContext) PowerAssign() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPowerAssign, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (p *JavaScriptParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, JavaScriptParserRULE_assignmentOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(917)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(JavaScriptParserMultiplyAssign-46))|(1<<(JavaScriptParserDivideAssign-46))|(1<<(JavaScriptParserModulusAssign-46))|(1<<(JavaScriptParserPlusAssign-46))|(1<<(JavaScriptParserMinusAssign-46))|(1<<(JavaScriptParserLeftShiftArithmeticAssign-46))|(1<<(JavaScriptParserRightShiftArithmeticAssign-46))|(1<<(JavaScriptParserRightShiftLogicalAssign-46))|(1<<(JavaScriptParserBitAndAssign-46))|(1<<(JavaScriptParserBitXorAssign-46))|(1<<(JavaScriptParserBitOrAssign-46))|(1<<(JavaScriptParserPowerAssign-46)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNullLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBooleanLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserStringLiteral, 0)
}

func (s *LiteralContext) TemplateStringLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTemplateStringLiteral, 0)
}

func (s *LiteralContext) RegularExpressionLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserRegularExpressionLiteral, 0)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *LiteralContext) BigintLiteral() IBigintLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBigintLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBigintLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *JavaScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, JavaScriptParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(926)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(919)
			p.Match(JavaScriptParserNullLiteral)
		}

	case JavaScriptParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(920)
			p.Match(JavaScriptParserBooleanLiteral)
		}

	case JavaScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(921)
			p.Match(JavaScriptParserStringLiteral)
		}

	case JavaScriptParserTemplateStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(922)
			p.Match(JavaScriptParserTemplateStringLiteral)
		}

	case JavaScriptParserRegularExpressionLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(923)
			p.Match(JavaScriptParserRegularExpressionLiteral)
		}

	case JavaScriptParserDecimalLiteral, JavaScriptParserHexIntegerLiteral, JavaScriptParserOctalIntegerLiteral, JavaScriptParserOctalIntegerLiteral2, JavaScriptParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(924)
			p.NumericLiteral()
		}

	case JavaScriptParserBigHexIntegerLiteral, JavaScriptParserBigOctalIntegerLiteral, JavaScriptParserBigBinaryIntegerLiteral, JavaScriptParserBigDecimalIntegerLiteral:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(925)
			p.BigintLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDecimalLiteral, 0)
}

func (s *NumericLiteralContext) HexIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserHexIntegerLiteral, 0)
}

func (s *NumericLiteralContext) OctalIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOctalIntegerLiteral, 0)
}

func (s *NumericLiteralContext) OctalIntegerLiteral2() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserOctalIntegerLiteral2, 0)
}

func (s *NumericLiteralContext) BinaryIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBinaryIntegerLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *JavaScriptParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, JavaScriptParserRULE_numericLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(928)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-61)&-(0x1f+1)) == 0 && ((1<<uint((_la-61)))&((1<<(JavaScriptParserDecimalLiteral-61))|(1<<(JavaScriptParserHexIntegerLiteral-61))|(1<<(JavaScriptParserOctalIntegerLiteral-61))|(1<<(JavaScriptParserOctalIntegerLiteral2-61))|(1<<(JavaScriptParserBinaryIntegerLiteral-61)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBigintLiteralContext is an interface to support dynamic dispatch.
type IBigintLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBigintLiteralContext differentiates from other interfaces.
	IsBigintLiteralContext()
}

type BigintLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBigintLiteralContext() *BigintLiteralContext {
	var p = new(BigintLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_bigintLiteral
	return p
}

func (*BigintLiteralContext) IsBigintLiteralContext() {}

func NewBigintLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BigintLiteralContext {
	var p = new(BigintLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_bigintLiteral

	return p
}

func (s *BigintLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BigintLiteralContext) BigDecimalIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBigDecimalIntegerLiteral, 0)
}

func (s *BigintLiteralContext) BigHexIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBigHexIntegerLiteral, 0)
}

func (s *BigintLiteralContext) BigOctalIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBigOctalIntegerLiteral, 0)
}

func (s *BigintLiteralContext) BigBinaryIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBigBinaryIntegerLiteral, 0)
}

func (s *BigintLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BigintLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BigintLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterBigintLiteral(s)
	}
}

func (s *BigintLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitBigintLiteral(s)
	}
}

func (p *JavaScriptParser) BigintLiteral() (localctx IBigintLiteralContext) {
	localctx = NewBigintLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, JavaScriptParserRULE_bigintLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(930)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(JavaScriptParserBigHexIntegerLiteral-66))|(1<<(JavaScriptParserBigOctalIntegerLiteral-66))|(1<<(JavaScriptParserBigBinaryIntegerLiteral-66))|(1<<(JavaScriptParserBigDecimalIntegerLiteral-66)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGetterContext is an interface to support dynamic dispatch.
type IGetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetterContext differentiates from other interfaces.
	IsGetterContext()
}

type GetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetterContext() *GetterContext {
	var p = new(GetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_getter
	return p
}

func (*GetterContext) IsGetterContext() {}

func NewGetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetterContext {
	var p = new(GetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_getter

	return p
}

func (s *GetterContext) GetParser() antlr.Parser { return s.parser }

func (s *GetterContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *GetterContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *GetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterGetter(s)
	}
}

func (s *GetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitGetter(s)
	}
}

func (p *JavaScriptParser) Getter() (localctx IGetterContext) {
	localctx = NewGetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, JavaScriptParserRULE_getter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(932)

	if !(p.n("get")) {
		panic(antlr.NewFailedPredicateException(p, "p.n(\"get\")", ""))
	}
	{
		p.SetState(933)
		p.Identifier()
	}
	{
		p.SetState(934)
		p.PropertyName()
	}

	return localctx
}

// ISetterContext is an interface to support dynamic dispatch.
type ISetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetterContext differentiates from other interfaces.
	IsSetterContext()
}

type SetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetterContext() *SetterContext {
	var p = new(SetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_setter
	return p
}

func (*SetterContext) IsSetterContext() {}

func NewSetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetterContext {
	var p = new(SetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_setter

	return p
}

func (s *SetterContext) GetParser() antlr.Parser { return s.parser }

func (s *SetterContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SetterContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *SetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterSetter(s)
	}
}

func (s *SetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitSetter(s)
	}
}

func (p *JavaScriptParser) Setter() (localctx ISetterContext) {
	localctx = NewSetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, JavaScriptParserRULE_setter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(936)

	if !(p.n("set")) {
		panic(antlr.NewFailedPredicateException(p, "p.n(\"set\")", ""))
	}
	{
		p.SetState(937)
		p.Identifier()
	}
	{
		p.SetState(938)
		p.PropertyName()
	}

	return localctx
}

// IIdentifierNameContext is an interface to support dynamic dispatch.
type IIdentifierNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierNameContext differentiates from other interfaces.
	IsIdentifierNameContext()
}

type IdentifierNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierNameContext() *IdentifierNameContext {
	var p = new(IdentifierNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_identifierName
	return p
}

func (*IdentifierNameContext) IsIdentifierNameContext() {}

func NewIdentifierNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierNameContext {
	var p = new(IdentifierNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_identifierName

	return p
}

func (s *IdentifierNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierNameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierNameContext) ReservedWord() IReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *IdentifierNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterIdentifierName(s)
	}
}

func (s *IdentifierNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitIdentifierName(s)
	}
}

func (p *JavaScriptParser) IdentifierName() (localctx IIdentifierNameContext) {
	localctx = NewIdentifierNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, JavaScriptParserRULE_identifierName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(942)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 107, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(940)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(941)
			p.ReservedWord()
		}

	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIdentifier, 0)
}

func (s *IdentifierContext) NonStrictLet() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNonStrictLet, 0)
}

func (s *IdentifierContext) Async() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAsync, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *JavaScriptParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, JavaScriptParserRULE_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(944)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-105)&-(0x1f+1)) == 0 && ((1<<uint((_la-105)))&((1<<(JavaScriptParserAsync-105))|(1<<(JavaScriptParserNonStrictLet-105))|(1<<(JavaScriptParserIdentifier-105)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReservedWordContext is an interface to support dynamic dispatch.
type IReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedWordContext differentiates from other interfaces.
	IsReservedWordContext()
}

type ReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedWordContext() *ReservedWordContext {
	var p = new(ReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_reservedWord
	return p
}

func (*ReservedWordContext) IsReservedWordContext() {}

func NewReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedWordContext {
	var p = new(ReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_reservedWord

	return p
}

func (s *ReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedWordContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *ReservedWordContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNullLiteral, 0)
}

func (s *ReservedWordContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBooleanLiteral, 0)
}

func (s *ReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterReservedWord(s)
	}
}

func (s *ReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitReservedWord(s)
	}
}

func (p *JavaScriptParser) ReservedWord() (localctx IReservedWordContext) {
	localctx = NewReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, JavaScriptParserRULE_reservedWord)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(949)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserBreak, JavaScriptParserDo, JavaScriptParserInstanceof, JavaScriptParserTypeof, JavaScriptParserCase, JavaScriptParserElse, JavaScriptParserNew, JavaScriptParserVar, JavaScriptParserCatch, JavaScriptParserFinally, JavaScriptParserReturn, JavaScriptParserVoid, JavaScriptParserContinue, JavaScriptParserFor, JavaScriptParserSwitch, JavaScriptParserWhile, JavaScriptParserDebugger, JavaScriptParserFunction, JavaScriptParserThis, JavaScriptParserWith, JavaScriptParserDefault, JavaScriptParserIf, JavaScriptParserThrow, JavaScriptParserDelete, JavaScriptParserIn, JavaScriptParserTry, JavaScriptParserAs, JavaScriptParserFrom, JavaScriptParserClass, JavaScriptParserEnum, JavaScriptParserExtends, JavaScriptParserSuper, JavaScriptParserConst, JavaScriptParserExport, JavaScriptParserImport, JavaScriptParserAsync, JavaScriptParserAwait, JavaScriptParserImplements, JavaScriptParserStrictLet, JavaScriptParserNonStrictLet, JavaScriptParserPrivate, JavaScriptParserPublic, JavaScriptParserInterface, JavaScriptParserPackage, JavaScriptParserProtected, JavaScriptParserStatic, JavaScriptParserYield:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(946)
			p.Keyword()
		}

	case JavaScriptParserNullLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(947)
			p.Match(JavaScriptParserNullLiteral)
		}

	case JavaScriptParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(948)
			p.Match(JavaScriptParserBooleanLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Break() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserBreak, 0)
}

func (s *KeywordContext) Do() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDo, 0)
}

func (s *KeywordContext) Instanceof() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserInstanceof, 0)
}

func (s *KeywordContext) Typeof() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTypeof, 0)
}

func (s *KeywordContext) Case() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCase, 0)
}

func (s *KeywordContext) Else() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserElse, 0)
}

func (s *KeywordContext) New() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNew, 0)
}

func (s *KeywordContext) Var() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserVar, 0)
}

func (s *KeywordContext) Catch() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserCatch, 0)
}

func (s *KeywordContext) Finally() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFinally, 0)
}

func (s *KeywordContext) Return() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserReturn, 0)
}

func (s *KeywordContext) Void() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserVoid, 0)
}

func (s *KeywordContext) Continue() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserContinue, 0)
}

func (s *KeywordContext) For() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFor, 0)
}

func (s *KeywordContext) Switch() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSwitch, 0)
}

func (s *KeywordContext) While() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserWhile, 0)
}

func (s *KeywordContext) Debugger() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDebugger, 0)
}

func (s *KeywordContext) Function() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFunction, 0)
}

func (s *KeywordContext) This() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserThis, 0)
}

func (s *KeywordContext) With() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserWith, 0)
}

func (s *KeywordContext) Default() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDefault, 0)
}

func (s *KeywordContext) If() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIf, 0)
}

func (s *KeywordContext) Throw() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserThrow, 0)
}

func (s *KeywordContext) Delete() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserDelete, 0)
}

func (s *KeywordContext) In() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserIn, 0)
}

func (s *KeywordContext) Try() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserTry, 0)
}

func (s *KeywordContext) Class() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserClass, 0)
}

func (s *KeywordContext) Enum() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEnum, 0)
}

func (s *KeywordContext) Extends() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserExtends, 0)
}

func (s *KeywordContext) Super() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSuper, 0)
}

func (s *KeywordContext) Const() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserConst, 0)
}

func (s *KeywordContext) Export() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserExport, 0)
}

func (s *KeywordContext) Import() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserImport, 0)
}

func (s *KeywordContext) Implements() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserImplements, 0)
}

func (s *KeywordContext) Let_() ILet_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILet_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILet_Context)
}

func (s *KeywordContext) Private() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPrivate, 0)
}

func (s *KeywordContext) Public() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPublic, 0)
}

func (s *KeywordContext) Interface() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserInterface, 0)
}

func (s *KeywordContext) Package() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserPackage, 0)
}

func (s *KeywordContext) Protected() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserProtected, 0)
}

func (s *KeywordContext) Static() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserStatic, 0)
}

func (s *KeywordContext) Yield() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserYield, 0)
}

func (s *KeywordContext) Async() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAsync, 0)
}

func (s *KeywordContext) Await() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAwait, 0)
}

func (s *KeywordContext) From() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserFrom, 0)
}

func (s *KeywordContext) As() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserAs, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *JavaScriptParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, JavaScriptParserRULE_keyword)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(997)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavaScriptParserBreak:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(951)
			p.Match(JavaScriptParserBreak)
		}

	case JavaScriptParserDo:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(952)
			p.Match(JavaScriptParserDo)
		}

	case JavaScriptParserInstanceof:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(953)
			p.Match(JavaScriptParserInstanceof)
		}

	case JavaScriptParserTypeof:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(954)
			p.Match(JavaScriptParserTypeof)
		}

	case JavaScriptParserCase:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(955)
			p.Match(JavaScriptParserCase)
		}

	case JavaScriptParserElse:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(956)
			p.Match(JavaScriptParserElse)
		}

	case JavaScriptParserNew:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(957)
			p.Match(JavaScriptParserNew)
		}

	case JavaScriptParserVar:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(958)
			p.Match(JavaScriptParserVar)
		}

	case JavaScriptParserCatch:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(959)
			p.Match(JavaScriptParserCatch)
		}

	case JavaScriptParserFinally:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(960)
			p.Match(JavaScriptParserFinally)
		}

	case JavaScriptParserReturn:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(961)
			p.Match(JavaScriptParserReturn)
		}

	case JavaScriptParserVoid:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(962)
			p.Match(JavaScriptParserVoid)
		}

	case JavaScriptParserContinue:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(963)
			p.Match(JavaScriptParserContinue)
		}

	case JavaScriptParserFor:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(964)
			p.Match(JavaScriptParserFor)
		}

	case JavaScriptParserSwitch:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(965)
			p.Match(JavaScriptParserSwitch)
		}

	case JavaScriptParserWhile:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(966)
			p.Match(JavaScriptParserWhile)
		}

	case JavaScriptParserDebugger:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(967)
			p.Match(JavaScriptParserDebugger)
		}

	case JavaScriptParserFunction:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(968)
			p.Match(JavaScriptParserFunction)
		}

	case JavaScriptParserThis:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(969)
			p.Match(JavaScriptParserThis)
		}

	case JavaScriptParserWith:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(970)
			p.Match(JavaScriptParserWith)
		}

	case JavaScriptParserDefault:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(971)
			p.Match(JavaScriptParserDefault)
		}

	case JavaScriptParserIf:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(972)
			p.Match(JavaScriptParserIf)
		}

	case JavaScriptParserThrow:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(973)
			p.Match(JavaScriptParserThrow)
		}

	case JavaScriptParserDelete:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(974)
			p.Match(JavaScriptParserDelete)
		}

	case JavaScriptParserIn:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(975)
			p.Match(JavaScriptParserIn)
		}

	case JavaScriptParserTry:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(976)
			p.Match(JavaScriptParserTry)
		}

	case JavaScriptParserClass:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(977)
			p.Match(JavaScriptParserClass)
		}

	case JavaScriptParserEnum:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(978)
			p.Match(JavaScriptParserEnum)
		}

	case JavaScriptParserExtends:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(979)
			p.Match(JavaScriptParserExtends)
		}

	case JavaScriptParserSuper:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(980)
			p.Match(JavaScriptParserSuper)
		}

	case JavaScriptParserConst:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(981)
			p.Match(JavaScriptParserConst)
		}

	case JavaScriptParserExport:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(982)
			p.Match(JavaScriptParserExport)
		}

	case JavaScriptParserImport:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(983)
			p.Match(JavaScriptParserImport)
		}

	case JavaScriptParserImplements:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(984)
			p.Match(JavaScriptParserImplements)
		}

	case JavaScriptParserStrictLet, JavaScriptParserNonStrictLet:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(985)
			p.Let_()
		}

	case JavaScriptParserPrivate:
		p.EnterOuterAlt(localctx, 36)
		{
			p.SetState(986)
			p.Match(JavaScriptParserPrivate)
		}

	case JavaScriptParserPublic:
		p.EnterOuterAlt(localctx, 37)
		{
			p.SetState(987)
			p.Match(JavaScriptParserPublic)
		}

	case JavaScriptParserInterface:
		p.EnterOuterAlt(localctx, 38)
		{
			p.SetState(988)
			p.Match(JavaScriptParserInterface)
		}

	case JavaScriptParserPackage:
		p.EnterOuterAlt(localctx, 39)
		{
			p.SetState(989)
			p.Match(JavaScriptParserPackage)
		}

	case JavaScriptParserProtected:
		p.EnterOuterAlt(localctx, 40)
		{
			p.SetState(990)
			p.Match(JavaScriptParserProtected)
		}

	case JavaScriptParserStatic:
		p.EnterOuterAlt(localctx, 41)
		{
			p.SetState(991)
			p.Match(JavaScriptParserStatic)
		}

	case JavaScriptParserYield:
		p.EnterOuterAlt(localctx, 42)
		{
			p.SetState(992)
			p.Match(JavaScriptParserYield)
		}

	case JavaScriptParserAsync:
		p.EnterOuterAlt(localctx, 43)
		{
			p.SetState(993)
			p.Match(JavaScriptParserAsync)
		}

	case JavaScriptParserAwait:
		p.EnterOuterAlt(localctx, 44)
		{
			p.SetState(994)
			p.Match(JavaScriptParserAwait)
		}

	case JavaScriptParserFrom:
		p.EnterOuterAlt(localctx, 45)
		{
			p.SetState(995)
			p.Match(JavaScriptParserFrom)
		}

	case JavaScriptParserAs:
		p.EnterOuterAlt(localctx, 46)
		{
			p.SetState(996)
			p.Match(JavaScriptParserAs)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILet_Context is an interface to support dynamic dispatch.
type ILet_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLet_Context differentiates from other interfaces.
	IsLet_Context()
}

type Let_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLet_Context() *Let_Context {
	var p = new(Let_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_let_
	return p
}

func (*Let_Context) IsLet_Context() {}

func NewLet_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_Context {
	var p = new(Let_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_let_

	return p
}

func (s *Let_Context) GetParser() antlr.Parser { return s.parser }

func (s *Let_Context) NonStrictLet() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserNonStrictLet, 0)
}

func (s *Let_Context) StrictLet() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserStrictLet, 0)
}

func (s *Let_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterLet_(s)
	}
}

func (s *Let_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitLet_(s)
	}
}

func (p *JavaScriptParser) Let_() (localctx ILet_Context) {
	localctx = NewLet_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, JavaScriptParserRULE_let_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(999)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JavaScriptParserStrictLet || _la == JavaScriptParserNonStrictLet) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavaScriptParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavaScriptParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserSemiColon, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(JavaScriptParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavaScriptParserListener); ok {
		listenerT.ExitEos(s)
	}
}

func (p *JavaScriptParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, JavaScriptParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(1005)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 110, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(1001)
			p.Match(JavaScriptParserSemiColon)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(1002)
			p.Match(JavaScriptParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(1003)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(1004)

		if !(p.closeBrace()) {
			panic(antlr.NewFailedPredicateException(p, "p.closeBrace()", ""))
		}

	}

	return localctx
}

func (p *JavaScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *ExpressionStatementContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionStatementContext)
		}
		return p.ExpressionStatement_Sempred(t, predIndex)

	case 21:
		var t *IterationStatementContext = nil
		if localctx != nil {
			t = localctx.(*IterationStatementContext)
		}
		return p.IterationStatement_Sempred(t, predIndex)

	case 23:
		var t *ContinueStatementContext = nil
		if localctx != nil {
			t = localctx.(*ContinueStatementContext)
		}
		return p.ContinueStatement_Sempred(t, predIndex)

	case 24:
		var t *BreakStatementContext = nil
		if localctx != nil {
			t = localctx.(*BreakStatementContext)
		}
		return p.BreakStatement_Sempred(t, predIndex)

	case 25:
		var t *ReturnStatementContext = nil
		if localctx != nil {
			t = localctx.(*ReturnStatementContext)
		}
		return p.ReturnStatement_Sempred(t, predIndex)

	case 26:
		var t *YieldStatementContext = nil
		if localctx != nil {
			t = localctx.(*YieldStatementContext)
		}
		return p.YieldStatement_Sempred(t, predIndex)

	case 34:
		var t *ThrowStatementContext = nil
		if localctx != nil {
			t = localctx.(*ThrowStatementContext)
		}
		return p.ThrowStatement_Sempred(t, predIndex)

	case 42:
		var t *ClassElementContext = nil
		if localctx != nil {
			t = localctx.(*ClassElementContext)
		}
		return p.ClassElement_Sempred(t, predIndex)

	case 57:
		var t *SingleExpressionContext = nil
		if localctx != nil {
			t = localctx.(*SingleExpressionContext)
		}
		return p.SingleExpression_Sempred(t, predIndex)

	case 67:
		var t *GetterContext = nil
		if localctx != nil {
			t = localctx.(*GetterContext)
		}
		return p.Getter_Sempred(t, predIndex)

	case 68:
		var t *SetterContext = nil
		if localctx != nil {
			t = localctx.(*SetterContext)
		}
		return p.Setter_Sempred(t, predIndex)

	case 74:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JavaScriptParser) ExpressionStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.notOpenBraceAndNotFunction()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) IterationStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.p("of")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) ContinueStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.notLineTerminator()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) BreakStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.notLineTerminator()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) ReturnStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.notLineTerminator()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) YieldStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.notLineTerminator()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) ThrowStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.notLineTerminator()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) ClassElement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.n("static")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 22:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 23:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 24:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 25:
		return p.Precpred(p.GetParserRuleContext(), 44)

	case 26:
		return p.Precpred(p.GetParserRuleContext(), 43)

	case 27:
		return p.Precpred(p.GetParserRuleContext(), 42)

	case 28:
		return p.Precpred(p.GetParserRuleContext(), 39)

	case 29:
		return p.notLineTerminator()

	case 30:
		return p.Precpred(p.GetParserRuleContext(), 38)

	case 31:
		return p.notLineTerminator()

	case 32:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) Getter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 33:
		return p.n("get")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) Setter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 34:
		return p.n("set")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 35:
		return p.lineTerminatorAhead()

	case 36:
		return p.closeBrace()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
