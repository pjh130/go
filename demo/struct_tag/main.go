/*
tag能用来干什么？
tag一般用于表示一个映射关系，最常见的是json解析中：
*/

package main

import (
	"fmt"
	"reflect" // 这里引入reflect模块
)

type User struct {
	Name   string `json:"name"`
	Passwd string `json:"password"`
}

/*
tag定义必须用键盘ESC键下面的那个吗？
不是，用双引号也可以：
*/
type User1 struct {
	Name   string "user name"
	Passwd string "user passsword"
}

func main() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}

	user1 := &User1{"chronos1", "pass1"}
	s1 := reflect.TypeOf(user1).Elem() //通过反射获取type定义
	for i := 0; i < s1.NumField(); i++ {
		fmt.Println(s1.Field(i).Tag) //将tag输出出来
	}
}
