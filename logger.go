package sclog

import (
	"fmt"
	"github.com/fatih/color"
	"math"
	"time"
)

// 定义级别常量
const (
	TRACE = 0
	DEBUG = 1
	INFO  = 2
	WARN  = 3
	ERROR = 4
	OFF   = math.MaxInt
)

// 每个级别常量对应的名称列表
var levelMap = map[int]string{
	TRACE: "TRACE",
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

// 最长的日志名称长度（用于对齐）
var maxLevelNameLength = 5

// SetLevelName 设定对应级别的显示名称
func SetLevelName(level int, name string) {
	levelMap[level] = name
	// 重新计算最长的名称
	maxLevelNameLength = 0
	for _, name := range levelMap {
		if len(name) > maxLevelNameLength {
			maxLevelNameLength = len(name)
		}
	}
}

// LineConfig 关于日志的每行输出的配置
type LineConfig struct {
	// 关于时间部分输出的配置
	Time struct {
		// 是否显示时间
		Enabled bool
		// 时间显示格式，为Go语言时间格式形式
		//
		// 例如：
		// "2006-01-02 15:04:05"
		// "2006-01-02 15:04:05.000"
		Pattern string
		// 输出颜色配置
		Color *color.Color
	}
	// 关于级别部分输出的配置
	Level struct {
		// 是否显示日志级别
		Enabled bool
		// 输出颜色配置
		Color *color.Color
	}
	// 关于具体日志消息部分的配置
	Message struct {
		// 是否显示日志消息
		Enabled bool
		// 消息部分前缀
		Prefix string
		// 输出颜色配置
		Color *color.Color
	}
}

// NewLineConfig 日志行内容配置构造函数
func NewLineConfig() *LineConfig {
	config := new(LineConfig)
	// 设定默认值
	// 时间
	config.Time.Enabled = true
	config.Time.Pattern = "2006-01-02 15:04:05.000"
	config.Time.Color = color.New(color.FgHiWhite)
	// 级别
	config.Level.Enabled = true
	config.Level.Color = color.New(color.FgHiWhite)
	// 消息
	config.Message.Enabled = true
	config.Message.Prefix = ""
	config.Message.Color = color.New(color.FgHiWhite)
	return config
}

// Logger 日志输出器对象
type Logger struct {
	// 最低日志级别
	Level int
	// 每个级别分别对应的配置的列表
	LevelConfig map[int]*LineConfig
}

// NewLogger 日志输出器构造函数
func NewLogger() *Logger {
	logger := &Logger{Level: INFO}
	logger.LevelConfig = make(map[int]*LineConfig)
	// 每一级别默认配置
	logger.LevelConfig[TRACE] = NewLineConfig()
	logger.LevelConfig[TRACE].Level.Color = color.New(color.FgWhite)
	logger.LevelConfig[DEBUG] = NewLineConfig()
	logger.LevelConfig[DEBUG].Level.Color = color.New(color.FgBlue)
	logger.LevelConfig[INFO] = NewLineConfig()
	logger.LevelConfig[INFO].Level.Color = color.New(color.FgGreen)
	logger.LevelConfig[WARN] = NewLineConfig()
	logger.LevelConfig[WARN].Level.Color = color.New(color.FgYellow)
	logger.LevelConfig[ERROR] = NewLineConfig()
	logger.LevelConfig[ERROR].Level.Color = color.New(color.FgRed)
	return logger
}

// ConfigAll 一键设置全部级别的输出配置
//
// config 行输出配置
func (logger *Logger) ConfigAll(config *LineConfig) {
	for level := range logger.LevelConfig {
		logger.ConfigLevel(level, config)
	}
}

// ConfigLevel 对某个级别的行输出单独进行配置
//
// level 要配置的级别
// config 传入对应配置
func (logger *Logger) ConfigLevel(level int, config *LineConfig) {
	// 复制一个配置对象
	copyConfig := *config
	logger.LevelConfig[level] = &copyConfig
}

// 根据对应级别的配置，打印一行日志到控制台，该打印不会换行
//
// level 当前日志级别
// message 要输出的日志的格式化字符串
// args 格式化字符串的参数
func (logger *Logger) printLog(level int, formatMessage string, args ...interface{}) {
	// 小于当前级别则不进行输出
	if level < logger.Level {
		return
	}
	// 获取当前级别配置
	config := logger.LevelConfig[level]
	// 打印时间部分
	if config.Time.Enabled {
		_, _ = config.Time.Color.Printf("%s ", time.Now().Format(config.Time.Pattern))
	}
	// 打印级别部分
	if config.Level.Enabled {
		_, _ = config.Level.Color.Printf(fmt.Sprintf("%%-%ds ", maxLevelNameLength), levelMap[level])
	}
	// 打印日志消息部分
	if config.Message.Enabled {
		// 输出前缀
		_, _ = config.Message.Color.Printf("%s", config.Message.Prefix)
		// 输出主体
		_, _ = config.Message.Color.Printf(formatMessage, args...)
	}
}

// Trace 打印一行TRACE级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *Logger) Trace(formatMessage string, args ...interface{}) {
	logger.printLog(TRACE, formatMessage, args...)
}

// TraceLine 打印一行TRACE级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *Logger) TraceLine(messageLine string) {
	logger.printLog(TRACE, messageLine+"\n")
}

// Debug 打印一行DEBUG级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *Logger) Debug(formatMessage string, args ...interface{}) {
	logger.printLog(DEBUG, formatMessage, args...)
}

// DebugLine 打印一行DEBUG级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *Logger) DebugLine(messageLine string) {
	logger.printLog(DEBUG, messageLine+"\n")
}

// Info 打印一行INFO级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *Logger) Info(formatMessage string, args ...interface{}) {
	logger.printLog(INFO, formatMessage, args...)
}

// InfoLine 打印一行INFO级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *Logger) InfoLine(messageLine string) {
	logger.printLog(INFO, messageLine+"\n")
}

// Warn 打印一行WARN级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *Logger) Warn(formatMessage string, args ...interface{}) {
	logger.printLog(WARN, formatMessage, args...)
}

// WarnLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *Logger) WarnLine(messageLine string) {
	logger.printLog(WARN, messageLine+"\n")
}

// Error 打印一行ERROR级别日志
//
// formatMessage 格式化的消息字符串
// args 格式化消息参数
func (logger *Logger) Error(formatMessage string, args ...interface{}) {
	logger.printLog(ERROR, formatMessage, args...)
}

// ErrorLine 打印一行Warn级别日志，并换行
//
// formatMessage 格式化的消息字符串
func (logger *Logger) ErrorLine(messageLine string) {
	logger.printLog(ERROR, messageLine+"\n")
}