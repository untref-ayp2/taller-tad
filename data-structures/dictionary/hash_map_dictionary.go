package dictionary

import "github.com/untref-ayp2/data-structures/hashtable"

// HashMapDictionary implementa Dictionary[K, V] sobre una HashTable.
// Delega todas las operaciones en la tabla de hash recibida por constructor.
type HashMapDictionary[K comparable, V any] struct {
	table hashtable.HashTable[K, V]
}

// NewHashMapDictionary crea un diccionario vacío que usa la tabla de hash
// proporcionada como almacenamiento subyacente.
func NewHashMapDictionary[K comparable, V any](table hashtable.HashTable[K, V]) *HashMapDictionary[K, V] {
	// Completar
	return nil
}

// Set almacena un par clave-valor. Si la clave ya existe, actualiza su valor.
func (d *HashMapDictionary[K, V]) Set(key K, value V) {
	// Completar
}

// Get devuelve el valor asociado a la clave.
// El booleano indica si la clave existe (sigue el idiom coma-ok de Go).
func (d *HashMapDictionary[K, V]) Get(key K) (V, bool) {
	// Completar
	var zero V
	return zero, false
}

// Delete elimina la clave y su valor del diccionario. Retorna error si no existe.
func (d *HashMapDictionary[K, V]) Delete(key K) error {
	// Completar
	return nil
}

// Contains devuelve true si la clave está presente en el diccionario.
func (d *HashMapDictionary[K, V]) Contains(key K) bool {
	// Completar
	return false
}

// Size devuelve la cantidad total de pares clave-valor almacenados.
func (d *HashMapDictionary[K, V]) Size() int {
	// Completar
	return 0
}

// Keys devuelve un slice con todas las claves presentes en el diccionario.
func (d *HashMapDictionary[K, V]) Keys() []K {
	// Completar
	return nil
}

// Values devuelve un slice con todos los valores presentes en el diccionario.
func (d *HashMapDictionary[K, V]) Values() []V {
	// Completar
	return nil
}

// String devuelve una representación en cadena del diccionario.
func (d *HashMapDictionary[K, V]) String() string {
	// Completar
	return "{}"
}
