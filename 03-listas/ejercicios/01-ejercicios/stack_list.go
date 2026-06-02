package ejercicios

// StackList implementa una pila sobre una List[T].
//
// Complexidades esperadas:
//
//	Push: O(?)  (justificar)
//	Pop:  O(?)
//	Top:  O(?)
type StackList[T comparable] struct {
	// Completar
}

// NewStackList crea una nueva pila vacía.
func NewStackList[T comparable]() *StackList[T] {
	// Completar
	return nil
}

// Push agrega un elemento al tope.
func (s *StackList[T]) Push(val T) {
	// Completar
}

// Pop elimina y devuelve el elemento del tope.
// Retorna error si la pila está vacía.
func (s *StackList[T]) Pop() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// Top devuelve el elemento del tope sin eliminarlo.
// Retorna error si la pila está vacía.
func (s *StackList[T]) Top() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

// IsEmpty devuelve true si la pila no tiene elementos.
func (s *StackList[T]) IsEmpty() bool {
	// Completar
	return false
}
