package ejercicios

import "github.com/untref-ayp2/data-structures/queue"

// QueueList implementa Queue[T] usando internamente SinglyLinkedList.
type QueueList[T any] struct {
	// Completar
}

// NewQueueList crea una cola vacía basada en lista.
func NewQueueList[T any]() *QueueList[T] {
	// Completar
	return nil
}

func (q *QueueList[T]) Enqueue(val T) {
	// Completar
}

func (q *QueueList[T]) Dequeue() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

func (q *QueueList[T]) Front() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

func (q *QueueList[T]) IsEmpty() bool {
	// Completar
	return false
}

var _ queue.Queue[any] = (*QueueList[any])(nil)
