// Package afinidad ordena usuarios por intereses compartidos.
package afinidad

import "github.com/untref-ayp2/data-structures/set"

// RankingAfinidad recibe un mapa de usuario → intereses y un usuario
// objetivo, y devuelve el resto de los usuarios ordenados por cantidad
// de intereses compartidos con el objetivo (descendente).
//
// En caso de empate, ordena alfabéticamente por nombre de usuario.
func RankingAfinidad(intereses map[string]set.Set[string], objetivo string) []string {
	// Completar
	return nil
}
