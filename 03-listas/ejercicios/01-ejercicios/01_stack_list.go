package ejercicios

import "github.com/untref-ayp2/data-structures/stack"

// StackList implementa Stack[T] usando internamente SinglyLinkedList.
type StackList[T any] struct {
	// Completar
}

// NewStackList crea una pila vacía basada en lista.
func NewStackList[T any]() *StackList[T] {
	// Completar
	return nil
}

func (s *StackList[T]) Push(val T) {
	// Completar
}

func (s *StackList[T]) Pop() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

func (s *StackList[T]) Top() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

func (s *StackList[T]) IsEmpty() bool {
	// Completar
	return false
}

var _ stack.Stack[any] = (*StackList[any])(nil)
