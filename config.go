package sclog

import (
	"github.com/fatih/color"
	"math"
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