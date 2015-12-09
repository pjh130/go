package main

import (
	"log"
)

func main() {
	log.Println("====tcp server====")

	startServer()

	bExit := make(chan bool)

	<-bExit
}
