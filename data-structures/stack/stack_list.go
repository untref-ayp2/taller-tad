package stack

// ListStack implementa Stack[T] sobre una List[T].
type ListStack[T comparable] struct {
	// Completar
}

// NewListStack crea una nueva pila vacía.
func NewListStack[T comparable]() *ListStack[T] {
	// Completar
	return nil
}

// Push agrega un elemento al tope de la pila.
func (s *ListStack[T]) Push(val T) {
	// Completar
}

// Pop elimina y devuelve el elemento del tope.
// Error si la pila está vacía.
func (s *ListStack[T]) Pop() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Top devuelve el elemento del tope sin eliminarlo.
// Error si la pila está vacía.
func (s *ListStack[T]) Top() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// IsEmpty devuelve true si la pila no tiene elementos.
func (s *ListStack[T]) IsEmpty() bool {
	// Completar
	return false
}
