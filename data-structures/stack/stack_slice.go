package stack

// SliceStack implementa Stack[T] sobre un slice.
type SliceStack[T any] struct {
	// Completar
	data []T
}

// NewSliceStack crea una nueva pila vacía.
func NewSliceStack[T any]() *SliceStack[T] {
	// Completar
	return nil
}

// Push agrega un elemento al tope de la pila.
func (s *SliceStack[T]) Push(val T) {
	// Completar
}

// Pop elimina y devuelve el elemento del tope.
// Error si la pila está vacía.
func (s *SliceStack[T]) Pop() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Top devuelve el elemento del tope sin eliminarlo.
// Error si la pila está vacía.
func (s *SliceStack[T]) Top() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// IsEmpty devuelve true si la pila no tiene elementos.
func (s *SliceStack[T]) IsEmpty() bool {
	// Completar
	return false
}
