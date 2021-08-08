package main

import (
	_ "net/http/pprof"
	"whosbugPack"
)

func main() {
	whosbugPack.Analysis()
}
