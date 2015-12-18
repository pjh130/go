// 打印调试
package debuglib

import (
	"log"
)

// 错误调试
var Debug bool = true

func SetDebug(bSet bool) {
	Debug = bSet
}
func Printf(format string, v ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(format, v...)
}

func Println(v ...interface{}) {
	if !Debug {
		return
	}
	log.Println(v...)
}

func Fatal(v ...interface{}) {
	if !Debug {
		return
	}
	log.Fatal(v...)
}
