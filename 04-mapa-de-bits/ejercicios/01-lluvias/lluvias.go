package lluvias

import "github.com/untref-ayp2/data-structures/bitmap"

type Mes uint8

const (
	Enero Mes = iota
	Febrero
	Marzo
	Abril
	Mayo
	Junio
	Julio
	Agosto
	Septiembre
	Octubre
	Noviembre
	Diciembre
)

func validaFecha(d uint8, m Mes) bool {
	switch m {
	case Abril, Junio, Septiembre, Noviembre:
		return d >= 1 && d <= 30
	case Febrero:
		return d >= 1 && d <= 29
	default:
		return d >= 1 && d <= 31
	}
}

type Lluvias struct {
	meses [12]*bitmap.BitMap
}

func NewLluvias() *Lluvias {
	return nil
}

func (l *Lluvias) Registrar(d uint8, m Mes) {}

func (l *Lluvias) Desregistrar(d uint8, m Mes) {}

func (l *Lluvias) Llovio(d uint8, m Mes) bool {
	return false
}

func (l *Lluvias) ListarMes(m Mes) []uint8 {
	return nil
}

func (l *Lluvias) ListarDiasEnAmbosMeses(m1, m2 Mes) []uint8 {
	return nil
}
