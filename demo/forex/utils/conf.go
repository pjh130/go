package utils

import (
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var (
	WAIT_TIME int
)

const (
	DB_FILE          = "data/data.db"
	DB_CONNECT_FOREX = "default"
)

func InitConfig() {
	cfg, err := goconfig.LoadConfigFile("conf/conf.ini")
	if nil != err {
		log.Println(err)
		os.Exit(-1)
		return
	}

	// config app
	beego.RunMode = cfg.MustValue("app", "runmode")
	beego.HttpPort = cfg.MustInt("app", "httpport")
	beego.AppName = cfg.MustValue("app", "appname")
	//	beego.AutoRender = cfg.MustBool("app", "autorender")
	//	beego.StaticDir["/assets"] = "assets"

	WAIT_TIME = cfg.MustInt("task", "wait_time")
}

func InitDatabase() {
	// orm.DefaultTimeLoc = time.UTC
	err := orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
	if nil != err {
		log.Println("RegisterDriver err: " + err.Error())
		os.Exit(2)
		return
	}

	// set default database
	err = orm.RegisterDataBase(DB_CONNECT_FOREX, "sqlite3", DB_FILE)
	if nil != err {
		log.Println("RegisterDataBase err: " + err.Error())
		os.Exit(2)
		return
	}

	//判断如果表不存在，创建表
	if false == IsTableExist("forex") {
		if err := CreateCodeTable(); nil != err {
			log.Println("CreateCodeTable err: " + err.Error())
			os.Exit(2)
			return
		}
	}
}

func IsTableExist(tableName string) bool {
	o := orm.NewOrm()
	err := o.Using(DB_CONNECT_FOREX)
	if nil != err {
		log.Println("IsTableExist err1: ", err)
		return false
	}
	var count int

	err = o.Raw("SELECT count(*) FROM sqlite_master WHERE type='table' AND name= ?", tableName).QueryRow(&count)
	if nil != err {
		log.Println("IsTableExist err2: ", err)
		return false
	} else {
		if count > 0 {
			log.Println("Table", tableName, "exist!")
			return true
		}
	}

	return false
}

func CreateCodeTable() error {
	sqlStmt := `
	CREATE TABLE forex (id integer not null PRIMARY KEY, 
	country varchar(32),
	name varchar(32),
	money_code varchar(3) UNIQUE,
	rate float,
	modify datetime);
	`
	o := orm.NewOrm()
	err := o.Using(DB_CONNECT_FOREX)
	if nil != err {
		log.Println("CreateCodeTable err: ", err)
		return err
	}
	_, err = o.Raw(sqlStmt).Exec()

	return err
}
