package directorios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/set"
)

func orderedSet(elems ...string) set.Set[string] {
	less := func(a, b string) bool { return a < b }
	s := set.NewOrderedSet(less)
	for _, e := range elems {
		s.Add(e)
	}
	return s
}

func values(s set.Set[string]) []string {
	return s.Values()
}

func TestFusionarDirectorios(t *testing.T) {
	a := orderedSet("Ana", "Pedro", "Luis")
	b := orderedSet("Luis", "Maria", "Juan")

	unificada := FusionarDirectorios(a, b)
	if unificada.Size() != 5 {
		t.Fatalf("FusionarDirectorios Size = %d, esperaba 5", unificada.Size())
	}

	vals := values(unificada)
	esperado := []string{"Ana", "Juan", "Luis", "Maria", "Pedro"}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("Values()[%d] = %s, esperaba %s", i, v, esperado[i])
		}
	}
}

func TestFusionarDirectoriosConRepetidos(t *testing.T) {
	a := orderedSet("Ana", "Pedro", "Luis")
	b := orderedSet("Pedro", "Luis", "Maria")

	unificada := FusionarDirectorios(a, b)
	if unificada.Size() != 4 {
		t.Fatalf("FusionarDirectorios Size = %d, esperaba 4 (sin repetidos)", unificada.Size())
	}
}

func TestEmpleadosEnComun(t *testing.T) {
	a := orderedSet("Ana", "Pedro", "Luis", "Maria")
	b := orderedSet("Pedro", "Juan", "Maria")

	comun := EmpleadosEnComun(a, b)
	vals := values(comun)
	esperado := []string{"Maria", "Pedro"}
	if len(vals) != 2 {
		t.Fatalf("EmpleadosEnComun len = %d, esperaba 2: %v", len(vals), vals)
	}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("EmpleadosEnComun[%d] = %s, esperaba %s", i, v, esperado[i])
		}
	}
}

func TestEmpleadosExclusivos(t *testing.T) {
	a := orderedSet("Ana", "Pedro", "Luis")
	b := orderedSet("Pedro", "Maria")

	exclusivos := EmpleadosExclusivos(a, b)
	vals := values(exclusivos)
	esperado := []string{"Ana", "Luis", "Maria"}
	if len(vals) != 3 {
		t.Fatalf("EmpleadosExclusivos len = %d, esperaba 3: %v", len(vals), vals)
	}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("EmpleadosExclusivos[%d] = %s, esperaba %s", i, v, esperado[i])
		}
	}
}

func TestCoberturaCompleta(t *testing.T) {
	a := orderedSet("Ana", "Pedro", "Luis", "Maria")
	b := orderedSet("Ana", "Pedro")

	if !CoberturaCompleta(a, b) {
		t.Error("a debería cubrir completamente a b")
	}
	if CoberturaCompleta(b, a) {
		t.Error("b no debería cubrir completamente a a")
	}
}
