package tree

import (
	"bytes"
	"container/list"
	"fmt"
	"reflect"
	"strings"
)

type BSTree struct {
	root *Node
}

type Node struct {
	left  *Node
	right *Node
	value interface{}
}

type CompareFunc func(interface{}, interface{}) int

var compare CompareFunc = nil

//小于 -1
//大于 1
//等于 0
const (
	Less  = -1
	More  = 1
	Equal = 0
)

func CompareDefault(a, b interface{}) int {
	t1 := reflect.TypeOf(a)
	t2 := reflect.TypeOf(a)

	//判断如果不是相同的比较类型，那么直接报错
	if t1 != t2 {
		panic("a and b must same type")
	}

	switch a.(type) {
	case *int:
		v1 := a.(*int)
		v2 := b.(*int)
		if *v1 < *v2 {
			return Less
		} else if *v1 > *v2 {
			return More
		} else {
			return Equal
		}
	case int:
		v1 := a.(int)
		v2 := b.(int)
		if v1 < v2 {
			return Less
		} else if v1 > v2 {
			return More
		} else {
			return Equal
		}
	case *uint:
		v1 := a.(*uint)
		v2 := b.(*uint)
		if *v1 < *v2 {
			return Less
		} else if *v1 > *v2 {
			return More
		} else {
			return Equal
		}
	case uint:
		v1 := a.(uint)
		v2 := b.(uint)
		if v1 < v2 {
			return Less
		} else if v1 > v2 {
			return More
		} else {
			return Equal
		}
	case *int64:
		v1 := a.(*int64)
		v2 := b.(*int64)
		if *v1 < *v2 {
			return Less
		} else if *v1 > *v2 {
			return More
		} else {
			return Equal
		}
	case int64:
		v1 := a.(int64)
		v2 := b.(int64)
		if v1 < v2 {
			return Less
		} else if v1 > v2 {
			return More
		} else {
			return Equal
		}
	case *uint64:
		v1 := a.(*uint64)
		v2 := b.(*uint64)
		if *v1 < *v2 {
			return Less
		} else if *v1 > *v2 {
			return More
		} else {
			return Equal
		}
	case uint64:
		v1 := a.(uint64)
		v2 := b.(uint64)
		if v1 < v2 {
			return Less
		} else if v1 > v2 {
			return More
		} else {
			return Equal
		}
	case string:
		v1 := a.(string)
		v2 := b.(string)
		if strings.EqualFold(v1, v2) {
			return Equal
		} else {
			return Less
		}
	case []byte:
		v1 := a.([]byte)
		v2 := b.([]byte)
		return bytes.Compare(v1, v2)
	default:
		panic("Unsupport type")
	}

	return Less
}

func initCompare(c CompareFunc) {
	if nil == c {
		compare = CompareDefault
	} else {
		compare = c
	}
}

// NewBSTree 创建树
func NewBSTree(c CompareFunc) *BSTree {
	initCompare(c)
	return &BSTree{}
}

func (t *BSTree) Insert(value interface{}) {
	var parent *Node
	z := &Node{value: value}
	x := t.root
	for x != nil {
		parent = x
		if compare(z.value, x.value) == Less {
			x = x.left
		} else {
			x = x.right
		}
	}
	if parent == nil { //该树为空
		t.root = z
	} else if compare(z.value, parent.value) == Less {
		parent.left = z
	} else {
		parent.right = z
	}
}

func (t *BSTree) Search(x interface{}) *Node {
	node := t.root
	for node != nil {
		if node.value == x {
			return node
		} else if compare(x, node.value) == Less {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func (t *BSTree) Delete(x interface{}) bool {
	var parent *Node
	node := t.root
	isFind := false
	for node != nil {
		if node.value == x {
			isFind = true
			break
		} else if compare(x, node.value) == Less {
			parent = node
			node = node.left
		} else {
			parent = node
			node = node.right
		}
	}
	if isFind == false {
		return false
	}
	//情况一：node为叶节点
	if node.left == nil && node.right == nil {
		if parent == nil {
			t.root = nil
		} else {
			if parent.left == node {
				parent.left = nil
			} else {
				parent.right = nil
			}
		}
		return true
	}
	//情况二:左孩子边为空或右边孩子为空
	if node.left == nil || node.right == nil {
		if parent == nil {
			if node.left == nil {
				t.root = node.right
			} else {
				t.root = node.left
			}
		} else {
			if parent.left == node {
				if node.left == nil {
					parent.left = node.right
				} else {
					parent.left = node.left
				}
			} else {
				if node.left == nil {
					parent.right = node.right
				} else {
					parent.right = node.left
				}
			}
		}
		return true
	}
	//情况三：两个孩子都不为空
	re := node.left
	re_parent := node
	for re.right != nil { //找到前驱节点和前驱节点的父节点
		re_parent = re
		re = re.right
	}
	node.value = re.value
	if node == re_parent {
		node.left = re.left
	} else {
		re_parent.right = re.left
	}
	return true
}

//PrintTree1 递归结构
func (t *Node) PrintTree1() {
	if t.left != nil {
		t.left.PrintTree1()
	}
	fmt.Print(t.value, " ")
	if t.right != nil {
		t.right.PrintTree1()
	}
}

//PrintTree2 非递归结构
func (t *Node) PrintTree2() {
	stack := list.New()
	stack.PushBack(t)
	for {
		node := stack.Back()
		if node == nil {
			return
		}
		stack.Remove(node)
		v, _ := node.Value.(*Node)
		fmt.Println(v.value)
		if v.left != nil {
			stack.PushBack(v.left)
		}
		if v.right != nil {
			stack.PushBack(v.right)
		}
	}
}
