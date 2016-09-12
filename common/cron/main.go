package main

import (
	"fmt"

	cronB "github.com/jakecoffman/cron"
	cronA "github.com/robfig/cron"
)

func main() {
	UseCorn1()
	UseCorn2()

	select {}
}

//https://github.com/robfig/cron/blob/master/doc.go
func UseCorn1() {
	c := cronA.New()
	c.AddFunc("30 * * * * *", func() { fmt.Println("Every minute on the half minute") })
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	// Funcs are invoked in their own goroutine, asynchronously.

	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Inspect the cron job entries' next and previous run times.
	//	inspect(c.Entries())

	// Stop the scheduler (does not stop any jobs already running).
	//	c.Stop()
}

//第二个是借鉴了第一个例子
func UseCorn2() {
	c := cronB.New()
	c.AddFunc("0 5 * * * *", func() { fmt.Println("Every 5 minutes") }, "Often")
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") }, "Frequent")
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") }, "Less Frequent")
	c.Start()

	// Funcs are invoked in their own goroutine, asynchronously.

	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") }, "My Job")

	// Inspect the cron job entries' next and previous run times.
	//inspect(c.Entries())

	// Remove an entry from the cron by name.
	c.RemoveJob("My Job")

	//c.Stop()
}
