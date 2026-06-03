// Package conteo implementa un contador de frecuencias de palabras usando
// una tabla de hash.
package conteo

import "github.com/untref-ayp2/data-structures/hashtable"

// ContadorDePalabras permite contar la frecuencia de cada palabra en un texto.
type ContadorDePalabras struct {
	frecuencias hashtable.HashTable[string, int]
}

// NuevoContador crea un contador vacío.
func NuevoContador() *ContadorDePalabras {
	return nil
}

// AgregarTexto procesa un texto y actualiza las frecuencias de cada palabra.
// Las palabras se separan por espacios. No distingue mayúsculas/minúsculas.
func (c *ContadorDePalabras) AgregarTexto(texto string) {
}

// Frecuencia devuelve la frecuencia de una palabra específica.
func (c *ContadorDePalabras) Frecuencia(palabra string) int {
	return 0
}

// TotalPalabras devuelve la cantidad total de palabras procesadas.
func (c *ContadorDePalabras) TotalPalabras() int {
	return 0
}

// PalabrasUnicas devuelve la cantidad de palabras distintas encontradas.
func (c *ContadorDePalabras) PalabrasUnicas() int {
	return 0
}

// TopN devuelve las n palabras más frecuentes y su frecuencia.
// Si hay menos de n palabras distintas, devuelve todas.
// El orden es de mayor a menor frecuencia.
type ParFrecuencia struct {
	Palabra   string
	Frecuencia int
}

func (c *ContadorDePalabras) TopN(n int) []ParFrecuencia {
	return nil
}

// ListarPalabras devuelve todas las palabras registradas.
func (c *ContadorDePalabras) ListarPalabras() []string {
	return nil
}
