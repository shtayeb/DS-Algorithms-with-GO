package list

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
	}
}

func (l *LinkedList) Append(item int) {
	l.Size++
	node := &Node{
		value: item,
	}

	if l.Head == nil {
		node.next = nil
		node.prev = nil
		l.Head = node
		return
	}

	node.prev = l.Head
	node.next = nil

	l.Head = node
}

func (l *LinkedList) Get(index int) int {
	current := l.Head

	for i := l.Size - 1; i >= 0; i-- {
		fmt.Printf("i = %v -- Size=%v", i, index)
		current = current.next
		if i == index {
			break
		}
	}

	if current == nil {
		return -1
	}

	return current.value
}

func (l *LinkedList) Pop() int {
	// get the last item = head
	l.Size--
	item := l.Head
	// set the head to the prev item in the list
	l.Head = l.Head.prev
	// return the poped item
	return item.value
}

func (l *LinkedList) RemoveAt(index int) {
	// direct remove
}

func (l *LinkedList) Remove(item int) {
	// search
}

func (l *LinkedList) InsertAt(index, item int) {
}

func (l *LinkedList) Clear() {
}
