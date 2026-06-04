# Ejercicios: Tablas de Hash

Antes de comenzar, implementá `HashTableOpenAddressing` y `HashTableChaining`
en tu repositorio
[`data-structures`](https://github.com/untref-ayp2/data-structures)
(el que se creó al aceptar la tarea en GitHub Classroom),
paquete `hashtable/`. Para la función de _hash_ con `K comparable` podés usar
el paquete [`hash/maphash`](https://pkg.go.dev/hash/maphash) de la biblioteca
estándar de Go.

Luego completá los siguientes ejercicios usando `HashTable[string, V]`:

---

## 1. Traductor Español-Inglés

Implementar un diccionario bilingüe español-inglés que permita traducir
palabras y oraciones completas.

**Operaciones**:

- `NuevoDiccionario() *Diccionario` — crea un diccionario vacío
- `Agregar(espanol, ingles string)` — agrega un par español → inglés
  (si la palabra ya existe, actualiza su traducción)
- `Traducir(palabra string) string` — traduce una palabra del español al inglés.
  Si no encuentra la palabra, devuelve la palabra original sin traducir.
- `TraducirOracion(oracion string) string` — traduce una oración completa
  palabra por palabra (separando por espacios). Si una palabra no está en el
  diccionario, se deja sin traducir.
- `ListarPalabras() []string` — devuelve todas las palabras en español registradas
- `Cantidad() int` — devuelve la cantidad de palabras registradas

**Invariante**:

- El diccionario nunca debe contener claves duplicadas (ya lo garantiza la
  tabla de _hash_).

**Preguntas**:

- ¿Qué implementación de `HashTable` elegiste y por qué?
- ¿Qué complejidad tiene `TraducirOracion` en función de la cantidad de
  palabras de la oración?
- ¿Cómo cambiarías el diseño si quisieras dar de alta y baja de idiomas
  dinámicamente (no solo español-inglés)?

→ `01-diccionario/`

---

## 2. Conteo de Frecuencias

Implementar un contador de frecuencias de palabras, similar a los que se usan
en análisis de texto y motores de búsqueda.

**Operaciones**:

- `NuevoContador() *ContadorDePalabras` — crea un contador vacío
- `AgregarTexto(texto string)` — procesa un texto y actualiza las frecuencias
  de cada palabra (separando por espacios, sin distinguir mayúsculas/minúsculas)
- `Frecuencia(palabra string) int` — devuelve la frecuencia de una palabra
  específica (0 si no existe)
- `TotalPalabras() int` — devuelve la cantidad total de palabras procesadas
- `PalabrasUnicas() int` — devuelve la cantidad de palabras distintas encontradas
- `TopN(n int) []ParFrecuencia` — devuelve las `n` palabras más frecuentes
  ordenadas de mayor a menor frecuencia (si hay menos de `n` distintas,
  devuelve todas)
- `ListarPalabras() []string` — devuelve todas las palabras registradas

**Invariante**:

- `TotalPalabras()` debe ser igual a la suma de todas las frecuencias.
- `TopN` debe devolver como máximo `n` elementos.

**Preguntas**:

- ¿Qué complejidad tiene `TopN`? ¿Cómo la implementaste?
- ¿Qué ventaja tiene usar una tabla de _hash_ frente a un slice para
  almacenar las frecuencias?
- Si el texto fuera muy largo (millones de palabras), ¿cómo optimizarías
  el procesamiento?

→ `02-conteo-palabras/`

---

## 3. Agenda Telefónica

Implementar una agenda de contactos que permita gestionar nombres, teléfonos
y correos electrónicos.

**Operaciones**:

- `NuevaAgenda() *Agenda` — crea una agenda vacía
- `AgregarContacto(contacto Contacto)` — agrega un nuevo contacto
  (si el nombre ya existe, actualiza sus datos)
- `BuscarContacto(nombre string) (Contacto, error)` — devuelve un contacto
  por su nombre (error si no existe)
- `EliminarContacto(nombre string) error` — elimina un contacto por su nombre
  (error si no existe)
- `ExisteContacto(nombre string) bool` — verifica si un nombre está registrado
- `ListarContactos() []Contacto` — devuelve todos los contactos
- `Cantidad() int` — devuelve la cantidad de contactos

**Estructura `Contacto`**:

```go
type Contacto struct {
    Nombre   string
    Telefono string
    Email    string
}
```

**Invariante**:

- Dos contactos no pueden tener el mismo nombre (la tabla de _hash_ lo
  garantiza al usar el nombre como clave).
- Después de eliminar un contacto, no debe aparecer en los listados.

**Preguntas**:

- ¿Por qué usamos el nombre como clave y no el teléfono o el email?
- ¿Qué pasaría si dos personas distintas tuvieran el mismo nombre?
  ¿Cómo rediseñarías la estructura para evitarlo?
- ¿Qué complejidad tiene `BuscarContacto`? ¿Y `ListarContactos`?

→ `03-agenda/`
