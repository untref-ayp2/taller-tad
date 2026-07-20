package conjunto_ordenado

import (
	"cmp"
	"slices"
	"testing"
)

func TestNuevoConjuntoVacio(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](cmp.Compare[int])
	if c.Cantidad() != 0 {
		t.Error("Conjunto nuevo debería tener cantidad 0")
	}
}

func TestAgregarYContiene(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](cmp.Compare[int])
	c.Agregar(5)
	c.Agregar(3)
	c.Agregar(7)

	if !c.Contiene(5) {
		t.Error("Debería contener 5")
	}
	if !c.Contiene(3) {
		t.Error("Debería contener 3")
	}
	if !c.Contiene(7) {
		t.Error("Debería contener 7")
	}
	if c.Contiene(9) {
		t.Error("No debería contener 9")
	}
}

func TestAgregarDuplicado(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](cmp.Compare[int])
	c.Agregar(5)
	c.Agregar(5)
	c.Agregar(5)

	if c.Cantidad() != 1 {
		t.Errorf("Cantidad esperada: 1, obtenida: %d", c.Cantidad())
	}
}

func TestEliminar(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](cmp.Compare[int])
	c.Agregar(5)
	c.Agregar(3)
	c.Agregar(7)

	c.Eliminar(3)
	if c.Contiene(3) {
		t.Error("3 no debería estar después de eliminar")
	}
	if c.Cantidad() != 2 {
		t.Errorf("Cantidad esperada: 2, obtenida: %d", c.Cantidad())
	}
}

func TestEliminarInexistente(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](cmp.Compare[int])
	c.Agregar(5)
	c.Eliminar(99)

	if c.Cantidad() != 1 {
		t.Errorf("Cantidad esperada: 1, obtenida: %d", c.Cantidad())
	}
}

func TestValoresOrdenados(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](cmp.Compare[int])
	valores := []int{7, 3, 9, 1, 5, 8, 10}
	for _, v := range valores {
		c.Agregar(v)
	}

	resultado := c.Valores()
	esperado := []int{1, 3, 5, 7, 8, 9, 10}
	if !slices.Equal(resultado, esperado) {
		t.Errorf("Valores esperados: %v, obtenidos: %v", esperado, resultado)
	}
}

func TestValoresStrings(t *testing.T) {
	c := NuevoConjuntoOrdenado[string](cmp.Compare[string])
	c.Agregar("zorro")
	c.Agregar("arbol")
	c.Agregar("casa")
	c.Agregar("bosque")

	resultado := c.Valores()
	esperado := []string{"arbol", "bosque", "casa", "zorro"}
	if !slices.Equal(resultado, esperado) {
		t.Errorf("Valores esperados: %v, obtenidos: %v", esperado, resultado)
	}
}

func TestConjuntoVacioValores(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](cmp.Compare[int])
	if r := c.Valores(); len(r) != 0 {
		t.Error("Conjunto vacío debería devolver slice vacío")
	}
}

func TestOrdenDescendente(t *testing.T) {
	c := NuevoConjuntoOrdenado[int](func(a, b int) int { return cmp.Compare(b, a) })
	c.Agregar(3)
	c.Agregar(1)
	c.Agregar(2)

	resultado := c.Valores()
	esperado := []int{3, 2, 1}
	if !slices.Equal(resultado, esperado) {
		t.Errorf("Valores esperados: %v, obtenidos: %v", esperado, resultado)
	}
}
