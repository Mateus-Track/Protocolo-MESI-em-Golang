package models

import (
	"MESI/config"
	"errors"
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
