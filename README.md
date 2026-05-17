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
```

Cada directorio `ejercicios/` contiene un esqueleto incompleto y tests para que implementes la solución.

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
