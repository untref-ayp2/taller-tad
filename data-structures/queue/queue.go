package queue

import "errors"

// ErrQueueEmpty se retorna cuando se opera sobre una cola vacía.
var ErrQueueEmpty = errors.New("la cola está vacía")

// Queue es la interfaz que deben implementar todas las colas.
type Queue[T any] interface {
	// Enqueue agrega un elemento al final de la cola.
	Enqueue(val T)
	// Dequeue elimina y devuelve el elemento del frente.
	// Error si la cola está vacía.
	Dequeue() (T, error)
	// Front devuelve el elemento del frente sin eliminarlo.
	// Error si la cola está vacía.
	Front() (T, error)
	// IsEmpty devuelve true si la cola no tiene elementos.
	IsEmpty() bool
}
