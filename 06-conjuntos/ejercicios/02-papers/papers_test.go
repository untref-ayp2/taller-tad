package papers

import (
	"testing"

	"github.com/untref-ayp2/data-structures/set"
)

func newSet(elems ...string) set.Set[string] {
	s := set.NewMapSet[string]()
	for _, e := range elems {
		s.Add(e)
	}
	return s
}

func TestBusquedaExacta(t *testing.T) {
	papers := map[string]set.Set[string]{
		"Go en la práctica":   newSet("go", "programación", "concurrencia"),
		"Estructuras de datos": newSet("programación", "algoritmos", "go"),
		"Machine Learning":     newSet("ia", "datos", "estadística"),
	}

	resultado := BusquedaExacta(papers, newSet("go", "programación"))
	if len(resultado) != 2 {
		t.Fatalf("BusquedaExacta() = %v, longitud %d, esperaba 2", resultado, len(resultado))
	}
}

func TestBusquedaExactaSinResultados(t *testing.T) {
	papers := map[string]set.Set[string]{
		"Go en la práctica": newSet("go", "programación"),
	}

	resultado := BusquedaExacta(papers, newSet("python"))
	if len(resultado) != 0 {
		t.Fatalf("BusquedaExacta() = %v, longitud %d, esperaba 0", resultado, len(resultado))
	}
}

func TestBusquedaRelajada(t *testing.T) {
	papers := map[string]set.Set[string]{
		"Go en la práctica":   newSet("go", "programación", "concurrencia"),
		"Machine Learning":     newSet("ia", "datos", "estadística"),
		"Redes Neuronales":     newSet("ia", "deep-learning", "datos"),
	}

	resultado := BusquedaRelajada(papers, newSet("ia", "datos"))
	if len(resultado) != 2 {
		t.Fatalf("BusquedaRelajada() = %v, longitud %d, esperaba 2", resultado, len(resultado))
	}
	// Machine Learning (2 coincidencias) debe ir antes que Redes Neuronales (1)
	if resultado[0] != "Machine Learning" {
		t.Errorf("BusquedaRelajada()[0] = %s, esperaba Machine Learning", resultado[0])
	}
	if resultado[1] != "Redes Neuronales" {
		t.Errorf("BusquedaRelajada()[1] = %s, esperaba Redes Neuronales", resultado[1])
	}
}

func TestSugerirKeywords(t *testing.T) {
	papers := map[string]set.Set[string]{
		"Go en la práctica":   newSet("go", "programación", "concurrencia"),
		"Estructuras de datos": newSet("programación", "algoritmos", "go"),
		"Machine Learning":     newSet("ia", "datos", "estadística"),
	}

	sugerencias := SugerirKeywords(papers, newSet("go"))
	if len(sugerencias) == 0 {
		t.Fatal("SugerirKeywords() no debería estar vacío")
	}
	// Debería sugerir: "algoritmos", "concurrencia", "programación"
	for _, s := range []string{"programación", "concurrencia", "algoritmos"} {
		encontrado := false
		for _, sug := range sugerencias {
			if sug == s {
				encontrado = true
				break
			}
		}
		if !encontrado {
			t.Errorf("SugerirKeywords() debería incluir %s, obtuvo %v", s, sugerencias)
		}
	}
}
