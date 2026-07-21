package list

// Nodo para lista enlazada simple.
type nodeS[T any] struct {
	data T
	next *nodeS[T]
}

// Nodo para lista enlazada doble.
type nodeD[T any] struct {
	data T
	prev *nodeD[T]
	next *nodeD[T]
}
