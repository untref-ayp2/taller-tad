package colacircular

import "testing"

func TestNuevaColaCircular(t *testing.T) {
	c := NewColaCircular[int](5)
	if !c.IsEmpty() {
		t.Error("cola recién creada debería estar vacía")
	}
	if c.IsFull() {
		t.Error("cola recién creada no debería estar llena")
	}
	if c.Size() != 0 {
		t.Errorf("Size() = %d, esperaba 0", c.Size())
	}
}

func TestEnqueueDequeue(t *testing.T) {
	c := NewColaCircular[int](3)

	if err := c.Enqueue(10); err != nil {
		t.Fatalf("Enqueue(10): %v", err)
	}
	if err := c.Enqueue(20); err != nil {
		t.Fatalf("Enqueue(20): %v", err)
	}
	if err := c.Enqueue(30); err != nil {
		t.Fatalf("Enqueue(30): %v", err)
	}

	if !c.IsFull() {
		t.Error("cola con 3 elementos en capacidad 3 debería estar llena")
	}
	if c.Size() != 3 {
		t.Errorf("Size() = %d, esperaba 3", c.Size())
	}

	val, err := c.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error: %v", err)
	}
	if val != 10 {
		t.Errorf("Dequeue() = %d, esperaba 10", val)
	}

	val, err = c.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error: %v", err)
	}
	if val != 20 {
		t.Errorf("Dequeue() = %d, esperaba 20", val)
	}

	val, err = c.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error: %v", err)
	}
	if val != 30 {
		t.Errorf("Dequeue() = %d, esperaba 30", val)
	}

	if !c.IsEmpty() {
		t.Error("cola después de extraer todo debería estar vacía")
	}
}

func TestEnqueueCompleta(t *testing.T) {
	c := NewColaCircular[string](2)
	c.Enqueue("a")
	c.Enqueue("b")

	err := c.Enqueue("c")
	if err == nil {
		t.Error("Enqueue en cola llena debería devolver error")
	}
}

func TestDequeueVacia(t *testing.T) {
	c := NewColaCircular[int](3)
	_, err := c.Dequeue()
	if err == nil {
		t.Error("Dequeue en cola vacía debería devolver error")
	}
}

func TestFront(t *testing.T) {
	c := NewColaCircular[int](3)
	c.Enqueue(42)
	c.Enqueue(99)

	front, err := c.Front()
	if err != nil {
		t.Fatalf("Front() error: %v", err)
	}
	if front != 42 {
		t.Errorf("Front() = %d, esperaba 42", front)
	}

	if c.Size() != 2 {
		t.Errorf("Size() = %d, esperaba 2 después de Front()", c.Size())
	}
}

func TestFrontVacia(t *testing.T) {
	c := NewColaCircular[int](3)
	_, err := c.Front()
	if err == nil {
		t.Error("Front en cola vacía debería devolver error")
	}
}

func TestCircular(t *testing.T) {
	c := NewColaCircular[int](3)

	c.Enqueue(1)
	c.Enqueue(2)
	c.Enqueue(3)

	v1, _ := c.Dequeue()
	v2, _ := c.Dequeue()
	if v1 != 1 || v2 != 2 {
		t.Fatalf("Dequeue: esperaba 1 y 2, obtuve %d y %d", v1, v2)
	}

	c.Enqueue(4)
	c.Enqueue(5)

	if c.Size() != 3 {
		t.Errorf("Size() = %d, esperaba 3", c.Size())
	}

	v3, _ := c.Dequeue()
	if v3 != 3 {
		t.Errorf("Dequeue() = %d, esperaba 3", v3)
	}
	v4, _ := c.Dequeue()
	if v4 != 4 {
		t.Errorf("Dequeue() = %d, esperaba 4", v4)
	}
	v5, _ := c.Dequeue()
	if v5 != 5 {
		t.Errorf("Dequeue() = %d, esperaba 5", v5)
	}
}

func TestSizeCambia(t *testing.T) {
	c := NewColaCircular[int](5)

	c.Enqueue(1)
	if c.Size() != 1 {
		t.Errorf("Size() = %d, esperaba 1", c.Size())
	}

	c.Enqueue(2)
	c.Dequeue()
	if c.Size() != 1 {
		t.Errorf("Size() = %d, esperaba 1 después de Dequeue", c.Size())
	}
}

func TestCapacidadUno(t *testing.T) {
	c := NewColaCircular[int](1)

	if err := c.Enqueue(100); err != nil {
		t.Fatalf("Enqueue(100) en capacidad 1: %v", err)
	}
	if !c.IsFull() {
		t.Error("cola de capacidad 1 con 1 elemento debería estar llena")
	}

	val, _ := c.Dequeue()
	if val != 100 {
		t.Errorf("Dequeue() = %d, esperaba 100", val)
	}
	if !c.IsEmpty() {
		t.Error("cola de capacidad 1 después de Dequeue debería estar vacía")
	}
}
