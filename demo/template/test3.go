package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"time"

	"fmt"
)

func Test3() {
	t := template.New("第一个模板").Delims("[[", "]]") //创建一个模板,设置模板边界
	t, _ = t.Parse("hello,[[.UserName]]\n")       //解析模板文件
	data := map[string]interface{}{"UserName": template.HTML("<script>alert('you have been pwned')</script>")}
	t.Execute(os.Stdout, data) //执行模板的merger操作，并输出到控制台

	t2 := template.New("新的模板")                         //创建模板
	t2.Funcs(map[string]interface{}{"tihuan": tihuan}) //向模板中注入函数
	bytes, _ := ioutil.ReadFile("test2.html")          //读文件
	template.Must(t2.Parse(string(bytes)))             //将字符串读作模板
	t2.Execute(os.Stdout, map[string]interface{}{"UserName": "你好世界"})
	fmt.Println("\n", t2.Name(), "\n")

	t3, _ := template.ParseFiles("test1.html") //将一个文件读作模板
	t3.Execute(os.Stdout, data)
	fmt.Println(t3.Name(), "\n") //模板名称

	t4, _ := template.ParseGlob("test1.html") //将一个文件读作模板
	t4.Execute(os.Stdout, data)
	fmt.Println(t4.Name())

}

//注入模板的函数
func tihuan(str string) string {
	return str + "-------" + time.Now().Format("2006-01-02")
}
