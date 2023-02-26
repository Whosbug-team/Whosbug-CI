package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"
	"runtime/pprof"

	"git.woa.com/bkdevops/whosbug-ci/internal/env"
	"git.woa.com/bkdevops/whosbug-ci/internal/zaplog"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug"
	"git.woa.com/bkdevops/whosbug-ci/pkg/whosbug/config"
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
		cpuFile, _ := os.Create("./cpu.prof")
		heapFile, _ := os.Create("./heap.prof")

		defer cpuFile.Close()
		defer heapFile.Close()
		defer pprof.StopCPUProfile()

		go func() {
			pprof.StartCPUProfile(cpuFile)
			pprof.WriteHeapProfile(heapFile)
		}()
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
	whosbugConfig = new(config.Config)
	zaplog.Logger.Info("[initConfig] init whosbug config")

	if env.DevFlag {
		zaplog.Logger.Info("[initConfig] we are in dev mode")
		inputJSONBytes, err = ioutil.ReadFile("./input.json")
		if err != nil {
			zaplog.Logger.Fatal(err.Error())
		}

		err = json.Unmarshal(inputJSONBytes, &whosbugConfig)
		if err != nil {
			zaplog.Logger.Fatal(err.Error())
		} else {
			zaplog.Logger.Info("[initConfig] get config from input.json succeed!")
		}
	}
	return
}
