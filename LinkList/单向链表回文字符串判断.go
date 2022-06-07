package LinkList

func Palindrome(l *LinkList) bool {
	if l.head.next == nil || l.head.next.next == nil {
		return true
	}
	slow := l.head
	fast := l.head
	for {
		if fast == nil || fast.next == nil {
			break
		} else {
			slow = slow.next
			fast = fast.next.next
		}
	}
	// 翻转
	reversedLinkList := NewLinkList([]string{})
	reversedLinkList.head.next = slow.next
	reversedLinkList = reversedLinkList.Reverse()
	p2 := reversedLinkList.head.next
	p1 := l.head.next
	for {
		if p2 == nil {
			return true
		}
		if p2.data != p1.data {
			return false
		}
		p1 = p1.next
		p2 = p2.next
	}
}
