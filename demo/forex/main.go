package main

import (
	"fmt"
	"time"
)

func main() {

	InitMoneyCode()

	err := InitDb()
	if nil != err {
		fmt.Println(err)
		return
	}

	err = InitCodeTable()
	if nil != err {
		fmt.Println(err)
		return
	}

	//开始收集数据
	go StartCollect()

	if false {
		code1, err1 := GetCode("HKD")
		if nil != err1 {
			fmt.Println(err1)
		} else {
			fmt.Println(code1)
		}

		code2, err2 := GetCode("JPY")
		if nil != err2 {
			fmt.Println(err2)
		} else {
			fmt.Println(code2)
		}

		fmt.Println(code1.Rate / code2.Rate)
	}

	//只是测试阻塞用
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
		}
	}
}
