package componentes

import "time"

type Livro struct {
	Reservas [][2]time.Time
	Nome     string
	Secao    string
}
