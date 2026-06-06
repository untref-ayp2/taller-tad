package stock

import "github.com/untref-ayp2/data-structures/dictionary"

// Stock gestiona las cantidades de productos en un depósito.
type Stock struct {
	productos dictionary.Dictionary[string, int]
}

// NuevoStock crea un registro de stock vacío.
func NuevoStock() *Stock {
	// Completar
	return nil
}

// IngresarStock agrega cantidad unidades al producto.
// Si el producto no existe, lo crea con esa cantidad.
func (s *Stock) IngresarStock(codigo string, cantidad int) {
	// Completar
}

// RetirarStock descuenta cantidad unidades del producto.
// Error si el producto no existe o si el stock resultante sería negativo.
func (s *Stock) RetirarStock(codigo string, cantidad int) error {
	// Completar
	return nil
}

// StockActual devuelve la cantidad en stock de un producto.
// Error si el producto no existe.
func (s *Stock) StockActual(codigo string) (int, error) {
	// Completar
	return 0, nil
}

// ProductoConMenorStock devuelve el código del producto con menor stock.
// Error si no hay productos registrados.
func (s *Stock) ProductoConMenorStock() (string, error) {
	// Completar
	return "", nil
}

// ProductosConStockCero devuelve los códigos de productos sin stock,
// ordenados alfabéticamente.
func (s *Stock) ProductosConStockCero() []string {
	// Completar
	return nil
}
