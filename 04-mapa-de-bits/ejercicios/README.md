# Ejercicios: Mapa de Bits

Antes de comenzar, implementá `BitMap` en tu repositorio
[`data-structures`](https://github.com/untref-ayp2/data-structures)
(el que se creó al aceptar la tarea en GitHub Classroom),
paquete `bitmap/`. El TAD debe soportar mapas de bits de 8, 32 y 64 bits.

Luego completá los siguientes ejercicios usando `BitMap`:

---

## 1. Registro de Lluvias

Implementar un registro de lluvias anuales, donde por cada mes se registran los
días en que hubo precipitaciones.

**Valores**: 12 meses, cada uno con un mapa de bits de 32 posiciones
(días 1 a 31).

**Operaciones**:

- `NewLluvias() *Lluvias` — crea un nuevo registro vacío
- `Registrar(d uint8, m Mes)` — registra que llovió el día `d` del mes `m`
  (si la fecha es inválida, no hace nada)
- `Desregistrar(d uint8, m Mes)` — desregistra un día lluvioso
- `Llovio(d uint8, m Mes) bool` — devuelve si llovió el día `d` del mes `m`
- `ListarMes(m Mes) []uint8` — devuelve los días con lluvia del mes `m`
- `ListarDiasEnAmbosMeses(m1, m2 Mes) []uint8` — devuelve los días que
  llovieron tanto en `m1` como en `m2`

**Invariante**:

- Cada mes debe verificar que el día esté en el rango válido
  (meses con 28/29/30/31 días según corresponda).
- Un mismo día registrado múltiples veces no debe duplicarse en los listados.

**Preguntas**:

- ¿Por qué usar un `BitMap` en lugar de un `[12][32]bool`? ¿Qué ventajas
  tiene en memoria y en operaciones como `ListarDiasEnAmbosMeses`?
- Si quisieras extender el registro a varios años, ¿cómo modificarías la
  estructura?
- `ListarDiasEnAmbosMeses` puede implementarse como una intersección de
  conjuntos. ¿Cómo harías esa intersección usando solo operaciones de bits
  (AND)?

→ `01-lluvias/`

---

## 2. Registro de Asistencia

Implementar un registro de asistencia para un curso que tiene `N` alumnos y `M`
clases.

**Valores**: `N` alumnos, cada uno con un mapa de bits de `M` posiciones
(una por clase).

**Operaciones**:

- `NewAsistencias(cantAlumnos uint, cantClases uint8) *Asistencias` —
  crea un nuevo registro
- `RegistrarAsistencia(alumno uint, clase uint8)` — registra la asistencia
  del alumno a la clase (si alumno o clase son inválidos, no hace nada)
- `Asistio(alumno uint, clase uint8) bool` — devuelve si el alumno asistió
  a la clase
- `ListarClase(clase uint8) []uint` — devuelve los alumnos que asistieron a
  la clase indicada
- `ListarAlumno(alumno uint) []uint8` — devuelve las clases a las que
  asistió el alumno indicado
- `ListarAsistencias() []uint8` — devuelve las clases a las que asistieron
  **todos** los alumnos
- `ListarAsistenciaPerfecta() []uint` — devuelve los alumnos que asistieron a
  **todas** las clases

**Invariante**:

- Todos los mapas de bits deben tener el mismo tamaño (`cantClases`).
- Las posiciones fuera de rango deben ignorarse.
- `cantClases` no debe superar los 64 bits.

**Preguntas**:

- `ListarAsistencias()` pide las clases comunes a todos los alumnos. ¿Cómo
  harías esa intersección con operaciones de bits?
- `ListarAsistenciaPerfecta()` pide los alumnos que nunca faltaron. ¿Cómo
  determinás si un alumno tiene todos los bits encendidos?
- ¿Qué orden de complejidad tiene cada operación? ¿Cómo se compara con usar
  un `[][]bool`?

→ `02-asistencia/`
