package ejercicios

// UndoRedo implementa un historial de acciones con deshacer/rehacer.
type UndoRedo struct {
	// Completar
}

// NewUndoRedo crea un historial vacío.
func NewUndoRedo() *UndoRedo {
	// Completar
	return nil
}

// Do ejecuta una nueva acción. Las acciones rehechas se descartan.
func (u *UndoRedo) Do(accion string) {
	// Completar
}

// Undo deshace la última acción. Devuelve la acción deshecha.
// Si no hay acciones para deshacer devuelve "".
func (u *UndoRedo) Undo() string {
	// Completar
	return ""
}

// Redo rehace la última acción deshecha. Devuelve la acción rehecha.
// Si no hay acciones para rehacer devuelve "".
func (u *UndoRedo) Redo() string {
	// Completar
	return ""
}

// Current devuelve la acción actual (la última ejecutada o rehecha).
// Si no hay acciones devuelve "".
func (u *UndoRedo) Current() string {
	// Completar
	return ""
}
