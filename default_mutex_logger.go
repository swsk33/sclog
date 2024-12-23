package sclog

import "sync"

// 全局默认互斥线程安全日志输出器对象
var defaultMutexLogger *MutexLogger

// 懒加载全局的互斥线程安全日志输出器
var mutexLoggerOnce sync.Once

// 懒加载获取默认的MutexLogger
func getDefaultMutexLogger() *MutexLogger {
	if defaultMutexLogger == nil {
		mutexLoggerOnce.Do(func() {
			defaultMutexLogger = NewMutexLogger()
			defaultMutexLogger.Level = TRACE
		})
	}
	return defaultMutexLogger
}

// MutexTrace 打印一行TRACE级别日志
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func MutexTrace(formatMessage string, args ...interface{}) {
	getDefaultMutexLogger().Trace(formatMessage, args...)
}

// MutexTraceLine 打印一行TRACE级别日志，并换行
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
func MutexTraceLine(messageLine string) {
	getDefaultMutexLogger().TraceLine(messageLine)
}

// MutexDebug 打印一行DEBUG级别日志
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func MutexDebug(formatMessage string, args ...interface{}) {
	getDefaultMutexLogger().Debug(formatMessage, args...)
}

// MutexDebugLine 打印一行DEBUG级别日志，并换行
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
func MutexDebugLine(messageLine string) {
	getDefaultMutexLogger().DebugLine(messageLine)
}

// MutexInfo 打印一行INFO级别日志
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func MutexInfo(formatMessage string, args ...interface{}) {
	getDefaultMutexLogger().Info(formatMessage, args...)
}

// MutexInfoLine 打印一行INFO级别日志，并换行
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
func MutexInfoLine(messageLine string) {
	getDefaultMutexLogger().InfoLine(messageLine)
}

// MutexWarn 打印一行WARN级别日志
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func MutexWarn(formatMessage string, args ...interface{}) {
	getDefaultMutexLogger().Warn(formatMessage, args...)
}

// MutexWarnLine 打印一行Warn级别日志，并换行
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
func MutexWarnLine(messageLine string) {
	getDefaultMutexLogger().WarnLine(messageLine)
}

// MutexError 打印一行ERROR级别日志
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func MutexError(formatMessage string, args ...interface{}) {
	getDefaultMutexLogger().Error(formatMessage, args...)
}

// MutexErrorLine 打印一行Warn级别日志，并换行
// 使用互斥锁保证线程安全性
//
// formatMessage 格式化的消息字符串
func MutexErrorLine(messageLine string) {
	getDefaultMutexLogger().ErrorLine(messageLine)
}