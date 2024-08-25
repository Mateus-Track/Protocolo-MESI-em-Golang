package componentes

import "fmt"

type Processador struct {
	Id         int
	Cachezinha Cache //precisa ser maiúsculo o nome da variável..
}

func InicializaProcessador(id int) Processador {
	return Processador{
		Id:         id,
		Cachezinha: InicializaCache(id),
	}
}

func (proc *Processador) RealizarLeitura(linha int, mp *Memoria, bp *BancoProcessadores) {
	cache_linha := proc.Cachezinha.ProcurarLinha(linha)

	if cache_linha != nil && cache_linha.Mesi != I {
		fmt.Println("Bloco encontrado na Cache! Read Hit")
		proc.Cachezinha.ReadHit(linha)
	} else {
		fmt.Println("Bloco não encontrado! Read Miss")
		proc.Cachezinha.ReadMiss(linha, mp, bp)
	}
}

func (proc *Processador) RealizarEscrita(linha int, reserva Reserva, mp *Memoria, bp *BancoProcessadores) {
	cache_linha := proc.Cachezinha.ProcurarLinha(linha)

	if cache_linha != nil {
		fmt.Println("Bloco encontrado na Cache! Write Hit")
		proc.Cachezinha.WriteHit(linha, reserva, mp, bp)
	} else {
		fmt.Println("Bloco não encontrado! Write Miss")
		proc.Cachezinha.WriteMiss(linha, reserva, mp, bp)
	}
}
