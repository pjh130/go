package main

import (
	"fmt"
	"time"
)

var i int = 1

func main() {
	//	channelTest1()
	//	channelTest2()
	//	channelTest3()
	channelTest4()
}

func channelTest1() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
		//Count(chs[i])
	}

	for k, ch := range chs {
		vv := <-ch
		fmt.Println(vv)
		fmt.Println("k: ", k)
	}
	//time.Sleep(3 * time.Second)
	//go func() {

	//	for _, ch := range chs {
	//		time.Sleep(3 * time.Second)
	//		//fmt.Sprint("%d", <-ch)
	//		vv := <-ch
	//		fmt.Println(vv)
	//		//fmt.Println(k)
	//	}
	//}()

	for {
		time.Sleep(1 * time.Second)
	}
}

func Count(ch chan int) {
	ch <- i
	i++
	fmt.Println("Counting: ", i)
}

func channelTest2() {
	counterA := createCounter(2)
	counterB := createCounter(102)

	for i := 0; i < 5; i++ {
		a := <-counterA

		fmt.Printf("a->%d, B->%d\n", a, <-counterB)
	}
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	return next
}

func channelTest3() {
	count := 20
	ch := make(chan int, count)
	for i := 0; i < count; i++ {
		//ch <- i
		ch <- i
	}

	select {
	case ch <- 1:
	default:
		fmt.Println("channel is full:")
	}

	//for i := range ch {
	//	fmt.Println("Received:", i)
	//}

	var bBreak bool = false

	for {
		select {
		case a := <-ch:
			fmt.Println("Received: ", a)
		default:
			bBreak = true
		}
		if true == bBreak {
			fmt.Println("break")
			break
		}
	}
}
func channelTest4() {
	tab := []int{1, 3, 0, 5}

	ch := make(chan int)
	for _, value := range tab {
		go func(val int) {
			time.Sleep(time.Duration(val) * time.Second)
			fmt.Println(val)
			ch <- val
		}(value)
	}

	for _ = range tab {
		v := <-ch
		fmt.Println("v: ", v)
	}

}
