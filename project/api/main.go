package main

import (
	_ "github.com/pjh130/go/project/api/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.Listen.EnableAdmin = true
		beego.BConfig.Listen.AdminAddr = "localhost"
		beego.BConfig.Listen.AdminPort = 8088
	}

	beego.Run()
}
