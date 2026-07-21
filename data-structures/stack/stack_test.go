package stack

import (
	"testing"
)

func TestSliceStackPushAndPop(t *testing.T) {
	s := NewSliceStack[int]()
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

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 1 {
		t.Errorf("esperaba 1, obtuve %d", val)
	}
}

func TestSliceStackPopWhenEmpty(t *testing.T) {
	s := NewSliceStack[int]()

	_, err := s.Pop()
	if err == nil {
		t.Error("esperaba error al hacer Pop de pila vacía")
	}
}

func TestSliceStackTop(t *testing.T) {
	s := NewSliceStack[string]()
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

func TestSliceStackTopWhenEmpty(t *testing.T) {
	s := NewSliceStack[int]()

	_, err := s.Top()
	if err == nil {
		t.Error("esperaba error al hacer Top de pila vacía")
	}
}

func TestSliceStackIsEmpty(t *testing.T) {
	s := NewSliceStack[float64]()

	if !s.IsEmpty() {
		t.Error("esperaba pila vacía recién creada")
	}

	s.Push(1.0)
	if s.IsEmpty() {
		t.Error("esperaba pila no vacía después de Push")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Error("esperaba pila vacía después de Pop del único elemento")
	}
}
