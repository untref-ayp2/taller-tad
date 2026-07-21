package queue

import (
	"testing"
)

func TestSliceQueueEnqueueDequeue(t *testing.T) {
	q := NewSliceQueue[int]()
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

func TestSliceQueueDequeueWhenEmpty(t *testing.T) {
	q := NewSliceQueue[int]()

	_, err := q.Dequeue()
	if err == nil {
		t.Error("esperaba error al hacer Dequeue de cola vacía")
	}
}

func TestSliceQueueFront(t *testing.T) {
	q := NewSliceQueue[string]()
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

func TestSliceQueueFrontWhenEmpty(t *testing.T) {
	q := NewSliceQueue[int]()

	_, err := q.Front()
	if err == nil {
		t.Error("esperaba error al hacer Front de cola vacía")
	}
}

func TestSliceQueueIsEmpty(t *testing.T) {
	q := NewSliceQueue[float64]()

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
