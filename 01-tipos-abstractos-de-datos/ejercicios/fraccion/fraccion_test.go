package fraccion

import (
	"testing"
)

func TestNuevaFraccion(t *testing.T) {
	f, err := NuevaFraccion(1, 2)
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if f.String() != "1/2" {
		t.Errorf("esperado '1/2', got '%s'", f.String())
	}
}

func TestDenominadorCero(t *testing.T) {
	_, err := NuevaFraccion(1, 0)
	if err == nil {
		t.Error("se esperaba error con denominador 0")
	}
}

func TestSimplificar(t *testing.T) {
	f, _ := NuevaFraccion(4, 6)
	if f.String() != "2/3" {
		t.Errorf("esperado '2/3', got '%s'", f.String())
	}
}

func TestSumar(t *testing.T) {
	a, _ := NuevaFraccion(1, 4)
	b, _ := NuevaFraccion(1, 4)
	r := a.Sumar(b)
	if r.String() != "1/2" {
		t.Errorf("1/4 + 1/4 esperado '1/2', got '%s'", r.String())
	}
}

func TestRestar(t *testing.T) {
	a, _ := NuevaFraccion(1, 2)
	b, _ := NuevaFraccion(1, 4)
	r := a.Restar(b)
	if r.String() != "1/4" {
		t.Errorf("1/2 - 1/4 esperado '1/4', got '%s'", r.String())
	}
}

func TestMultiplicar(t *testing.T) {
	a, _ := NuevaFraccion(2, 3)
	b, _ := NuevaFraccion(3, 4)
	r := a.Multiplicar(b)
	if r.String() != "1/2" {
		t.Errorf("2/3 * 3/4 esperado '1/2', got '%s'", r.String())
	}
}

func TestDividir(t *testing.T) {
	a, _ := NuevaFraccion(1, 2)
	b, _ := NuevaFraccion(3, 4)
	r, err := a.Dividir(b)
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if r.String() != "2/3" {
		t.Errorf("1/2 / (3/4) esperado '2/3', got '%s'", r.String())
	}
}

func TestValor(t *testing.T) {
	f, _ := NuevaFraccion(1, 4)
	if f.Valor() != 0.25 {
		t.Errorf("valor esperado 0.25, got %f", f.Valor())
	}
}
