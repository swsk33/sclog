package sclog

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	Trace("这是格式化Trace日志！日志级别常量：%d\n", TRACE)
	TraceLine("这是换行的Trace日志！")
	Debug("这是格式化Debug日志！日志级别常量：%d\n", DEBUG)
	DebugLine("这是换行的Debug日志！")
	Info("这是格式化Info日志！日志级别常量：%d\n", INFO)
	InfoLine("这是换行的Info日志！")
	Warn("这是格式化Warn日志！日志级别常量：%d\n", WARN)
	WarnLine("这是换行的Warn日志！")
	Error("这是格式化Error日志！日志级别常量：%d\n", ERROR)
	ErrorLine("这是换行的Error日志！")
	fmt.Println("----------------------------------")
}

func TestCustomLogger(t *testing.T) {
	// 创建一个自定义的日志输出器对象
	logger := NewLogger()
	// 创建一个配置对象
	config := NewLineConfig(false, true)
	config.Message.Color = color.New(color.FgCyan)
	// 设定全部级别为自定义配置
	logger.ConfigAll(config)
	// 打印日志
	logger.InfoLine("Info级别日志")
	logger.WarnLine("Warn级别日志")
	logger.ErrorLine("Error级别日志")
	fmt.Println("----------------------------------")
}