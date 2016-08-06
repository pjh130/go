package example

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"github.com/pjh130/go/common/file"
)

func TestGetSubDirs(t *testing.T) {
	dirs, err := file.GetSubDirs("../", false)
	if nil != err {
		t.Error(err)
		return
	}
	log.Println(dirs)
}

func TestGetSubDirsAll(t *testing.T) {
	_, err := file.GetSubDirsAll("../", true)
	if nil != err {
		t.Error(err)
		return
	}
	//	log.Println(dirs)
}

func TestGetSubFilesAll(t *testing.T) {
	path, _ := filepath.Abs("./")
	log.Println(path)
	list, err := file.GetSubFilesAll(path, false)
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
	err := file.CopyDir(src, dest)
	if nil != err {
		t.Error(err)
		return
	}

	err = os.RemoveAll(dest)
	if nil != err {
		log.Println(err)
	}
}

func TestGetCurrPath(t *testing.T) {
	log.Println(file.GetCurrPath())
}