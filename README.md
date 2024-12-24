# 简单彩色日志-Go

## 1，介绍
sclog（Simple Color Log）是一个简单的Go日志库，可用于在控制台输出彩色字符的日志。

![image-20240929220540622](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929220540622.png)

该日志库的功能非常简单，仅用于在控制台上输出不同级别的日志，并支持自定义一些输出日志的配置，适用于一些简单的需要在控制台输出日志的场景或者小型程序。

如果你对性能有一定的要求，或者需要更加强大的日志功能，请考虑使用其它更加专业的结构化日志库，例如[Zap](https://github.com/uber-go/zap)、[Logrus](https://github.com/sirupsen/logrus)等等。

可通过下列命令集成该日志库：

```bash
go get gitee.com/swsk33/sclog
```

## 2，输出日志的基本方法

在导入`gitee.com/swsk33/sclog`包后，直接调用对应的日志级别方法即可：

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 格式化输出日志
	sclog.Trace("这是%s级别日志\n", "Trace")
	sclog.Debug("这是%s级别日志\n", "Debug")
	sclog.Info("这是%s级别日志\n", "Info")
	sclog.Warn("这是%s级别日志\n", "Warn")
	sclog.Error("这是%s级别日志\n", "Error")
}
```

结果：

![image-20240929225052530](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929225052530.png)

每一个级别的方法例如`Info`等等的使用方式和`fmt.Printf`一样，第一个参数都是格式化字符串，后面可以传入不定长个参数表示格式化字符串中每个占位符的值。

除了格式化输出日志之外，还可以输出单行日志，例如：

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 输出一行日志
	sclog.TraceLine("这是一行Trace日志")
	sclog.DebugLine("这是一行Debug日志")
	sclog.InfoLine("这是一行Info日志")
	sclog.WarnLine("这是一行Warn日志")
	sclog.ErrorLine("这是一行Error日志")
}
```

结果：

![image-20240929225429444](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929225429444.png)

以`Info`级别日志输出方法为例，可见：

- `Info`方法为格式化输出，并且不会自动换行，需要通过`\n`实现换行
- `InfoLine`方法为单行字符串输出，会自动换行

## 3，日志级别

目前有下列日志级别：

- `TRACE` 跟踪级别
- `DEBUG` 调试级别
- `INFO` 普通信息级别
- `WARN` 警告级别
- `ERROR` 错误级别

不同级别的权重大小为：

```
TRACE < DEBUG < INFO < WARN < ERROR
```

在`sclog`包下，每个级别都定义了一个常量，在后续部分涉及到设定日志最小级别等其它自定义操作时，就可以直接调用这些日志级别常量：

```go
// 表示INFO级别的常量
sclog.INFO
// 表示ERROR级别的常量
sclog.ERROR
```

## 4，自定义日志输出器

上述通过直接调用`sclog`包中的函数（例如`Info`、`Debug`等等）其实是调用的默认内置的日志输出器对象实现日志打印，默认的日志输出器是不能进行自定义和配置的，因此如果需要进行一些定制，可以自己创建新的日志输出器对象。

通过下列代码可以定义一个新的日志输出器`Logger`对象，并设定其最低级别为`DEBUG`级别：

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 创建一个新的日志输出器
	myLogger := sclog.NewLogger()
	// 设定最低级别为DEBUG
	myLogger.Level = sclog.DEBUG
	// 调用该日志输出器的对应方法输出日志
	myLogger.Trace("这是%s级别日志\n", "Trace")
	myLogger.Debug("这是%s级别日志\n", "Debug")
	myLogger.Info("这是%s级别日志\n", "Info")
	myLogger.Warn("这是%s级别日志\n", "Warn")
	myLogger.Error("这是%s级别日志\n", "Error")
}
```

结果：

![image-20240929231429759](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929231429759.png)

通过`NewLogger`方法能够返回一个日志输出器的指针对象，且包含默认的配置，调用该日志输出器对象的对应级别方法（例如`Info`、`InfoLine`等等）就可以输出日志，使用方式与`sclog`包下的函数一样，其默认的最低级别为`INFO`。可见上述由于设定了最低级别为`DEBUG`，因此低于该级别的`TRACE`日志将不会被打印。

如果想直接关闭全部级别的日志输出，则设定其`Level`为`sclog.OFF`即可。

## 5，输出内容的自定义配置

每一行日志的输出内容，其结构都是固定的，具体如下：

```
2024-09-29 23:13:37.347 INFO  这是Info级别日志
\_____________________/ \__/  \_____________/
           |              |          |
          时间           级别    日志消息内容
```

在`sclog`中还定义了**针对输出内容的自定义配置模型**，该模型结构体表示如下：

```go
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
```

这个结构体表示了**输出的每一行日志的结构、颜色等配置信息**，每一个级别都会对应一个这样的配置对象，可以通过`sclog`包的函数`NewLineConfig`创建一个`LineConfig`对象，刚创建出来的`LineConfig`对象都是包含默认值的。

在`LineConfig`对象中：

- `Time`属性为结构体，用于自定义日志**时间**部分的输出行为，其中：
	- `Enabled` 是否输出时间部分内容，默认为`true`
	- `Pattern` 格式化时间字符串，使用的是`time.Format`的`Layout`字符串，默认为`2006-01-02 15:04:05.000`
	- `Color` 时间部分的颜色，默认为白色
- `Level`属性为结构体，用于自定义日志**级别**部分的输出行为，其中：
	- `Enabled` 是否输出级别部分内容，默认为`true`
	- `Color` 时间部分的颜色，默认为白色
- `Message`属性为结构体，用于自定义日志**消息**部分的输出行为，其中：
	- `Enabled` 是否输出消息部分内容，默认为`true`
	- `Prefix` 消息部分的前缀字符串，默认为空字符串
	- `Color` 时间部分的颜色，默认为白色

通过创建`LineConfig`对象，修改需要的配置属性，并配置给日志输出器`Logger`的对应级别即可实现自定义输出内容，下面将通过几个示例进一步地实现日志输出的自定义。

### (1) 配置`LineConfig`的方式

在定义了`LineConfig`对象后，我们可以将其直接设定到`Logger`的全部级别上：

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 创建一个配置对象
	myConfig := sclog.NewLineConfig()
	// 修改时间部分配置
	// 关闭时间显示
	myConfig.Time.Enabled = false
	// 创建一个新的日志输出器
	myLogger := sclog.NewLogger()
	// 将自定义配置对象应用到全部的级别
	myLogger.ConfigAll(myConfig)
	// 调用该日志输出器的对应方法输出日志
	myLogger.Info("这是%s级别日志\n", "Info")
	myLogger.Warn("这是%s级别日志\n", "Warn")
	myLogger.Error("这是%s级别日志\n", "Error")
}
```

结果：

![image-20240929234254617](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929234254617.png)

通过`Logger`对象的`ConfigAll`方法可以直接将一个自定义的`LineConfig`配置对象应用到全部级别，如果说只想设定到其中一个级别，可以用`ConfigLevel`方法：

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 创建一个配置对象
	myConfig := sclog.NewLineConfig()
	// 修改时间部分配置
	// 关闭时间显示
	myConfig.Time.Enabled = false
	// 创建一个新的日志输出器
	myLogger := sclog.NewLogger()
	// 将自定义配置对象只应用到INFO级别
	myLogger.ConfigLevel(sclog.INFO, myConfig)
	// 调用该日志输出器的对应方法输出日志
	myLogger.Info("这是%s级别日志\n", "Info")
	myLogger.Warn("这是%s级别日志\n", "Warn")
	myLogger.Error("这是%s级别日志\n", "Error")
}
```

结果：

![image-20240929234553860](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929234553860.png)

此时，只有`INFO`级别应用了自定义的配置。

如果不想自己创建配置，只是想基于`Logger`默认的配置进行修改，则可以通过直接修改`Logger`中`LevelConfig`列表中对应级别的配置对象属性即可。

`Logger`的`LevelConfig`属性是一个`map`类型，键为不同的级别（例如`sclog.INFO`），值为该级别对应的配置对象，每个`Logger`中的`LevelConfig`中都包含了全部级别的配置对象。

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 创建一个新的日志输出器
	myLogger := sclog.NewLogger()
	// 直接修改Logger中INFO级别的配置
	// 修改INFO级别的时间Layout
	myLogger.LevelConfig[sclog.INFO].Time.Pattern = "2006-01-02 15:04:05"
	// 调用该日志输出器的对应方法输出日志
	myLogger.Info("这是%s级别日志\n", "Info")
	myLogger.Warn("这是%s级别日志\n", "Warn")
	myLogger.Error("这是%s级别日志\n", "Error")
}
```

结果：

![image-20240929235118311](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929235118311.png)

### (2) 消息内容前缀

通过设定`LineConfig`对象的`Message`结构体的`Prefix`属性，可以实现在日志内容部分加上自定义前缀：

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 创建一个新的日志输出器
	myLogger := sclog.NewLogger()
	// 修改INFO级别的消息前缀
	myLogger.LevelConfig[sclog.INFO].Message.Prefix = "[主线程] "
	// 调用该日志输出器的对应方法输出日志
	myLogger.Info("这是%s级别日志\n", "Info")
}
```

结果：

![image-20240929235742651](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240929235742651.png)

### (3) 自定义字符颜色

在默认的`Logger`中除了级别部分信息为彩色之外，其余部分消息均为白色字符，可通过设定`LineConfig`对象中对应结构体的`Color`属性实现设定颜色。

`LineConfig`对象中`Color`属性类型为`github.com/fatih/color`包下的`Color`指针类型，通过这个包下的`New`函数即可创建颜色对象，例如：

```go
package main

import (
	"gitee.com/swsk33/sclog"
	"github.com/fatih/color"
)

func main() {
	// 创建颜色对象
	// 前景色为白色，背景色为绿色
	infoColor := color.New(color.FgWhite, color.BgGreen)
	// 创建日志输出器对象
	logger := sclog.NewLogger()
	// 修改日志输出器对象中，INFO级别配置对象的级别部分颜色为我们自定义的颜色对象
	logger.LevelConfig[sclog.INFO].Level.Color = infoColor
	// 输出日志
	logger.InfoLine("这是Info级别日志")
}
```

结果：

![image-20240930183109048](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240930183109048.png)

颜色对象`color.Color`通过`New`构造函数创建，函数中可以**传入不定长个颜色常量作为参数以组合颜色**，在`color`包下定义了所有可用颜色的常量，其中`Fg`开头的表示前景色，即文字颜色，而`Bg`开头的表示背景色，默认情况下前景色为白色，背景色为透明。

此外，还可以定义加粗或者带下划线的颜色对象：

```go
package main

import (
	"gitee.com/swsk33/sclog"
	"github.com/fatih/color"
)

func main() {
	// 黄色加粗
	warnColor := color.New(color.FgHiYellow, color.Bold)
	// 红色带下划线
	errorColor := color.New(color.FgHiRed, color.Underline)
	// 配置自定义日志对象
	logger := sclog.NewLogger()
	// 配置对应级别消息部分颜色
	logger.LevelConfig[sclog.WARN].Message.Color = warnColor
	logger.LevelConfig[sclog.ERROR].Message.Color = errorColor
	// 打印日志
	logger.WarnLine("这是Warn日志")
	logger.ErrorLine("这是Error日志")
}
```

结果：

![image-20240930185525356](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240930185525356.png)

更多关于`color.Color`对象的创建，请参考`color`包的文档：[传送门](https://github.com/fatih/color)

## 6，修改级别显示名称

默认情况下，每个级别部分使用大写英文名称表示，例如`INFO`、`ERROR`等等，可以通过`SetLevelName`函数自定义每个级别的显示名：

```go
package main

import (
	"gitee.com/swsk33/sclog"
)

func main() {
	// 自定义每个级别的名字
	sclog.SetLevelName(sclog.INFO, "普通")
	sclog.SetLevelName(sclog.WARN, "警告")
	sclog.SetLevelName(sclog.ERROR, "错误")
	a := 1
	sclog.Info("a的值为：%d\n", a)
	sclog.Warn("a的值为：%d\n", a)
	sclog.Error("a的值为：%d\n", a)
}
```

结果：

![image-20240930185944883](https://swsk33-note.oss-cn-shanghai.aliyuncs.com/image-20240930185944883.png)

通过`SetLevelName`函数自定义的级别名称将应用于全局，无论是默认的`Logger`还是自定义的。

## 7，线程安全的日志

默认的日志输出器`Logger`不是线程安全的，这可能导致在多线程同时调用一个日志输出时发生输出消息错乱的情况。

因此，该类库还提供了下列线程安全的日志输出器实现：

- `MutexLogger` 基于互斥锁的线程安全的日志实现
- `BufferLogger` 基于通道缓冲区的线程安全的日志实现

上述线程安全的日志输出器都继承于`Logger`对象，因此可以使用上述方式对其进行配置例如级别、输出内容和颜色等。

### (1) `MutexLogger`的使用

可以直接调用内部默认的`MutexLogger`函数实现：

```go
package main

import "gitee.com/swsk33/sclog"

func main() {
	// 格式化输出
	sclog.MutexTrace("基于互斥锁的线程安全日志，级别：%d\n", sclog.TRACE)
	sclog.MutexDebug("基于互斥锁的线程安全日志，级别：%d\n", sclog.DEBUG)
	sclog.MutexInfo("基于互斥锁的线程安全日志，级别：%d\n", sclog.INFO)
	sclog.MutexWarn("基于互斥锁的线程安全日志，级别：%d\n", sclog.WARN)
	sclog.MutexError("基于互斥锁的线程安全日志，级别：%d\n", sclog.ERROR)
	// 单行输出
	sclog.MutexTraceLine("基于互斥锁的线程安全日志，单行输出")
	sclog.MutexDebugLine("基于互斥锁的线程安全日志，单行输出")
	sclog.MutexInfoLine("基于互斥锁的线程安全日志，单行输出")
	sclog.MutexWarnLine("基于互斥锁的线程安全日志，单行输出")
	sclog.MutexErrorLine("基于互斥锁的线程安全日志，单行输出")
}
```

只需在对应级别的方法名（例如`Info`）前面加上`Mutex`（例如`MutexInfo`）即可调用基于互斥锁的线程安全日志实现。

此外，也可以使用构造函数`NewMutexLogger`创建自定义的互斥锁线程安全日志输出器：

```go
package main

import (
	"gitee.com/swsk33/sclog"
	"sync"
)

func main() {
	// 创建基于互斥锁的线程安全日志输出器
	mutexLogger := sclog.NewMutexLogger()
	mutexLogger.Level = sclog.INFO
	// 调用对应级别方法即可
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		for i := 0; i <= 10; i += 2 {
			mutexLogger.Info("[线程1] 当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 1; i <= 11; i += 2 {
			mutexLogger.Info("[线程2] 当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
```

此外，还可以使用`NewMutexLoggerShareLock`传入一个现有的`sync.Mutex`对象，指定一个现有的互斥锁，实现多个日志输出器绑定到一个互斥锁上，这样即使创建多个`MutexLogger`对象，在多个线程中同时运行，也能够保证几个日志输出器的线程安全：

```go
package main

import (
	"gitee.com/swsk33/sclog"
	"sync"
)

func main() {
	// 互斥锁
	lock := &sync.Mutex{}
	// 创建两个日志输出器，使用同一个互斥锁
	l1, l2 := sclog.NewMutexLoggerShareLock(lock), sclog.NewMutexLoggerShareLock(lock)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		for i := 0; i <= 10; i += 2 {
			l1.Info("[线程1] 当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 1; i <= 11; i += 2 {
			l2.Info("[线程2] 当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
```

基于互斥锁的线程安全日志输出器虽然能够保证线程安全，但是在并发调用数量过多时可能发生锁竞争导致性能有所下降。

### (2) `BufferLogger`的使用

同样地可以直接调用内部函数实现：

```go
package main

import (
	"gitee.com/swsk33/sclog"
	"time"
)

func main() {
	// 格式化输出
	sclog.BufferTrace("基于缓冲区通道的线程安全日志，级别：%d\n", sclog.TRACE)
	sclog.BufferDebug("基于缓冲区通道的线程安全日志，级别：%d\n", sclog.DEBUG)
	sclog.BufferInfo("基于缓冲区通道的线程安全日志，级别：%d\n", sclog.INFO)
	sclog.BufferWarn("基于缓冲区通道的线程安全日志，级别：%d\n", sclog.WARN)
	sclog.BufferError("基于缓冲区通道的线程安全日志，级别：%d\n", sclog.ERROR)
	// 单行输出
	sclog.BufferTraceLine("基于缓冲区通道的线程安全日志，单行输出")
	sclog.BufferDebugLine("基于缓冲区通道的线程安全日志，单行输出")
	sclog.BufferInfoLine("基于缓冲区通道的线程安全日志，单行输出")
	sclog.BufferWarnLine("基于缓冲区通道的线程安全日志，单行输出")
	sclog.BufferErrorLine("基于缓冲区通道的线程安全日志，单行输出")
	// 防止主线程提前退出
	time.Sleep(100 * time.Millisecond)
}
```

当然，也可以使用`NewBufferLogger`创建一个自定义的缓冲区线程安全日志输出器对象：

```go
package main

import (
	"gitee.com/swsk33/sclog"
	"sync"
	"time"
)

func main() {
	// 创建基于缓冲区通道的线程安全日志输出器
	bufferLogger := sclog.NewBufferLogger(10)
	bufferLogger.Level = sclog.INFO
	// 调用对应级别方法即可
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		for i := 0; i <= 10; i += 2 {
			bufferLogger.Info("[线程1] 当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 1; i <= 11; i += 2 {
			bufferLogger.Info("[线程2] 当前输出：%d\n", i)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
	// 防止主线程提前退出
	time.Sleep(1 * time.Second)
}
```

`NewBufferLogger`的参数表示日志消息缓冲区的大小，若同一时间消息超过了缓冲区大小且未被及时输出，则可能导致调用日志时发生阻塞。

基于缓冲区通道的线程安全日志比起基于互斥锁的线程安全日志输出器通常会有更好的性能，但是可能造成额外的内存开销。