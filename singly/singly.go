package singly

type node struct {
	data interface{}
	next *node
}

// List is our Linked List
type List struct {
	head *node
}

// Push will create a new node and insert it into our List
func (l *List) Push(data ...interface{}) {
	for _, key := range data {
		currentNode := &node{
			data: key,
			next: l.head,
		}
		l.head = currentNode
	}
}

// Pop will pop the recently pushed data
func (l *List) Pop() interface{} {
	poppedNode := l.head
	l.head = l.head.next
	return poppedNode.data
}

// Read will collect all data in our List to a slice
func (l *List) Read() []interface{} {
	tempList := l.head
	var result []interface{}
	for tempList != nil {
		result = append(result, tempList.data)
		tempList = tempList.next
	}
	return result
}

// Reverse will reverse our linked list
func (l *List) Reverse() {
	tempList := l.head
	var newList List
	for tempList != nil {
		currentNode := &node{
			data: tempList.data,
			next: newList.head,
		}
		newList.head = currentNode
		tempList = tempList.next
	}
	l.head = newList.head
}
