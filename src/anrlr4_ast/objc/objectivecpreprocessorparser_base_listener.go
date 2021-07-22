// Generated from ObjectiveCPreprocessorParser.g4 by ANTLR 4.7.

package parser // ObjectiveCPreprocessorParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseObjectiveCPreprocessorParserListener is a complete listener for a parse tree produced by ObjectiveCPreprocessorParser.
type BaseObjectiveCPreprocessorParserListener struct{}

var _ ObjectiveCPreprocessorParserListener = &BaseObjectiveCPreprocessorParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObjectiveCPreprocessorParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObjectiveCPreprocessorParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPreprocessorImport is called when production preprocessorImport is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorImport(ctx *PreprocessorImportContext) {
}

// ExitPreprocessorImport is called when production preprocessorImport is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorImport(ctx *PreprocessorImportContext) {
}

// EnterPreprocessorConditional is called when production preprocessorConditional is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorConditional(ctx *PreprocessorConditionalContext) {
}

// ExitPreprocessorConditional is called when production preprocessorConditional is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorConditional(ctx *PreprocessorConditionalContext) {
}

// EnterPreprocessorDef is called when production preprocessorDef is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorDef(ctx *PreprocessorDefContext) {
}

// ExitPreprocessorDef is called when production preprocessorDef is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorDef(ctx *PreprocessorDefContext) {}

// EnterPreprocessorPragma is called when production preprocessorPragma is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorPragma(ctx *PreprocessorPragmaContext) {
}

// ExitPreprocessorPragma is called when production preprocessorPragma is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorPragma(ctx *PreprocessorPragmaContext) {
}

// EnterPreprocessorError is called when production preprocessorError is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorError(ctx *PreprocessorErrorContext) {
}

// ExitPreprocessorError is called when production preprocessorError is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorError(ctx *PreprocessorErrorContext) {
}

// EnterPreprocessorWarning is called when production preprocessorWarning is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorWarning(ctx *PreprocessorWarningContext) {
}

// ExitPreprocessorWarning is called when production preprocessorWarning is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorWarning(ctx *PreprocessorWarningContext) {
}

// EnterPreprocessorDefine is called when production preprocessorDefine is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorDefine(ctx *PreprocessorDefineContext) {
}

// ExitPreprocessorDefine is called when production preprocessorDefine is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorDefine(ctx *PreprocessorDefineContext) {
}

// EnterDirectiveText is called when production directiveText is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterDirectiveText(ctx *DirectiveTextContext) {}

// ExitDirectiveText is called when production directiveText is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitDirectiveText(ctx *DirectiveTextContext) {}

// EnterPreprocessorParenthesis is called when production preprocessorParenthesis is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorParenthesis(ctx *PreprocessorParenthesisContext) {
}

// ExitPreprocessorParenthesis is called when production preprocessorParenthesis is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorParenthesis(ctx *PreprocessorParenthesisContext) {
}

// EnterPreprocessorNot is called when production preprocessorNot is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorNot(ctx *PreprocessorNotContext) {
}

// ExitPreprocessorNot is called when production preprocessorNot is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorNot(ctx *PreprocessorNotContext) {}

// EnterPreprocessorBinary is called when production preprocessorBinary is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorBinary(ctx *PreprocessorBinaryContext) {
}

// ExitPreprocessorBinary is called when production preprocessorBinary is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorBinary(ctx *PreprocessorBinaryContext) {
}

// EnterPreprocessorConstant is called when production preprocessorConstant is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorConstant(ctx *PreprocessorConstantContext) {
}

// ExitPreprocessorConstant is called when production preprocessorConstant is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorConstant(ctx *PreprocessorConstantContext) {
}

// EnterPreprocessorConditionalSymbol is called when production preprocessorConditionalSymbol is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorConditionalSymbol(ctx *PreprocessorConditionalSymbolContext) {
}

// ExitPreprocessorConditionalSymbol is called when production preprocessorConditionalSymbol is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorConditionalSymbol(ctx *PreprocessorConditionalSymbolContext) {
}

// EnterPreprocessorDefined is called when production preprocessorDefined is entered.
func (s *BaseObjectiveCPreprocessorParserListener) EnterPreprocessorDefined(ctx *PreprocessorDefinedContext) {
}

// ExitPreprocessorDefined is called when production preprocessorDefined is exited.
func (s *BaseObjectiveCPreprocessorParserListener) ExitPreprocessorDefined(ctx *PreprocessorDefinedContext) {
}
