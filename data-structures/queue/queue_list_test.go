package queue

import "testing"

func TestNewListQueueIsEmpty(t *testing.T) {
	q := NewListQueue[int]()
	if !q.IsEmpty() {
		t.Error("esperaba cola vacía recién creada")
	}
}

func TestListQueueEnqueueDequeue(t *testing.T) {
	q := NewListQueue[int]()
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
}

func TestListQueueDequeueWhenEmpty(t *testing.T) {
	q := NewListQueue[int]()
	_, err := q.Dequeue()
	if err == nil {
		t.Error("esperaba error al hacer Dequeue de cola vacía")
	}
}

func TestListQueueFront(t *testing.T) {
	q := NewListQueue[string]()
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

func TestListQueueFrontWhenEmpty(t *testing.T) {
	q := NewListQueue[int]()
	_, err := q.Front()
	if err == nil {
		t.Error("esperaba error al hacer Front de cola vacía")
	}
}
