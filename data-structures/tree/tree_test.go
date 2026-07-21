package tree

import "testing"

func buildTree() *BinaryTree[int] {
	bt := NewBinaryTree[int]()
	bt.Root = &TreeNode[int]{Value: 2}
	bt.Root.Left = &TreeNode[int]{Value: 1}
	bt.Root.Right = &TreeNode[int]{Value: 3}
	return bt
}

func TestNewBinaryTree(t *testing.T) {
	bt := NewBinaryTree[int]()
	if bt == nil {
		t.Fatal("NewBinaryTree no debería retornar nil")
	}
	if bt.Root != nil {
		t.Error("El árbol nuevo debería tener raíz nil")
	}
}

func TestHeight(t *testing.T) {
	bt := NewBinaryTree[int]()
	if bt.Height() != 0 {
		t.Error("Árbol vacío debería tener altura 0")
	}

	bt = buildTree()
	if bt.Height() != 2 {
		t.Errorf("Altura esperada: 2, obtenida: %d", bt.Height())
	}
}

func TestSize(t *testing.T) {
	bt := NewBinaryTree[int]()
	if bt.Size() != 0 {
		t.Error("Árbol vacío debería tener size 0")
	}

	bt = buildTree()
	if bt.Size() != 3 {
		t.Errorf("Size esperado: 3, obtenido: %d", bt.Size())
	}
}

func TestInorderTraversal(t *testing.T) {
	bt := buildTree()
	result := bt.InorderTraversal()
	expected := []int{1, 2, 3}
	if len(result) != len(expected) {
		t.Fatalf("Longitud esperada: %d, obtenida: %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Posición %d: esperaba %d, obtuve %d", i, expected[i], v)
		}
	}
}

func TestPreorderTraversal(t *testing.T) {
	bt := buildTree()
	result := bt.PreorderTraversal()
	expected := []int{2, 1, 3}
	if len(result) != len(expected) {
		t.Fatalf("Longitud esperada: %d, obtenida: %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Posición %d: esperaba %d, obtuve %d", i, expected[i], v)
		}
	}
}

func TestPostorderTraversal(t *testing.T) {
	bt := buildTree()
	result := bt.PostorderTraversal()
	expected := []int{1, 3, 2}
	if len(result) != len(expected) {
		t.Fatalf("Longitud esperada: %d, obtenida: %d", len(expected), len(result))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Posición %d: esperaba %d, obtuve %d", i, expected[i], v)
		}
	}
}

func TestEmptyTreeTraversals(t *testing.T) {
	bt := NewBinaryTree[int]()
	if r := bt.InorderTraversal(); len(r) != 0 {
		t.Error("Inorder de árbol vacío debería ser []")
	}
	if r := bt.PreorderTraversal(); len(r) != 0 {
		t.Error("Preorder de árbol vacío debería ser []")
	}
	if r := bt.PostorderTraversal(); len(r) != 0 {
		t.Error("Postorder de árbol vacío debería ser []")
	}
}

func TestSingleNode(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Root = &TreeNode[int]{Value: 42}
	if bt.Height() != 1 {
		t.Errorf("Altura esperada: 1, obtenida: %d", bt.Height())
	}
	if bt.Size() != 1 {
		t.Errorf("Size esperado: 1, obtenido: %d", bt.Size())
	}
	r := bt.InorderTraversal()
	if len(r) != 1 || r[0] != 42 {
		t.Error("Inorder de nodo único debería ser [42]")
	}
}

func TestUnbalancedTree(t *testing.T) {
	bt := NewBinaryTree[int]()
	bt.Root = &TreeNode[int]{Value: 1}
	bt.Root.Right = &TreeNode[int]{Value: 2}
	bt.Root.Right.Right = &TreeNode[int]{Value: 3}
	bt.Root.Right.Right.Right = &TreeNode[int]{Value: 4}

	if bt.Height() != 4 {
		t.Errorf("Altura esperada: 4, obtenida: %d", bt.Height())
	}
	if bt.Size() != 4 {
		t.Errorf("Size esperado: 4, obtenido: %d", bt.Size())
	}

	r := bt.InorderTraversal()
	expected := []int{1, 2, 3, 4}
	for i, v := range r {
		if v != expected[i] {
			t.Errorf("Posición %d: esperaba %d, obtuve %d", i, expected[i], v)
		}
	}
}
