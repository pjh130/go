package utils

import (
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	"log"
	"os"
)

var (
	WAIT_TIME int
)

func InitConfig() {
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	if nil != err {
		log.Fatal(err)
		os.Exit(-1)
		return
	}

	// config app
	beego.RunMode = cfg.MustValue("app", "runmode")
	beego.HttpPort = cfg.MustInt("app", "httpport")
	beego.AppName = cfg.MustValue("app", "appname")
	//	beego.AutoRender = cfg.MustBool("app", "autorender")
	//	beego.StaticDir["/assets"] = "assets"

	WAIT_TIME = cfg.MustInt("task", "wait_time")
}
