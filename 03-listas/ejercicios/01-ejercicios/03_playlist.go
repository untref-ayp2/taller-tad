package ejercicios

// Playlist implementa un reproductor con repetición circular.
type Playlist struct {
	// Completar
}

// NewPlaylist crea una playlist vacía.
func NewPlaylist() *Playlist {
	// Completar
	return nil
}

// Add agrega una canción al final de la playlist.
func (p *Playlist) Add(cancion string) {
	// Completar
}

// Current devuelve la canción actual. Si la playlist está vacía devuelve "".
func (p *Playlist) Current() string {
	// Completar
	return ""
}

// Next avanza a la siguiente canción. Si la playlist está vacía no hace nada.
func (p *Playlist) Next() {
	// Completar
}

// Prev retrocede a la canción anterior. Si la playlist está vacía no hace nada.
func (p *Playlist) Prev() {
	// Completar
}
