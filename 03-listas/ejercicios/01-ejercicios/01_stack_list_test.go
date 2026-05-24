package ejercicios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/stack"
)

func TestStackListPushAndPop(t *testing.T) {
	s := NewStackList[int]()
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

func TestStackListPopWhenEmpty(t *testing.T) {
	s := NewStackList[int]()
	_, err := s.Pop()
	if err == nil {
		t.Error("esperaba error al hacer Pop de pila vacía")
	}
}

func TestStackListTop(t *testing.T) {
	s := NewStackList[string]()
	s.Push("a")
	s.Push("b")

	val, err := s.Top()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != "b" {
		t.Errorf("esperaba 'b', obtuve %s", val)
	}

	val, err = s.Top()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != "b" {
		t.Errorf("esperaba 'b', obtuve %s", val)
	}
}

func TestStackListTopWhenEmpty(t *testing.T) {
	s := NewStackList[int]()
	_, err := s.Top()
	if err == nil {
		t.Error("esperaba error al hacer Top de pila vacía")
	}
}

func TestStackListIsEmpty(t *testing.T) {
	s := NewStackList[float64]()

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

func TestStackListSatisfaceInterface(t *testing.T) {
	var s stack.Stack[int] = NewStackList[int]()
	s.Push(1)
	s.Push(2)
	val, err := s.Pop()
	if err != nil {
		t.Fatalf("Pop error: %v", err)
	}
	if val != 2 {
		t.Errorf("esperaba 2, obtuve %d", val)
	}
}
