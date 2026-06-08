// Package guia_telefonica implementa una guía telefónica ordenada usando
// un ABB como estructura subyacente. A diferencia de la versión con tabla
// de hash (ver 05-hashing/ejercicios/03-agenda), esta implementación
// mantiene los contactos ordenados alfabéticamente por nombre.
package guia_telefonica

import (
	"cmp"
	"github.com/untref-ayp2/data-structures/binarysearchtree"
)

// Contacto representa la información de un contacto.
type Contacto struct {
	Nombre   string
	Telefono string
	Email    string
}

// GuiaTelefonica permite gestionar contactos ordenados alfabéticamente.
// Utiliza un ABB de Contacto, comparando por nombre.
type GuiaTelefonica struct {
	arbol *binarysearchtree.BinarySearchTree[Contacto]
}

// NuevaGuiaTelefonica crea una guía telefónica vacía.
// La función de comparación ordena los contactos por nombre.
func NuevaGuiaTelefonica() *GuiaTelefonica {
	cmpContacto := func(a, b Contacto) int {
		return cmp.Compare(a.Nombre, b.Nombre)
	}
	return &GuiaTelefonica{
		arbol: binarysearchtree.NewBinarySearchTree(cmpContacto),
	}
}

// AgregarContacto agrega un nuevo contacto.
// Si el contacto ya existe, actualiza sus datos.
func (g *GuiaTelefonica) AgregarContacto(contacto Contacto) {
	// Completar
}

// BuscarContacto devuelve un contacto por su nombre.
// Devuelve error si no existe.
func (g *GuiaTelefonica) BuscarContacto(nombre string) (Contacto, error) {
	// Completar
	var zero Contacto
	return zero, nil
}

// EliminarContacto elimina un contacto por su nombre.
// Devuelve error si no existe.
func (g *GuiaTelefonica) EliminarContacto(nombre string) error {
	// Completar
	return nil
}

// ExisteContacto devuelve true si el nombre está registrado.
func (g *GuiaTelefonica) ExisteContacto(nombre string) bool {
	// Completar
	return false
}

// ListarContactos devuelve todos los contactos en orden alfabético.
func (g *GuiaTelefonica) ListarContactos() []Contacto {
	// Completar
	return nil
}

// Cantidad devuelve la cantidad de contactos en la guía.
func (g *GuiaTelefonica) Cantidad() int {
	// Completar
	return 0
}
