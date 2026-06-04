# Papers

Implementar un buscador de papers académicos que permita búsqueda exacta,
relajada y sugerencia de palabras clave.

Completar las funciones:

```go
func BusquedaExacta(papers map[string]set.Set[string], busqueda set.Set[string]) []string
func BusquedaRelajada(papers map[string]set.Set[string], busqueda set.Set[string]) []string
func SugerirKeywords(papers map[string]set.Set[string], busqueda set.Set[string]) []string
```

- `BusquedaExacta`: el paper debe contener **todas** las keywords.
- `BusquedaRelajada`: el paper debe contener **al menos una** keyword,
  ordenado por cantidad de coincidencias descendente.
- `SugerirKeywords`: tomar los 5 papers con más coincidencias y devolver las
  keywords que no estaban en la búsqueda original, sin repetidas, ordenadas
  alfabéticamente.
