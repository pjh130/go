package main

import (
	"log"
    "encoding/binary"
    "net"
	"io"
)

func Client(ip string) {
    conn, err := net.Dial("tcp", ip)
    if err != nil {
        panic(err)
    }

    data := []byte("hello world")

    // len + data
    m := make([]byte, 2+len(data))

    // 默认使用大端序
    binary.BigEndian.PutUint16(m, uint16(len(data)))

    copy(m[2:], data)

    // 发送消息
    conn.Write(m)
	
//	b := make([]byte, 1024)
//	_, err = conn.Read(b)
//	if nil == err {
//		log.Println(string(b))
//	} else {
//		log.Println(err)
//	}

	BufLength := 1024
	
		buf := make([]byte, BufLength)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				log.Println(err)
			}
			log.Println(string(buf[2:n]))
		}
}