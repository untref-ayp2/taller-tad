package guia_telefonica

import (
	"slices"
	"testing"
)

func TestNuevaGuiaVacia(t *testing.T) {
	g := NuevaGuiaTelefonica()
	if g.Cantidad() != 0 {
		t.Error("Guía nueva debería tener cantidad 0")
	}
}

func TestAgregarYBuscar(t *testing.T) {
	g := NuevaGuiaTelefonica()
	g.AgregarContacto(Contacto{Nombre: "Juan", Telefono: "1234", Email: "juan@test.com"})

	contacto, err := g.BuscarContacto("Juan")
	if err != nil {
		t.Fatalf("Error inesperado: %v", err)
	}
	if contacto.Telefono != "1234" {
		t.Errorf("Teléfono esperado: 1234, obtenido: %s", contacto.Telefono)
	}
}

func TestBuscarInexistente(t *testing.T) {
	g := NuevaGuiaTelefonica()
	_, err := g.BuscarContacto("Inexistente")
	if err == nil {
		t.Error("Debería devolver error para contacto inexistente")
	}
}

func TestActualizarContacto(t *testing.T) {
	g := NuevaGuiaTelefonica()
	g.AgregarContacto(Contacto{Nombre: "Ana", Telefono: "1111"})
	g.AgregarContacto(Contacto{Nombre: "Ana", Telefono: "2222"})

	contacto, _ := g.BuscarContacto("Ana")
	if contacto.Telefono != "2222" {
		t.Errorf("Teléfono esperado: 2222, obtenido: %s", contacto.Telefono)
	}
	if g.Cantidad() != 1 {
		t.Errorf("Cantidad esperada: 1, obtenida: %d", g.Cantidad())
	}
}

func TestEliminarContacto(t *testing.T) {
	g := NuevaGuiaTelefonica()
	g.AgregarContacto(Contacto{Nombre: "Pedro", Telefono: "3333"})
	err := g.EliminarContacto("Pedro")
	if err != nil {
		t.Fatalf("Error inesperado: %v", err)
	}
	if g.ExisteContacto("Pedro") {
		t.Error("Pedro no debería existir después de eliminar")
	}
	if g.Cantidad() != 0 {
		t.Errorf("Cantidad esperada: 0, obtenida: %d", g.Cantidad())
	}
}

func TestEliminarInexistente(t *testing.T) {
	g := NuevaGuiaTelefonica()
	err := g.EliminarContacto("Inexistente")
	if err == nil {
		t.Error("Debería devolver error al eliminar contacto inexistente")
	}
}

func TestExisteContacto(t *testing.T) {
	g := NuevaGuiaTelefonica()
	g.AgregarContacto(Contacto{Nombre: "Luis"})
	if !g.ExisteContacto("Luis") {
		t.Error("Debería existir Luis")
	}
	if g.ExisteContacto("Maria") {
		t.Error("No debería existir Maria")
	}
}

func TestListarOrdenado(t *testing.T) {
	g := NuevaGuiaTelefonica()
	g.AgregarContacto(Contacto{Nombre: "Zara"})
	g.AgregarContacto(Contacto{Nombre: "Ana"})
	g.AgregarContacto(Contacto{Nombre: "Carlos"})
	g.AgregarContacto(Contacto{Nombre: "Beatriz"})

	contactos := g.ListarContactos()
	nombres := make([]string, len(contactos))
	for i, c := range contactos {
		nombres[i] = c.Nombre
	}

	esperado := []string{"Ana", "Beatriz", "Carlos", "Zara"}
	if !slices.Equal(nombres, esperado) {
		t.Errorf("Orden esperado: %v, obtenido: %v", esperado, nombres)
	}
}

func TestListarVacio(t *testing.T) {
	g := NuevaGuiaTelefonica()
	contactos := g.ListarContactos()
	if len(contactos) != 0 {
		t.Error("Guía vacía debería devolver slice vacío")
	}
}
