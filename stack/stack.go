package main

import (
	"fmt"
	"reflect"
)

type Stack interface {
	Push(value interface{}) interface{}
	Pop() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}

type stack struct {
	list          []interface{}
	top           int
	stackDataType reflect.Type
}

func NewStack() Stack {
	return &stack{make([]interface{}, 0), -1, nil}
}

func (s *stack) Push(value interface{}) interface{} {
	if value == nil {
		panic("cannot push nil value in stack")
	}
	if s.stackDataType != nil && s.stackDataType != reflect.TypeOf(value) {
		panic("cannot push value of different types in stack")
	}
	s.stackDataType = reflect.TypeOf(value)
	s.list = append(s.list, value)
	s.top++
	return value
}

func (s *stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("cannot perform pop operation on empty stack")
	} else {
		value := s.list[s.top]
		s.list = s.list[:s.top]
		s.top--
		return value
	}
}

func (s *stack) Peek() interface{} {
	if s.IsEmpty() {
		panic("cannot perform peek operation on empty stack")
	} else {
		value := s.list[s.top]
		return value
	}
}

func (s *stack) Size() int {
	return len(s.list)
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}

func main() {

	stack := NewStack()
	fmt.Println("Push operation : ", stack.Push(10))
	fmt.Println("Push operation : ", stack.Push(20))
	fmt.Println("Push operation : ", stack.Push(30))
	fmt.Println("Pop operation : ", stack.Pop())
	fmt.Println("Pop operation : ", stack.Pop())
	fmt.Println("Push operation : ", stack.Push(40))
	fmt.Println("Peek operation : ", stack.Peek())
	fmt.Println("Check size of stack : ", stack.Size())
	fmt.Println("Check if stack is empty : ", stack.IsEmpty())
	fmt.Println("Print Stack : ", stack)
}
