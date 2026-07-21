# Cómo trabajar en este repositorio

## Ejecutar los tests

Ejecutá todos los tests:

    make test

Ejecutar los tests con detalle de cada caso:

    make test-v

Ejecutar los tests de un paquete específico:

    make test-pkg PKG=ruta/alpaquete

Ejemplo — solo tests de stack:

    make test-pkg PKG=stack/

## Feedback automático (CI)

Cada vez que hacés `git push`, GitHub Actions ejecuta los tests automáticamente.
Vas a ver el resultado en la pestaña **Actions** de tu repositorio.

Si los tests pasan: se muestra un tick verde.
Si algún test falla: se muestra una cruz roja — revisá el error, corregí el código y volvé a pushear.

## Tu repositorio tiene un PR "Feedback"

Cuando empezaste la assignment, GitHub Classroom creó automáticamente un Pull Request
llamado **"Feedback"** en tu repositorio. Este PR es tu canal de comunicación con los docentes.

**No cierres este PR. Si lo cerrás, perdés el canal de ayuda.**

### Cómo funcionan los commits en el PR "Feedback"

Cada vez que hacés `git push`, tus commits aparecen automáticamente en el PR "Feedback".
No necesitás hacer nada especial — el PR se actualiza solo.

### Cómo pedir ayuda

1. **Commitear y pushear** tu intento (aunque no esté completo):

       git add .
       git commit -m "intento ejercicio stack"
       git push

2. **Ir al PR "Feedback"** — en la pestaña Pull Requests de tu repositorio,
   click en el PR llamado "Feedback".

3. **Agregar un comentario** en el PR explicando:
   - Qué estructura estás implementando.
   - Qué intentaste.
   - Qué error estás viendo.

4. **Mencionar al docente** — Para mencionar a un docente podes arrobarlo
   con el nombre de usuario o el correo electrónico, en el comentario para
   que reciba una notificación: 

       @nombredeusuario ayuda con la implementación de Stack

### Normas para pedir ayuda

- **Un solo PR por repositorio** — no necesitás abrir más PRs; el PR "Feedback"
  ya existe y es el canal correcto.
- **Antes de pedir ayuda** — intentá resolver el ejercicio por tu cuenta al
  menos 3 veces. En el comentario explicá qué intentaste.
- **No pedir que te lo resuelvan** — el objetivo es que aprendas, no que te
  den la solución.
- **No cerrar el PR** — el PR "Feedback" debe mantenerse abierto para que los
  docentes puedan verte y responderte.

## Convenciones de código

- **Errores**: usá errores predefinidos del paquete
  (ej. `var ErrStackEmpty = errors.New("la pila esta vacia")`).
- **Genéricos**: usá `[T any]` para tipos que no requieren comparabilidad;
  `[T comparable]` cuando sí la necesitás.
- **Nombres**: las interfaces van en singular (`Stack`, `Queue`, `List`).
  Las implementaciones tienen un prefijo descriptivo (`SliceStack`, `ListStack`).
- **Indentación**: tabs (ya configurado en `.editorconfig`).

## Estructura de paquetes

Cada paquete define una interfaz que los alumnos deben implementar. La convención de nombres:

- `stack/`, `queue/`: `*_slice.go` para la impl. sobre slices, `*_list.go` para la impl. sobre listas enlazadas
- `list/`: `*_linked.go` según el tipo de enlace
- `set/`: `*_map.go`, `*_hashtable.go`, `*_ordered.go` según la estructura interna
- `heap/`: `*_slice.go` para la impl. sobre slice

Paquetes disponibles:

```
stack/       # interface Stack[T] + SliceStack + ListStack + tests
queue/       # interface Queue[T] + SliceQueue + ListQueue + tests
list/        # interface List[T] + SinglyLinked + DoublyLinked + tests
bitmap/      # BitMap (8/32/64 bits) + tests
hashtable/   # interface HashTable[K, V] + implementations + tests
set/         # interface Set[T] + implementations + tests
dictionary/  # interface Dictionary[K, V] + implementation + tests
heap/        # interface Heap[T] + SliceHeap (min/max) + tests
priorityqueue/ # PriorityQueue[T] (composición con Heap[T]) + tests
tree/        # TreeNode[T] + BinaryTree[T] + tests
binarysearchtree/ # BinarySearchTree[T] (ABB) + tests
avltree/     # AVLTree[T] + AVLNode[T] + tests
```

## Comandos útiles

    make fmt    # formatear todo el código
    make lint   # verificar estilo con linter
    make build  # compilar todo sin ejecutar
    make clean  # limpiar archivos generados

## Requisitos

- Go 1.22 o superior.
- Opcional: golangci-lint (https://golangci-lint.run/) para verificar estilo
  localmente.
