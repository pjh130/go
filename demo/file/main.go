package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	//	"strings"
	"strconv"

	Myfile "github.com/pjh130/go/common/file"
)

func main() {
	var Tokens []string
	Tokens = append(Tokens, "SELECT", "SELECT")
	Tokens = append(Tokens, "LIMIT", strconv.Itoa(123))
	Tokens = append(Tokens, "ASC")
	fmt.Println(Tokens)
	return

	fmt.Println("Hello world!")
	fmt.Println("os.PathSeparator:", string(os.PathSeparator))
	v := "E:\\qq_chat"
	//	files, err := filelib.WalkDirAll(v)
	//	files, err := filelib.WalkDirFiles(v, "")
	files, err := Myfile.GetSubFiles(v, "", true)
	if nil == err {

		//		for i := 0; i < len(files); i++ {
		//			fmt.Println(files[i])
		//		}
		for _, s := range files {
			fmt.Println(s)
		}
		fmt.Println("length of files is ", len(files))
	}

	ReadFileLine()

	Md5OfFile("test.txt")
}

func ReadFileLine() {
	//打开文件，并进行相关处理
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	//文件关闭
	defer f.Close()

	//将文件作为一个io.Reader对象进行buffered I/O操作
	br := bufio.NewReader(f)
	for {
		//每次读取一行
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else {
			fmt.Println(line)
		}
	}
}

func Md5OfFile(filePath string) string {
	fi, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)

	h := md5.New()

	h.Write(fd)

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}
