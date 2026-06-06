package directorio

import (
	"testing"
)

func TestNuevoDirectorio(t *testing.T) {
	d := NuevoDirectorio()
	if d == nil {
		t.Fatal("NuevoDirectorio no debería devolver nil")
	}
}

func TestAgregarYBuscarPorNombre(t *testing.T) {
	d := NuevoDirectorio()
	err := d.AgregarContacto("Juan", "1234-5678")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	c, err := d.BuscarPorNombre("Juan")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if c.Nombre != "Juan" || c.Telefono != "1234-5678" {
		t.Errorf("esperaba (Juan, 1234-5678), obtuve (%s, %s)", c.Nombre, c.Telefono)
	}
}

func TestBuscarPorNombreInexistente(t *testing.T) {
	d := NuevoDirectorio()
	_, err := d.BuscarPorNombre("Inexistente")
	if err == nil {
		t.Error("esperaba error al buscar nombre inexistente")
	}
}

func TestBuscarPorTelefono(t *testing.T) {
	d := NuevoDirectorio()
	d.AgregarContacto("Maria", "9876-5432")

	c, err := d.BuscarPorTelefono("9876-5432")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if c.Nombre != "Maria" {
		t.Errorf("esperaba Maria, obtuve %s", c.Nombre)
	}
}

func TestBuscarPorTelefonoInexistente(t *testing.T) {
	d := NuevoDirectorio()
	_, err := d.BuscarPorTelefono("0000-0000")
	if err == nil {
		t.Error("esperaba error al buscar teléfono inexistente")
	}
}

func TestAgregarContactoNombreDuplicado(t *testing.T) {
	d := NuevoDirectorio()
	d.AgregarContacto("Juan", "1234-5678")
	err := d.AgregarContacto("Juan", "9999-9999")
	if err == nil {
		t.Error("esperaba error por nombre duplicado")
	}
}

func TestAgregarContactoTelefonoDuplicado(t *testing.T) {
	d := NuevoDirectorio()
	d.AgregarContacto("Juan", "1234-5678")
	err := d.AgregarContacto("Maria", "1234-5678")
	if err == nil {
		t.Error("esperaba error por teléfono duplicado")
	}
}

func TestActualizarTelefono(t *testing.T) {
	d := NuevoDirectorio()
	d.AgregarContacto("Juan", "1234-5678")

	err := d.ActualizarTelefono("Juan", "1111-1111")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	c, _ := d.BuscarPorNombre("Juan")
	if c.Telefono != "1111-1111" {
		t.Errorf("esperaba 1111-1111, obtuve %s", c.Telefono)
	}

	// El teléfono viejo ya no debe aparecer en búsqueda inversa
	_, err = d.BuscarPorTelefono("1234-5678")
	if err == nil {
		t.Error("el teléfono viejo no debería seguir siendo buscable")
	}

	// El teléfono nuevo debe aparecer
	_, err = d.BuscarPorTelefono("1111-1111")
	if err != nil {
		t.Errorf("el teléfono nuevo debería ser buscable: %v", err)
	}
}

func TestActualizarTelefonoInexistente(t *testing.T) {
	d := NuevoDirectorio()
	err := d.ActualizarTelefono("Inexistente", "1111-1111")
	if err == nil {
		t.Error("esperaba error al actualizar nombre inexistente")
	}
}

func TestActualizarTelefonoYaEnUso(t *testing.T) {
	d := NuevoDirectorio()
	d.AgregarContacto("Juan", "1234-5678")
	d.AgregarContacto("Maria", "9876-5432")

	err := d.ActualizarTelefono("Juan", "9876-5432")
	if err == nil {
		t.Error("esperaba error al asignar teléfono ya en uso")
	}

	// El teléfono de Juan no debe haber cambiado
	c, _ := d.BuscarPorNombre("Juan")
	if c.Telefono != "1234-5678" {
		t.Errorf("el teléfono de Juan no debería cambiar, obtuve %s", c.Telefono)
	}
}

func TestContactosConNombreLargo(t *testing.T) {
	d := NuevoDirectorio()
	d.AgregarContacto("Juan", "1111-1111")
	d.AgregarContacto("Maria", "2222-2222")
	d.AgregarContacto("Josefina", "3333-3333")

	resultado := d.ContactosConNombreLargo(4)
	if len(resultado) != 2 {
		t.Fatalf("esperaba 2 contactos con nombre >4, obtuve %d", len(resultado))
	}

	nombres := make(map[string]bool)
	for _, c := range resultado {
		nombres[c.Nombre] = true
	}
	if !nombres["Maria"] || !nombres["Josefina"] {
		t.Errorf("resultado inesperado: %v", resultado)
	}
}

func TestContactosConNombreLargoVacio(t *testing.T) {
	d := NuevoDirectorio()
	d.AgregarContacto("Juan", "1111-1111")

	resultado := d.ContactosConNombreLargo(10)
	if len(resultado) != 0 {
		t.Errorf("esperaba lista vacía, obtuve %v", resultado)
	}
}
