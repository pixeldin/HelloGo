package Stack

import "fmt"

//链表构造栈
type StackNode struct {
	next *StackNode
	val interface{}
}

type StackList struct {
	top *StackNode
}

func InitStack() *StackList{
	return &StackList{nil}
}

func (s *StackList) IsEmpty() bool {
	return s.top == nil
}

func (s *StackList) Push(val interface{})  {
	//n := &StackNode{next:s.top, val:val}
	//s.top = n
	s.top = &StackNode{next:s.top, val:val}
}

//返回栈顶值
func (s *StackList) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.top.val
}

//pop处栈顶节点, 并返回其值
func (s *StackList) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	pn := s.top
	s.top = s.top.next
	return pn.val
}

func (s *StackList) Clean() {
	if !s.IsEmpty() {
		//for s.top != nil {
		//	s.top = s.top.next
		//}
		s.top = nil
	}
}

//隐式保证 Stack所有方法都被StackList所实现
var _ Stack = &StackList{}

/* ============================================ */
func (s *StackList) Travel() {
	p := s.top
	fmt.Print("Stack: top|->")
	for p != nil {
		fmt.Printf("%v->", p.val)
		p = p.next
	}
	fmt.Print("|tail\n")
}

