package list

// List define las operaciones que debe implementar una lista enlazada.
type List[T comparable] interface {
	// Consulta
	Size() int
	IsEmpty() bool
	Contains(data T) bool
	Head() (T, bool)
	Tail() (T, bool)

	// Inserción
	Prepend(data T)
	Append(data T)
	InsertAfter(target, data T) bool
	InsertBefore(target, data T) bool

	// Eliminación
	RemoveFirst() bool
	RemoveLast() bool
	Remove(data T) bool

	// Recorrido
	Values() []T

	// Utilidad
	Clear()
	String() string
}
