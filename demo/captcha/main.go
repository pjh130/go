package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

const HTML_TPL string = `
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Captcha by Golang</title>
</head>

<body>
<form method="post">
    <p><img src="/captcha/{{.CaptchaId}}.png" /></p>
    <p><input name="captcha" placeholder="请输入验证码" type="text" /></p>
    <input name="captcha_id" type="hidden" value="{{.CaptchaId}}" />
    <input type="submit" />
</form>
</body>
</html>`

// Main 控制器
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	captchaId := captcha.NewLen(6) //验证码长度为6
	fmt.Println(captchaId)
	html := strings.Replace(HTML_TPL, "{{.CaptchaId}}", captchaId, -1)
	this.Ctx.WriteString(html)
}

func (this *MainController) Post() {
	id, value := this.GetString("captcha_id"), this.GetString("captcha")
	b := captcha.VerifyString(id, value) //验证码校验
	this.Ctx.WriteString(strconv.FormatBool(b))
}

//设置beego和路由注册
func init() {
	beego.AutoRender = false //禁止 beego 的模板自动渲染

	beego.Router("/", &MainController{})
	beego.Handler("/captcha/*.png", captcha.Server(240, 80)) //注册验证码服务，验证码图片的宽高为240 x 80
}

func main() {
	beego.Run() //打开浏览器并访问 http://localhost:8080
}
