# Directorios

Una empresa tiene dos sucursales, cada una con su propio directorio de
empleados. Unificar la información usando conjuntos ordenados.

Completar las funciones:

```go
func FusionarDirectorios(a, b set.Set[string]) set.Set[string]
func EmpleadosEnComun(a, b set.Set[string]) set.Set[string]
func EmpleadosExclusivos(a, b set.Set[string]) set.Set[string]
func CoberturaCompleta(a, b set.Set[string]) bool
```

- Usar `OrderedSet` (con orden alfabético) como tipo concreto de `a` y `b`.
- `FusionarDirectorios`: `Union`.
- `EmpleadosEnComun`: `Intersection`.
- `EmpleadosExclusivos`: `SymmetricDifference`.
- `CoberturaCompleta`: `Superset`.
- Todas las operaciones binarias se benefician del merge $O(n+m)$ por estar
  ordenadas.
