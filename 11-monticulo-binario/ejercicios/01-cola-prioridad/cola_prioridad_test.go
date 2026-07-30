package colaprioridad

import "testing"

func TestNuevaColaPrioridad(t *testing.T) {
	cp := NuevaColaPrioridad()
	if !cp.EstaVacia() {
		t.Error("cola nueva debería estar vacía")
	}
	if cp.Cantidad() != 0 {
		t.Errorf("Cantidad() = %d, esperaba 0", cp.Cantidad())
	}
}

func TestAgregarAtender(t *testing.T) {
	cp := NuevaColaPrioridad()
	cp.Agregar(Persona{"Ana", 25})

	p, err := cp.Atender()
	if err != nil {
		t.Fatalf("Atender() error: %v", err)
	}
	if p.nombre != "Ana" || p.edad != 25 {
		t.Errorf("Atender() = %+v, esperaba {Ana 25}", p)
	}
	if !cp.EstaVacia() {
		t.Error("cola debería estar vacía después de atender")
	}
}

func TestPrioridadPorEdad(t *testing.T) {
	cp := NuevaColaPrioridad()
	cp.Agregar(Persona{"Ana", 25})
	cp.Agregar(Persona{"Bob", 30})
	cp.Agregar(Persona{"Eva", 20})

	p1, _ := cp.Atender()
	if p1.nombre != "Bob" {
		t.Errorf("primero debería ser Bob (30), fue %s", p1.nombre)
	}

	p2, _ := cp.Atender()
	if p2.nombre != "Ana" {
		t.Errorf("segundo debería ser Ana (25), fue %s", p2.nombre)
	}

	p3, _ := cp.Atender()
	if p3.nombre != "Eva" {
		t.Errorf("tercero debería ser Eva (20), fue %s", p3.nombre)
	}
}

func TestOrdenDeLlegadaMismaEdad(t *testing.T) {
	cp := NuevaColaPrioridad()
	cp.Agregar(Persona{"Ana", 30})
	cp.Agregar(Persona{"Bob", 30})

	p1, _ := cp.Atender()
	if p1.nombre != "Ana" {
		t.Errorf("con misma edad, debería respetar orden: esperaba Ana, fue %s", p1.nombre)
	}
}

func TestSiguienteSinEliminar(t *testing.T) {
	cp := NuevaColaPrioridad()
	cp.Agregar(Persona{"Ana", 25})

	p, err := cp.Siguiente()
	if err != nil {
		t.Fatalf("Siguiente() error: %v", err)
	}
	if p.nombre != "Ana" {
		t.Errorf("Siguiente() = %s, esperaba Ana", p.nombre)
	}
	if cp.Cantidad() != 1 {
		t.Error("Siguiente() no debería eliminar el elemento")
	}
}

func TestAtenderColaVacia(t *testing.T) {
	cp := NuevaColaPrioridad()
	_, err := cp.Atender()
	if err == nil {
		t.Error("Atender() de cola vacía debería devolver error")
	}
}
