package main

import (
	"net/http"
	_ "net/http/pprof"
	"whosbugPack"
)

func main() {
	go func() {
		whosbugPack.Analysis()
	}()
	defer close(whosbugPack.ObjectChan)
	panic(http.ListenAndServe("0.0.0.0:6060", nil))
}
