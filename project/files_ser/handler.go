package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

//封装一层安全机制处理的函数
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				//log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var MyTemplates = make(map[string]*template.Template)

//给templates加载所有views文件夹下的模版文件
func loadTmpl() {
	fileInfoArr, err := ioutil.ReadDir(RootDir + VIEWS_DIR)

	check(err)

	var temlateName, temlatePath string
	for _, fileInfo := range fileInfoArr {
		temlateName = fileInfo.Name()

		//检查后缀名
		if ext := path.Ext(temlateName); ext != ".html" {
			continue
		}
		temlatePath = RootDir + VIEWS_DIR + "/" + temlateName
		log.Println("Loadtmpl: ", temlatePath)
		t := template.Must(template.ParseFiles(temlatePath))
		MyTemplates[temlateName] = t
	}
}

//渲染模版
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := MyTemplates[tmpl+".html"].Execute(w, locals)
	check(err)
}

//一个简单的上传WEB
func uploadHtmlHandler(w http.ResponseWriter, r *http.Request) {
	renderHtml(w, "upload", nil)
}

//上传的文件资源
func filesHandler(w http.ResponseWriter, r *http.Request) {
	//文件服务器
	had := http.StripPrefix("/files", http.FileServer(http.Dir(RootDir+UPLOAD_DIR)))
	had.ServeHTTP(w, r)
}

//上传的接口
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	log.Println("Method:", method)
	if "GET" == method {
		w.Write([]byte("Please used POST method!"))
		return
	} else if "POST" == method {
		//默认的上传文件大小defaultMaxMemory是32MB（32 << 20）
		//如果有需要更大的文件上传，需要设置内存大小
		err := r.ParseMultipartForm(32 << 20)
		if nil != err {
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		//寻找表单中名为file的文件域
		file, header, err := r.FormFile("upload")
		defer file.Close()
		if nil != err {
			log.Println("FormFile: ", err)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		} else {
			//如果有中文字符，需要先在获取的内容上转码
			v, err := url.QueryUnescape(header.Filename)
			//出错
			if nil != err {
				log.Println("QueryUnescape: ", err)
				w.Write([]byte(fmt.Sprintf("%s", err)))
				return
			} else {
				log.Println(v)
			}

			//当天日期为文件夹并检查文件夹是否存在
			date := time.Now().Format("20060102")
			upDir := RootDir + UPLOAD_DIR + "/" + date
			_, err = os.Stat(upDir)
			if err != nil {
				err = os.Mkdir(upDir, os.ModePerm)
				//如果创建文件夹失败
				if err != nil {
					w.Write([]byte(fmt.Sprintf("%s", err)))
					return
				}
			}

			name := upDir + "/" + v

			//保存文件
			fileW, err := os.Create(name)
			defer fileW.Close()

			//出错
			if nil != err {
				log.Println("Create: ", err)
				w.Write([]byte(fmt.Sprintf("%s", err)))
				return
			}

			_, err = io.Copy(fileW, file)
			if nil != err {
				log.Println("Copy: ", err)
				w.Write([]byte(fmt.Sprintf("%s", err)))
				return
			} else {
				http.Redirect(w, r, "/files/"+date, http.StatusFound)
			}
		}
	} else {
		w.Write([]byte("Not support method!"))
		return
	}
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
