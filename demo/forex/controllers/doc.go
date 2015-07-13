package controllers

import (
	"github.com/astaxie/beego"
)

type ForexDoc struct {
	beego.Controller
}

func (this *ForexDoc) All() {
	str := "Hello world"

	this.Ctx.WriteString(str)
}
