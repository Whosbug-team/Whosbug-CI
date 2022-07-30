package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"git.woa.com/bkdevops/golang-atom-sdk/api"
	"git.woa.com/bkdevops/golang-atom-sdk/log"
	"git.woa.com/bkdevops/whosbug"
	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/util"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("panic: ", err)
			api.FinishBuild(api.StatusError, "panic occurs")
		}
	}()

	runtime.GOMAXPROCS(4)

	// 插件输入参数
	_config := initConfig()
	util.GLogger.Infof(fmt.Sprintf("Whosbug config:\n%+v", _config))

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
	util.GLogger.Info("Init whosbug...")

	inputJSONBytes, err = ioutil.ReadFile("./input.json")
	if err != nil {
		util.GLogger.Emergency(err.Error())
		os.Exit(-1)
	}

	err = json.Unmarshal(inputJSONBytes, &whosbugConfig)
	if err != nil {
		util.GLogger.Emergency(err.Error())
		os.Exit(-1)
	} else {
		util.GLogger.Info("Get input-config succeed!")
	}
	return
}
