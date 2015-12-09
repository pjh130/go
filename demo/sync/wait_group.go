package main

import (
	"log"
	"sync"
)

const (
	MAX_NUM = 5
)

func waitGroup1() {
	log.Println("waitGroup1 start ...")
	var v sync.WaitGroup
	for i := 0; i < MAX_NUM; i++ {
		v.Add(1)
		go func() {
			log.Println("i: ", i)

			v.Done()
		}()
	}
	v.Wait()
	log.Println("waitGroup1 end   ...")
}

var wait sync.WaitGroup

func showNum(i int) {
	log.Println("i: ", i)
	wait.Done()
}

func waitGroup2() {
	log.Println("waitGroup2 start ...")

	for i := 0; i < MAX_NUM; i++ {
		wait.Add(1)
		go showNum(i)
	}
	wait.Wait()
	log.Println("waitGroup2 end   ...")
}
