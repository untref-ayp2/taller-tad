# Taller de TAD

Repositorio complementario de la sección **Tipos Abstractos de Datos (TAD)** de los apuntes de Algoritmos y Programación II.

## Estructura

```
01-tipos-abstractos-de-datos/   # ← capítulo 3-1
├── ejemplos/
│   └── contador/               # ejemplo resuelto del apunte
└── ejercicios/
    ├── 01-fraccion/            # esqueleto con tests
    └── 02-reloj/               # esqueleto con tests

02-pilas-colas/                 # ← capítulo 3-2
└── ejercicios/
    └── 01-ejercicios/          # ejercicios que usan pilas y colas

03-listas/                      # ← capítulo 3-3
└── ejercicios/
    └── 01-ejercicios/          # implementaciones y ejercicios con listas

04-mapa-de-bits/                # ← capítulo 3-4
└── ejercicios/
    ├── 01-lluvias/             # registro de lluvias con bitmap
    └── 02-asistencia/          # registro de asistencias con bitmap

05-hashing/                     # ← capítulo 3-5
└── ejercicios/
    ├── 01-diccionario/         # traductor español-inglés
    ├── 02-conteo-palabras/     # contador de frecuencias
    └── 03-agenda/              # agenda de contactos

06-conjuntos/                   # ← capítulo 3-6
└── ejercicios/
    ├── 01-afinidad/            # ranking por intereses compartidos
    ├── 02-papers/              # buscador de papers por keywords
    └── 03-directorios/         # fusión de directorios ordenados

07-diccionarios/                # ← capítulo 3-7
└── ejercicios/
    ├── 01-stock/               # registro de stock de productos
    └── 02-directorio/          # directorio con búsqueda inversa

08-arboles/                     # ← capítulo 3-8
├── parser/                     # parser de expresiones → árbol binario
└── ejercicios/
    ├── 01-ejercicios/          # recorridos + evaluación de expresiones
    └── 01-altura/              # altura de árboles balanceado y degenerado

09-abb/                         # ← capítulo 3-9 (ABB)
└── ejercicios/
    ├── 01-conjunto-ordenado/   # conjunto ordenado sobre ABB
    └── 02-guia-telefonica/     # guía telefónica ordenada sobre ABB

10-arboles-balanceados/         # ← capítulo 3-10 (AVL)
└── ejercicios/
    └── README.md               # ejercicios de lápiz y papel

10-monticulo-binario/           # ← capítulo 3-11
└── ejercicios/
    ├── 01-merge-listas/        # merge de K listas ordenadas con PQ
    └── 02-triage/              # sistema de triage hospitalario con PQ
```

Cada directorio `ejercicios/` contiene esqueletos incompletos y tests para que implementes la solución.

## Dependencias

Los ejercicios de `02-pilas-colas/`, `03-listas/`, `04-mapa-de-bits/`, `05-hashing/`,
`06-conjuntos/`, `07-diccionarios/`, `08-arboles/`, `09-abb/` y `10-monticulo-binario/`
importan las interfaces de [data-structures](https://github.com/untref-ayp2/data-structures).
Asegurate de tener el repositorio clonado en `../data-structures`
(ver `go.mod` para el `replace`).

`01-tipos-abstractos-de-datos` no depende de data-structures. Los capítulos que se agreguen
en el futuro y dependan de él deben incluirse en esta lista.

## Cómo usar

```bash
# Ejecutar todos los tests
make test

# Con detalle de cada caso
make test-v

# Tests de un ejercicio específico
go test -v ./01-tipos-abstractos-de-datos/ejercicios/01-fraccion/...
```

Para ejecutar un ejemplo:

```bash
go run ./01-tipos-abstractos-de-datos/ejemplos/contador
```

**Nota:** este repositorio depende de `data-structures`. Asegurate de tenerlo
clonado en `../data-structures` (lo crea classroom50 automáticamente al aceptar
ambas asignaciones).

Para más información, ver [CONTRIBUTING.md](CONTRIBUTING.md).

## Requisitos

- Go 1.22 o superior
- Opcional: [golangci-lint](https://golangci-lint.run/) para linting local
