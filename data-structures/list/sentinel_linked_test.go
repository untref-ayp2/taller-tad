package list

import "testing"

func TestSentinelLinkedListNew(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if l.Size() != 0 {
		t.Errorf("Size(): esperaba 0, obtuvo %d", l.Size())
	}
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true para lista vacía")
	}
	_, ok := l.Head()
	if ok {
		t.Error("Head(): esperaba ok = false para lista vacía")
	}
	_, ok = l.Tail()
	if ok {
		t.Error("Tail(): esperaba ok = false para lista vacía")
	}
}

func TestSentinelLinkedListPrepend(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Prepend(2)
	l.Prepend(1)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

func TestSentinelLinkedListAppend(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	l.Append(2)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

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

func TestSentinelLinkedListRemoveFirst(t *testing.T) {
	l := NewSentinelLinkedList[int]()
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

func TestSentinelLinkedListRemoveFirstEmpty(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if l.RemoveFirst() {
		t.Error("RemoveFirst(): esperaba false en lista vacía")
	}
}

func TestSentinelLinkedListRemoveLast(t *testing.T) {
	l := NewSentinelLinkedList[int]()
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

func TestSentinelLinkedListRemoveLastSingleElement(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	l.RemoveLast()
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true después de remover único elemento")
	}
}

func TestSentinelLinkedListRemoveLastEmpty(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if l.RemoveLast() {
		t.Error("RemoveLast(): esperaba false en lista vacía")
	}
}

func TestSentinelLinkedListRemove(t *testing.T) {
	l := NewSentinelLinkedList[int]()
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

func TestSentinelLinkedListRemoveHead(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 2, 3)
}

func TestSentinelLinkedListRemoveTail(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Remove(3)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

func TestSentinelLinkedListRemoveNonExistent(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	l.Append(2)
	if l.Remove(99) {
		t.Error("Remove(99): esperaba false para valor inexistente")
	}
	if l.Size() != 2 {
		t.Errorf("Size(): no debería cambiar, esperaba 2, obtuvo %d", l.Size())
	}
}

func TestSentinelLinkedListContains(t *testing.T) {
	l := NewSentinelLinkedList[int]()
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

func TestSentinelLinkedListContainsEmpty(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if l.Contains(1) {
		t.Error("Contains(1): esperaba false en lista vacía")
	}
}

func TestSentinelLinkedListValues(t *testing.T) {
	l := NewSentinelLinkedList[int]()
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

func TestSentinelLinkedListValuesEmpty(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	vals := l.Values()
	if len(vals) != 0 {
		t.Errorf("Values(): esperaba 0 elementos, obtuvo %d", len(vals))
	}
}

func TestSentinelLinkedListClear(t *testing.T) {
	l := NewSentinelLinkedList[int]()
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

func TestSentinelLinkedListInsertAfter(t *testing.T) {
	l := NewSentinelLinkedList[int]()
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

func TestSentinelLinkedListInsertAfterTail(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(1)
	if !l.InsertAfter(1, 2) {
		t.Error("InsertAfter(1, 2): esperaba true")
	}
	checkHeadTail(t, l, 1, 2)
}

func TestSentinelLinkedListInsertAfterEmpty(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if l.InsertAfter(1, 2) {
		t.Error("InsertAfter(1, 2): esperaba false en lista vacía")
	}
}

func TestSentinelLinkedListInsertBefore(t *testing.T) {
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

func TestSentinelLinkedListInsertBeforeHead(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	l.Append(2)
	if !l.InsertBefore(2, 1) {
		t.Error("InsertBefore(2, 1): esperaba true")
	}
	checkHeadTail(t, l, 1, 2)
}

func TestSentinelLinkedListInsertBeforeEmpty(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if l.InsertBefore(1, 2) {
		t.Error("InsertBefore(1, 2): esperaba false en lista vacía")
	}
}

func TestSentinelLinkedListIsEmpty(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true para lista vacía")
	}
	l.Append(1)
	if l.IsEmpty() {
		t.Error("IsEmpty(): esperaba false para lista no vacía")
	}
}

func TestSentinelLinkedListString(t *testing.T) {
	l := NewSentinelLinkedList[int]()
	if s := l.String(); s != "[]" {
		t.Errorf("String(): esperaba '[]', obtuvo '%s'", s)
	}
	l.Append(1)
	l.Append(2)
	if s := l.String(); s != "[1 2]" {
		t.Errorf("String(): esperaba '[1 2]', obtuvo '%s'", s)
	}
}
