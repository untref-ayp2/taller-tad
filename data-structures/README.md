# Data Structures — AyP2

Contratos (interfaces) y esqueletos para las estructuras de datos de
Algoritmos y Programación II.

## Estructura

```
stack/       # interface Stack[T] + SliceStack[T] + StackList[T] + tests
queue/       # interface Queue[T] + SliceQueue[T] + QueueList[T] + tests
list/        # interface List[T] + SinglyLinkedList + DoublyLinkedList +
             #   SentinelLinkedList + CircularLinkedList + tests
bitmap/      # BitMap (8/32/64 bits) + tests
hashtable/   # interface HashTable[K, V] + HashTableOpenAddressing +
             #   HashTableChaining + tests (usar hash/maphash para
             #   la función de hash con K comparable)
set/         # interface Set[T] + MapSet[T] + HashTableSet[T] + OrderedSet[T] + tests
dictionary/  # interface Dictionary[K, V] + HashMapDictionary[K, V] + tests
heap/        # interface Heap[T] + SliceHeap[T] (min/max) + tests
priorityqueue/ # PriorityQueue[T] (composición con Heap[T]) + tests
tree/        # TreeNode[T] + BinaryTree[T] + tests
binarysearchtree/  # BinarySearchTree[T] (ABB) por composición + tests
avltree/     # AVLTree[T] + AVLNode[T] + tests
```

Cada paquete define una **interfaz** que los alumnos deben implementar en su
fork. La convención de nombres varía según el paquete:

- `stack/`, `queue/`: `*_slice.go` para la impl. sobre slices, `*_list.go` para la impl. sobre listas enlazadas
- `list/`: `*_linked.go` según el tipo de enlace
- `set/`: `*_map.go`, `*_hashtable.go`, `*_ordered.go` según la estructura interna
- `hashtable/`: nombres completos descriptivos
- `heap/`: `*_slice.go` para la impl. sobre slice

## Requisitos

- Go 1.22 o superior
- Opcional: [golangci-lint](https://golangci-lint.run/) para linting local

## Cómo usar

```bash
git clone https://github.com/untref-ayp2/data-structures.git
cd data-structures
```

Para verificar que las interfaces compilan:

```bash
go build ./...
```

Para ejecutar los tests:

```bash
make test
```

O con más detalle:

```bash
make test-v
```

Para ejecutar los tests de un paquete específico:

```bash
make test-pkg PKG=stack/
```

Para ejecutar el linter (si tenés golangci-lint instalado):

```bash
make lint
```

Para más información sobre cómo trabajar, ver [CONTRIBUTING.md](CONTRIBUTING.md).

## Uso con taller-tad

Si también forkaste [taller-tad](https://github.com/untref-ayp2/taller-tad),
los ejercicios de ese repo importan este paquete. Para que funcionen
localmente, agregá esta directiva `replace` en el `go.mod` de taller-tad:

```
replace github.com/tu-usuario/data-structures => ../data-structures
```

(o ajustá la ruta según dónde hayas clonado cada repo).

## Licencia

MIT
