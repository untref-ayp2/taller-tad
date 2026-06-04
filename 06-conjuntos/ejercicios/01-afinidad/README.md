# Afinidad

Dado un mapa de usuarios con sus intereses (`Set[string]`), ordenar el resto
de los usuarios según cuántos intereses comparten con un usuario objetivo.

Completar la función `RankingAfinidad`:

```go
func RankingAfinidad(intereses map[string]set.Set[string], objetivo string) []string
```

- Usar `Intersection` para calcular los intereses compartidos.
- El resultado debe ordenarse por cantidad de coincidencias (descendente).
- En caso de empate, orden alfabético por nombre de usuario.
- No incluir al usuario `objetivo` en el resultado.
