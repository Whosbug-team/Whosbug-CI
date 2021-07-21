package whosbugAssigns

import (
	"bytes"
	"os/exec"
	"path"
)

var supportLans = []string{"*.java"}

// @title execCommandOutput
// @description 执行输入的字符串命令并返回标准输出
// @author KevinMatt
// @param command string 输入的命令
// @return output.String() string输出信息
func execCommandOutput(command string) string {

	cmd := exec.Command(command)
	output := bytes.Buffer{}
	cmd.Stdout = &output
	err := cmd.Run()
	errorHandler(err, "exec command error:")
	return output.String()
}

func lanFilter(fileName string) bool {
	for _, supportLan := range supportLans {
		if path.Ext(fileName) == supportLan {
			return true
		}
	}
	return false
}
