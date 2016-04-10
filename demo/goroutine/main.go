package main

import (
	"log"
	"sync"
	"time"
)

var mux sync.Mutex
var Count int = 0

func main() {
	log.Println("Begin")

	for i := 0; i < 100; i++ {
		go OneRoutine(i)
	}

	time.Sleep(30 * time.Second)
	log.Println("End")
}

func OneRoutine(i int) {
	mux.Lock()
	log.Println(i, ":", Count)
	Count++
	mux.Unlock()
}
