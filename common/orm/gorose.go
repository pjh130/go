package orm

import (
	"fmt"

	"github.com/gohouse/gorose"
	_ "github.com/mattn/go-sqlite3"
)

var err error
var engin *gorose.Engin

func init() {
	// 全局初始化数据库,并复用
	// 这里的engin需要全局保存,可以用全局变量,也可以用单例
	// 配置&gorose.Config{}是单一数据库配置
	// 如果配置读写分离集群,则使用&gorose.ConfigCluster{}
	engin, err = gorose.Open(&gorose.Config{Driver: "sqlite3", Dsn: "./db.sqlite"})
	// mysql示例, 记得导入mysql驱动 github.com/go-sql-driver/mysql
	// engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true"})
}
func DB() gorose.IOrm {
	return engin.NewOrm()
}

func TestGorose() {
	// 原生sql, 直接返回结果集
	res, err := DB().Query("select * from users where uid>? limit 2", 1)
	fmt.Println(res)
	affected_rows, err := DB().Execute("delete from users where uid=?", 1)
	fmt.Println(affected_rows, err)

	// orm链式操作,查询单条数据
	res, err = DB().Table("users").First()
	// res 类型为 map[string]interface{}
	fmt.Println(res)

	// orm链式操作,查询多条数据
	res2, _ := DB().Table("users").Get()
	// res2 类型为 []map[string]interface{}
	fmt.Println(res2)
}
