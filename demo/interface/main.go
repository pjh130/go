package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	v := new(myInstance)
	v.Hello()
	//	v.Bye()

	vv := new(newInstance)
	vv.Bye()
	vv.ByeBye()
}
