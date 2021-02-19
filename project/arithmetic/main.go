package main

import (
	"fmt"
	//	"util/stack"
)

type tree struct {
	data int
	l    *tree
	r    *tree
}

type list struct {
	data int
	next *list
}

//阶乘
func fact(n uint32) uint32 {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//二分查找
func BSearch(a []int, element, low, height int) int {
	if low > height {
		return -1
	}

	mid := (low + height) / 2

	if a[mid] == element {
		return 0
	} else if a[mid] > element {
		return BSearch(a, element, low, mid-1)
	} else {
		return BSearch(a, element, mid+1, height)
	}
}

//斐波那契数列--递归
func f(n int) int {
	if n < 3 {
		return 1
	}
	return f(n-1) + f(n-2)
}

//斐波那契数列--非递归
func f1(n int) int {

	var s, s1, s2 int = 1, 1, 1

	for i := 3; i <= n; i++ {
		s = s1 + s2
		s2 = s1
		s1 = s
	}
	return s
}

//冒泡
func bubblesort(a []int, n int) {
	//n个数据需要n-1趟
	for i := 1; i < n; i++ {
		//每趟需要比较 元素个数-1 次
		for j := 0; j < n-i; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}

//反转链表
func reverseList(head *list) *list {
	if head == nil {
		return head
	}
	ph := reverseList(head.next)
	head.next.next = head
	head.next = nil
	return ph
}

func reverseList1(head *list) *list {
	if head == nil || head.next == nil {
		return head
	}
	node1, node2, node3 := head, head.next, head.next.next
	for node2 != nil {
		node2.next = node1
		node1.next = nil

		node1 = node2
		node2 = node3
		node3 = node3.next
	}

	return node2

}

//将二叉搜索树转换成双向链表
func treeToDoubleList(head, tail *tree, root *tree) {
	var ltail, rhead *tree
	if root == nil {
		head = nil
		tail = nil
		return
	}
	treeToDoubleList(head, ltail, root.l)
	treeToDoubleList(rhead, tail, root.r)

	if ltail != nil {
		ltail.r = root
		root.l = ltail
	} else {
		head = root
	}

	if rhead != nil {
		root.r = rhead
		rhead.l = root
	} else {
		tail = root
	}

}

//比较两个树是否相等
func compareTree(t1, t2 *tree) int {
	if t1 == nil && t2 == nil {
		return 1
	}
	if t1 == nil || t2 == nil {
		return -1
	}
	if t1.data != t2.data {
		return -1
	}

	if compareTree(t1.l, t2.l) == 1 && compareTree(t1.r, t2.r) == 1 {
		return 1
	} else {
		return -1
	}

}

//在二元树中找出和为某一值的所有路径
func getValuePath(root *tree, n int, sum int) {
	if root == nil {
		return
	}

	s.Push(root.data)

	sum += root.data

	if sum == n {
		fmt.Println(s)
		s.Pop()
		return
	}
	getValuePath(root.l, n, sum)
	getValuePath(root.r, n, sum)

	s.Pop()

}

//获取最大子数组
func getMaxSubArray(a []int, n int) {
	if n < 0 || a == nil {
		return
	}
	var sum, begin, end, max int = 0, 0, 0, -1 << 31

	for i := 0; i < n; i++ {
		sum += a[i]
		if sum < 0 {
			sum = 0
			begin = i + 1
		}
		if sum > max {
			max = sum
			end = i
		}
	}
	fmt.Println(begin, end, max)
}

//判断整数序列是不是二元查找树的后序遍历结果
func checkPostOrder(a []int, l, h int, b *bool) {
	if l >= h {
		return
	}
	parent := a[h]
	//fmt.Println(parent, l, h)
	i := l
	for ; i < h-1; i++ {
		if a[i] > parent {
			break
		}
	}

	for j := i + 1; j < h; j++ {
		if a[j] <= parent {
			*b = false
			break
		}
	}

	if *b {
		checkPostOrder(a, l, i-1, b)
		checkPostOrder(a, i, h-1, b)
	}
}

//^-反转句子里面单词的顺序
func swap(b []byte, begin, end int) {
	if begin > end {
		return
	}
	for begin <= end {
		tmp := b[begin]
		b[begin] = b[end]
		b[end] = tmp

		begin++
		end--

	}
}

func reverseStr(str *string) {

	begin := 0

	b := []byte(*str)

	for i := 0; i < len(*str); i++ {
		if b[i] == 32 {
			swap(b, begin, i-1)
			begin = i + 1
		}
	}
	swap(b, begin, len(*str)-1)

	//反转整个字符串
	swap(b, 0, len(*str)-1)

	fmt.Println(string(b))

}

//$-反转句子里面单词的顺序
func main() {
	/*
		//var b uint32 = 6
		//fmt.Println(b, "!=", fact(b))
		//a := []int{1, 2, 4, 6, 7, 8, 9, 10, 11, 12}
		//c := BSearch(a, 4, 0, 9)
		//fmt.Println(c)
		//t := []int{3, 2, 1, 5, 7, 6, 9, 0}
		//bubblesort(t, 8)
		//fmt.Println(t)
		//t := []int{1, -2, 3, 10, -4, 7, 2, -5}

		//getMaxSubArray(t1, len(t1))

				t1 := tree{4, nil, nil}
				t2 := tree{7, nil, nil}
				t3 := tree{5, &t1, &t2}
				t4 := tree{12, nil, nil}
				root := tree{10, &t3, &t4}

				getValuePath(&root, 22, 0)

			t := []int{5, 7, 6, 9, 11, 20, 1, 8}
			//t := []int{7, 4, 6, 5}
			b := true
			checkPostOrder(t, 0, 7, &b)
			fmt.Println(b)
	*/
	str := "i am a student."
	reverseStr(&str)

}
