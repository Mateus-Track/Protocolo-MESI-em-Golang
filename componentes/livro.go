package componentes

import (
	"MESI/config"
	"errors"
	"fmt"
	"time"
)

type Reserva struct {
	data_inicio time.Time
	data_fim    time.Time
	id_pessoa   int
}

func InicializaReserva(data_inicio, data_fim time.Time, id_pessoa int) (Reserva, error) {
	if data_inicio.After(data_fim) {
		return Reserva{}, errors.New("InvalidDate")
	}

	return Reserva{
		data_inicio,
		data_fim,
		id_pessoa,
	}, nil
}

func (reserva *Reserva) ToString() string {
	var str = ""

	str += reserva.data_inicio.Format(config.TIME_LAYOUT) +
		" até " +
		reserva.data_fim.Format(config.TIME_LAYOUT)

	return str
}

type Livro struct {
	Reservas []Reserva
	Nome     string
	Secao    string
}

func InicializaLivro() Livro {
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
