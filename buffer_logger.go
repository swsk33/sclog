package sclog

import (
	"fmt"
	"time"
)

// 缓存日志消息
type bufferMessage struct {
	// 日志等级
	level int
	// 日志产生时间
	time time.Time
	// 消息内容格式化字符串
	formatMessage string
	// 格式化参数切片
	args []interface{}
}

// BufferLogger 基于通道缓冲区的线程安全日志输出器
type BufferLogger struct {
	// 继承普通Logger
	*Logger
	// 缓存消息的通道
	messageChannel chan *bufferMessage
}

// NewBufferLogger 构造函数
//
// size 缓冲区大小
// 当缓冲区内的消息数量未被消费（输出）且已到达设定大小时，再次调用日志输出就会导致阻塞
func NewBufferLogger(size int) *BufferLogger {
	// 创建带缓冲区的日志对象
	logger := &BufferLogger{
		Logger:         NewLogger(),
		messageChannel: make(chan *bufferMessage, size),
	}
	// 在一个新的线程中消费日志消息
	go func() {
		for message := range logger.messageChannel {
			logger.printBufferMessage(message)
		}
	}()
	return logger
}

// 打印缓存消息对象
func (logger *BufferLogger) printBufferMessage(message *bufferMessage) {
	// 小于当前级别则不进行输出
	if message.level < logger.Level {
		return
	}
	// 获取当前级别配置
	config := logger.LevelConfig[message.level]
	// 打印时间部分
	if config.Time.Enabled {
		_, _ = config.Time.Color.Printf("%s ", message.time.Format(config.Time.Pattern))
	}
	// 打印级别部分
	if config.Level.Enabled {
		_, _ = config.Level.Color.Printf(fmt.Sprintf("%%-%ds ", maxLevelNameLength), levelMap[message.level])
	}
	// 打印日志消息部分
	if config.Message.Enabled {
		// 输出前缀
		_, _ = config.Message.Color.Printf("%s", config.Message.Prefix)
		// 输出主体
		_, _ = config.Message.Color.Printf(message.formatMessage, message.args...)
	}
}

// Trace 打印一行TRACE级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *BufferLogger) Trace(formatMessage string, args ...interface{}) {
	logger.messageChannel <- &bufferMessage{
		level:         TRACE,
		time:          time.Now(),
		formatMessage: formatMessage,
		args:          args,
	}
}

// TraceLine 打印一行TRACE级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *BufferLogger) TraceLine(messageLine string) {
	logger.messageChannel <- &bufferMessage{
		level:         TRACE,
		time:          time.Now(),
		formatMessage: messageLine + "\n",
		args:          []interface{}{},
	}
}

// Debug 打印一行DEBUG级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *BufferLogger) Debug(formatMessage string, args ...interface{}) {
	logger.messageChannel <- &bufferMessage{
		level:         DEBUG,
		time:          time.Now(),
		formatMessage: formatMessage,
		args:          args,
	}
}

// DebugLine 打印一行DEBUG级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *BufferLogger) DebugLine(messageLine string) {
	logger.messageChannel <- &bufferMessage{
		level:         DEBUG,
		time:          time.Now(),
		formatMessage: messageLine + "\n",
		args:          []interface{}{},
	}
}

// Info 打印一行INFO级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *BufferLogger) Info(formatMessage string, args ...interface{}) {
	logger.messageChannel <- &bufferMessage{
		level:         INFO,
		time:          time.Now(),
		formatMessage: formatMessage,
		args:          args,
	}
}

// InfoLine 打印一行INFO级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *BufferLogger) InfoLine(messageLine string) {
	logger.messageChannel <- &bufferMessage{
		level:         INFO,
		time:          time.Now(),
		formatMessage: messageLine + "\n",
		args:          []interface{}{},
	}
}

// Warn 打印一行WARN级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *BufferLogger) Warn(formatMessage string, args ...interface{}) {
	logger.messageChannel <- &bufferMessage{
		level:         WARN,
		time:          time.Now(),
		formatMessage: formatMessage,
		args:          args,
	}
}

// WarnLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *BufferLogger) WarnLine(messageLine string) {
	logger.messageChannel <- &bufferMessage{
		level:         WARN,
		time:          time.Now(),
		formatMessage: messageLine + "\n",
		args:          []interface{}{},
	}
}

// Error 打印一行ERROR级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *BufferLogger) Error(formatMessage string, args ...interface{}) {
	logger.messageChannel <- &bufferMessage{
		level:         ERROR,
		time:          time.Now(),
		formatMessage: formatMessage,
		args:          args,
	}
}

// ErrorLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *BufferLogger) ErrorLine(messageLine string) {
	logger.messageChannel <- &bufferMessage{
		level:         ERROR,
		time:          time.Now(),
		formatMessage: messageLine + "\n",
		args:          []interface{}{},
	}
}

// Close 关闭缓冲区日志，释放资源
func (logger *BufferLogger) Close() {
	close(logger.messageChannel)
}