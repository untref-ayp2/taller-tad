package ejercicios

import "testing"

func TestJosephus(t *testing.T) {
	casos := []struct {
		n, k, esperado int
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 2, 3},
		{7, 3, 4},
		{5, 2, 3},
		{10, 3, 4},
		{41, 3, 31}, // caso histórico
	}

	for _, c := range casos {
		resultado := Josephus(c.n, c.k)
		if resultado != c.esperado {
			t.Errorf("Josephus(%d, %d) = %d, esperaba %d", c.n, c.k, resultado, c.esperado)
		}
	}
}
