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

## 5，输出内容的自定义

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

- `Time`属性为结构体，用于自定义日志**时间**部分的输出，其中：
	- `Enabled` 是否输出时间部分内容，默认为`true`
	- `Pattern` 格式化时间字符串，使用的是`time.Format`的`Layout`字符串，默认为`2006-01-02 15:04:05.000`
	- `Color` 时间部分的颜色，默认为白色
- `Level`属性为结构体，用于自定义日志**级别**部分的输出，其中：
	- `Enabled` 是否输出级别部分内容，默认为`true`
	- `Color` 时间部分的颜色，默认为白色
- `Message`属性为结构体，用于自定义日志**消息**部分的输出，其中：
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