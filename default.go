package sclog

// 全局默认日志输出器对象
var defaultLogger = NewLogger()

func init() {
	// 初始化部分配置
	defaultLogger.Level = TRACE
}

// Trace 打印一行TRACE级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Trace(formatMessage string, args ...interface{}) {
	defaultLogger.Trace(formatMessage, args...)
}

// TraceLine 打印一行TRACE级别日志，并换行
//
// formatMessage 格式化的消息字符串
func TraceLine(messageLine string) {
	defaultLogger.TraceLine(messageLine)
}

// Debug 打印一行DEBUG级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Debug(formatMessage string, args ...interface{}) {
	defaultLogger.Debug(formatMessage, args...)
}

// DebugLine 打印一行DEBUG级别日志，并换行
//
// formatMessage 格式化的消息字符串
func DebugLine(messageLine string) {
	defaultLogger.DebugLine(messageLine)
}

// Info 打印一行INFO级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Info(formatMessage string, args ...interface{}) {
	defaultLogger.Info(formatMessage, args...)
}

// InfoLine 打印一行INFO级别日志，并换行
//
// formatMessage 格式化的消息字符串
func InfoLine(messageLine string) {
	defaultLogger.InfoLine(messageLine)
}

// Warn 打印一行WARN级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Warn(formatMessage string, args ...interface{}) {
	defaultLogger.Warn(formatMessage, args...)
}

// WarnLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func WarnLine(messageLine string) {
	defaultLogger.WarnLine(messageLine)
}

// Error 打印一行ERROR级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func Error(formatMessage string, args ...interface{}) {
	defaultLogger.Error(formatMessage, args...)
}

// ErrorLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func ErrorLine(messageLine string) {
	defaultLogger.ErrorLine(messageLine)
}