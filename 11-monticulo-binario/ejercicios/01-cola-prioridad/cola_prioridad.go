package colaprioridad

import (
	_ "github.com/untref-ayp2/data-structures/heap"
)

// Persona representa una persona con nombre y edad.
type Persona struct {
	nombre string
	edad   int
}

// ColaPrioridad es una cola de prioridad de personas.
// La prioridad se define por la edad: mayor edad = mayor prioridad.
// A igual edad, se respeta el orden de llegada.
type ColaPrioridad struct {
	// Completar
}

// NuevaColaPrioridad crea una cola de prioridad vacía.
func NuevaColaPrioridad() *ColaPrioridad {
	// Completar
	return nil
}

// Agregar agrega una persona a la cola con prioridad según su edad.
func (cp *ColaPrioridad) Agregar(p Persona) {
	// Completar
}

// Atender elimina y devuelve la persona con mayor prioridad.
func (cp *ColaPrioridad) Atender() (Persona, error) {
	// Completar
	return Persona{}, nil
}

// Siguiente devuelve la persona con mayor prioridad sin eliminarla.
func (cp *ColaPrioridad) Siguiente() (Persona, error) {
	// Completar
	return Persona{}, nil
}

// Cantidad devuelve la cantidad de personas en la cola.
func (cp *ColaPrioridad) Cantidad() int {
	// Completar
	return 0
}

// EstaVacia devuelve true si la cola no tiene personas.
func (cp *ColaPrioridad) EstaVacia() bool {
	// Completar
	return false
}
