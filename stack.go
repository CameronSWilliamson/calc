package main

type node struct {
	value int
	next  *node
}

type stack struct {
	top *node
}

func newStack() *stack {
	return &stack{nil}
}

func (s *stack) Push(value int) {
	s.top = &node{value: value, next: s.top}
}

func (s *stack) Pop() int {
	if s.top == nil {
		return -1
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *stack) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.value
}

func (s *stack) IsEmpty() bool {
	return s.top == nil
}