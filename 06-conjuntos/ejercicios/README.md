# Ejercicios: Conjuntos

Antes de comenzar, implementá `MapSet`, `HashTableSet` y `OrderedSet` en tu repositorio
data-structures, paquete `set/`.

Luego completá los siguientes ejercicios usando `Set[T]`:

---

## 1. Ranking por afinidad

Dado un conjunto de usuarios, cada uno con sus intereses (representados como
un `Set[string]`), ordenar el resto de los usuarios según cuántos intereses
comparten con un usuario objetivo.

**Valores**: un mapa `usuario → Set[string]` de intereses.

**Operaciones**:

- `RankingAfinidad(intereses map[string]set.Set[string], objetivo string) []string`
  — devuelve los demás usuarios ordenados por cantidad de intereses compartidos
  con `objetivo` (descendente). En caso de empate, orden alfabético por nombre.

**Invariante**:

- El usuario `objetivo` está en el mapa de `intereses`.
- No se debe incluir al `objetivo` en el resultado.

**Preguntas**:

- ¿Qué operación de `Set` usaste para calcular los intereses compartidos?
- ¿Qué complejidad tiene el ranking completo?
- ¿Cómo cambiaría la solución si los intereses fuesen `Set[int]` en lugar de
  `Set[string]`?

→ `01-afinidad/`

---

## 2. Buscador de papers por palabras clave

Implementar un buscador de papers académicos que permita búsqueda exacta,
relajada y sugerencia de palabras clave.

**Valores**: un mapa `paper → Set[string]` de palabras clave asociadas a cada
paper.

**Operaciones**:

- `BusquedaExacta(papers map[string]set.Set[string], busqueda set.Set[string]) []string`
  — devuelve los papers que contienen **todas** las keywords buscadas.
- `BusquedaRelajada(papers map[string]set.Set[string], busqueda set.Set[string]) []string`
  — devuelve los papers que contienen **al menos una** keyword, ordenados por
  cantidad de coincidencias (descendente).
- `SugerirKeywords(papers map[string]set.Set[string], busqueda set.Set[string]) []string`
  — encuentra los 5 papers con más coincidencias y devuelve las keywords de
  esos papers que **no estaban** en la búsqueda original, sin repeticiones y
  ordenadas alfabéticamente.

**Invariante**:

- Cada paper tiene al menos una keyword.
- No hay papers repetidos en el mapa.

**Preguntas**:

- ¿Qué operaciones de `Set` usaste para cada función?
- `SugerirKeywords` necesita combinar varios conjuntos. ¿Cómo evitás
  repeticiones?
- ¿Qué complejidad tiene cada operación?

→ `02-papers/`

---

## 3. Fusión de directorios

Una empresa tiene dos sucursales, cada una con su propio directorio de
empleados. Se necesita unificar la información usando conjuntos ordenados.

**Valores**: dos `OrderedSet[string]` con nombres de empleados.

**Operaciones**:

- `FusionarDirectorios(a, b set.Set[string]) set.Set[string]` — devuelve un
  nuevo `OrderedSet` con todos los empleados de ambos directorios, ordenados
  alfabéticamente.
- `EmpleadosEnComun(a, b set.Set[string]) set.Set[string]` — devuelve los
  empleados que están en ambos directorios, ordenados alfabéticamente.
- `EmpleadosExclusivos(a, b set.Set[string]) set.Set[string]` — devuelve los
  empleados que están en uno solo de los directorios, ordenados
  alfabéticamente.
- `CoberturaCompleta(a, b set.Set[string]) bool` — verifica si `a` contiene a
  todos los empleados de `b`.

**Invariante**:

- Los nombres de empleados son únicos dentro de cada directorio.
- Los `OrderedSet` recibidos respetan el invariante de conjunto ordenado
  (ordenados según `less`, sin repetidos).

**Preguntas**:

- ¿Qué operaciones de `Set` implementan cada una de estas funciones?
- ¿Por qué `FusionarDirectorios`, `EmpleadosEnComun` y `EmpleadosExclusivos`
  se benefician de usar `OrderedSet` en lugar de `MapSet`?
- ¿Cómo cambiaría la complejidad si los directorios fueran `MapSet` en lugar
  de `OrderedSet`?

→ `03-directorios/`
