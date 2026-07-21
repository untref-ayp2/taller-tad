package priorityqueue

import (
	"fmt"

	"github.com/untref-ayp2/data-structures/heap"
)

// PriorityQueue implementa una cola de prioridad
// inyectando un heap binario por composición.
type PriorityQueue[T any] struct {
	// Completar
}

// NewPriorityQueue crea una PriorityQueue que usa el heap provisto.
func NewPriorityQueue[T any](h heap.Heap[T]) *PriorityQueue[T] {
	// Completar
	return nil
}

// Enqueue agrega un elemento a la cola de prioridad.
func (pq *PriorityQueue[T]) Enqueue(val T) {
	// Completar
}

// Dequeue remueve y retorna el elemento de mayor prioridad.
func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Front retorna el elemento de mayor prioridad sin removerlo.
func (pq *PriorityQueue[T]) Front() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Size retorna la cantidad de elementos.
func (pq *PriorityQueue[T]) Size() int {
	// Completar
	return 0
}

// IsEmpty retorna true si la cola está vacía.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	// Completar
	return false
}

// ErrorColaVacia se retorna al operar sobre una cola vacía.
var ErrorColaVacia = fmt.Errorf("la cola de prioridad está vacía")
