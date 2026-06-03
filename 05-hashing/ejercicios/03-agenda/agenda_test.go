package agenda

import (
	"testing"
)

func TestAgregarYBuscarContacto(t *testing.T) {
	a := NuevaAgenda()
	a.AgregarContacto(Contacto{
		Nombre:   "Juan",
		Telefono: "1234-5678",
		Email:    "juan@example.com",
	})

	contacto, err := a.BuscarContacto("Juan")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if contacto.Telefono != "1234-5678" {
		t.Errorf("Teléfono = '%s', esperaba '1234-5678'", contacto.Telefono)
	}
	if contacto.Email != "juan@example.com" {
		t.Errorf("Email = '%s', esperaba 'juan@example.com'", contacto.Email)
	}
}

func TestBuscarContactoInexistente(t *testing.T) {
	a := NuevaAgenda()
	_, err := a.BuscarContacto("Inexistente")
	if err == nil {
		t.Error("esperaba error al buscar contacto inexistente")
	}
}

func TestEliminarContacto(t *testing.T) {
	a := NuevaAgenda()
	a.AgregarContacto(Contacto{Nombre: "Juan", Telefono: "1234"})

	err := a.EliminarContacto("Juan")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	if a.ExisteContacto("Juan") {
		t.Error("ExisteContacto('Juan') debería ser false después de eliminar")
	}
}

func TestEliminarContactoInexistente(t *testing.T) {
	a := NuevaAgenda()
	err := a.EliminarContacto("Inexistente")
	if err == nil {
		t.Error("esperaba error al eliminar contacto inexistente")
	}
}

func TestActualizarContacto(t *testing.T) {
	a := NuevaAgenda()
	a.AgregarContacto(Contacto{Nombre: "Juan", Telefono: "1234", Email: "juan@viejo.com"})
	a.AgregarContacto(Contacto{Nombre: "Juan", Telefono: "5678", Email: "juan@nuevo.com"})

	contacto, err := a.BuscarContacto("Juan")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if contacto.Telefono != "5678" {
		t.Errorf("Teléfono = '%s', esperaba '5678'", contacto.Telefono)
	}
	if contacto.Email != "juan@nuevo.com" {
		t.Errorf("Email = '%s', esperaba 'juan@nuevo.com'", contacto.Email)
	}
}

func TestExisteContacto(t *testing.T) {
	a := NuevaAgenda()
	a.AgregarContacto(Contacto{Nombre: "Ana", Telefono: "0000"})

	if !a.ExisteContacto("Ana") {
		t.Error("ExisteContacto('Ana') debería ser true")
	}
	if a.ExisteContacto("Pedro") {
		t.Error("ExisteContacto('Pedro') debería ser false")
	}
}

func TestListarContactos(t *testing.T) {
	a := NuevaAgenda()
	a.AgregarContacto(Contacto{Nombre: "A", Telefono: "1"})
	a.AgregarContacto(Contacto{Nombre: "B", Telefono: "2"})
	a.AgregarContacto(Contacto{Nombre: "C", Telefono: "3"})

	contactos := a.ListarContactos()
	if len(contactos) != 3 {
		t.Fatalf("ListarContactos() = %v, longitud %d, esperaba 3", contactos, len(contactos))
	}
}

func TestCantidad(t *testing.T) {
	a := NuevaAgenda()
	if a.Cantidad() != 0 {
		t.Error("esperaba 0 en agenda recién creada")
	}

	a.AgregarContacto(Contacto{Nombre: "X", Telefono: "0"})
	if a.Cantidad() != 1 {
		t.Error("esperaba 1 después de agregar un contacto")
	}

	a.EliminarContacto("X")
	if a.Cantidad() != 0 {
		t.Error("esperaba 0 después de eliminar el único contacto")
	}
}
