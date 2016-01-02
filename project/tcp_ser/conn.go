package main

import (
	"bufio"
	"errors"
	"github.com/coocood/freecache"
	"github.com/pjh130/go/common/uuidlib"
	"log"
	"net"
	"sync"
	"time"
)

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
	mu       sync.Mutex
	identity string
	conn     net.Conn
	err      error
	// Read
	readTimeout time.Duration
	br          *bufio.Reader

	// Write
	writeTimeout time.Duration
	bw           *bufio.Writer

	recv  DataRecv
	reply chan []byte
	cache *freecache.Cache
}

func NewClient(netConn net.Conn, readTimeout, writeTimeout time.Duration) *Client {
	c := new(Client)
	c.identity = uuidlib.NewV4().String()
	c.conn = netConn
	c.bw = bufio.NewWriter(netConn)
	c.br = bufio.NewReader(netConn)
	c.readTimeout = readTimeout
	c.writeTimeout = writeTimeout
	c.reply = make(chan []byte, 100)

	return c
}

func (c *Client) Close() error {
	c.mu.Lock()
	err := c.err
	if c.err == nil {
		c.err = errors.New("client: closed")
		err = c.conn.Close()
	}
	c.mu.Unlock()
	return err
}

func (c *Client) fatal(err error) error {
	c.mu.Lock()
	if c.err == nil {
		c.err = err
		// Close connection to force errors on subsequent calls and to unblock
		// other reader or writer.
		c.conn.Close()
	}
	c.mu.Unlock()
	return err
}

func (c *Client) readLine() ([]byte, error) {
	return c.br.ReadSlice('\n')
}

func (c *Client) readBytes(v []byte) (int, error) {
	return c.br.Read(v)
}

func (c *Client) writeString(s string) (int, error) {
	return c.bw.WriteString(s)
}

func (c *Client) writeBytes(p []byte) (int, error) {
	return c.bw.Write(p)
}

func (c *Client) Start() {
	go c.readLoop()
	go c.writeLoop()
}

func (c *Client) readLoop() {
	var v []byte = make([]byte, 1024)
	for {
		n, err := c.readBytes(v)
		//		n, err := c.conn.Read(v)
		//发生错误就清理资源退出循环
		if nil != err {
			log.Printf("[%v] read err: %v\n", c.identity, err)
			//关闭通道会触发writeLoop的错误，退出writeLoop的循环
			close(c.reply)
			return
		}
		if n > 0 {
			//添加数据
			c.recv.appendData(v[:n])

			//打印数据
			c.recv.printData()

			//处理数据
			_, err := c.recv.parseData(nil)
			if nil == err {
				//				c.cache.Set([]byte("key"), data, 0)
			}

			//处理返回
			reply := []byte("OK")
			c.reply <- reply
		}
	}
}

func (c *Client) writeLoop() {
	for {
		select {
		case reply, ok := <-c.reply:
			if !ok {
				log.Printf("[%v] reply chan fail\n", c.identity)
				//readLoop中发生错误会关闭通道触发
				c.Close()
				return
			}

			_, err := c.writeBytes(reply)
			//			_, err := c.conn.Write(reply)
			if nil != err {
				log.Printf("[%v] write err: %v\n", c.identity, err)
				//关闭链接会触发readLoop关闭通道
				c.Close()
				return
			}
		}
	}
}
