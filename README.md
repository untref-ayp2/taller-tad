# Taller de TAD

Repositorio complementario de la sección **Tipos Abstractos de Datos (TAD)** de los apuntes de Algoritmos y Programación II.

## Estructura

```
01-tipos-abstractos-de-datos/   # ← capítulo 3-1
├── ejemplos/
│   └── contador/               # ejemplo resuelto del apunte
└── ejercicios/
    ├── fraccion/               # esqueleto con tests
    └── reloj/                  # esqueleto con tests

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
```

Cada directorio `ejercicios/` contiene esqueletos incompletos y tests para que implementes la solución.

## Dependencias

Los ejercicios de `02-pilas-colas/`, `03-listas/`, `04-mapa-de-bits/` y `05-hashing/` importan
las interfaces de [data-structures](https://github.com/untref-ayp2/data-structures).
Asegurate de tener el repositorio clonado en `../data-structures`
(ver `go.mod` para el `replace`).

## Cómo usar

```bash
git clone https://github.com/untref-ayp2/taller-tad.git
cd taller-tad
```

Para ejecutar un ejemplo:

```bash
go run ./01-tipos-abstractos-de-datos/ejemplos/contador
```

Para correr los tests de un ejercicio:

```bash
go test ./01-tipos-abstractos-de-datos/ejercicios/01-fraccion/...
```

Para correr todos los tests:

```bash
go test ./...
```

## Requisitos

Go 1.22 o superior.
