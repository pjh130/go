package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/elazarl/goproxy"
)

func main() {
	log.Println("Enter")

	//server
	go ProxyServer()

	//client
	go ProxyClient()

	go func() {
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		log.Fatal(http.ListenAndServe(":8888", proxy))
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-signalChan

	log.Println("Quit")
}

func ProxyClient() {
	proxy_addr := "http://localhost:33333"
	url_addr := "http://www.baidu.com"
	//	url_addr := "http://localhost:60000"

	request, err := http.NewRequest("GET", url_addr, nil)
	if err != nil {
		log.Println("NewRequest: ", err)
		return
	}

	if false {
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Println("Do: ", err)
			return
		}

		//		log.Println(resp)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	} else {
		proxy, err := url.Parse(proxy_addr)
		if err != nil {
			log.Println("Parse: ", err)
			return
		}
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
		resp, err := client.Do(request)
		if err != nil {
			log.Println("Do: ", err)
			return
		}

		//		log.Println(resp)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}
}

func ProxyServer() {
	http.HandleFunc("/", Handler)
	port := "33333"
	log.Println("Starting agent: ", port)
	http.ListenAndServe(":"+port, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	log.Println("URL:", r.URL.String())
	log.Println("URL.Path:", r.URL.Path)
	log.Println("URL.Scheme:", r.URL.Scheme)

	req, _ := http.NewRequest(r.Method, "", r.Body)
	req.URL = r.URL
	req.URL.Host = r.Host
	req.URL.Scheme = "http"

	for _, v := range r.Cookies() {
		req.AddCookie(v)
	}

	//req.Header = r.Header 这里的Header就不要使用了,使用的话他会自动跳转到https,代理就出问题了.

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Here:", err)
		return
	}
	for k, v := range resp.Header {
		for _, value := range v {
			w.Header().Add(k, value)
		}
	}
	for _, cookie := range resp.Cookies() {
		w.Header().Add("Set-Cookie", cookie.Raw)
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()

	r.Body.Close()

	log.Println("==========end==========")
}
