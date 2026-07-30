# Ejercicios: Pilas y Colas

Antes de comenzar, implementá `SliceStack[T]` y `SliceQueue[T]` en `data-structures/`.

Luego completá las siguientes funciones usando pilas y colas.

---

## 1. InvertirCadena

Escribir una función que reciba una cadena de caracteres y devuelva la cadena
invertida. Analizar el orden.

---

## 2. EsPalindromo

Escribir una función que verifique si una cadena es palíndromo (es igual a su
inversa). Por ejemplo: `"1456541"` y `"145541"` son palíndromos.

---

## 3. EstaBalanceada

Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves
está bien balanceada. Por ejemplo: `[()]{}{[()()]()}` debe devolver `true`,
mientras que `[(])` debe devolver `false`.

---

## 4. UnirColas

Escribir una función que, dadas dos colas, construya una cola con el resultado
de poner una a continuación de la otra.

Ejemplo: si `q1 = [1, 2, 3]` (1 al frente) y `q2 = [5, 7]`, el resultado es
`[1, 2, 3, 5, 7]` (1 al frente).

---

## 5. EvaluarRPN

Escribir una función que reciba una cadena que representa una expresión en
[notación polaca inversa (RPN)](https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa)
y devuelva el resultado.

Ejemplo: `"2 3 + 5 *"` debe devolver `25`.
Operadores válidos: `+`, `-`, `*`, `/`.

---

## 6. EsPosibleConPila

Dadas dos secuencias de números enteros (entrada y salida),
determinar si es posible transformar la secuencia de entrada en la secuencia
de salida usando una pila. Es el problema clásico de ordenamiento con pila
(*stack sorting*).

Ejemplo: entrada `[1, 2, 3]`, salida `[3, 2, 1]` → `true`
         entrada `[1, 2, 3]`, salida `[3, 1, 2]` → `false`

→ `01-ejercicios/`

---

## 7. ColaCircular

Implementar una cola circular genérica (`ColaCircular[T]`) sobre un arreglo de
tamaño fijo, también conocida como *ring buffer*.

A diferencia de `SliceQueue[T]`, esta cola:

- Tiene una capacidad máxima fija definida al crearla.
- Reutiliza el espacio del arreglo cuando se extraen elementos (el frente y el
  final "dan la vuelta").
- La operación `Dequeue` es $O(1)$ real (no desplaza elementos como el slice).

### Métodos requeridos

`NewColaCircular[T](capacidad int) *ColaCircular[T]`
: Crea una cola circular vacía con la capacidad indicada.

`Enqueue(val T) error`
: Agrega un elemento al final. Error si la cola está llena.

`Dequeue() (T, error)`
: Extrae y devuelve el elemento del frente. Error si la cola está vacía.

`Front() (T, error)`
: Devuelve el elemento del frente sin extraerlo. Error si la cola está vacía.

`IsEmpty() bool`
: Devuelve `true` si la cola no tiene elementos.

`IsFull() bool`
: Devuelve `true` si la cola alcanzó su capacidad máxima.

`Size() int`
: Devuelve la cantidad de elementos actualmente en la cola.

### Pistas

- Usá un *slice* de tamaño fijo como almacenamiento subyacente.
- Mantené dos índices: `front` (frente) y `rear` (final). Inicializá `rear` en `-1`.
- Para avanzar un índice de forma circular: `(indice + 1) % capacidad`.
- La cola está vacía cuando `size == 0` y llena cuando `size == capacidad`.
- No es necesario "limpiar" las celdas al hacer `Dequeue`; basta con avanzar `front` y decrementar `size`.

→ `02-cola-circular/`
