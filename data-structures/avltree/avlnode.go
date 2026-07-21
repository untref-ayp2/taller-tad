package avltree

// AVLNode representa un nodo de un árbol AVL.
// Almacena la altura del subárbol para calcular el factor de balanceo.
type AVLNode[T any] struct {
	value  T
	left   *AVLNode[T]
	right  *AVLNode[T]
	height int
}

// NewAVLNode crea un nuevo nodo hoja con el valor dado.
// La altura de una hoja es 0 (un árbol vacío tiene altura -1).
func NewAVLNode[T any](value T) *AVLNode[T] {
	return &AVLNode[T]{value: value, height: 0}
}

// Value devuelve el valor almacenado en el nodo.
func (n *AVLNode[T]) Value() T {
	return n.value
}

// Left devuelve el hijo izquierdo del nodo.
func (n *AVLNode[T]) Left() *AVLNode[T] {
	return n.left
}

// Right devuelve el hijo derecho del nodo.
func (n *AVLNode[T]) Right() *AVLNode[T] {
	return n.right
}

// SetLeft asigna el hijo izquierdo del nodo.
func (n *AVLNode[T]) SetLeft(child *AVLNode[T]) {
	n.left = child
}

// SetRight asigna el hijo derecho del nodo.
func (n *AVLNode[T]) SetRight(child *AVLNode[T]) {
	n.right = child
}

// Height devuelve la altura del subárbol enraizado en este nodo.
// Un nodo nil tiene altura -1 (árbol vacío).
func (n *AVLNode[T]) Height() int {
	if n == nil {
		return -1
	}
	return n.height
}

// BalanceFactor calcula el factor de balanceo del nodo:
// altura(hijo izquierdo) - altura(hijo derecho).
func (n *AVLNode[T]) BalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

// updateHeight actualiza la altura del nodo basándose en la altura de sus hijos.
func (n *AVLNode[T]) updateHeight() {
	if n == nil {
		return
	}
	lh := n.left.Height()
	rh := n.right.Height()
	if lh > rh {
		n.height = lh + 1
	} else {
		n.height = rh + 1
	}
}
