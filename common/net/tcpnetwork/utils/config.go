package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	MaxMsgLen       uint32 `json:"max_msg_len"`
	MinMsgLen       uint32 `json:"min_msg_len"`
	Timeout         int    `json:"time_out"`
	LenMsgLen       int    `json:"len_msg_len"`
	LittleEndian    bool   `json:"little_endian"`
	MaxConnNum      int    `json:"max_conn_num"`
	PendingWriteNum int    `json:"pending_write_num"`
	ServerAddr      string `json:"server_addr"`
}

var Cfg *Config

func InitConfig() {
	//解析配置文件
	var err error
	Cfg, err = ParseConfig("./conf/config.json")
	if nil != err {
		log.Println(err)
		os.Exit(-1)
		return
	}
}

/*
解析配置文件
*/
func ParseConfig(path string) (*Config, error) {
	config := new(Config)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		return config, err
	}

	return config, nil
}
