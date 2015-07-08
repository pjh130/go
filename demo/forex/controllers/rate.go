package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pjh130/go/demo/forex/models"
	"github.com/pjh130/go/demo/forex/utils"
	//	"log"
)

type ForexBase struct {
	beego.Controller
}

type ForexPair struct {
	beego.Controller
}

func (this *ForexBase) All() {
	callback := this.Input().Get("callback")

	code := this.Input().Get("code")
	data, err := models.GetCode(code)

	var result utils.Result
	if nil == err {
		result.Code = 0
		data.Rate = data.Rate / 100
		result.Data = data
	} else {
		result.Code = -1
		result.Msg = fmt.Sprintf("%s", err)
	}

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

func (this *ForexPair) All() {
	callback := this.Input().Get("callback")
	var result utils.Result

	code1 := this.Input().Get("code1")
	data1, err1 := models.GetCode(code1)

	code2 := this.Input().Get("code2")
	data2, err2 := models.GetCode(code2)

	if nil != err1 || nil != err2 {
		result.Code = -1
		if nil != err1 {
			result.Msg = fmt.Sprintf("%s", err1)
		} else {
			result.Msg = fmt.Sprintf("%s", err2)
		}
	} else {
		result.Code = 0
		result.Data = data1.Rate / data2.Rate
	}

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
