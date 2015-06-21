package main

import (
	"fmt"
)

type item struct {
	key int
}

type tree struct {
	lchild, rchild *tree
	item           item
	count          int
}

func compare(x, y item) int {
	var ret int
	switch {
	case x.key > y.key:
		ret = 1
	case x.key == y.key:
		ret = 0
	case x.key < y.key:
		ret = -1
	}
	return ret
}

func create(T *tree, x item) *tree {
	if T == nil {
		T = new(tree)
		T.item = x
		T.count = 1
	} else if compare(T.item, x) == 1 {
		T.lchild = create(T.lchild, x)
	} else if compare(T.item, x) == 0 {
		T.count++
	} else {
		T.rchild = create(T.rchild, x)
	}
	return T
}

func search(T *tree, x item) *tree {
	if T == nil {
		return nil
	} else if compare(T.item, x) == 1 {
		return search(T.lchild, x)
	} else if compare(T.item, x) == -1 {
		return search(T.rchild, x)
	}
	return T
}

func main() {
	var root *tree
	t := create(root, item{89})
	root = t

	iarr := []int{1, 89, 44, 98, 54, 24, 96, 34, 74, 69, 96, 4, 0}

	for _, i := range iarr {
		//fmt.Println(i)
		create(root, item{i})
	}

	s := search(root, item{96})
	fmt.Println(s)
	s = search(root, item{4})
	fmt.Println(s)
	s = search(root, item{0})
	fmt.Println(s)
	s = search(root, item{989})
	fmt.Println(s)

	fmt.Println("")
	fmt.Println(root.item.key)
	fmt.Println(root.lchild.item.key)
	fmt.Println(root.rchild.item.key)
}
