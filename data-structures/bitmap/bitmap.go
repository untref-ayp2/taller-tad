package bitmap

// BitMap representa un mapa de bits parametrizable por tamaño.
// Soporta mapas de bits de 8, 32 y 64 bits.
type BitMap struct {
	bits uint64
	n    uint8
}

// NewBitMap crea un nuevo BitMap con el tamaño especificado (8, 32 o 64).
func NewBitMap(n uint8) (*BitMap, error) {
	// Completar
	return nil, nil
}

// On activa el bit en la posición pos.
func (bm *BitMap) On(pos uint8) error {
	// Completar
	return nil
}

// Off desactiva el bit en la posición pos.
func (bm *BitMap) Off(pos uint8) error {
	// Completar
	return nil
}

// IsOn devuelve true si el bit en la posición pos está activo.
func (bm *BitMap) IsOn(pos uint8) (bool, error) {
	// Completar
	return false, nil
}

// IsOff devuelve true si el bit en la posición pos está inactivo.
func (bm *BitMap) IsOff(pos uint8) (bool, error) {
	// Completar
	return false, nil
}

// CountOn devuelve la cantidad de bits activos.
func (bm *BitMap) CountOn() uint8 {
	// Completar
	return 0
}

// String devuelve una representación en cadena del mapa de bits
// mostrando exactamente n bits (ej: para n=8, "00001100").
func (bm *BitMap) String() string {
	// Completar
	return ""
}

// Bits devuelve el valor interno enmascarado según el tamaño.
func (bm *BitMap) Bits() uint64 {
	// Completar
	return 0
}

// Size devuelve el tamaño del mapa de bits (8, 32 o 64).
func (bm *BitMap) Size() uint8 {
	// Completar
	return 0
}
