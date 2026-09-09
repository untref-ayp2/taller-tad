package ejercicios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/list"
)

func TestMergeListas(t *testing.T) {
	cargar := func(vals []int) list.List[int] {
		l := list.NewSinglyLinkedList[int]()
		for _, v := range vals {
			l.Append(v)
		}
		return l
	}

	casos := []struct {
		l1, l2, esperado []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{2}, []int{2}},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, c := range casos {
		l1 := cargar(c.l1)
		l2 := cargar(c.l2)
		resultado := MergeListas[int](l1, l2)

		vals := resultado.Values()
		if len(vals) != len(c.esperado) {
			t.Fatalf("MergeListas(%v, %v) = %v, esperaba %v", c.l1, c.l2, vals, c.esperado)
		}
		for i := range vals {
			if vals[i] != c.esperado[i] {
				t.Errorf("MergeListas(%v, %v) = %v, esperaba %v", c.l1, c.l2, vals, c.esperado)
				break
			}
		}
	}
}
