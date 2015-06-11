package main

import (
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8888"
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Println("[UDPERR]", "net.ResolveUDPAddr: ", err.Error())
		return
	}
	udpconn, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		log.Println("[UDPERR]", "net.ListenUDP: ", err.Error())
		return
	} else {
		log.Println("net.ListenUDP: ", addr)
	}
	defer udpconn.Close()

	var buf [2048]byte
	for {
		n, remote_addr, err := udpconn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Println("[UDPERR]", "udp receiveFrom ", remote_addr, "faild: ", err.Error())
			continue
		} else {
			log.Println("[UDP]", remote_addr, "Recv:", n, "Bytes")
			log.Println("[UDP]", remote_addr, string(buf[:n]))
			_, err = udpconn.WriteTo([]byte("ok"), remote_addr)
			if nil != err {
				log.Println("[UDPERR] ", err)
			}
		}
	}
}
