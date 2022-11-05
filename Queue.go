package main

import "fmt"

type QueueNode struct {
	value int
	next  *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func (q *Queue) back() int {
	return q.tail.value
}

func (q *Queue) front() int {
	return q.head.value
}

func (q *Queue) push(value int) {
	newNode := &QueueNode{value, nil}
	if q.head == nil {
		q.tail = newNode
	}
	newNode.next = q.head
	q.head = newNode
}

func (q *Queue) pop() {
	if q.head != nil {
		q.tail = q.tail.next
	} else {
		fmt.Println("Очередь содержит 1 элемент")
	}
}

func (q *Queue) empty() bool {
	if q.head != nil {
		return false
	}else{
		return true
	}
}
