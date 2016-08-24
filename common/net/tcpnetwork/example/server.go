package main

import (
	"log"
	"time"

	"github.com/pjh130/go/common/net/tcpnetwork"
)

func main() {
	log.Println("Start...")

	//	//新建一个server
	ser := tcpnetwork.CreateServer(tcpnetwork.MyAgentFunc)

	ser.Start()

	//	//客户端发送消息
	go Client(ser.Addr)

	//设置延迟退出，让服务端和客户端完成交互
	select {
	case <-time.After(5 * time.Second):
	}

	log.Println("End...")
}
