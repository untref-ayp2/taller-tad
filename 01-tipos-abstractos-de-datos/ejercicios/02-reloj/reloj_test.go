package reloj

import "testing"

func TestNuevoReloj(t *testing.T) {
	r, err := NuevoReloj(10, 30, 15)
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if r.String() != "10:30:15" {
		t.Errorf("esperado '10:30:15', got '%s'", r.String())
	}
}

func TestHoraInvalida(t *testing.T) {
	_, err := NuevoReloj(24, 0, 0)
	if err == nil {
		t.Error("se esperaba error con hora 24")
	}
}

func TestMinutoInvalido(t *testing.T) {
	_, err := NuevoReloj(0, 60, 0)
	if err == nil {
		t.Error("se esperaba error con minuto 60")
	}
}

func TestSegundoInvalido(t *testing.T) {
	_, err := NuevoReloj(0, 0, 60)
	if err == nil {
		t.Error("se esperaba error con segundo 60")
	}
}

func TestAvanzarUnSegundo(t *testing.T) {
	r, _ := NuevoReloj(10, 30, 15)
	r.AvanzarUnSegundo()
	if r.String() != "10:30:16" {
		t.Errorf("esperado '10:30:16', got '%s'", r.String())
	}
}

func TestAvanzarCambioDeMinuto(t *testing.T) {
	r, _ := NuevoReloj(10, 30, 59)
	r.AvanzarUnSegundo()
	if r.String() != "10:31:00" {
		t.Errorf("esperado '10:31:00', got '%s'", r.String())
	}
}

func TestAvanzarCambioDeHora(t *testing.T) {
	r, _ := NuevoReloj(10, 59, 59)
	r.AvanzarUnSegundo()
	if r.String() != "11:00:00" {
		t.Errorf("esperado '11:00:00', got '%s'", r.String())
	}
}

func TestAvanzarCambioDeDia(t *testing.T) {
	r, _ := NuevoReloj(23, 59, 59)
	r.AvanzarUnSegundo()
	if r.String() != "00:00:00" {
		t.Errorf("esperado '00:00:00', got '%s'", r.String())
	}
}

func TestGetters(t *testing.T) {
	r, _ := NuevoReloj(8, 15, 30)
	if r.Hora() != 8 || r.Minuto() != 15 || r.Segundo() != 30 {
		t.Errorf("getters incorrectos: %d:%d:%d", r.Hora(), r.Minuto(), r.Segundo())
	}
}
