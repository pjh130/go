package main

import (
	"fmt"
	"github.com/pjh130/go/common/databaselib"
	_ "github.com/pjh130/go/common/databaselib/sqlite"
)

func CreateTable() {
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
