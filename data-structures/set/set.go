// Package set define la interfaz que deben implementar todos los conjuntos.
package set

// Set es la interfaz que deben implementar todos los conjuntos.
type Set[T comparable] interface {
	// Contains verifica si un elemento pertenece al conjunto.
	Contains(element T) bool
	// Add agrega un elemento al conjunto. Si ya existe, no tiene efecto.
	Add(element T)
	// Remove elimina un elemento del conjunto. Si no existe, no tiene efecto.
	Remove(element T)
	// Size devuelve la cantidad de elementos del conjunto.
	Size() int
	// Values devuelve un slice con todos los elementos, sin orden definido.
	Values() []T
	// String devuelve una representación textual del conjunto.
	String() string

	// Union devuelve un nuevo conjunto con los elementos de ambos conjuntos.
	Union(other Set[T]) Set[T]
	// Intersection devuelve un nuevo conjunto con los elementos comunes.
	Intersection(other Set[T]) Set[T]
	// Difference devuelve un nuevo conjunto con los elementos del receptor
	// que no están en other.
	Difference(other Set[T]) Set[T]
	// SymmetricDifference devuelve un nuevo conjunto con los elementos que
	// están en uno de los conjuntos pero no en ambos.
	SymmetricDifference(other Set[T]) Set[T]
	// Subset verifica si el receptor es subconjunto de other.
	Subset(other Set[T]) bool
	// Superset verifica si el receptor es superconjunto de other.
	Superset(other Set[T]) bool
}
