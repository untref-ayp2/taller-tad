package ejercicios

import "testing"

func TestUndoRedoDoAndCurrent(t *testing.T) {
	u := NewUndoRedo()
	if u.Current() != "" {
		t.Error("UndoRedo vacío debería tener Current vacío")
	}

	u.Do("acción 1")
	if u.Current() != "acción 1" {
		t.Errorf("esperaba 'acción 1', obtuve %q", u.Current())
	}

	u.Do("acción 2")
	if u.Current() != "acción 2" {
		t.Errorf("esperaba 'acción 2', obtuve %q", u.Current())
	}
}

func TestUndoRedoUndo(t *testing.T) {
	u := NewUndoRedo()
	u.Do("a")
	u.Do("b")
	u.Do("c")

	val := u.Undo()
	if val != "c" {
		t.Errorf("Undo esperaba 'c', obtuve %q", val)
	}
	if u.Current() != "b" {
		t.Errorf("Current después de Undo esperaba 'b', obtuve %q", u.Current())
	}

	val = u.Undo()
	if val != "b" {
		t.Errorf("Undo esperaba 'b', obtuve %q", val)
	}
	if u.Current() != "a" {
		t.Errorf("Current después de Undo esperaba 'a', obtuve %q", u.Current())
	}

	val = u.Undo()
	if val != "a" {
		t.Errorf("Undo esperaba 'a', obtuve %q", val)
	}
	if u.Current() != "" {
		t.Errorf("Current después de deshacer todo esperaba '', obtuve %q", u.Current())
	}

	val = u.Undo()
	if val != "" {
		t.Errorf("Undo en vacío esperaba '', obtuve %q", val)
	}
}

func TestUndoRedoRedo(t *testing.T) {
	u := NewUndoRedo()
	u.Do("a")
	u.Do("b")
	u.Do("c")

	u.Undo()
	u.Undo()

	val := u.Redo()
	if val != "b" {
		t.Errorf("Redo esperaba 'b', obtuve %q", val)
	}
	if u.Current() != "b" {
		t.Errorf("Current después de Redo esperaba 'b', obtuve %q", u.Current())
	}

	val = u.Redo()
	if val != "c" {
		t.Errorf("Redo esperaba 'c', obtuve %q", val)
	}
	if u.Current() != "c" {
		t.Errorf("Current después de Redo esperaba 'c', obtuve %q", u.Current())
	}

	val = u.Redo()
	if val != "" {
		t.Errorf("Redo en tope esperaba '', obtuve %q", val)
	}
}

func TestUndoRedoDoClearsRedo(t *testing.T) {
	u := NewUndoRedo()
	u.Do("a")
	u.Do("b")
	u.Do("c")

	u.Undo()
	u.Undo()

	u.Do("x")

	if u.Current() != "x" {
		t.Errorf("Current esperaba 'x', obtuve %q", u.Current())
	}

	val := u.Redo()
	if val != "" {
		t.Errorf("Redo después de nueva acción esperaba '', obtuve %q", val)
	}
}
