package ejercicios

// QueueList implementa una cola sobre una List[T].
//
// Complexidades esperadas:
//
//	Enqueue: O(?)  (justificar)
//	Dequeue: O(?)
//	Front:   O(?)
type QueueList[T comparable] struct {
	// Completar
}

// NewQueueList crea una nueva cola vacía.
func NewQueueList[T comparable]() *QueueList[T] {
	// Completar
	return nil
}

// Enqueue agrega un elemento al final.
func (q *QueueList[T]) Enqueue(val T) {
	// Completar
}

// Dequeue elimina y devuelve el elemento del frente.
// Retorna error si la cola está vacía.
func (q *QueueList[T]) Dequeue() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Front devuelve el elemento del frente sin eliminarlo.
// Retorna error si la cola está vacía.
func (q *QueueList[T]) Front() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// IsEmpty devuelve true si la cola no tiene elementos.
func (q *QueueList[T]) IsEmpty() bool {
	// Completar
	return false
}
