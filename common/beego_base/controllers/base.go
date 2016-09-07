package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

// Controller基类继承封装
type BaseController struct {
	beego.Controller
}

// run before
func (this *BaseController) Prepare() {
	// login status
	//	user := this.GetSession("username")
	if beego.BConfig.RunMode == beego.DEV {
		log.Println("This is " + beego.DEV)
	} else {
		log.Println("This is " + beego.PROD)
	}
}
