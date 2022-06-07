package LinkList

func (l *LinkList) HasCycle() bool {
	slow := l.head
	fast := l.head
	for {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
}
