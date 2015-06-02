package main

import (
	"fmt"
)

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice1 []int = myArray[:5]
	var mySlice2 []int = myArray[2:5]
	var mySlice3 []int = myArray[5:]
	var mySlice4 []int = myArray[:]

	fmt.Println("数组中的元素：")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println("数组切片1中的元素：")
	for _, v := range mySlice1 {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println("数组切片2中的元素：")
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println("数组切片3中的元素：")
	for _, v := range mySlice3 {
		fmt.Print(v, " ")
	}

	fmt.Println()
	fmt.Println("数组切片4中的元素：")
	for _, v := range mySlice4 {
		fmt.Print(v, " ")
	}
}
