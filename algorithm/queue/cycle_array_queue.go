package queue

type cycleArrayQueue struct {
	data []int

	head   int // 队列头
	tail   int //队列尾
	length int //队列长度
	cap    int // 容量
}

// NewCycleArrayQueue new cycle array queue
func NewCycleArrayQueue(cap int) Queue {
	return &cycleArrayQueue{
		data:   make([]int, cap, cap),
		head:   0,
		tail:   0,
		length: 0,
		cap:    cap,
	}
}

func (q *cycleArrayQueue) Enqueue(data int) error {
	if q.Len() == q.cap && q.head == q.tail {
		return ErrQueueFull
	}

	q.data[q.head] = data
	q.length++
	q.head = (q.head + 1) / q.cap
	return nil
}

func (q *cycleArrayQueue) Dequeue() (int, error) {
	if q.Len() == 0 && q.head == q.tail {
		return 0, ErrQueueEmpty
	}
	data := q.data[q.tail]
	q.length--
	q.tail = (q.tail + 1) / q.cap
	return data, nil
}

func (q *cycleArrayQueue) Len() int {
	return q.length
}

func (q *cycleArrayQueue) Empty() error {
	q.length = 0
	q.head = 0
	q.tail = 0
	// q.data = make([]int, 0)
	return nil
}
