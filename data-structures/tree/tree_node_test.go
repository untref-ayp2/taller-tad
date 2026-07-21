package tree

import "testing"

func TestNewTreeNode(t *testing.T) {
	n := NewTreeNode(42)
	if n == nil {
		t.Fatal("NewTreeNode devolvió nil")
	}
	if n.Value != 42 {
		t.Errorf("esperaba 42, obtuve %d", n.Value)
	}
	if !n.IsLeaf() {
		t.Error("un nodo nuevo debería ser hoja")
	}
}

func TestTreeNodeIsLeaf(t *testing.T) {
	root := NewTreeNode(1)
	if !root.IsLeaf() {
		t.Error("nodo sin hijos debería ser hoja")
	}

	root.Left = NewTreeNode(2)
	if root.IsLeaf() {
		t.Error("nodo con hijo izquierdo no debería ser hoja")
	}

	root.Right = NewTreeNode(3)
	if root.IsLeaf() {
		t.Error("nodo con ambos hijos no debería ser hoja")
	}
}
