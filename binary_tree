// binary_tree 二叉树
package Algorithm

import (
	"reflect"
)

// 二叉树定义
type BinaryTree struct {
	Data   interface{}
	Lchild *BinaryTree
	Rchild *BinaryTree
}

// 构造方法
func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

// 先序遍历
func (bt *BinaryTree) PreOrder() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			res = append(res, t.Data)
			stack.Push(t)
			t = t.Lchild
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*BinaryTree)
			t = t.Rchild
		}
	}
	return res
}

// 中序遍历
func (bt *BinaryTree) InOrder() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			t = t.Lchild
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*BinaryTree)
			res = append(res, t.Data)
			t = t.Rchild
		}
	}
	return res
}

// 后续遍历
func (bt *BinaryTree) PostOrder() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	s := NewStack(reflect.TypeOf(true))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			s.Push(false)
			t = t.Lchild
		}
		for flag, _ := s.Top(); !stack.Empty() && flag.(bool); {
			s.Pop()
			v, _ := stack.Pop()
			res = append(res, v.(*BinaryTree).Data)
			flag, _ = s.Top()
		}
		if !stack.Empty() {
			s.Pop()
			s.Push(true)
			v, _ := stack.Top()
			t = v.(*BinaryTree)
			t = t.Rchild
		}
	}
	return res
}
