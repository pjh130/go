package databaselib

/*
使用说明必须添加
_ "github.com/pjh130/go/common/databaselib/postgres"
才能注册Provider
*/

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pjh130/go/common/databaselib"
	"log"
)

type PqProvider struct {
	savePath string
	bOk      bool
	db       *sql.DB
}

var pqpder = &PqProvider{}

func (this *PqProvider) DbInit(config string) error {
	this.bOk = false
	//db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	db, err := sql.Open("postgres", config)
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

func (this *PqProvider) GetDb() *sql.DB {
	if this.bOk == true {
		return this.db
	}

	return nil
}

func (this *PqProvider) Exec(query string, args ...interface{}) (sql.Result, error) {

	return this.db.Exec(query, args...)
}

func init() {
	databaselib.Register("postgres", pqpder)
}
