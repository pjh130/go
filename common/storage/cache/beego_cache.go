package cache

import (
	"log"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
)

/*
文档地址: http://beego.me/docs/module/cache.md
*/

func ExampleBeego() {
	bm, err := cache.NewCache("memory", `{"interval":60}`)
	//	bm, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	//	bm, err := cache.NewCache("redis", `{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}`)
	//	bm, err := cache.NewCache("memcache", `{"conn":"127.0.0.1:11211"}`)
	if nil != err {
		log.Println(err)
		return
	}
	bm.Put("astaxie", 1, 10*time.Second)
	bm.Get("astaxie")
	bm.IsExist("astaxie")
	bm.Delete("astaxie")
}
