package main

import (
	"fmt"

	"github.com/pjh130/go/common/sort"
)

type Obj interface {
	Max() bool
}

//交换模板
func swapT(a []Obj, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func isMax(a, b Obj) bool {
	return true
}

func MyTest(a []Obj) []Obj {
	q := len(a)
	f := true
	for i := 0; i < q-1; i++ { //有多少个数字需要比较 注意不和自己比较所以减一个
		{
			for j := 0; j < q-i-1; j++ { //某个数字需要和别的多少个数字比较 排好序的不用比较
				if isMax(a[j], a[j+1]) {
					swapT(a, j, j+1)
					f = false
				}
			}
			if f == true {
				return a
			}
		}
	}
	return a
}

type Boy struct {
	Name string
	Age  int
}

func (boy *Boy) Max() bool {
	return true
}

func main() {
	if true {
		var a []Obj
		var b1 Obj
		b1 = &Boy{"zhangsan", 21}

		var b2 Obj
		b2 = &Boy{"lisi", 22}

		var b3 Obj
		b3 = &Boy{"zhangsan", 23}

		var b4 Obj
		b4 = &Boy{"wangmazi", 24}

		a = append(a, b1)
		a = append(a, b2)
		a = append(a, b3)
		a = append(a, b4)

		fmt.Println(a)
		fmt.Println(MyTest(a))

		// return
	}
	var a []int
	var b, d int
	// fmt.Scanf("%d", &b)
	for i := 0; i < 10; i++ {
		d = i + 1
		// fmt.Println(d)
		a = append(a, d)
	}
	//冒泡
	fmt.Println(sort.BubbleSort(a))
	//快速
	fmt.Println(sort.QuickSort(a, 0, b-1))
	//选择
	fmt.Println(sort.SelectionSort(a))
	//插入排序
	fmt.Println(sort.InsertionSort(a))
	//希尔排序
	fmt.Println(sort.ShellSort(a))
	//归并排序
	fmt.Println(sort.MergeSort(a))
	//堆排序
	fmt.Println(sort.HeapSort(a))
	//计数排序
	fmt.Println(sort.CountingSort(a))
	//桶排序
	fmt.Println(sort.CountingSort(a))
	//基数排序
	fmt.Println(sort.RadixSort(a))
}
