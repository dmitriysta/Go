//Твой код

package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

func (l *List) pushFront(value int) {
	newNode = &ListNode{value, nil}
	newNode.next = l.head
	l.head = newNode
}

func (l *List) front() int {
	return l.head.value
}

func (l *List) popFront() {
	if l.head != nil {
		l.head = l.head.next
	}
}

// Mой код:
// pushBack
func (l *List) pushBack(value int) {
	newNode := &ListNode{value, nil}
	l.tail.next = newNode
	l.tail = newNode
}

// back
func (l *List) back() int {
	return l.tail.value
}

// popBack
func (l *List) popBack() {
	if l.head != nil {
		previousBack := l.head
		for previousBack != l.tail {
			previousBack = previousBack.next
		}
		previousBack.next = nil
	} else {
		fmt.Println("Связанный список содержит 1 элемент")
	}
}
