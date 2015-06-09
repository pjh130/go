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

	t := T{23, "skidoo", 100}
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

}
