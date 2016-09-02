package main

import (
	"log"

	"github.com/pjh130/go/common/net/tcp_ser"
)

func main() {
	cfg, err := tcp_ser.ReadConfig("./conf/config.ini")
	if nil != err {
		log.Println(err)
		return
	}
	my := new(tcp_ser.MyParser)
	tcp_ser.StartServer(*cfg, my, tcp_ser.MyDo)
}
