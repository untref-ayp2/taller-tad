package list

// SentinelLinkedList implementa List[T] con enlaces dobles y nodos centinela.
type SentinelLinkedList[T comparable] struct {
	// Completar
}

// NewSentinelLinkedList crea una nueva lista vacía.
func NewSentinelLinkedList[T comparable]() *SentinelLinkedList[T] {
	// Completar
	return nil
}

func (l *SentinelLinkedList[T]) Size() int {
	// Completar
	return 0
}

func (l *SentinelLinkedList[T]) IsEmpty() bool {
	// Completar
	return false
}

func (l *SentinelLinkedList[T]) Contains(data T) bool {
	// Completar
	return false
}

func (l *SentinelLinkedList[T]) Head() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *SentinelLinkedList[T]) Tail() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *SentinelLinkedList[T]) Prepend(data T) {
	// Completar
}

func (l *SentinelLinkedList[T]) Append(data T) {
	// Completar
}

func (l *SentinelLinkedList[T]) InsertAfter(target, data T) bool {
	// Completar
	return false
}

func (l *SentinelLinkedList[T]) InsertBefore(target, data T) bool {
	// Completar
	return false
}

func (l *SentinelLinkedList[T]) RemoveFirst() bool {
	// Completar
	return false
}

func (l *SentinelLinkedList[T]) RemoveLast() bool {
	// Completar
	return false
}

func (l *SentinelLinkedList[T]) Remove(data T) bool {
	// Completar
	return false
}

func (l *SentinelLinkedList[T]) Values() []T {
	// Completar
	return nil
}

func (l *SentinelLinkedList[T]) Clear() {
	// Completar
}

func (l *SentinelLinkedList[T]) String() string {
	// Completar
	return ""
}
