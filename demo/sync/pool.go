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

func pool1() {
	log.Println("pool1 start ...")
	log.Println("pool1  end  ...")
}
