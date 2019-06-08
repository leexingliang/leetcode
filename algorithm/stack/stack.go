package stack

import "errors"

// Stack stack
type Stack interface {
	Push(data int) error
	Pop() (int, error)
	Len() int
	Empty() error
}

// ErrStackFull full error
var ErrStackFull = errors.New("stack is full")

// ErrStackEmpty empty error
var ErrStackEmpty = errors.New("stack is empty")

// ErrOperationFail operation fail
var ErrOperationFail = errors.New("operation fail")
