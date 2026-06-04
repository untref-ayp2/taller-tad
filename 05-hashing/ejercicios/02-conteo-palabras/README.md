# Conteo de Frecuencias

Implementar un contador de frecuencias usando `HashTable[string, int]` de
`data-structures/hashtable`.

Completar los métodos de `ContadorDePalabras`:

```go
func NuevoContador() *ContadorDePalabras
func (c *ContadorDePalabras) AgregarTexto(texto string)
func (c *ContadorDePalabras) Frecuencia(palabra string) int
func (c *ContadorDePalabras) TotalPalabras() int
func (c *ContadorDePalabras) PalabrasUnicas() int
func (c *ContadorDePalabras) TopN(n int) []ParFrecuencia
func (c *ContadorDePalabras) ListarPalabras() []string
```

- No distingue mayúsculas/minúsculas.
- Las palabras se separan por espacios.
- `TopN` debe devolver las `n` palabras más frecuentes, ordenadas de mayor
  a menor frecuencia.
