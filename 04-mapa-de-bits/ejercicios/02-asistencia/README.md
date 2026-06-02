# Asistencias

Implementar un registro de asistencia para un curso que tiene `N` alumnos y `M`
clases. Usar el `BitMap` de `data-structures/bitmap` como contenedor interno.

Completar los métodos de `Asistencias`:

```go
func NewAsistencias(cantAlumnos uint, cantClases uint8) *Asistencias
func (a *Asistencias) RegistrarAsistencia(alumno uint, clase uint8)
func (a *Asistencias) Asistio(alumno uint, clase uint8) bool
func (a *Asistencias) ListarClase(clase uint8) []uint
func (a *Asistencias) ListarAlumno(alumno uint) []uint8
func (a *Asistencias) ListarAsistencias() []uint8
func (a *Asistencias) ListarAsistenciaPerfecta() []uint
```

- `ListarClase(clase)` devuelve los números de alumno que asistieron a esa clase.
- `ListarAlumno(alumno)` devuelve las clases a las que asistió ese alumno.
- `ListarAsistencias()` devuelve las clases a las que asistieron **todos** los alumnos.
- `ListarAsistenciaPerfecta()` devuelve los alumnos que asistieron a **todas** las clases.
