package mssql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func TestMssql() {
	db, err := sql.Open("mssql", "server=localhost;user id=sa")
	if nil != err {
		log.Println(err)
		return
	}
	db.Query("SELECT * FROM t WHERE a = ?3, b = ?2, c = ?1", "x", "y", "z")
}

//proc is the proc name
//declare is the proc declare with the return values
//in is the params in
//out is the params out
//outparas is the select parameters
func GetProcSql(proc, declare, in, out string, outparas ...string) string {

	_sql := fmt.Sprintf("%v;exec %v %v", declare, proc, in)

	var outparam string
	for _, out := range outparas {
		outparam = fmt.Sprintf("%v,%v=%v OUTPUT", outparam, out, out)
	}

	outparam = fmt.Sprintf("%v;", outparam)

	if out != "" {
		_sql = fmt.Sprintf("%v%vselect %v;", _sql, outparam, out)
	} else {
		_sql = fmt.Sprintf("%v%v", _sql, outparam)
	}

	return _sql

}
