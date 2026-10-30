package list

import "testing"

func TestCircularLinkedListNew(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListPrepend(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Prepend(2)
	l.Prepend(1)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

func TestCircularLinkedListAppend(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

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

func TestCircularLinkedListRemoveFirst(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListRemoveFirstSingleElement(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.RemoveFirst()
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true después de remover único elemento")
	}
}

func TestCircularLinkedListRemoveFirstEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	if l.RemoveFirst() {
		t.Error("RemoveFirst(): esperaba false en lista vacía")
	}
}

func TestCircularLinkedListRemoveLast(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListRemoveLastSingleElement(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.RemoveLast()
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true después de remover único elemento")
	}
}

func TestCircularLinkedListRemoveLastEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	if l.RemoveLast() {
		t.Error("RemoveLast(): esperaba false en lista vacía")
	}
}

func TestCircularLinkedListRemove(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListRemoveHead(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 2, 3)
}

func TestCircularLinkedListRemoveTail(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Remove(3)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 1, 2)
}

func TestCircularLinkedListRemoveNonExistent(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	if l.Remove(99) {
		t.Error("Remove(99): esperaba false para valor inexistente")
	}
	if l.Size() != 2 {
		t.Errorf("Size(): no debería cambiar, esperaba 2, obtuvo %d", l.Size())
	}
}

func TestCircularLinkedListRemoveSingleElement(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	if !l.Remove(1) {
		t.Error("Remove(1): esperaba true")
	}
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true tras eliminar el único elemento")
	}
	if l.Size() != 0 {
		t.Errorf("Size(): esperaba 0, obtuvo %d", l.Size())
	}
	if _, ok := l.Head(); ok {
		t.Error("Head(): esperaba false en lista vacía")
	}
}

func TestCircularLinkedListRemoveSingleElementMismatch(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	if l.Remove(99) {
		t.Error("Remove(99): esperaba false si el único valor no coincide")
	}
	if l.Size() != 1 {
		t.Errorf("Size(): esperaba 1, obtuvo %d", l.Size())
	}
}

func TestCircularLinkedListContains(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListContainsEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	if l.Contains(1) {
		t.Error("Contains(1): esperaba false en lista vacía")
	}
}

func TestCircularLinkedListValues(t *testing.T) {
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

func TestCircularLinkedListValuesEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	vals := l.Values()
	if len(vals) != 0 {
		t.Errorf("Values(): esperaba 0 elementos, obtuvo %d", len(vals))
	}
}

func TestCircularLinkedListClear(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListInsertAfter(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListInsertAfterTail(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	if !l.InsertAfter(1, 2) {
		t.Error("InsertAfter(1, 2): esperaba true")
	}
	checkHeadTail(t, l, 1, 2)
}

func TestCircularLinkedListInsertAfterEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	if l.InsertAfter(1, 2) {
		t.Error("InsertAfter(1, 2): esperaba false en lista vacía")
	}
}

func TestCircularLinkedListInsertBefore(t *testing.T) {
	l := NewCircularLinkedList[int]()
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

func TestCircularLinkedListInsertBeforeHead(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(2)
	if !l.InsertBefore(2, 1) {
		t.Error("InsertBefore(2, 1): esperaba true")
	}
	checkHeadTail(t, l, 1, 2)
}

func TestCircularLinkedListInsertBeforeEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	if l.InsertBefore(1, 2) {
		t.Error("InsertBefore(1, 2): esperaba false en lista vacía")
	}
}

func TestCircularLinkedListIsEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	if !l.IsEmpty() {
		t.Error("IsEmpty(): esperaba true para lista vacía")
	}
	l.Append(1)
	if l.IsEmpty() {
		t.Error("IsEmpty(): esperaba false para lista no vacía")
	}
}

func TestCircularLinkedListString(t *testing.T) {
	l := NewCircularLinkedList[int]()
	if s := l.String(); s != "[]" {
		t.Errorf("String(): esperaba '[]', obtuvo '%s'", s)
	}
	l.Append(1)
	l.Append(2)
	if s := l.String(); s != "[1 2]" {
		t.Errorf("String(): esperaba '[1 2]', obtuvo '%s'", s)
	}
}

func TestCircularLinkedListReuseAfterClear(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Clear()
	l.Append(3)
	l.Append(4)
	if l.Size() != 2 {
		t.Errorf("Size(): esperaba 2, obtuvo %d", l.Size())
	}
	checkHeadTail(t, l, 3, 4)
}

func TestCircularLinkedListForward(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Forward(1)
	if got := l.Values(); len(got) != 4 || got[0] != 2 || got[1] != 3 || got[2] != 4 || got[3] != 1 {
		t.Errorf("Forward(1): esperaba [2 3 4 1], obtuvo %v", got)
	}
	checkHeadTail(t, l, 2, 1)
}

func TestCircularLinkedListForwardTwo(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Forward(2)
	if got := l.Values(); len(got) != 4 || got[0] != 3 || got[1] != 4 || got[2] != 1 || got[3] != 2 {
		t.Errorf("Forward(2): esperaba [3 4 1 2], obtuvo %v", got)
	}
	checkHeadTail(t, l, 3, 2)
}

func TestCircularLinkedListForwardFullCycle(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Forward(4)
	if got := l.Values(); len(got) != 4 || got[0] != 1 {
		t.Errorf("Forward(4): esperaba [1 2 3 4], obtuvo %v", got)
	}
	checkHeadTail(t, l, 1, 4)
}

func TestCircularLinkedListForwardOverflow(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Forward(6)
	if got := l.Values(); len(got) != 4 || got[0] != 3 || got[1] != 4 || got[2] != 1 || got[3] != 2 {
		t.Errorf("Forward(6): esperaba [3 4 1 2], obtuvo %v", got)
	}
}

func TestCircularLinkedListForwardNonPositive(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Forward(0)
	if got := l.Values(); len(got) != 3 || got[0] != 1 {
		t.Errorf("Forward(0): esperaba [1 2 3], obtuvo %v", got)
	}
	l.Forward(-2)
	if got := l.Values(); len(got) != 3 || got[0] != 1 {
		t.Errorf("Forward(-2): esperaba [1 2 3], obtuvo %v", got)
	}
}

func TestCircularLinkedListForwardEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Forward(3)
	if !l.IsEmpty() {
		t.Error("Forward(3): esperaba lista vacía intacta")
	}
}

func TestCircularLinkedListForwardSingleElement(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(42)
	l.Forward(5)
	checkHeadTail(t, l, 42, 42)
	if l.Size() != 1 {
		t.Errorf("Size(): esperaba 1, obtuvo %d", l.Size())
	}
}

func TestCircularLinkedListBackward(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Backward(1)
	if got := l.Values(); len(got) != 4 || got[0] != 4 || got[1] != 1 || got[2] != 2 || got[3] != 3 {
		t.Errorf("Backward(1): esperaba [4 1 2 3], obtuvo %v", got)
	}
	checkHeadTail(t, l, 4, 3)
}

func TestCircularLinkedListBackwardTwo(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Backward(2)
	if got := l.Values(); len(got) != 4 || got[0] != 3 || got[1] != 4 || got[2] != 1 || got[3] != 2 {
		t.Errorf("Backward(2): esperaba [3 4 1 2], obtuvo %v", got)
	}
	checkHeadTail(t, l, 3, 2)
}

func TestCircularLinkedListBackwardFullCycle(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Backward(4)
	if got := l.Values(); len(got) != 4 || got[0] != 1 {
		t.Errorf("Backward(4): esperaba [1 2 3 4], obtuvo %v", got)
	}
	checkHeadTail(t, l, 1, 4)
}

func TestCircularLinkedListBackwardOverflow(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Backward(6)
	if got := l.Values(); len(got) != 4 || got[0] != 3 || got[1] != 4 || got[2] != 1 || got[3] != 2 {
		t.Errorf("Backward(6): esperaba [3 4 1 2], obtuvo %v", got)
	}
}

func TestCircularLinkedListBackwardNonPositive(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Backward(0)
	if got := l.Values(); len(got) != 3 || got[0] != 1 {
		t.Errorf("Backward(0): esperaba [1 2 3], obtuvo %v", got)
	}
	l.Backward(-1)
	if got := l.Values(); len(got) != 3 || got[0] != 1 {
		t.Errorf("Backward(-1): esperaba [1 2 3], obtuvo %v", got)
	}
}

func TestCircularLinkedListBackwardEmpty(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Backward(3)
	if !l.IsEmpty() {
		t.Error("Backward(3): esperaba lista vacía intacta")
	}
}

func TestCircularLinkedListBackwardSingleElement(t *testing.T) {
	l := NewCircularLinkedList[int]()
	l.Append(42)
	l.Backward(5)
	checkHeadTail(t, l, 42, 42)
	if l.Size() != 1 {
		t.Errorf("Size(): esperaba 1, obtuvo %d", l.Size())
	}
}
