package antlr

// AstParser antlr ast 解析器
//
//	@author kevineluo
//	@update 2023-02-27 09:15:02
type AstParser interface {
	AstParse(input string) AstInfo
}
