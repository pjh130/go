package main

import (
	"fmt"
)

//实例化接口
type myInstance struct {
	ins myInterface
}

//这可以检查在接口申明中的函数是否全部都实现,比如Hello和Bye
var _ myInterface = new(myInstance)

func (m *myInstance) Hello() {
	fmt.Println("hello")
}

func (m *myInstance) Bye() {
	fmt.Println("bye")
}

type newInstance struct {
	myInstance // 匿名字段，那么默认newInstance就包含了myInstance的所有字段
}

func (m *newInstance) ByeBye() {
	fmt.Println("bye bye")
}

//实际测试会重载匿名字段中的Bye()函数
//func (m *newInstance) Bye() {
//	fmt.Println("bye ...")
//}
