package main

/*
WaitGroup的用途：它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
官方对它的说明如下：
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add
to set the number of goroutines to wait for. Then each of the goroutines runs and
calls Done when finished. At the same time, Wait can be used to block until all
goroutines have finished.

sync.WaitGroup只有3个方法，Add()，Done()，Wait()。
其中Done()是Add(-1)的别名。简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
*/

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	num = 10000000
)

func main() {
	TestFunc("testchan", TestChan)
}

func TestFunc(name string, f func()) {
	st := time.Now().UnixNano()
	f()
	fmt.Printf("task %s cost %d \r\n", name, (time.Now().UnixNano()-st)/int64(time.Millisecond))
}

func TestChan() {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(1)

	go func() {
		iCount := 0
		if true {
			//			for _ = range c {
			//			}
			for vv := range c {
				iCount++
				if iCount >= num {
					fmt.Printf("last v is : %s\r\n", vv)
					break
				}
			}
		} else {
			for {
				//判断channel是否close
				vv, ok := <-c
				if ok {
				} else {
					println("channel closed")
					break
				}

				iCount++
				if iCount >= num {
					fmt.Printf("last v is : %s\r\n", vv)
					//					break
				}
			}
		}

		fmt.Printf("range count %d\r\n", iCount)
		wg.Done()
	}()

	for i := 0; i < num; i++ {
		c <- strconv.Itoa(i)
	}

	close(c)
	wg.Wait()

}
