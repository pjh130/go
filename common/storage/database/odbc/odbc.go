package databaselib

/*
使用说明必须添加
_ "github.com/pjh130/go/common/storage/database/odbc"
才能注册Provider
*/

import (
	"database/sql"
	"github.com/pjh130/go/common/storage/database"
	_ "github.com/weigj/go-odbc"
	"log"
)

type OdbcProvider struct {
	savePath string
	bOk      bool
	db       *sql.DB
}

var odbcpder = &OdbcProvider{}

func (this *OdbcProvider) DbInit(config string) error {
	this.bOk = false
	db, err := sql.Open("odbc", config)
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

func (this *OdbcProvider) GetDb() *sql.DB {
	if this.bOk == true {
		return this.db
	}

	return nil
}

func (this *OdbcProvider) Exec(query string, args ...interface{}) (sql.Result, error) {

	return this.db.Exec(query, args...)
}

func init() {
	databaselib.Register("odbc", odbcpder)
}
