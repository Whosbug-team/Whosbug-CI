package whosbugAssigns

import (
	"fmt"
	"time"
)

var workPath string = "C:\\Users\\KevinMatt\\Desktop\\whosbug-Golang"

func MainProcess() {
	t := time.Now()

	fmt.Println("All cost: ", time.Since(t))
}
