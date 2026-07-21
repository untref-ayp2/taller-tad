package hashtable

// HashTableOpenAddressing implementa HashTable usando sondeo lineal
// (linear probing) sobre un arreglo.
//
// Para hashear una clave genérica K comparable, se recomienda usar el paquete
// hash/maphash de la biblioteca estándar (import "hash/maphash").
type HashTableOpenAddressing[K comparable, V any] struct {
	// Completar
}

// NewHashTableOpenAddressing crea una nueva tabla de hash vacía
// con capacidad inicial 16 y factor de carga 0.75.
func NewHashTableOpenAddressing[K comparable, V any]() *HashTableOpenAddressing[K, V] {
	// Completar
	return nil
}

// Put agrega un nuevo par clave-valor.
// Si la clave ya existe, actualiza el valor asociado.
func (ht *HashTableOpenAddressing[K, V]) Put(key K, value V) {
	// Completar
}

// Get devuelve el valor asociado a la clave.
// Error si la clave no existe.
func (ht *HashTableOpenAddressing[K, V]) Get(key K) (V, error) {
	// Completar
	var zero V
	return zero, nil
}

// Delete elimina el par clave-valor asociado a la clave.
// Error si la clave no existe.
func (ht *HashTableOpenAddressing[K, V]) Delete(key K) error {
	// Completar
	return nil
}

// Contains devuelve true si la clave existe en la tabla.
func (ht *HashTableOpenAddressing[K, V]) Contains(key K) bool {
	// Completar
	return false
}

// Size devuelve la cantidad de elementos en la tabla.
func (ht *HashTableOpenAddressing[K, V]) Size() int {
	// Completar
	return 0
}

// IsEmpty devuelve true si la tabla no tiene elementos.
func (ht *HashTableOpenAddressing[K, V]) IsEmpty() bool {
	// Completar
	return false
}

// Keys devuelve una lista con todas las claves.
func (ht *HashTableOpenAddressing[K, V]) Keys() []K {
	// Completar
	return nil
}

// Values devuelve una lista con todos los valores.
func (ht *HashTableOpenAddressing[K, V]) Values() []V {
	// Completar
	return nil
}

// Clear elimina todos los elementos de la tabla.
func (ht *HashTableOpenAddressing[K, V]) Clear() {
	// Completar
}

// String devuelve una representación estilo Python dict.
// Formato: {clave1: valor1, clave2: valor2}
func (ht *HashTableOpenAddressing[K, V]) String() string {
	// Completar
	return "{}"
}
