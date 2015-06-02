package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type jieguo struct {
	go_num int
	Lock   sync.Mutex
}

var searchName string
var search_result chan string = make(chan string, 1)
var nums chan int = make(chan int, 1)
var x jieguo = jieguo{go_num: 0}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("args not enough!")
		return
	} else {
		searchName = os.Args[1]
	}

	list := []string{"D:", "E:", "F"}
	for _, i := range list {
		fmt.Println(i, searchName)
		go pan(i)
	}
	for {
		select {
		case x := <-search_result:
			fmt.Println(x)
		case y := <-nums:
			if y == 0 {
				os.Exit(0)
			}
		}
	}
}

func pan(path string) {
	a := exec.Command("ls", path)
	result, e := a.Output()
	if e != nil {
		fmt.Println(e)
	}
	z := bytes.Split(result, []byte{10})
	for _, i := range z[:len(z)-1] {
		s := path + "\\" + string(i)
		fmt.Println(s)
		e = os.Chdir(s)
		if e == nil {
			x.Lock.Lock()
			x.go_num = x.go_num + 1
			x.Lock.Unlock()
			go func() {
				filepath.Walk(s, walk)
				x.Lock.Lock()
				x.go_num = x.go_num - 1
				nums <- x.go_num
				x.Lock.Unlock()
			}()
		} else {
			fmt.Println(e)
			if string(i) == searchName {
				fmt.Println(s)
			}
		}
	}
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && info.Name() == searchName {
		search_result <- path
		return nil
	}
	return nil
}
