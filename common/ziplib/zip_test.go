package ziplib

import (
	"log"
	"os"
	"testing"
)

func TestItemsZip(t *testing.T) {
	lst, err := ItemsZip("test.zip")
	if nil != err {
		t.Error(err)
		return
	}
	log.Println("length: ", len(lst))
}

func TestUnpackZip(t *testing.T) {
	dir := "./unzip"
	err := UnpackZip("test.zip", dir)
	if nil != err {
		t.Error(err)
		return
	}

	os.RemoveAll(dir)
	if nil != err {
		t.Error(err)
		return
	}
}

func TestPackZip(t *testing.T) {
	fileName := "c:\\test.zip"
	err := PackZip(fileName, "./")
	if nil != err {
		t.Error(err)
		return
	}
	return
	os.RemoveAll(fileName)
	if nil != err {
		t.Error(err)
		return
	}
}
