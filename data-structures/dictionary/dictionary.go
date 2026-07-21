// Package dictionary define la interfaz que deben implementar todos los diccionarios.
package dictionary

// Dictionary define las operaciones abstractas de un diccionario clave-valor.
type Dictionary[K comparable, V any] interface {
	// Set almacena un par clave-valor. Si la clave ya existe, actualiza su valor.
	Set(key K, value V)

	// Get devuelve el valor asociado a la clave.
	// El booleano indica si la clave existe en el diccionario
	// (sigue el idiom coma-ok de Go: key, ok := dict.Get("clave")).
	Get(key K) (V, bool)

	// Delete elimina la clave y su valor del diccionario. Retorna error si no existe.
	Delete(key K) error

	// Contains devuelve true si la clave está presente en el diccionario.
	Contains(key K) bool

	// Size devuelve la cantidad total de pares clave-valor almacenados.
	Size() int

	// Keys devuelve un slice con todas las claves presentes en el diccionario.
	Keys() []K

	// Values devuelve un slice con todos los valores presentes en el diccionario.
	Values() []V

	// String devuelve una representación en cadena del diccionario.
	String() string
}
