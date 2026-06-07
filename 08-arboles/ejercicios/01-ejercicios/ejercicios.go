package ejercicios

import (
	"github.com/untref-ayp2/data-structures/tree"
)

// Preorden devuelve un slice con los valores de los nodos en orden preorden
// (raíz, izquierdo, derecho).
func Preorden[T any](root *tree.TreeNode[T]) []T {
	// Completar
	return nil
}

// Inorden devuelve un slice con los valores de los nodos en orden inorden
// (izquierdo, raíz, derecho).
func Inorden[T any](root *tree.TreeNode[T]) []T {
	// Completar
	return nil
}

// Postorden devuelve un slice con los valores de los nodos en orden postorden
// (izquierdo, derecho, raíz).
func Postorden[T any](root *tree.TreeNode[T]) []T {
	// Completar
	return nil
}

// Evaluar evalúa un árbol de expresión aritmética y devuelve el resultado.
// El árbol debe tener operadores (+, -, *, /) en los nodos internos y
// números enteros en las hojas. Usa recorrido postorden con una pila.
func Evaluar(arbol *tree.TreeNode[string]) (int, error) {
	// Completar
	return 0, nil
}
