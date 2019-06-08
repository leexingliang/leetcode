package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayQueue(t *testing.T) {
	q := NewArrayQueue(10)
	q.Enqueue(10)
	q.Enqueue(20)

	data, err := q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 10, data)
}

func Test_LinkQueue(t *testing.T) {
	q := NewArrayQueue(10)
	q.Enqueue(10)
	q.Enqueue(20)

	data, err := q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 10, data)
}
