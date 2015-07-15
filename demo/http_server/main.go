package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	go Server1()
	go Server2()
	go Server3()

	//循环不退出
	for {
		time.Sleep(1 * time.Second)
	}
}

/////////////////////////////////////////////////////////////////
func Server1() {
	http.HandleFunc("/", SayHello1)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func SayHello1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world 11"))
	//io.WriteString(w, "hello world 1")
}

/////////////////////////////////////////////////////////////////
func Server2() {
	mux := http.NewServeMux()
	mux.Handle("/", &Server2Handler{})
	mux.HandleFunc("/hello", SayHello2)

	wd, err := os.Getwd()
	if nil != err {
		log.Fatal(err)
	}
	//静态文件服务器
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8082", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type Server2Handler struct{}

func (*Server2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "url: "+r.URL.String())
}

func SayHello2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world 2")
}

/////////////////////////////////////////////////////////////////
func Server3() {
	server := http.Server{
		Addr:         ":8083",
		Handler:      &myHandler3{},
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	router = make(map[string]func(w http.ResponseWriter, r *http.Request))
	router["/router1"] = Router1
	router["/router2"] = Router2
	router["/router3"] = Router2

	err := server.ListenAndServe()
	if nil != err {
		log.Fatal(err)
	}
}

type myHandler3 struct{}

var router map[string]func(w http.ResponseWriter, r *http.Request)

func (*myHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "hello world 3")

	if h, ok := router[r.URL.String()]; ok {
		h(w, r)
	} else if strings.HasPrefix(r.URL.String(), "/static") {
		//文件服务器
		had := http.StripPrefix("/static", http.FileServer(http.Dir("static")))
		had.ServeHTTP(w, r)
	} else {
		io.WriteString(w, "no router")
	}
}

func Router1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "router 1")
}

func Router2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "router 2")
}

func Router3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "router 3")
}
