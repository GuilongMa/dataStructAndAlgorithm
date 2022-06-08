package StackAndQueue

import "fmt"

type NodeQueue struct {
	data interface{}
	pre  *NodeQueue
	next *NodeQueue
}

type LinkListQueue struct {
	head *NodeQueue
	tail *NodeQueue
}

func NewLinkListQueue() *LinkListQueue {
	q := &LinkListQueue{
		head: &NodeQueue{},
		tail: &NodeQueue{},
	}
	q.head.next = q.tail
	q.tail.pre = q.head

	// 增加此代码，则为循环队列
	q.head.pre = q.tail
	q.tail.next = q.head
	return q
}

func (q *LinkListQueue) Push(data interface{}) {
	node := &NodeQueue{
		data: data,
	}
	node.pre = q.tail.pre
	node.next = q.tail
	q.tail.pre.next = node
	q.tail.pre = node
}

func (q *LinkListQueue) Get() interface{} {
	if q.IsEmpty() {
		panic("LinkListQueue is empty")
	}
	node := q.head.next
	q.head.next = q.head.next.next
	q.head.next.pre = q.head
	return node.data
}

func (q *LinkListQueue) IsEmpty() bool {
	if q.head.next == q.tail || q.tail.pre == q.head {
		return true
	}
	return false
}

func (q *LinkListQueue) String() string {
	var s string
	var p = q.head.next
	for {
		if p == q.tail {
			s += "nil"
			break
		}
		s += fmt.Sprintf("node(%s)<->", p.data)
		p = p.next
	}
	return s
}
