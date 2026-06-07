# Parser de expresiones aritméticas → árbol binario

Este paquete convierte una cadena como `"(3+4)*(5-2)"` en un árbol binario
donde los nodos internos son operadores y las hojas son números enteros.

```
        *
       / \
      +   2
     / \
    3   4
```

Se usa en los ejercicios de `08-arboles/` para que los estudiantes puedan
armar árboles de expresión sin tener que escribir un parser ellos mismos.

## Uso

```go
import "github.com/untref-ayp2/taller-tad/08-arboles/parser"

arbol, err := parser.Parse("(3+4)*(5-2)")
if err != nil {
    // manejar error
}
// arbol es un *tree.TreeNode[string]
```

## Cómo funciona internamente

### Tokenización

El parser primero convierte la cadena en una lista de tokens con un barrido
lineal (`tokenize`). Reconoce cuatro tipos de token:

| Token     | Ejemplos       |
|-----------|----------------|
| Número    | `42`, `0`, `7` |
| Operador  | `+`, `-`, `*`, `/` |
| Abrir par | `(`             |
| Cerrar par| `)`             |

Los espacios en blanco se ignoran. Cualquier otro carácter produce un error
(`ErrCaracterInesperado`).

### Gramática (descendente recursiva)

El parser implementa un analizador sintáctico **descendente recursivo** (*recursive descent*)
que sigue esta gramática:

```
expr   → term (('+' | '-') term)*
term   → factor (('*' | '/') factor)*
factor → NUMBER | '(' expr ')'
```

La separación en `expr` y `term` es la clave para respetar la **prioridad
de operadores**. Funciona así:

- `parseExpr` solo reconoce `+` y `-`. Cada vez que encuentra uno, toma lo
  que ya parseó como operando izquierdo y pide un `term` como operando derecho.
- `parseTerm` solo reconoce `*` y `/`. Misma lógica, pero pide `factor` como
  operandos.
- `parseFactor` reconoce números sueltos o expresiones entre paréntesis.

Como `parseExpr` delega en `parseTerm` para obtener sus operandos, y
`parseTerm` agarra todos los `*` y `/` que pueda antes de devolver el control,
la multiplicación/división siempre se agrupa antes que la suma/resta.

Cada función no terminal (`parseExpr`, `parseTerm`, `parseFactor`) parsea su
regla y construye el árbol de forma bottom-up. Cuando encuentra un operador,
crea un nodo con ese operador, asigna el subárbol izquierdo (lo que ya parseó)
y el subárbol derecho (el resultado de parsear el siguiente factor/término).

### Construcción del árbol

#### Ejemplo 1: `3+4*2` (la prioridad hace que `4*2` se agrupe primero)

1. `parseExpr` llama a `parseTerm`
2. `parseTerm` llama a `parseFactor` → obtiene `3`
3. `parseTerm` ve `*` (es de su competencia), así que consume el `*`
4. Llama a `parseFactor` → obtiene `4`
5. Construye `(* 4 ...)`, llama a `parseFactor` → obtiene `2`
6. Cierra `parseTerm`, devuelve `(* 4 2)`
7. Vuelve a `parseExpr`, ve `+` (es de su competencia), consume el `+`
8. Llama a `parseTerm` → como no hay más `*` ni `/`, devuelve `2` (parseFactor)
9. Construye `(+ 3 2)` **pero** el 2 que recibió no es el 2 literal,
   sino el resultado de parseTerm, que en el paso 6 devolvió `(* 4 2)`.
   El árbol final es `(+ 3 (* 4 2))`.

Resultado: `3 + (4 * 2) = 11`

```
      +
     / \
    3   *
       / \
      4   2
```

#### Ejemplo 2: `(3+4)*2` (los paréntesis fuerzan la suma primero)

1. `parseExpr` llama a `parseTerm`
2. `parseTerm` llama a `parseFactor`
3. `parseFactor` ve `(`, consume el `(`, llama a `parseExpr`
4. Dentro del paréntesis, `parseExpr` ve `3+4`:
   - `parseTerm` devuelve `3`
   - ve `+`, consume, `parseTerm` devuelve `4`
   - construye `(+ 3 4)`
5. `parseFactor` ve `)`, consume el `)`, devuelve `(+ 3 4)`
6. `parseTerm` ve `*` (es de su competencia), consume el `*`
7. `parseTerm` llama a `parseFactor` → obtiene `2`
8. Construye `(* (+ 3 4) 2)`
9. `parseExpr` no ve más operadores, devuelve el árbol

Resultado: `(3 + 4) * 2 = 14`

```
      *
     / \
    +   2
   / \
  3   4
```

### Manejo de errores

- `ErrCaracterInesperado`: la expresión contiene un caracter no válido.
- `ErrExpresionInvalida`: la expresión está mal formada (faltan operandos,
  paréntesis sin cerrar, operadores sin los dos operandos, etc.).

## Limitaciones

- Solo opera con números enteros (no decimales ni negativos).
- No soporta operadores unarios (como `-3`).
- No valida división por cero en tiempo de parseo (eso se hace al evaluar).
