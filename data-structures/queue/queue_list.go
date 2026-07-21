package queue

// ListQueue implementa Queue[T] sobre una List[T].
type ListQueue[T comparable] struct {
	// Completar
}

// NewListQueue crea una nueva cola vacía.
func NewListQueue[T comparable]() *ListQueue[T] {
	// Completar
	return nil
}

// Enqueue agrega un elemento al final de la cola.
func (q *ListQueue[T]) Enqueue(val T) {
	// Completar
}

// Dequeue elimina y devuelve el elemento del frente.
// Error si la cola está vacía.
func (q *ListQueue[T]) Dequeue() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Front devuelve el elemento del frente sin eliminarlo.
// Error si la cola está vacía.
func (q *ListQueue[T]) Front() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// IsEmpty devuelve true si la cola no tiene elementos.
func (q *ListQueue[T]) IsEmpty() bool {
	// Completar
	return false
}
