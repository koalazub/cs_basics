package linkedlist

import (
	"fmt"

	"golang.org/x/exp/slog"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int) {
	newNode := &Node{data: value}

	// newNode is the head
	if l.head == nil {
		l.head = newNode
		return
	}

	// keep assigning current from next
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	// we're trying to set it to nil
	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func SinglyLinkedList() {
	list := &List{}
	for i := 0; i < 4; i++ {
		list.add(i)
	}

	fmt.Println("Initial")
	printList(list)

	list.remove(2)
	fmt.Println("removed")
	printList(list)

	list.remove(4)
	fmt.Println("after removing 4")
	printList(list)
}

func printList(l *List) {
	curr := l.head

	for curr != nil {
		slog.Info("curr.next: ", curr)
		curr = curr.next
	}
	slog.Info("done iterating")
}
