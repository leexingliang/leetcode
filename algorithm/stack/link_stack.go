package stack

type node struct {
	data int

	prev *node
	next *node
}

type linkStack struct {
	head *node
	top  *node

	length int
	cap    int
}

// NewLinkStack new link stack
func NewLinkStack(cap int) Stack {
	s := &linkStack{
		head:   new(node),
		length: 0,
		cap:    cap,
	}
	s.head.next = nil
	s.head.prev = nil
	s.top = s.head

	return s
}

func (s *linkStack) Push(data int) error {
	if s.Len() == s.cap {
		return ErrStackFull
	}

	node := new(node)
	node.data = data

	s.top.next = node
	node.prev = s.top

	s.top = node
	s.length++

	return nil
}

func (s *linkStack) Pop() (int, error) {
	if s.Len() == 0 {
		return 0, ErrStackEmpty
	}
	data := s.top.data

	s.top = s.top.prev
	s.top.next = nil
	s.length--
	return data, nil
}

func (s *linkStack) Len() int {
	return s.length
}

func (s *linkStack) Empty() error {
	s.head.next = nil
	s.head.prev = nil
	s.top = s.head
	s.length = 0

	return nil
}
