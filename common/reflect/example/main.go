package main

import (
	// "log"

	"github.com/pjh130/go/common/reflect"
)

func main() {
	Test1()
}

type Orange struct {
	size   int    `kitty:"size"`
	Weight int    `kitty:"wgh"`
	From   string `kitty:"source"`
}

func (this Orange) GetWeight() int {
	return this.Weight
}

func Test1() {
	orange := Orange{1, 18, "Shanghai"}
	out := Orange{1, 18, "Shanghai"}
	reflect.ParseStruct(orange, out)
	// log.Println("out: ", out)
}
