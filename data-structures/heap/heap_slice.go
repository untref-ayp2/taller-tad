package heap

import "fmt"

// SliceHeap implementa Heap[T] sobre un slice.
type SliceHeap[T any] struct {
	data []T
	cmp  func(T, T) int
	dir  int // 1 min-heap, -1 max-heap
}

// NewMinHeap crea un SliceHeap que mantiene el mínimo en la cima.
func NewMinHeap[T any](cmp func(T, T) int) *SliceHeap[T] {
	// Completar
	return nil
}

// NewMaxHeap crea un SliceHeap que mantiene el máximo en la cima.
func NewMaxHeap[T any](cmp func(T, T) int) *SliceHeap[T] {
	// Completar
	return nil
}

func (h *SliceHeap[T]) Insert(element T) {
	// Completar
}

func (h *SliceHeap[T]) Remove() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

func (h *SliceHeap[T]) Top() (T, error) {
	// Completar
	var zero T
	return zero, nil
}

func (h *SliceHeap[T]) Size() int {
	// Completar
	return 0
}

func (h *SliceHeap[T]) IsEmpty() bool {
	// Completar
	return false
}

func (h *SliceHeap[T]) upHeap(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if h.dir*h.cmp(h.data[parent], h.data[idx]) > 0 {
			h.data[parent], h.data[idx] = h.data[idx], h.data[parent]
			idx = parent
		} else {
			break
		}
	}
}

func (h *SliceHeap[T]) downHeap(idx int) {
	n := len(h.data)
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		best := idx

		if left < n && h.dir*h.cmp(h.data[best], h.data[left]) > 0 {
			best = left
		}
		if right < n && h.dir*h.cmp(h.data[best], h.data[right]) > 0 {
			best = right
		}
		if best == idx {
			break
		}
		h.data[best], h.data[idx] = h.data[idx], h.data[best]
		idx = best
	}
}

// ensure SliceHeap implements Heap
var _ Heap[any] = (*SliceHeap[any])(nil)

// ErrorHeapVacio se retorna cuando se opera sobre un heap vacío.
var ErrorHeapVacio = fmt.Errorf("el montículo está vacío")
