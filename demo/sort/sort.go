package main

import (
	"fmt"
	"sort"
)

func TestSort() {
	intList := []int{5, 6, 8, 7, 1, 3, 4, 9, 2, 0}
	float64List := []float64{5.5, 6.6, 8.8, 7.7, 1.1, 3.3, 4.4, 9.9, 2.2, 0.0}
	stringList := []string{"a", "b", "s", "f", "d", "c", "r", "e", "y", "u", "i", "o"}

	sort.Ints(intList)
	sort.Float64s(float64List)
	sort.Strings(stringList)

	fmt.Println(intList)
	fmt.Println(float64List)
	fmt.Println(stringList)
}

type MyData struct {
	Age  int
	Name string
}

type MySlice []MyData

func (a MySlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a MySlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a MySlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Age < a[i].Age
}

func TestSort1() {
	my := []MyData{
		{11, "11"},
		{22, "22"},
		{33, "33"},
		{44, "44"},
	}

	fmt.Println(my)

	sort.Sort(MySlice(my))

	fmt.Println(my)
}
