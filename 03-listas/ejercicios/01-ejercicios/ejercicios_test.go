package ejercicios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/list"
)

func TestPrependAll(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	PrependAll(l, []int{3, 2, 1})
	if l.Size() != 3 {
		t.Fatalf("Size(): esperaba 3, obtuvo %d", l.Size())
	}
	vals := l.Values()
	esperado := []int{1, 2, 3}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("Values()[%d]: esperaba %d, obtuvo %d", i, esperado[i], v)
		}
	}
}

func TestAppendAll(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	AppendAll(l, []int{1, 2, 3})
	if l.Size() != 3 {
		t.Fatalf("Size(): esperaba 3, obtuvo %d", l.Size())
	}
	vals := l.Values()
	esperado := []int{1, 2, 3}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("Values()[%d]: esperaba %d, obtuvo %d", i, esperado[i], v)
		}
	}
}

func TestListaInvertida(t *testing.T) {
	original := list.NewSinglyLinkedList[int]()
	original.Append(1)
	original.Append(2)
	original.Append(3)

	invertida := ListaInvertida(original)
	if invertida == nil || invertida.Size() != 3 {
		t.Fatalf("Size(): esperaba 3, obtuvo una lista incorrecta o nil")
	}
	vals := invertida.Values()
	esperado := []int{3, 2, 1}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("Values()[%d]: esperaba %d, obtuvo %d", i, esperado[i], v)
		}
	}

	// La original no debe modificarse
	origVals := original.Values()
	if len(origVals) != 3 || origVals[0] != 1 {
		t.Error("la lista original fue modificada")
	}
}

func TestRemoverTodos(t *testing.T) {
	l := list.NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(1)
	l.Append(3)
	l.Append(1)

	removidos := RemoverTodos(l, 1)
	if removidos != 3 {
		t.Errorf("RemoverTodos(): esperaba 3, obtuvo %d", removidos)
	}
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	vals := l.Values()
	esperado := []int{2, 3}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("Values()[%d]: esperaba %d, obtuvo %d", i, esperado[i], v)
		}
	}
}

func TestCombinarListas(t *testing.T) {
	a := list.NewSinglyLinkedList[int]()
	a.Append(1)
	a.Append(2)

	b := list.NewSinglyLinkedList[int]()
	b.Append(3)
	b.Append(4)

	c := CombinarListas(a, b)
	if c == nil || c.Size() != 4 {
		t.Fatalf("Size(): esperaba 4, obtuvo una lista incorrecta o nil")
	}
	vals := c.Values()
	esperado := []int{1, 2, 3, 4}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("Values()[%d]: esperaba %d, obtuvo %d", i, esperado[i], v)
		}
	}
}
