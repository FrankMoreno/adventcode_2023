package collections

type Stack[T any] struct {
	stack []T
}

func NewStack[T any](items ...T) Stack[T] {
	var stack Stack[T]
	stack.Fill(items...)

	return stack
}

func (s *Stack[T]) Fill(items ...T) int {
	for _, item := range items {
		s.Push(item)
	}

	return s.Len()
}
func (s *Stack[T]) Len() int {
	return len(s.stack)
}

func (s *Stack[T]) Pop() T {
	temp := s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]

	return temp
}

func (s *Stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *Stack[T]) Peek() T {
	return s.stack[s.Len()-1]
}
