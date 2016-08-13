package main

import (
	"bytes"
	"log"
	"sync"
)

var bp sync.Pool

func init() {
	bp.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

func bufferPoolGet() *bytes.Buffer {
	return bp.Get().(*bytes.Buffer)
}

func bufferPoolPut(b *bytes.Buffer) {
	bp.Put(b)
}

func main() {
	b := bp.Get().(*bytes.Buffer)
	//第一次 初始化和填充数据
	b.Reset()
	b.WriteString("hello world!")
	bp.Put(b)

	b = bp.Get().(*bytes.Buffer)
	log.Println(b.String())

	//重置但是不Put数据进去
	b.Reset()
	b = bp.Get().(*bytes.Buffer)
	log.Println(b.String())

	//重置和改变数据
	b.Reset()
	b.WriteString("Bye Bye!")
	bp.Put(b)

	b = bp.Get().(*bytes.Buffer)
	log.Println(b.String())

	//不重置继续添加数据
	b.WriteString("Ok Ok!")
	bp.Put(b)

	b = bp.Get().(*bytes.Buffer)
	log.Println(b.String())
}
