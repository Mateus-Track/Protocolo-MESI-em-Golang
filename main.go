package main

import (
	"MESI/componentes"
	"MESI/constantes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Bem vindo à Biblioteca Matheus Meu Deus!")
	// fmt.Println("Carregando Memória Principal")
	mp := componentes.PreencherLivros()
	componentes.Printar_MP(mp)
	bp := componentes.InicializaBP(constantes.QUANTIDADE_USUARIOS)

	for !should_exit() {
		cache_num := escolher_cache()

		linha := decide_linha()
		leitura := decide_operacao()

		//componentes.Printa_Cache(bp.BP[constantes.Cache_escolhida_int].Cachezinha)
		cache := &bp.BP[cache_num].Cachezinha

		// componentes.Printa_Cache(*cache_escolhida)
		cache_index := cache.Procura_Cache(linha)

		fmt.Printf("Cache index = %d", cache_index)
		switch {
		case cache_index >= 0 && leitura:
			fmt.Println("Bloco encontrado na Cache! Read Hit")
			cache.Read_Hit(linha)
		case cache_index >= 0:
			fmt.Println("Bloco encontrado na Cache! Write Hit")
			//componentes.Write_Hit(&cache_escolhida, linha)
		case leitura:
			fmt.Println("Bloco não encontrado! Read Miss")
			cache.Read_Miss(linha, &mp, &bp)
		case cache_index < 0 && !leitura:
			fmt.Println("Bloco não encontrado! Write Miss")
			//return constantes.WM
		default:
			panic("Erro, não encontrada essa transação.")
			// return -1
		}

		//transacao := componentes.Define_Transacao(cache_index >= 0, leitura) //read ou write, hit ou miss. Está dando certo para miss, testar com HIT.

		//omponentes.Realiza_Transacao(transacao, &cache_escolhida, linha, mp, cache_index)

		if leitura {
			fmt.Printf("Operacao de Leitura na linha %d", linha)
		} else {
			fmt.Printf("Operacao de Escrita na linha %d", linha)
		}
		//verificar se esse bloco está na Cache (cache_escolhida_int),
	}

}

func escolher_cache() uint {
	var cache_num uint

	fmt.Printf("\nQual usuário da biblioteca você gostaria de controlar? Roger Waters (0), Slash (1), Jonathan Davis (2) ou Anitta (3)\n")
	_, err := fmt.Scan(&cache_num)

	for err != nil || cache_num >= constantes.QUANTIDADE_USUARIOS {
		fmt.Printf("Usuário inexistente! Selecione um usuário válido, de 0 a %d\n", (constantes.QUANTIDADE_USUARIOS - 1))
		_, err = fmt.Scan(&cache_num)
	}

	fmt.Print(cache_num)
	return cache_num
}

func should_exit() bool {
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

// func verificacao() bool {
// 	estatos, err := should_exit()

// 	if estatos {
// 		escolher_cache()
// 		return true
// 	} else {
// 		return false //não quer escolher.
// 	}
// }

func decide_operacao() bool {
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

func decide_linha() int {
	var linha_escolhida string

	fmt.Printf("\nQual livro da biblioteca você gostaria de acessar? Selecione de 0 a %d\n", (constantes.QUANTIDADE_LIVROS - 1))
	fmt.Scan(&linha_escolhida)
	linha_escolhida_int, err := strconv.Atoi(linha_escolhida)
	for linha_escolhida_int >= constantes.QUANTIDADE_LIVROS || linha_escolhida_int < 0 || err != nil {
		fmt.Printf("Livro inexistente! Selecione um livro válido, de 0 a %d\n", (constantes.QUANTIDADE_LIVROS - 1))
		fmt.Scan(&linha_escolhida)
		linha_escolhida_int, err = strconv.Atoi(linha_escolhida)
	}
	//fmt.Print(cache_escolhida_int)
	return linha_escolhida_int
}
