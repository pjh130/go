package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

var testTpl string = `<!DOCTYPE html>
<html>
<head>
	<title>template</title>
</head>
<body>
receive test data: {{.vv}} <br>
</body>
</html>
`

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

func Test4() {
	go func() {
		http.HandleFunc("/", Handler)
		http.HandleFunc("/test", test)
		http.ListenAndServe(":8080", nil)
	}()
}

func test(w http.ResponseWriter, r *http.Request) {
	var v []byte = make([]byte, 1024)
	n, _ := r.Body.Read(v)
	Vv := v[:n]
	fmt.Println(string(Vv))

	t := template.New("Person template")
	t, err := t.Parse(testTpl)

	checkError(err)
	data := map[string]interface{}{"vv": string(Vv)}
	err = t.Execute(w, data)
	// w.Write([]byte("receive post test"))
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
	onlineUser.LoginTime = time.Now().String()

	t := template.New("Person template")
	t, err := t.Parse(templ)
	//	t, err := template.ParseFiles("tmpl.html")

	checkError(err)

	err = t.Execute(w, onlineUser)
	checkError(err)
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
			LoginTime {{.LoginTime}} <br>
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
