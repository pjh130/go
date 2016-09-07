package routers

import (
	"github.com/astaxie/beego"
	"github.com/pjh130/go/common/beego_base/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/any", &controllers.MyConntroller{}, "*:Any")
}
