package stack

import "errors"

// ErrStackEmpty se retorna cuando se opera sobre una pila vacía.
var ErrStackEmpty = errors.New("la pila está vacía")

// Stack es la interfaz que deben implementar todas las pilas.
type Stack[T any] interface {
	// Push agrega un elemento al tope de la pila.
	Push(val T)
	// Pop elimina y devuelve el elemento del tope.
	// Error si la pila está vacía.
	Pop() (T, error)
	// Top devuelve el elemento del tope sin eliminarlo.
	// Error si la pila está vacía.
	Top() (T, error)
	// IsEmpty devuelve true si la pila no tiene elementos.
	IsEmpty() bool
}
