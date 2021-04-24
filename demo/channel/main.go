package main

import (
	"fmt"
	"time"
)

var i int = 1

func main() {
	// channelTest()
	// channelTest1()
	// channelTest2()
	channelTest3()
	// channelTest4()
}

func channelTest() {

	var c1 chan string = make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	fmt.Println("I am here")
	fmt.Println("c1 is", <-c1)

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
	count := 200
	ch := make(chan int, count)
	go func() {
		for i := 0; i < count; i++ {
			ch <- i + 1
		}
		//		close(ch)
	}()

	go func() {
		for i := range ch {
			fmt.Println("Received:", i)
		}
	}()

	//	tick := time.NewTicker(3*time.Second)
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println(time.Now())
		}
	}

	//	var bBreak bool = false

	//	for {
	//		select {
	//		case a := <-ch:
	//			fmt.Println("Received: ", a)
	//		default:
	//			bBreak = true
	//		}
	//		if true == bBreak {
	//			fmt.Println("break")
	//			break
	//		}
	//	}
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
