package main

import (
	"io"
	"log"
	//	"mime/multipart"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	UPLOAD_DIR = "./files"
	VIEWS_DIR  = "./views"
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
				name := UPLOAD_DIR + "/" + v

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

//封装一层安全机制处理的函数
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				//                log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

//上传的文件
func files(w http.ResponseWriter, r *http.Request) {
	//文件服务器
	had := http.StripPrefix("/files", http.FileServer(http.Dir("files")))
	had.ServeHTTP(w, r)
}

func uploadHtml(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(VIEWS_DIR + "/upload.html")
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, "")
}

func listHtml(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles(VIEWS_DIR + "/list.html")
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := []string{}

	filepath.Walk(UPLOAD_DIR, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		files = append(files, filename)
		return nil
	})
	log.Println("listHtml files len:", len(files))

	t.Execute(w, files)
}

//构建一个静态文件监控目录
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		log.Println("static: ", file)
		//判断本地文件是否存在
		if _, err := os.Stat(file); nil != err {
			http.NotFound(w, r)
			log.Println(file, "not found")
			return
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	//	staticDirHandler(mux, "/test/", UPLOAD_DIR)

	mux.HandleFunc("/", safeHandler(uploadHtml))
	mux.HandleFunc("/list.html", safeHandler(listHtml))
	mux.HandleFunc("/upload", safeHandler(upload))
	mux.HandleFunc("/files/", safeHandler(files))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
