package ejercicios

import "testing"

// ----- StackList -----

func TestStackListPushAndPop(t *testing.T) {
	s := NewStackList[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	val, err := s.Pop()
	if err != nil {
		t.Fatalf("Pop(): error inesperado: %v", err)
	}
	if val != 3 {
		t.Errorf("Pop(): esperaba 3, obtuvo %d", val)
	}

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("Pop(): error inesperado: %v", err)
	}
	if val != 2 {
		t.Errorf("Pop(): esperaba 2, obtuvo %d", val)
	}
}

func TestStackListPopWhenEmpty(t *testing.T) {
	s := NewStackList[int]()
	_, err := s.Pop()
	if err == nil {
		t.Error("Pop(): esperaba error en pila vacía")
	}
}

func TestStackListTop(t *testing.T) {
	s := NewStackList[string]()
	s.Push("a")
	s.Push("b")

	val, err := s.Top()
	if err != nil {
		t.Fatalf("Top(): error inesperado: %v", err)
	}
	if val != "b" {
		t.Errorf("Top(): esperaba 'b', obtuvo %s", val)
	}

	// Top no debe eliminar
	val, err = s.Top()
	if err != nil {
		t.Fatalf("Top(): error inesperado: %v", err)
	}
	if val != "b" {
		t.Errorf("Top(): esperaba 'b', obtuvo %s", val)
	}
}

func TestStackListIsEmpty(t *testing.T) {
	s := NewStackList[int]()
	if !s.IsEmpty() {
		t.Error("IsEmpty(): esperaba true en pila vacía")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Error("IsEmpty(): esperaba false después de Push")
	}
	s.Pop()
	if !s.IsEmpty() {
		t.Error("IsEmpty(): esperaba true después de Pop del único elemento")
	}
}

// ----- QueueList -----

func TestQueueListEnqueueDequeue(t *testing.T) {
	q := NewQueueList[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue(): error inesperado: %v", err)
	}
	if val != 1 {
		t.Errorf("Dequeue(): esperaba 1, obtuvo %d", val)
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue(): error inesperado: %v", err)
	}
	if val != 2 {
		t.Errorf("Dequeue(): esperaba 2, obtuvo %d", val)
	}
}

func TestQueueListDequeueWhenEmpty(t *testing.T) {
	q := NewQueueList[int]()
	_, err := q.Dequeue()
	if err == nil {
		t.Error("Dequeue(): esperaba error en cola vacía")
	}
}

func TestQueueListFront(t *testing.T) {
	q := NewQueueList[string]()
	q.Enqueue("a")
	q.Enqueue("b")

	val, err := q.Front()
	if err != nil {
		t.Fatalf("Front(): error inesperado: %v", err)
	}
	if val != "a" {
		t.Errorf("Front(): esperaba 'a', obtuvo %s", val)
	}

	// Front no debe eliminar
	val, err = q.Front()
	if err != nil {
		t.Fatalf("Front(): error inesperado: %v", err)
	}
	if val != "a" {
		t.Errorf("Front(): esperaba 'a', obtuvo %s", val)
	}
}

func TestQueueListIsEmpty(t *testing.T) {
	q := NewQueueList[int]()
	if !q.IsEmpty() {
		t.Error("IsEmpty(): esperaba true en cola vacía")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("IsEmpty(): esperaba false después de Enqueue")
	}
	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("IsEmpty(): esperaba true después de Dequeue del único elemento")
	}
}
