package queue

// SliceQueue implementa Queue[T] sobre un slice.
type SliceQueue[T any] struct {
	// Completar
	data []T
}

// NewSliceQueue crea una nueva cola vacía.
func NewSliceQueue[T any]() *SliceQueue[T] {
	// Completar
	return nil
}

// Enqueue agrega un elemento al final de la cola.
func (q *SliceQueue[T]) Enqueue(val T) {
	// Completar
}

// Dequeue elimina y devuelve el elemento del frente.
// Error si la cola está vacía.
func (q *SliceQueue[T]) Dequeue() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Front devuelve el elemento del frente sin eliminarlo.
// Error si la cola está vacía.
func (q *SliceQueue[T]) Front() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// IsEmpty devuelve true si la cola no tiene elementos.
func (q *SliceQueue[T]) IsEmpty() bool {
	// Completar
	return false
}
