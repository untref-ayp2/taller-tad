package heap

// Heap define las operaciones de un montículo binario.
type Heap[T any] interface {
	Insert(element T)
	Remove() (T, error)
	Top() (T, error)
	Size() int
	IsEmpty() bool
}
