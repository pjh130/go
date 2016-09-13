package mysql

/*
使用说明必须添加
_ "github.com/pjh130/go/common/storage/database/mysql"
才能注册Provider
*/

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pjh130/go/common/storage/database"
)

type MysqlProvider struct {
	savePath string
	bOk      bool
	db       *sql.DB
}

var mysqlpder = &MysqlProvider{}

func (this *MysqlProvider) DbInit(config string) error {
	this.bOk = false
	db, err := sql.Open("mysql", config)
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

func (this *MysqlProvider) GetDb() *sql.DB {
	if this.bOk == true {
		return this.db
	}

	return nil
}

func (this *MysqlProvider) Exec(query string, args ...interface{}) (sql.Result, error) {

	return this.db.Exec(query, args...)
}

func init() {
	databaselib.Register("mysql", mysqlpder)
}
