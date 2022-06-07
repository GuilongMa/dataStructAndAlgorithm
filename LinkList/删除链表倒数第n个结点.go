package LinkList

func (l *LinkList) DeleteNode(n int) {
	slow := l.head
	fast := l.head
	for i := 0; i < n; i++ {
		fast = fast.next
	}
	for {
		if fast.next == nil {
			break
		}
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}
