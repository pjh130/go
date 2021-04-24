package http_ser

import (
	"fmt"
	"log"
	"github.com/valyala/fasthttp"
)

func fasthttpExample() {
	flag.Parse()

	addr: = "12.0.0.1:8080"
	if err := fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")
}

func TestFasthttpClient1() {
	statusCode, body, err := fasthttp.GetTimeout(nil, "http://localhost:60000", 10*time.Second)
	if err != nil {
		log.Fatalf("Error w: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	log.Println(len(body))
}

func TestFasthttpClient() {
	c := &fasthttp.Client{
		MaxIdleConnDuration: 10,
	}
	// statusCode, body, err := c.Get(nil, "http://www.baidu.com/")
	// statusCode, body, err := c.Get(nil, "http://localhost:60000")
	statusCode, body, err := c.GetTimeout(nil, "http://localhost:60000", 10*time.Second)
	if err != nil {
		log.Fatalf("Error w: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}

	log.Println(string(body))
}
