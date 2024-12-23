package sclog

import "sync"

// 全局默认日志输出器对象
var defaultLogger *Logger

// 懒加载日志输出器
var defaultLoggerOnce sync.Once

// 懒加载获取默认的Logger
func getDefaultLogger() *Logger {
	if defaultLogger == nil {
		defaultLoggerOnce.Do(func() {
			defaultLogger = NewLogger()
			defaultLogger.Level = TRACE
		})
	}
	return defaultLogger
}

// Trace 打印一行TRACE级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Trace(formatMessage string, args ...interface{}) {
	getDefaultLogger().Trace(formatMessage, args...)
}

// TraceLine 打印一行TRACE级别日志，并换行
//
// formatMessage 格式化的消息字符串
func TraceLine(messageLine string) {
	getDefaultLogger().TraceLine(messageLine)
}

// Debug 打印一行DEBUG级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Debug(formatMessage string, args ...interface{}) {
	getDefaultLogger().Debug(formatMessage, args...)
}

// DebugLine 打印一行DEBUG级别日志，并换行
//
// formatMessage 格式化的消息字符串
func DebugLine(messageLine string) {
	getDefaultLogger().DebugLine(messageLine)
}

// Info 打印一行INFO级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Info(formatMessage string, args ...interface{}) {
	getDefaultLogger().Info(formatMessage, args...)
}

// InfoLine 打印一行INFO级别日志，并换行
//
// formatMessage 格式化的消息字符串
func InfoLine(messageLine string) {
	getDefaultLogger().InfoLine(messageLine)
}

// Warn 打印一行WARN级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Warn(formatMessage string, args ...interface{}) {
	getDefaultLogger().Warn(formatMessage, args...)
}

// WarnLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func WarnLine(messageLine string) {
	getDefaultLogger().WarnLine(messageLine)
}

// Error 打印一行ERROR级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Error(formatMessage string, args ...interface{}) {
	getDefaultLogger().Error(formatMessage, args...)
}

// ErrorLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func ErrorLine(messageLine string) {
	getDefaultLogger().ErrorLine(messageLine)
}