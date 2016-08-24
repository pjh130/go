package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("====tcp server====")

	server := NewServer()
	server.Start(":60000")

	ch := make(chan os.Signal, 1)

	//	signal.Notify(ch, os.Interrupt, os.Kill)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	c := <-ch
	log.Println(c)
	//	bExit := make(chan bool)
	//	<-bExit
}
