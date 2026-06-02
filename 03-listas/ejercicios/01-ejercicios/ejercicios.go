package ejercicios

import "github.com/untref-ayp2/data-structures/list"

// PrependAll agrega todos los elementos al inicio de la lista, en el mismo
// orden en que aparecen en el slice.
func PrependAll[T comparable](l list.List[T], elementos []T) {
	// Completar
}

// AppendAll agrega todos los elementos al final de la lista, en el mismo orden
// en que aparecen en el slice.
func AppendAll[T comparable](l list.List[T], elementos []T) {
	// Completar
}

// ListaInvertida devuelve una nueva lista con los elementos en orden inverso.
// La lista original no se modifica.
func ListaInvertida[T comparable](original list.List[T]) list.List[T] {
	// Completar
	return nil
}

// RemoverTodos elimina todas las ocurrencias del valor en la lista.
// Devuelve la cantidad de elementos removidos.
func RemoverTodos[T comparable](l list.List[T], valor T) int {
	// Completar
	return 0
}

// CombinarListas devuelve una nueva lista que contiene los elementos de a
// seguidos por los de b, en ese orden.
func CombinarListas[T comparable](a, b list.List[T]) list.List[T] {
	// Completar
	return nil
}
