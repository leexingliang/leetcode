package graph

// Queue 循环队列
type Queue struct {
	val []int

	capacity int

	head int // 队列游标, 指向可用的下标
	tail int //
}

// NewQueue new
func NewQueue(l int) Queue {

	return Queue{
		val:      make([]int, 0, l),
		capacity: l,
	}
}

// EnQueue 入队
func (q *Queue) EnQueue(v int) {
	if q.IsFull() {
		return
	}

	q.val[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
}

// DeQueue 出队
func (q *Queue) DeQueue() int {
	if q.IsEmpty() {
		return -1
	}
	val := q.val[q.head]
	q.head = (q.head + 1) % q.capacity
	return val
}

// IsEmpty 队列判空
func (q *Queue) IsEmpty() bool {

	return q.head == q.tail
}

// IsFull 队列是否满了
func (q *Queue) IsFull() bool {
	return q.head == (q.tail+1)%q.capacity
}

// Len 队列长度
func (q *Queue) Len() int {
	return (q.head + q.capacity - q.tail) % q.capacity
}
