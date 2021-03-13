package main

import (
	"log"
	"reflect"
	"time"

	"github.com/pjh130/go/common/workerpool"
	// "github.com/sniperHW/kendynet/asyn"
)

type Student struct {
	Age  int
	Name string
}

// func (t *Student) SayHello() {
// 	log.Println("Hello world!")
// }

func myWorker(t interface{}) error {
	log.Println("I am working function:", t)
	tType := reflect.TypeOf(t)
	tValue := reflect.ValueOf(t)

	log.Println("tType: ", tType.Kind())
	log.Println("tValue: ", tValue.Kind())

	kind := tType.Kind()
	switch kind {
	case reflect.Bool:
		break
	case reflect.Slice:
		// keys := tValue.MapKeys()
		// if keys == nil {
		// 	return nil
		// }
		// vv := reflect.ValueOf(tValue)
		// num := tValue.NumField()
		// log.Println("Addr:", tValue.Pointer())

		refValue := reflect.ValueOf(tValue) // value
		refType := reflect.TypeOf(tValue)   // type

		if refValue.Kind() == reflect.Struct {
			log.Println("Kind: ", refValue.Kind())
			num := refValue.NumField()
			log.Println("fieldCount:", num)
			for i := 0; i < num; i++ {
				fieldType := refType.Field(i)   // field type
				fieldValue := refValue.Field(i) // field vlaue

				log.Println("fieldType:", fieldType.Name)
				log.Println("fieldValue:", fieldValue)
				log.Println("")
			}
		} else {
			log.Println("else Kind: ", refValue.Kind())
		}
		break
	}
	return nil
	// log.Println("numIn: ", numIn)
	// var out []reflect.Value
	numIn := tType.NumIn()
	if numIn == 0 {
		log.Println("numIn: is nil")
		// out =
	} else {
		// log.Println("numIn: ", numIn)
	}

	// log.Println("I am working function Name: ", t.(Student).Name)
	// log.Println("I am working function Age: ", t.(Student).Age)
	// log.Println("I am working function Age: ", t.(Student).SayHello())

	return nil
}

func main() {

	// test1()
	test2()

	for {
		time.Sleep(time.Second * 3)
	}

}

func test1() {
	wp := &workerpool.WorkerPool{
		WorkerFunc:      myWorker,
		MaxWorkersCount: 100,
	}

	//	var c int = 1024
	//	var c string = "hello world"
	c := Student{
		Age:  20,
		Name: "panpan",
	}

	wp.Start()

	wp.Working(c)
}

func test2() {
	c := Student{
		Age:  20,
		Name: "panpan",
	}

	if true {
		p := workerpool.NewRoutinePool(10)
		p.AddTask(myWorker, c)
	} else {
		p := workerpool.NewRoutinePool(10)
		p.AddTask(func() { log.Println("hello") })
	}

}
