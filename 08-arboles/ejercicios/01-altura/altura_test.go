package altura

import "testing"

func TestArbolBalanceado(t *testing.T) {
	h := ArbolBalanceado()
	if h != 3 {
		t.Errorf("Altura esperada: 3, obtenida: %d", h)
	}
}

func TestArbolDegenerado(t *testing.T) {
	h := ArbolDegenerado()
	if h != 5 {
		t.Errorf("Altura esperada: 5, obtenida: %d", h)
	}
}
