package stack

import "testing"

func TestNewListStackIsEmpty(t *testing.T) {
	s := NewListStack[int]()
	if !s.IsEmpty() {
		t.Error("esperaba pila vacía recién creada")
	}
}

func TestListStackPushAndPop(t *testing.T) {
	s := NewListStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	val, err := s.Pop()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 3 {
		t.Errorf("esperaba 3, obtuve %d", val)
	}

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 2 {
		t.Errorf("esperaba 2, obtuve %d", val)
	}
}

func TestListStackPopWhenEmpty(t *testing.T) {
	s := NewListStack[int]()
	_, err := s.Pop()
	if err == nil {
		t.Error("esperaba error al hacer Pop de pila vacía")
	}
}

func TestListStackTop(t *testing.T) {
	s := NewListStack[string]()
	s.Push("a")
	s.Push("b")

	val, err := s.Top()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != "b" {
		t.Errorf("esperaba 'b', obtuve %s", val)
	}

	// Top no debe eliminar el elemento
	val, err = s.Top()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != "b" {
		t.Errorf("esperaba 'b', obtuve %s", val)
	}
}

func TestListStackTopWhenEmpty(t *testing.T) {
	s := NewListStack[int]()
	_, err := s.Top()
	if err == nil {
		t.Error("esperaba error al hacer Top de pila vacía")
	}
}
