package env

import (
	"os"
	"strings"
)

var (
	// ModuleName 完整模块名（如whosbug-CI）
	ModuleName = GetEnv("MODULE_NAME", "whosbug-ci")

	// ServiceName 去除后缀的模块名（如whosbug）
	ServiceName = strings.Split(ModuleName, "-")[0]

	// DevFlag 是否是开发模式
	DevFlag = GetEnv("IS_DEV", "false") == "true"
)

// GetEnv 获取环境变量，获取不到则使用默认值
//
//	@param key string
//	@param defaultVal string
//	@return val string
//	@author: Kevineluo 2022-04-19 03:34:22
func GetEnv(key, defaultVal string) (val string) {
	val = os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return
}
