package colacircular

// ColaCircular implementa una cola circular (ring buffer) sobre un arreglo de
// tamaño fijo. Los elementos se agregan al final y se extraen del frente.
// Cuando el arreglo se llena, no se pueden agregar más elementos hasta que
// se extraiga alguno.
type ColaCircular[T any] struct {
	// Completar
}

// NewColaCircular crea una cola circular vacía con la capacidad indicada.
// La capacidad debe ser mayor o igual a 1.
func NewColaCircular[T any](capacidad int) *ColaCircular[T] {
	// Completar
	return nil
}

// Enqueue agrega un elemento al final de la cola.
// Devuelve error si la cola está llena.
func (c *ColaCircular[T]) Enqueue(val T) error {
	// Completar
	return nil
}

// Dequeue extrae y devuelve el elemento del frente.
// Devuelve error si la cola está vacía.
func (c *ColaCircular[T]) Dequeue() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Front devuelve el elemento del frente sin extraerlo.
// Devuelve error si la cola está vacía.
func (c *ColaCircular[T]) Front() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// IsEmpty devuelve true si la cola no tiene elementos.
func (c *ColaCircular[T]) IsEmpty() bool {
	// Completar
	return false
}

// IsFull devuelve true si la cola alcanzó su capacidad máxima.
func (c *ColaCircular[T]) IsFull() bool {
	// Completar
	return false
}

// Size devuelve la cantidad de elementos actualmente en la cola.
func (c *ColaCircular[T]) Size() int {
	// Completar
	return 0
}
