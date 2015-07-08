package main

import (
	"github.com/astaxie/beego"
	_ "github.com/pjh130/go/demo/forex/router"
	"log"
	//	"time"
)

func main() {

	InitMoneyCode()

	err := InitDb()
	if nil != err {
		log.Println(err)
		return
	}

	err = InitCodeTable()
	if nil != err {
		log.Println(err)
		return
	}

	//开始收集数据
	//	go StartCollect()

	if true {
		code1, err1 := GetCode("HKD")
		if nil != err1 {
			log.Println(err1)
		} else {
			log.Println(code1)
		}

		code2, err2 := GetCode("JPY")
		if nil != err2 {
			log.Println(err2)
		} else {
			log.Println(code2)
		}

		log.Println(code1.Rate / code2.Rate)
	}

	//	beego.RunMode = "prod"
	beego.Run()

	//只是测试阻塞用
	//	ticker := time.NewTicker(3 * time.Second)
	//	for {
	//		select {
	//		case <-ticker.C:
	//		}
	//	}
}
