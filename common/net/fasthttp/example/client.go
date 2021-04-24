package main

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func TestClient1() {
	statusCode, body, err := fasthttp.GetTimeout(nil, "http://localhost:60000", 10*time.Second)
	if err != nil {
		log.Printf("Error w: %s", err)
		return
	}
	if statusCode != fasthttp.StatusOK {
		log.Printf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
		return
	}
	log.Println(len(body))
}

func TestClient() {
	c := &fasthttp.Client{
		MaxIdleConnDuration: 10,
	}
	// statusCode, body, err := c.Get(nil, "http://www.baidu.com/")
	// statusCode, body, err := c.Get(nil, "http://localhost:60000")
	statusCode, body, err := c.GetTimeout(nil, "http://localhost:60000", 10*time.Second)
	if err != nil {
		log.Printf("Error w: %s", err)
		return
	}
	if statusCode != fasthttp.StatusOK {
		log.Printf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
		return
	}

	log.Println(string(body))
}
