package ziplib

import (
	"log"
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	//打包
	fileName := "c:\\gozip.zip"
	err := PackZip(fileName, "c:\\debug\\gozip")
	if nil != err {
		t.Error(err)
		return
	}

	//显示
	lst, err := ItemsZip(fileName)
	if nil != err {
		t.Error(err)
		return
	}
	for i, v := range lst {
		log.Println(i, v)
	}

	//解压
	dir := "d:\\gotest\\unzip"
	err = UnpackZip(fileName, dir)
	if nil != err {
		t.Error(err)
		return
	}

	//	return
	//删除测试文件
	os.RemoveAll(fileName)
	os.RemoveAll(dir)

}
