// Generated from ObjectiveCPreprocessorParser.g4 by ANTLR 4.7.

package objcParser // ObjectiveCPreprocessorParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 224, 97,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 44, 10, 2, 5, 2, 46, 10, 2, 3, 3, 6,
	3, 49, 10, 3, 13, 3, 14, 3, 50, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 76, 10, 4, 5, 4, 78, 10, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4,
	92, 10, 4, 12, 4, 14, 4, 95, 11, 4, 3, 4, 2, 3, 6, 5, 2, 4, 6, 2, 6, 3,
	2, 188, 189, 3, 2, 223, 224, 3, 2, 207, 208, 3, 2, 211, 214, 2, 119, 2,
	45, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 77, 3, 2, 2, 2, 8, 9, 7, 184, 2,
	2, 9, 10, 9, 2, 2, 2, 10, 46, 5, 4, 3, 2, 11, 12, 7, 184, 2, 2, 12, 13,
	7, 193, 2, 2, 13, 46, 5, 6, 4, 2, 14, 15, 7, 184, 2, 2, 15, 16, 7, 194,
	2, 2, 16, 46, 5, 6, 4, 2, 17, 18, 7, 184, 2, 2, 18, 46, 7, 195, 2, 2, 19,
	20, 7, 184, 2, 2, 20, 46, 7, 199, 2, 2, 21, 22, 7, 184, 2, 2, 22, 23, 7,
	197, 2, 2, 23, 46, 7, 216, 2, 2, 24, 25, 7, 184, 2, 2, 25, 26, 7, 198,
	2, 2, 26, 46, 7, 216, 2, 2, 27, 28, 7, 184, 2, 2, 28, 29, 7, 196, 2, 2,
	29, 46, 7, 216, 2, 2, 30, 31, 7, 184, 2, 2, 31, 32, 7, 190, 2, 2, 32, 46,
	5, 4, 3, 2, 33, 34, 7, 184, 2, 2, 34, 35, 7, 202, 2, 2, 35, 46, 5, 4, 3,
	2, 36, 37, 7, 184, 2, 2, 37, 38, 7, 203, 2, 2, 38, 46, 5, 4, 3, 2, 39,
	40, 7, 184, 2, 2, 40, 41, 7, 191, 2, 2, 41, 43, 7, 216, 2, 2, 42, 44, 5,
	4, 3, 2, 43, 42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45,
	8, 3, 2, 2, 2, 45, 11, 3, 2, 2, 2, 45, 14, 3, 2, 2, 2, 45, 17, 3, 2, 2,
	2, 45, 19, 3, 2, 2, 2, 45, 21, 3, 2, 2, 2, 45, 24, 3, 2, 2, 2, 45, 27,
	3, 2, 2, 2, 45, 30, 3, 2, 2, 2, 45, 33, 3, 2, 2, 2, 45, 36, 3, 2, 2, 2,
	45, 39, 3, 2, 2, 2, 46, 3, 3, 2, 2, 2, 47, 49, 9, 3, 2, 2, 48, 47, 3, 2,
	2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 5,
	3, 2, 2, 2, 52, 53, 8, 4, 1, 2, 53, 78, 7, 200, 2, 2, 54, 78, 7, 201, 2,
	2, 55, 78, 7, 217, 2, 2, 56, 78, 7, 215, 2, 2, 57, 62, 7, 216, 2, 2, 58,
	59, 7, 205, 2, 2, 59, 60, 5, 6, 4, 2, 60, 61, 7, 206, 2, 2, 61, 63, 3,
	2, 2, 2, 62, 58, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 78, 3, 2, 2, 2, 64,
	65, 7, 205, 2, 2, 65, 66, 5, 6, 4, 2, 66, 67, 7, 206, 2, 2, 67, 78, 3,
	2, 2, 2, 68, 69, 7, 204, 2, 2, 69, 78, 5, 6, 4, 8, 70, 75, 7, 192, 2, 2,
	71, 76, 7, 216, 2, 2, 72, 73, 7, 205, 2, 2, 73, 74, 7, 216, 2, 2, 74, 76,
	7, 206, 2, 2, 75, 71, 3, 2, 2, 2, 75, 72, 3, 2, 2, 2, 76, 78, 3, 2, 2,
	2, 77, 52, 3, 2, 2, 2, 77, 54, 3, 2, 2, 2, 77, 55, 3, 2, 2, 2, 77, 56,
	3, 2, 2, 2, 77, 57, 3, 2, 2, 2, 77, 64, 3, 2, 2, 2, 77, 68, 3, 2, 2, 2,
	77, 70, 3, 2, 2, 2, 78, 93, 3, 2, 2, 2, 79, 80, 12, 7, 2, 2, 80, 81, 9,
	4, 2, 2, 81, 92, 5, 6, 4, 8, 82, 83, 12, 6, 2, 2, 83, 84, 7, 209, 2, 2,
	84, 92, 5, 6, 4, 7, 85, 86, 12, 5, 2, 2, 86, 87, 7, 210, 2, 2, 87, 92,
	5, 6, 4, 6, 88, 89, 12, 4, 2, 2, 89, 90, 9, 5, 2, 2, 90, 92, 5, 6, 4, 5,
	91, 79, 3, 2, 2, 2, 91, 82, 3, 2, 2, 2, 91, 85, 3, 2, 2, 2, 91, 88, 3,
	2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94,
	7, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 10, 43, 45, 50, 62, 75, 77, 91, 93,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'auto'", "'break'", "'case'", "'char'", "'const'", "'continue'", "'default'",
	"'do'", "'double'", "", "'enum'", "'extern'", "'float'", "'for'", "'goto'",
	"", "'inline'", "'int'", "'long'", "'register'", "'restrict'", "'return'",
	"'short'", "'signed'", "'sizeof'", "'static'", "'struct'", "'switch'",
	"'typedef'", "'union'", "'unsigned'", "'void'", "'volatile'", "'while'",
	"'_Bool'", "'_Complex'", "'_Imaginery'", "'true'", "'false'", "'BOOL'",
	"'Class'", "'bycopy'", "'byref'", "'id'", "'IMP'", "'in'", "'inout'", "'nil'",
	"'NO'", "'NULL'", "'oneway'", "'out'", "'Protocol'", "'SEL'", "'self'",
	"'super'", "'YES'", "'@autoreleasepool'", "'@catch'", "'@class'", "'@dynamic'",
	"'@encode'", "'@end'", "'@finally'", "'@implementation'", "'@interface'",
	"'@import'", "'@package'", "'@protocol'", "'@optional'", "'@private'",
	"'@property'", "'@protected'", "'@public'", "'@required'", "'@selector'",
	"'@synchronized'", "'@synthesize'", "'@throw'", "'@try'", "'atomic'", "'nonatomic'",
	"'retain'", "'__attribute__'", "'__autoreleasing'", "'__block'", "'__bridge'",
	"'__bridge_retained'", "'__bridge_transfer'", "'__covariant'", "'__contravariant'",
	"'__deprecated'", "'__kindof'", "'__strong'", "", "'__unsafe_unretained'",
	"'__unused'", "'__weak'", "", "", "", "'null_resettable'", "'NS_INLINE'",
	"'NS_ENUM'", "'NS_OPTIONS'", "'assign'", "'copy'", "'getter'", "'setter'",
	"'strong'", "'readonly'", "'readwrite'", "'weak'", "'unsafe_unretained'",
	"'IBOutlet'", "'IBOutletCollection'", "'IBInspectable'", "'IB_DESIGNABLE'",
	"", "", "", "", "", "'__TVOS_PROHIBITED'", "", "", "", "'{'", "'}'", "'['",
	"']'", "';'", "','", "'.'", "'->'", "'@'", "'='", "", "", "", "'~'", "'?'",
	"':'", "", "", "", "", "", "", "'++'", "'--'", "'+'", "'-'", "'*'", "'/'",
	"'&'", "'|'", "'^'", "'%'", "'+='", "'-='", "'*='", "'/='", "'&='", "'|='",
	"'^='", "'%='", "'<<='", "'>>='", "'...'", "", "", "", "", "", "", "",
	"", "", "", "'\\'", "", "", "", "", "", "", "", "", "'defined'", "", "'elif'",
	"", "'undef'", "'ifdef'", "'ifndef'", "'endif'",
}
var symbolicNames = []string{
	"", "AUTO", "BREAK", "CASE", "CHAR", "CONST", "CONTINUE", "DEFAULT", "DO",
	"DOUBLE", "ELSE", "ENUM", "EXTERN", "FLOAT", "FOR", "GOTO", "IF", "INLINE",
	"INT", "LONG", "REGISTER", "RESTRICT", "RETURN", "SHORT", "SIGNED", "SIZEOF",
	"STATIC", "STRUCT", "SWITCH", "TYPEDEF", "UNION", "UNSIGNED", "VOID", "VOLATILE",
	"WHILE", "BOOL_", "COMPLEX", "IMAGINERY", "TRUE", "FALSE", "BOOL", "Class",
	"BYCOPY", "BYREF", "ID", "IMP", "IN", "INOUT", "NIL", "NO", "NULL_", "ONEWAY",
	"OUT", "PROTOCOL_", "SEL", "SELF", "SUPER", "YES", "AUTORELEASEPOOL", "CATCH",
	"CLASS", "DYNAMIC", "ENCODE", "END", "FINALLY", "IMPLEMENTATION", "INTERFACE",
	"IMPORT", "PACKAGE", "PROTOCOL", "OPTIONAL", "PRIVATE", "PROPERTY", "PROTECTED",
	"PUBLIC", "REQUIRED", "SELECTOR", "SYNCHRONIZED", "SYNTHESIZE", "THROW",
	"TRY", "ATOMIC", "NONATOMIC", "RETAIN", "ATTRIBUTE", "AUTORELEASING_QUALIFIER",
	"BLOCK", "BRIDGE", "BRIDGE_RETAINED", "BRIDGE_TRANSFER", "COVARIANT", "CONTRAVARIANT",
	"DEPRECATED", "KINDOF", "STRONG_QUALIFIER", "TYPEOF", "UNSAFE_UNRETAINED_QUALIFIER",
	"UNUSED", "WEAK_QUALIFIER", "NULL_UNSPECIFIED", "NULLABLE", "NONNULL",
	"NULL_RESETTABLE", "NS_INLINE", "NS_ENUM", "NS_OPTIONS", "ASSIGN", "COPY",
	"GETTER", "SETTER", "STRONG", "READONLY", "READWRITE", "WEAK", "UNSAFE_UNRETAINED",
	"IB_OUTLET", "IB_OUTLET_COLLECTION", "IB_INSPECTABLE", "IB_DESIGNABLE",
	"NS_ASSUME_NONNULL_BEGIN", "NS_ASSUME_NONNULL_END", "EXTERN_SUFFIX", "IOS_SUFFIX",
	"MAC_SUFFIX", "TVOS_PROHIBITED", "IDENTIFIER", "LP", "RP", "LBRACE", "RBRACE",
	"LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "STRUCTACCESS", "AT", "ASSIGNMENT",
	"GT", "LT", "BANG", "TILDE", "QUESTION", "COLON", "EQUAL", "LE", "GE",
	"NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV", "BITAND",
	"BITOR", "BITXOR", "MOD", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN",
	"AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN",
	"RSHIFT_ASSIGN", "ELIPSIS", "CHARACTER_LITERAL", "STRING_START", "HEX_LITERAL",
	"OCTAL_LITERAL", "BINARY_LITERAL", "DECIMAL_LITERAL", "FLOATING_POINT_LITERAL",
	"WS", "MULTI_COMMENT", "SINGLE_COMMENT", "BACKSLASH", "SHARP", "STRING_NEWLINE",
	"STRING_END", "STRING_VALUE", "DIRECTIVE_IMPORT", "DIRECTIVE_INCLUDE",
	"DIRECTIVE_PRAGMA", "DIRECTIVE_DEFINE", "DIRECTIVE_DEFINED", "DIRECTIVE_IF",
	"DIRECTIVE_ELIF", "DIRECTIVE_ELSE", "DIRECTIVE_UNDEF", "DIRECTIVE_IFDEF",
	"DIRECTIVE_IFNDEF", "DIRECTIVE_ENDIF", "DIRECTIVE_TRUE", "DIRECTIVE_FALSE",
	"DIRECTIVE_ERROR", "DIRECTIVE_WARNING", "DIRECTIVE_BANG", "DIRECTIVE_LP",
	"DIRECTIVE_RP", "DIRECTIVE_EQUAL", "DIRECTIVE_NOTEQUAL", "DIRECTIVE_AND",
	"DIRECTIVE_OR", "DIRECTIVE_LT", "DIRECTIVE_GT", "DIRECTIVE_LE", "DIRECTIVE_GE",
	"DIRECTIVE_STRING", "DIRECTIVE_ID", "DIRECTIVE_DECIMAL_LITERAL", "DIRECTIVE_FLOAT",
	"DIRECTIVE_NEWLINE", "DIRECTIVE_MULTI_COMMENT", "DIRECTIVE_SINGLE_COMMENT",
	"DIRECTIVE_BACKSLASH_NEWLINE", "DIRECTIVE_TEXT_NEWLINE", "DIRECTIVE_TEXT",
}

var ruleNames = []string{
	"directive", "directiveText", "preprocessorExpression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ObjectiveCPreprocessorParser struct {
	*antlr.BaseParser
}

func NewObjectiveCPreprocessorParser(input antlr.TokenStream) *ObjectiveCPreprocessorParser {
	this := new(ObjectiveCPreprocessorParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ObjectiveCPreprocessorParser.g4"

	return this
}

// ObjectiveCPreprocessorParser tokens.
const (
	ObjectiveCPreprocessorParserEOF                         = antlr.TokenEOF
	ObjectiveCPreprocessorParserAUTO                        = 1
	ObjectiveCPreprocessorParserBREAK                       = 2
	ObjectiveCPreprocessorParserCASE                        = 3
	ObjectiveCPreprocessorParserCHAR                        = 4
	ObjectiveCPreprocessorParserCONST                       = 5
	ObjectiveCPreprocessorParserCONTINUE                    = 6
	ObjectiveCPreprocessorParserDEFAULT                     = 7
	ObjectiveCPreprocessorParserDO                          = 8
	ObjectiveCPreprocessorParserDOUBLE                      = 9
	ObjectiveCPreprocessorParserELSE                        = 10
	ObjectiveCPreprocessorParserENUM                        = 11
	ObjectiveCPreprocessorParserEXTERN                      = 12
	ObjectiveCPreprocessorParserFLOAT                       = 13
	ObjectiveCPreprocessorParserFOR                         = 14
	ObjectiveCPreprocessorParserGOTO                        = 15
	ObjectiveCPreprocessorParserIF                          = 16
	ObjectiveCPreprocessorParserINLINE                      = 17
	ObjectiveCPreprocessorParserINT                         = 18
	ObjectiveCPreprocessorParserLONG                        = 19
	ObjectiveCPreprocessorParserREGISTER                    = 20
	ObjectiveCPreprocessorParserRESTRICT                    = 21
	ObjectiveCPreprocessorParserRETURN                      = 22
	ObjectiveCPreprocessorParserSHORT                       = 23
	ObjectiveCPreprocessorParserSIGNED                      = 24
	ObjectiveCPreprocessorParserSIZEOF                      = 25
	ObjectiveCPreprocessorParserSTATIC                      = 26
	ObjectiveCPreprocessorParserSTRUCT                      = 27
	ObjectiveCPreprocessorParserSWITCH                      = 28
	ObjectiveCPreprocessorParserTYPEDEF                     = 29
	ObjectiveCPreprocessorParserUNION                       = 30
	ObjectiveCPreprocessorParserUNSIGNED                    = 31
	ObjectiveCPreprocessorParserVOID                        = 32
	ObjectiveCPreprocessorParserVOLATILE                    = 33
	ObjectiveCPreprocessorParserWHILE                       = 34
	ObjectiveCPreprocessorParserBOOL_                       = 35
	ObjectiveCPreprocessorParserCOMPLEX                     = 36
	ObjectiveCPreprocessorParserIMAGINERY                   = 37
	ObjectiveCPreprocessorParserTRUE                        = 38
	ObjectiveCPreprocessorParserFALSE                       = 39
	ObjectiveCPreprocessorParserBOOL                        = 40
	ObjectiveCPreprocessorParserClass                       = 41
	ObjectiveCPreprocessorParserBYCOPY                      = 42
	ObjectiveCPreprocessorParserBYREF                       = 43
	ObjectiveCPreprocessorParserID                          = 44
	ObjectiveCPreprocessorParserIMP                         = 45
	ObjectiveCPreprocessorParserIN                          = 46
	ObjectiveCPreprocessorParserINOUT                       = 47
	ObjectiveCPreprocessorParserNIL                         = 48
	ObjectiveCPreprocessorParserNO                          = 49
	ObjectiveCPreprocessorParserNULL_                       = 50
	ObjectiveCPreprocessorParserONEWAY                      = 51
	ObjectiveCPreprocessorParserOUT                         = 52
	ObjectiveCPreprocessorParserPROTOCOL_                   = 53
	ObjectiveCPreprocessorParserSEL                         = 54
	ObjectiveCPreprocessorParserSELF                        = 55
	ObjectiveCPreprocessorParserSUPER                       = 56
	ObjectiveCPreprocessorParserYES                         = 57
	ObjectiveCPreprocessorParserAUTORELEASEPOOL             = 58
	ObjectiveCPreprocessorParserCATCH                       = 59
	ObjectiveCPreprocessorParserCLASS                       = 60
	ObjectiveCPreprocessorParserDYNAMIC                     = 61
	ObjectiveCPreprocessorParserENCODE                      = 62
	ObjectiveCPreprocessorParserEND                         = 63
	ObjectiveCPreprocessorParserFINALLY                     = 64
	ObjectiveCPreprocessorParserIMPLEMENTATION              = 65
	ObjectiveCPreprocessorParserINTERFACE                   = 66
	ObjectiveCPreprocessorParserIMPORT                      = 67
	ObjectiveCPreprocessorParserPACKAGE                     = 68
	ObjectiveCPreprocessorParserPROTOCOL                    = 69
	ObjectiveCPreprocessorParserOPTIONAL                    = 70
	ObjectiveCPreprocessorParserPRIVATE                     = 71
	ObjectiveCPreprocessorParserPROPERTY                    = 72
	ObjectiveCPreprocessorParserPROTECTED                   = 73
	ObjectiveCPreprocessorParserPUBLIC                      = 74
	ObjectiveCPreprocessorParserREQUIRED                    = 75
	ObjectiveCPreprocessorParserSELECTOR                    = 76
	ObjectiveCPreprocessorParserSYNCHRONIZED                = 77
	ObjectiveCPreprocessorParserSYNTHESIZE                  = 78
	ObjectiveCPreprocessorParserTHROW                       = 79
	ObjectiveCPreprocessorParserTRY                         = 80
	ObjectiveCPreprocessorParserATOMIC                      = 81
	ObjectiveCPreprocessorParserNONATOMIC                   = 82
	ObjectiveCPreprocessorParserRETAIN                      = 83
	ObjectiveCPreprocessorParserATTRIBUTE                   = 84
	ObjectiveCPreprocessorParserAUTORELEASING_QUALIFIER     = 85
	ObjectiveCPreprocessorParserBLOCK                       = 86
	ObjectiveCPreprocessorParserBRIDGE                      = 87
	ObjectiveCPreprocessorParserBRIDGE_RETAINED             = 88
	ObjectiveCPreprocessorParserBRIDGE_TRANSFER             = 89
	ObjectiveCPreprocessorParserCOVARIANT                   = 90
	ObjectiveCPreprocessorParserCONTRAVARIANT               = 91
	ObjectiveCPreprocessorParserDEPRECATED                  = 92
	ObjectiveCPreprocessorParserKINDOF                      = 93
	ObjectiveCPreprocessorParserSTRONG_QUALIFIER            = 94
	ObjectiveCPreprocessorParserTYPEOF                      = 95
	ObjectiveCPreprocessorParserUNSAFE_UNRETAINED_QUALIFIER = 96
	ObjectiveCPreprocessorParserUNUSED                      = 97
	ObjectiveCPreprocessorParserWEAK_QUALIFIER              = 98
	ObjectiveCPreprocessorParserNULL_UNSPECIFIED            = 99
	ObjectiveCPreprocessorParserNULLABLE                    = 100
	ObjectiveCPreprocessorParserNONNULL                     = 101
	ObjectiveCPreprocessorParserNULL_RESETTABLE             = 102
	ObjectiveCPreprocessorParserNS_INLINE                   = 103
	ObjectiveCPreprocessorParserNS_ENUM                     = 104
	ObjectiveCPreprocessorParserNS_OPTIONS                  = 105
	ObjectiveCPreprocessorParserASSIGN                      = 106
	ObjectiveCPreprocessorParserCOPY                        = 107
	ObjectiveCPreprocessorParserGETTER                      = 108
	ObjectiveCPreprocessorParserSETTER                      = 109
	ObjectiveCPreprocessorParserSTRONG                      = 110
	ObjectiveCPreprocessorParserREADONLY                    = 111
	ObjectiveCPreprocessorParserREADWRITE                   = 112
	ObjectiveCPreprocessorParserWEAK                        = 113
	ObjectiveCPreprocessorParserUNSAFE_UNRETAINED           = 114
	ObjectiveCPreprocessorParserIB_OUTLET                   = 115
	ObjectiveCPreprocessorParserIB_OUTLET_COLLECTION        = 116
	ObjectiveCPreprocessorParserIB_INSPECTABLE              = 117
	ObjectiveCPreprocessorParserIB_DESIGNABLE               = 118
	ObjectiveCPreprocessorParserNS_ASSUME_NONNULL_BEGIN     = 119
	ObjectiveCPreprocessorParserNS_ASSUME_NONNULL_END       = 120
	ObjectiveCPreprocessorParserEXTERN_SUFFIX               = 121
	ObjectiveCPreprocessorParserIOS_SUFFIX                  = 122
	ObjectiveCPreprocessorParserMAC_SUFFIX                  = 123
	ObjectiveCPreprocessorParserTVOS_PROHIBITED             = 124
	ObjectiveCPreprocessorParserIDENTIFIER                  = 125
	ObjectiveCPreprocessorParserLP                          = 126
	ObjectiveCPreprocessorParserRP                          = 127
	ObjectiveCPreprocessorParserLBRACE                      = 128
	ObjectiveCPreprocessorParserRBRACE                      = 129
	ObjectiveCPreprocessorParserLBRACK                      = 130
	ObjectiveCPreprocessorParserRBRACK                      = 131
	ObjectiveCPreprocessorParserSEMI                        = 132
	ObjectiveCPreprocessorParserCOMMA                       = 133
	ObjectiveCPreprocessorParserDOT                         = 134
	ObjectiveCPreprocessorParserSTRUCTACCESS                = 135
	ObjectiveCPreprocessorParserAT                          = 136
	ObjectiveCPreprocessorParserASSIGNMENT                  = 137
	ObjectiveCPreprocessorParserGT                          = 138
	ObjectiveCPreprocessorParserLT                          = 139
	ObjectiveCPreprocessorParserBANG                        = 140
	ObjectiveCPreprocessorParserTILDE                       = 141
	ObjectiveCPreprocessorParserQUESTION                    = 142
	ObjectiveCPreprocessorParserCOLON                       = 143
	ObjectiveCPreprocessorParserEQUAL                       = 144
	ObjectiveCPreprocessorParserLE                          = 145
	ObjectiveCPreprocessorParserGE                          = 146
	ObjectiveCPreprocessorParserNOTEQUAL                    = 147
	ObjectiveCPreprocessorParserAND                         = 148
	ObjectiveCPreprocessorParserOR                          = 149
	ObjectiveCPreprocessorParserINC                         = 150
	ObjectiveCPreprocessorParserDEC                         = 151
	ObjectiveCPreprocessorParserADD                         = 152
	ObjectiveCPreprocessorParserSUB                         = 153
	ObjectiveCPreprocessorParserMUL                         = 154
	ObjectiveCPreprocessorParserDIV                         = 155
	ObjectiveCPreprocessorParserBITAND                      = 156
	ObjectiveCPreprocessorParserBITOR                       = 157
	ObjectiveCPreprocessorParserBITXOR                      = 158
	ObjectiveCPreprocessorParserMOD                         = 159
	ObjectiveCPreprocessorParserADD_ASSIGN                  = 160
	ObjectiveCPreprocessorParserSUB_ASSIGN                  = 161
	ObjectiveCPreprocessorParserMUL_ASSIGN                  = 162
	ObjectiveCPreprocessorParserDIV_ASSIGN                  = 163
	ObjectiveCPreprocessorParserAND_ASSIGN                  = 164
	ObjectiveCPreprocessorParserOR_ASSIGN                   = 165
	ObjectiveCPreprocessorParserXOR_ASSIGN                  = 166
	ObjectiveCPreprocessorParserMOD_ASSIGN                  = 167
	ObjectiveCPreprocessorParserLSHIFT_ASSIGN               = 168
	ObjectiveCPreprocessorParserRSHIFT_ASSIGN               = 169
	ObjectiveCPreprocessorParserELIPSIS                     = 170
	ObjectiveCPreprocessorParserCHARACTER_LITERAL           = 171
	ObjectiveCPreprocessorParserSTRING_START                = 172
	ObjectiveCPreprocessorParserHEX_LITERAL                 = 173
	ObjectiveCPreprocessorParserOCTAL_LITERAL               = 174
	ObjectiveCPreprocessorParserBINARY_LITERAL              = 175
	ObjectiveCPreprocessorParserDECIMAL_LITERAL             = 176
	ObjectiveCPreprocessorParserFLOATING_POINT_LITERAL      = 177
	ObjectiveCPreprocessorParserWS                          = 178
	ObjectiveCPreprocessorParserMULTI_COMMENT               = 179
	ObjectiveCPreprocessorParserSINGLE_COMMENT              = 180
	ObjectiveCPreprocessorParserBACKSLASH                   = 181
	ObjectiveCPreprocessorParserSHARP                       = 182
	ObjectiveCPreprocessorParserSTRING_NEWLINE              = 183
	ObjectiveCPreprocessorParserSTRING_END                  = 184
	ObjectiveCPreprocessorParserSTRING_VALUE                = 185
	ObjectiveCPreprocessorParserDIRECTIVE_IMPORT            = 186
	ObjectiveCPreprocessorParserDIRECTIVE_INCLUDE           = 187
	ObjectiveCPreprocessorParserDIRECTIVE_PRAGMA            = 188
	ObjectiveCPreprocessorParserDIRECTIVE_DEFINE            = 189
	ObjectiveCPreprocessorParserDIRECTIVE_DEFINED           = 190
	ObjectiveCPreprocessorParserDIRECTIVE_IF                = 191
	ObjectiveCPreprocessorParserDIRECTIVE_ELIF              = 192
	ObjectiveCPreprocessorParserDIRECTIVE_ELSE              = 193
	ObjectiveCPreprocessorParserDIRECTIVE_UNDEF             = 194
	ObjectiveCPreprocessorParserDIRECTIVE_IFDEF             = 195
	ObjectiveCPreprocessorParserDIRECTIVE_IFNDEF            = 196
	ObjectiveCPreprocessorParserDIRECTIVE_ENDIF             = 197
	ObjectiveCPreprocessorParserDIRECTIVE_TRUE              = 198
	ObjectiveCPreprocessorParserDIRECTIVE_FALSE             = 199
	ObjectiveCPreprocessorParserDIRECTIVE_ERROR             = 200
	ObjectiveCPreprocessorParserDIRECTIVE_WARNING           = 201
	ObjectiveCPreprocessorParserDIRECTIVE_BANG              = 202
	ObjectiveCPreprocessorParserDIRECTIVE_LP                = 203
	ObjectiveCPreprocessorParserDIRECTIVE_RP                = 204
	ObjectiveCPreprocessorParserDIRECTIVE_EQUAL             = 205
	ObjectiveCPreprocessorParserDIRECTIVE_NOTEQUAL          = 206
	ObjectiveCPreprocessorParserDIRECTIVE_AND               = 207
	ObjectiveCPreprocessorParserDIRECTIVE_OR                = 208
	ObjectiveCPreprocessorParserDIRECTIVE_LT                = 209
	ObjectiveCPreprocessorParserDIRECTIVE_GT                = 210
	ObjectiveCPreprocessorParserDIRECTIVE_LE                = 211
	ObjectiveCPreprocessorParserDIRECTIVE_GE                = 212
	ObjectiveCPreprocessorParserDIRECTIVE_STRING            = 213
	ObjectiveCPreprocessorParserDIRECTIVE_ID                = 214
	ObjectiveCPreprocessorParserDIRECTIVE_DECIMAL_LITERAL   = 215
	ObjectiveCPreprocessorParserDIRECTIVE_FLOAT             = 216
	ObjectiveCPreprocessorParserDIRECTIVE_NEWLINE           = 217
	ObjectiveCPreprocessorParserDIRECTIVE_MULTI_COMMENT     = 218
	ObjectiveCPreprocessorParserDIRECTIVE_SINGLE_COMMENT    = 219
	ObjectiveCPreprocessorParserDIRECTIVE_BACKSLASH_NEWLINE = 220
	ObjectiveCPreprocessorParserDIRECTIVE_TEXT_NEWLINE      = 221
	ObjectiveCPreprocessorParserDIRECTIVE_TEXT              = 222
)

// ObjectiveCPreprocessorParser rules.
const (
	ObjectiveCPreprocessorParserRULE_directive              = 0
	ObjectiveCPreprocessorParserRULE_directiveText          = 1
	ObjectiveCPreprocessorParserRULE_preprocessorExpression = 2
)

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectiveCPreprocessorParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectiveCPreprocessorParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) CopyFrom(ctx *DirectiveContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PreprocessorDefContext struct {
	*DirectiveContext
}

func NewPreprocessorDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorDefContext {
	var p = new(PreprocessorDefContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PreprocessorDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorDefContext) SHARP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserSHARP, 0)
}

func (s *PreprocessorDefContext) DIRECTIVE_IFDEF() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_IFDEF, 0)
}

func (s *PreprocessorDefContext) DIRECTIVE_ID() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ID, 0)
}

func (s *PreprocessorDefContext) DIRECTIVE_IFNDEF() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_IFNDEF, 0)
}

func (s *PreprocessorDefContext) DIRECTIVE_UNDEF() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_UNDEF, 0)
}

func (s *PreprocessorDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorDef(s)
	}
}

func (s *PreprocessorDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorDef(s)
	}
}

type PreprocessorErrorContext struct {
	*DirectiveContext
}

func NewPreprocessorErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorErrorContext {
	var p = new(PreprocessorErrorContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PreprocessorErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorErrorContext) SHARP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserSHARP, 0)
}

func (s *PreprocessorErrorContext) DIRECTIVE_ERROR() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ERROR, 0)
}

func (s *PreprocessorErrorContext) DirectiveText() IDirectiveTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveTextContext)
}

func (s *PreprocessorErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorError(s)
	}
}

func (s *PreprocessorErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorError(s)
	}
}

type PreprocessorConditionalContext struct {
	*DirectiveContext
}

func NewPreprocessorConditionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorConditionalContext {
	var p = new(PreprocessorConditionalContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PreprocessorConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorConditionalContext) SHARP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserSHARP, 0)
}

func (s *PreprocessorConditionalContext) DIRECTIVE_IF() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_IF, 0)
}

func (s *PreprocessorConditionalContext) PreprocessorExpression() IPreprocessorExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreprocessorExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreprocessorExpressionContext)
}

func (s *PreprocessorConditionalContext) DIRECTIVE_ELIF() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ELIF, 0)
}

func (s *PreprocessorConditionalContext) DIRECTIVE_ELSE() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ELSE, 0)
}

func (s *PreprocessorConditionalContext) DIRECTIVE_ENDIF() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ENDIF, 0)
}

func (s *PreprocessorConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorConditional(s)
	}
}

func (s *PreprocessorConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorConditional(s)
	}
}

type PreprocessorImportContext struct {
	*DirectiveContext
}

func NewPreprocessorImportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorImportContext {
	var p = new(PreprocessorImportContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PreprocessorImportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorImportContext) SHARP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserSHARP, 0)
}

func (s *PreprocessorImportContext) DirectiveText() IDirectiveTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveTextContext)
}

func (s *PreprocessorImportContext) DIRECTIVE_IMPORT() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_IMPORT, 0)
}

func (s *PreprocessorImportContext) DIRECTIVE_INCLUDE() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_INCLUDE, 0)
}

func (s *PreprocessorImportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorImport(s)
	}
}

func (s *PreprocessorImportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorImport(s)
	}
}

type PreprocessorPragmaContext struct {
	*DirectiveContext
}

func NewPreprocessorPragmaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorPragmaContext {
	var p = new(PreprocessorPragmaContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PreprocessorPragmaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorPragmaContext) SHARP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserSHARP, 0)
}

func (s *PreprocessorPragmaContext) DIRECTIVE_PRAGMA() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_PRAGMA, 0)
}

func (s *PreprocessorPragmaContext) DirectiveText() IDirectiveTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveTextContext)
}

func (s *PreprocessorPragmaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorPragma(s)
	}
}

func (s *PreprocessorPragmaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorPragma(s)
	}
}

type PreprocessorDefineContext struct {
	*DirectiveContext
}

func NewPreprocessorDefineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorDefineContext {
	var p = new(PreprocessorDefineContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PreprocessorDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorDefineContext) SHARP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserSHARP, 0)
}

func (s *PreprocessorDefineContext) DIRECTIVE_DEFINE() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_DEFINE, 0)
}

func (s *PreprocessorDefineContext) DIRECTIVE_ID() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ID, 0)
}

func (s *PreprocessorDefineContext) DirectiveText() IDirectiveTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveTextContext)
}

func (s *PreprocessorDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorDefine(s)
	}
}

func (s *PreprocessorDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorDefine(s)
	}
}

type PreprocessorWarningContext struct {
	*DirectiveContext
}

func NewPreprocessorWarningContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorWarningContext {
	var p = new(PreprocessorWarningContext)

	p.DirectiveContext = NewEmptyDirectiveContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DirectiveContext))

	return p
}

func (s *PreprocessorWarningContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorWarningContext) SHARP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserSHARP, 0)
}

func (s *PreprocessorWarningContext) DIRECTIVE_WARNING() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_WARNING, 0)
}

func (s *PreprocessorWarningContext) DirectiveText() IDirectiveTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveTextContext)
}

func (s *PreprocessorWarningContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorWarning(s)
	}
}

func (s *PreprocessorWarningContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorWarning(s)
	}
}

func (p *ObjectiveCPreprocessorParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ObjectiveCPreprocessorParserRULE_directive)
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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPreprocessorImportContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		p.SetState(7)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ObjectiveCPreprocessorParserDIRECTIVE_IMPORT || _la == ObjectiveCPreprocessorParserDIRECTIVE_INCLUDE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(8)
			p.DirectiveText()
		}

	case 2:
		localctx = NewPreprocessorConditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(10)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_IF)
		}
		{
			p.SetState(11)
			p.preprocessorExpression(0)
		}

	case 3:
		localctx = NewPreprocessorConditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(13)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ELIF)
		}
		{
			p.SetState(14)
			p.preprocessorExpression(0)
		}

	case 4:
		localctx = NewPreprocessorConditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(15)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(16)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ELSE)
		}

	case 5:
		localctx = NewPreprocessorConditionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(17)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(18)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ENDIF)
		}

	case 6:
		localctx = NewPreprocessorDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(19)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(20)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_IFDEF)
		}
		{
			p.SetState(21)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ID)
		}

	case 7:
		localctx = NewPreprocessorDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(22)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(23)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_IFNDEF)
		}
		{
			p.SetState(24)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ID)
		}

	case 8:
		localctx = NewPreprocessorDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(25)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(26)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_UNDEF)
		}
		{
			p.SetState(27)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ID)
		}

	case 9:
		localctx = NewPreprocessorPragmaContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(28)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(29)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_PRAGMA)
		}
		{
			p.SetState(30)
			p.DirectiveText()
		}

	case 10:
		localctx = NewPreprocessorErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(31)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(32)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ERROR)
		}
		{
			p.SetState(33)
			p.DirectiveText()
		}

	case 11:
		localctx = NewPreprocessorWarningContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(34)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(35)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_WARNING)
		}
		{
			p.SetState(36)
			p.DirectiveText()
		}

	case 12:
		localctx = NewPreprocessorDefineContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(37)
			p.Match(ObjectiveCPreprocessorParserSHARP)
		}
		{
			p.SetState(38)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_DEFINE)
		}
		{
			p.SetState(39)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ID)
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ObjectiveCPreprocessorParserDIRECTIVE_TEXT_NEWLINE || _la == ObjectiveCPreprocessorParserDIRECTIVE_TEXT {
			{
				p.SetState(40)
				p.DirectiveText()
			}

		}

	}

	return localctx
}

// IDirectiveTextContext is an interface to support dynamic dispatch.
type IDirectiveTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveTextContext differentiates from other interfaces.
	IsDirectiveTextContext()
}

type DirectiveTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveTextContext() *DirectiveTextContext {
	var p = new(DirectiveTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectiveCPreprocessorParserRULE_directiveText
	return p
}

func (*DirectiveTextContext) IsDirectiveTextContext() {}

func NewDirectiveTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveTextContext {
	var p = new(DirectiveTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectiveCPreprocessorParserRULE_directiveText

	return p
}

func (s *DirectiveTextContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveTextContext) AllDIRECTIVE_TEXT() []antlr.TerminalNode {
	return s.GetTokens(ObjectiveCPreprocessorParserDIRECTIVE_TEXT)
}

func (s *DirectiveTextContext) DIRECTIVE_TEXT(i int) antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_TEXT, i)
}

func (s *DirectiveTextContext) AllDIRECTIVE_TEXT_NEWLINE() []antlr.TerminalNode {
	return s.GetTokens(ObjectiveCPreprocessorParserDIRECTIVE_TEXT_NEWLINE)
}

func (s *DirectiveTextContext) DIRECTIVE_TEXT_NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_TEXT_NEWLINE, i)
}

func (s *DirectiveTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterDirectiveText(s)
	}
}

func (s *DirectiveTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitDirectiveText(s)
	}
}

func (p *ObjectiveCPreprocessorParser) DirectiveText() (localctx IDirectiveTextContext) {
	localctx = NewDirectiveTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ObjectiveCPreprocessorParserRULE_directiveText)
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ObjectiveCPreprocessorParserDIRECTIVE_TEXT_NEWLINE || _la == ObjectiveCPreprocessorParserDIRECTIVE_TEXT {
		p.SetState(45)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ObjectiveCPreprocessorParserDIRECTIVE_TEXT_NEWLINE || _la == ObjectiveCPreprocessorParserDIRECTIVE_TEXT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPreprocessorExpressionContext is an interface to support dynamic dispatch.
type IPreprocessorExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreprocessorExpressionContext differentiates from other interfaces.
	IsPreprocessorExpressionContext()
}

type PreprocessorExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreprocessorExpressionContext() *PreprocessorExpressionContext {
	var p = new(PreprocessorExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ObjectiveCPreprocessorParserRULE_preprocessorExpression
	return p
}

func (*PreprocessorExpressionContext) IsPreprocessorExpressionContext() {}

func NewPreprocessorExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreprocessorExpressionContext {
	var p = new(PreprocessorExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ObjectiveCPreprocessorParserRULE_preprocessorExpression

	return p
}

func (s *PreprocessorExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PreprocessorExpressionContext) CopyFrom(ctx *PreprocessorExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PreprocessorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PreprocessorParenthesisContext struct {
	*PreprocessorExpressionContext
}

func NewPreprocessorParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorParenthesisContext {
	var p = new(PreprocessorParenthesisContext)

	p.PreprocessorExpressionContext = NewEmptyPreprocessorExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PreprocessorExpressionContext))

	return p
}

func (s *PreprocessorParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorParenthesisContext) DIRECTIVE_LP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_LP, 0)
}

func (s *PreprocessorParenthesisContext) PreprocessorExpression() IPreprocessorExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreprocessorExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreprocessorExpressionContext)
}

func (s *PreprocessorParenthesisContext) DIRECTIVE_RP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_RP, 0)
}

func (s *PreprocessorParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorParenthesis(s)
	}
}

func (s *PreprocessorParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorParenthesis(s)
	}
}

type PreprocessorNotContext struct {
	*PreprocessorExpressionContext
}

func NewPreprocessorNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorNotContext {
	var p = new(PreprocessorNotContext)

	p.PreprocessorExpressionContext = NewEmptyPreprocessorExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PreprocessorExpressionContext))

	return p
}

func (s *PreprocessorNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorNotContext) DIRECTIVE_BANG() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_BANG, 0)
}

func (s *PreprocessorNotContext) PreprocessorExpression() IPreprocessorExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreprocessorExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreprocessorExpressionContext)
}

func (s *PreprocessorNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorNot(s)
	}
}

func (s *PreprocessorNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorNot(s)
	}
}

type PreprocessorBinaryContext struct {
	*PreprocessorExpressionContext
	op antlr.Token
}

func NewPreprocessorBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorBinaryContext {
	var p = new(PreprocessorBinaryContext)

	p.PreprocessorExpressionContext = NewEmptyPreprocessorExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PreprocessorExpressionContext))

	return p
}

func (s *PreprocessorBinaryContext) GetOp() antlr.Token { return s.op }

func (s *PreprocessorBinaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *PreprocessorBinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorBinaryContext) AllPreprocessorExpression() []IPreprocessorExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPreprocessorExpressionContext)(nil)).Elem())
	var tst = make([]IPreprocessorExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPreprocessorExpressionContext)
		}
	}

	return tst
}

func (s *PreprocessorBinaryContext) PreprocessorExpression(i int) IPreprocessorExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreprocessorExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPreprocessorExpressionContext)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_EQUAL() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_EQUAL, 0)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_NOTEQUAL, 0)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_AND() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_AND, 0)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_OR() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_OR, 0)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_LT() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_LT, 0)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_GT() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_GT, 0)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_LE() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_LE, 0)
}

func (s *PreprocessorBinaryContext) DIRECTIVE_GE() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_GE, 0)
}

func (s *PreprocessorBinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorBinary(s)
	}
}

func (s *PreprocessorBinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorBinary(s)
	}
}

type PreprocessorConstantContext struct {
	*PreprocessorExpressionContext
}

func NewPreprocessorConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorConstantContext {
	var p = new(PreprocessorConstantContext)

	p.PreprocessorExpressionContext = NewEmptyPreprocessorExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PreprocessorExpressionContext))

	return p
}

func (s *PreprocessorConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorConstantContext) DIRECTIVE_TRUE() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_TRUE, 0)
}

func (s *PreprocessorConstantContext) DIRECTIVE_FALSE() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_FALSE, 0)
}

func (s *PreprocessorConstantContext) DIRECTIVE_DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_DECIMAL_LITERAL, 0)
}

func (s *PreprocessorConstantContext) DIRECTIVE_STRING() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_STRING, 0)
}

func (s *PreprocessorConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorConstant(s)
	}
}

func (s *PreprocessorConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorConstant(s)
	}
}

type PreprocessorConditionalSymbolContext struct {
	*PreprocessorExpressionContext
}

func NewPreprocessorConditionalSymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorConditionalSymbolContext {
	var p = new(PreprocessorConditionalSymbolContext)

	p.PreprocessorExpressionContext = NewEmptyPreprocessorExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PreprocessorExpressionContext))

	return p
}

func (s *PreprocessorConditionalSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorConditionalSymbolContext) DIRECTIVE_ID() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ID, 0)
}

func (s *PreprocessorConditionalSymbolContext) DIRECTIVE_LP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_LP, 0)
}

func (s *PreprocessorConditionalSymbolContext) PreprocessorExpression() IPreprocessorExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreprocessorExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreprocessorExpressionContext)
}

func (s *PreprocessorConditionalSymbolContext) DIRECTIVE_RP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_RP, 0)
}

func (s *PreprocessorConditionalSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorConditionalSymbol(s)
	}
}

func (s *PreprocessorConditionalSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorConditionalSymbol(s)
	}
}

type PreprocessorDefinedContext struct {
	*PreprocessorExpressionContext
}

func NewPreprocessorDefinedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreprocessorDefinedContext {
	var p = new(PreprocessorDefinedContext)

	p.PreprocessorExpressionContext = NewEmptyPreprocessorExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PreprocessorExpressionContext))

	return p
}

func (s *PreprocessorDefinedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreprocessorDefinedContext) DIRECTIVE_DEFINED() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_DEFINED, 0)
}

func (s *PreprocessorDefinedContext) DIRECTIVE_ID() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_ID, 0)
}

func (s *PreprocessorDefinedContext) DIRECTIVE_LP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_LP, 0)
}

func (s *PreprocessorDefinedContext) DIRECTIVE_RP() antlr.TerminalNode {
	return s.GetToken(ObjectiveCPreprocessorParserDIRECTIVE_RP, 0)
}

func (s *PreprocessorDefinedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.EnterPreprocessorDefined(s)
	}
}

func (s *PreprocessorDefinedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjectiveCPreprocessorParserListener); ok {
		listenerT.ExitPreprocessorDefined(s)
	}
}

func (p *ObjectiveCPreprocessorParser) PreprocessorExpression() (localctx IPreprocessorExpressionContext) {
	return p.preprocessorExpression(0)
}

func (p *ObjectiveCPreprocessorParser) preprocessorExpression(_p int) (localctx IPreprocessorExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPreprocessorExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPreprocessorExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, ObjectiveCPreprocessorParserRULE_preprocessorExpression, _p)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ObjectiveCPreprocessorParserDIRECTIVE_TRUE:
		localctx = NewPreprocessorConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(51)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_TRUE)
		}

	case ObjectiveCPreprocessorParserDIRECTIVE_FALSE:
		localctx = NewPreprocessorConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_FALSE)
		}

	case ObjectiveCPreprocessorParserDIRECTIVE_DECIMAL_LITERAL:
		localctx = NewPreprocessorConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_DECIMAL_LITERAL)
		}

	case ObjectiveCPreprocessorParserDIRECTIVE_STRING:
		localctx = NewPreprocessorConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_STRING)
		}

	case ObjectiveCPreprocessorParserDIRECTIVE_ID:
		localctx = NewPreprocessorConditionalSymbolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ID)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(56)
				p.Match(ObjectiveCPreprocessorParserDIRECTIVE_LP)
			}
			{
				p.SetState(57)
				p.preprocessorExpression(0)
			}
			{
				p.SetState(58)
				p.Match(ObjectiveCPreprocessorParserDIRECTIVE_RP)
			}

		}

	case ObjectiveCPreprocessorParserDIRECTIVE_LP:
		localctx = NewPreprocessorParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_LP)
		}
		{
			p.SetState(63)
			p.preprocessorExpression(0)
		}
		{
			p.SetState(64)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_RP)
		}

	case ObjectiveCPreprocessorParserDIRECTIVE_BANG:
		localctx = NewPreprocessorNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_BANG)
		}
		{
			p.SetState(67)
			p.preprocessorExpression(6)
		}

	case ObjectiveCPreprocessorParserDIRECTIVE_DEFINED:
		localctx = NewPreprocessorDefinedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(ObjectiveCPreprocessorParserDIRECTIVE_DEFINED)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ObjectiveCPreprocessorParserDIRECTIVE_ID:
			{
				p.SetState(69)
				p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ID)
			}

		case ObjectiveCPreprocessorParserDIRECTIVE_LP:
			{
				p.SetState(70)
				p.Match(ObjectiveCPreprocessorParserDIRECTIVE_LP)
			}
			{
				p.SetState(71)
				p.Match(ObjectiveCPreprocessorParserDIRECTIVE_ID)
			}
			{
				p.SetState(72)
				p.Match(ObjectiveCPreprocessorParserDIRECTIVE_RP)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPreprocessorBinaryContext(p, NewPreprocessorExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ObjectiveCPreprocessorParserRULE_preprocessorExpression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(78)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*PreprocessorBinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ObjectiveCPreprocessorParserDIRECTIVE_EQUAL || _la == ObjectiveCPreprocessorParserDIRECTIVE_NOTEQUAL) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*PreprocessorBinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(79)
					p.preprocessorExpression(6)
				}

			case 2:
				localctx = NewPreprocessorBinaryContext(p, NewPreprocessorExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ObjectiveCPreprocessorParserRULE_preprocessorExpression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(81)

					var _m = p.Match(ObjectiveCPreprocessorParserDIRECTIVE_AND)

					localctx.(*PreprocessorBinaryContext).op = _m
				}
				{
					p.SetState(82)
					p.preprocessorExpression(5)
				}

			case 3:
				localctx = NewPreprocessorBinaryContext(p, NewPreprocessorExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ObjectiveCPreprocessorParserRULE_preprocessorExpression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(84)

					var _m = p.Match(ObjectiveCPreprocessorParserDIRECTIVE_OR)

					localctx.(*PreprocessorBinaryContext).op = _m
				}
				{
					p.SetState(85)
					p.preprocessorExpression(4)
				}

			case 4:
				localctx = NewPreprocessorBinaryContext(p, NewPreprocessorExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ObjectiveCPreprocessorParserRULE_preprocessorExpression)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(87)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*PreprocessorBinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-209)&-(0x1f+1)) == 0 && ((1<<uint((_la-209)))&((1<<(ObjectiveCPreprocessorParserDIRECTIVE_LT-209))|(1<<(ObjectiveCPreprocessorParserDIRECTIVE_GT-209))|(1<<(ObjectiveCPreprocessorParserDIRECTIVE_LE-209))|(1<<(ObjectiveCPreprocessorParserDIRECTIVE_GE-209)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*PreprocessorBinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(88)
					p.preprocessorExpression(3)
				}

			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ObjectiveCPreprocessorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *PreprocessorExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PreprocessorExpressionContext)
		}
		return p.PreprocessorExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ObjectiveCPreprocessorParser) PreprocessorExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
