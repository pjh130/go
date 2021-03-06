package main

import (
	"fmt"

	"github.com/pjh130/go/common/storage/database"
	_ "github.com/pjh130/go/common/storage/database/sqlite"
)

func main() {
	m, err := databaselib.NewManager("sqlite3", "./test.db")
	if nil == err {
		fmt.Println("ok")
		_, err = m.GetProvider().Exec("create table foo (id integer not null primary key, name text);")
		if nil != err {
			fmt.Println(err)
		} else {
			fmt.Println("exec ok")
		}
	} else {
		fmt.Println(err)
	}
}
