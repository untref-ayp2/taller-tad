package asistencia

import (
	"testing"
)

func TestRegistrarAsistencia(t *testing.T) {
	a := NewAsistencias(3, 5)
	a.RegistrarAsistencia(0, 0)
	a.RegistrarAsistencia(1, 0)
	a.RegistrarAsistencia(1, 1)
	a.RegistrarAsistencia(1, 2)
	a.RegistrarAsistencia(1, 3)
	a.RegistrarAsistencia(1, 4)
	a.RegistrarAsistencia(2, 4)

	casos := []struct {
		alumno   uint
		clase    uint8
		esperado bool
	}{
		{0, 0, true},
		{1, 0, true},
		{1, 1, true},
		{1, 2, true},
		{1, 3, true},
		{1, 4, true},
		{2, 4, true},
		{0, 1, false},
		{0, 2, false},
		{0, 3, false},
		{0, 4, false},
		{2, 0, false},
		{2, 1, false},
		{2, 2, false},
		{2, 3, false},
	}

	for _, c := range casos {
		if a.Asistio(c.alumno, c.clase) != c.esperado {
			t.Errorf("Asistio(%d, %d) = %v, esperaba %v",
				c.alumno, c.clase, a.Asistio(c.alumno, c.clase), c.esperado)
		}
	}
}

func TestListarClase(t *testing.T) {
	a := NewAsistencias(3, 5)
	a.RegistrarAsistencia(0, 0)
	a.RegistrarAsistencia(1, 0)
	a.RegistrarAsistencia(1, 1)
	a.RegistrarAsistencia(1, 2)
	a.RegistrarAsistencia(1, 3)
	a.RegistrarAsistencia(1, 4)
	a.RegistrarAsistencia(2, 4)

	clase0 := a.ListarClase(0)
	if len(clase0) != 2 {
		t.Fatalf("ListarClase(0) = %v, longitud %d, esperaba 2", clase0, len(clase0))
	}
	if !contieneUint(clase0, 0) {
		t.Errorf("ListarClase(0) no contiene alumno 0, resultado: %v", clase0)
	}
	if !contieneUint(clase0, 1) {
		t.Errorf("ListarClase(0) no contiene alumno 1, resultado: %v", clase0)
	}

	clase1 := a.ListarClase(1)
	if len(clase1) != 1 {
		t.Fatalf("ListarClase(1) = %v, longitud %d, esperaba 1", clase1, len(clase1))
	}
	if !contieneUint(clase1, 1) {
		t.Errorf("ListarClase(1) no contiene alumno 1, resultado: %v", clase1)
	}

	clase4 := a.ListarClase(4)
	if len(clase4) != 2 {
		t.Fatalf("ListarClase(4) = %v, longitud %d, esperaba 2", clase4, len(clase4))
	}
	if !contieneUint(clase4, 1) {
		t.Errorf("ListarClase(4) no contiene alumno 1, resultado: %v", clase4)
	}
	if !contieneUint(clase4, 2) {
		t.Errorf("ListarClase(4) no contiene alumno 2, resultado: %v", clase4)
	}
}

func TestListarAlumno(t *testing.T) {
	a := NewAsistencias(3, 5)
	a.RegistrarAsistencia(0, 0)
	a.RegistrarAsistencia(0, 1)
	a.RegistrarAsistencia(0, 2)
	a.RegistrarAsistencia(0, 3)
	a.RegistrarAsistencia(0, 4)
	a.RegistrarAsistencia(1, 0)
	a.RegistrarAsistencia(1, 1)
	a.RegistrarAsistencia(1, 2)
	a.RegistrarAsistencia(1, 3)
	a.RegistrarAsistencia(1, 4)
	a.RegistrarAsistencia(2, 4)

	clases := a.ListarAlumno(1)
	if len(clases) != 5 {
		t.Fatalf("ListarAlumno(1) = %v, longitud %d, esperaba 5", clases, len(clases))
	}
	esperadas := []uint8{0, 1, 2, 3, 4}
	for _, c := range esperadas {
		if !contieneUint8(clases, c) {
			t.Errorf("ListarAlumno(1) no contiene clase %d, resultado: %v", c, clases)
		}
	}

	clases = a.ListarAlumno(2)
	if len(clases) != 1 {
		t.Fatalf("ListarAlumno(2) = %v, longitud %d, esperaba 1", clases, len(clases))
	}
	if !contieneUint8(clases, 4) {
		t.Errorf("ListarAlumno(2) no contiene clase 4, resultado: %v", clases)
	}
}

func TestListarAsistencias(t *testing.T) {
	a := NewAsistencias(3, 5)
	a.RegistrarAsistencia(0, 0)
	a.RegistrarAsistencia(0, 1)
	a.RegistrarAsistencia(0, 2)
	a.RegistrarAsistencia(0, 3)
	a.RegistrarAsistencia(0, 4)
	a.RegistrarAsistencia(1, 0)
	a.RegistrarAsistencia(1, 1)
	a.RegistrarAsistencia(1, 2)
	a.RegistrarAsistencia(1, 3)
	a.RegistrarAsistencia(1, 4)
	a.RegistrarAsistencia(2, 2)

	clases := a.ListarAsistencias()
	if len(clases) != 1 {
		t.Fatalf("ListarAsistencias() = %v, longitud %d, esperaba 1", clases, len(clases))
	}
	if !contieneUint8(clases, 2) {
		t.Errorf("ListarAsistencias() no contiene clase 2, resultado: %v", clases)
	}
}

func TestListarAsistenciaPerfecta(t *testing.T) {
	a := NewAsistencias(3, 5)
	a.RegistrarAsistencia(0, 0)
	a.RegistrarAsistencia(0, 1)
	a.RegistrarAsistencia(0, 2)
	a.RegistrarAsistencia(0, 3)
	a.RegistrarAsistencia(0, 4)
	a.RegistrarAsistencia(1, 3)
	a.RegistrarAsistencia(1, 4)
	a.RegistrarAsistencia(2, 0)
	a.RegistrarAsistencia(2, 1)
	a.RegistrarAsistencia(2, 3)
	a.RegistrarAsistencia(2, 4)

	alumnos := a.ListarAsistenciaPerfecta()
	if len(alumnos) != 1 {
		t.Fatalf("ListarAsistenciaPerfecta() = %v, longitud %d, esperaba 1", alumnos, len(alumnos))
	}
	if !contieneUint(alumnos, 0) {
		t.Errorf("ListarAsistenciaPerfecta() no contiene alumno 0, resultado: %v", alumnos)
	}
}

func contieneUint(s []uint, v uint) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func contieneUint8(s []uint8, v uint8) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}
