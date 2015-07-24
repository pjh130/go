package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)
	fileName := "./test.txt"
	f, err := os.Stat(fileName)
	if nil != err {
		log.Println(err)
		return
	}

	fileW, err := bodyWrite.CreateFormFile("upload", f.Name())
	if nil != err {
		log.Println(err)
		return
	}

	fileR, err := os.Open(fileName)
	if nil != err {
		log.Println(err)
		return
	}

	io.Copy(fileW, fileR)

	contextType := bodyWrite.FormDataContentType()
	bodyWrite.Close()

	resp, err := http.Post("http://localhost:8080/upload", contextType, bodyBuf)
	if nil != err {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(resp_body))
}
