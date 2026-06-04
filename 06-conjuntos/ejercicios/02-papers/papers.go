// Package papers implementa un buscador de papers por palabras clave.
package papers

import "github.com/untref-ayp2/data-structures/set"

// BusquedaExacta devuelve los papers que contienen todas las keywords
// buscadas.
func BusquedaExacta(papers map[string]set.Set[string], busqueda set.Set[string]) []string {
	// Completar
	return nil
}

// BusquedaRelajada devuelve los papers que contienen al menos una keyword
// de la búsqueda, ordenados por cantidad de coincidencias descendente.
func BusquedaRelajada(papers map[string]set.Set[string], busqueda set.Set[string]) []string {
	// Completar
	return nil
}

// SugerirKeywords recibe un conjunto de keywords buscadas, encuentra los
// 5 papers con más coincidencias y devuelve las keywords de esos papers
// que no estaban en la búsqueda original, sin repeticiones y ordenadas
// alfabéticamente.
func SugerirKeywords(papers map[string]set.Set[string], busqueda set.Set[string]) []string {
	// Completar
	return nil
}
