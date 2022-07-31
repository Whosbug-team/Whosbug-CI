package util

import "fmt"

// ErrorMessage
//	@Description: 只打印错误信息，不打印堆栈
//	@param err
//	@return string
//	@author KevinMatt 2021-08-08 16:14:42
//	@function_mark PASS
func ErrorMessage(err error) string {
	return err.Error()
}

// ErrorStack
//	@Description: 打印含堆栈的错误信息
//	@param err 错误
//	@return string 字符串
//	@author KevinMatt 2021-08-08 16:13:58
//	@function_mark PASS
func ErrorStack(err error) string {
	errMsg := fmt.Sprintf("%+v", err)
	return CleanPath(errMsg)
}
