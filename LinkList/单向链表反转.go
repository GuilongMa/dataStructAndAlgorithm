package LinkList

func (l *LinkList) Reverse() *LinkList {
	tmp := l.head.next
	newLinkList := NewLinkList([]string{})
	for {
		if tmp == nil {
			return newLinkList
		} else {
			node := tmp
			tmp = tmp.next
			node.next = newLinkList.head.next
			newLinkList.head.next = node
		}
	}
}
