package collections

type Set[T comparable] struct {
	set map[T]interface{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	var set Set[T]
	set.set = map[T]interface{}{}
	set.Fill(items...)

	return &set
}

func (s *Set[T]) Add(item T) {
	s.set[item] = nil
}

func (s *Set[T]) Fill(items ...T) {
	for _, item := range items {
		s.Add(item)
	}
}
func (s *Set[T]) Remove(item T) {
	delete(s.set, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.set[item]

	return ok
}

func (s *Set[T]) Len() int {
	return len(s.set)
}
