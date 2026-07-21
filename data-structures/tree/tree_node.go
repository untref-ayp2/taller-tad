package tree

// TreeNode representa un nodo de un árbol binario.
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// NewTreeNode crea un nuevo nodo hoja con el valor dado.
func NewTreeNode[T any](value T) *TreeNode[T] {
	// Completar
	return nil
}

// IsLeaf devuelve true si el nodo no tiene hijos.
func (n *TreeNode[T]) IsLeaf() bool {
	// Completar
	return false
}

// Height devuelve la altura del subárbol que tiene a este nodo como raíz.
func (n *TreeNode[T]) Height() int {
	// Completar
	return 0
}

// Size devuelve la cantidad de nodos en el subárbol que tiene a este nodo como raíz.
func (n *TreeNode[T]) Size() int {
	// Completar
	return 0
}

// Preorder agrega los valores del subárbol en recorrido preorder al slice resultado.
func (n *TreeNode[T]) Preorder(result *[]T) {
	// Completar
}

// Inorder agrega los valores del subárbol en recorrido inorder al slice resultado.
func (n *TreeNode[T]) Inorder(result *[]T) {
	// Completar
}

// Postorder agrega los valores del subárbol en recorrido postorder al slice resultado.
func (n *TreeNode[T]) Postorder(result *[]T) {
	// Completar
}
