package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func Test4() {
	s := "this is string"
	fmt.Println("1: ", reflect.TypeOf(s))
	fmt.Println("2: ", reflect.ValueOf(s))
	fmt.Println("-------------------")

	var x float64 = 3.4
	fmt.Println("3: ", reflect.ValueOf(x))
	fmt.Println("-------------------")

	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)

	fmt.Println("4: ", typ.NumMethod())
	//	fmt.Println("44: ", typ.NumOut())
	fmt.Println("-------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	//	fmt.Println("5: ", b[0])
	fmt.Println("5: ", b)

}
