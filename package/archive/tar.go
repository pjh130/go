package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func ExampleTar() {
	fileinfo, err := os.Stat("./doc.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	h, err := tar.FileInfoHeader(fileinfo, "")
	h.Linkname = "haha"
	h.Gname = "test"
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(h.AccessTime, h.ChangeTime, h.Devmajor, h.Devminor, h.Gid, h.Gname, h.Linkname, h.ModTime, h.Mode, h.Name, h.Size, h.Typeflag, h.Uid, h.Uname, h.Xattrs)

}
