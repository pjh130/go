package main

import (
	"log"

	"github.com/pjh130/go/common/workerpool"
)

type Student struct {
	Age  int
	Name string
}

func (t *Student) SayHello() {
	log.Println("Hello world!")
}

func myWorker(t interface{}) error {
	// log.Println("I am working function:", t)

	log.Println("I am working function Name: ", t.(Student).Name)
	log.Println("I am working function Age: ", t.(Student).Age)
	// log.Println("I am working function Age: ", t.(Student).SayHello())

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

	select {}
}
