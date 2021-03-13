package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/pjh130/go/common/debug/pprof"
)

func main() {
	v := 4
	switch v {
	case 1:
		break
	case 2:
		Test2() //输出结构体和map
		break
	case 3:
		Test3() //html回调go中的函数
		break
	case 4:
		Test4() //传值给html页面
		break
	case 5:
		break
	default:

	}
	// return
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
