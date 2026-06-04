# Traductor Español-Inglés

Implementar un diccionario bilingüe usando `HashTable[string, string]` de
`data-structures/hashtable`.

Completar los métodos de `Diccionario`:

```go
func NuevoDiccionario() *Diccionario
func (d *Diccionario) Agregar(espanol, ingles string)
func (d *Diccionario) Traducir(palabra string) string
func (d *Diccionario) TraducirOracion(oracion string) string
func (d *Diccionario) ListarPalabras() []string
func (d *Diccionario) Cantidad() int
```

- `Traducir` devuelve la palabra original si no encuentra la traducción.
- `TraducirOracion` separa por espacios. Las palabras no encontradas
  se dejan sin traducir.
- `Agregar` actualiza la traducción si la palabra ya existe.
