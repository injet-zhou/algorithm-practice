package stack

import "fmt"

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		val int
		pre *node
	}
)

func NewStack() *Stack {
	return &Stack{top: nil,length: 0}
}

func (s *Stack) Peek() int {
	if s.length == 0 {
		return 0
	}
	return s.top.val
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Pop() int {
	if s.length <= 0 {
		return 0
	}
	n := s.top
	val := n.val
	s.top = n.pre
	s.length--
	n = nil
	return val
}

func (s *Stack) Push(val int) {
	oldNode := s.top
	newNode := &node{
		val: val,
	}
	newNode.pre = oldNode
	s.top = newNode
	s.length++
}

func (s *Stack) IsEmpty() bool {
	if s.length <= 0 {
		return true
	}
	return false
}

func InitStack(args ...int) *Stack {
	s := NewStack()
	for _, arg := range args {
		s.Push(arg)
	}
	return s
}

func PrintStack(s *Stack) {
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}
