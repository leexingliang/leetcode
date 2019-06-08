package queue

type node struct {
	data int

	next *node
}

type linkQueue struct {
	head *node
	tail *node

	length int
	cap    int
}

// NewLinkQueue new link queue
func NewLinkQueue(cap int) Queue {
	q := &linkQueue{
		cap:    cap,
		length: 0,
	}
	// q.head = new(node)
	// q.head.next = nil
	// q.tail = q.head

	return q
}

func (q *linkQueue) Enqueue(data int) error {
	if q.Len() == q.cap {
		return ErrQueueFull
	}

	// first data
	if q.head == nil {
		node := new(node)
		node.data = data

		q.head = node
		q.tail = node
		q.length++
		return nil
	}
	node := new(node)
	node.data = data

	q.head.next = node
	q.head = node
	q.length++

	return nil
}

func (q *linkQueue) Dequeue() (int, error) {
	if q.Len() == 0 && q.head == q.tail {
		return 0, ErrQueueEmpty
	}
	q.tail = q.tail.next
	q.length--

	return q.tail.data, nil
}

func (q *linkQueue) Len() int {
	return q.length
}

func (q *linkQueue) Empty() error {
	q.head = new(node)
	q.head.next = nil
	q.tail = q.head
	q.length = 0
	return nil
}
