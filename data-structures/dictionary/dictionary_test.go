package dictionary

import (
	"testing"

	"github.com/untref-ayp2/data-structures/hashtable"
)

func newTestDict[K comparable, V any]() *HashMapDictionary[K, V] {
	table := hashtable.NewHashTableChaining[K, V]()
	return NewHashMapDictionary[K, V](table)
}

func TestHashMapDictionarySetAndGet(t *testing.T) {
	d := newTestDict[string, int]()
	d.Set("a", 1)
	d.Set("b", 2)

	val, ok := d.Get("a")
	if !ok {
		t.Fatal("Get('a') debería devolver ok = true")
	}
	if val != 1 {
		t.Errorf("esperaba 1, obtuve %d", val)
	}

	val, ok = d.Get("b")
	if !ok {
		t.Fatal("Get('b') debería devolver ok = true")
	}
	if val != 2 {
		t.Errorf("esperaba 2, obtuve %d", val)
	}
}

func TestHashMapDictionaryGetNonExistent(t *testing.T) {
	d := newTestDict[string, int]()
	_, ok := d.Get("nonexistent")
	if ok {
		t.Error("Get de clave inexistente debería devolver ok = false")
	}
}

func TestHashMapDictionarySetUpdate(t *testing.T) {
	d := newTestDict[string, int]()
	d.Set("a", 1)
	d.Set("a", 99)

	val, ok := d.Get("a")
	if !ok {
		t.Fatal("Get('a') debería devolver ok = true")
	}
	if val != 99 {
		t.Errorf("esperaba 99, obtuve %d", val)
	}
}

func TestHashMapDictionaryDelete(t *testing.T) {
	d := newTestDict[string, int]()
	d.Set("a", 1)
	d.Set("b", 2)

	err := d.Delete("a")
	if err != nil {
		t.Fatalf("error inesperado al eliminar 'a': %v", err)
	}

	if d.Contains("a") {
		t.Error("'a' no debería estar después de eliminarla")
	}
	if !d.Contains("b") {
		t.Error("'b' debería seguir estando")
	}
}

func TestHashMapDictionaryDeleteNonExistent(t *testing.T) {
	d := newTestDict[string, int]()
	err := d.Delete("nonexistent")
	if err == nil {
		t.Error("esperaba error al eliminar clave inexistente")
	}
}

func TestHashMapDictionaryContains(t *testing.T) {
	d := newTestDict[string, int]()
	d.Set("a", 1)

	if !d.Contains("a") {
		t.Error("debería contener 'a'")
	}
	if d.Contains("b") {
		t.Error("no debería contener 'b'")
	}
}

func TestHashMapDictionarySize(t *testing.T) {
	d := newTestDict[string, int]()

	if d.Size() != 0 {
		t.Errorf("diccionario vacío debería tener size 0, obtuve %d", d.Size())
	}

	d.Set("a", 1)
	d.Set("b", 2)
	if d.Size() != 2 {
		t.Errorf("esperaba size 2, obtuve %d", d.Size())
	}

	d.Delete("a")
	if d.Size() != 1 {
		t.Errorf("esperaba size 1, obtuve %d", d.Size())
	}
}

func TestHashMapDictionaryKeys(t *testing.T) {
	d := newTestDict[string, int]()
	d.Set("a", 1)
	d.Set("b", 2)
	d.Set("c", 3)

	keys := d.Keys()
	if len(keys) != 3 {
		t.Fatalf("esperaba 3 claves, obtuve %d", len(keys))
	}

	m := make(map[string]bool)
	for _, k := range keys {
		m[k] = true
	}
	if !m["a"] || !m["b"] || !m["c"] {
		t.Error("Keys no contiene todas las claves esperadas")
	}
}

func TestHashMapDictionaryValues(t *testing.T) {
	d := newTestDict[string, int]()
	d.Set("a", 10)
	d.Set("b", 20)

	values := d.Values()
	if len(values) != 2 {
		t.Fatalf("esperaba 2 valores, obtuve %d", len(values))
	}

	sum := 0
	for _, v := range values {
		sum += v
	}
	if sum != 30 {
		t.Errorf("suma de valores esperada 30, obtuve %d", sum)
	}
}

func TestHashMapDictionaryEmptyKeys(t *testing.T) {
	d := newTestDict[string, int]()
	keys := d.Keys()
	if len(keys) != 0 {
		t.Errorf("diccionario vacío debería devolver 0 keys, obtuvo %d", len(keys))
	}
}

func TestHashMapDictionaryEmptyValues(t *testing.T) {
	d := newTestDict[string, int]()
	values := d.Values()
	if len(values) != 0 {
		t.Errorf("diccionario vacío debería devolver 0 values, obtuvo %d", len(values))
	}
}

func TestHashMapDictionaryString(t *testing.T) {
	d := newTestDict[string, int]()
	d.Set("a", 1)
	d.Set("b", 2)

	s := d.String()
	if len(s) == 0 {
		t.Error("String no debería estar vacío")
	}
}

func TestHashMapDictionaryWithIntKeys(t *testing.T) {
	d := newTestDict[int, string]()
	d.Set(1, "uno")
	d.Set(2, "dos")

	val, ok := d.Get(1)
	if !ok {
		t.Fatal("Get(1) debería devolver ok = true")
	}
	if val != "uno" {
		t.Errorf("esperaba 'uno', obtuve %s", val)
	}

	if !d.Contains(2) {
		t.Error("debería contener clave 2")
	}
}
