package tcp_ser

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"net"
)

//业务处理函数样例
func MyDo(req MsgResquest) []MsgResponse {
	var res []MsgResponse
	r := MsgResponse{
		Key:  req.Key,
		Data: []byte("OK"),
	}
	res = append(res, r)

	return res
}

//////////////////解析数据包的接口实现样例1//////////////////
type MyParser struct {
}

func (this *MyParser) Decode(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 4096)
	_, err := conn.Read(buf)
	return buf, err
}

func (this *MyParser) Encode(data []byte) ([]byte, error) {
	return data, nil
}

//////////////////解析数据包的接口实现样例2//////////////////
//数据长度类型
const (
	MSG_LEN_UINT8  = 1
	MSG_LEN_UINT16 = 2
	MSG_LEN_UINT32 = 4
)

// --------------
// | len | data |
// --------------
type MsgParser struct {
	lenMsgLen    int
	minMsgLen    uint32
	maxMsgLen    uint32
	littleEndian bool
}

func NewMsgParser() *MsgParser {
	p := new(MsgParser)
	p.lenMsgLen = 2
	p.minMsgLen = 1
	p.maxMsgLen = 4096
	p.littleEndian = false

	return p
}

//消息读取
func (p *MsgParser) Decode(conn net.Conn) ([]byte, error) {
	var b [4]byte
	bufMsgLen := b[:p.lenMsgLen]

	// read len
	if _, err := io.ReadFull(conn, bufMsgLen); err != nil {
		return nil, err
	}

	// parse len
	var msgLen uint32
	switch p.lenMsgLen {
	case MSG_LEN_UINT8:
		msgLen = uint32(bufMsgLen[0])
	case MSG_LEN_UINT16:
		if p.littleEndian {
			msgLen = uint32(binary.LittleEndian.Uint16(bufMsgLen))
		} else {
			msgLen = uint32(binary.BigEndian.Uint16(bufMsgLen))
		}
	case MSG_LEN_UINT32:
		if p.littleEndian {
			msgLen = binary.LittleEndian.Uint32(bufMsgLen)
		} else {
			msgLen = binary.BigEndian.Uint32(bufMsgLen)
		}
	}

	// check len
	if msgLen > p.maxMsgLen {
		return nil, errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return nil, errors.New("message too short")
	}

	// data
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}

// goroutine safe
func (p *MsgParser) Encode(data []byte) ([]byte, error) {

	// get len
	var msgLen uint32
	msgLen = uint32(len(data))

	// check len
	if msgLen > p.maxMsgLen {
		return nil, errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return nil, errors.New("message too short")
	}

	msg := make([]byte, uint32(p.lenMsgLen)+msgLen)

	//计算消息头的长度并赋值
	switch p.lenMsgLen {
	case MSG_LEN_UINT8:
		msg[0] = byte(msgLen)
	case MSG_LEN_UINT16:
		if p.littleEndian {
			binary.LittleEndian.PutUint16(msg, uint16(msgLen))
		} else {
			binary.BigEndian.PutUint16(msg, uint16(msgLen))
		}
	case MSG_LEN_UINT32:
		if p.littleEndian {
			binary.LittleEndian.PutUint32(msg, msgLen)
		} else {
			binary.BigEndian.PutUint32(msg, msgLen)
		}
	}

	//添加原始数据
	l := p.lenMsgLen
	copy(msg[l:], data)

	return msg, nil
}

// It's dangerous to call the method on reading or writing
func (p *MsgParser) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// It's dangerous to call the method on reading or writing
func (p *MsgParser) SetMsgLen(lenMsgLen int, minMsgLen uint32, maxMsgLen uint32) {
	if lenMsgLen == 1 || lenMsgLen == 2 || lenMsgLen == 4 {
		p.lenMsgLen = lenMsgLen
	}
	if minMsgLen != 0 {
		p.minMsgLen = minMsgLen
	}
	if maxMsgLen != 0 {
		p.maxMsgLen = maxMsgLen
	}

	var max uint32
	switch p.lenMsgLen {
	case MSG_LEN_UINT8:
		max = math.MaxUint8
	case MSG_LEN_UINT16:
		max = math.MaxUint16
	case MSG_LEN_UINT32:
		max = math.MaxUint32
	}
	if p.minMsgLen > max {
		p.minMsgLen = max
	}
	if p.maxMsgLen > max {
		p.maxMsgLen = max
	}
}
