package heap

import "testing"

func intCmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TestMinHeapInsertAndRemove(t *testing.T) {
	h := NewMinHeap[int](intCmp)
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	if h.Size() != 4 {
		t.Errorf("Size esperado 4, obtuve %d", h.Size())
	}

	val, _ := h.Remove()
	if val != 1 {
		t.Errorf("Remove esperaba 1, obtuve %d", val)
	}

	val, _ = h.Remove()
	if val != 3 {
		t.Errorf("Remove esperaba 3, obtuve %d", val)
	}

	val, _ = h.Remove()
	if val != 5 {
		t.Errorf("Remove esperaba 5, obtuve %d", val)
	}

	val, _ = h.Remove()
	if val != 8 {
		t.Errorf("Remove esperaba 8, obtuve %d", val)
	}

	if !h.IsEmpty() {
		t.Error("El heap debería estar vacío después de remover todos los elementos")
	}
}

func TestMaxHeapInsertAndRemove(t *testing.T) {
	h := NewMaxHeap[int](intCmp)
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	val, _ := h.Remove()
	if val != 8 {
		t.Errorf("Remove esperaba 8, obtuve %d", val)
	}

	val, _ = h.Remove()
	if val != 5 {
		t.Errorf("Remove esperaba 5, obtuve %d", val)
	}

	val, _ = h.Remove()
	if val != 3 {
		t.Errorf("Remove esperaba 3, obtuve %d", val)
	}

	val, _ = h.Remove()
	if val != 1 {
		t.Errorf("Remove esperaba 1, obtuve %d", val)
	}
}

func TestHeapRemoveEmpty(t *testing.T) {
	h := NewMinHeap[int](intCmp)
	_, err := h.Remove()
	if err == nil {
		t.Error("Remove sobre heap vacío debería retornar error")
	}
}

func TestHeapTop(t *testing.T) {
	h := NewMaxHeap[int](intCmp)
	h.Insert(10)
	h.Insert(5)

	val, _ := h.Top()
	if val != 10 {
		t.Errorf("Top esperaba 10, obtuve %d", val)
	}
	if h.Size() != 2 {
		t.Errorf("Top no debe remover, Size esperado 2, obtuve %d", h.Size())
	}
}

func TestHeapTopEmpty(t *testing.T) {
	h := NewMinHeap[int](intCmp)
	_, err := h.Top()
	if err == nil {
		t.Error("Top sobre heap vacío debería retornar error")
	}
}

func TestHeapIsEmpty(t *testing.T) {
	h := NewMinHeap[int](intCmp)
	if !h.IsEmpty() {
		t.Error("Heap nuevo debería estar vacío")
	}
	h.Insert(1)
	if h.IsEmpty() {
		t.Error("Heap con elementos no debería estar vacío")
	}
}

func TestHeapInsertDuplicates(t *testing.T) {
	h := NewMaxHeap[int](intCmp)
	h.Insert(5)
	h.Insert(5)
	h.Insert(3)

	if h.Size() != 3 {
		t.Errorf("Size esperado 3, obtuve %d", h.Size())
	}

	val, _ := h.Remove()
	if val != 5 {
		t.Errorf("Remove esperaba 5, obtuve %d", val)
	}
	val, _ = h.Remove()
	if val != 5 {
		t.Errorf("Remove esperaba 5, obtuve %d", val)
	}
	val, _ = h.Remove()
	if val != 3 {
		t.Errorf("Remove esperaba 3, obtuve %d", val)
	}
}

func TestHeapSingleElement(t *testing.T) {
	h := NewMinHeap[int](intCmp)
	h.Insert(42)

	if h.Size() != 1 {
		t.Errorf("Size esperado 1, obtuve %d", h.Size())
	}

	val, _ := h.Top()
	if val != 42 {
		t.Errorf("Top esperaba 42, obtuve %d", val)
	}

	val, _ = h.Remove()
	if val != 42 {
		t.Errorf("Remove esperaba 42, obtuve %d", val)
	}

	if !h.IsEmpty() {
		t.Error("Heap debería estar vacío")
	}
}

func TestMinHeapWithStructs(t *testing.T) {
	type persona struct {
		nombre string
		edad   int
	}

	cmp := func(a, b persona) int {
		if a.edad < b.edad {
			return -1
		}
		if a.edad > b.edad {
			return 1
		}
		return 0
	}

	h := NewMinHeap[persona](cmp)
	h.Insert(persona{"Ana", 30})
	h.Insert(persona{"Juan", 25})
	h.Insert(persona{"Luis", 35})

	p, _ := h.Remove()
	if p.nombre != "Juan" || p.edad != 25 {
		t.Errorf("Esperaba Juan/25, obtuve %s/%d", p.nombre, p.edad)
	}
}
