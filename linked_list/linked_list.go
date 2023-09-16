package linkedlist

import (
	"golang.org/x/exp/slog"
)

type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
}

func SinglyLinkedList() {
	list := &List{}
	for i := 0; i < 5; i++ {
		list.add(i)
	}
	printList("Adding: ", list)

	list.remove(2)
	printList("after removal:", list)

}

func printList(message string, l *List) {
	curr := l.head
	for curr != nil {
		slog.Info(message, curr.data)
		curr = curr.next
	}
	slog.Info("done!")
}

// we create, and append a new value to the node at the end of the list
func (l *List) add(value int) {
	// new node is created and value is injected
	node := &Node{data: value}

	if l.head == nil {
		l.head = node
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = node
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
	for curr.next != nil && curr.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
		return
	}

}
