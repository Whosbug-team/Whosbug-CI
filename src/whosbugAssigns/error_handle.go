package whosbugAssigns

import (
	"fmt"
	"log"
)

/* errorHandler
/* @Description: 通用错误处理函数
 * @param err 传入的错误类型
 * @param message 错误处理的消息
 * @author KevinMatt 2021-07-26 20:48:38
 * @function_mark PASS
*/
func errorHandler(err error, message ...string) {
	if err != nil {
		fmt.Println(message, ":", err.Error())
		log.Fatal(err)
	}
}
