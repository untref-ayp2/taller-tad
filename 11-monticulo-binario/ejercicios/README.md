# Ejercicios: Montículo Binario y Cola de Prioridad

Antes de comenzar, implementá `heap` y `priorityqueue` en `data-structures/`.

---

## 1. Cola de Prioridad de Personas

Implementar una cola de prioridad de personas donde la prioridad se define por
la edad (mayor edad = mayor prioridad). A igual edad, se respeta el orden de
llegada. Usar un `MaxHeap` de `data-structures/heap`.

**Operaciones**:

- `NuevaColaPrioridad() *ColaPrioridad` — crea una cola vacía.
- `Agregar(p Persona)` — agrega una persona.
- `Atender() (Persona, error)` — elimina y devuelve la de mayor prioridad.
- `Siguiente() (Persona, error)` — devuelve la de mayor prioridad sin eliminarla.
- `Cantidad() int` — cantidad de personas en la cola.
- `EstaVacia() bool` — true si no hay personas.

→ `01-cola-prioridad/`

---

## 2. Merge de K listas ordenadas

Dadas K listas enlazadas ordenadas, fusionarlas en una única lista ordenada
utilizando una cola de prioridad.

**Estrategia**:

1. Insertar el primer elemento de cada lista en una cola de prioridad (mínimo).
2. Mientras la cola no esté vacía:
   - Extraer el menor elemento y agregarlo a la lista resultado.
   - Si la lista de origen tiene más elementos, insertar el siguiente en la cola.

**Valores**: K listas enlazadas de tipo `list.List[T]`.

**Operaciones**:

- `MergeKListas[T any](listas []list.List[T], cmp func(T, T) int) list.List[T]` — fusiona las listas en una sola ordenada.

**Preguntas**:

- ¿Qué complejidad temporal tiene el algoritmo si hay N elementos totales y K listas?
- ¿Qué pasa si una o más listas están vacías?

→ `02-merge-listas/`

---

## 3. Sistema de triage hospitalario

Un hospital recibe pacientes con distintos niveles de gravedad (1 = más grave,
5 = menos grave). Se debe implementar una función que determine el orden de
atención utilizando una cola de prioridad.

**Reglas**:

1. Se atiende primero al paciente de mayor gravedad (menor número).
2. En caso de igual gravedad, se atiende al que llegó primero (menor orden de llegada).

**Valores**:

- `Paciente` con los campos `Nombre string`, `Gravedad int` y `Llegada int`.

**Operaciones**:

- `Atender(pacientes []Paciente) []Paciente` — retorna los pacientes en el orden en que deben ser atendidos.

**Preguntas**:

- ¿Usaste un min-heap o un max-heap? ¿Por qué?
- ¿Qué complejidad temporal tiene atender a N pacientes?

→ `03-triage/`

---

**Nota para el alumno**: las respuestas a las preguntas teóricas deben
incluirse como comentarios al final del archivo `.go` de implementación, en un
bloque encabezado con `// === PREGUNTAS TEÓRICAS ===`.
