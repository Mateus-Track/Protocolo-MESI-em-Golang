package models

import (
	"errors"
	"fmt"
)

type Livro struct {
	Reservas []Reserva
	Nome     string
	Secao    string
}

func InicializarLivro() Livro {
	return Livro{
		Reservas: []Reserva{},
		Nome:     "",
		Secao:    "",
	}
}

func (livro *Livro) AdicionarReserva(reserva Reserva) error {
	for _, r := range livro.Reservas {
		if !(r.data_fim.After(reserva.data_inicio) || r.data_inicio.Before(reserva.data_fim)) {
			return errors.New("AlreadyExists")
		}
	}

	livro.Reservas = append(livro.Reservas, reserva)

	return nil
}

func (livro *Livro) ToString() string {
	var str = ""

	for _, r := range livro.Reservas {
		str += "\t" + r.ToString() + "\n"
	}

	return fmt.Sprintf("Livro: %s\n"+
		"Secao: %s\n"+
		"Reservas: \n%s\n",
		livro.Nome,
		livro.Secao,
		str,
	)
}
