package main

import (
	"log"

	"github.com/pjh130/go/common/workerpool"
)

type Student struct {
	Age  int
	Name string
}

func myWorker(t interface{}) error {
	log.Println("I am working function:", t)

	//	log.Println("I am working function:", t.(Student).Name)

	return nil
}

func main() {

	wp := &workerpool.WorkerPool{
		WorkerFunc:      myWorker,
		MaxWorkersCount: 100,
	}

	//	var c int = 1024
	//	var c string = "hello world"
	c := Student{
		Age:  20,
		Name: "panpan",
	}

	wp.Start()

	wp.Working(c)

	ch := make(chan int)
	<-ch
}
