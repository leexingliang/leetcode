package valid_parentheses

import "errors"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := &arrayStack{
		data:   make([]byte, 1024, 1024),
		length: 0,
		top:    0,
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack.Push(s[i])
		case ')':
			tmp, err := stack.Pop()
			if err != nil || tmp != '(' {
				return false
			}
		case ']':
			tmp, err := stack.Pop()
			if err != nil || tmp != '[' {
				return false
			}
		case '}':
			tmp, err := stack.Pop()
			if err != nil || tmp != '{' {
				return false
			}

		default:
			return false
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

type arrayStack struct {
	data []byte

	length int // 大小
	top    int // 栈顶
}

func (s *arrayStack) Push(data byte) error {
	// if s.Len() == s.cap {
	// 	return ErrStackFull
	// }

	s.data[s.top] = data
	s.length++
	s.top++

	return nil
}

func (s *arrayStack) Pop() (byte, error) {
	if s.Len() == 0 {
		return 0, ErrStackEmpty
	}

	s.top--
	tmp := s.data[s.top]
	s.length--

	return tmp, nil
}

func (s *arrayStack) Len() int {
	return s.length
}

var ErrStackEmpty = errors.New("stack is empty")
