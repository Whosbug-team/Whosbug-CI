package logpack

import (
	"os"
	"whosbugPack/global_type"
)

var LogFile, _ = os.OpenFile(global_type.WorkPath+"log.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
