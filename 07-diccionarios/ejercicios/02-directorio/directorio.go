package directorio

import "github.com/untref-ayp2/data-structures/dictionary"

// Contacto representa los datos de una persona en el directorio.
type Contacto struct {
	Nombre   string
	Telefono string
}

// Directorio permite buscar contactos por nombre y por teléfono.
type Directorio struct {
	porNombre   dictionary.Dictionary[string, Contacto]
	porTelefono dictionary.Dictionary[string, string]
}

// NuevoDirectorio crea un directorio vacío.
func NuevoDirectorio() *Directorio {
	// Completar
	return nil
}

// AgregarContacto agrega un contacto. Error si el nombre ya existe o si el
// teléfono ya está asociado a otro contacto.
func (d *Directorio) AgregarContacto(nombre, telefono string) error {
	// Completar
	return nil
}

// BuscarPorNombre devuelve el contacto por su nombre. Error si no existe.
func (d *Directorio) BuscarPorNombre(nombre string) (Contacto, error) {
	// Completar
	return Contacto{}, nil
}

// BuscarPorTelefono devuelve el contacto por su número. Error si no existe.
func (d *Directorio) BuscarPorTelefono(telefono string) (Contacto, error) {
	// Completar
	return Contacto{}, nil
}

// ActualizarTelefono cambia el teléfono de un contacto.
// Error si el nombre no existe o el nuevo teléfono ya está en uso.
func (d *Directorio) ActualizarTelefono(nombre string, nuevoTelefono string) error {
	// Completar
	return nil
}

// ContactosConNombreLargo devuelve los contactos cuyo nombre tiene más de n
// caracteres.
func (d *Directorio) ContactosConNombreLargo(n int) []Contacto {
	// Completar
	return nil
}
