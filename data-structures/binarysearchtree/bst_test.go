package binarysearchtree

import (
	"cmp"
	"slices"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	if bst == nil {
		t.Fatal("NewBinarySearchTree no debería retornar nil")
	}
}

func TestInsertAndSearch(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)

	if !bst.Search(4) {
		t.Error("Debería encontrar 4")
	}
	if bst.Search(9) {
		t.Error("No debería encontrar 9")
	}
	if !bst.Search(5) {
		t.Error("Debería encontrar la raíz 5")
	}
}

func TestInsertDuplicate(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(5)

	if bst.Size() != 1 {
		t.Errorf("Size esperado: 1, obtenido: %d", bst.Size())
	}
}

func TestDeleteLeaf(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		bst.Insert(v)
	}
	bst.Delete(2)
	if bst.Search(2) {
		t.Error("2 no debería estar después de eliminar")
	}
	if !bst.Search(3) {
		t.Error("3 debería seguir estando")
	}
}

func TestDeleteOneChild(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(2)

	bst.Delete(3)
	if bst.Search(3) {
		t.Error("3 no debería estar después de eliminar")
	}
	if !bst.Search(2) {
		t.Error("2 debería seguir estando")
	}
}

func TestDeleteTwoChildren(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		bst.Insert(v)
	}
	bst.Delete(5)
	if bst.Search(5) {
		t.Error("5 no debería estar después de eliminar")
	}
	if !bst.Search(3) || !bst.Search(7) {
		t.Error("Los demás nodos deberían seguir existiendo")
	}
}

func TestHeight(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	if bst.Height() != 0 {
		t.Error("Árbol vacío debería tener altura 0")
	}

	bst.Insert(5)
	if bst.Height() != 1 {
		t.Errorf("Altura esperada: 1, obtenida: %d", bst.Height())
	}

	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)
	if bst.Height() != 3 {
		t.Errorf("Altura esperada: 3, obtenida: %d", bst.Height())
	}
}

func TestTraversals(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		bst.Insert(v)
	}

	inorder := bst.InorderTraversal()
	expected := []int{2, 3, 4, 5, 6, 7, 8}
	if !slices.Equal(inorder, expected) {
		t.Errorf("Inorder esperado: %v, obtenido: %v", expected, inorder)
	}

	preorder := bst.PreorderTraversal()
	expected = []int{5, 3, 2, 4, 7, 6, 8}
	if !slices.Equal(preorder, expected) {
		t.Errorf("Preorder esperado: %v, obtenido: %v", expected, preorder)
	}

	postorder := bst.PostorderTraversal()
	expected = []int{2, 4, 3, 6, 8, 7, 5}
	if !slices.Equal(postorder, expected) {
		t.Errorf("Postorder esperado: %v, obtenido: %v", expected, postorder)
	}
}

func TestStringTree(t *testing.T) {
	bst := NewBinarySearchTree[string](cmp.Compare[string])
	bst.Insert("banana")
	bst.Insert("manzana")
	bst.Insert("naranja")

	inorder := bst.InorderTraversal()
	expected := []string{"banana", "manzana", "naranja"}
	if !slices.Equal(inorder, expected) {
		t.Errorf("Inorder esperado: %v, obtenido: %v", expected, inorder)
	}

	if !bst.Search("manzana") {
		t.Error("Debería encontrar manzana")
	}
	if bst.Search("pera") {
		t.Error("No debería encontrar pera")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	bst.Insert(5)
	bst.Delete(99)

	if !bst.Search(5) {
		t.Error("5 debería seguir estando después de eliminar un valor inexistente")
	}
}

func TestSize(t *testing.T) {
	bst := NewBinarySearchTree[int](cmp.Compare[int])
	if bst.Size() != 0 {
		t.Errorf("Size esperado: 0, obtenido: %d", bst.Size())
	}

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	if bst.Size() != 3 {
		t.Errorf("Size esperado: 3, obtenido: %d", bst.Size())
	}

	bst.Delete(3)
	if bst.Size() != 2 {
		t.Errorf("Size esperado: 2, obtenido: %d", bst.Size())
	}
}

func TestCustomCompare(t *testing.T) {
	// Orden descendente (inverso)
	bst := NewBinarySearchTree[int](func(a, b int) int { return cmp.Compare(b, a) })
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)

	inorder := bst.InorderTraversal()
	expected := []int{3, 2, 1}
	if !slices.Equal(inorder, expected) {
		t.Errorf("Inorder (desc) esperado: %v, obtenido: %v", expected, inorder)
	}
}

func TestStructWithCompare(t *testing.T) {
	type Persona struct {
		Nombre string
		Edad   int
	}
	// Ordenar por edad
	bst := NewBinarySearchTree[Persona](func(a, b Persona) int {
		return cmp.Compare(a.Edad, b.Edad)
	})
	bst.Insert(Persona{"Ana", 30})
	bst.Insert(Persona{"Luis", 25})
	bst.Insert(Persona{"Zoe", 35})

	inorder := bst.InorderTraversal()
	if inorder[0].Nombre != "Luis" || inorder[2].Nombre != "Zoe" {
		t.Errorf("Esperaba Luis, Ana, Zoe; obtuve: %v", inorder)
	}
}
