package hashtable

// HashTableChaining implementa HashTable usando encadenamiento
// (separate chaining) con listas enlazadas.
//
// Para hashear una clave genérica K comparable, se recomienda usar el paquete
// hash/maphash de la biblioteca estándar (import "hash/maphash").
type HashTableChaining[K comparable, V any] struct {
	// Completar
}

// NewHashTableChaining crea una nueva tabla de hash vacía
// con capacidad inicial 16 y factor de carga 0.75.
func NewHashTableChaining[K comparable, V any]() *HashTableChaining[K, V] {
	// Completar
	return nil
}

// Put agrega un nuevo par clave-valor.
// Si la clave ya existe, actualiza el valor asociado.
func (ht *HashTableChaining[K, V]) Put(key K, value V) {
	// Completar
}

// Get devuelve el valor asociado a la clave.
// Error si la clave no existe.
func (ht *HashTableChaining[K, V]) Get(key K) (V, error) {
	// Completar
	var zero V
	return zero, nil
}

// Delete elimina el par clave-valor asociado a la clave.
// Error si la clave no existe.
func (ht *HashTableChaining[K, V]) Delete(key K) error {
	// Completar
	return nil
}

// Contains devuelve true si la clave existe en la tabla.
func (ht *HashTableChaining[K, V]) Contains(key K) bool {
	// Completar
	return false
}

// Size devuelve la cantidad de elementos en la tabla.
func (ht *HashTableChaining[K, V]) Size() int {
	// Completar
	return 0
}

// IsEmpty devuelve true si la tabla no tiene elementos.
func (ht *HashTableChaining[K, V]) IsEmpty() bool {
	// Completar
	return false
}

// Keys devuelve una lista con todas las claves.
func (ht *HashTableChaining[K, V]) Keys() []K {
	// Completar
	return nil
}

// Values devuelve una lista con todos los valores.
func (ht *HashTableChaining[K, V]) Values() []V {
	// Completar
	return nil
}

// Clear elimina todos los elementos de la tabla.
func (ht *HashTableChaining[K, V]) Clear() {
	// Completar
}

// String devuelve una representación estilo Python dict.
// Formato: {clave1: valor1, clave2: valor2}
func (ht *HashTableChaining[K, V]) String() string {
	// Completar
	return "{}"
}
