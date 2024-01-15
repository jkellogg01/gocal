package main

type Stack[t any] struct {
	Head   *Node[t]
	Length int
}

type Node[t any] struct {
	Val  *t
	Next *Node[t]
}

func (s *Stack[t]) Push(value t) {
	next := &Node[t]{
		Val:  &value,
		Next: s.Head,
	}

	s.Head = next
	s.Length++
}

func (s *Stack[t]) Pop() *t {
	if s.Length < 1 {
		return nil
	}

	val := s.Head.Val

	s.Head = s.Head.Next
	s.Length--

	return val
}

func (s *Stack[t]) Peek() *t {
	if s.Length < 1 {
		return nil
	}

	return s.Head.Val
}
