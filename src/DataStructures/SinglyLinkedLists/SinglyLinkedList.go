package singlylinkedlists

import "fmt"

type Node struct {
	value int
	next *Node
}

type SinglyLinkedList struct {
	size int
	head *Node
}

func NewList(val int) *SinglyLinkedList {
	node := Node{value: val, next: nil}
	l := SinglyLinkedList{size: 1, head: &node}
	return &l
}

// Add node to the end of the list
func (l *SinglyLinkedList) Append(val int) {
	n := Node{value: val, next: nil}

	if l.size == 0 {
		l.head = &n
		l.size++
		return
	}
	ptr := l.head
	for i := 0; i < l.size; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.size++
			return
		}
		ptr = ptr.next
	}
}

func (l *SinglyLinkedList) Remove(val int) bool {
	if l.size == 0 {
		return false
	}
	if l.head.value == val {
		l.head = l.head.next
		l.size--
		return true
	}

	cur := l.head
	for cur.next != nil {
		fmt.Printf("Current node value is: %v\n", cur.value)
		fmt.Printf("Cur.next is equal to: %v\n", cur.next)
		if cur.next.value == val {
			cur.next = cur.next.next
			l.size--
			return true
		}
		cur = cur.next
	}

	return false
}