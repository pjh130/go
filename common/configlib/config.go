package configlib

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func example1() {
	cfg, err := goconfig.LoadConfigFile("test.ini")
	if nil != err {
		fmt.Println("open config err: ", err)
		return
	}

	v1 := cfg.MustValue("xxx", "bbb")
	fmt.Println(v1)
}
