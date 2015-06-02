package main

import (
	"fmt"
	"time"
)

func main() {
	timeTest1()
}

func timeTest1() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().String())
		}
	}
}
