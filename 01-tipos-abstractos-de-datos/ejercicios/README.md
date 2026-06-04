# Ejercicios: Tipos Abstractos de Datos

## 1. TAD Fracción

Implementar un TAD `Fraccion` que represente un número racional.

**Valores**: numerador y denominador (enteros).

**Operaciones**:

- `NuevaFraccion(num, den int) (*Fraccion, error)` — constructor
- `Sumar(f *Fraccion) *Fraccion` — devuelve una nueva fracción con la suma
- `Restar(f *Fraccion) *Fraccion` — devuelve una nueva fracción con la resta
- `Multiplicar(f *Fraccion) *Fraccion` — devuelve una nueva fracción con el producto
- `Dividir(f *Fraccion) (*Fraccion, error)` — devuelve una nueva fracción con el cociente
- `Valor() float64` — devuelve el valor en punto flotante
- `String() string` — devuelve la representación como cadena `"num/den"`

**Invariante**:

- El denominador nunca debe ser 0.
- La fracción debe mantenerse siempre simplificada (usar el algoritmo de Euclides para
  calcular el máximo común divisor).

**Preguntas**:

- ¿En qué momento de las operaciones podría no cumplirse el invariante?
- ¿Cómo garantiza la atomicidad en cada primitiva?

→ `01-fraccion/`

---

## 2. TAD Reloj

Implementar un TAD `Reloj` que represente la hora del día en formato de 24 horas.

**Valores**: hora, minuto y segundo (enteros).

**Operaciones**:

- `NuevoReloj(h, m, s int) (*Reloj, error)` — constructor con validación
- `AvanzarUnSegundo()` — incrementa el reloj en 1 segundo, manejando correctamente el
  cambio de minuto, hora y el reinicio a `00:00:00` al llegar a `23:59:59`
- `Hora() int` — devuelve la hora
- `Minuto() int` — devuelve el minuto
- `Segundo() int` — devuelve el segundo
- `String() string` — devuelve la hora formateada como `"HH:MM:SS"`

**Invariante**:

- `0 <= hora <= 23`
- `0 <= minuto <= 59`
- `0 <= segundo <= 59`

**Preguntas**:

- En `AvanzarUnSegundo()`, cuando se pasa de `23:59:59` a `00:00:00`, ¿hay algún
  momento en que el invariante no se cumpla?
- ¿Cómo harías para que la primitiva sea atómica?

→ `02-reloj/`
