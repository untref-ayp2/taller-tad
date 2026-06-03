package conteo

import (
	"testing"
)

func TestAgregarTextoYFrecuencia(t *testing.T) {
	c := NuevoContador()
	c.AgregarTexto("hola mundo hola")

	if c.Frecuencia("hola") != 2 {
		t.Errorf("Frecuencia('hola') = %d, esperaba 2", c.Frecuencia("hola"))
	}
	if c.Frecuencia("mundo") != 1 {
		t.Errorf("Frecuencia('mundo') = %d, esperaba 1", c.Frecuencia("mundo"))
	}
}

func TestPalabraInexistente(t *testing.T) {
	c := NuevoContador()
	if c.Frecuencia("foo") != 0 {
		t.Error("Frecuencia('foo') debería ser 0 para palabra no agregada")
	}
}

func TestTotalPalabras(t *testing.T) {
	c := NuevoContador()
	c.AgregarTexto("a b c a b a")

	if c.TotalPalabras() != 6 {
		t.Errorf("TotalPalabras() = %d, esperaba 6", c.TotalPalabras())
	}
}

func TestPalabrasUnicas(t *testing.T) {
	c := NuevoContador()
	c.AgregarTexto("a b c a b a")

	if c.PalabrasUnicas() != 3 {
		t.Errorf("PalabrasUnicas() = %d, esperaba 3", c.PalabrasUnicas())
	}
}

func TestTopN(t *testing.T) {
	c := NuevoContador()
	c.AgregarTexto("a b c a b a")

	top := c.TopN(2)
	if len(top) != 2 {
		t.Fatalf("TopN(2) = %v, longitud %d, esperaba 2", top, len(top))
	}
	if top[0].Palabra != "a" || top[0].Frecuencia != 3 {
		t.Errorf("TopN[0] = {%s, %d}, esperaba {a, 3}", top[0].Palabra, top[0].Frecuencia)
	}
}

func TestCaseInsensitive(t *testing.T) {
	c := NuevoContador()
	c.AgregarTexto("Hola hola HOLa")

	if c.Frecuencia("hola") != 3 {
		t.Errorf("Frecuencia('hola') = %d, esperaba 3 (case insensitive)", c.Frecuencia("hola"))
	}
}

func TestListarPalabras(t *testing.T) {
	c := NuevoContador()
	c.AgregarTexto("a b c")

	palabras := c.ListarPalabras()
	if len(palabras) != 3 {
		t.Fatalf("ListarPalabras() = %v, longitud %d, esperaba 3", palabras, len(palabras))
	}
}
