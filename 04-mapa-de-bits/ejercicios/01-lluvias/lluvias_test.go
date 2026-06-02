package lluvias

import (
	"testing"
)

func TestRegistrar(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, Enero)
	lluvias.Registrar(2, Enero)
	lluvias.Registrar(10, Enero)
	lluvias.Registrar(1, Enero)
	lluvias.Registrar(87, Enero)

	if !lluvias.Llovio(1, Enero) {
		t.Error("Llovio(1, Enero) = false, esperaba true")
	}
	if !lluvias.Llovio(2, Enero) {
		t.Error("Llovio(2, Enero) = false, esperaba true")
	}
	if !lluvias.Llovio(10, Enero) {
		t.Error("Llovio(10, Enero) = false, esperaba true")
	}
	if lluvias.Llovio(87, Enero) {
		t.Error("Llovio(87, Enero) = true, esperaba false (día inválido)")
	}
}

func TestNoContaminaMeses(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, Enero)
	if lluvias.Llovio(2, Enero) {
		t.Error("Llovio(2, Enero) = true, esperaba false")
	}
	if !lluvias.Llovio(1, Enero) {
		t.Error("Llovio(1, Enero) = false, esperaba true")
	}
	lluvias.Registrar(1, Diciembre)
	if !lluvias.Llovio(1, Diciembre) {
		t.Error("Llovio(1, Diciembre) = false, esperaba true")
	}
}

func TestListarMes(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, Enero)
	lluvias.Registrar(2, Enero)
	lluvias.Registrar(10, Enero)
	lluvias.Registrar(1, Enero)
	lluvias.Registrar(9, Enero)
	lluvias.Registrar(25, Enero)

	dias := lluvias.ListarMes(Enero)
	if len(dias) != 5 {
		t.Fatalf("ListarMes() = %v, longitud %d, esperaba 5", dias, len(dias))
	}
	esperados := []uint8{1, 2, 9, 10, 25}
	for _, d := range esperados {
		if !contiene(dias, d) {
			t.Errorf("ListarMes() no contiene día %d, resultado: %v", d, dias)
		}
	}
}

func TestListarDiasEnAmbosMeses(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, Enero)
	lluvias.Registrar(2, Enero)
	lluvias.Registrar(10, Enero)
	lluvias.Registrar(1, Marzo)
	lluvias.Registrar(10, Marzo)
	lluvias.Registrar(25, Mayo)
	lluvias.Registrar(25, Noviembre)

	dias := lluvias.ListarDiasEnAmbosMeses(Enero, Marzo)
	if len(dias) != 2 {
		t.Fatalf("ListarDiasEnAmbosMeses(Enero, Marzo) = %v, longitud %d, esperaba 2", dias, len(dias))
	}
	if !contiene(dias, 1) {
		t.Errorf("ListarDiasEnAmbosMeses() no contiene día 1, resultado: %v", dias)
	}
	if !contiene(dias, 10) {
		t.Errorf("ListarDiasEnAmbosMeses() no contiene día 10, resultado: %v", dias)
	}

	dias = lluvias.ListarDiasEnAmbosMeses(Noviembre, Mayo)
	if len(dias) != 1 {
		t.Fatalf("ListarDiasEnAmbosMeses(Noviembre, Mayo) = %v, longitud %d, esperaba 1", dias, len(dias))
	}
	if !contiene(dias, 25) {
		t.Errorf("ListarDiasEnAmbosMeses() no contiene día 25, resultado: %v", dias)
	}
}

func contiene(s []uint8, v uint8) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}
