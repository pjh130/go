package main

import (
	"fmt"
	"github.com/pjh130/go/common/pproflib"
	"os"
	"os/signal"
)

func main() {
	//在标准设备中输出
	if false {
		Test1()
	}

	//输出结构体和map
	if true {
		Test2()
	}

	//html回调go中的函数
	if false {
		Test3()
	}

	//传值给html页面
	if true {
		Test4()
	}

	//服务器性能查看
	go pproflib.StartAdminHttp(":8081")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)

	//	bExit := make(chan bool)
	//	<-bExit
}
