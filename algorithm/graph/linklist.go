package graph

// LinkNode 链表节点
type LinkNode struct {
	next *LinkNode
	prev *LinkNode

	val int
}

// LinkList 链表结构
type LinkList struct {
	root *LinkNode
	tail *LinkNode
}

// NewLinkList new
func NewLinkList(v int) LinkList {
	return LinkList{root: &LinkNode{}}
}

// AddNode 添加一个节点
func (l *LinkList) AddNode(v int) {
	node := &LinkNode{val: v}

	if l.tail == nil {
		l.tail = node
		l.root.next = l.tail
		l.tail.prev = l.root
		return
	}
	l.tail.next = node
	node.prev = l.tail
	l.tail = node
}

// DelNode 删除一个节点
func (l *LinkList) DelNode(node *LinkNode) {
	prev := l.root.next

	for prev != nil && prev != node {
		prev = prev.next
	}
	if prev == nil {
		return
	}
	prev.prev.next = prev.next
	if prev.next != nil {
		prev.next.prev = prev.prev
	}
}
