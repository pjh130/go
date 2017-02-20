package routers

import (
	"github.com/astaxie/beego"
	"github.com/pjh130/go/project/myfile/controllers"
)

func init() {
	// index
	beego.Router("/", &controllers.Index{})
}
