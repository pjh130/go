package databaselib

/*
使用说明必须添加
_ "github.com/pjh130/go/common/databaselib/sqlite"
才能注册Provider
*/

/*
Supported Types
+------------------------------+
|go        | sqlite3           |
|----------|-------------------|
|nil       | null              |
|int       | integer           |
|int64     | integer           |
|float64   | float             |
|bool      | integer           |
|[]byte    | blob              |
|string    | text              |
|time.Time | timestamp/datetime|
+------------------------------+
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pjh130/go/common/databaselib"
	"log"
)

type SQLiteProvider struct {
	savePath string
	bOk      bool
	db       *sql.DB
}

var sqlitesqlpder = &SQLiteProvider{}

func (this *SQLiteProvider) DbInit(config string) error {
	this.bOk = false
	db, err := sql.Open("sqlite3", config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	//判断数据库是否可用
	err = db.Ping()
	if nil != err {
		log.Fatal(err)
		return err
	}

	this.db = db
	this.bOk = true
	this.savePath = config
	return nil
}

func (this *SQLiteProvider) GetDb() *sql.DB {
	if this.bOk == true {
		return this.db
	}

	return nil
}

func (this *SQLiteProvider) Exec(query string, args ...interface{}) (sql.Result, error) {

	return this.db.Exec(query, args...)
}

func init() {
	fmt.Println("Register sqlite3")
	databaselib.Register("sqlite3", sqlitesqlpder)
}
