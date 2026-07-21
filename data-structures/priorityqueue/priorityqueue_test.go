package priorityqueue

import (
	"testing"

	"github.com/untref-ayp2/data-structures/heap"
)

func intCmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func TestPriorityQueueEnqueueDequeue(t *testing.T) {
	h := heap.NewMinHeap[int](intCmp)
	pq := NewPriorityQueue[int](h)

	pq.Enqueue(5)
	pq.Enqueue(1)
	pq.Enqueue(3)

	if pq.Size() != 3 {
		t.Errorf("Size esperado 3, obtuve %d", pq.Size())
	}

	val, _ := pq.Dequeue()
	if val != 1 {
		t.Errorf("Dequeue esperaba 1, obtuve %d", val)
	}

	val, _ = pq.Dequeue()
	if val != 3 {
		t.Errorf("Dequeue esperaba 3, obtuve %d", val)
	}

	val, _ = pq.Dequeue()
	if val != 5 {
		t.Errorf("Dequeue esperaba 5, obtuve %d", val)
	}

	if !pq.IsEmpty() {
		t.Error("La cola debería estar vacía después de vaciarla")
	}
}

func TestPriorityQueueFront(t *testing.T) {
	h := heap.NewMinHeap[int](intCmp)
	pq := NewPriorityQueue[int](h)
	pq.Enqueue(10)
	pq.Enqueue(5)

	val, _ := pq.Front()
	if val != 5 {
		t.Errorf("Front esperaba 5, obtuve %d", val)
	}
	if pq.Size() != 2 {
		t.Errorf("Front no debe remover, Size esperado 2, obtuve %d", pq.Size())
	}
}

func TestPriorityQueueDequeueEmpty(t *testing.T) {
	h := heap.NewMinHeap[int](intCmp)
	pq := NewPriorityQueue[int](h)

	_, err := pq.Dequeue()
	if err == nil {
		t.Error("Dequeue sobre cola vacía debería retornar error")
	}
}

func TestPriorityQueueFrontEmpty(t *testing.T) {
	h := heap.NewMinHeap[int](intCmp)
	pq := NewPriorityQueue[int](h)

	_, err := pq.Front()
	if err == nil {
		t.Error("Front sobre cola vacía debería retornar error")
	}
}

func TestPriorityQueueIsEmpty(t *testing.T) {
	h := heap.NewMinHeap[int](intCmp)
	pq := NewPriorityQueue[int](h)
	if !pq.IsEmpty() {
		t.Error("Cola nueva debería estar vacía")
	}
	pq.Enqueue(1)
	if pq.IsEmpty() {
		t.Error("Cola con elementos no debería estar vacía")
	}
}

func TestPriorityQueueMaxHeap(t *testing.T) {
	h := heap.NewMaxHeap[int](intCmp)
	pq := NewPriorityQueue[int](h)

	pq.Enqueue(10)
	pq.Enqueue(5)
	pq.Enqueue(20)

	val, _ := pq.Dequeue()
	if val != 20 {
		t.Errorf("Dequeue con max-heap esperaba 20, obtuve %d", val)
	}

	val, _ = pq.Dequeue()
	if val != 10 {
		t.Errorf("Dequeue con max-heap esperaba 10, obtuve %d", val)
	}
}
