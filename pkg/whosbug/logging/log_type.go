package logging

import (
	"os"
	"strings"
	"time"
)

var LogFile *os.File

func init() {
	if _, err := os.Stat("/data/log"); !os.IsNotExist(err) {
		os.RemoveAll("/data/log/")
	}
	os.MkdirAll("/data/log", os.ModePerm)
	LogFile, _ = os.OpenFile("/data/log"+"/whosbug-"+strings.ReplaceAll(time.Now().Local().Format("2006-01-02 15:04:05"), " ", "--")+".log", os.O_RDWR|os.O_CREATE, os.ModePerm)
}
