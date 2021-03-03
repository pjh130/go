package main

import (
	"log"
	"net/http"
)

func main() {
	//初始化配置
	initConfig()

	//加载模版
	loadTmpl()

	mux := http.NewServeMux()

	//自带上传工具(也是找不到的默认页面)
	mux.HandleFunc("/", safeHandler(uploadHtmlHandler))

	//上传接口
	mux.HandleFunc("/upload", safeHandler(uploadHandler))

	//静态文件查看
	mux.HandleFunc("/files/", safeHandler(filesHandler))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
