package main

import (
	"fmt"
	"strings"
	"time"
)

var formatReplacer = strings.NewReplacer([]string{
	"yyyy", "2006",
	"yy", "06",
	"mm", "01",
	"dd", "02",
	"HH", "15",
	"MM", "04",
	"SS", "05",
}...)

func main() {
	s := "yyyy-mm-dd HH:MM:SS"
	fmt.Println(formatReplacer.Replace(s))
	s = "2006-01-02 15:04:05"
	fmt.Println(formatReplacer.Replace(s))

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
