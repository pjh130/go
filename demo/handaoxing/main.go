package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup
	strUrl := "http://www.biquku.com/0/761/"

	lst, err := GetPartAll(strUrl)
	if nil == err {
		log.Println("len:", len(lst))
	} else {
		log.Println(err)
	}

	n := 0
	root := "./" + time.Now().Format("2006-01-02_15-04-05")
	os.MkdirAll(root, 0755)
	for key, value := range lst {
		os.MkdirAll(root+"/"+key, 0755)
		for _, item := range value {
			log.Println(key, item.name, item.url)

			w.Add(1)
			n++

			content, err := GetContent(item.url)
			if nil == err && len(content) > 0 {
				ioutil.WriteFile(root+"/"+key+"/"+item.name+".txt", []byte(content), 0755)
			} else {
				log.Println("content err:", err)
			}
			w.Done()

			if n > 5 {
				time.Sleep(3 * time.Second)
				n = 0
			}
		}
	}

	w.Wait()
}

func TestHex() {
	v, _ := hex.DecodeString("c4a0")
	log.Println(v)
}
