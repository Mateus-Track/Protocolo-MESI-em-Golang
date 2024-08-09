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

	for {
		var status bool = verificacao() //+ cache escolhida.
		if !status {
			return //quitar do sistema
		}

		linha := decide_linha()
		leitura := decide_operacao()                //se leitura = false, operação é de escrita.
		bloco := componentes.Linha_Conversao(linha) //converte a linha para qual o bloco que ela deve ser acessado.
		componentes.Printa_Cache(bp.BP[constantes.Cache_escolhida_int].Cachezinha)
		encontrado := componentes.Procura_Cache(bp.BP[constantes.Cache_escolhida_int].Cachezinha, bloco)

		transacao := componentes.Define_Transacao(encontrado, leitura, bloco, linha) //read ou write, hit ou miss. Está dando certo para miss, testar com HIT.

		componentes.Realiza_Transacao(transacao, &bp.BP[constantes.Cache_escolhida_int].Cachezinha, bloco, linha, mp)
		//bp.BP[cache_escolhida_int].cache.linhas[0].bloco)
		if leitura {
			fmt.Printf("Operacao de Leitura na linha %d", linha)
		} else {
			fmt.Printf("Operacao de Escrita na linha %d", linha)
		}
		//verificar se esse bloco está na Cache (cache_escolhida_int),
	}

}

func escolher_cache() {
	var cache_escolhida string = strconv.Itoa(constantes.QUANTIDADE_USUARIOS)
	//inválido por padrão.
	fmt.Printf("\nQual usuário da biblioteca você gostaria de controlar? Roger Waters (0), Slash (1), Jonathan Davis (2) ou Anitta (3)\n")
	fmt.Scan(&cache_escolhida)
	var err error
	constantes.Cache_escolhida_int, err = strconv.Atoi(cache_escolhida)
	for constantes.Cache_escolhida_int >= constantes.QUANTIDADE_USUARIOS || constantes.Cache_escolhida_int < 0 || err != nil {
		fmt.Printf("Usuário inexistente! Selecione um usuário válido, de 0 a %d\n", (constantes.QUANTIDADE_USUARIOS - 1))
		fmt.Scan(&cache_escolhida)
		constantes.Cache_escolhida_int, err = strconv.Atoi(cache_escolhida)
	}
	//fmt.Print(cache_escolhida_int)1
}

func exit() bool {
	var saida string = "2"
	for {
		fmt.Print("\nDeseja realizar uma operação no sistema? (1) Ou deseja finalizar? (0)\n")
		fmt.Scan(&saida)
		//print(saida)
		if saida != "0" && saida != "1" {
			fmt.Print("Não há uma operação com esse input! Por favor selecione uma opção válida!")
			saida = "2"
			continue
		} else if saida == "0" {
			return false
		} else {
			return true
		}
	}
}

func verificacao() bool {
	var estatos bool = exit()

	if estatos {
		escolher_cache()
		return true
	} else {
		return false //não quer escolher.
	}
}

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
