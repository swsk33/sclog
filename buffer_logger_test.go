package sclog

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"testing"
	"time"
)

func TestDefaultBufferLogger(t *testing.T) {
	BufferTrace("这是格式化Trace日志！日志级别常量：%d\n", TRACE)
	BufferTraceLine("这是换行的Trace日志！")
	BufferDebug("这是格式化Debug日志！日志级别常量：%d\n", DEBUG)
	BufferDebugLine("这是换行的Debug日志！")
	BufferInfo("这是格式化Info日志！日志级别常量：%d\n", INFO)
	BufferInfoLine("这是换行的Info日志！")
	BufferWarn("这是格式化Warn日志！日志级别常量：%d\n", WARN)
	BufferWarnLine("这是换行的Warn日志！")
	BufferError("这是格式化Error日志！日志级别常量：%d\n", ERROR)
	BufferErrorLine("这是换行的Error日志！")
	fmt.Println("----------------------------------")
	time.Sleep(100 * time.Millisecond)
}

func TestCustomBufferLogger(t *testing.T) {
	// 创建一个自定义的日志输出器对象
	logger := NewBufferLogger(10)
	// 创建一个配置对象
	config := NewLineConfig()
	config.Time.Enabled = false
	config.Message.Color = color.New(color.FgCyan)
	// 设定全部级别为自定义配置
	logger.ConfigAll(config)
	// 打印日志
	logger.InfoLine("Info级别日志")
	logger.WarnLine("Warn级别日志")
	logger.ErrorLine("Error级别日志")
	fmt.Println("----------------------------------")
	time.Sleep(100 * time.Millisecond)
}

func TestBufferCustomLevelName(t *testing.T) {
	// 自定义每个级别的名字
	SetLevelName(INFO, "普通")
	SetLevelName(WARN, "警告")
	SetLevelName(ERROR, "错误")
	// 打印日志
	a := 1
	BufferInfo("a的值为：%d\n", a)
	BufferWarn("a的值为：%d\n", a)
	BufferError("a的值为：%d\n", a)
	fmt.Println("----------------------------------")
	time.Sleep(100 * time.Millisecond)
}

func TestBufferLoggerConcurrent(t *testing.T) {
	SetLevelName(INFO, "INFO")
	SetLevelName(WARN, "WARN")
	SetLevelName(ERROR, "ERROR")
	// 自定义配置及其前缀
	logger := NewBufferLogger(20)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		for i := 0; i <= 10; i += 2 {
			logger.Info("[线程1]当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 1; i <= 11; i += 2 {
			logger.Info("[线程2]当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
	fmt.Println("----------------------------------")
	time.Sleep(100 * time.Millisecond)
}