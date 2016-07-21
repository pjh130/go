package main

import (
    "encoding/binary"
    "net"
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
}