package config

import (
	//	"flag"
	"fmt"
	"time"

	"os"

	"github.com/Unknwon/goconfig"
	"github.com/caarlos0/env"
	"github.com/go-ini/ini"
	_ "github.com/spf13/viper"
)

func exampleIni() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	// 我们可以做一些候选值限制的操作
	fmt.Println("Server Protocol:",
		cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// 如果读取的值不在候选列表内，则会回退使用提供的默认值
	fmt.Println("Email Protocol:",
		cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 试一试自动类型转换
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 差不多了，修改某个值然后进行保存
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my.ini.local")
}

func exampleGoconfig() {
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

func exampleEnv() {
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
