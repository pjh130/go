package main

import (
	"github.com/astaxie/beego"
	"github.com/pjh130/go/demo/forex/models"
	_ "github.com/pjh130/go/demo/forex/router"
	"github.com/pjh130/go/demo/forex/utils"
)

func main() {
	//初始化配置文件
	utils.InitConfig()

	//初始化数据库和表
	utils.InitDatabase()

	//初始化要查询的原始数据
	models.InitMoneyCode()

	//开始收集数据
	models.StartCollect()

	beego.Run()
}
