package binarysearchtree

import "github.com/untref-ayp2/data-structures/tree"

// BinarySearchTree representa un árbol binario de búsqueda construido
// sobre un árbol binario genérico. La función de comparación permite
// trabajar con cualquier tipo de dato sin restringir a cmp.Ordered.
type BinarySearchTree[T any] struct {
	bt  *tree.BinaryTree[T]
	cmp func(T, T) int
}

// NewBinarySearchTree crea un árbol binario de búsqueda vacío.
// La función cmp debe devolver un valor negativo si a < b,
// cero si a == b, y positivo si a > b.
func NewBinarySearchTree[T any](cmp func(T, T) int) *BinarySearchTree[T] {
	// Completar
	return nil
}

// Insert inserta un valor en el árbol. No inserta duplicados.
func (bst *BinarySearchTree[T]) Insert(value T) {
	// Completar
}

// Search busca un valor en el árbol.
func (bst *BinarySearchTree[T]) Search(value T) bool {
	// Completar
	return false
}

// Delete elimina un valor del árbol. No hace nada si no existe.
func (bst *BinarySearchTree[T]) Delete(value T) {
	// Completar
}

// Height devuelve la altura del árbol.
func (bst *BinarySearchTree[T]) Height() int {
	// Completar
	return 0
}

// Size devuelve la cantidad de nodos en el árbol.
func (bst *BinarySearchTree[T]) Size() int {
	// Completar
	return 0
}

// InorderTraversal devuelve un slice con los valores en recorrido inorder.
func (bst *BinarySearchTree[T]) InorderTraversal() []T {
	// Completar
	return nil
}

// PreorderTraversal devuelve un slice con los valores en recorrido preorder.
func (bst *BinarySearchTree[T]) PreorderTraversal() []T {
	// Completar
	return nil
}

// PostorderTraversal devuelve un slice con los valores en recorrido postorder.
func (bst *BinarySearchTree[T]) PostorderTraversal() []T {
	// Completar
	return nil
}
