package LinkList

func (l *LinkList) MidPoint() (slow *Node) {
	slow = l.head.next
	fast := l.head.next.next
	for {
		if fast == nil || fast.next == nil {
			return
		} else {
			slow = slow.next
			fast = fast.next.next
		}
	}
}
