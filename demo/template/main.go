package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/pjh130/go/common/debug/pprof"
)

func main() {
	//输出结构体和map
	if true {
		Test2()
		return
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
	go pprof.StartAdminHttp(":8081")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)

	//	bExit := make(chan bool)
	//	<-bExit
}
