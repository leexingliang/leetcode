package graph

// Stack 栈
type Stack struct {
	data []int

	top int
	cap int
}

// NewStack new
func NewStack(c int) Stack {
	return Stack{
		data: make([]int, c),
		cap:  c,
		top:  -1,
	}
}

// Push 入栈
func (s *Stack) Push(v int) {
	if s.IsFull() {
		return
	}

	s.top++
	s.data[s.top] = v
}

// Pop 出栈
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	val := s.data[s.top]
	s.top--
	return val
}

// Top 返回栈顶元素
func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	val := s.data[s.top]
	return val
}

// IsEmpty 判空
func (s *Stack) IsEmpty() bool {

	return s.top == -1
}

// IsFull 判满
func (s *Stack) IsFull() bool {
	return s.top+1 == s.cap
}
