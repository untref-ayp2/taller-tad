package ejercicios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/list"
)

func TestInvertirLista(t *testing.T) {
	casos := []struct {
		entrada  []int
		esperado []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{5, 10, 15, 20}, []int{20, 15, 10, 5}},
	}

	for _, c := range casos {
		l := list.NewSinglyLinkedList[int]()
		for _, v := range c.entrada {
			l.Append(v)
		}
		resultado := InvertirLista[int](l)
		if len(resultado) != len(c.esperado) {
			t.Fatalf("InvertirLista(%v) = %v, esperaba %v", c.entrada, resultado, c.esperado)
		}
		for i := range resultado {
			if resultado[i] != c.esperado[i] {
				t.Errorf("InvertirLista(%v) = %v, esperaba %v", c.entrada, resultado, c.esperado)
				break
			}
		}
	}
}
