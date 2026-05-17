package fraccion

// Fraccion representa un número racional.
type Fraccion struct {
	numerador   int
	denominador int
}

// NuevaFraccion crea una fracción simplificada.
// El denominador debe ser distinto de cero.
// TODO: implementar
func NuevaFraccion(num, den int) (*Fraccion, error) {
	return nil, nil
}

// Sumar devuelve una nueva fracción con la suma de f y g.
// TODO: implementar
func (f *Fraccion) Sumar(g *Fraccion) *Fraccion {
	return nil
}

// Restar devuelve una nueva fracción con la resta de f y g.
// TODO: implementar
func (f *Fraccion) Restar(g *Fraccion) *Fraccion {
	return nil
}

// Multiplicar devuelve una nueva fracción con el producto de f y g.
// TODO: implementar
func (f *Fraccion) Multiplicar(g *Fraccion) *Fraccion {
	return nil
}

// Dividir devuelve una nueva fracción con el cociente de f y g.
// El denominador de g debe ser distinto de cero.
// TODO: implementar
func (f *Fraccion) Dividir(g *Fraccion) (*Fraccion, error) {
	return nil, nil
}

// Valor devuelve el valor en punto flotante de la fracción.
// TODO: implementar
func (f *Fraccion) Valor() float64 {
	return 0
}

// String devuelve la representación como cadena "num/den".
// TODO: implementar
func (f *Fraccion) String() string {
	return ""
}

// mcd calcula el máximo común divisor usando el algoritmo de Euclides.
// TODO: implementar
func mcd(a, b int) int {
	return 0
}

// simplificar reduce la fracción a su forma más simple dividiendo
// numerador y denominador por su mcd.
// TODO: implementar
func (f *Fraccion) simplificar() {
}
