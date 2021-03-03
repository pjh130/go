package main

import (
	"fmt"
	"strconv"

	"github.com/pjh130/go/common/sort"
)

type Boy struct {
	Name string
	Age  int
}

func (boy *Boy) GetValue() int {
	return boy.Age
}

func getTestData() []int {
	var a []int
	var b int
	for i := 10; i > 0; i-- {
		b = i
		a = append(a, b)
	}
	return a
}

func main() {
	if true {
		var a []sort.Obj
		for i := 0; i < 10; i++ {
			var b sort.Obj
			b = &Boy{"zhangsan" + strconv.Itoa(i+1), i + 1}
			a = append(a, b)
		}
		fmt.Println(a)
		fmt.Println(sort.BubbleSortT(a))
		// return
	}

	var a []int
	//冒泡
	a = getTestData()
	fmt.Println("BubbleSort")
	fmt.Println(a)
	fmt.Println(sort.BubbleSort(a))
	fmt.Println("")
	//快速
	a = getTestData()
	fmt.Println("QuickSort")
	fmt.Println(a)
	fmt.Println(sort.QuickSort(a, 0, len(a)-1))
	fmt.Println("")
	//选择
	a = getTestData()
	fmt.Println("SelectionSort")
	fmt.Println(a)
	fmt.Println(sort.SelectionSort(a))
	fmt.Println("")
	//插入排序
	a = getTestData()
	fmt.Println("InsertionSort")
	fmt.Println(a)
	fmt.Println(sort.InsertionSort(a))
	fmt.Println("")
	//希尔排序
	a = getTestData()
	fmt.Println("ShellSort")
	fmt.Println(a)
	fmt.Println(sort.ShellSort(a))
	fmt.Println("")
	//归并排序
	a = getTestData()
	fmt.Println("MergeSort")
	fmt.Println(a)
	fmt.Println(sort.MergeSort(a))
	fmt.Println("")
	//堆排序
	a = getTestData()
	fmt.Println("HeapSort")
	fmt.Println(a)
	fmt.Println(sort.HeapSort(a))
	fmt.Println("")
	//计数排序
	a = getTestData()
	fmt.Println("CountingSort")
	fmt.Println(a)
	fmt.Println(sort.CountingSort(a))
	fmt.Println("")
	//桶排序
	a = getTestData()
	fmt.Println("CountingSort")
	fmt.Println(a)
	fmt.Println(sort.CountingSort(a))
	fmt.Println("")
	//基数排序
	a = getTestData()
	sort.Swap(a, 1, 9)
	sort.Swap(a, 3, 5)
	sort.Swap(a, 6, 8)
	sort.Swap(a, 2, 4)
	fmt.Println("RadixSort")
	fmt.Println(a)
	fmt.Println(sort.RadixSort(a))
}
