# Ejercicios: Montículo Binario y Cola de Prioridad

Antes de comenzar, implementá `heap` y `priorityqueue` en tu repositorio data-structures.

---

## 1. Merge de K listas ordenadas

Dadas K listas enlazadas ordenadas, fusionarlas en una única lista ordenada utilizando una cola de prioridad.

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

→ `01-merge-listas/`

---

## 2. Sistema de triage hospitalario

Un hospital recibe pacientes con distintos niveles de gravedad (1 = más grave, 5 = menos grave). Se debe implementar una función que determine el orden de atención utilizando una cola de prioridad.

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

→ `02-triage/`
