package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8"
	db, err := OpenDb(dsn)
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println("open ok")

		err := DeleteTable(db, "session")
		fmt.Println("Delete result:", err)
		err = CreateTable(db, "create_table.sql")
		fmt.Println("Create result:", err)
		//		InsertData(db)
		//		ReadData(db)
		if false {
			ReadDatas(db)
		}
	}
}

func OpenDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		//		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		return nil, err
	}
	//	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		//		panic(err.Error()) // proper error handling instead of panic in your app
		return nil, err
	}

	return db, err
}

func InsertData(db *sql.DB) error {
	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		//		panic(err.Error()) // proper error handling instead of panic in your app
		return err
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		res, err := stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
		if err != nil {
			//			panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println(err)
			break
		} else {
			id, _ := res.LastInsertId()
			fmt.Println("LastInsertId: ", id)
		}
	}
	return err
}

func DeleteTable(db *sql.DB, name string) error {
	s := "DROP TABLE " + name
	_, err := db.Exec(s)

	return err
}

func CreateTable(db *sql.DB, sqlFile string) error {
	f, err := os.Open(sqlFile)
	if nil != err {
		return err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if nil != err {
		return err
	}
	fmt.Println(string(fd))
	_, err = db.Exec(string(fd))
	return err
}

func ReadData(db *sql.DB) error {
	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT uid FROM order_no WHERE id = ?")
	if err != nil {
		//		panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(err)
		return err
	}
	defer stmtOut.Close()

	var squareNum int64 // we "scan" the result in here

	err = stmtOut.QueryRow(15).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		//		panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(err)
		return err
	} else {
		fmt.Println(squareNum)
	}

	err = stmtOut.QueryRow(38).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		//		panic(err.Error()) // proper error handling instead of panic in your app
		fmt.Println(err)
		return err
	} else {
		fmt.Println(squareNum)
	}

	return err
}

func ReadDatas(db *sql.DB) error {

	rows, err := db.Query("SELECT * FROM order_no WHERE uid = ?", 4)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println(columns)
	}

	for rows.Next() {
		refs := make([]interface{}, 0, len(columns))
		//		for _, col := range columns {
		//			if len(col) > 0 {
		//			}
		//			var ref interface{}
		//			refs = append(refs, &ref)
		//		}

		//		var refs []interface{}
		for i := 0; i < len(columns); i++ {
			var ref interface{}
			refs = append(refs, &ref)
		}

		if err := rows.Scan(refs...); err != nil {
			fmt.Println("11: ", err)
			return err
		} else {
			fmt.Println("len: ", len(refs))
			for i := 0; i < len(refs); i++ {
				val := reflect.ValueOf(refs[i])
				//				fmt.Println("val.Kind: ", val.Kind())
				if val.IsValid() {
					ind := reflect.Indirect(val)
					if ind.IsValid() {
						el := ind.Elem()
						fmt.Println("ind.Kind():", ind.Kind())
						if el.IsValid() {
							//							fmt.Println("Interface: ", el.Interface())
							name := el.Type().String()
							//							fmt.Println("type name: ", name)
							switch name {
							case "int64":
								fmt.Println(columns[i]+": ", el.Int())
							case "[]uint8":
								fmt.Println(columns[i]+": ", string(el.Bytes()))
							default:
								//								fmt.Println("value: ", el.String())
							}
						}
					}
				}
			}
		}
		fmt.Println("----------------------------------")
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return err
}
