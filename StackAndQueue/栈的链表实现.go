package StackAndQueue

type Node struct {
	data interface{}
	next *Node
}

type LinkListStack struct {
	head *Node
}

func NewLinkListStack() *LinkListStack {
	return &LinkListStack{
		head: &Node{},
	}
}

func (l *LinkListStack) Push(data interface{}) {
	newNode := &Node{data: data}
	newNode.next = l.head.next
	l.head.next = newNode
}

func (l *LinkListStack) Pop() interface{} {
	data := l.head.next.data
	l.head.next = l.head.next.next
	return data
}

func (l *LinkListStack) Top() interface{} {
	return l.head.next.data
}

func (l *LinkListStack) IsEmpty() bool {
	return l.head.next == nil
}
