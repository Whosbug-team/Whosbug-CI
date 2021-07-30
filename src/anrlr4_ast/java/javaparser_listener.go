// Code generated from JavaParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package javaParser // JavaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JavaParserListener is a complete listener for a parse tree produced by JavaParser.
type JavaParserListener interface {
	antlr.ParseTreeListener

	ExitMethodDeclaration(c *MethodDeclarationContext)

	EnterMethodCall(c *MethodCallContext)

	EnterClassDeclaration(c *ClassDeclarationContext)

	EnterFieldDeclaration(c *FieldDeclarationContext)

	EnterPackageDeclaration(c *PackageDeclarationContext)
}
