package main

import (
	"github.com/pjh130/go/common/net/tcpnetwork"
	"log"
	"time"
)

func main(){
	log.Println("Start...")

//	//新建一个server
	tcpServer := new(tcpnetwork.TCPServer)
	tcpServer.Addr = tcpnetwork.ServerAddr
		tcpServer.MaxConnNum = tcpnetwork.MaxConnNum
		tcpServer.PendingWriteNum = tcpnetwork.PendingWriteNum
		tcpServer.LenMsgLen = tcpnetwork.LenMsgLen
		tcpServer.MaxMsgLen = tcpnetwork.MaxMsgLen
		tcpServer.LittleEndian = tcpnetwork.LittleEndian
		tcpServer.NewAgent = func(conn *tcpnetwork.TCPConn) tcpnetwork.Agent {
			a := &tcpnetwork.MyAgent{Connect:conn}
			return a
		}
		
	tcpServer.Start()
	
//	//客户端发送消息
	go Client(tcpnetwork.ServerAddr)
	
	//设置延迟退出，让服务端和客户端完成交互
	select {
		case <-time.After(5*time.Second):
	}
	
	log.Println("End...")
}
	