package config

import (
	//	"flag"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/caarlos0/env"
	_ "github.com/go-ini/ini"
	_ "github.com/spf13/viper"
	"time"
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

func example2() {
	//	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//	pflag.Parse()
}

func example3() {
	type config struct {
		Home         string        `env:"HOME"`
		Port         int           `env:"PORT" envDefault:"3000"`
		IsProduction bool          `env:"PRODUCTION"`
		Hosts        []string      `env:"HOSTS" envSeparator:":"`
		Duration     time.Duration `env:"DURATION"`
	}

	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", cfg)
}
