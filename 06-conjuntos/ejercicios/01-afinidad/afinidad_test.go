package afinidad

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

func TestRankingAfinidadBasico(t *testing.T) {
	intereses := map[string]set.Set[string]{
		"alice": newSet("música", "cine", "viajes", "lectura"),
		"bob":   newSet("música", "deportes", "cine"),
		"carol": newSet("viajes", "lectura", "cocina"),
		"dave":  newSet("deportes", "juegos"),
	}

	resultado := RankingAfinidad(intereses, "alice")

	esperado := []string{"bob", "carol", "dave"}
	if len(resultado) != 3 {
		t.Fatalf("RankingAfinidad() = %v, longitud %d, esperaba 3", resultado, len(resultado))
	}
	for i, r := range resultado {
		if r != esperado[i] {
			t.Errorf("RankingAfinidad()[%d] = %s, esperaba %s", i, r, esperado[i])
		}
	}
}

func TestRankingAfinidadEmpate(t *testing.T) {
	intereses := map[string]set.Set[string]{
		"alice": newSet("a", "b"),
		"bob":   newSet("a", "c"),
		"zara":  newSet("a", "c"),
	}

	resultado := RankingAfinidad(intereses, "alice")

	// bob y zara comparten 1 interés cada uno, deben salir ordenados alfabéticamente
	esperado := []string{"bob", "zara"}
	if len(resultado) != 2 {
		t.Fatalf("RankingAfinidad() = %v, longitud %d, esperaba 2", resultado, len(resultado))
	}
	for i, r := range resultado {
		if r != esperado[i] {
			t.Errorf("RankingAfinidad()[%d] = %s, esperaba %s", i, r, esperado[i])
		}
	}
}

func TestRankingAfinidadSinCoincidencias(t *testing.T) {
	intereses := map[string]set.Set[string]{
		"alice": newSet("a", "b"),
		"bob":   newSet("c", "d"),
	}

	resultado := RankingAfinidad(intereses, "alice")
	if len(resultado) != 1 {
		t.Fatalf("RankingAfinidad() = %v, longitud %d, esperaba 1", resultado, len(resultado))
	}
	if resultado[0] != "bob" {
		t.Errorf("RankingAfinidad()[0] = %s, esperaba bob", resultado[0])
	}
}
