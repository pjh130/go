package main

import (
	// "io"
	// "io/ioutil"
	// "log"
	// "net"
	// "net/http"

	"github.com/pjh130/go/common/net/tcp_ser"
)

func main() {
	// TestServer()

	myParser := new(tcp_ser.MyParser)

	tcp_ser.StartServer("./conf/config.ini", myParser, tcp_ser.MyDo)
}

//实现一个百万连接的服务器，采用每个连接一个goroutine的模式(goroutine-per-conn)
// var epoller *epoll

// func TestServer() {
// 	setLimit()

// 	ln, err := net.Listen("tcp", ":8972")

// 	if err != nil {
// 		panic(err)

// 	}

// 	go func() {
// 		if err := http.ListenAndServe(":6060", nil); err != nil {
// 			log.Fatalf("pprof failed: %v", err)

// 		}

// 	}()

// 	epoller, err = MkEpoll()

// 	if err != nil {
// 		panic(err)

// 	}

// 	go start()

// 	for {
// 		conn, e := ln.Accept()

// 		if e != nil {
// 			if ne, ok := e.(net.Error); ok && ne.Temporary() {
// 				log.Printf("accept temp err: %v", ne)

// 				continue

// 			}

// 			log.Printf("accept err: %v", e)

// 			return

// 		}

// 		if err := epoller.Add(conn); err != nil {
// 			log.Printf("failed to add connection %v", err)

// 			conn.Close()

// 		}

// 	}

// }

// func start() {
// 	var buf = make([]byte, 8)

// 	for {
// 		connections, err := epoller.Wait()

// 		if err != nil {
// 			log.Printf("failed to epoll wait %v", err)

// 			continue

// 		}

// 		for _, conn := range connections {
// 			if conn == nil {
// 				break

// 			}

// 			if _, err := conn.Read(buf); err != nil {
// 				if err := epoller.Remove(conn); err != nil {
// 					log.Printf("failed to remove %v", err)

// 				}

// 				conn.Close()

// 			}

// 		}

// 	}

// }
