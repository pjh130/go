package main

import (
	"fmt"
	"reflect"
)

func Test2(i interface{}) {
	if nil == i {
		fmt.Println("i is nil")
		return
	}
	//
	v := reflect.ValueOf(i)
	if v.IsValid() {
		fmt.Println("v.IsValid true")
	} else {
		fmt.Println("v.IsValid false")
		return
	}

	if v.Kind() == reflect.Ptr {
		s := v.Elem()
		typeOfT := s.Type()
		if s.IsValid() {
			fmt.Println("s.IsValid: ", s.NumField())
			//		return
		} else {
			fmt.Println("s.IsValid false")
			return
		}
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("NumField: %d  Name: %s  Type: %s  Interface:%v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())

		}
	} else if v.Kind() == reflect.Struct {

	} else {

	}

}
