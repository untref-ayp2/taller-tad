package ejercicios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/tree"
	"github.com/untref-ayp2/taller-tad/08-arboles/parser"
)

// buildExampleTree construye el árbol del apunte:
//
//	      +
//	     / \
//	    a   *
//	       / \
//	      -   d
//	     / \
//	    b   c
func buildExampleTree() *tree.TreeNode[string] {
	root := tree.NewTreeNode("+")

	a := tree.NewTreeNode("a")
	root.SetLeft(a)

	star := tree.NewTreeNode("*")
	root.SetRight(star)

	minus := tree.NewTreeNode("−")
	star.SetLeft(minus)

	d := tree.NewTreeNode("d")
	star.SetRight(d)

	b := tree.NewTreeNode("b")
	minus.SetLeft(b)

	c := tree.NewTreeNode("c")
	minus.SetRight(c)

	return root
}

func slicesIguales[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPreorden(t *testing.T) {
	root := buildExampleTree()
	result := Preorden(root)
	expected := []string{"+", "a", "*", "−", "b", "c", "d"}
	if !slicesIguales(result, expected) {
		t.Errorf("Preorden\nesperaba: %v\nobtuve:   %v", expected, result)
	}
}

func TestPreordenNil(t *testing.T) {
	result := Preorden[*string](nil)
	if len(result) != 0 {
		t.Errorf("Preorden de nil debería devolver slice vacío, obtuve %v", result)
	}
}

func TestInorden(t *testing.T) {
	root := buildExampleTree()
	result := Inorden(root)
	expected := []string{"a", "+", "b", "−", "c", "*", "d"}
	if !slicesIguales(result, expected) {
		t.Errorf("Inorden\nesperaba: %v\nobtuve:   %v", expected, result)
	}
}

func TestInordenNil(t *testing.T) {
	result := Inorden[*string](nil)
	if len(result) != 0 {
		t.Errorf("Inorden de nil debería devolver slice vacío, obtuve %v", result)
	}
}

func TestPostorden(t *testing.T) {
	root := buildExampleTree()
	result := Postorden(root)
	expected := []string{"a", "b", "c", "−", "d", "*", "+"}
	if !slicesIguales(result, expected) {
		t.Errorf("Postorden\nesperaba: %v\nobtuve:   %v", expected, result)
	}
}

func TestPostordenNil(t *testing.T) {
	result := Postorden[*string](nil)
	if len(result) != 0 {
		t.Errorf("Postorden de nil debería devolver slice vacío, obtuve %v", result)
	}
}

func TestEvaluarSimple(t *testing.T) {
	arbol, err := parser.Parse("3+4")
	if err != nil {
		t.Fatalf("error al parsear: %v", err)
	}
	result, err := Evaluar(arbol)
	if err != nil {
		t.Fatalf("error al evaluar: %v", err)
	}
	if result != 7 {
		t.Errorf("3+4 esperaba 7, obtuve %d", result)
	}
}

func TestEvaluarConPrecedencia(t *testing.T) {
	arbol, err := parser.Parse("3+4*2")
	if err != nil {
		t.Fatalf("error al parsear: %v", err)
	}
	result, err := Evaluar(arbol)
	if err != nil {
		t.Fatalf("error al evaluar: %v", err)
	}
	if result != 11 {
		t.Errorf("3+4*2 esperaba 11, obtuve %d", result)
	}
}

func TestEvaluarConParentesis(t *testing.T) {
	arbol, err := parser.Parse("(3+4)*2")
	if err != nil {
		t.Fatalf("error al parsear: %v", err)
	}
	result, err := Evaluar(arbol)
	if err != nil {
		t.Fatalf("error al evaluar: %v", err)
	}
	if result != 14 {
		t.Errorf("(3+4)*2 esperaba 14, obtuve %d", result)
	}
}

func TestEvaluarResta(t *testing.T) {
	arbol, err := parser.Parse("10-3")
	if err != nil {
		t.Fatalf("error al parsear: %v", err)
	}
	result, err := Evaluar(arbol)
	if err != nil {
		t.Fatalf("error al evaluar: %v", err)
	}
	if result != 7 {
		t.Errorf("10-3 esperaba 7, obtuve %d", result)
	}
}

func TestEvaluarDivision(t *testing.T) {
	arbol, err := parser.Parse("8/2")
	if err != nil {
		t.Fatalf("error al parsear: %v", err)
	}
	result, err := Evaluar(arbol)
	if err != nil {
		t.Fatalf("error al evaluar: %v", err)
	}
	if result != 4 {
		t.Errorf("8/2 esperaba 4, obtuve %d", result)
	}
}

func TestEvaluarMultiplesOperadores(t *testing.T) {
	arbol, err := parser.Parse("(3+4)*(5-2)")
	if err != nil {
		t.Fatalf("error al parsear: %v", err)
	}
	result, err := Evaluar(arbol)
	if err != nil {
		t.Fatalf("error al evaluar: %v", err)
	}
	if result != 21 {
		t.Errorf("(3+4)*(5-2) esperaba 21, obtuve %d", result)
	}
}

func TestEvaluarDivisionPorCero(t *testing.T) {
	arbol, err := parser.Parse("5/0")
	if err != nil {
		t.Fatalf("error al parsear: %v", err)
	}
	_, err = Evaluar(arbol)
	if err == nil {
		t.Error("esperaba error por división por cero")
	}
}
