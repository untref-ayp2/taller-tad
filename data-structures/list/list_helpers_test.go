package list

import "testing"

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
