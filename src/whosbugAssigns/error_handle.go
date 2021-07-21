package whosbugAssigns

import (
	"fmt"
	"log"
)

// @title errorHandler
// @description 用于错误处理，可传入错误日志和打印的附加字段
// @auth KevinMatt
// @param err error "传入的错误类型"
func errorHandler(err error, message ...string) {

	if err != nil {
		fmt.Println(message[0], err.Error())
		log.Fatal(err)
	}
}
