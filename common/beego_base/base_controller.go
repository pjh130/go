package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type CheckLoginFunc func([]string) bool

// Controller基类继承封装
type BaseController struct {
	beego.Controller
	IsLogin bool
}

// run before
func (this *BaseController) Prepare() {
	this.IsLogin = false

	if beego.BConfig.RunMode == beego.DEV {
		log.Println("This is " + beego.DEV)
	} else {
		log.Println("This is " + beego.PROD)
	}
}

func (this *BaseController) CheckLogin(check CheckLoginFunc, items []string) ([]string, bool) {
	this.IsLogin = false
	var values []string

	if check == nil {
		return values, this.IsLogin
	}

	for _, item := range items {
		//从参数中获取
		v := this.GetString(item)
		//如果参数中没获取到，从cookie中获取
		if len(v) <= 0 {
			v = this.Ctx.GetCookie(item)
		}

		values = append(values, v)
	}

	this.IsLogin = check(values)

	return values, this.IsLogin
}
