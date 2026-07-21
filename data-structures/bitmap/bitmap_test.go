package bitmap

import "testing"

func TestNewBitMap(t *testing.T) {
	sizes := []uint8{8, 32, 64}
	for _, n := range sizes {
		bm, err := NewBitMap(n)
		if err != nil {
			t.Fatalf("NewBitMap(%d) error inesperado: %v", n, err)
		}
		if bm.Size() != n {
			t.Errorf("NewBitMap(%d).Size() = %d, esperaba %d", n, bm.Size(), n)
		}
	}
}

func TestNewBitMapInvalidSize(t *testing.T) {
	sizes := []uint8{0, 1, 7, 9, 16, 31, 128}
	for _, n := range sizes {
		_, err := NewBitMap(n)
		if err == nil {
			t.Errorf("NewBitMap(%d) debería haber fallado", n)
		}
	}
}

func TestOnAndOff(t *testing.T) {
	bm, _ := NewBitMap(8)
	if err := bm.On(3); err != nil {
		t.Fatalf("On(3) error: %v", err)
	}
	ok, err := bm.IsOn(3)
	if err != nil {
		t.Fatalf("IsOn(3) error: %v", err)
	}
	if !ok {
		t.Error("IsOn(3) = false después de On(3), esperaba true")
	}
	bm.Off(3)
	ok, _ = bm.IsOn(3)
	if ok {
		t.Error("IsOn(3) = true después de Off(3), esperaba false")
	}
}

func TestOutOfRange(t *testing.T) {
	bm, _ := NewBitMap(8)
	if err := bm.On(8); err == nil {
		t.Error("On(8) debería haber fallado para tamaño 8")
	}
	if _, err := bm.IsOn(8); err == nil {
		t.Error("IsOn(8) debería haber fallado para tamaño 8")
	}
	if err := bm.Off(15); err == nil {
		t.Error("Off(15) debería haber fallado para tamaño 8")
	}
}

func TestCountOn(t *testing.T) {
	bm, _ := NewBitMap(32)
	if bm.CountOn() != 0 {
		t.Errorf("CountOn() para bitmap vacío = %d, esperaba 0", bm.CountOn())
	}
	bm.On(0)
	if bm.CountOn() != 1 {
		t.Errorf("CountOn() después de On(0) = %d, esperaba 1", bm.CountOn())
	}
	bm.On(5)
	if bm.CountOn() != 2 {
		t.Errorf("CountOn() después de On(5) = %d, esperaba 2", bm.CountOn())
	}
	bm.Off(0)
	if bm.CountOn() != 1 {
		t.Errorf("CountOn() después de Off(0) = %d, esperaba 1", bm.CountOn())
	}
}

func TestString(t *testing.T) {
	bm, _ := NewBitMap(8)
	if s := bm.String(); s != "00000000" {
		t.Errorf("String() = %q, esperaba %q", s, "00000000")
	}
	bm.On(0)
	if s := bm.String(); s != "00000001" {
		t.Errorf("String() después de On(0) = %q, esperaba %q", s, "00000001")
	}
}

func TestStringSize(t *testing.T) {
	bm, _ := NewBitMap(32)
	if s := bm.String(); len(s) != 32 {
		t.Errorf("String() length = %d para tamaño 32, esperaba 32", len(s))
	}
}

func TestIsOff(t *testing.T) {
	bm, _ := NewBitMap(8)
	ok, err := bm.IsOff(0)
	if err != nil {
		t.Fatalf("IsOff(0) error: %v", err)
	}
	if !ok {
		t.Error("IsOff(0) = false para bitmap vacío, esperaba true")
	}
	bm.On(0)
	ok, _ = bm.IsOff(0)
	if ok {
		t.Error("IsOff(0) = true después de On(0), esperaba false")
	}
}
