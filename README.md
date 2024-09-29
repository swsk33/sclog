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

