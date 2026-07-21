package hashtable

import (
	"testing"
)

func TestChainingPutAndGet(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("uno", 1)
	ht.Put("dos", 2)
	ht.Put("tres", 3)

	val, err := ht.Get("uno")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 1 {
		t.Errorf("esperaba 1, obtuve %d", val)
	}

	val, err = ht.Get("dos")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 2 {
		t.Errorf("esperaba 2, obtuve %d", val)
	}

	val, err = ht.Get("tres")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 3 {
		t.Errorf("esperaba 3, obtuve %d", val)
	}
}

func TestChainingGetNonExistent(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	_, err := ht.Get("inexistente")
	if err == nil {
		t.Error("esperaba error al obtener clave inexistente")
	}
}

func TestChainingUpdateExisting(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("clave", 10)
	ht.Put("clave", 20)

	val, err := ht.Get("clave")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 20 {
		t.Errorf("esperaba 20, obtuve %d", val)
	}
}

func TestChainingDelete(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("a", 1)
	ht.Put("b", 2)

	err := ht.Delete("a")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	_, err = ht.Get("a")
	if err == nil {
		t.Error("esperaba error al obtener clave eliminada")
	}

	val, err := ht.Get("b")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if val != 2 {
		t.Errorf("esperaba 2, obtuve %d", val)
	}
}

func TestChainingDeleteNonExistent(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	err := ht.Delete("inexistente")
	if err == nil {
		t.Error("esperaba error al eliminar clave inexistente")
	}
}

func TestChainingContains(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("clave", 42)

	if !ht.Contains("clave") {
		t.Error("esperaba Contains true para clave existente")
	}
	if ht.Contains("otra") {
		t.Error("esperaba Contains false para clave inexistente")
	}
}

func TestChainingSizeAndIsEmpty(t *testing.T) {
	ht := NewHashTableChaining[string, int]()

	if !ht.IsEmpty() {
		t.Error("esperaba tabla vacía recién creada")
	}
	if ht.Size() != 0 {
		t.Errorf("esperaba size 0, obtuve %d", ht.Size())
	}

	ht.Put("a", 1)
	ht.Put("b", 2)

	if ht.IsEmpty() {
		t.Error("esperaba tabla no vacía después de insertar")
	}
	if ht.Size() != 2 {
		t.Errorf("esperaba size 2, obtuve %d", ht.Size())
	}

	ht.Delete("a")
	if ht.Size() != 1 {
		t.Errorf("esperaba size 1 después de eliminar, obtuve %d", ht.Size())
	}
}

func TestChainingKeys(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("a", 1)
	ht.Put("b", 2)
	ht.Put("c", 3)

	keys := ht.Keys()
	if len(keys) != 3 {
		t.Fatalf("esperaba 3 claves, obtuve %d", len(keys))
	}

	m := make(map[string]bool)
	for _, k := range keys {
		m[k] = true
	}
	if !m["a"] || !m["b"] || !m["c"] {
		t.Error("las claves no coinciden con las insertadas")
	}
}

func TestChainingValues(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("a", 10)
	ht.Put("b", 20)

	values := ht.Values()
	if len(values) != 2 {
		t.Fatalf("esperaba 2 valores, obtuve %d", len(values))
	}

	sum := 0
	for _, v := range values {
		sum += v
	}
	if sum != 30 {
		t.Errorf("suma de valores esperaba 30, obtuve %d", sum)
	}
}

func TestChainingClear(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("a", 1)
	ht.Put("b", 2)
	ht.Clear()

	if !ht.IsEmpty() {
		t.Error("esperaba tabla vacía después de Clear")
	}
	if ht.Size() != 0 {
		t.Errorf("esperaba size 0 después de Clear, obtuve %d", ht.Size())
	}
}

func TestChainingString(t *testing.T) {
	ht := NewHashTableChaining[string, int]()
	ht.Put("a", 1)
	ht.Put("b", 2)

	s := ht.String()
	if s == "{}" {
		t.Error("String no debe devolver {} para tabla con elementos")
	}
	if s != "{a: 1, b: 2}" && s != "{b: 2, a: 1}" {
		t.Errorf("formato inesperado: %s", s)
	}
}

func TestChainingCollision(t *testing.T) {
	ht := NewHashTableChaining[int, string]()

	for i := 0; i < 10; i++ {
		ht.Put(i, "valor")
	}

	for i := 0; i < 10; i++ {
		val, err := ht.Get(i)
		if err != nil {
			t.Fatalf("error al obtener clave %d: %v", i, err)
		}
		if val != "valor" {
			t.Errorf("clave %d: esperaba 'valor', obtuve '%s'", i, val)
		}
	}
}

func TestChainingResize(t *testing.T) {
	ht := NewHashTableChaining[int, int]()

	for i := 0; i < 100; i++ {
		ht.Put(i, i*2)
	}

	if ht.Size() != 100 {
		t.Errorf("esperaba size 100, obtuve %d", ht.Size())
	}

	for i := 0; i < 100; i++ {
		val, err := ht.Get(i)
		if err != nil {
			t.Fatalf("error al obtener clave %d: %v", i, err)
		}
		if val != i*2 {
			t.Errorf("clave %d: esperaba %d, obtuve %d", i, i*2, val)
		}
	}
}
