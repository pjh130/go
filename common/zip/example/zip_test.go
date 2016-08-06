package example

import (
	"log"
	"os"
	"testing"
	"github.com/pjh130/go/common/zip"
)

func TestAll(t *testing.T) {
	//打包
	fileName := "c:\\gozip.zip"
	err := zip.PackZip(fileName, "c:\\debug\\gozip_test")
//	err := zip.PackZip(fileName, "c:\\debug\\1.1")
	if nil != err {
		t.Error(err)
		return
	}

	//显示
	lst, err := zip.ItemsZip(fileName)
	if nil != err {
		t.Error(err)
		return
	}
	for i, v := range lst {
		log.Println(i, v)
	}

	//解压
	dir := "d:\\gotest\\unzip"
	err = zip.UnpackZip(fileName, dir)
	if nil != err {
		t.Error(err)
		return
	}

		
	//删除测试文件
	os.RemoveAll(fileName)
	os.RemoveAll(dir)

}
