package list

// SinglyLinkedList implementa List[T] con nodos de enlace simple.
type SinglyLinkedList[T comparable] struct {
	head *nodeS[T]
	tail *nodeS[T]
	size int
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) Size() int {
	// Completar
	return 0
}

func (l *SinglyLinkedList[T]) IsEmpty() bool {
	// Completar
	return false
}

func (l *SinglyLinkedList[T]) Contains(data T) bool {
	// Completar
	return false
}

func (l *SinglyLinkedList[T]) Head() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *SinglyLinkedList[T]) Tail() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *SinglyLinkedList[T]) Prepend(data T) {
	// Completar
}

func (l *SinglyLinkedList[T]) Append(data T) {
	// Completar
}

func (l *SinglyLinkedList[T]) InsertAfter(target, data T) bool {
	// Completar
	return false
}

func (l *SinglyLinkedList[T]) InsertBefore(target, data T) bool {
	// Completar
	return false
}

func (l *SinglyLinkedList[T]) RemoveFirst() bool {
	// Completar
	return false
}

func (l *SinglyLinkedList[T]) RemoveLast() bool {
	// Completar
	return false
}

func (l *SinglyLinkedList[T]) Remove(data T) bool {
	// Completar
	return false
}

func (l *SinglyLinkedList[T]) Values() []T {
	// Completar
	return nil
}

func (l *SinglyLinkedList[T]) Clear() {
	// Completar
}

func (l *SinglyLinkedList[T]) String() string {
	// Completar
	return "[]"
}
