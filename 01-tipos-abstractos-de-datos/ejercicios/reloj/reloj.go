package reloj

// Reloj representa la hora del día en formato de 24 horas.
type Reloj struct {
	hora    int
	minuto  int
	segundo int
}

// NuevoReloj crea un reloj con la hora indicada.
// Valida que hora (0-23), minuto (0-59) y segundo (0-59).
// TODO: implementar
func NuevoReloj(h, m, s int) (*Reloj, error) {
	return nil, nil
}

// AvanzarUnSegundo incrementa el reloj en 1 segundo.
// Maneja correctamente el cambio de minuto, hora y el reinicio
// a 00:00:00 al llegar a 23:59:59.
// TODO: implementar
func (r *Reloj) AvanzarUnSegundo() {
}

// Hora devuelve la hora actual.
// TODO: implementar
func (r *Reloj) Hora() int {
	return 0
}

// Minuto devuelve el minuto actual.
// TODO: implementar
func (r *Reloj) Minuto() int {
	return 0
}

// Segundo devuelve el segundo actual.
// TODO: implementar
func (r *Reloj) Segundo() int {
	return 0
}

// String devuelve la hora formateada como "HH:MM:SS".
// TODO: implementar
func (r *Reloj) String() string {
	return ""
}
