package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"

	"runtime/pprof"

	"git.woa.com/bkdevops/whosbug"
	"git.woa.com/bkdevops/whosbug/config"
	"git.woa.com/bkdevops/whosbug/env"
	"git.woa.com/bkdevops/whosbug/zaplog"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			zaplog.Logger.Error("panic in main", zaplog.Any("error", err))
		}
	}()

	// TODO: 暂定为4，待优化
	runtime.GOMAXPROCS(4)

	if env.DevFlag {
		zaplog.Logger.Info("is test env")
		cpufile, _ := os.Create("./cpufile.prof")
		heapfile, _ := os.Create("./heapfile.prof")

		go func() {
			pprof.StartCPUProfile(cpufile)
			pprof.WriteHeapProfile(heapfile)
		}()

		defer cpufile.Close()
		defer heapfile.Close()
		defer pprof.StopCPUProfile()
	}

	// 插件输入参数
	_config := initConfig()
	zaplog.Logger.Info("get whosbug config", zaplog.Any("config", _config))

	whosbug.Analysis(_config)

}

// initConfig 初始化配置
//
//	@return whosbugConfig *config.Config
//	@author: Kevineluo 2022-07-30 07:42:34
func initConfig() (whosbugConfig *config.Config) {
	var (
		err            error
		inputJSONBytes []byte
	)
	zaplog.Logger.Info("Init whosbug...")

	if env.DevFlag {
		zaplog.Logger.Info("is dev env")
		inputJSONBytes, err = ioutil.ReadFile("./input.json")
		if err != nil {
			zaplog.Logger.Fatal(err.Error())
		}

		err = json.Unmarshal(inputJSONBytes, &whosbugConfig)
		if err != nil {
			zaplog.Logger.Emergency(err.Error())
			os.Exit(-1)
		} else {
			zaplog.Logger.Info("Get input-config succeed!")
		}
	} else {
		zaplog.Logger.Info("no dev env")
		whosbugConfig = new(config.Config)

		whosbugConfig.WebServerHost = "http://119.29.46.189:8083"
		whosbugConfig.WebServerKey = "whosbug2022"
		whosbugConfig.WebServerUserName = "admin"
		whosbugConfig.CryptoKey = ""
		whosbugConfig.BranchName = os.Getenv("BRANCH_NAME")
		whosbugConfig.ProjectID = os.Getenv("GIT_HTTP_URL")
		whosbugConfig.ProjectURL = "/root/workspace/"
		whosbugConfig.ReleaseVersion = os.Getenv("GIT_COMMIT")
	}
	return
}
