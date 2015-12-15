package main

import (
	"github.com/coocood/freecache"
	"log"
	"net"
	"sync"
	"time"
)

type Server struct {
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

		client := new(Client)
		client.conn = conn
		client.addr = conn.RemoteAddr().String()
		client.reply = make(chan []byte, 100)

		log.Println("Accept new client:", client.addr)

		go client.readLoop()
		go client.writeLoop()
	}
}

type DataRecv struct {
	mutex sync.Mutex
	recv  []byte
}

func (this *DataRecv) appendData(data []byte) {
	this.mutex.Lock()
	this.recv = append(this.recv, data...)
	this.mutex.Unlock()
}

func (this *DataRecv) parseData(value interface{}) ([]byte, error) {
	return nil, nil
}

func (this *DataRecv) printData() {
	this.mutex.Lock()
	log.Println(string(this.recv))
	this.mutex.Unlock()
}

type Client struct {
	conn  net.Conn
	addr  string
	recv  DataRecv
	reply chan []byte
	cache *freecache.Cache
}

func (this *Client) readLoop() {
	var v []byte = make([]byte, 1024)
	for {
		n, err := this.conn.Read(v)
		//发生错误就清理资源退出循环
		if nil != err {
			log.Println("read err:", err)
			//关闭通道会触发writeLoop的错误，退出writeLoop的循环
			close(this.reply)
			return
		}
		if n > 0 {
			//添加数据
			this.recv.appendData(v[:n])

			//打印数据
			this.recv.printData()

			//处理数据
			data, err := this.recv.parseData(nil)
			if nil == err {
				this.cache.Set([]byte("key"), data, 0)
			}

			//处理返回
			reply := []byte("OK")
			this.reply <- reply
		}
	}
}

func (this *Client) writeLoop() {
	for {
		select {
		case reply, ok := <-this.reply:
			if !ok {
				log.Println("reply chan fail")
				//readLoop中发生错误会关闭通道触发
				this.conn.Close()
				return
			}

			_, err := this.conn.Write(reply)
			if nil != err {
				log.Println("write err:", err)
				//关闭链接会触发readLoop关闭通道
				this.conn.Close()
				return
			}
		}
	}
}
