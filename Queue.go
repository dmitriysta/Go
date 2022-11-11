package main

type QueueNode struct {
	value int
	next  *QueueNode
}

type MyQueue struct {
	head *QueueNode
	tail *QueueNode
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(value int) {
	newNode := &QueueNode{value, nil}
	if q.tail == nil {
		q.tail, q.head = newNode, newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *MyQueue) Pop() int {
	temp := q.head
	if temp != nil {
		q.head = q.head.next
		if q.head == nil {
			q.tail = nil
		}
		return temp.value
	} else {
		return 0
	}
}

func (q *MyQueue) Peek() int {
	if q.head != nil {
		return q.head.value
	}
	return 0
}

func (q *MyQueue) Empty() bool {
	if q.head != nil {
		return false
	} else {
		return true
	}
}
