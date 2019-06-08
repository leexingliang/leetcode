package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayStack(t *testing.T) {
	s := NewArrayStack(10)

	s.Push(1)
	s.Push(2)
	s.Push(10)

	data, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 10, data)
}

func Test_LinkStack(t *testing.T) {
	s := NewLinkStack(10)

	s.Push(1)
	s.Push(2)
	s.Push(10)

	data, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 10, data)
}
