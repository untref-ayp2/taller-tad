package avltree

// AVLTree representa un árbol AVL balanceado.
type AVLTree[T any] struct {
	root *AVLNode[T]
	size int
	cmp  func(T, T) int
}

// NewAVLTree crea un árbol AVL vacío.
// La función cmp debe devolver un valor negativo si a < b,
// cero si a == b, y positivo si a > b.
func NewAVLTree[T any](cmp func(T, T) int) *AVLTree[T] {
	return &AVLTree[T]{cmp: cmp}
}

// Insert inserta un valor en el árbol. No inserta duplicados.
func (t *AVLTree[T]) Insert(value T) {
	// Completar
}

// Search busca un valor en el árbol. Devuelve el nodo o nil si no existe.
func (t *AVLTree[T]) Search(value T) *AVLNode[T] {
	// Completar
	return nil
}

// Delete elimina un valor del árbol. No hace nada si no existe.
func (t *AVLTree[T]) Delete(value T) {
	t.root = t.deleteRec(t.root, value)
}

func (t *AVLTree[T]) deleteRec(n *AVLNode[T], value T) *AVLNode[T] {
	if n == nil {
		return nil
	}

	c := t.cmp(value, n.value)
	if c < 0 {
		n.left = t.deleteRec(n.left, value)
	} else if c > 0 {
		n.right = t.deleteRec(n.right, value)
	} else {
		t.size--
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		succ := minValue(n.right)
		n.value = succ.value
		n.right = t.deleteRec(n.right, succ.value)
	}

	n.updateHeight()
	return rebalance(n)
}

// minValue encuentra el nodo con el valor mínimo en el subárbol.
func minValue[T any](n *AVLNode[T]) *AVLNode[T] {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

// Size devuelve la cantidad de nodos en el árbol.
func (t *AVLTree[T]) Size() int {
	return t.size
}

// rotateRight rota el subárbol a la derecha.
// Completar: reasignar punteros y actualizar alturas.
func rotateRight[T any](y *AVLNode[T]) *AVLNode[T] {
	// Completar
	return y
}

// rotateLeft rota el subárbol a la izquierda.
// Completar: reasignar punteros y actualizar alturas.
func rotateLeft[T any](x *AVLNode[T]) *AVLNode[T] {
	// Completar
	return x
}

// rebalance verifica el factor de balanceo del nodo y aplica
// la rotación correspondiente si está desbalanceado.
func rebalance[T any](n *AVLNode[T]) *AVLNode[T] {
	fb := n.BalanceFactor()
	if fb > 1 {
		if n.left.BalanceFactor() < 0 {
			n.left = rotateLeft(n.left)
		}
		return rotateRight(n)
	}
	if fb < -1 {
		if n.right.BalanceFactor() > 0 {
			n.right = rotateRight(n.right)
		}
		return rotateLeft(n)
	}
	return n
}

// InorderTraversal devuelve un slice con los valores en recorrido inorder.
func (t *AVLTree[T]) InorderTraversal() []T {
	// Completar
	return nil
}

// PreorderTraversal devuelve un slice con los valores en recorrido preorder.
func (t *AVLTree[T]) PreorderTraversal() []T {
	// Completar
	return nil
}

// PostorderTraversal devuelve un slice con los valores en recorrido postorder.
func (t *AVLTree[T]) PostorderTraversal() []T {
	// Completar
	return nil
}
