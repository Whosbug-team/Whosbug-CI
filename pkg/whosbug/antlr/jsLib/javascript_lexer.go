// Code generated from JavaScriptLexer.g4 by ANTLR 4.9.2. DO NOT EDIT.

package jsLib

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

type JavaScriptLexer struct {
	JavaScriptLexerBase
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewJavaScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *JavaScriptLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJavaScriptLexer(input antlr.CharStream) *JavaScriptLexer {
	l := new(JavaScriptLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "JavaScriptLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JavaScriptLexer tokens.
const (
	JavaScriptLexerHashBangLine               = 1
	JavaScriptLexerMultiLineComment           = 2
	JavaScriptLexerSingleLineComment          = 3
	JavaScriptLexerRegularExpressionLiteral   = 4
	JavaScriptLexerOpenBracket                = 5
	JavaScriptLexerCloseBracket               = 6
	JavaScriptLexerOpenParen                  = 7
	JavaScriptLexerCloseParen                 = 8
	JavaScriptLexerOpenBrace                  = 9
	JavaScriptLexerCloseBrace                 = 10
	JavaScriptLexerSemiColon                  = 11
	JavaScriptLexerComma                      = 12
	JavaScriptLexerAssign                     = 13
	JavaScriptLexerQuestionMark               = 14
	JavaScriptLexerColon                      = 15
	JavaScriptLexerEllipsis                   = 16
	JavaScriptLexerDot                        = 17
	JavaScriptLexerPlusPlus                   = 18
	JavaScriptLexerMinusMinus                 = 19
	JavaScriptLexerPlus                       = 20
	JavaScriptLexerMinus                      = 21
	JavaScriptLexerBitNot                     = 22
	JavaScriptLexerNot                        = 23
	JavaScriptLexerMultiply                   = 24
	JavaScriptLexerDivide                     = 25
	JavaScriptLexerModulus                    = 26
	JavaScriptLexerPower                      = 27
	JavaScriptLexerNullCoalesce               = 28
	JavaScriptLexerHashtag                    = 29
	JavaScriptLexerRightShiftArithmetic       = 30
	JavaScriptLexerLeftShiftArithmetic        = 31
	JavaScriptLexerRightShiftLogical          = 32
	JavaScriptLexerLessThan                   = 33
	JavaScriptLexerMoreThan                   = 34
	JavaScriptLexerLessThanEquals             = 35
	JavaScriptLexerGreaterThanEquals          = 36
	JavaScriptLexerEquals_                    = 37
	JavaScriptLexerNotEquals                  = 38
	JavaScriptLexerIdentityEquals             = 39
	JavaScriptLexerIdentityNotEquals          = 40
	JavaScriptLexerBitAnd                     = 41
	JavaScriptLexerBitXOr                     = 42
	JavaScriptLexerBitOr                      = 43
	JavaScriptLexerAnd                        = 44
	JavaScriptLexerOr                         = 45
	JavaScriptLexerMultiplyAssign             = 46
	JavaScriptLexerDivideAssign               = 47
	JavaScriptLexerModulusAssign              = 48
	JavaScriptLexerPlusAssign                 = 49
	JavaScriptLexerMinusAssign                = 50
	JavaScriptLexerLeftShiftArithmeticAssign  = 51
	JavaScriptLexerRightShiftArithmeticAssign = 52
	JavaScriptLexerRightShiftLogicalAssign    = 53
	JavaScriptLexerBitAndAssign               = 54
	JavaScriptLexerBitXorAssign               = 55
	JavaScriptLexerBitOrAssign                = 56
	JavaScriptLexerPowerAssign                = 57
	JavaScriptLexerARROW                      = 58
	JavaScriptLexerNullLiteral                = 59
	JavaScriptLexerBooleanLiteral             = 60
	JavaScriptLexerDecimalLiteral             = 61
	JavaScriptLexerHexIntegerLiteral          = 62
	JavaScriptLexerOctalIntegerLiteral        = 63
	JavaScriptLexerOctalIntegerLiteral2       = 64
	JavaScriptLexerBinaryIntegerLiteral       = 65
	JavaScriptLexerBigHexIntegerLiteral       = 66
	JavaScriptLexerBigOctalIntegerLiteral     = 67
	JavaScriptLexerBigBinaryIntegerLiteral    = 68
	JavaScriptLexerBigDecimalIntegerLiteral   = 69
	JavaScriptLexerBreak                      = 70
	JavaScriptLexerDo                         = 71
	JavaScriptLexerInstanceof                 = 72
	JavaScriptLexerTypeof                     = 73
	JavaScriptLexerCase                       = 74
	JavaScriptLexerElse                       = 75
	JavaScriptLexerNew                        = 76
	JavaScriptLexerVar                        = 77
	JavaScriptLexerCatch                      = 78
	JavaScriptLexerFinally                    = 79
	JavaScriptLexerReturn                     = 80
	JavaScriptLexerVoid                       = 81
	JavaScriptLexerContinue                   = 82
	JavaScriptLexerFor                        = 83
	JavaScriptLexerSwitch                     = 84
	JavaScriptLexerWhile                      = 85
	JavaScriptLexerDebugger                   = 86
	JavaScriptLexerFunction                   = 87
	JavaScriptLexerThis                       = 88
	JavaScriptLexerWith                       = 89
	JavaScriptLexerDefault                    = 90
	JavaScriptLexerIf                         = 91
	JavaScriptLexerThrow                      = 92
	JavaScriptLexerDelete                     = 93
	JavaScriptLexerIn                         = 94
	JavaScriptLexerTry                        = 95
	JavaScriptLexerAs                         = 96
	JavaScriptLexerFrom                       = 97
	JavaScriptLexerClass                      = 98
	JavaScriptLexerEnum                       = 99
	JavaScriptLexerExtends                    = 100
	JavaScriptLexerSuper                      = 101
	JavaScriptLexerConst                      = 102
	JavaScriptLexerExport                     = 103
	JavaScriptLexerImport                     = 104
	JavaScriptLexerAsync                      = 105
	JavaScriptLexerAwait                      = 106
	JavaScriptLexerImplements                 = 107
	JavaScriptLexerStrictLet                  = 108
	JavaScriptLexerNonStrictLet               = 109
	JavaScriptLexerPrivate                    = 110
	JavaScriptLexerPublic                     = 111
	JavaScriptLexerInterface                  = 112
	JavaScriptLexerPackage                    = 113
	JavaScriptLexerProtected                  = 114
	JavaScriptLexerStatic                     = 115
	JavaScriptLexerYield                      = 116
	JavaScriptLexerIdentifier                 = 117
	JavaScriptLexerStringLiteral              = 118
	JavaScriptLexerTemplateStringLiteral      = 119
	JavaScriptLexerWhiteSpaces                = 120
	JavaScriptLexerLineTerminator             = 121
	JavaScriptLexerHtmlComment                = 122
	JavaScriptLexerCDataComment               = 123
	JavaScriptLexerUnexpectedCharacter        = 124
)

// JavaScriptLexerERROR is the JavaScriptLexer channel.
const JavaScriptLexerERROR = 2

func (l *JavaScriptLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 8:
		l.OpenBrace_Action(localctx, actionIndex)

	case 9:
		l.CloseBrace_Action(localctx, actionIndex)

	case 117:
		l.StringLiteral_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *JavaScriptLexer) OpenBrace_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		l.ProcessOpenBrace()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *JavaScriptLexer) CloseBrace_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
		l.ProcessCloseBrace()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *JavaScriptLexer) StringLiteral_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 2:
		l.ProcessStringLiteral()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}

func (l *JavaScriptLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		return l.HashBangLine_Sempred(localctx, predIndex)

	case 3:
		return l.RegularExpressionLiteral_Sempred(localctx, predIndex)

	case 62:
		return l.OctalIntegerLiteral_Sempred(localctx, predIndex)

	case 106:
		return l.Implements_Sempred(localctx, predIndex)

	case 107:
		return l.StrictLet_Sempred(localctx, predIndex)

	case 108:
		return l.NonStrictLet_Sempred(localctx, predIndex)

	case 109:
		return l.Private_Sempred(localctx, predIndex)

	case 110:
		return l.Public_Sempred(localctx, predIndex)

	case 111:
		return l.Interface_Sempred(localctx, predIndex)

	case 112:
		return l.Package_Sempred(localctx, predIndex)

	case 113:
		return l.Protected_Sempred(localctx, predIndex)

	case 114:
		return l.Static_Sempred(localctx, predIndex)

	case 115:
		return l.Yield_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JavaScriptLexer) HashBangLine_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.IsStartOfFile()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) RegularExpressionLiteral_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.IsRegexPossible()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) OctalIntegerLiteral_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return !p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Implements_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) StrictLet_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) NonStrictLet_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return !p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Private_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Public_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Interface_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Package_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Protected_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Static_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JavaScriptLexer) Yield_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.IsStrictMode()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
