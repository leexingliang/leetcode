package queue

import "errors"

// Queue queue
type Queue interface {
	Enqueue(data int) error
	Dequeue() (int, error)
	Len() int
	Empty() error
}

// ErrQueueFull full error
var ErrQueueFull = errors.New("queue is full")

// ErrQueueEmpty empty error
var ErrQueueEmpty = errors.New("queue is empty")

// ErrOperationFail operation fail
var ErrOperationFail = errors.New("operation fail")
