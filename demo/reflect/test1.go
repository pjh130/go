package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
	C int64
}

func Test1() {
	fmt.Println("===========Test1 begin===========")
	//传入的是指针才能赋值
	t := T{23, "skidoo", 100}
	value := reflect.ValueOf(&t)
	fmt.Println("value CanSet", value.CanSet())

	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("NumField: %d  Name: %s  Type: %s  Interface:%v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())

	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	s.Field(2).SetInt(88)
	fmt.Println("t is now", t)

	fmt.Println("")
}
