package example

import (
	"github.com/pjh130/go/common/archiver"
	"log"
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	//打包
	fileName := "c:\\gozip.zip"
	err := archiver.Compress(fileName, []string{"c:\\debug\\gozip_test"})
	//	err := zip.Compress(fileName, "c:\\debug\\1.1")
	if nil != err {
		t.Error(err)
		return
	}

	//显示
	lst, err := archiver.ZipItems(fileName)
	if nil != err {
		t.Error(err)
		return
	}
	for i, v := range lst {
		log.Println(i, v)
	}

	//解压
	dir := "d:\\gotest\\unzip"
	err = archiver.Uncompress(fileName, dir)
	if nil != err {
		t.Error(err)
		return
	}

	//删除测试文件
	os.RemoveAll(fileName)
	os.RemoveAll(dir)

}
