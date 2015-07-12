package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

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
	dumx := Person{
		Name:    "zoro",
		Age:     27,
		Emails:  []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role:    "SE"}

	chxd := Person{Name: "chxd", Age: 27, Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	//	t, err := template.ParseFiles("tmpl.html")

	checkError(err)

	err = t.Execute(w, onlineUser)
	checkError(err)
}

func main() {
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
