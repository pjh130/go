package main

import (
	"io"
	"log"
	//	"mime/multipart"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func upload(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	log.Println("Method:", method)
	if "GET" == method {
		w.Write([]byte("Please used POST method!"))
		return
	} else if "POST" == method {
		//判断是否之上传资源
		if true == strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
			//默认的上传文件大小defaultMaxMemory是32MB（32 << 20）
			//如果有需要更大的文件上传，需要设置内存大小
			err := r.ParseMultipartForm(32 << 20)
			if nil != err {
				w.Write([]byte(fmt.Sprintf("%s", err)))
				return
			}

			file, header, err := r.FormFile("upload")
			if nil != err {
				w.Write([]byte(fmt.Sprintf("%s", err)))
				return
			} else {
				defer file.Close()
				//如果有中文字符，需要先在获取的内容上转码
				v, err := url.QueryUnescape(header.Filename)
				//出错
				if nil != err {
					w.Write([]byte(fmt.Sprintf("%s", err)))
					return
				} else {
					log.Println(v)
				}
				name := "./files/" + v

				//保存文件
				fileW, err := os.Create(name)
				//出错
				if nil != err {
					w.Write([]byte(fmt.Sprintf("%s", err)))
					return
				}

				defer fileW.Close()

				_, err = io.Copy(fileW, file)
				if nil != err {
					w.Write([]byte(fmt.Sprintf("%s", err)))
					return
				} else {
					resp := "Success"
					w.Write([]byte(resp))
					log.Println(resp)
				}
			}
		} else {
			w.Write([]byte("Not support form-data!"))
			return
		}
	} else {
		w.Write([]byte("Not support method!"))
		return
	}
}

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
