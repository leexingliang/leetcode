package queue

type arrayQueue struct {
	data []int

	head   int // 队列头
	tail   int //队列尾
	length int //队列长度
	cap    int // 容量
}

// NewArrayQueue new array queue
func NewArrayQueue(cap int) Queue {
	return &arrayQueue{
		head:   0,
		tail:   0,
		length: 0,
		cap:    cap,
	}
}

func (q *arrayQueue) Enqueue(data int) error {
	if q.Len() == q.cap {
		return ErrQueueFull
	}

	q.data = append(q.data, data)
	q.head++
	q.length++
	return nil
}

func (q *arrayQueue) Dequeue() (int, error) {
	if q.Len() == 0 && q.head == q.tail {
		return 0, ErrQueueEmpty
	}
	data := q.data[q.tail]
	q.tail++
	q.length--
	return data, nil
}

func (q *arrayQueue) Len() int {
	return q.length
}

func (q *arrayQueue) Empty() error {
	q.length = 0
	q.head = 0
	q.tail = 0
	q.data = make([]int, 0)
	return nil
}
