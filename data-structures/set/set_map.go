package set

// MapSet implementa Set[T] sobre un map nativo de Go.
type MapSet[T comparable] struct {
	elements map[T]struct{}
}

// NewMapSet crea un conjunto vacío.
func NewMapSet[T comparable]() *MapSet[T] {
	// Completar
	return nil
}

func (s *MapSet[T]) Contains(element T) bool {
	// Completar
	return false
}

func (s *MapSet[T]) Add(element T) {
	// Completar
}

func (s *MapSet[T]) Remove(element T) {
	// Completar
}

func (s *MapSet[T]) Size() int {
	// Completar
	return 0
}

func (s *MapSet[T]) Values() []T {
	// Completar
	return nil
}

func (s *MapSet[T]) String() string {
	// Completar
	return ""
}

func (s *MapSet[T]) Union(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *MapSet[T]) Intersection(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *MapSet[T]) Difference(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *MapSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *MapSet[T]) Subset(other Set[T]) bool {
	// Completar
	return false
}

func (s *MapSet[T]) Superset(other Set[T]) bool {
	// Completar
	return false
}
