package controllers

import (
	"github.com/astaxie/beego"
)

type ForexDoc struct {
	beego.Controller
}

func (this *ForexDoc) All() {
	str := `
	1、查询单个汇率和人民币的比例
		localhost:8088/forex/one?money_code=HKD
		
	2、查询两个货币之间的比例
		localhost:8088/forex/pair?money_code1=HKD&money_code2=USD
		
	3、查询汇率列表，汇率以人民币为基础
		localhost:8088/forex/list
	`

	this.Ctx.WriteString(str)
}
