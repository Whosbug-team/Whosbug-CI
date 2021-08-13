package main

import (
	"fmt"
	"os"
	"whosbugPack"
	"whosbugPack/logpack"
)

func main() {
	logpack.ExecCommandOutput("git", "clone", "https://gitee.com/egzosn/pay-java-parent.git")
	st := logpack.ExecCommandOutput("ls")
	fmt.Println(st)
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	whosbugPack.Analysis()
}
