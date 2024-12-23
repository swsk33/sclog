package sclog

import "sync"

// 全局默认带缓冲区的线程安全日志输出器对象
var defaultBufferLogger *BufferLogger

// 懒加载全局的带缓冲区线程安全日志输出器
var bufferLoggerOnce sync.Once

// 懒加载获取默认的BufferLogger
func getDefaultBufferLogger() *BufferLogger {
	if defaultBufferLogger == nil {
		bufferLoggerOnce.Do(func() {
			defaultBufferLogger = NewBufferLogger(233)
			defaultBufferLogger.Level = TRACE
		})
	}
	return defaultBufferLogger
}

// BufferTrace 打印一行TRACE级别日志
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func BufferTrace(formatMessage string, args ...interface{}) {
	getDefaultBufferLogger().Trace(formatMessage, args...)
}

// BufferTraceLine 打印一行TRACE级别日志，并换行
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
func BufferTraceLine(messageLine string) {
	getDefaultBufferLogger().TraceLine(messageLine)
}

// BufferDebug 打印一行DEBUG级别日志
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func BufferDebug(formatMessage string, args ...interface{}) {
	getDefaultBufferLogger().Debug(formatMessage, args...)
}

// BufferDebugLine 打印一行DEBUG级别日志，并换行
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
func BufferDebugLine(messageLine string) {
	getDefaultBufferLogger().DebugLine(messageLine)
}

// BufferInfo 打印一行INFO级别日志
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func BufferInfo(formatMessage string, args ...interface{}) {
	getDefaultBufferLogger().Info(formatMessage, args...)
}

// BufferInfoLine 打印一行INFO级别日志，并换行
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
func BufferInfoLine(messageLine string) {
	getDefaultBufferLogger().InfoLine(messageLine)
}

// BufferWarn 打印一行WARN级别日志
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func BufferWarn(formatMessage string, args ...interface{}) {
	getDefaultBufferLogger().Warn(formatMessage, args...)
}

// BufferWarnLine 打印一行Warn级别日志，并换行
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
func BufferWarnLine(messageLine string) {
	getDefaultBufferLogger().WarnLine(messageLine)
}

// BufferError 打印一行ERROR级别日志
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func BufferError(formatMessage string, args ...interface{}) {
	getDefaultBufferLogger().Error(formatMessage, args...)
}

// BufferErrorLine 打印一行Warn级别日志，并换行
// 使用通道作为缓冲区保证线程安全性
//
// formatMessage 格式化的消息字符串
func BufferErrorLine(messageLine string) {
	getDefaultBufferLogger().ErrorLine(messageLine)
}