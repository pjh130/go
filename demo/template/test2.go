package main

import (
	"html/template"
	"os"
)

type Actor struct {
	UserName string
}

func tpl_merger_structdata() {
	t := template.New("struct data demo template") //创建一个模板
	t, _ = t.Parse("hello, {{.UserName}}! \n")     //解析模板文件
	actor := Actor{UserName: "jsrush@structMap"}   // 创建一个数据对象
	t.Execute(os.Stdout, actor)                    //执行模板的merger操作，并输出到控制台
}

func tpl_merger_mapdata() {
	t := template.New("map data demo template")
	t, _ = t.Parse("Hello, {{.userName}}!\n")
	actorMap := make(map[string]string)
	actorMap["userName"] = "jsrush@actorMap"
	t.Execute(os.Stdout, actorMap)
}

func Test2() {
	tpl_merger_structdata() // 数据类型为Struct
	tpl_merger_mapdata()    // 数据类型为Map
}
