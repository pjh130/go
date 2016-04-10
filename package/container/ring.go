package main

import (
	"container/ring"
	"fmt"
)

/*
ring包实现了环形链表的操作。
type Ring  //Ring类型代表环形链表的一个元素，同时也代表链表本身。环形链表没有头尾；指向环形
链表任一元素的指针都可以作为整个环形链表看待。Ring零值是具有一个（Value字段为nil的）元素的链表。
*/

func RingFunc() {
	r := ring.New(10) //初始长度10
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
	r = r.Move(6)
	fmt.Println(r.Value) //6
	r1 := r.Unlink(19)   //移除19%10=9个元素
	for i := 0; i < r1.Len(); i++ {
		fmt.Println(r1.Value)
		r1 = r1.Next()
	}
	fmt.Println(r.Len())  //10-9=1
	fmt.Println(r1.Len()) //9
}
