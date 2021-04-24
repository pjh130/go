package log

import (
	"github.com/Sirupsen/logrus"
	"github.com/aiwuTech/fileLogger"
	"github.com/astaxie/beego/logs"
	"github.com/cihub/seelog"
	"github.com/davyxu/golog"
	"github.com/golang/glog"
	"github.com/jeanphorn/log4go"
)

/*
fileLogger是一个基于Go开发的可自动分割文件进行备份的异步日志库
1、日志文件可按文件大小进行备份,可定制文件大小和数量
2、日志文件可按日期进行备份
*/
func exampleLileLogger() {
	logFile := fileLogger.NewDefaultLogger("./log", "test.log")
	//	logFile := fileLogger.NewDailyLogger("./log", "test.log")
	//	logFile := fileLogger.NewSizeLogger("./log", "test.log")
	logFile.SetLogLevel(fileLogger.INFO) //trace log will not be print

	for i := 1; i <= 1000; i++ {
		logFile.T("This is the No[%v] TRACE log using fileLogger that written by aiwuTech.", i)
		logFile.I("This is the No[%v] INFO log using fileLogger that written by aiwuTech.", i)
		logFile.W("This is the No[%v] WARN log using fileLogger that written by aiwuTech.", i)
		logFile.E("This is the No[%v] ERROR log using fileLogger that written by aiwuTech.", i)
	}
}

/*
这是一个用来处理日志的库，目前支持的引擎有 file、console、net、smtp
*/
func exampleLogs() {
	//初始化 log 变量（10000 表示缓存的大小）：
	log := logs.NewLogger(10000)

	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	log.EnableFuncCallDepth(true)

	//如果你的应用自己封装了调用log包,那么需要设置SetLogFuncCallDepth,默认是2,也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	log.SetLogFuncCallDepth(3)

	//添加输出引擎（log 支持同时输出到多个引擎）
	log.SetLogger("console", `{"level":1}`)
	log.SetLogger("file", `{"filename":"test.log"}`)
	log.SetLogger("conn", `{"net":"tcp","addr":":7020"}`)
	log.SetLogger("smtp", `{"username":"beegotest@gmail.com","password":"xxxxxxxx","host":"smtp.gmail.com:587","sendTos":["xiemengjun@gmail.com"]}`)

	//使用
	log.Trace("trace %s %s", "param1", "param2")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")
	log.Critical("critical")
}

/*
Console writer
File writer
Buffered writer (Chunk writer)
Rolling log writer (Logging with rotation)
SMTP writer

https://github.com/cihub/seelog-examples
https://github.com/cihub/seelog/wiki/Example-config
*/
func exampleSeelog() {
	defer seelog.Flush()
	seelog.Info("Hello from Seelog!")
}

func exampleLogrus() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

/*
Leveled execution logs for Go.

This is an efficient pure Go implementation of leveled logs in the
manner of the open source C++ package
*/
func example5() {
	glog.Info("Prepare to repel boarders")

	//	glog.Fatalf("Initialization failed: %s", err)
}

func exampleGolog() {
	// // 基本使用
	// var log *golog.Logger = golog.New("test")
	// log.Debugln("hello world")

	// // 层级设置
	// golog.SetLevelByString("test", "info")
	var log = golog.New("test2")
	colorStyle := `
	{
		"Rule":[
			{"Text":"panic:","Color":"Red"},
			{"Text":"[DB]","Color":"Green"},
			{"Text":"#http.listen","Color":"Blue"},
			{"Text":"#http.recv","Color":"Blue"},
			{"Text":"#http.send","Color":"Purple"},
	
			{"Text":"#tcp.listen","Color":"Blue"},
			{"Text":"#tcp.accepted","Color":"Blue"},
			{"Text":"#tcp.closed","Color":"Blue"},
			{"Text":"#tcp.recv","Color":"Blue"},
			{"Text":"#tcp.send","Color":"Purple"},
			{"Text":"#tcp.connected","Color":"Blue"},
	
			{"Text":"#udp.listen","Color":"Blue"},
			{"Text":"#udp.recv","Color":"Blue"},
			{"Text":"#udp.send","Color":"Purple"},
	
			{"Text":"#rpc.recv","Color":"Blue"},
			{"Text":"#rpc.send","Color":"Purple"},
	
			{"Text":"#relay.recv","Color":"Blue"},
			{"Text":"#relay.send","Color":"Purple"}
		]
	}
	`
	golog.SetColorDefine(".", colorStyle)

	// 默认颜色是关闭的
	log.SetParts()
	golog.EnableColorLogger(".", true)
	log.Debugln("关闭所有部分样式")

	log.SetParts(golog.LogPart_CurrLevel)
	log.SetColor("blue")
	log.Debugln("蓝色的字+级别")

	log.SetParts(golog.LogPart_CurrLevel, golog.LogPart_Name)
	// 颜色只会影响一行
	log.SetColor("red")
	log.Warnf("级别颜色高于手动设置 + 日志名字")

	log.SetParts(golog.LogPart_CurrLevel, golog.LogPart_Name, golog.LogPart_Time, golog.LogPart_ShortFileName)
	log.Debugln()
	log.Debugf("[DB] DB日志是绿色的，从文件读取，按文字匹配的， 完整的日志样式")

	log.SetParts(golog.LogPart_TimeMS, golog.LogPart_LongFileName, func(l *golog.Logger) {
		l.WriteRawString("固定头部: ")
	})

	log.SetColor("purple")

	log.Debugf("自定义紫色 + 固定头部内容")
	log.Debugf("自定义紫色 + 固定头部内容2")
}

/*
日志输出到终端
日志输出到文件，支持按大小和时间切片
日志输出到网络
日志异步输出
支持json文件配置
日志分类
不同类别的日志，输出到不同的printer中.
兼容老的日志方式
*/
func exampleLog4go() {
	// load config file, it's optional
	// or log.LoadConfiguration("./example.json", "json")
	// config file could be json or xml
	log4go.LoadConfiguration("./example.json")

	log4go.LOGGER("Test").Info("category Test info test ...")
	log4go.LOGGER("Test").Info("category Test info test message: %s", "new test msg")
	log4go.LOGGER("Test").Debug("category Test debug test ...")

	// Other category not exist, test
	log4go.LOGGER("Other").Debug("category Other debug test ...")

	// socket log test
	log4go.LOGGER("TestSocket").Debug("category TestSocket debug test ...")

	// original log4go test
	log4go.Info("normal info test ...")
	log4go.Debug("normal debug test ...")

	log4go.Close()
}
