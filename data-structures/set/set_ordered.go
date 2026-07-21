package set

import "github.com/untref-ayp2/data-structures/list"

// OrderedSet implementa Set[T] sobre una List, manteniendo los elementos
// ordenados según la función less.
//
// Add/Remove/Contains son O(n). Union/Intersection/Difference aprovechan
// el merge de listas ordenadas (O(n+m)).
type OrderedSet[T comparable] struct {
	elements list.List[T]
	less     func(a, b T) bool
}

// NewOrderedSet crea un conjunto ordenado vacío.
// La función less define el criterio de orden: less(a, b) es true si a
// debe ir antes que b.
func NewOrderedSet[T comparable](less func(a, b T) bool) *OrderedSet[T] {
	// Completar
	return nil
}

func (s *OrderedSet[T]) Contains(element T) bool {
	// Completar
	return false
}

func (s *OrderedSet[T]) Add(element T) {
	// Completar
}

func (s *OrderedSet[T]) Remove(element T) {
	// Completar
}

func (s *OrderedSet[T]) Size() int {
	// Completar
	return 0
}

func (s *OrderedSet[T]) Values() []T {
	// Completar
	return nil
}

func (s *OrderedSet[T]) String() string {
	// Completar
	return ""
}

func (s *OrderedSet[T]) Union(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *OrderedSet[T]) Intersection(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *OrderedSet[T]) Difference(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *OrderedSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *OrderedSet[T]) Subset(other Set[T]) bool {
	// Completar
	return false
}

func (s *OrderedSet[T]) Superset(other Set[T]) bool {
	// Completar
	return false
}
