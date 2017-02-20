package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	FILE_CONFIG = "G:\\Code\\config\\myfile\\config.ini"
)

type Conf struct {
	Files   []string
	Methods map[int]int
	Qiniu   map[string]string
}

var Config Conf

func init() {

	//	var conf Conf
	//	conf.Files = append(conf.Files, "D:\\gotest")
	//	conf.Files = append(conf.Files, "b.txt")
	//	conf.Files = append(conf.Files, "c.txt")

	//	conf.Methods = make(map[int]int)
	//	conf.Methods[1] = 111
	//	conf.Methods[2] = 222
	//	conf.Methods[3] = 333

	//	data, err := json.Marshal(conf)
	//	if nil == err {
	//		log.Println(string(data))
	//	}
}

func InitConf() {
	data, err := ioutil.ReadFile(FILE_CONFIG)
	if nil != err {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(data, &Config)
	if nil != err {
		log.Fatal(err)
		return
	}
	log.Println(Config)

	////////////////////////////////////////////////////////////////////////////
}
