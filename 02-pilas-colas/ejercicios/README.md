# Ejercicios: Pilas y Colas

Antes de comenzar, implementá `SliceStack[T]` y `SliceQueue[T]` en tu fork de
[data-structures](https://github.com/untref-ayp2/data-structures).

Luego completá las siguientes funciones usando pilas y colas:

### 01. InvertirCadena

Escribir una función que reciba una cadena de caracteres y devuelva la cadena
invertida. Analizar el orden.

### 02. EsPalindromo

Escribir una función que verifique si una cadena es palíndromo (es igual a su
inversa). Por ejemplo: `"1456541"` y `"145541"` son palíndromos.

### 03. EstaBalanceada

Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves
está bien balanceada. Por ejemplo: `[()]{}{[()()]()}` debe devolver `true`,
mientras que `[(])` debe devolver `false`.

### 04. UnirColas

Escribir una función que, dadas dos colas, construya una cola con el resultado
de poner una a continuación de la otra.

Ejemplo: si `q1 = [1, 2, 3]` (1 al frente) y `q2 = [5, 7]`, el resultado es
`[1, 2, 3, 5, 7]` (1 al frente).

### 05. EvaluarRPN

Escribir una función que reciba una cadena que representa una expresión en
[notación polaca inversa (RPN)](https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa)
y devuelva el resultado.

Ejemplo: `"2 3 + 5 *"` debe devolver `25`.
Operadores válidos: `+`, `-`, `*`, `/`.

### 06. EsPosibleConPila

Dadas dos secuencias de números enteros (entrada y salida),
determinar si es posible transformar la secuencia de entrada en la secuencia
de salida usando una pila. Es el problema clásico de ordenamiento con pila
(*stack sorting*).

Ejemplo: entrada `[1, 2, 3]`, salida `[3, 2, 1]` → `true`
         entrada `[1, 2, 3]`, salida `[3, 1, 2]` → `false`
