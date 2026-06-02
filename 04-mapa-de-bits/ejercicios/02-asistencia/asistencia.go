package asistencia

type Asistencias struct {
}

func NewAsistencias(cantAlumnos uint, cantClases uint8) *Asistencias {
	return nil
}

func (a *Asistencias) RegistrarAsistencia(alumno uint, clase uint8) {}

func (a *Asistencias) Asistio(alumno uint, clase uint8) bool {
	return false
}

func (a *Asistencias) ListarClase(clase uint8) []uint {
	return nil
}

func (a *Asistencias) ListarAlumno(alumno uint) []uint8 {
	return nil
}

func (a *Asistencias) ListarAsistencias() []uint8 {
	return nil
}

func (a *Asistencias) ListarAsistenciaPerfecta() []uint {
	return nil
}
