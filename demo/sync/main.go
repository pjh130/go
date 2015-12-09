package main

import (
	"time"
)

func main() {

	waitGroup1()

	waitGroup2()

	mutex1()

	<-time.After(5 * time.Second)
}
