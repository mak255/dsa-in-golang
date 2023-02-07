package lists

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func NewList() *LinkedList {
	return &LinkedList{}
}

func NewNode(data int) *Node {
	return &Node{data, nil}
}

func (l *LinkedList) Push(data int) {
	node := NewNode(data)
	node.next = l.head
	l.head = node
	l.length++
}

func (l *LinkedList) Append(data int) {
	if l.length == 0 {
		l.Push(data)
		return
	}
	tmpNode := l.head
	for tmpNode.next != nil {
		tmpNode = tmpNode.next
	}

	lastNode := NewNode(data)
	tmpNode.next = lastNode
	l.length++
}

func (l *LinkedList) Println() {
	if l.length == 0 {
		fmt.Println("Given linked list is empty")
		return
	}
	counter := 1
	curentNode := l.head
	for curentNode != nil {
		fmt.Printf("item No %d is: %d \n", counter, curentNode.data)
		curentNode = curentNode.next
		counter++
	}
}
