package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pjh130/go/demo/forex/utils"
	"log"
	"time"
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Forex))
}

const (
	DEFAULT_COUNT = 20 //默认获取的条数
)

type Forex struct {
	Id      int
	Country string
	Name    string
	Code    string
	Rate    float64
	Modify  time.Time
}

func GetCode(code string) (Forex, error) {
	var v Forex
	o := orm.NewOrm()
	err := o.Using(utils.DB_CONNECT_FOREX)
	if nil != err {
		log.Println(err)
		return v, err
	}

	err = o.Raw("SELECT * FROM forex WHERE code = ?", code).QueryRow(&v)

	return v, err
}

func GetCodes(begin int, count int) ([]Forex, error) {
	var v []Forex
	o := orm.NewOrm()
	err := o.Using(utils.DB_CONNECT_FOREX)
	if nil != err {
		log.Println(err)
		return v, err
	}

	//输出参数校验
	if begin < 0 {
		begin = 0
	}

	if count <= 0 {
		count = DEFAULT_COUNT
	}

	_, err = o.Raw("SELECT * FROM forex ORDER BY modify DESC limit ?,?", begin, count).QueryRows(&v)

	return v, err
}

func InsertCode(add Forex) error {
	//输出参数校验

	o := orm.NewOrm()
	err := o.Using(utils.DB_CONNECT_FOREX)
	if nil != err {
		log.Println(err)
		return err
	}

	//判断是否存在数据
	bFind := false
	var v Forex
	err = o.Raw("SELECT * FROM forex WHERE code = ?", add.Code).QueryRow(&v)
	if err != nil {
		if err != orm.ErrNoRows {
			log.Println(err)
			return err
		}
	} else {
		bFind = true
	}

	if bFind {
		//更新原有数据
		result, err := o.Raw("UPDATE forex SET rate = ?, modify = ? WHERE code = ?", add.Rate, add.Modify, add.Code).Exec()
		if nil != err {
			if err != orm.ErrNoRows {
				log.Println(err)
				return err
			}
		} else {
			return err
			aft, err := result.RowsAffected()
			//如果是一条数据都没更新
			if aft <= 0 {
				err = errors.New("0 rows affected")
				log.Println(err)
				return err
			}
		}

	} else {
		//插入新数据
		_, err := o.Insert(&add)
		if nil != err {
			log.Println(err)
			return err
		}
	}

	return nil
}
