package key_storage

// mysql session support need create table as sql:
//	CREATE TABLE `session` (
//	`session_key` char(64) NOT NULL,
//	`session_data` blob,
//	`session_expiry` int(11) unsigned NOT NULL,
//	PRIMARY KEY (`session_key`)
//	) ENGINE=MyISAM DEFAULT CHARSET=utf8;
//

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"bytes"
	"encoding/gob"
	"errors"
	"time"
)


var DriveMysqlStorage string = "mysql-storage"

type mysqlStorage struct {
	Maxlifetime int64
	SavePath    string
}

func NewMysqlStorage(config string) Storage {
	var maxlifetime int64
	var savePath    string
	
	return &mysqlStorage{
		Maxlifetime: maxlifetime,
		SavePath:  savePath,
	}
}

// connect to mysql
func (this *mysqlStorage) connectInit() *sql.DB {
	db, e := sql.Open("mysql", this.SavePath)
	if e != nil {
		return nil
	}
	return db
}

func (this *mysqlStorage) Driver() string {
	return DriveMysqlStorage
}

func (this *mysqlStorage) Get(key string, dst interface{}) error {
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	if nil == err {
		if len(sessiondata) == 0 {
			return errors.New("No data")
		} else {
			dst, err = DecodeGob(sessiondata)
			if err != nil {
				return  err
			} else {
				return  nil
			}
		}
	}
	
	return  err
}

func (this *mysqlStorage) GetRaw(key string) (interface{}, error) {
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	if nil == err {
		var kv interface{}
		if len(sessiondata) == 0 {
			return nil, errors.New("No data")
		} else {
			kv, err = DecodeGob(sessiondata)
			if err != nil {
				return nil, err
			} else {
				return kv, nil
			}
		}
	}
	
	return nil, err
}

func (this *mysqlStorage) Set(key string, v interface{}) error {
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	//不存在就插入新数据
	if err == sql.ErrNoRows {
		c.Exec("insert into session(`session_key`,`session_data`,`session_expiry`) values(?,?,?)",
			key, "", time.Now().Unix())
	}
	
	//如果是已经存在数据，则更新
	if nil == err {
		b, err := EncodeGob(v)
		if err != nil {
			return err
		}
		c.Exec("UPDATE session set `session_data`=?, `session_expiry`=? where session_key=?",
			b, time.Now().Unix(), key)
	}
	
	return  err
}

func (this *mysqlStorage) GetBool(key string) (bool, error){
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	if nil == err {
		var v interface{}
		if len(sessiondata) == 0 {
			return false, errors.New("No data")
		} else {
			v, err = DecodeGob(sessiondata)
			if err != nil {
				return false, err
			} else {
				if v2, ok := v.(bool); ok {
				return v2, nil
				}
			}
		}
	}
	
	return false, err
}

func (this *mysqlStorage) GetInt(key string) (int, error) {
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	if nil == err {
		var v interface{}
		if len(sessiondata) == 0 {
			return 0, errors.New("No data")
		} else {
			v, err = DecodeGob(sessiondata)
			if err != nil {
				return 0, err
			} else {
				if v2, ok := v.(int); ok {
				return v2, nil
				}
			}
		}
	}
	
	return 0, err
}

func (this *mysqlStorage) GetInt64(key string) (int64, error) {
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	if nil == err {
		var v interface{}
		if len(sessiondata) == 0 {
			return 0, errors.New("No data")
		} else {
			v, err = DecodeGob(sessiondata)
			if err != nil {
				return 0, err
			} else {
				if v2, ok := v.(int64); ok {
				return v2, nil
				}
			}
		}
	}
	
	return 0, err
}

func (this *mysqlStorage) GetString(key string) (string, error){
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	if nil == err {
		var v interface{}
		if len(sessiondata) == 0 {
			return "", errors.New("No data")
		} else {
			v, err = DecodeGob(sessiondata)
			if err != nil {
				return "", err
			} else {
				if v2, ok := v.(string); ok {
				return v2, nil
				}
			}
		}
	}
	
	return "", err
}

func (this *mysqlStorage) GetFloat64(key string) (float64, error) {
	c := this.connectInit()
	defer c.Close()
	
	row := c.QueryRow("select session_data from session where session_key=?", key)
	var sessiondata []byte
	err := row.Scan(&sessiondata)
	
	if nil == err {
		var v interface{}
		if len(sessiondata) == 0 {
			return 0, errors.New("No data")
		} else {
			v, err = DecodeGob(sessiondata)
			if err != nil {
				return 0, err
			} else {
				if v2, ok := v.(float64); ok {
				return v2, nil
				}
			}
		}
	}
	
	return 0, err
}

func (this *mysqlStorage) Del(key string) {
	c := this.connectInit()
	defer c.Close()
	c.Exec("DELETE FROM session where session_key=?", key)
}

func (this *mysqlStorage) SetExpire(key string, v interface{}, seconds int64) error {
	return nil
}

func EncodeGob(obj interface{}) ([]byte, error) {
	gob.Register(obj)
		
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(obj)
	if err != nil {
		return []byte(""), err
	}
	return buf.Bytes(), nil
}

func DecodeGob(encoded []byte) (interface{}, error) {
	buf := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(buf)
	var out interface{}
	err := dec.Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
