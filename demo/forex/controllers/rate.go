package controllers

import (
	"github.com/astaxie/beego"
	"github.com/pjh130/go/demo/forex/utils"
	"log"
)

type ForexBase struct {
	beego.Controller
}

func (this *ForexBase) All() {
	callback := this.Input().Get("callback")
	log.Println("Method: ", this.Controller.Ctx.Request.Method)
	this.Controller.Ctx.WriteString("hello world")

	var result utils.Result
	if len(callback) > 0 {
		this.Data["jsonp"] = result
		this.ServeJsonp()
		return
	} else {
		this.Data["json"] = result
		this.ServeJson()
		return
	}

}

func (this *ForexBase) Get() {
	this.Controller.Ctx.WriteString("Get: hello world")
}

func (this *ForexBase) Post() {
	this.Controller.Ctx.WriteString("Post: hello world")
}
