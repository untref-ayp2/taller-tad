package parser

import (
	"testing"

	"github.com/untref-ayp2/data-structures/tree"
)

func checkTree(t *testing.T, node *tree.TreeNode[string], expectedVal string, hasLeft, hasRight bool) {
	t.Helper()
	if node == nil {
		t.Fatal("nodo es nil")
	}
	if node.Value() != expectedVal {
		t.Errorf("esperaba valor %q, obtuve %q", expectedVal, node.Value())
	}
	if hasLeft && node.Left() == nil {
		t.Errorf("nodo %q debería tener hijo izquierdo", expectedVal)
	}
	if !hasLeft && node.Left() != nil {
		t.Errorf("nodo %q no debería tener hijo izquierdo", expectedVal)
	}
	if hasRight && node.Right() == nil {
		t.Errorf("nodo %q debería tener hijo derecho", expectedVal)
	}
	if !hasRight && node.Right() != nil {
		t.Errorf("nodo %q no debería tener hijo derecho", expectedVal)
	}
}

func TestParseNumber(t *testing.T) {
	node, err := Parse("42")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	checkTree(t, node, "42", false, false)
}

func TestParseSimpleSum(t *testing.T) {
	node, err := Parse("3+4")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	checkTree(t, node, "+", true, true)
	checkTree(t, node.Left(), "3", false, false)
	checkTree(t, node.Right(), "4", false, false)
}

func TestParsePrecedence(t *testing.T) {
	node, err := Parse("3+4*2")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	checkTree(t, node, "+", true, true)
	checkTree(t, node.Left(), "3", false, false)
	checkTree(t, node.Right(), "*", true, true)
	checkTree(t, node.Right().Left(), "4", false, false)
	checkTree(t, node.Right().Right(), "2", false, false)
}

func TestParseParentheses(t *testing.T) {
	node, err := Parse("(3+4)*2")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	checkTree(t, node, "*", true, true)
	checkTree(t, node.Left(), "+", true, true)
	checkTree(t, node.Left().Left(), "3", false, false)
	checkTree(t, node.Left().Right(), "4", false, false)
	checkTree(t, node.Right(), "2", false, false)
}

func TestParseComplex(t *testing.T) {
	node, err := Parse("(3+4)*(5-2)")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	checkTree(t, node, "*", true, true)
	checkTree(t, node.Left(), "+", true, true)
	checkTree(t, node.Right(), "-", true, true)
}

func TestParseError(t *testing.T) {
	_, err := Parse("3+@4")
	if err == nil {
		t.Error("esperaba error por caracter inválido")
	}
}
