package tcp_ser

import (
	"net"
)

/*
读消息
*/
type MsgResquest struct {
	Key  string //连接唯一标识符
	Data []byte //内容
}

/*
写消息
*/
type MsgResponse struct {
	Key  string //连接唯一标识符
	Data []byte //内容
}

//消息解析接口
type Parser interface {
	Decode(conn net.Conn) ([]byte, error)
	Encode(data []byte) ([]byte, error)
}
