package LinkList

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

func (node *Node) String() string {
	return fmt.Sprintf("node:(%v)", node.data)
}

type LinkList struct {
	head *Node
}

func NewLinkList(dataList []string) *LinkList {
	head := &Node{}
	l := &LinkList{head: head}
	for _, data := range dataList {
		l.BackInsert(data)
	}
	return l
}

// PreInsert  链表头插入
func (l *LinkList) PreInsert(data interface{}) {
	newNode := &Node{data: data}
	newNode.next = l.head.next
	l.head.next = newNode
}

// BackInsert 链表尾部插入
func (l *LinkList) BackInsert(data interface{}) {
	p := l.head
	for {
		if p.next == nil {
			break
		} else {
			p = p.next
		}
	}
	newNode := &Node{data: data}
	p.next = newNode
}

// Delete 删除
func (l *LinkList) Delete(data interface{}) {
	p := l.head
	for {
		if p.next.data == data {
			break
		} else {
			p = p.next
		}
	}
	p.next = p.next.next
}

// Update 修改
func (l *LinkList) Update(oldData, newData interface{}) {
	p := l.head.next
	for {
		if p.data == oldData {
			break
		} else {
			p = p.next
		}
	}
	p.data = newData
}

// Search 查询
func (l *LinkList) Search(data interface{}) bool {
	p := l.head.next
	for {
		if p.data == data {
			return true
		} else {
			p = p.next
			if p == nil {
				break
			}
		}
	}

	return false
}

func (l *LinkList) String() string {
	p := l.head.next
	var s string
	for {
		if p == nil {
			s += fmt.Sprintf("nil")
			break
		}
		s += fmt.Sprintf("%s->", p)
		p = p.next
	}
	return s
}

func (l *LinkList) IsEmpty() bool {
	return l.head.next == nil
}
