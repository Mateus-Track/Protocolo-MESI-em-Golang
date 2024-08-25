package main

import (
	"MESI/componentes"
	"MESI/constantes"
	"fmt"
	"strconv"
	"time"
)

const QUANTIDADE_LIVROS = 50

func main() {
	fmt.Println("Bem vindo à Biblioteca Matheus Meu Deus!")

	mp := componentes.InicializaMemoria()
	mp.PreencherLivros()
	mp.Print()

	bp := componentes.InicializaBP()

	for !deveSair() {
		usuario := escolherProcessador(&bp)

		linha := decidirLinha()
		leitura := decidirOperacao()

		if leitura {
			usuario.RealizarLeitura(linha, &mp, &bp)

			fmt.Printf("Operacao de Leitura na linha %d", linha)
		} else {
			reserva := lerReserva(usuario)
			usuario.RealizarEscrita(linha, reserva, &mp, &bp)

			fmt.Printf("Operacao de Escrita na linha %d", linha)
		}
	}

}

func escolherProcessador(bp *componentes.BancoProcessadores) *componentes.Processador {
	var proc_num int

	var proc *componentes.Processador
	var proc_err error

	fmt.Printf("\nQual usuário da biblioteca você gostaria de controlar? Roger Waters (0), Slash (1), Jonathan Davis (2) ou Anitta (3)\n")
	_, err := fmt.Scan(&proc_num)

	if err == nil {
		proc, proc_err = bp.SelecionarProcessador(proc_num)
	}

	for err != nil || proc_err != nil {
		fmt.Printf("Usuário inexistente! Selecione um usuário válido, de 0 a 4 (inclusivo)\n")
		_, err = fmt.Scan(&proc_num)
	}

	fmt.Print(proc_num)
	return proc
}

func deveSair() bool {
	var saida int = -1

	for saida < 0 {
		fmt.Print("\nDeseja realizar uma operação no sistema? (1) Ou deseja finalizar? (0)\n")
		_, err := fmt.Scan(&saida)

		if err != nil || (saida != 0 && saida != 1) {
			fmt.Print("Não há uma operação com esse input! Por favor selecione uma opção válida!")
			saida = -1
		}
	}

	return saida == 0
}

func decidirOperacao() bool {
	var saida string = "2"
	for {
		fmt.Print("\nDeseja consultar as reservas do livro (1 - Operação de Leitura) ou registrar uma nova reserva nele (0 - Operação de Escrita)\n")
		fmt.Scan(&saida)
		//print(saida)
		if saida != "0" && saida != "1" {
			fmt.Print("Não há uma operação com esse input! Por favor selecione uma opção válida!")
			saida = "2"
			continue
		} else if saida == "0" {
			return false //é escrita.
		} else {
			return true //é leitura
		}
	}

}

func decidirLinha() int {
	var linha_escolhida string

	fmt.Printf("\nQual livro da biblioteca você gostaria de acessar? Selecione de 0 a %d\n", (QUANTIDADE_LIVROS - 1))
	fmt.Scan(&linha_escolhida)
	linha_escolhida_int, err := strconv.Atoi(linha_escolhida)

	for linha_escolhida_int >= QUANTIDADE_LIVROS || linha_escolhida_int < 0 || err != nil {
		fmt.Printf("Livro inexistente! Selecione um livro válido, de 0 a %d\n", (QUANTIDADE_LIVROS - 1))
		fmt.Scan(&linha_escolhida)
		linha_escolhida_int, err = strconv.Atoi(linha_escolhida)
	}

	return linha_escolhida_int
}

func lerReserva(proc *componentes.Processador) componentes.Reserva {
	var data_inicio, data_fim time.Time
	var buffer string
	var err error

	for {
		fmt.Printf("\nQual a data de início da reserva? Escreva da seguinte maneira (DD-MM-YYYY):\n")

		for {
			fmt.Scan(&buffer)
			data_inicio, err = time.Parse(constantes.TIME_LAYOUT, buffer)

			if err != nil {
				fmt.Printf("\nInsira uma data válida (DD-MM-YYYY):\n")
				continue
			}
			break
		}

		for {
			fmt.Printf("\nQual a data de fim da reserva? Escreva da seguinte maneira (DD-MM-YYYY):\n")
			fmt.Scan(&buffer)
			data_fim, err = time.Parse(constantes.TIME_LAYOUT, buffer)

			if err != nil {
				fmt.Printf("\nInsira uma data válida (DD-MM-YYYY):\n")
				continue
			}
			break
		}

		reserva, err := componentes.InicializaReserva(data_inicio, data_fim, proc.Id)
		if err != nil {
			panic(err)
		}

		return reserva
	}
}
