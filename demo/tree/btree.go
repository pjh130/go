package main

import (
	"container/list"
	"fmt"
)

type BSTree struct {
	root *Node
}

type Node struct {
	left  *Node
	right *Node
	value int
}

// NewBSTree 创建树
func NewBSTree() *BSTree {
	return &BSTree{}
}

func (t *BSTree) Insert(value int) {
	var parent *Node
	z := &Node{value: value}
	x := t.root
	for x != nil {
		parent = x
		if z.value < x.value {
			x = x.left
		} else {
			x = x.right
		}
	}
	if parent == nil { //该树为空
		t.root = z
	} else if z.value < parent.value {
		parent.left = z
	} else {
		parent.right = z
	}
}

func (t *BSTree) Search(x int) *Node {
	node := t.root
	for node != nil {
		if node.value == x {
			return node
		} else if x < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func (t *BSTree) Delete(x int) bool {
	var parent *Node
	node := t.root
	isFind := false
	for node != nil {
		if node.value == x {
			isFind = true
			break
		} else if x < node.value {
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
