package sclog

import (
	"sync"
)

// MutexLogger 基于互斥锁的线程安全日志输出器
type MutexLogger struct {
	// 继承普通Logger
	*Logger
	// 互斥锁
	lock sync.Mutex
}

// NewMutexLogger 构造函数
func NewMutexLogger() *MutexLogger {
	return &MutexLogger{
		Logger: NewLogger(),
		lock:   sync.Mutex{},
	}
}

// 根据对应级别的配置，打印一行日志到控制台，该打印不会换行
// 互斥日志的线程安全重写
//
// level 当前日志级别
// message 要输出的日志的格式化字符串
// args 格式化字符串的参数
func (logger *MutexLogger) printLog(level int, formatMessage string, args ...interface{}) {
	// 上锁
	logger.lock.Lock()
	defer logger.lock.Unlock()
	// 调用父类方法
	logger.Logger.printLog(level, formatMessage, args...)
}

// Trace 打印一行TRACE级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *MutexLogger) Trace(formatMessage string, args ...interface{}) {
	logger.printLog(TRACE, formatMessage, args...)
}

// TraceLine 打印一行TRACE级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *MutexLogger) TraceLine(messageLine string) {
	logger.printLog(TRACE, messageLine+"\n")
}

// Debug 打印一行DEBUG级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *MutexLogger) Debug(formatMessage string, args ...interface{}) {
	logger.printLog(DEBUG, formatMessage, args...)
}

// DebugLine 打印一行DEBUG级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *MutexLogger) DebugLine(messageLine string) {
	logger.printLog(DEBUG, messageLine+"\n")
}

// Info 打印一行INFO级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *MutexLogger) Info(formatMessage string, args ...interface{}) {
	logger.printLog(INFO, formatMessage, args...)
}

// InfoLine 打印一行INFO级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *MutexLogger) InfoLine(messageLine string) {
	logger.printLog(INFO, messageLine+"\n")
}

// Warn 打印一行WARN级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *MutexLogger) Warn(formatMessage string, args ...interface{}) {
	logger.printLog(WARN, formatMessage, args...)
}

// WarnLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *MutexLogger) WarnLine(messageLine string) {
	logger.printLog(WARN, messageLine+"\n")
}

// Error 打印一行ERROR级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *MutexLogger) Error(formatMessage string, args ...interface{}) {
	logger.printLog(ERROR, formatMessage, args...)
}

// ErrorLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *MutexLogger) ErrorLine(messageLine string) {
	logger.printLog(ERROR, messageLine+"\n")
}