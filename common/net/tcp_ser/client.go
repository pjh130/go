package tcp_ser

import (
	"log"
	"net"
)

const (
	MSG_LEN = 10
)

/*
 客户端结构体
*/
type Client struct {
	// 连接信息
	Key  string           //客户端连接的唯标志
	Conn net.Conn         //连接
	In   chan MsgResquest //输入消息
	Out  chan MsgResponse //输出消息
	Quit chan *Client     //退出
	Par  Parser           //需要自己实现的消息解析
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
		Key:  key,
		Conn: conn,
		In:   make(chan MsgResquest, MSG_LEN),
		Out:  make(chan MsgResponse, MSG_LEN),
		Quit: make(chan *Client),
		Par:  parser,
	}

	//开始工作
	client.Listen()

	return client
}

/*
 自动读入或者写出消息
*/
func (this *Client) Listen() {

	if nil == this.Par {
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

	this.Conn.Close()
}

/*
 读取消息
*/
func (this *Client) read() {
	for {
		data, err := this.Par.Decode(this.Conn)
		if err == nil {
			//解析出来消息放入缓存中
			req := MsgResquest{
				Key:  this.Key,
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
		send, err := this.Par.Encode(resp.Data)
		if nil != err {
			log.Printf("封装发送的消息失败[%s]", err)
			return
		}
		_, err = this.Conn.Write(send)
		if nil != err {
			//发生错误
			log.Printf("Write error: %s\n", err)
			this.Quiting()
			return
		}
	}
}

/*
 设置输出消息
*/
func (this *Client) PutOut(resp MsgResponse) {
	if resp.Key == this.Key {
		this.Out <- resp
	}
}
