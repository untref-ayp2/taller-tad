package triage

import "testing"

func checkNombres(t *testing.T, result []Paciente, expected []string) {
	t.Helper()
	if len(result) != len(expected) {
		t.Fatalf("Cantidad esperada %d, obtuve %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i].Nombre != v {
			t.Errorf("Posición %d: esperaba %s, obtuve %s", i, v, result[i].Nombre)
		}
	}
}

func TestAtenderOrdenCorrecto(t *testing.T) {
	pacientes := []Paciente{
		{"Ana", 3, 1},
		{"Juan", 1, 2},
		{"Luis", 2, 3},
		{"Eva", 1, 4},
	}

	result := Atender(pacientes)
	checkNombres(t, result, []string{"Juan", "Eva", "Luis", "Ana"})
}

func TestAtenderMismaGravedad(t *testing.T) {
	pacientes := []Paciente{
		{"Pedro", 2, 5},
		{"Maria", 2, 2},
		{"Jorge", 2, 8},
	}

	result := Atender(pacientes)
	checkNombres(t, result, []string{"Maria", "Pedro", "Jorge"})
}

func TestAtenderVacio(t *testing.T) {
	result := Atender([]Paciente{})
	if len(result) != 0 {
		t.Error("Resultado debería estar vacío para entrada vacía")
	}
}

func TestAtenderUnPaciente(t *testing.T) {
	pacientes := []Paciente{
		{"Sofia", 1, 1},
	}

	result := Atender(pacientes)
	if len(result) != 1 || result[0].Nombre != "Sofia" {
		t.Errorf("Esperaba Sofia, obtuve %+v", result)
	}
}

func TestAtenderGravedadInversa(t *testing.T) {
	pacientes := []Paciente{
		{"A", 5, 1},
		{"B", 4, 2},
		{"C", 3, 3},
		{"D", 2, 4},
		{"E", 1, 5},
	}

	result := Atender(pacientes)
	checkNombres(t, result, []string{"E", "D", "C", "B", "A"})
}

func TestAtenderMismoOrdenLlegadaGravedad(t *testing.T) {
	pacientes := []Paciente{
		{"Luis", 1, 2},
		{"Ana", 1, 2},
	}

	result := Atender(pacientes)

	seen := make(map[string]bool)
	for _, p := range result {
		seen[p.Nombre] = true
	}
	if !seen["Luis"] || !seen["Ana"] {
		t.Error("Deben aparecer ambos pacientes")
	}
}
