package main

import (
	"fmt"
	"reflect"
)

type Queue interface {
	Add(value interface{}) interface{}
	Remove() interface{}
	Get() interface{}
	Size() int
	IsEmpty() bool
}

type queue struct {
	list          []interface{}
	size          int
	queueDataType reflect.Type
}

func (q *queue) Add(value interface{}) interface{} {
	if value == nil {
		panic("cannot add nil value in the queue")
	}
	if !q.IsEmpty() && q.queueDataType != reflect.TypeOf(value) {
		panic("cannot add value of different types in queue")
	}
	q.queueDataType = reflect.TypeOf(value)
	q.list = append(q.list, value)
	q.size++
	return value
}

func (q *queue) Remove() interface{} {
	if q.IsEmpty() {
		panic("cannot perform remove operation on empty queue")
	}
	value := q.list[0]
	q.list = q.list[1:]
	q.size--
	return value
}

func (q *queue) Get() interface{} {
	if q.IsEmpty() {
		panic("cannot perform get operation on empty queue")
	}
	return q.list[0]
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func NewQueue() Queue {
	return &queue{
		make([]interface{}, 0), 0, nil}
}

func main() {
	queue := NewQueue()
	fmt.Println("Add operation : ", queue.Add(10))
	fmt.Println("Add operation : ", queue.Add(20))
	fmt.Println("Add operation : ", queue.Add(30))
	fmt.Println("Remove operation : ", queue.Remove())
	fmt.Println("Remove operation : ", queue.Remove())
	fmt.Println("Add operation : ", queue.Add(40))
	fmt.Println("Remove operation : ", queue.Remove())
	fmt.Println("Get operation : ", queue.Get())
	fmt.Println("Check size of queue : ", queue.Size())
	fmt.Println("Check if queue is empty : ", queue.IsEmpty())
	fmt.Println("Print queue : ", queue)
}
