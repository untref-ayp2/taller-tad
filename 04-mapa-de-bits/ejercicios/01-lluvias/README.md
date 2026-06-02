# Lluvias

Implementar un registro de lluvias anuales, donde por cada mes se registran los
días en que hubo precipitaciones. Usar el `BitMap` de `data-structures/bitmap`
como contenedor interno.

Completar los métodos de `Lluvias`:

```go
func NewLluvias() *Lluvias
func (l *Lluvias) Registrar(d uint8, m Mes)
func (l *Lluvias) Desregistrar(d uint8, m Mes)
func (l *Lluvias) Llovio(d uint8, m Mes) bool
func (l *Lluvias) ListarMes(m Mes) []uint8
func (l *Lluvias) ListarDiasEnAmbosMeses(m1, m2 Mes) []uint8
```

- Los días se numeran del 1 al 31 según el mes.
- Las fechas inválidas deben ignorarse (no causar error, simplemente no hacer nada).
- `ListarMes` devuelve un slice con los días (orden ascendente) en que llovió.
- `ListarDiasEnAmbosMeses` devuelve los días que llovieron en ambos meses.
