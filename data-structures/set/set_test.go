package set

import (
	"testing"
)

// --- MapSet tests ---

func TestMapSetNewIsEmpty(t *testing.T) {
	s := NewMapSet[int]()
	if s.Size() != 0 {
		t.Error("MapSet recién creado debería tener Size 0")
	}
}

func TestMapSetAddAndContains(t *testing.T) {
	s := NewMapSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if !s.Contains(1) {
		t.Error("Debería contener 1")
	}
	if !s.Contains(2) {
		t.Error("Debería contener 2")
	}
	if s.Contains(4) {
		t.Error("No debería contener 4")
	}
}

func TestMapSetAddDuplicate(t *testing.T) {
	s := NewMapSet[int]()
	s.Add(1)
	s.Add(1)
	s.Add(1)

	if s.Size() != 1 {
		t.Errorf("Size = %d, esperaba 1 (elementos duplicados no deberían agregarse)", s.Size())
	}
}

func TestMapSetRemove(t *testing.T) {
	s := NewMapSet[int]()
	s.Add(1)
	s.Add(2)
	s.Remove(1)

	if s.Contains(1) {
		t.Error("Debería haber eliminado 1")
	}
	if !s.Contains(2) {
		t.Error("Debería conservar 2")
	}
}

func TestMapSetRemoveNonExistent(t *testing.T) {
	s := NewMapSet[int]()
	s.Add(1)
	s.Remove(99) // no debería panic

	if s.Size() != 1 {
		t.Error("Remove de elemento inexistente no debería cambiar Size")
	}
}

func TestMapSetSize(t *testing.T) {
	s := NewMapSet[string]()
	s.Add("a")
	s.Add("b")
	s.Add("c")

	if s.Size() != 3 {
		t.Errorf("Size = %d, esperaba 3", s.Size())
	}

	s.Remove("a")
	if s.Size() != 2 {
		t.Errorf("Size = %d, esperaba 2", s.Size())
	}
}

func TestMapSetValues(t *testing.T) {
	s := NewMapSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	vals := s.Values()
	if len(vals) != 3 {
		t.Fatalf("Values() len = %d, esperaba 3", len(vals))
	}

	m := make(map[int]bool)
	for _, v := range vals {
		m[v] = true
	}
	if !m[1] || !m[2] || !m[3] {
		t.Errorf("Values() = %v, debería contener 1, 2, 3", vals)
	}
}

func TestMapSetString(t *testing.T) {
	s := NewMapSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	str := s.String()
	if len(str) == 0 {
		t.Error("String() no debería estar vacío")
	}
}

func TestMapSetUnion(t *testing.T) {
	a := NewMapSet[int]()
	a.Add(1)
	a.Add(2)

	b := NewMapSet[int]()
	b.Add(2)
	b.Add(3)

	u := a.Union(b)
	if u.Size() != 3 {
		t.Errorf("Union Size = %d, esperaba 3", u.Size())
	}
	if !u.Contains(1) || !u.Contains(2) || !u.Contains(3) {
		t.Error("Union debería contener 1, 2, 3")
	}
}

func TestMapSetIntersection(t *testing.T) {
	a := NewMapSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewMapSet[int]()
	b.Add(2)
	b.Add(3)
	b.Add(4)

	i := a.Intersection(b)
	if i.Size() != 2 {
		t.Errorf("Intersection Size = %d, esperaba 2", i.Size())
	}
	if !i.Contains(2) || !i.Contains(3) {
		t.Error("Intersection debería contener 2, 3")
	}
	if i.Contains(1) || i.Contains(4) {
		t.Error("Intersection no debería contener 1 ni 4")
	}
}

func TestMapSetDifference(t *testing.T) {
	a := NewMapSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewMapSet[int]()
	b.Add(2)

	d := a.Difference(b)
	if d.Size() != 2 {
		t.Errorf("Difference Size = %d, esperaba 2", d.Size())
	}
	if !d.Contains(1) || !d.Contains(3) {
		t.Error("Difference debería contener 1, 3")
	}
	if d.Contains(2) {
		t.Error("Difference no debería contener 2")
	}
}

func TestMapSetSymmetricDifference(t *testing.T) {
	a := NewMapSet[int]()
	a.Add(1)
	a.Add(2)

	b := NewMapSet[int]()
	b.Add(2)
	b.Add(3)

	sd := a.SymmetricDifference(b)
	if sd.Size() != 2 {
		t.Errorf("SymmetricDifference Size = %d, esperaba 2", sd.Size())
	}
	if !sd.Contains(1) || !sd.Contains(3) {
		t.Error("SymmetricDifference debería contener 1, 3")
	}
	if sd.Contains(2) {
		t.Error("SymmetricDifference no debería contener 2")
	}
}

func TestMapSetSubset(t *testing.T) {
	a := NewMapSet[int]()
	a.Add(1)
	a.Add(2)

	b := NewMapSet[int]()
	b.Add(1)
	b.Add(2)
	b.Add(3)

	if !a.Subset(b) {
		t.Error("{1,2} debería ser subconjunto de {1,2,3}")
	}
	if b.Subset(a) {
		t.Error("{1,2,3} no debería ser subconjunto de {1,2}")
	}
}

func TestMapSetSuperset(t *testing.T) {
	a := NewMapSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewMapSet[int]()
	b.Add(1)
	b.Add(2)

	if !a.Superset(b) {
		t.Error("{1,2,3} debería ser superconjunto de {1,2}")
	}
	if b.Superset(a) {
		t.Error("{1,2} no debería ser superconjunto de {1,2,3}")
	}
}

// --- OrderedSet tests ---

func lessInt(a, b int) bool { return a < b }

func TestOrderedSetNewIsEmpty(t *testing.T) {
	s := NewOrderedSet(lessInt)
	if s.Size() != 0 {
		t.Error("OrderedSet recién creado debería tener Size 0")
	}
}

func TestOrderedSetAddAndContains(t *testing.T) {
	s := NewOrderedSet(lessInt)
	s.Add(3)
	s.Add(1)
	s.Add(2)

	if !s.Contains(1) {
		t.Error("Debería contener 1")
	}
	if !s.Contains(2) {
		t.Error("Debería contener 2")
	}
	if s.Contains(4) {
		t.Error("No debería contener 4")
	}
}

func TestOrderedSetValuesOrdered(t *testing.T) {
	s := NewOrderedSet(lessInt)
	s.Add(3)
	s.Add(1)
	s.Add(2)

	vals := s.Values()
	esperado := []int{1, 2, 3}
	if len(vals) != 3 {
		t.Fatalf("Values() len = %d, esperaba 3", len(vals))
	}
	for i, v := range vals {
		if v != esperado[i] {
			t.Errorf("Values()[%d] = %d, esperaba %d", i, v, esperado[i])
		}
	}
}

func TestOrderedSetUnion(t *testing.T) {
	a := NewOrderedSet(lessInt)
	a.Add(1)
	a.Add(3)

	b := NewOrderedSet(lessInt)
	b.Add(2)
	b.Add(4)

	u := a.Union(b)
	if u.Size() != 4 {
		t.Errorf("Union Size = %d, esperaba 4", u.Size())
	}
	for _, v := range []int{1, 2, 3, 4} {
		if !u.Contains(v) {
			t.Errorf("Union debería contener %d", v)
		}
	}
}

func TestOrderedSetSubset(t *testing.T) {
	a := NewOrderedSet(lessInt)
	a.Add(1)
	a.Add(2)

	b := NewOrderedSet(lessInt)
	b.Add(1)
	b.Add(2)
	b.Add(3)

	if !a.Subset(b) {
		t.Error("{1,2} debería ser subconjunto de {1,2,3}")
	}
}
