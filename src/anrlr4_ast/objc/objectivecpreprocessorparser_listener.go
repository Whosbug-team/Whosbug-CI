// Generated from ObjectiveCPreprocessorParser.g4 by ANTLR 4.7.

package objcParser // ObjectiveCPreprocessorParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ObjectiveCPreprocessorParserListener is a complete listener for a parse tree produced by ObjectiveCPreprocessorParser.
type ObjectiveCPreprocessorParserListener interface {
	antlr.ParseTreeListener

	// EnterPreprocessorImport is called when entering the preprocessorImport production.
	EnterPreprocessorImport(c *PreprocessorImportContext)

	// EnterPreprocessorConditional is called when entering the preprocessorConditional production.
	EnterPreprocessorConditional(c *PreprocessorConditionalContext)

	// EnterPreprocessorDef is called when entering the preprocessorDef production.
	EnterPreprocessorDef(c *PreprocessorDefContext)

	// EnterPreprocessorPragma is called when entering the preprocessorPragma production.
	EnterPreprocessorPragma(c *PreprocessorPragmaContext)

	// EnterPreprocessorError is called when entering the preprocessorError production.
	EnterPreprocessorError(c *PreprocessorErrorContext)

	// EnterPreprocessorWarning is called when entering the preprocessorWarning production.
	EnterPreprocessorWarning(c *PreprocessorWarningContext)

	// EnterPreprocessorDefine is called when entering the preprocessorDefine production.
	EnterPreprocessorDefine(c *PreprocessorDefineContext)

	// EnterDirectiveText is called when entering the directiveText production.
	EnterDirectiveText(c *DirectiveTextContext)

	// EnterPreprocessorParenthesis is called when entering the preprocessorParenthesis production.
	EnterPreprocessorParenthesis(c *PreprocessorParenthesisContext)

	// EnterPreprocessorNot is called when entering the preprocessorNot production.
	EnterPreprocessorNot(c *PreprocessorNotContext)

	// EnterPreprocessorBinary is called when entering the preprocessorBinary production.
	EnterPreprocessorBinary(c *PreprocessorBinaryContext)

	// EnterPreprocessorConstant is called when entering the preprocessorConstant production.
	EnterPreprocessorConstant(c *PreprocessorConstantContext)

	// EnterPreprocessorConditionalSymbol is called when entering the preprocessorConditionalSymbol production.
	EnterPreprocessorConditionalSymbol(c *PreprocessorConditionalSymbolContext)

	// EnterPreprocessorDefined is called when entering the preprocessorDefined production.
	EnterPreprocessorDefined(c *PreprocessorDefinedContext)

	// ExitPreprocessorImport is called when exiting the preprocessorImport production.
	ExitPreprocessorImport(c *PreprocessorImportContext)

	// ExitPreprocessorConditional is called when exiting the preprocessorConditional production.
	ExitPreprocessorConditional(c *PreprocessorConditionalContext)

	// ExitPreprocessorDef is called when exiting the preprocessorDef production.
	ExitPreprocessorDef(c *PreprocessorDefContext)

	// ExitPreprocessorPragma is called when exiting the preprocessorPragma production.
	ExitPreprocessorPragma(c *PreprocessorPragmaContext)

	// ExitPreprocessorError is called when exiting the preprocessorError production.
	ExitPreprocessorError(c *PreprocessorErrorContext)

	// ExitPreprocessorWarning is called when exiting the preprocessorWarning production.
	ExitPreprocessorWarning(c *PreprocessorWarningContext)

	// ExitPreprocessorDefine is called when exiting the preprocessorDefine production.
	ExitPreprocessorDefine(c *PreprocessorDefineContext)

	// ExitDirectiveText is called when exiting the directiveText production.
	ExitDirectiveText(c *DirectiveTextContext)

	// ExitPreprocessorParenthesis is called when exiting the preprocessorParenthesis production.
	ExitPreprocessorParenthesis(c *PreprocessorParenthesisContext)

	// ExitPreprocessorNot is called when exiting the preprocessorNot production.
	ExitPreprocessorNot(c *PreprocessorNotContext)

	// ExitPreprocessorBinary is called when exiting the preprocessorBinary production.
	ExitPreprocessorBinary(c *PreprocessorBinaryContext)

	// ExitPreprocessorConstant is called when exiting the preprocessorConstant production.
	ExitPreprocessorConstant(c *PreprocessorConstantContext)

	// ExitPreprocessorConditionalSymbol is called when exiting the preprocessorConditionalSymbol production.
	ExitPreprocessorConditionalSymbol(c *PreprocessorConditionalSymbolContext)

	// ExitPreprocessorDefined is called when exiting the preprocessorDefined production.
	ExitPreprocessorDefined(c *PreprocessorDefinedContext)
}
