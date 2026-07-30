package merge_listas

import (
	"github.com/untref-ayp2/data-structures/heap"
	"github.com/untref-ayp2/data-structures/list"
	"github.com/untref-ayp2/data-structures/priorityqueue"
)

type elemento[T any] struct {
	valor    T
	listaIdx int
}

// MergeKListas fusiona K listas ordenadas en una sola lista ordenada
// usando una cola de prioridad. Cada lista interna está ordenada según cmp.
func MergeKListas[T comparable](listas []list.List[T], cmp func(T, T) int) list.List[T] {
	h := heap.NewMinHeap[elemento[T]](func(a, b elemento[T]) int {
		return cmp(a.valor, b.valor)
	})
	pq := priorityqueue.NewPriorityQueue[elemento[T]](h)
	_ = pq

	result := list.NewSinglyLinkedList[T]()

	// Completar:
	// 1. Convertí cada lista a slice con l.Values()
	// 2. Insertá el primer elemento de cada slice en la PQ
	// 3. Mientras la PQ no esté vacía, extraé el mínimo,
	//    agregalo a result y si la lista de origen tiene más
	//    elementos, insertá el siguiente.

	return result
}
