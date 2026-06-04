# Ejercicios: Listas Enlazadas

Antes de comenzar, implementá las cuatro variantes de lista en tu repositorio
data-structures:
`SinglyLinkedList`, `DoublyLinkedList`, `CircularLinkedList` y `SentinelLinkedList`.

Completá las siguientes funciones y tipos en este repositorio (`taller-tad`),
usando las listas que implementaste en `data-structures`.

---

## 1. StackList

Implementar una pila (`Stack[T]`) usando internamente `SinglyLinkedList[T]`. La
pila debe respetar la interfaz `Stack[T]` definida en
`github.com/untref-ayp2/data-structures/stack`.

El tipo debe llamarse `StackList[T]` y su constructor `NewStackList[T]()`.

---

## 2. QueueList

Implementar una cola (`Queue[T]`) usando internamente `SinglyLinkedList[T]`. La
cola debe respetar la interfaz `Queue[T]` definida en
`github.com/untref-ayp2/data-structures/queue`.

El tipo debe llamarse `QueueList[T]` y su constructor `NewQueueList[T]()`.

---

## 3. InvertirLista

Escribir una función que reciba una `List[T]` y devuelva un slice con los
elementos en orden inverso.

```go
func InvertirLista[T comparable](l list.List[T]) []T
```

Ejemplo: si la lista contiene `[1, 2, 3]`, debe devolver `[3, 2, 1]`.

---

## 4. MergeListas

Escribir una función que dadas dos listas **ordenadas** devuelva una nueva lista
con los elementos de ambas en orden (merge).

```go
func MergeListas[T comparable](l1, l2 list.List[T]) list.List[T]
```

Ejemplo: `l1 = [1, 3, 5]`, `l2 = [2, 4, 6]` → resultado `[1, 2, 3, 4, 5, 6]`.

---

## 5. PlaylistCircular

Implementar un reproductor de música con repetición (`repeat`) usando una
`CircularLinkedList`. Debe permitir:

- Agregar canciones (strings)
- Avanzar a la siguiente canción (`Next`)
- Retroceder a la canción anterior (`Prev`)
- Obtener la canción actual (`Current`)

El tipo debe llamarse `Playlist` y su constructor `NewPlaylist()`.

---

## 6. Josephus

Resolver el [problema de
Josephus](https://es.wikipedia.org/wiki/Problema_de_Flavio_Josefo) usando una
lista circular.

La función recibe la cantidad de personas `n` y el salto `k`, y devuelve la
posición del sobreviviente (1-indexed).

```go
func Josephus(n, k int) int
```

Ejemplo: `Josephus(7, 3)` debe devolver `4`.

---

## 7. UndoRedo

Implementar un historial de acciones con deshacer/rehacer usando una
`DoublyLinkedList`. Debe permitir:

- Ejecutar una acción (`Do`)
- Deshacer la última acción (`Undo`)
- Rehacer la última acción deshecha (`Redo`)
- Consultar la acción actual (`Current`)

Cada acción es un string que describe la operación realizada.

El tipo debe llamarse `UndoRedo` y su constructor `NewUndoRedo()`.

→ `01-ejercicios/`
