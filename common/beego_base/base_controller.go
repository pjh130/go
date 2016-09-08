package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type CheckLoginFunc func(...string) bool

// Controller基类继承封装
type BaseController struct {
	beego.Controller
}

// 如果有需要在这里预先处理
func (this *BaseController) Prepare() {

	if beego.BConfig.RunMode == beego.DEV {
		log.Println("This is " + beego.DEV)
	} else {
		log.Println("This is " + beego.PROD)
	}
}

// 检查登录的接口
// checkItems的值和
func (this *BaseController) CheckLogin(check CheckLoginFunc, checkItems []string) (map[string]string, bool) {
	IsLogin := false
	CheckValues := make(map[string]string)
	args := make([]string, 0)

	if check == nil {
		return CheckValues, IsLogin
	}

	for _, item := range checkItems {
		//从参数中获取
		v := this.GetString(item)
		//如果参数中没获取到，从cookie中获取
		if len(v) <= 0 {
			v = this.Ctx.GetCookie(item)
		}

		if len(v) > 0 {
			args = append(args, v)
			CheckValues[item] = v
		}
	}

	IsLogin = check(args...)

	return CheckValues, IsLogin
}
