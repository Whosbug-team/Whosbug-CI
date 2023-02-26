package logging

import (
	"os"
	"strings"
	"time"
)

var LogFile *os.File

func init() {
	LogFile, _ = func() (*os.File, error) {
		if _, err := os.Stat("/data/Whosbuglog"); !os.IsNotExist(err) {
			os.RemoveAll("/data/Whosbuglog/")
		}
		os.MkdirAll("/data/Whosbuglog", os.ModePerm)
		return os.OpenFile("/data/Whosbuglog"+"/log-"+strings.ReplaceAll(time.Now().Local().Format("2006-01-02 15:04:05"), " ", "--")+".txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	}()
}
