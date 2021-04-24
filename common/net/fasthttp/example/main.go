package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	foobar string
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	if false {
		//高并发的server，那这里的取到的request body就会有丢失一部分数据
		//复制一个新的request来解决这个问题
		newrequest := &fasthttp.Request{}
		ctx.Request.CopyTo(newrequest)
		//获取body
		body := newrequest.Body()
		fmt.Println(string(body))
	}

	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

// index 页
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome")
}

// 简单路由页
func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello")
}

// 获取GET请求json数据
// 使用 ctx.QueryArgs() 方法
// Peek类似与python中dict的pop方法，取某个键对应的值
func TestGet(ctx *fasthttp.RequestCtx) {
	values := ctx.QueryArgs()
	fmt.Fprint(ctx, string(values.Peek("abc"))) // 不加string返回的byte数组
}

// 获取post的请求json数据
// 这里就有点坑是，查了很多网页说可以用 ctx.PostArgs() 取post的参数，返现不行，返回空
// 后来用 ctx.FormValue() 取表单数据就好了，难道是版本升级的问题？
// ctx.PostBody() 在上传文件的时候比较有用
func TestPost(ctx *fasthttp.RequestCtx) {
	//postValues := ctx.PostArgs()
	//fmt.Fprint(ctx, string(postValues))

	// 获取表单数据
	fmt.Fprint(ctx, string(ctx.FormValue("abc")))

	// 这两行可以获取PostBody数据，在上传数据文件的时候有用
	postBody := ctx.PostBody()
	fmt.Fprint(ctx, string(postBody))
}
func fooHandlerFunc(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "foo")
}

func barHandlerFunc(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "bar")
}

func bazHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "baz")
}

func main() {
	go ExampleFS()

	go TestClient()
	go TestClient1()

	i := 3
	switch i {
	case 0:
		// pass bound struct method to fasthttp
		myHandler := &MyHandler{
			foobar: "foobar",
		}
		fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
		break
	case 1:
		// pass plain function to fasthttp
		fasthttp.ListenAndServe(":8081", fastHTTPHandler)
		break
	case 2:
		// 创建路由
		router := fasthttprouter.New()

		// 不同的路由执行不同的处理函数
		router.GET("/", Index)
		router.GET("/hello", Hello)
		router.GET("/get", TestGet)
		// post方法
		router.POST("/post", TestPost)

		// 启动web服务器，监听 0.0.0.0:12345
		log.Fatal(fasthttp.ListenAndServe(":8082", router.Handler))
		break
	case 3:
		// the corresponding fasthttp code
		m := func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/foo":
				fooHandlerFunc(ctx)
			case "/bar":
				barHandlerFunc(ctx)
			case "/baz":
				bazHandler(ctx)
			default:
				ctx.Error("not found", fasthttp.StatusNotFound)
			}
		}

		fasthttp.ListenAndServe(":8083", m)
		break
	default:
	}
	select {}
}

func ExampleFS() {
	fs := &fasthttp.FS{
		// Path to directory to serve.
		Root: "d:/test",

		// Generate index pages if client requests directory contents.
		GenerateIndexPages: true,

		// Enable transparent compression to save network traffic.
		Compress: true,
	}

	// Create request handler for serving static files.
	h := fs.NewRequestHandler()

	// Start the server.
	if err := fasthttp.ListenAndServe(":8088", h); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
