# Agenda Telefónica

Implementar una agenda de contactos usando `HashTable[string, Contacto]` de
`data-structures/hashtable`.

Completar los métodos de `Agenda`:

```go
func NuevaAgenda() *Agenda
func (a *Agenda) AgregarContacto(contacto Contacto)
func (a *Agenda) BuscarContacto(nombre string) (Contacto, error)
func (a *Agenda) EliminarContacto(nombre string) error
func (a *Agenda) ExisteContacto(nombre string) bool
func (a *Agenda) ListarContactos() []Contacto
func (a *Agenda) Cantidad() int
```

- `BuscarContacto` devuelve error si el contacto no existe.
- `EliminarContacto` devuelve error si el contacto no existe.
- `AgregarContacto` actualiza los datos si el nombre ya existe.
