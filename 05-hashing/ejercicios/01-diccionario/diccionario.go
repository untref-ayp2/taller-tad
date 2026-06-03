// Package diccionario implementa un traductor básico español-inglés usando
// una tabla de hash.
package diccionario

import "github.com/untref-ayp2/data-structures/hashtable"

// Diccionario permite almacenar traducciones español-inglés y traducir
// oraciones palabra por palabra.
type Diccionario struct {
	traducciones hashtable.HashTable[string, string]
}

// NuevoDiccionario crea un diccionario vacío.
func NuevoDiccionario() *Diccionario {
	return nil
}

// Agregar agrega un par (español → inglés) al diccionario.
// Si la palabra ya existe, actualiza su traducción.
func (d *Diccionario) Agregar(espanol, ingles string) {
}

// Traducir traduce una palabra del español al inglés.
// Devuelve la palabra original si no encuentra la traducción.
func (d *Diccionario) Traducir(palabra string) string {
	return ""
}

// TraducirOracion traduce una oración completa palabra por palabra.
// Las palabras se separan por espacios. Si una palabra no está en el
// diccionario, se deja sin traducir.
func (d *Diccionario) TraducirOracion(oracion string) string {
	return ""
}

// ListarPalabras devuelve todas las palabras en español registradas.
func (d *Diccionario) ListarPalabras() []string {
	return nil
}

// Cantidad devuelve la cantidad de palabras registradas.
func (d *Diccionario) Cantidad() int {
	return 0
}
