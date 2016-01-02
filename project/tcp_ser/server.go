package main

import (
	"errors"
	//	"github.com/pjh130/go/common/uuidlib"
	"log"
	"net"
	"sync"
	"time"
)

type Server struct {
	mutex    sync.Mutex
	connPool map[string]*Client
}

func NewServer() (server *Server) {
	server = new(Server)
	return
}

func (this *Server) Start(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	log.Println("Listening on port", addr)

	var tempDelay time.Duration
	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("accept error: %v; retrying in %v", err, tempDelay)
				continue
			}
			log.Println(err)
			return
		}
		tempDelay = 0

		client := NewClient(conn, 0, 0)

		this.mutex.Lock()
		//		this.connPool[client.identity] = client
		this.mutex.Unlock()

		log.Println("Accept new client:", conn.RemoteAddr())

		go client.readLoop()
		go client.writeLoop()
	}
}

func (this *Server) CloseConn(identity string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	client, ok := this.connPool[identity]
	if false == ok {
		return
	}
	log.Printf("[%v] close\n", identity)
	delete(this.connPool, identity)
	client.cache = nil
	client.conn.Close()
	//	close(client.reply)
}

func (this *Server) SendMsg(identity string, data []byte) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	client, ok := this.connPool[identity]
	if false == ok {
		return errors.New("identity not exist")
	}

	//防止写入超时
	select {
	case <-time.After(time.Second * 3):
		log.Printf("[%v] write channel timeout\n", identity)
		return errors.New("write channel timeout")
	case client.reply <- data:
	}
	return nil
}
