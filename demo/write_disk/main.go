package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pjh130/go/common/file"
	"github.com/pjh130/go/common/rand"
)

var formatReplacer = strings.NewReplacer([]string{
	"yyyy", "2006",
	"yy", "06",
	"mm", "01",
	"dd", "02",
	"HH", "15",
	"MM", "04",
	"SS", "05",
}...)

var dir = file.GetCurrPath() + string(os.PathSeparator) + "mydisk"
var file1 = dir + string(os.PathSeparator) + "1.txt"
var file2 = dir + string(os.PathSeparator) + "2.txt"
var file3 = dir + string(os.PathSeparator) + "log.txt"

func main() {
	//	s := "yyyy-mm-dd HH:MM:SS"
	//	fmt.Println(formatReplacer.Replace(s))
	//	s = "2006-01-02 15:04:05"
	//	fmt.Println(formatReplacer.Replace(s))

	//每次运行前先删除日志文件
	os.Remove(file3)

	startTimer()
}

func startTimer() {
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			working()
		}
	}
}

func working() {
	/*
		在当前程序下新建一个文件夹mydisk
		策略 随机名字为 1.txt 2.txt
		如果文件存在就删除其他文件,文件内写入时间戳
		如果不存在就新建，其他两个文件不管
	*/

	//文件夹不存在
	if !file.IsExist(dir) {
		err := file.MkDir(dir)
		if nil != err {
			fmt.Println(err)
			return
		} else {

		}
	}

	i := rand.RandInt(1, 3)

	str := time.Now().String()

	if i == 1 {
		str = str + " remove 2.txt and write 1.txt"
		err := os.Remove(file2)
		if nil != err {
			//fmt.Println(err)
		}

		f, err := os.Create(file1)
		if err == nil {
			f.Write([]byte(str))
			f.Close()
		} else {
			fmt.Println(err)
		}
	} else {
		str = str + " remove 1.txt and write 2.txt"
		err := os.Remove(file1)
		if nil != err {
			//fmt.Println(err)
		}

		f, err := os.Create(file2)
		if err == nil {
			f.Write([]byte(str))
			f.Close()
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println(str)
	f, err := os.OpenFile(file3, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		str = str + "\r\n"
		f.Write([]byte(str))
		f.Close()
	} else {
		fmt.Println(err)
	}

}
