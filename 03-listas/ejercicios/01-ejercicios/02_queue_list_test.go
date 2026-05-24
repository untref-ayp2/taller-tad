package ejercicios

import (
	"testing"

	"github.com/untref-ayp2/data-structures/queue"
)

func TestQueueListEnqueueDequeue(t *testing.T) {
	q := NewQueueList[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 1 {
		t.Errorf("esperaba 1, obtuve %d", val)
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 2 {
		t.Errorf("esperaba 2, obtuve %d", val)
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 3 {
		t.Errorf("esperaba 3, obtuve %d", val)
	}
}

func TestQueueListDequeueWhenEmpty(t *testing.T) {
	q := NewQueueList[int]()
	_, err := q.Dequeue()
	if err == nil {
		t.Error("esperaba error al hacer Dequeue de cola vacía")
	}
}

func TestQueueListFront(t *testing.T) {
	q := NewQueueList[string]()
	q.Enqueue("a")
	q.Enqueue("b")

	val, err := q.Front()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != "a" {
		t.Errorf("esperaba 'a', obtuve %s", val)
	}

	// Front no debe eliminar el elemento
	val, err = q.Front()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != "a" {
		t.Errorf("esperaba 'a', obtuve %s", val)
	}
}

func TestQueueListFrontWhenEmpty(t *testing.T) {
	q := NewQueueList[int]()
	_, err := q.Front()
	if err == nil {
		t.Error("esperaba error al hacer Front de cola vacía")
	}
}

func TestQueueListIsEmpty(t *testing.T) {
	q := NewQueueList[float64]()

	if !q.IsEmpty() {
		t.Error("esperaba cola vacía recién creada")
	}

	q.Enqueue(1.0)
	if q.IsEmpty() {
		t.Error("esperaba cola no vacía después de Enqueue")
	}

	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("esperaba cola vacía después de Dequeue del único elemento")
	}
}

func TestQueueListSatisfaceInterface(t *testing.T) {
	var q queue.Queue[string] = NewQueueList[string]()
	q.Enqueue("a")
	q.Enqueue("b")
	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue error: %v", err)
	}
	if val != "a" {
		t.Errorf("esperaba 'a', obtuve %s", val)
	}
}
