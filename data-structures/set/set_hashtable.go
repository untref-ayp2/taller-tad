package set

import "github.com/untref-ayp2/data-structures/hashtable"

// HashTableSet implementa Set[T] sobre una HashTable.
// El valor asociado a cada clave es struct{} (sin uso).
type HashTableSet[T comparable] struct {
	table hashtable.HashTable[T, struct{}]
}

// NewHashTableSet crea un conjunto vacío.
func NewHashTableSet[T comparable]() *HashTableSet[T] {
	// Completar
	return nil
}

func (s *HashTableSet[T]) Contains(element T) bool {
	// Completar
	return false
}

func (s *HashTableSet[T]) Add(element T) {
	// Completar
}

func (s *HashTableSet[T]) Remove(element T) {
	// Completar
}

func (s *HashTableSet[T]) Size() int {
	// Completar
	return 0
}

func (s *HashTableSet[T]) Values() []T {
	// Completar
	return nil
}

func (s *HashTableSet[T]) String() string {
	// Completar
	return ""
}

func (s *HashTableSet[T]) Union(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *HashTableSet[T]) Intersection(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *HashTableSet[T]) Difference(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *HashTableSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	// Completar
	return nil
}

func (s *HashTableSet[T]) Subset(other Set[T]) bool {
	// Completar
	return false
}

func (s *HashTableSet[T]) Superset(other Set[T]) bool {
	// Completar
	return false
}
