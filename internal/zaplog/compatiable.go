package zaplog

import (
	"go.uber.org/zap/zapcore"
)

// Emergency 兼容性方法，不建议在新的模块中使用
//  @receiver l *ZapLogger
//  @param msg string
//  @author kevinmatthe
func (l *ZapLogger) Emergency(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

// Emergencyf 兼容性方法，不建议在新的模块中使用
//  @receiver l *ZapLogger
//  @param format string
//  @param args ...interface{}
func (l *ZapLogger) Emergencyf(format string, args ...interface{}) {
	l.logger.Sugar().Warnf(format, args...)
}

// Warning 兼容性方法，不建议在新的模块中使用
//  @receiver l *ZapLogger
//  @param msg
func (l *ZapLogger) Warning(msg string) {
	l.logger.Warn(msg)
}
