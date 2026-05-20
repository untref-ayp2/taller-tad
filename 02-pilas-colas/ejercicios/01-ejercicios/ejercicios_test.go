package ejercicios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/queue"
	"github.com/untref-ayp2/data-structures/stack"
)

func TestInvertirCadena(t *testing.T) {
	casos := []struct {
		entrada, esperado string
	}{
		{"hola", "aloh"},
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"12345", "54321"},
		{"radar", "radar"},
	}

	for _, c := range casos {
		resultado := InvertirCadena(c.entrada)
		if resultado != c.esperado {
			t.Errorf("InvertirCadena(%q) = %q, esperaba %q", c.entrada, resultado, c.esperado)
		}
	}
}

func TestEsPalindromo(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"1456541", true},
		{"145541", true},
		{"abc", false},
		{"hola", false},
		{"radar", true},
	}

	for _, c := range casos {
		resultado := EsPalindromo(c.entrada)
		if resultado != c.esperado {
			t.Errorf("EsPalindromo(%q) = %v, esperaba %v", c.entrada, resultado, c.esperado)
		}
	}
}

func TestEstaBalanceada(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado bool
	}{
		{"", true},
		{"[]", true},
		{"()", true},
		{"{}", true},
		{"[()]{}{[()()]()}", true},
		{"[(])", false},
		{"{", false},
		{"]", false},
		{"{[()]}", true},
		{"({[}])", false},
	}

	for _, c := range casos {
		resultado := EstaBalanceada(c.entrada)
		if resultado != c.esperado {
			t.Errorf("EstaBalanceada(%q) = %v, esperaba %v", c.entrada, resultado, c.esperado)
		}
	}
}

func TestUnirColas(t *testing.T) {
	q1 := queue.NewSliceQueue[int]()
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	q2 := queue.NewSliceQueue[int]()
	q2.Enqueue(5)
	q2.Enqueue(7)

	resultado := UnirColas[int](q1, q2)

	esperados := []int{1, 2, 3, 5, 7}
	for _, esp := range esperados {
		val, err := resultado.Dequeue()
		if err != nil {
			t.Fatalf("error inesperado al hacer Dequeue: %v", err)
		}
		if val != esp {
			t.Errorf("esperaba %d, obtuve %d", esp, val)
		}
	}

	if !resultado.IsEmpty() {
		t.Error("la cola resultado debería estar vacía después de extraer todos los elementos")
	}
}

func TestEvaluarRPN(t *testing.T) {
	casos := []struct {
		expresion string
		esperado  int
	}{
		{"2 3 + 5 *", 25},
		{"5 3 +", 8},
		{"10 2 /", 5},
		{"4 2 * 3 +", 11},
		{"10 3 - 2 *", 14},
		{"3 4 + 2 * 7 /", 2},
	}

	for _, c := range casos {
		resultado, err := EvaluarRPN(c.expresion)
		if err != nil {
			t.Errorf("EvaluarRPN(%q) error inesperado: %v", c.expresion, err)
			continue
		}
		if resultado != c.esperado {
			t.Errorf("EvaluarRPN(%q) = %d, esperaba %d", c.expresion, resultado, c.esperado)
		}
	}
}

func TestEsPosibleConPila(t *testing.T) {
	casos := []struct {
		entrada, salida []int
		esperado        bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{2, 1, 3}, true},
		{[]int{1, 2, 3}, []int{3, 1, 2}, false},
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}, true},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}, false},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, true},
	}

	for _, c := range casos {
		resultado := EsPosibleConPila(c.entrada, c.salida)
		if resultado != c.esperado {
			t.Errorf("EsPosibleConPila(%v, %v) = %v, esperaba %v",
				c.entrada, c.salida, resultado, c.esperado)
		}
	}
}

// Test helpers to ensure SliceStack and SliceQueue are usable
func TestUsaSliceStack(t *testing.T) {
	s := stack.NewSliceStack[int]()
	s.Push(1)
	s.Push(2)
	val, err := s.Top()
	if err != nil {
		t.Fatalf("Top error: %v", err)
	}
	if val != 2 {
		t.Errorf("esperaba 2, obtuve %d", val)
	}
}

func TestUsaSliceQueue(t *testing.T) {
	q := queue.NewSliceQueue[string]()
	q.Enqueue("a")
	q.Enqueue("b")
	val, err := q.Front()
	if err != nil {
		t.Fatalf("Front error: %v", err)
	}
	if val != "a" {
		t.Errorf("esperaba 'a', obtuve %s", val)
	}
}
