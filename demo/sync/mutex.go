package main

import (
	"log"
	"sync"
)

var mutex sync.Mutex
var num int = 0

func mutex1() {
	for i := 0; i < MAX_NUM; i++ {
		go addMutex()
	}
}

func addMutex() {
	mutex.Lock()
	num++
	log.Println("num: ", num)
	mutex.Unlock()
}
