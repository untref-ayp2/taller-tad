package tree

// BinaryTree representa un árbol binario genérico.
type BinaryTree[T any] struct {
	Root *TreeNode[T]
}

// NewBinaryTree crea un árbol binario vacío.
func NewBinaryTree[T any]() *BinaryTree[T] {
	// Completar
	return nil
}

// Height devuelve la altura del árbol.
func (bt *BinaryTree[T]) Height() int {
	// Completar
	return 0
}

// Size devuelve la cantidad de nodos en el árbol.
func (bt *BinaryTree[T]) Size() int {
	// Completar
	return 0
}

// InorderTraversal devuelve un slice con los valores en recorrido inorder.
func (bt *BinaryTree[T]) InorderTraversal() []T {
	// Completar
	return nil
}

// PreorderTraversal devuelve un slice con los valores en recorrido preorder.
func (bt *BinaryTree[T]) PreorderTraversal() []T {
	// Completar
	return nil
}

// PostorderTraversal devuelve un slice con los valores en recorrido postorder.
func (bt *BinaryTree[T]) PostorderTraversal() []T {
	// Completar
	return nil
}
