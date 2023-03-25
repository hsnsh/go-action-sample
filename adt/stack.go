package adt

type Stack struct {
	isEmpty bool
}

func (s *Stack) Empty() bool {
	return s.isEmpty
}

func (s *Stack) Add(item string) {
	s.isEmpty = false
}

func NewStack() *Stack {
	return &Stack{true}
}
