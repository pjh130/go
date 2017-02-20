package main

import (
	"github.com/astaxie/beego"
	"github.com/pjh130/go/project/myfile/controllers"
	"github.com/pjh130/go/project/myfile/utils"

	_ "github.com/pjh130/go/project/myfile/routers"
)

func main() {
	//初始化配置文件
	utils.InitConf()

	//监控文件目录
	controllers.WatchFiles()

	beego.Run()
}
