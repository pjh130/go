// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/pjh130/go/project/api/controllers"
)

func init() {
	if true {
		ns := beego.NewNamespace("/v1",
			beego.NSNamespace("/object",
				beego.NSInclude(
					&controllers.ObjectController{},
				),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
		)
		beego.AddNamespace(ns)
	}

	return

	//判断请求是否非法
	beego.InsertFilter("/*", beego.BeforeExec, func(ctx *context.Context) {
		if true {
			ctx.Output.Body([]byte("Check CheckHandleOk!"))
			//go CheckHandleOk(ctx)
		}
	}, false)

	//添加日志打印适配
	beego.InsertFilter("/*", beego.FinishRouter, func(ctx *context.Context) {
		ctx.Output.Body([]byte("WriteHandleLogs"))
		//		go WriteHandleLogs(ctx)
	}, false)

}
