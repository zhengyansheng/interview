package main

import (
	"fmt"
)

func main() {
	// stack
	stack := NewStack(2)
	fmt.Println(stack)
	stack.Push("h")
	stack.Push("e")
	stack.Push("l")
	fmt.Println("->", stack.IsFull())
	stack.Pop()
	stack.Pop()
	
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.data, stack.size)
}

/*
	栈
		Push: Add an element to the top of a stack
		Pop: Remove an element from the top of a stack
		IsEmpty: Check if the stack is empty
		IsFull: Check if the stack is full
		Peek: Get the value of the top element without removing it
*/

type Stack struct {
	data []string // 栈的存储
	size int     // 栈的大小
}

func NewStack(size ...int) *Stack {
	if len(size) == 0 {
		return &Stack{}
	} else if len(size) == 1 {
		return &Stack{size: size[0]}
	} else {
		return &Stack{}
	}
}

func (s *Stack) Push(item string) {
	if s.IsFull() {
		fmt.Println("Stack is full. ")
		return
	}
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("Stack is empty. ")
	}
	v := s.data[0]
	s.data = s.data[1:]
	return v, nil
}

func (s *Stack) Peek() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("Stack is empty. ")
	}
	return s.data[0], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) IsFull() bool {
	if s.size == 0 {
		return false
	}
	return len(s.data) >= s.size
}

// 队列

// 优先队列
