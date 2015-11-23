package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

/*
Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，这和Java或者C++中的this类
似，如果要访问当前对象的字段通过{{.FieldName}},但是需要注意一点：这个字段必须是导出的(字段首字母必
须是大写的),否则在渲染的时候就会报错
*/

type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

type OnlineUser struct {
	User      []*Person
	LoginTime string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	zoro := Person{
		Name:    "zoro",
		Age:     27,
		Emails:  []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role:    "SE"}

	chxd := Person{Name: "chxd", Age: 27, Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&zoro, &chxd}}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	//	t, err := template.ParseFiles("tmpl.html")

	checkError(err)

	err = t.Execute(w, onlineUser)
	checkError(err)
}

func main() {
	Test3()
	return

	Test1()

	Test2()

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

var templ string = `<html>
		<head>
		</head>
		<body>
			<form action="/test" method="POST">
					{{with .User}}
						{{range .}}
							<input type="radio" name="test" value={{.Name}}/>{{.Name}}<br/>
						{{end}}
					{{end}}
					<input type="submit" value="submit"/>
			</form>
		</body>
</html>`

/*
html file
======================================
<html>
		<head>
		</head>
		<body>
			<form action="/test" method="POST">
					{{with .User}}
						{{range .}}
							<input type="radio" name="test" value={{.Name}}/>{{.Name}}<br/>
						{{end}}
					{{end}}
					<input type="submit" value="submit"/>
			</form>
		</body>
</html>
*/
