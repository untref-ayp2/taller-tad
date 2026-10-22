package altura

import "testing"

func TestArbolBalanceado(t *testing.T) {
	h := ArbolBalanceado()
	if h != 2 {
		t.Errorf("Altura esperada: 2, obtenida: %d", h)
	}
}

func TestArbolDegenerado(t *testing.T) {
	h := ArbolDegenerado()
	if h != 4 {
		t.Errorf("Altura esperada: 4, obtenida: %d", h)
	}
}
