package componentes

import "time"

type Livro struct {
	Reservas [][2]time.Time
	Nome     string
	Secao    string
}

func InicializaLivro() Livro {
	return Livro{
		Reservas: make([][2]time.Time, 4),
		Nome:     "",
		Secao:    "",
	}
}
