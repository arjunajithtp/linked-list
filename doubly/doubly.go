package doubly

type node struct {
	data interface{}
	next *node
	prev *node
}

// List is our doubly linked list
type List struct {
	head *node
	tail *node
}

// Push will insert data into linked list
func (l *List) Push(data ...interface{}) {
	for _, key := range data {
		currentNode := &node{
			data: key,
			next: l.head,
		}
		if l.head != nil {
			l.head.prev = currentNode
		}

		l.head = currentNode

		tempNode := l.head
		for tempNode.next != nil {
			tempNode = tempNode.next
		}
		l.tail = tempNode
	}
}

// ReadLIFO will collect all data in our List to a slice in LIFO order
func (l *List) ReadLIFO() []interface{} {
	tempList := l.head
	var result []interface{}
	for tempList != nil {
		result = append(result, tempList.data)
		tempList = tempList.next
	}
	return result
}

// ReadFIFO will collect all data in our List to a slice in FIFO order
func (l *List) ReadFIFO() []interface{} {
	tempList := l.tail
	var result []interface{}
	for tempList != nil {
		result = append(result, tempList.data)
		tempList = tempList.prev
	}
	return result
}

// PopHead will pop the head element
func (l *List) PopHead() interface{} {
	poppedNode := l.head
	l.head = l.head.next
	l.head.prev = nil
	return poppedNode.data
}

// PopTail will pop the head element
func (l *List) PopTail() interface{} {
	poppedNode := l.tail
	l.tail = l.tail.prev
	l.tail.next = nil
	return poppedNode.data
}
