package stock

import (
	"testing"
)

func TestNuevoStock(t *testing.T) {
	s := NuevoStock()
	if s == nil {
		t.Fatal("NuevoStock no debería devolver nil")
	}
}

func TestIngresarStock(t *testing.T) {
	s := NuevoStock()
	s.IngresarStock("P001", 10)
	s.IngresarStock("P002", 5)

	stock, err := s.StockActual("P001")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if stock != 10 {
		t.Errorf("esperaba 10, obtuve %d", stock)
	}
}

func TestIngresarStockAcumula(t *testing.T) {
	s := NuevoStock()
	s.IngresarStock("P001", 5)
	s.IngresarStock("P001", 3)

	stock, err := s.StockActual("P001")
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if stock != 8 {
		t.Errorf("esperaba 8, obtuve %d", stock)
	}
}

func TestRetirarStock(t *testing.T) {
	s := NuevoStock()
	s.IngresarStock("P001", 10)

	err := s.RetirarStock("P001", 3)
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}

	stock, _ := s.StockActual("P001")
	if stock != 7 {
		t.Errorf("esperaba 7, obtuve %d", stock)
	}
}

func TestRetirarStockProductoInexistente(t *testing.T) {
	s := NuevoStock()
	err := s.RetirarStock("P999", 1)
	if err == nil {
		t.Error("esperaba error al retirar de producto inexistente")
	}
}

func TestRetirarStockStockInsuficiente(t *testing.T) {
	s := NuevoStock()
	s.IngresarStock("P001", 5)

	err := s.RetirarStock("P001", 10)
	if err == nil {
		t.Error("esperaba error por stock insuficiente")
	}

	stock, _ := s.StockActual("P001")
	if stock != 5 {
		t.Errorf("el stock no debería modificarse tras error, obtuve %d", stock)
	}
}

func TestStockActualInexistente(t *testing.T) {
	s := NuevoStock()
	_, err := s.StockActual("P999")
	if err == nil {
		t.Error("esperaba error para producto inexistente")
	}
}

func TestProductoConMenorStock(t *testing.T) {
	s := NuevoStock()
	s.IngresarStock("P001", 10)
	s.IngresarStock("P002", 3)
	s.IngresarStock("P003", 7)

	codigo, err := s.ProductoConMenorStock()
	if err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if codigo != "P002" {
		t.Errorf("esperaba P002 (stock 3), obtuve %s", codigo)
	}
}

func TestProductoConMenorStockVacio(t *testing.T) {
	s := NuevoStock()
	_, err := s.ProductoConMenorStock()
	if err == nil {
		t.Error("esperaba error con stock vacío")
	}
}

func TestProductosConStockCero(t *testing.T) {
	s := NuevoStock()
	s.IngresarStock("P003", 0)
	s.IngresarStock("P001", 5)
	s.IngresarStock("P002", 0)

	cero := s.ProductosConStockCero()
	if len(cero) != 2 {
		t.Fatalf("esperaba 2 productos con stock cero, obtuve %d", len(cero))
	}
	if cero[0] != "P002" || cero[1] != "P003" {
		t.Errorf("esperaba [P002 P003] ordenados, obtuve %v", cero)
	}
}

func TestProductosConStockCeroNinguno(t *testing.T) {
	s := NuevoStock()
	s.IngresarStock("P001", 5)
	s.IngresarStock("P002", 3)

	cero := s.ProductosConStockCero()
	if len(cero) != 0 {
		t.Errorf("esperaba lista vacía, obtuve %v", cero)
	}
}
