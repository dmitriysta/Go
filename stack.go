package main

import "fmt"

type StackNode struct {
	value int
	next  *StackNode
}

type Stack struct {
	bottom *StackNode
	top    *StackNode
}

func (s *Stack) add(value int) {
	newNode := &StackNode{value, nil}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) delete() {
	if s.top != nil {
		s.top = s.top.next
	} else {
		fmt.Println("Стек не содержит элементов")
	}

}

func (s *Stack) searchTop() int {
	return s.top.value
}

func (s *Stack) empty() bool {
	if s.top != nil {
		return false
	}else{
		return true
	}
}


