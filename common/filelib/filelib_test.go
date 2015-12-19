package filelib

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestGetSubDirs(t *testing.T) {
	dirs, err := GetSubDirs("../", false)
	if nil != err {
		t.Error(err)
		return
	}
	log.Println(dirs)
}

func TestGetSubDirsAll(t *testing.T) {
	_, err := GetSubDirsAll("../", true)
	if nil != err {
		t.Error(err)
		return
	}
	//	log.Println(dirs)
}

func TestGetSubFilesAll(t *testing.T) {
	path, _ := filepath.Abs("./")
	log.Println(path)
	list, err := GetSubFilesAll(path, false)
	if nil != err {
		t.Error(err)
		return
	}

	log.Println(len(list))
	for i, v := range list {
		log.Println(i, v)
	}
}

func TestCopyDir(t *testing.T) {
	src := "./"
	dest := "./new"
	err := CopyDir(src, dest)
	if nil != err {
		t.Error(err)
		return
	}

	err = os.RemoveAll(dest)
	if nil != err {
		log.Println(err)
	}
}
