package list

import "testing"

// checkHeadTail verifica que Head() y Tail() devuelvan el valor esperado.
func checkHeadTail(t *testing.T, l List[int], wantHead, wantTail int) {
	t.Helper()
	head, ok := l.Head()
	if !ok {
		t.Error("Head(): esperaba ok = true, obtuvo false")
	} else if head != wantHead {
		t.Errorf("Head(): esperaba %d, obtuvo %d", wantHead, head)
	}
	tail, ok := l.Tail()
	if !ok {
		t.Error("Tail(): esperaba ok = true, obtuvo false")
	} else if tail != wantTail {
		t.Errorf("Tail(): esperaba %d, obtuvo %d", wantTail, tail)
	}
}

// ----- SinglyLinkedList -----

func TestSinglyLinkedListPrepend(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Prepend(2)
	l.Prepend(1)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

func TestSinglyLinkedListAppend(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

func TestSinglyLinkedListRemoveFirst(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	if !l.RemoveFirst() {
		t.Error("RemoveFirst(): esperaba true")
	}
	if l.Size() != 1 {
		t.Errorf("Size(): esperaba 1, obtuvo %d", l.Size())
	}
	head, ok := l.Head()
	if !ok || head != 2 {
		t.Errorf("Head(): esperaba 2, true; obtuvo %d, %v", head, ok)
	}
}

func TestSinglyLinkedListRemoveLast(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	if !l.RemoveLast() {
		t.Error("RemoveLast(): esperaba true")
	}
	if l.Size() != 1 {
		t.Errorf("Size(): esperaba 1, obtuvo %d", l.Size())
	}
	tail, ok := l.Tail()
	if !ok || tail != 1 {
		t.Errorf("Tail(): esperaba 1, true; obtuvo %d, %v", tail, ok)
	}
}

func TestSinglyLinkedListRemove(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	if !l.Remove(2) {
		t.Error("Remove(2): esperaba true")
	}
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 3)
}

func TestSinglyLinkedListContains(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(10)
	l.Append(20)
	if !l.Contains(10) {
		t.Error("Contains(10): esperaba true")
	}
	if !l.Contains(20) {
		t.Error("Contains(20): esperaba true")
	}
	if l.Contains(30) {
		t.Error("Contains(30): esperaba false")
	}
}

func TestSinglyLinkedListValues(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	vals := l.Values()
	if len(vals) != 3 {
		t.Fatalf("Values(): esperaba 3 elementos, obtuvo %d", len(vals))
	}
	for i, v := range vals {
		if v != i+1 {
			t.Errorf("Values()[%d]: esperaba %d, obtuvo %d", i, i+1, v)
		}
	}
}

func TestSinglyLinkedListClear(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Clear()
	if !l.IsEmpty() {
		t.Error("Clear(): la lista debería estar vacía")
	}
	if l.Size() != 0 {
		t.Errorf("Size(): esperaba 0, obtuvo %d", l.Size())
	}
}

func TestSinglyLinkedListInsertAfter(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(3)
	if !l.InsertAfter(1, 2) {
		t.Error("InsertAfter(1, 2): esperaba true")
	}
	vals := l.Values()
	if len(vals) != 3 || vals[0] != 1 || vals[1] != 2 || vals[2] != 3 {
		t.Errorf("Values(): esperaba [1 2 3], obtuvo %v", vals)
	}
}

func TestSinglyLinkedListInsertBefore(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	l.Append(1)
	l.Append(3)
	if !l.InsertBefore(3, 2) {
		t.Error("InsertBefore(3, 2): esperaba true")
	}
	vals := l.Values()
	if len(vals) != 3 || vals[0] != 1 || vals[1] != 2 || vals[2] != 3 {
		t.Errorf("Values(): esperaba [1 2 3], obtuvo %v", vals)
	}
}

func TestSinglyLinkedListIsEmpty(t *testing.T) {
	l := NewSinglyLinkedList[int]()
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true para lista vacía")
	}
	l.Append(1)
	if l.IsEmpty() {
		t.Error("IsEmpty(): esperaba false para lista no vacía")
	}
}

// ----- DoublyLinkedList -----

func TestDoublyLinkedListPrependAppend(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	l.Prepend(2)
	l.Prepend(1)
	l.Append(3)
	l.Append(4)
	checkHeadTail(t, l, 1, 4)
	if l.Size() != 4 {
		t.Errorf("Size(): esperaba 4, obtuvo %d", l.Size())
	}
}

func TestDoublyLinkedListRemoveLast(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	if !l.RemoveLast() {
		t.Error("RemoveLast(): esperaba true")
	}
	checkHeadTail(t, l, 1, 2)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
}

// ----- SentinelLinkedList -----

func TestSentinelPrependAppend(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Prepend(2)
	l.Prepend(1)
	l.Append(3)
	l.Append(4)
	checkHeadTail(t, l, 1, 4)
	if l.Size() != 4 {
		t.Errorf("Size(): esperaba 4, obtuvo %d", l.Size())
	}
}

func TestSentinelInsertBefore(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	l.Append(3)
	if !l.InsertBefore(3, 2) {
		t.Error("InsertBefore(3, 2): esperaba true")
	}
	vals := l.Values()
	if len(vals) != 3 || vals[0] != 1 || vals[1] != 2 || vals[2] != 3 {
		t.Errorf("Values(): esperaba [1 2 3], obtuvo %v", vals)
	}
}

func TestSentinelRemoveLast(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.RemoveLast()
	checkHeadTail(t, l, 1, 2)
}

func TestSentinelRemoveOnlyElement(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	if !l.RemoveFirst() {
		t.Error("RemoveFirst(): esperaba true")
	}
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true después de remover único elemento")
	}
}

// ----- CircularLinkedList -----

func TestCircularPrependAppend(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Prepend(2)
	l.Prepend(1)
	l.Append(3)
	l.Append(4)
	checkHeadTail(t, l, 1, 4)
	if l.Size() != 4 {
		t.Errorf("Size(): esperaba 4, obtuvo %d", l.Size())
	}
}

func TestCircularValues(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	vals := l.Values()
	if len(vals) != 3 {
		t.Fatalf("Values(): esperaba 3, obtuvo %d", len(vals))
	}
	for i, v := range vals {
		if v != i+1 {
			t.Errorf("Values()[%d]: esperaba %d, obtuvo %d", i, i+1, v)
		}
	}
}
