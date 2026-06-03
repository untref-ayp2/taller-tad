package diccionario

import (
	"testing"
)

func TestAgregarYTraducir(t *testing.T) {
	d := NuevoDiccionario()
	d.Agregar("hola", "hello")
	d.Agregar("mundo", "world")

	if d.Traducir("hola") != "hello" {
		t.Error("Traducir('hola') debería devolver 'hello'")
	}
	if d.Traducir("mundo") != "world" {
		t.Error("Traducir('mundo') debería devolver 'world'")
	}
}

func TestTraducirPalabraInexistente(t *testing.T) {
	d := NuevoDiccionario()
	d.Agregar("hola", "hello")

	if d.Traducir("adiós") != "adiós" {
		t.Error("Traducir('adiós') debería devolver 'adiós' (palabra original)")
	}
}

func TestActualizarTraduccion(t *testing.T) {
	d := NuevoDiccionario()
	d.Agregar("rojo", "red")
	d.Agregar("rojo", "rouge")

	if d.Traducir("rojo") != "rouge" {
		t.Error("Traducir('rojo') debería devolver 'rouge' (última actualización)")
	}
}

func TestTraducirOracion(t *testing.T) {
	d := NuevoDiccionario()
	d.Agregar("hola", "hello")
	d.Agregar("mundo", "world")

	resultado := d.TraducirOracion("hola mundo")
	if resultado != "hello world" {
		t.Errorf("TraducirOracion() = '%s', esperaba 'hello world'", resultado)
	}
}

func TestTraducirOracionConPalabraFaltante(t *testing.T) {
	d := NuevoDiccionario()
	d.Agregar("hola", "hello")

	resultado := d.TraducirOracion("hola amigo")
	if resultado != "hello amigo" {
		t.Errorf("TraducirOracion() = '%s', esperaba 'hello amigo'", resultado)
	}
}

func TestListarPalabras(t *testing.T) {
	d := NuevoDiccionario()
	d.Agregar("hola", "hello")
	d.Agregar("mundo", "world")
	d.Agregar("sol", "sun")

	palabras := d.ListarPalabras()
	if len(palabras) != 3 {
		t.Fatalf("ListarPalabras() = %v, longitud %d, esperaba 3", palabras, len(palabras))
	}
	contiene := func(s string) bool {
		for _, p := range palabras {
			if p == s {
				return true
			}
		}
		return false
	}
	if !contiene("hola") || !contiene("mundo") || !contiene("sol") {
		t.Errorf("ListarPalabras() no contiene todas las palabras esperadas: %v", palabras)
	}
}

func TestCantidad(t *testing.T) {
	d := NuevoDiccionario()
	if d.Cantidad() != 0 {
		t.Error("esperaba 0 en diccionario recién creado")
	}

	d.Agregar("a", "b")
	if d.Cantidad() != 1 {
		t.Error("esperaba 1 después de agregar una palabra")
	}
}
