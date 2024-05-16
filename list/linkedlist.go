package list

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
		if i == index {
			break
		}
		current = current.prev
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

func (l *LinkedList) RemoveAt(index int) int {
	current := l.Head

	for i := l.Size - 1; i >= 0; i-- {
		if i == index {
			break
		}
		current = current.prev
	}

	if current == nil {
		return -1
	}

	l.Size--

	if current.prev != nil {
		current.prev.next = current.next
	}

	if current.next != nil {
		current.next.prev = current.prev
	}

	if current == l.Head {
		l.Head = current.prev
	}

	current.next = nil
	current.prev = nil

	return current.value
}

func (l *LinkedList) InsertAt(index, item int) {
	current := l.Head

	for i := l.Size - 1; i >= 0; i-- {
		if i == index {
			break
		}
		current = current.prev
	}

	// empty list or index out of range
	// append on those cases
	if current == nil {
		l.Append(item)
	}

	// do the thing
	l.Size++
	node := &Node{value: item}

	node.next = current
	node.prev = current.prev
	current.prev = node
}

func (l *LinkedList) Clear() {
	l.Size = 0
}
