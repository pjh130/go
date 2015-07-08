package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const (
	DB_FILE = "data/data.db"
)

type Forex struct {
	Id      int
	Country string
	Name    string
	Code    string
	Rate    float64
	Modify  time.Time
}

var dbbase *sql.DB = nil

func InitDb() error {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return err
	}

	//判断数据库是否可用
	err = db.Ping()
	if nil != err {
		db.Close()
		return err
	}

	dbbase = db
	return nil
}

func IsTableExist(tableName string) bool {
	if dbbase == nil {
		log.Fatal("dbbase is not init")
		return false
	}
	//SELECT count(*) FROM sqlite_master WHERE type='table' AND name='tableName';
	str := "SELECT count(*) FROM sqlite_master WHERE type='table' AND name= ?"

	var count int
	if false {
		rows, err := dbbase.Query(str, tableName)
		if err != nil {
			log.Fatal(err)
			return false
		}

		defer rows.Close()
		for rows.Next() {
			rows.Scan(&count)
		}
	} else {
		err := dbbase.QueryRow(str, tableName).Scan(&count)
		if err != nil {
			log.Fatal(err)
			return false
		}
	}

	if count > 0 {
		log.Println("Table", tableName, "exist!")
		return true
	}

	return false
}

func InitCodeTable() error {
	//判断表是否存在，如果存在就返回
	if true == IsTableExist("forex") {
		return nil
	}

	sqlStmt := `
	CREATE TABLE forex (id integer not null PRIMARY KEY, 
	code varchar(3) UNIQUE,
	country varchar(32),
	name varchar(32),
	rate float,
	modify datetime);
	`
	_, err := dbbase.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	return err
}

func GetCode(code string) (Forex, error) {
	var v Forex

	err := dbbase.QueryRow("SELECT * FROM forex WHERE code = ?", code).Scan(&v.Id,
		&v.Country, &v.Name, &v.Code, &v.Rate, &v.Modify)
	if err != nil {
		log.Fatal(err)
		return v, err
	}

	return v, err
}

func InsertCode(add Forex) error {
	//判断是否存在
	rows, err := dbbase.Query("SELECT * FROM forex WHERE code = ?", add.Code)
	if err != nil {
		log.Fatal(err)
		return err
	}

	bFind := false
	if rows.Next() {
		bFind = true
	}
	rows.Close()

	if bFind {
		stmt, err := dbbase.Prepare("UPDATE forex SET rate = ?, modify = ? WHERE code = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(add.Rate, add.Modify, add.Code)
		if err != nil {
			log.Fatal(err)
			return err
		}

	} else {
		stmt, err := dbbase.Prepare("insert into forex(country, name, code ,rate, modify) values(?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(add.Country, add.Name, add.Code, add.Rate, add.Modify)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
