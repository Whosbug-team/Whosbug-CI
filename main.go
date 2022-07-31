package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"

	"git.woa.com/bkdevops/whosbug"
	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/zaplog"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			zaplog.Logger.Error("panic in main", zaplog.Any("error", err))
		}
	}()

	runtime.GOMAXPROCS(4)

	// 插件输入参数
	_config := initConfig()
	zaplog.Logger.Info("get whosbug config", zaplog.Any("config", _config))

	whosbug.Analysis(_config)
}

// initConfig 初始化配置
//  @return whosbugConfig *config.Config
//  @author: Kevineluo 2022-07-30 07:42:34
func initConfig() (whosbugConfig *config.Config) {
	var (
		err            error
		inputJSONBytes []byte
	)
	zaplog.Logger.Info("Init whosbug...")

	inputJSONBytes, err = ioutil.ReadFile("./input.json")
	if err != nil {
		zaplog.Logger.Emergency(err.Error())
		os.Exit(-1)
	}

	err = json.Unmarshal(inputJSONBytes, &whosbugConfig)
	if err != nil {
		zaplog.Logger.Emergency(err.Error())
		os.Exit(-1)
	} else {
		zaplog.Logger.Info("Get input-config succeed!")
	}
	return
}
