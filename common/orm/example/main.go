package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Name     sql.NullString
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if nil != err {
		log.Println(err)
		return
	}

	// Check model `User`'s table exists or not
	if false == db.HasTable(&User{}) {
		log.Println("`User`'s table not exists")
		db.AutoMigrate(&User{})
	}

	defer db.Close()
}
