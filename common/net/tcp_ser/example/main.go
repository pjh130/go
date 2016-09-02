package main

import (
	"github.com/pjh130/go/common/net/tcp_ser"
)

func main() {
	my := new(tcp_ser.MyParser)

	tcp_ser.StartServer("./conf/config.ini", my, tcp_ser.MyDo)
}
