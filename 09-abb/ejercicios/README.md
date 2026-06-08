# Ejercicios: Árbol Binario de Búsqueda

Antes de comenzar, implementá `BinarySearchTree[T]` en tu repositorio data-structures.

---

## 1. Conjunto ordenado con ABB

Implementá el tipo `ConjuntoOrdenado[T any]` del paquete `conjunto_ordenado`.
Debe usar un `BinarySearchTree[T]` como estructura de almacenamiento subyacente.

Un ABB mantiene naturalmente los elementos ordenados: el recorrido inorder devuelve
los valores en orden ascendente. Esto hace que operaciones como `Valores()` sean
O(n) y devuelvan el resultado directamente ordenado, a diferencia de las
implementaciones sobre listas o tablas de hash.

**Valores**: elementos de tipo `T` con una función de comparación.

**Operaciones**:

- `Agregar(elemento T)` — inserta un elemento. Si ya existe, no tiene efecto.
- `Contiene(elemento T) bool` — verifica si el elemento existe.
- `Eliminar(elemento T)` — remueve un elemento. Si no existe, no tiene efecto.
- `Cantidad() int` — devuelve la cantidad de elementos.
- `Valores() []T` — devuelve todos los elementos en orden ascendente.

**Preguntas**:

- ¿Qué complejidad tienen Agregar, Contiene y Eliminar?
- ¿Cómo se compara con la implementación de `OrderedSet` sobre lista enlazada?
- ¿Por qué `Valores()` no necesita ordenar explícitamente?

→ `03-conjunto-ordenado/`

---

## 2. Guía telefónica ordenada

Implementá el tipo `GuiaTelefonica` del paquete `guia_telefonica`.
Debe usar un `BinarySearchTree[Contacto]` internamente, comparando los contactos
por su nombre. La función de comparación ya está provista en el constructor.

Compará esta implementación con la versión sobre tabla de hash del ejercicio
`03-agenda` del capítulo `05-hashing`.

**Valores**: contactos con nombre, teléfono y email.

**Operaciones**:

- `AgregarContacto(contacto Contacto)` — agrega o actualiza un contacto.
- `BuscarContacto(nombre string) (Contacto, error)` — busca por nombre exacto.
- `EliminarContacto(nombre string) error` — elimina un contacto por nombre.
- `ExisteContacto(nombre string) bool` — verifica si un nombre está registrado.
- `ListarContactos() []Contacto` — devuelve todos los contactos en orden alfabético.
- `Cantidad() int` — devuelve la cantidad de contactos.

**Preguntas**:

- ¿Qué complejidad tiene `ListarContactos` en cada versión (ABB vs hash)?
- ¿En qué casos conviene usar una tabla de hash? ¿En cuáles un ABB?
- ¿Cómo cambia la complejidad de `BuscarContacto` si la guía tiene 10 o 10.000 contactos?

→ `04-guia-telefonica/`
