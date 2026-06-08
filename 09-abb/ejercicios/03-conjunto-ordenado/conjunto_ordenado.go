// Package conjunto_ordenado implementa un conjunto ordenado usando un ABB.
package conjunto_ordenado

import "github.com/untref-ayp2/data-structures/binarysearchtree"

// ConjuntoOrdenado representa un conjunto que mantiene sus elementos
// ordenados. Está implementado sobre un ABB, lo que permite inserciones,
// búsquedas y eliminaciones en O(log n) promedio, y recorrido ordenado
// mediante inorder.
type ConjuntoOrdenado[T any] struct {
	arbol *binarysearchtree.BinarySearchTree[T]
}

// NuevoConjuntoOrdenado crea un conjunto ordenado vacío.
// La función cmp debe devolver un valor negativo si a < b,
// cero si a == b, y positivo si a > b.
func NuevoConjuntoOrdenado[T any](cmp func(T, T) int) *ConjuntoOrdenado[T] {
	// Completar
	return nil
}

// Agregar inserta un elemento en el conjunto.
// Si ya existe, no tiene efecto.
func (c *ConjuntoOrdenado[T]) Agregar(elemento T) {
	// Completar
}

// Contiene verifica si un elemento pertenece al conjunto.
func (c *ConjuntoOrdenado[T]) Contiene(elemento T) bool {
	// Completar
	return false
}

// Eliminar remueve un elemento del conjunto.
// Si no existe, no tiene efecto.
func (c *ConjuntoOrdenado[T]) Eliminar(elemento T) {
	// Completar
}

// Cantidad devuelve la cantidad de elementos del conjunto.
func (c *ConjuntoOrdenado[T]) Cantidad() int {
	// Completar
	return 0
}

// Valores devuelve un slice con todos los elementos en orden ascendente.
func (c *ConjuntoOrdenado[T]) Valores() []T {
	// Completar
	return nil
}
