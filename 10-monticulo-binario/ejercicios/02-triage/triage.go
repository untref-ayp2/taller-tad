package triage

// Paciente representa un paciente en el sistema de triaje.
type Paciente struct {
	Nombre   string
	Gravedad int
	Llegada  int
}

// Atender ordena los pacientes según gravedad (menor = más urgente)
// y, en caso de igual gravedad, por orden de llegada (menor = antes).
func Atender(pacientes []Paciente) []Paciente {
	// Completar
	return nil
}
