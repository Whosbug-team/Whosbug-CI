package whosbugPack

import (
	pythonparser "anrlr4_ast/Python"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type BasePython3ParserListener struct {
	Infos AnalysisInfoType
}

/* ExitFuncdef
/* @Description: 匹配函数定义
 * @receiver s
 * @param ctx
 * @author KevinMatt 2021-08-01 12:31:11
 * @function_mark
*/
func (s *BasePython3ParserListener) ExitFuncdef(ctx *pythonparser.FuncdefContext) {
	var methodInfo MethodInfoType
	MethodName := ctx.GetChild(1).(antlr.ParseTree).GetText()
	ReturnType := ctx.GetChild(0).(antlr.ParseTree).GetText()
	Params := getParams(ctx.GetChild(2).(antlr.ParseTree))
	//fmt.Println(Params[0].paramName, Params[0].paramType)
	methodInfo.ReturnType = ReturnType
	methodInfo.StartLine = ctx.GetStart().GetLine()
	methodInfo.EndLine = ctx.GetStop().GetLine()
	methodInfo.MethodName = MethodName
	methodInfo.MasterObject = masterObjectInfoType{}
	methodInfo.Params = append(methodInfo.Params, Params...)
	s.Infos.AstInfoList.Methods = append(s.Infos.AstInfoList.Methods, methodInfo)
}

/* EnterClassdef
/* @Description:
 * @receiver s
 * @param ctx
 * @author KevinMatt 2021-08-01 12:36:26
 * @function_mark
*/
func (s *BasePython3ParserListener) EnterClassdef(ctx *pythonparser.ClassdefContext) {
	// TODO 断点调试解析树节点结构
}
