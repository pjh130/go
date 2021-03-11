package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	// "os"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

type echoServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if true {
		fmt.Println(string(frame))
		// out = frame
		ret := []byte("I'm server: ")
		var b bytes.Buffer
		b.Write(ret)
		b.Write(frame)
		out = b.Bytes()
	} else {
		data := append([]byte{}, frame...)
		// Use ants pool to unblock the event-loop.
		_ = es.pool.Submit(func() {
			time.Sleep(1 * time.Second)
			c.AsyncWrite(data)
			// c.AsyncWrite(frame)
		})
	}

	return
}

func main() {
	go test1()
	select {}
}

func test1() {
	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo, "tcp://:60000", gnet.WithMulticore(true)))

}
