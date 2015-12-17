package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"time"

	"fmt"
)

func Test3() {
	fmt.Println("解析模版文件测试==============begin==============")
	t := template.New("第一个模板").Delims("[[", "]]") //创建一个模板,设置模板边界
	t, _ = t.Parse("hello,[[.UserName]]\n")       //解析模板文件
	data := map[string]interface{}{"UserName": template.HTML("<script>alert('you have been pwned')</script>")}
	t.Execute(os.Stdout, data) //执行模板的merger操作，并输出到控制台
	fmt.Println("解析模版文件测试============== end ==============\n")

	fmt.Println("注入函数测试==============begin==============")
	t2 := template.New("新的模板")                         //创建模板
	t2.Funcs(map[string]interface{}{"tihuan": tihuan}) //向模板中注入函数
	bytes, _ := ioutil.ReadFile("test3_2.html")        //读文件
	template.Must(t2.Parse(string(bytes)))             //将字符串读作模板
	t2.Execute(os.Stdout, map[string]interface{}{"UserName": "你好世界"})
	fmt.Println("注入函数测试============== end ==============\n")

	fmt.Println("读取文件模版测试==============begin==============")
	data1 := map[string]interface{}{"UserName": "data1"}
	t3, _ := template.ParseFiles("test3_1.html") //将一个文件读作模板
	t3.Execute(os.Stdout, data1)
	fmt.Println("读取文件模版测试============== end==============\n")

	fmt.Println("Glob 文件模版测试==============begin==============")
	var v TestStruct
	v.Name = "pan"
	v.Age = 20
	t4, _ := template.ParseGlob("test3_3.html") //将一个文件读作模板

	//如果回调方法中用的是(this *TestStruct)需要传入v的地址
	t4.Execute(os.Stdout, &v)
	//如果回调方法中用的是(this TestStruct)直接传值，不需要地址
	//t4.Execute(os.Stdout, v)
	fmt.Println("Glob 文件模版测试 Glob============== end ==============")
}

type TestStruct struct {
	Name string
	Age  int
}

func (this *TestStruct) SayHello(str string) string {
	return str + "-------" + time.Now().Format("2006-01-02")
}

//注入模板的函数
func tihuan(str string) string {
	return str + "-------" + time.Now().Format("2006-01-02")
}

var templ3 string = `<!DOCTYPE html>
<html>
<head>
	<title>template</title>
</head>
<body>
hello {{.UserName}}<br>
</body>
</html>
`
