package main

import (
	"fmt"
)

func main() {

	//		ScanTest()
	ScanfTest()
	//	ScanlnTest()
}

//Scan从标准输入扫描文本，将成功读取的空白分隔的值保存进成功传递给本函数的参数。换行视为空白。
//返回成功扫描的条目个数和遇到的任何错误。如果读取的条目比提供的参数少，会返回一个错误报告原因。
func ScanTest() {
	fmt.Println("Scan test start...")

	for {
		var s string
		fmt.Println("Please input something...")
		n, err := fmt.Scan(&s)
		if nil != err {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("n:", n)
		if "quit" == s {
			fmt.Println("Quit")
			break
		}
		fmt.Println("You input is:", s)
	}

	fmt.Println("Scan test end")
}

//Scanf从标准输入扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递
//给本函数的参数。返回成功扫描的条目个数和遇到的任何错误。
func ScanfTest() {
	fmt.Println("Scanf test start...")

	for {
		var s string
		var i int
		fmt.Println("Please input something...")
		n, err := fmt.Scanf("%s%d\n", &s, &i)
		if nil != err {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("n:", n)
		if "quit" == s {
			fmt.Println("Quit")
			break
		}
		fmt.Println("S:", s, " i:", i)
	}

	fmt.Println("Scanf test end")
}

//Scanln类似Scan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置。
func ScanlnTest() {
	fmt.Println("Scanln test start...")

	for {
		var s string
		var i int
		fmt.Println("Please input something...")
		n, err := fmt.Scanln(&s, &i)
		if nil != err {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("n:", n)
		if "quit" == s {
			fmt.Println("Quit")
			break
		}
		fmt.Println("s:", s, "i:", i)
	}
	fmt.Println("Scanln test end")
}
