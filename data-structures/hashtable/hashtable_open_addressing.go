package hashtable

import (
	"errors"
	"fmt"
)

// hashTableEntry representa una entrada en la tabla hash, que contiene una
// clave, su valor asociado y un marcador de eliminación.
//
// El campo deleted (tombstone) evita perder la posición al eliminar: en lugar
// de compactar el arreglo, se marca la entrada como eliminada para que las
// búsquedas lineales no se interrumpan.
type hashTableEntry[K comparable, V any] struct {
	key     K
	value   V
	deleted bool
}

// HashTableOpenAddressing es una implementación de hashing cerrado con sondeo
// lineal que utiliza un arreglo para almacenar pares clave-valor.
type HashTableOpenAddressing[K comparable, V any] struct {
	buckets    []*hashTableEntry[K, V]
	size       int
	capacity   int
	loadFactor float32
	threshold  int
}

// NewHashTableOpenAddressing crea una nueva tabla de hash cerrada con la
// capacidad y el factor de carga especificados.
//
// - Si la capacidad es igual a 0, se establece en 17.
//
// - Si el factor de carga es menor o igual a 0 o mayor que 1, se establece en
// 0.75.
//
// - Si la capacidad no es un número primo, se redimensiona a la siguiente
// capacidad primo mayor o igual a la capacidad especificada.
func NewHashTableOpenAddressing[K comparable, V any](capacity int, loadFactor float32) *HashTableOpenAddressing[K, V] {
	if capacity == 0 {
		capacity = 17
	}
	if loadFactor <= 0 || loadFactor > 1 {
		loadFactor = 0.75
	}
	if !isPrime(capacity) {
		capacity = nextPrime(capacity)
	}
	return &HashTableOpenAddressing[K, V]{
		buckets:    make([]*hashTableEntry[K, V], capacity),
		size:       0,
		capacity:   capacity,
		loadFactor: loadFactor,
		threshold:  int(float32(capacity) * loadFactor),
	}
}

// Put agrega un nuevo par clave-valor a la tabla de hash. Si la clave ya
// existe, actualiza el valor asociado a la clave.
func (ht *HashTableOpenAddressing[K, V]) Put(key K, value V) {
	if index, esta := ht.getIndex(key); esta {
		ht.buckets[index].value = value
		return
	}

	if ht.size >= ht.threshold {
		ht.resize()
	}

	index := ht.hash(key) % ht.capacity
	for {
		if ht.buckets[index] == nil || ht.buckets[index].deleted {
			ht.buckets[index] = &hashTableEntry[K, V]{key: key, value: value}
			ht.size++
			return
		}
		index = (index + 1) % ht.capacity
	}
}

// Get devuelve el valor asociado a la clave y nil. Si la clave no existe,
// devuelve el zero value de V y un error.
func (ht *HashTableOpenAddressing[K, V]) Get(key K) (V, error) {
	index, exists := ht.getIndex(key)
	if !exists {
		var zero V
		return zero, errors.New("key not found")
	}
	return ht.buckets[index].value, nil
}

// Delete elimina el par clave-valor asociado a la clave.
// Devuelve error si la clave no existe.
//
// Usa eliminación perezosa (tombstone): marca la entrada como deleted en lugar
// de eliminarla del arreglo. Esto mantiene la continuidad de las sondas
// lineales: si se eliminara la entrada, las búsquedas posteriores podrían no
// encontrar elementos que están más allá en la secuencia de sondeo.
func (ht *HashTableOpenAddressing[K, V]) Delete(key K) error {
	index, exists := ht.getIndex(key)
	if !exists {
		return errors.New("key not found")
	}
	ht.buckets[index].deleted = true
	var zero V
	ht.buckets[index].value = zero
	ht.size--
	return nil
}

// Contains devuelve true si la clave existe en la tabla.
func (ht *HashTableOpenAddressing[K, V]) Contains(key K) bool {
	_, exists := ht.getIndex(key)
	return exists
}

// Keys devuelve una lista de todas las claves en la tabla de hash.
//
// Pre-asigna el slice con capacidad ht.size para evitar realocaciones durante
// los append. Solo incluye entradas activas (no eliminadas).
func (ht *HashTableOpenAddressing[K, V]) Keys() []K {
	keys := make([]K, 0, ht.size)
	for _, node := range ht.buckets {
		if node != nil && !node.deleted {
			keys = append(keys, node.key)
		}
	}
	return keys
}

// Values devuelve una lista de todos los valores en la tabla de hash.
//
// Pre-asigna el slice con capacidad ht.size para evitar realocaciones. Solo
// incluye valores de entradas activas (no eliminadas).
func (ht *HashTableOpenAddressing[K, V]) Values() []V {
	values := make([]V, 0, ht.size)
	for _, node := range ht.buckets {
		if node != nil && !node.deleted {
			values = append(values, node.value)
		}
	}
	return values
}

// Size devuelve el número de elementos en la tabla de hash.
func (ht *HashTableOpenAddressing[K, V]) Size() int {
	return ht.size
}

// IsEmpty devuelve true si la tabla de hash está vacía, false en caso contrario.
func (ht *HashTableOpenAddressing[K, V]) IsEmpty() bool {
	return ht.size == 0
}

// Clear elimina todos los elementos de la tabla de hash.
func (ht *HashTableOpenAddressing[K, V]) Clear() {
	ht.buckets = make([]*hashTableEntry[K, V], ht.capacity)
	ht.size = 0
}

// String devuelve una representación en cadena de la tabla de hash.
//
// Formato: {clave1: valor1, clave2: valor2}. Tras recorrer todas las entradas
// activas se elimina la última coma y espacio para mantener el formato limpio.
func (ht *HashTableOpenAddressing[K, V]) String() string {
	result := "{"
	for _, node := range ht.buckets {
		if node != nil && !node.deleted {
			result += fmt.Sprintf("%v: %v", node.key, node.value) + ", "
		}
	}
	if len(result) > 1 {
		result = result[:len(result)-2]
	}
	result += "}"
	return result
}

// Funciones internas

// a es una constante utilizada para calcular el hash de un string
const a = 11

// hash calcula el índice del bucket para una clave dada.
//
// Se utiliza la técnica de Multiplicación Polinómica con el método de Horner:
// evalúa el polinomio c1*a^{n-1} + c2*a^{n-2} + ... + cn*a^0 aplicando
// la propiedad distributiva para evitar el cálculo costoso de potencias.
//
// La constante a=11 es un número primo pequeño que ofrece buena dispersión
// para cadenas de texto.
//
// Como K es cualquier tipo comparable, se convierte la clave a string con
// fmt.Sprintf para aplicar el algoritmo polinómico sobre su representación.
//
// El módulo por la capacidad NO se aplica acá sino en el llamante, para que
// este mismo valor pueda reusarse al buscar en distintas capacidades (ej.
// durante el redimensionamiento).
func (ht *HashTableOpenAddressing[K, V]) hash(key K) int {
	s := fmt.Sprintf("%v", key)
	var hash int
	for _, c := range s {
		hash = hash*a + int(c)
	}
	return hash
}

// getIndex devuelve el índice del bucket para una clave dada y un booleano que
// indica si la clave existe.
//
// Realiza una sonda lineal (linear probing): calcula el índice inicial con
// hash(key) % capacity y, si el bucket está ocupado por otra clave, avanza
// secuencialmente hasta encontrar la clave buscada o un bucket vacío (nil).
//
// La condición i < ht.capacity limita la búsqueda para evitar un bucle
// infinito si la tabla estuviera completamente llena (sin buckets nil).
// Al encontrar un bucket nil detiene la búsqueda porque en hashing cerrado
// con sondeo lineal los elementos no saltan buckets vacíos.
func (ht *HashTableOpenAddressing[K, V]) getIndex(key K) (int, bool) {
	for i, index := 0, ht.hash(key)%ht.capacity; i < ht.capacity && ht.buckets[index] != nil; i, index = i+1, (index+1)%ht.capacity {
		if !ht.buckets[index].deleted && ht.buckets[index].key == key {
			return index, true
		}
	}
	return 0, false
}

// resize redimensiona la tabla de hash y reubica todos los elementos en la
// nueva tabla.
//
// El nuevo tamaño es el siguiente número primo mayor o igual al doble de la
// capacidad actual.
//
// Se itera sobre todos los buckets del arreglo viejo y se reinsertan solo las
// entradas activas (no eliminadas). Las entradas marcadas como deleted se
// descartan, liberando así la memoria que ocupaban.
//
// Tras la reubicación se actualizan capacity y threshold. El factor de carga
// loadFactor se mantiene igual. Esto asegura que la próxima vez que se alcance
// el umbral se vuelva a redimensionar.
func (ht *HashTableOpenAddressing[K, V]) resize() {
	newCapacity := nextPrime(ht.capacity * 2)
	newBuckets := make([]*hashTableEntry[K, V], newCapacity)

	for _, node := range ht.buckets {
		if node != nil && !node.deleted {
			index := ht.hash(node.key) % newCapacity
			for newBuckets[index] != nil {
				index = (index + 1) % newCapacity
			}
			newBuckets[index] = node
		}
	}

	ht.buckets = newBuckets
	ht.capacity = newCapacity
	ht.threshold = int(float32(newCapacity) * ht.loadFactor)
}

// nextPrime devuelve el siguiente número primo mayor o igual a n.
//
// La búsqueda es secuencial: prueba cada número a partir de n hacia arriba
// hasta encontrar un primo. Como la frecuencia de primos es aproximadamente
// n / ln(n), rara vez recorre muchos números.
func nextPrime(n int) int {
	if n <= 1 {
		return 2
	}
	for i := n; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

// isPrime devuelve true si n es un número primo, false en caso contrario.
//
// Optimización 6k ± 1: todo primo mayor a 3 puede expresarse como 6k ± 1.
// Esto permite verificar divisibilidad solo con números de esa forma,
// reduciendo las iteraciones a ~1/3 de una verificación ingenua.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
