package ormlib

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/lunny/godbc"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/ziutek/mymysql/godrv"
)

// https://github.com/go-xorm/xorm
// http://xorm.io/docs

// _ "github.com/mattn/go-oci8" 需要安装 pkg-config

//获取数据库信息
func ExampleXorm() {

	engine, err := xorm.NewEngine("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer engine.Close()

	engine.ShowSQL = true   //则会在控制台打印出生成的SQL语句；
	engine.ShowDebug = true //则会在控制台打印调试信息；
	engine.ShowErr = true   //则会在控制台打印错误信息；
	engine.ShowWarn = true  //则会在控制台打印警告信息；

	//获取数据库信息
	engine.DBMetas()      //获取到数据库中所有的表，字段，索引的信息。
	engine.TableInfo(nil) //我们通过struct建模时希望数据库的表的结构信息,不是数据库当前的表结构信息

	//表操作
	engine.CreateTables(nil) //创建表
	engine.IsTableEmpty(nil) //判断表是否为空
	engine.IsTableExist(nil) //判断表是否存在
	engine.DropTables(nil)   //删除表

	//Dump数据库结构和数据
	//1. engine.DumpAll(w io.Writer)
	//2. engine.DumpAllFile(fpath string)

	//Import 执行数据库SQL脚本
	//1. engine.Import(r io.Reader)
	//2. engine.ImportFile(fpath string)

	//插入数据
	engine.Insert(nil)
	engine.InsertOne(nil)

	//查询条件方法
	//1. Id(interface{}) 传入一个主键字段的值，作为查询条件
	//2. Where(string, …interface{}) 和SQL中Where语句中的条件基本相同，作为条件
	//3. And(string, …interface{}) 和Where函数中的条件基本相同，作为条件
	//4. Or(string, …interface{}) 和Where函数中的条件基本相同，作为条件
	//5. Sql(string, …interface{}) 执行指定的Sql语句，并把结果映射到结构体
	//6. Asc(…string) 指定字段名正序排序
	//7. Desc(…string) 指定字段名逆序排序
	//8. OrderBy(string) 按照指定的顺序进行排序
	//9. In(string, …interface{}) 某字段在一些值中，这里需要注意必须是[]interface{}才可以展开，由于Go语言的限制，[]int64等不可以直接展开，而是通过传递一个slice。
	//10. Cols(…string) 只查询或更新某些指定的字段，默认是查询所有映射的字段或者根据Update的第一个参数来判断更新的字段
	//11. AllCols() 查询或更新所有字段，一般与Update配合使用，因为默认Update只更新非0，非”“，非bool的字段
	//12. MustCols(…string) 某些字段必须更新，一般与Update配合使用
	//13. Omit(…string) 和cols相反，此函数指定排除某些指定的字段。
	//14. Distinct(…string) 按照参数中指定的字段归类结果
	//15. Table(nameOrStructPtr interface{}) 传入表名称或者结构体指针，如果传入的是结构体指针，则按照IMapper的规则提取出表名
	//16. Limit(int, …int) 限制获取的数目，第一个参数为条数，第二个参数表示开始位置，如果不传则为0
	//17. Top(int) 相当于Limit(int, 0)
	//18. Join(string,string,string) 第一个参数为连接类型，当前支持INNER, LEFT OUTER, CROSS中的一个值，第二个参数为表名，第三个参数为连接条件
	//19. GroupBy(string) Groupby的参数字符串
	//20. Having(string) Having的参数字符串

	//Iterate方法
	//Iterate方法提供逐条执行查询到的记录的方法，他所能使用的条件和Find方法完全相同
	//	err := engine.Where("age > ? or name=?)", 30, "xlw").Iterate(new(Userinfo), func(i int, bean interface{}) error {
	//		user := bean.(*Userinfo)
	//		//do somthing use i and user
	//	})

	//Rows方法
	//Rows方法和Iterate方法类似，提供逐条执行查询到的记录的方法，不过Rows更加灵活好用。
	//user := new(User)
	//rows, err := engine.Where("id >?", 1).Rows(user)
	//if err != nil {
	//}
	//defer rows.Close()
	//for rows.Next() {
	//    err = rows.Scan(user)
	//    //...
	//}

	//事务处理
	session := engine.NewSession()
	defer session.Close()
	// add Begin() before any action
	err = session.Begin()
	//	_, err = session.Insert(&user1)
	//_, err = session.Where("id = ?", 2).Update(&user2)
	//_, err = session.Exec("delete from userinfo where username = ?", user2.Username)
	if err != nil {
		session.Rollback()
		return
	}
	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return
	}

}
