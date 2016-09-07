package log

import (
	"github.com/Sirupsen/logrus"
	"github.com/aiwuTech/fileLogger"
	"github.com/astaxie/beego/logs"
	"github.com/cihub/seelog"
	"github.com/golang/glog"
)

/*
fileLogger是一个基于Go开发的可自动分割文件进行备份的异步日志库
1、日志文件可按文件大小进行备份,可定制文件大小和数量
2、日志文件可按日期进行备份
*/
func example1() {
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
func example2() {
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
func example3() {
	defer seelog.Flush()
	seelog.Info("Hello from Seelog!")
}

func example4() {
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
