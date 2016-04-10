package main

import (
	"archive/zip"
	"fmt"
	"os"
	"time"
)

/*
zip包不支持跨硬盘进行操作为了向下兼容，FileHeader同时拥有32位和64位的Size字段。64位字段总
是包含正确的值，对普通格式的档案未见它们的值是相同的。对zip64格式的档案文件32位字段将是
0xffffffff，必须使用64位字段。
*/

func ExampleZip() {

	fileinfo, err := os.Stat("./doc.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileheader, err := zip.FileInfoHeader(fileinfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileheader.ModTime()) //2015-09-22 15:55:02 +0000 UTC
	fileheader.SetModTime(time.Now().AddDate(1, 1, 1))
	fmt.Println(fileheader.ModTime()) //2016-12-11 06:57:48 +0000 UTC
}
