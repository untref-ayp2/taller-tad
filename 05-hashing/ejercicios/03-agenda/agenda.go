// Package agenda implementa una agenda telefónica usando una tabla de hash.
package agenda

import "github.com/untref-ayp2/data-structures/hashtable"

// Contacto representa la información de un contacto.
type Contacto struct {
	Nombre   string
	Telefono string
	Email    string
}

// Agenda permite gestionar contactos.
// La clave es el nombre del contacto, el valor es un Contacto.
type Agenda struct {
	contactos hashtable.HashTable[string, Contacto]
}

// NuevaAgenda crea una agenda vacía.
func NuevaAgenda() *Agenda {
	return nil
}

// AgregarContacto agrega un nuevo contacto.
// Si el contacto ya existe, actualiza sus datos.
func (a *Agenda) AgregarContacto(contacto Contacto) {
}

// BuscarContacto devuelve un contacto por su nombre.
// Devuelve error si no existe.
func (a *Agenda) BuscarContacto(nombre string) (Contacto, error) {
	var zero Contacto
	return zero, nil
}

// EliminarContacto elimina un contacto por su nombre.
// Devuelve error si no existe.
func (a *Agenda) EliminarContacto(nombre string) error {
	return nil
}

// ExisteContacto devuelve true si el nombre está registrado.
func (a *Agenda) ExisteContacto(nombre string) bool {
	return false
}

// ListarContactos devuelve todos los contactos de la agenda.
func (a *Agenda) ListarContactos() []Contacto {
	return nil
}

// Cantidad devuelve la cantidad de contactos en la agenda.
func (a *Agenda) Cantidad() int {
	return 0
}
