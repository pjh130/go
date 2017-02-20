package controllers

import (
	"github.com/astaxie/beego"
)

type Index struct {
	beego.Controller
}

func (this *Index) Get() {
	this.Ctx.WriteString("hello world!")
}
