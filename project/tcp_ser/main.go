package main

import (
	"log"
)

func main() {
	log.Println("====tcp server====")

	server := NewServer()
	server.Start(":60000")

	bExit := make(chan bool)

	<-bExit
}
