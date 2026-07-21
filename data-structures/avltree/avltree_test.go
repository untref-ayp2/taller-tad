package avltree

import (
	"cmp"
	"slices"
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	if tree == nil {
		t.Fatal("NewAVLTree no deberia retornar nil")
	}
	if tree.Size() != 0 {
		t.Errorf("Size esperado: 0, obtenido: %d", tree.Size())
	}
}

// buildTree construye un arbol AVL manualmente con los valores dados
// en orden de insercion. Util para probar Delete sin depender de Insert.
// TODO: size++ se ejecuta incluso si insertRec skipea un duplicado.
// Si buildTree se usara con duplicados, habria que mover size++ dentro
// de insertRec o validar antes de incrementar.
func buildTree(t *testing.T, values ...int) *AVLTree[int] {
	tree := NewAVLTree[int](cmp.Compare[int])
	for _, v := range values {
		tree.root = insertRec(tree.root, v, tree.cmp)
		tree.size++
	}
	return tree
}

func insertRec[T any](n *AVLNode[T], value T, cmp func(T, T) int) *AVLNode[T] {
	if n == nil {
		return NewAVLNode(value)
	}
	c := cmp(value, n.value)
	if c < 0 {
		n.left = insertRec(n.left, value, cmp)
	} else if c > 0 {
		n.right = insertRec(n.right, value, cmp)
	} else {
		return n
	}
	n.updateHeight()
	return rebalance(n)
}

func collectInorder[T any](n *AVLNode[T]) []T {
	if n == nil {
		return nil
	}
	result := collectInorder(n.left)
	result = append(result, n.value)
	result = append(result, collectInorder(n.right)...)
	return result
}

func TestDeleteLeaf(t *testing.T) {
	tree := buildTree(t, 10, 5, 15)
	tree.Delete(5)
	if tree.Size() != 2 {
		t.Errorf("Size esperado: 2, obtenido: %d", tree.Size())
	}
	inorder := collectInorder(tree.root)
	if slices.Contains(inorder, 5) {
		t.Error("5 no deberia estar despues de eliminar")
	}
	if tree.root.BalanceFactor() < -1 || tree.root.BalanceFactor() > 1 {
		t.Errorf("Arbol desbalanceado luego de eliminar hoja, fb(raiz) = %d", tree.root.BalanceFactor())
	}
}

func TestDeleteOneChild(t *testing.T) {
	tree := buildTree(t, 10, 5, 3)
	tree.Delete(5)
	if tree.Size() != 2 {
		t.Errorf("Size esperado: 2, obtenido: %d", tree.Size())
	}
	inorder := collectInorder(tree.root)
	if slices.Contains(inorder, 5) {
		t.Error("5 no deberia estar despues de eliminar")
	}
	if !slices.Contains(inorder, 3) {
		t.Error("3 deberia seguir estando")
	}
}

func TestDeleteTwoChildren(t *testing.T) {
	tree := buildTree(t, 10, 5, 15, 3, 7)
	tree.Delete(10)
	if tree.Size() != 4 {
		t.Errorf("Size esperado: 4, obtenido: %d", tree.Size())
	}
	inorder := collectInorder(tree.root)
	if slices.Contains(inorder, 10) {
		t.Error("10 no deberia estar despues de eliminar")
	}
	if !slices.Contains(inorder, 3) || !slices.Contains(inorder, 5) ||
		!slices.Contains(inorder, 7) || !slices.Contains(inorder, 15) {
		t.Error("Los demas nodos deberian seguir existiendo")
	}
}

func TestDeleteRoot(t *testing.T) {
	tree := buildTree(t, 10, 5, 15)
	tree.Delete(10)
	if tree.Size() != 2 {
		t.Errorf("Size esperado: 2, obtenido: %d", tree.Size())
	}
	inorder := collectInorder(tree.root)
	if slices.Contains(inorder, 10) {
		t.Error("10 no deberia estar despues de eliminar")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	tree := buildTree(t, 10, 5, 15)
	tree.Delete(99)
	if tree.Size() != 3 {
		t.Errorf("Size esperado: 3, obtenido: %d", tree.Size())
	}
}

func TestDeleteFromEmptyTree(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	tree.Delete(10)
	if tree.Size() != 0 {
		t.Errorf("Size esperado: 0, obtenido: %d", tree.Size())
	}
}

func TestDeleteRebalances(t *testing.T) {
	tree := buildTree(t, 30, 20, 40, 10, 25, 35, 50, 5, 15)
	tree.Delete(40)
	inorder := collectInorder(tree.root)
	expected := []int{5, 10, 15, 20, 25, 30, 35, 50}
	if !slices.Equal(inorder, expected) {
		t.Errorf("Inorder esperado: %v, obtenido: %v", expected, inorder)
	}
	var checkBalance func(*AVLNode[int]) int
	checkBalance = func(n *AVLNode[int]) int {
		if n == nil {
			return 0
		}
		lh := checkBalance(n.left)
		rh := checkBalance(n.right)
		fb := lh - rh
		if fb < -1 || fb > 1 {
			t.Errorf("Nodo %d desbalanceado: fb = %d", n.value, fb)
		}
		if lh > rh {
			return lh + 1
		}
		return rh + 1
	}
	checkBalance(tree.root)
}

func TestDoubleRotationLeftRight(t *testing.T) {
	a := NewAVLNode(30)
	b := NewAVLNode(20)
	c := NewAVLNode(25)
	a.left = b
	b.right = c
	b.updateHeight()
	c.updateHeight()
	a.updateHeight()

	if a.BalanceFactor() != 2 {
		t.Fatalf("fb(30) esperado: 2, obtenido: %d", a.BalanceFactor())
	}

	a = rebalance(a)

	if a.value != 25 {
		t.Errorf("Raiz esperada: 25, obtenida: %d", a.value)
	}
	if a.BalanceFactor() < -1 || a.BalanceFactor() > 1 {
		t.Errorf("Raiz desbalanceada: fb = %d", a.BalanceFactor())
	}
}

func TestDoubleRotationRightLeft(t *testing.T) {
	a := NewAVLNode(30)
	b := NewAVLNode(40)
	c := NewAVLNode(35)
	a.right = b
	b.left = c
	b.updateHeight()
	c.updateHeight()
	a.updateHeight()

	if a.BalanceFactor() != -2 {
		t.Fatalf("fb(30) esperado: -2, obtenido: %d", a.BalanceFactor())
	}

	a = rebalance(a)

	if a.value != 35 {
		t.Errorf("Raiz esperada: 35, obtenida: %d", a.value)
	}
	if a.BalanceFactor() < -1 || a.BalanceFactor() > 1 {
		t.Errorf("Raiz desbalanceada: fb = %d", a.BalanceFactor())
	}
}

func TestInsertAndSearch(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	if tree.Search(7) == nil {
		t.Error("Deberia encontrar 7")
	}
	if tree.Search(10) == nil {
		t.Error("Deberia encontrar la raiz 10")
	}
	if tree.Search(99) != nil {
		t.Error("No deberia encontrar 99")
	}
	if tree.Size() != 5 {
		t.Errorf("Size esperado: 5, obtenido: %d", tree.Size())
	}
}

func TestInsertMaintainsBalance(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	values := []int{10, 20, 30, 40, 50, 5, 3, 7, 15, 25}
	for _, v := range values {
		tree.Insert(v)
	}

	var checkBalance func(*AVLNode[int]) int
	checkBalance = func(n *AVLNode[int]) int {
		if n == nil {
			return 0
		}
		lh := checkBalance(n.left)
		rh := checkBalance(n.right)
		fb := lh - rh
		if fb < -1 || fb > 1 {
			t.Errorf("Nodo %d desbalanceado: fb = %d", n.value, fb)
		}
		if lh > rh {
			return lh + 1
		}
		return rh + 1
	}
	checkBalance(tree.root)
}

func TestInsertDuplicate(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	tree.Insert(5)
	tree.Insert(5)
	tree.Insert(5)

	if tree.Size() != 1 {
		t.Errorf("Size esperado: 1, obtenido: %d", tree.Size())
	}
}

func TestInsertAscending(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	for i := 1; i <= 100; i++ {
		tree.Insert(i)
	}

	var checkBalance func(*AVLNode[int]) int
	checkBalance = func(n *AVLNode[int]) int {
		if n == nil {
			return 0
		}
		lh := checkBalance(n.left)
		rh := checkBalance(n.right)
		fb := lh - rh
		if fb < -1 || fb > 1 {
			t.Errorf("Nodo %d desbalanceado: fb = %d", n.value, fb)
		}
		if lh > rh {
			return lh + 1
		}
		return rh + 1
	}
	checkBalance(tree.root)

	height := tree.root.Height()
	maxExpected := 12
	if height > maxExpected {
		t.Errorf("Altura demasiado grande: %d, esperada <= %d (aprox 1.44*log2(n))", height, maxExpected)
	}
}

func TestInorderTraversal(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		tree.Insert(v)
	}

	inorder := tree.InorderTraversal()
	expected := []int{2, 3, 4, 5, 6, 7, 8}
	if !slices.Equal(inorder, expected) {
		t.Errorf("Inorder esperado: %v, obtenido: %v", expected, inorder)
	}
}

func TestPreorderTraversal(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		tree.Insert(v)
	}

	preorder := tree.PreorderTraversal()
	expected := []int{5, 3, 2, 4, 7, 6, 8}
	if !slices.Equal(preorder, expected) {
		t.Errorf("Preorder esperado: %v, obtenido: %v", expected, preorder)
	}
}

func TestPostorderTraversal(t *testing.T) {
	tree := NewAVLTree[int](cmp.Compare[int])
	for _, v := range []int{5, 3, 7, 2, 4, 6, 8} {
		tree.Insert(v)
	}

	postorder := tree.PostorderTraversal()
	expected := []int{2, 4, 3, 6, 8, 7, 5}
	if !slices.Equal(postorder, expected) {
		t.Errorf("Postorder esperado: %v, obtenido: %v", expected, postorder)
	}
}
