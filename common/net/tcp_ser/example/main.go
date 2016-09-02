package main

import (
	"log"
	"net"

	"github.com/pjh130/go/common/net/tcp_ser"
)

func main() {
	cfg, err := tcp_ser.ReadConfig("./conf/config.ini")
	if nil != err {
		log.Println(err)
		return
	}
	my := new(MyParser)
	tcp_ser.StartServer(*cfg, my)
}

type MyParser struct {
}

func (this *MyParser) Decode(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	return buf, err
}

func (this *MyParser) Encode(data []byte) ([]byte, error) {
	return data, nil
}
