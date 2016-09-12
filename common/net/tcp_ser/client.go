package tcp_ser

import (
	"errors"
	"log"
	"net"
)

const (
	MSG_LEN = 10
)

const (
	ConnStatus_Connected = iota
	ConnStatus_Disconnected
)

/*
 客户端结构体
*/
type Client struct {
	// 连接信息
	key    string //客户端连接的唯标志
	status int
	conn   net.Conn         //连接
	In     chan MsgRequest  //输入消息
	Out    chan MsgResponse //输出消息
	Quit   chan *Client     //退出
	parser Parser           //需要自己实现的消息解析
}

/*
 客户端列表
*/
type ClientTable map[string]*Client

/*
 创建客户端
*/
func CreateClient(key string, conn net.Conn, parser Parser) *Client {
	client := &Client{
		key:    key,
		conn:   conn,
		status: ConnStatus_Connected,
		In:     make(chan MsgRequest, MSG_LEN),
		Out:    make(chan MsgResponse, MSG_LEN),
		Quit:   make(chan *Client),
		parser: parser,
	}

	//开始工作
	client.Listen()

	return client
}

/*
 自动读入或者写出消息
*/
func (this *Client) Listen() {

	if nil == this.parser {
		log.Println("Parser 不能为空")
		this.Close()
		return
	}
	go this.read()
	go this.write()
}

/*
 退出了一个连接
*/
func (this *Client) Quiting() {
	this.Quit <- this
}

/*
 关闭连接通道
*/
func (this *Client) Close() {
	//	close(this.In)
	//	close(this.Out)
	//	close(this.Quit)
	if this.status != ConnStatus_Connected {
		return
	}

	this.status = ConnStatus_Disconnected
	this.conn.Close()
}

/*
 读取消息
*/
func (this *Client) read() {
	for {
		//链接不可用
		if this.status != ConnStatus_Connected {
			return
		}

		data, err := this.parser.Decode(this.conn)
		if err == nil {
			//解析出来消息放入缓存中
			req := MsgRequest{
				Key:  this.key,
				Data: data,
			}
			this.In <- req
		} else {
			this.Quiting()
			return
		}
	}
}

/*
 输出消息
*/
func (this *Client) write() {
	for resp := range this.Out {
		//链接不可用
		if this.status != ConnStatus_Connected {
			return
		}

		send, err := this.parser.Encode(resp.Data)
		if nil != err {
			log.Printf("封装发送的消息失败[%s]", err)
			return
		}
		_, err = this.conn.Write(send)
		if nil != err {
			//发生错误
			log.Printf("Write error: %s\n", err)
			this.Quiting()
			return
		}
	}
}

/*
获取输出消息
*/
func (this *Client) GetIn() MsgRequest {
	req := <-this.In
	return req
}

/*
 设置输出消息
*/
func (this *Client) PutOut(resp MsgResponse) error {
	if this.status != ConnStatus_Connected {
		return errors.New("Connection is closed.")
	}

	if resp.Key == this.key {
		this.Out <- resp
	}

	return nil
}

/*
获取key
*/
func (this *Client) GetKey() string {
	return this.key
}

/*
设置Key
*/
func (this *Client) SetKey(key string) {
	this.key = key
}

func (this *Client) GetStatus() int {
	return this.status
}
