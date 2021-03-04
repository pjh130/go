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

func (boy *Boy) SetValue(a int) {
	boy.Age = a
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
		testSortT()
	} else {
		testSort()
	}
}

func creatListT(count int) []sort.Obj {
	var a []sort.Obj
	sw := 0
	for i := count; i > 0; i-- {
		var b sort.Obj
		b = &Boy{"zhangsan" + strconv.Itoa(i), i}
		a = append(a, b)
		if sw >= 3 {
			sort.SwapT(a, 0, len(a)-1)
			sw = 0
		}
		sw++
	}

	return a
}

func creatList(count int) []int {
	var a []int
	sw := 0
	for i := count; i > 0; i-- {
		a = append(a, i)
		if sw >= 3 {
			sort.Swap(a, 0, len(a)-1)
			sw = 0
		}
		sw++
	}

	return a
}
func testSortT() {
	var b []sort.Obj
	count := 10
	//冒泡
	fmt.Println("BubbleSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.BubbleSortT(b))
	//快速
	fmt.Println("QuickSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.QuickSortT(b, 0, len(b)-1))
	//选择
	fmt.Println("SelectionSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.SelectionSortT(b))
	//插入排序
	fmt.Println("InsertionSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.InsertionSortT(b))
	//希尔排序
	fmt.Println("ShellSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.ShellSortT(b))
	//归并排序
	fmt.Println("MergeSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.MergeSortT(b))
	//堆排序
	fmt.Println("HeapSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.HeapSortT(b))
	//计数排序
	fmt.Println("CountingSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	var c sort.Obj
	c = &Boy{"", 0}
	sort.PrintObjList(sort.CountingSortT(b, c))
	//桶排序
	fmt.Println("SortT")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.SortT(b))
	//基数排序(暂时错的，需要定位修改)
	fmt.Println("RadixSort")
	b = creatListT(count)
	sort.PrintObjList(b)
	sort.PrintObjList(sort.RadixSortT(b))
}

func testSort() {
	var a []int
	count := 10
	//冒泡
	a = creatList(count)
	fmt.Println("BubbleSort")
	fmt.Println(a)
	fmt.Println(sort.BubbleSort(a))
	fmt.Println("")
	//快速
	a = creatList(count)
	fmt.Println("QuickSort")
	fmt.Println(a)
	fmt.Println(sort.QuickSort(a, 0, len(a)-1))
	fmt.Println("")
	//选择
	a = creatList(count)
	fmt.Println("SelectionSort")
	fmt.Println(a)
	fmt.Println(sort.SelectionSort(a))
	fmt.Println("")
	//插入排序
	a = creatList(count)
	fmt.Println("InsertionSort")
	fmt.Println(a)
	fmt.Println(sort.InsertionSort(a))
	fmt.Println("")
	//希尔排序
	a = creatList(count)
	fmt.Println("ShellSort")
	fmt.Println(a)
	fmt.Println(sort.ShellSort(a))
	fmt.Println("")
	//归并排序
	a = creatList(count)
	fmt.Println("MergeSort")
	fmt.Println(a)
	fmt.Println(sort.MergeSort(a))
	fmt.Println("")
	//堆排序
	a = creatList(count)
	fmt.Println("HeapSort")
	fmt.Println(a)
	fmt.Println(sort.HeapSort(a))
	fmt.Println("")
	//计数排序
	a = creatList(count)
	fmt.Println("CountingSort")
	fmt.Println(a)
	fmt.Println(sort.CountingSort(a))
	fmt.Println("")
	//桶排序
	a = creatList(count)
	fmt.Println("Sort")
	fmt.Println(a)
	fmt.Println(sort.Sort(a))
	fmt.Println("")
	//基数排序
	a = creatList(count)
	fmt.Println("RadixSort")
	fmt.Println(a)
	fmt.Println(sort.RadixSort(a))
}
