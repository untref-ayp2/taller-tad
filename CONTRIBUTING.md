# Cómo trabajar en este repositorio

## Ejecutar los tests

Ejecutá todos los tests:

    make test

Ejecutar los tests con detalle de cada caso:

    make test-v

Ejecutar los tests de un paquete específico:

    make test-pkg PKG=03-listas/ejercicios/...

Ejemplo — solo tests de listas:

    make test-pkg PKG=03-listas/ejercicios/...

## Feedback automático (CI)

Cada vez que hacés `git push`, GitHub Actions ejecuta los tests automáticamente.
Vas a ver el resultado en la pestaña **Actions** de tu repositorio.

Si los tests pasan: se muestra un tick verde.
Si algún test falla: se muestra una cruz roja — revisá el error, corregí el código y volvé a pushear.

## Tu repositorio tiene un PR "Feedback"

Cuando aceptaste la asignación con classroom50, se creó automáticamente un Pull Request
llamado **"Feedback"** en tu repositorio. Este PR es tu canal de comunicación con los docentes.

**No cierres este PR. Si lo cerrás, perdés el canal de ayuda.**

### Cómo funcionan los commits en el PR "Feedback"

Cada vez que hacés `git push`, tus commits aparecen automáticamente en el PR "Feedback".
No necesitás hacer nada especial — el PR se actualiza solo.

### Cómo pedir ayuda

1. **Commitear y pushear** tu intento (aunque no esté completo):

       git add .
       git commit -m "intento ejercicio fraccion"
       git push

2. **Ir al PR "Feedback"** — en la pestaña Pull Requests de tu repositorio,
   click en el PR llamado "Feedback".

3. **Agregar un comentario** en el PR explicando:
   - Qué ejercicio estás haciendo.
   - Qué intentaste.
   - Qué error estás viendo.

4. **Mencionar al docente** — Para mencionar a un docente podes arrobarlo
   con el nombre de usuario o el correo electrónico, en el comentario para
   que reciba una notificación: 

       @nombredeusuario ayuda con el ejercicio de fraccion

### Normas para pedir ayuda

- **Un solo PR por repositorio** — no necesitás abrir más PRs; el PR "Feedback"
  ya existe y es el canal correcto.
- **Antes de pedir ayuda** — intentá resolver el ejercicio por tu cuenta al
  menos 3 veces. En el comentario explicá qué intentaste.
- **No pedir que te lo resuelvan** — el objetivo es que aprendas, no que te
  den la solución.
- **No cerrar el PR** — el PR "Feedback" debe mantenerse abierto para que los
  docentes puedan verte y responderte.

## Dependencia con data-structures

`data-structures/` (Stack, Queue, List, etc.) viene incluido como subdirectorio de este repositorio.
No necesitás clonarlo aparte. El `replace` en el `go.mod` apunta a `./data-structures`.

## Convenciones de código

- **Errores**: usá errores predefinidos del paquete
  (ej. `var ErrStackEmpty = errors.New("la pila esta vacia")`).
- **Genéricos**: usá `[T any]` para tipos que no requieren comparabilidad;
  `[T comparable]` cuando sí la necesitás.
- **Nombres**: las interfaces van en singular (`Stack`, `Queue`, `List`).
- **Indentación**: tabs (ya configurado en `.editorconfig`).

## Estructura de capítulos

```
01-tipos-abstractos-de-datos/   # capítulo 3-1 (no depende de data-structures)
02-pilas-colas/                  # capítulo 3-2
03-listas/                       # capítulo 3-3
04-mapa-de-bits/                 # capítulo 3-4
05-hashing/                      # capítulo 3-5
06-conjuntos/                   # capítulo 3-6
07-diccionarios/                # capítulo 3-7
08-arboles/                     # capítulo 3-8
09-abb/                         # capítulo 3-9
10-arboles-balanceados/         # capítulo 3-10
10-monticulo-binario/           # capítulo 3-11
```

## Comandos útiles

    make fmt    # formatear todo el código
    make lint   # verificar estilo con linter
    make build  # compilar todo sin ejecutar
    make clean  # limpiar archivos generados

    go run ./01-tipos-abstractos-de-datos/ejemplos/contador  # ejecutar un ejemplo

## Requisitos

- Go 1.22 o superior.
- Opcional: golangci-lint (https://golangci-lint.run/) para verificar estilo
  localmente.
