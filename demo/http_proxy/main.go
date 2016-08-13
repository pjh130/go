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
)

func main() {
	log.Println("Enter")

	//server
	go ProxyServer()

	//client
	go ProxyClient()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-signalChan

	log.Println("Quit")
}

func ProxyClient() {
	proxy_addr := "http://localhost:33333"
	url_addr := "http://baidu.com"
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
	//	w.Write([]byte("Hello world!"))
	//	return
	log.Println("url:", r.URL.String())
	log.Println("url path:", r.URL.Path)

	//	log.Println(r)
	request, err := http.NewRequest(r.Method, r.URL.String(), nil)
	for k, v := range r.Header {
		for _, vv := range v {
			request.Header.Add(k, vv)
		}
	}

	for _, c := range r.Cookies() {
		request.Header.Add("Set-Cookie", c.Raw)
	}

	res, err := http.DefaultClient.Do(request)
	// res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("Do"))
		return
	}
	defer res.Body.Close()

	for k, v := range res.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	for _, c := range res.Cookies() {
		w.Header().Add("Set-Cookie", c.Raw)
	}
	w.WriteHeader(res.StatusCode)
	result, err := ioutil.ReadAll(res.Body)
	if err != nil && err != io.EOF {
		log.Println(err.Error())
		w.Write([]byte("ReadAll"))
		return
	}
	defer res.Body.Close()

	w.Write(result)
	log.Println("==========end==========")
	log.Println(string(result))
}
