package whosbugAssigns

import (
	"fmt"
	"log"
)

// errorHandler
/* @Description:
 * @param err 传入的错误类型
 * @param message 错误处理的消息
 * @author KevinMatt 2021-07-25 14:35:41
 * @function_mark
 */
func errorHandler(err error, message ...string) {
	if err != nil {
		fmt.Println(message[0], err.Error())
		log.Fatal(err)
	}
}
