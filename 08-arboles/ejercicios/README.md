# Ejercicios: Árboles Binarios

Antes de comenzar, implementá `TreeNode[T]` en tu fork de data-structures
(paquete `tree/`).

Luego completá las siguientes funciones.

---

## 1. Recorridos del árbol binario

Implementar las tres funciones de recorrido en profundidad.
Cada una debe devolver un slice con los valores de los nodos en el orden
correspondiente.

```go
func Preorden[T any](root *tree.TreeNode[T]) []T
func Inorden[T any](root *tree.TreeNode[T]) []T
func Postorden[T any](root *tree.TreeNode[T]) []T
```

- Si `root` es `nil`, devolver un slice vacío.
- No usar funciones auxiliares externas (la recursión es parte del ejercicio).
- Pista: crear un slice, visitar los nodos en orden, devolver el slice.

Ejemplo con el árbol del apunte:

```
      +
     / \
    a   *
       / \
      -   d
     / \
    b   c
```

| Recorrido | Resultado |
|---|---|
| Preorden  | `[+, a, *, −, b, c, d]` |
| Inorden   | `[a, +, b, −, c, *, d]` |
| Postorden | `[a, b, c, −, d, *, +]` |

---

## 2. Evaluar expresión aritmética

Dado un árbol de expresión aritmética (nodos internos = operadores, hojas = números),
evaluarlo usando un recorrido **postorden** con una **pila**.

```
        *
       / \
      +   2
     / \
    3   4
```

Evaluar con postorden: `[3, 4, +, 2, *]`

### Notación Polaca Inversa (RPN)

El recorrido postorden de un árbol de expresión produce la expresión en
[notación polaca inversa](https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa) (RPN).

Para evaluar RPN con una pila:

1. Recorrer los valores de izquierda a derecha.
2. Si el valor es un número, apilarlo.
3. Si el valor es un operador, desapilar los dos operandos (primero el tope es el
   operando derecho), aplicar la operación y apilar el resultado.
4. Al final, el tope de la pila contiene el resultado.

En este ejercicio **no** implementes RPN desde cero. Usá el recorrido `Postorden`
que implementaste en el ejercicio 1 para obtener la secuencia, y luego evaluala
con una pila (usá `stack.Stack[int]` de data-structures).

```go
func Evaluar(arbol *tree.TreeNode[string]) (int, error)
```

- Los operandos son números enteros (pueden tener varios dígitos).
- Operadores soportados: `+`, `-`, `*`, `/`.
- Si hay división por cero, devolver error.
- El paquete `parser` (en `08-arboles/parser/`) convierte una cadena como
  `"(3+4)*2"` en un árbol de expresión. Lo usamos en los tests para armar
  los árboles, pero la función `Evaluar` recibe directamente el árbol ya armado.

---

→ `01-ejercicios/`
