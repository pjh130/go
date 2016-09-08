package main

import (
	"fmt"
	"reflect"
)

type TT struct {
	A int
	B string
}

//简单类型反射
func Example1() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type :", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

//复杂类型反射
func Example2() {
	t := TT{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	//改变结构体的值
	s.Field(0).SetInt(203333333333)
	s.Field(1).SetString("mh203333333333")

	fmt.Println(t)
}

//通过反射获得的变量的可设置属性
func Example3() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	if true == v.CanSet() {
		fmt.Println("v is setable")
		//v.Set(4.1)
	} else {
		fmt.Println("v is not setable")
	}
}

func Example4() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settablitty of p :", p.CanSet())

	v := p.Elem()
	fmt.Println("settablitty of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
