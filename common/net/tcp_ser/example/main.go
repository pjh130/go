package main

import (
	"github.com/pjh130/go/common/net/tcp_ser"
)

func main() {
	myParser := new(tcp_ser.MyParser)

	tcp_ser.StartServer("./conf/config.ini", myParser, tcp_ser.MyDo)
}
