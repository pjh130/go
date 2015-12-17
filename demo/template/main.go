package main

import ()

func main() {
	//在标准设备中输出
	if false {
		Test1()
	}

	//输出结构体和map
	if false {
		Test2()
	}

	//html回调go中的函数
	if false {
		Test3()
	}

	//传值给html页面
	if true {
		Test4()
	}

	bExit := make(chan bool)
	<-bExit
}
