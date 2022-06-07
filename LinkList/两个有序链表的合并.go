package LinkList

func (l *LinkList) Merge(l2 *LinkList) *LinkList {
	l3 := NewLinkList([]string{})
	p1 := l.head.next
	p2 := l2.head.next
	p3 := l3.head
	for {
		if p1 == nil || p2 == nil {
			break
		}
		if p1.data.(string) <= p2.data.(string) {
			p3.next = p1
			p1 = p1.next
			p3 = p3.next
		} else {
			p3.next = p2
			p2 = p2.next
			p3 = p3.next
		}
	}
	if p1 == nil {
		p3.next = p2
	} else {
		p3.next = p1
	}
	return l3
}
