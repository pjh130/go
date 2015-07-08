package router

import (
	"github.com/astaxie/beego"
	"github.com/pjh130/go/demo/forex/controllers"
)

/*基础路由
所有的支持的基础函数
beego.Get(router, beego.FilterFunc)
beego.Post(router, beego.FilterFunc)
beego.Put(router, beego.FilterFunc)
beego.Head(router, beego.FilterFunc)
beego.Options(router, beego.FilterFunc)
beego.Delete(router, beego.FilterFunc)
beego.Any(router, beego.FilterFunc)

type FilterFunc func(*context.Context)
*/

/*
RESTful Controller 路由
*/

func init() {
	//	beego.Router("/forex", &controllers.ForexBase{}, "*:All")
	beego.Router("/forex", &controllers.ForexBase{})
}
