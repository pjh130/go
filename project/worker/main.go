package main

import (
	"time"
)

func main() {

	StartDispatcher()

	for i := 0; i < 200; i++ {
		// let's create a job with the payload
		work := Job{Id: i}
		// Push the work onto the queue.
		JobQueue <- work
	}

	ticker := time.NewTicker(6000 * time.Second)
	for {
		select {
		case <-ticker.C:
			return
		}
	}
}
