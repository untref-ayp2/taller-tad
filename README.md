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
```

Cada directorio `ejercicios/` contiene un esqueleto incompleto y tests para que implementes la solución.

## Dependencias

Los ejercicios de `02-pilas-colas/` importan las interfaces de
[data-structures](https://github.com/untref-ayp2/data-structures).
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
go test ./01-tipos-abstractos-de-datos/ejercicios/fraccion/...
```

Para correr todos los tests:

```bash
go test ./...
```

## Requisitos

Go 1.22 o superior.
