package list

// DoublyLinkedList implementa List[T] con nodos de enlace doble.
type DoublyLinkedList[T comparable] struct {
	head *nodeD[T]
	tail *nodeD[T]
	size int
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Size() int {
	// Completar
	return 0
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	// Completar
	return false
}

func (l *DoublyLinkedList[T]) Contains(data T) bool {
	// Completar
	return false
}

func (l *DoublyLinkedList[T]) Head() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *DoublyLinkedList[T]) Tail() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *DoublyLinkedList[T]) Prepend(data T) {
	// Completar
}

func (l *DoublyLinkedList[T]) Append(data T) {
	// Completar
}

func (l *DoublyLinkedList[T]) InsertAfter(target, data T) bool {
	// Completar
	return false
}

func (l *DoublyLinkedList[T]) InsertBefore(target, data T) bool {
	// Completar
	return false
}

func (l *DoublyLinkedList[T]) RemoveFirst() bool {
	// Completar
	return false
}

func (l *DoublyLinkedList[T]) RemoveLast() bool {
	// Completar
	return false
}

func (l *DoublyLinkedList[T]) Remove(data T) bool {
	// Completar
	return false
}

func (l *DoublyLinkedList[T]) Values() []T {
	// Completar
	return nil
}

func (l *DoublyLinkedList[T]) Clear() {
	// Completar
}

func (l *DoublyLinkedList[T]) String() string {
	// Completar
	return "[]"
}
