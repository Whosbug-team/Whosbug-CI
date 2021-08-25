package util

import (
	"os"
	"strings"

	go_logger "github.com/phachon/go-logger"
)

var (
	GLogger    *go_logger.Logger
	ModuleName = "whosbug"
	level      = os.Getenv("LOG_LEVEL")
)

func init() {
	GLogger = go_logger.NewLogger()

	// 日志等级默认为INFO
	logLevel := go_logger.LOGGER_LEVEL_INFO
	if strings.ToUpper(level) == "DEBUG" {
		logLevel = go_logger.LOGGER_LEVEL_DEBUG
	} else if strings.ToUpper(level) == "INFO" {
		logLevel = go_logger.LOGGER_LEVEL_INFO
	} else if strings.ToUpper(level) == "ERROR" {
		logLevel = go_logger.LOGGER_LEVEL_ERROR
	} else if strings.ToUpper(level) == "EMERGENCY" {
		logLevel = go_logger.LOGGER_LEVEL_EMERGENCY
	} else if strings.ToUpper(level) == "ALERT" {
		logLevel = go_logger.LOGGER_LEVEL_ALERT
	} else if strings.ToUpper(level) == "CRITICAL" {
		logLevel = go_logger.LOGGER_LEVEL_CRITICAL
	} else if strings.ToUpper(level) == "NOTICE" {
		logLevel = go_logger.LOGGER_LEVEL_NOTICE
	} else if strings.ToUpper(level) == "WARNING" {
		logLevel = go_logger.LOGGER_LEVEL_WARNING
	}
	consoleConfig := &go_logger.ConsoleConfig{
		Format: "%timestamp_format% [%level_string%] [%file%:%function%](line %line%): %body%",
	}
	GLogger.Detach("console")
	GLogger.Attach("console", logLevel, consoleConfig)

	// 日志文件配置
	// dirPath := "~/whosbug_logs/"
	// CreateDirIfNotExists(dirPath)
	// logFile := fmt.Sprintf("%s/%s.log", dirPath, ModuleName)
	// logFileErr := fmt.Sprintf("%s/%s_error.log", dirPath, ModuleName)
	// fileConfig := &go_logger.FileConfig{
	// 	Filename:   logFile,
	// 	MaxSize:    501 * 1024,
	// 	MaxLine:    0,
	// 	DateSlice:  "d",
	// 	JsonFormat: false,
	// 	Format:     "%timestamp_format% [%level_string%] %file%:%function% %line%: %body%",
	// 	LevelFileName: map[int]string{
	// 		GLogger.LoggerLevel("ERROR"):     logFileErr, // The error level log is written to the error.log file.
	// 		GLogger.LoggerLevel("EMERGENCY"): logFileErr, // The emergency level log is written to the error.log file.
	// 		GLogger.LoggerLevel("CRITICAL"):  logFileErr, // The critical level log is written to the error.log file.
	// 		GLogger.LoggerLevel("INFO"):      logFile,    // The info level log is written to the error.log file.
	// 		GLogger.LoggerLevel("ALERT"):     logFile,    // The alert level log is written to the error.log file.
	// 		GLogger.LoggerLevel("WARNING"):   logFile,    // The warning level log is written to the error.log file.
	// 		GLogger.LoggerLevel("NOTICE"):    logFile,    // The notice level log is written to the error.log file.
	// 		GLogger.LoggerLevel("DEBUG"):     logFile,    // The debug level log is written to the error.log file.
	// 	},
	// }
	// GLogger.Attach("file", logLevel, fileConfig)
}
