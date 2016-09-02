package tcp_ser

import (
	"encoding/json"
	"io/ioutil"
)

/*
配置文件为json格式
	{
	  "port": 9090,
	  "max_clients": -1
	}
*/
type Config struct {
	Port       int `json:"port"`        //服务端长连接监听端口
	MaxClients int `json:"max_clients"` //服务端长连接最大连接数
}

/*
读取配置文件
*/
func ReadConfig(path string) (*Config, error) {
	config := new(Config)
	err := config.Parse(path)
	if err != nil {
		return nil, err
	}
	return config, nil
}

/*
解析配置文件
*/
func (this *Config) Parse(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &this)
	if err != nil {
		return err
	}
	return nil
}
