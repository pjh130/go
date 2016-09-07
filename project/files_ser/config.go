package main

import (
	"log"
	"os"

	filelib "github.com/pjh130/go/common/file"
)

const (
	UPLOAD_DIR = "/uploads"
	VIEWS_DIR  = "/views"
	ASSERT_DIR = "/assert"
)

var (
	RootDir = ""
)

func initConfig() {
	RootDir, _ = filelib.GetAppCurrPath()
	log.Println("root: ", RootDir)

	checkUploadDir()
}

func checkUploadDir() {
	Uploadinfo, err := os.Stat(RootDir + UPLOAD_DIR)
	if err != nil {
		os.Mkdir(RootDir+UPLOAD_DIR, os.ModePerm)
		log.Println("os.Mkdir(" + RootDir + UPLOAD_DIR)
		return
	}
	if Uploadinfo.IsDir() {
		// it's a file
	} else {
		os.Mkdir(RootDir+UPLOAD_DIR, os.ModePerm)
	}

	_, err = os.Stat(RootDir + VIEWS_DIR)
	if err != nil {
		log.Fatal("views file not found! require template file")
		// no such file or dir
		return
	}

}
