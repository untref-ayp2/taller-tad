package list

// CircularLinkedList implementa List[T] con enlaces dobles y comportamiento circular.
type CircularLinkedList[T comparable] struct {
	// Completar
}

// NewCircularLinkedList crea una nueva lista vacía.
func NewCircularLinkedList[T comparable]() *CircularLinkedList[T] {
	// Completar
	return nil
}

func (l *CircularLinkedList[T]) Size() int {
	// Completar
	return 0
}

func (l *CircularLinkedList[T]) IsEmpty() bool {
	// Completar
	return false
}

func (l *CircularLinkedList[T]) Contains(data T) bool {
	// Completar
	return false
}

func (l *CircularLinkedList[T]) Head() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *CircularLinkedList[T]) Tail() (T, bool) {
	// Completar
	var zero T
	return zero, false
}

func (l *CircularLinkedList[T]) Prepend(data T) {
	// Completar
}

func (l *CircularLinkedList[T]) Append(data T) {
	// Completar
}

func (l *CircularLinkedList[T]) InsertAfter(target, data T) bool {
	// Completar
	return false
}

func (l *CircularLinkedList[T]) InsertBefore(target, data T) bool {
	// Completar
	return false
}

func (l *CircularLinkedList[T]) RemoveFirst() bool {
	// Completar
	return false
}

func (l *CircularLinkedList[T]) RemoveLast() bool {
	// Completar
	return false
}

func (l *CircularLinkedList[T]) Remove(data T) bool {
	// Completar
	return false
}

func (l *CircularLinkedList[T]) Values() []T {
	// Completar
	return nil
}

func (l *CircularLinkedList[T]) Clear() {
	// Completar
}

func (l *CircularLinkedList[T]) String() string {
	// Completar
	return ""
}

//// Se agregan primitivas Exclusivas para la lista circular

// Forward avanza la cabeza n posiciones en el sentido de next
// (hacia adelante). Si n <= 0, no hace nada.
func (l *CircularLinkedList[T]) Forward(n int) {
	// Completar
}

// Backward retrocede la cabeza n posiciones en el sentido de prev
// (hacia atrás). Si n <= 0, no hace nada.
func (l *CircularLinkedList[T]) Backward(n int) {
	// Completar
}
