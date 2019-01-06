package lru_cache

type list struct {
	next  *list
	prev  *list
	key   int
	value int
}

// LRUCache lru
type LRUCache struct {
	length   int
	capacity int
	head     *list
	tail     *list
}

func Constructor(capacity int) LRUCache {
	tmp := &list{
		next: nil,
		prev: nil,
	}
	return LRUCache{
		capacity: capacity,
		length:   0,
		head:     tmp,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.head.next == nil {
		return -1
	}
	node := this.head.next
	for node != nil && node.key != key {
		node = node.next
	}
	// not found
	if node == nil {
		return -1
	}

	// move node
	this.moveNode(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {

	// 是否需要添加节点
	isAddNewNode := true

	node := this.head
	for node.next != nil {
		if node.next.key == key {
			isAddNewNode = false
			break
		}
		node = node.next
	}
	// first
	if !isAddNewNode {
		// TODO: 搬移节点
		node.next.value = value
		this.moveNode(node.next)
		return
	}

	// sencond
	// cache 满了，需要删除数据
	if this.length == this.capacity {
		this.delTailNode()
	}

	insert := &list{
		next:  nil,
		prev:  nil,
		key:   key,
		value: value,
	}
	if this.head.next != nil {
		this.head.next.prev = insert
		insert.next = this.head.next
	}
	this.head.next = insert
	insert.prev = this.head
	this.length++
	if this.tail == nil {
		this.tail = insert
	}
}

func (this *LRUCache) delTailNode() {

	this.tail = this.tail.prev
	this.tail.next = nil

	if this.tail == this.head {
		this.tail = nil
	}
	this.length--
}

func (this *LRUCache) moveNode(node *list) {
	// 如果是第一个节点，不做操作
	if this.head.next == node {
		return
	}

	node.prev.next = node.next
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		this.tail = node.prev
	}

	this.head.next.prev = node
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
}
