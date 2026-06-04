// Package directorios implementa operaciones de fusión de directorios
// de empleados usando conjuntos ordenados.
package directorios

import "github.com/untref-ayp2/data-structures/set"

// FusionarDirectorios devuelve un nuevo OrderedSet con todos los empleados
// de ambos directorios, ordenados alfabéticamente.
func FusionarDirectorios(a, b set.Set[string]) set.Set[string] {
	// Completar
	return nil
}

// EmpleadosEnComun devuelve los empleados que están en ambos directorios,
// ordenados alfabéticamente.
func EmpleadosEnComun(a, b set.Set[string]) set.Set[string] {
	// Completar
	return nil
}

// EmpleadosExclusivos devuelve los empleados que están en uno de los dos
// directorios pero no en ambos, ordenados alfabéticamente.
func EmpleadosExclusivos(a, b set.Set[string]) set.Set[string] {
	// Completar
	return nil
}

// CoberturaCompleta verifica si el directorio a contiene a todos los
// empleados del directorio b (a es superconjunto de b).
func CoberturaCompleta(a, b set.Set[string]) bool {
	// Completar
	return false
}
