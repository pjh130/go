package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Files struct {
	gorm.Model
	Path   string
	Hash   string
	Key    string `sql:"index"`
	Status int
}

func (Files) TableName() string {
	return "files"
}

var db *gorm.DB = nil

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./resources/files.db")
	if nil != err {
		log.Fatal(err)
		return
	}

	//Create table
	db.AutoMigrate(&Files{})

	//	defer db.Close()
}

//树形结构存储目录和文件
func GetFiles() []Files {
	return nil
}
