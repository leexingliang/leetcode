package queue

// // type node struct {
// // 	data int

// // 	next *node
// // }

// type cycleLinkQueue struct {
// 	head *node
// 	tail *node

// 	length int
// 	cap    int
// }

// // NewCycleLinkQueue new cycle link queue
// func NewCycleLinkQueue(cap int) Queue {
// 	q := &cycleLinkQueue{
// 		cap:    cap,
// 		length: 0,
// 	}
// 	q.head = new(node)
// 	q.head.next = nil
// 	q.tail = q.head

// 	return q
// }

// func (q *cycleLinkQueue) Enqueue(data int) error {
// 	if q.Len() == q.cap {
// 		return ErrQueueFull
// 	}

// 	node := new(node)
// 	node.data = data

// 	q.head.next = node
// 	q.head = node
// 	q.length++

// 	return nil
// }

// func (q *cycleLinkQueue) Dequeue() (int, error) {
// 	if q.Len() == 0 && q.head == q.tail {
// 		return 0, ErrQueueEmpty
// 	}
// 	q.tail = q.tail.next
// 	q.length--

// 	return q.tail.data, nil
// }

// func (q *cycleLinkQueue) Len() int {
// 	return q.length
// }

// func (q *cycleLinkQueue) Empty() error {
// 	q.head = new(node)
// 	q.head.next = nil
// 	q.tail = q.head
// 	q.length = 0
// 	return nil
// }
