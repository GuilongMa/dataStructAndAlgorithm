package LinkList

import "sync"

const DefaultCap = 8

type LRU struct {
	m        sync.Map
	linkList *LinkList
	size     int
	cap      int
}

func (lru *LRU) Search(key interface{}) (value interface{}) {
	val, ok := lru.m.Load(key)
	if ok {
		l := val.(*Node).data.(*LinkList).head.next
		for {
			if l == nil {
				break
			}
			if l.data == key {
				// 删除元素
				tmp := lru.linkList.head
				var node *Node
				for {
					if tmp.next == val {
						node = tmp.next
						tmp.next = tmp.next.next
						break
					}
					tmp = tmp.next
				}

				// 插入头部
				node.next = lru.linkList.head.next
				lru.linkList.head.next = node
				return l.data
			}
			l = l.next
		}
	} else {
		l := NewLinkList([]string{key.(string)})
		node := &Node{data: l}
		node.next = lru.linkList.head.next
		lru.linkList.head.next = node
		lru.m.Store(key, l)
		if lru.size < lru.cap {
			lru.size += 1
		} else {
			tmp := lru.linkList.head
			// 删除尾节点
			for {
				if tmp.next.next == nil {
					tmp.next = tmp.next.next
					break
				}
				tmp = tmp.next
			}
		}
	}

	return nil
}
