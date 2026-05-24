package ejercicios

import "testing"

func TestPlaylistAddAndCurrent(t *testing.T) {
	p := NewPlaylist()
	if p.Current() != "" {
		t.Error("playlist vacía debería tener Current vacío")
	}

	p.Add("Canción 1")
	if p.Current() != "Canción 1" {
		t.Errorf("esperaba 'Canción 1', obtuve %q", p.Current())
	}
}

func TestPlaylistNext(t *testing.T) {
	p := NewPlaylist()
	p.Add("A")
	p.Add("B")
	p.Add("C")

	if p.Current() != "A" {
		t.Fatalf("esperaba 'A', obtuve %q", p.Current())
	}

	p.Next()
	if p.Current() != "B" {
		t.Errorf("esperaba 'B', obtuve %q", p.Current())
	}

	p.Next()
	if p.Current() != "C" {
		t.Errorf("esperaba 'C', obtuve %q", p.Current())
	}

	// Vuelve al principio (circular)
	p.Next()
	if p.Current() != "A" {
		t.Errorf("esperaba 'A' (circular), obtuve %q", p.Current())
	}
}

func TestPlaylistPrev(t *testing.T) {
	p := NewPlaylist()
	p.Add("A")
	p.Add("B")
	p.Add("C")

	p.Next()
	if p.Current() != "B" {
		t.Fatalf("esperaba 'B', obtuve %q", p.Current())
	}

	p.Prev()
	if p.Current() != "A" {
		t.Errorf("esperaba 'A', obtuve %q", p.Current())
	}

	// Retroceder desde A va al final (circular)
	p.Prev()
	if p.Current() != "C" {
		t.Errorf("esperaba 'C' (circular), obtuve %q", p.Current())
	}
}
