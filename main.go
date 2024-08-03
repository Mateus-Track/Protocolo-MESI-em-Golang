package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Carregando Memória Principal")
	mp := preencherLivros()
	Printar_MP(mp)

	for {
		var status bool = verificacao() //+ cache escolhida.
		if !status {
			return //quitar do sistema
		}

		linha := decide_linha()
		leitura := decide_operacao() //se leitura = false, operação é de escrita.

		if leitura {
			fmt.Printf("Operacao de Leitura na linha %d", linha)
		} else {
			fmt.Printf("Operacao de Escrita na linha %d", linha)
		}
		//verificar se esse bloco está na Cache (cache_escolhida_int),
	}

}

func escolher_cache() {
	var cache_escolhida string = strconv.Itoa(QUANTIDADE_USUARIOS)
	//inválido por padrão.
	fmt.Printf("\nQual usuário da biblioteca você gostaria de controlar? Selecione de 0 a %d\n", (QUANTIDADE_USUARIOS - 1))
	fmt.Scan(&cache_escolhida)
	var err error
	cache_escolhida_int, err = strconv.Atoi(cache_escolhida)
	for cache_escolhida_int >= QUANTIDADE_USUARIOS || cache_escolhida_int < 0 || err != nil {
		fmt.Printf("Usuário inexistente! Selecione um usuário válido, de 0 a %d\n", (QUANTIDADE_USUARIOS - 1))
		fmt.Scan(&cache_escolhida)
		cache_escolhida_int, err = strconv.Atoi(cache_escolhida)
	}
	//fmt.Print(cache_escolhida_int)
}

func exit() bool {
	var saida string = "2"
	for {
		fmt.Print("\nDeseja continuar operando no sistema (1) ou deseja finalizar (0)\n")
		fmt.Scan(&saida)
		//print(saida)
		if saida != "0" && saida != "1" {
			fmt.Print("ERRO! Selecione uma opção válida!")
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

func preencherLivros() MP {
	secoes := []string{
		"Tecnologia",
		"Matemática",
		"História",
		"Literatura",
		"Filosofia",
		"Ciência",
		"Arte",
		"Geografia",
		"Economia",
		"Psicologia",
	}

	var mp MP
	for i := 0; i < 50; i++ {
		secao := secoes[i/5]
		nome := fmt.Sprintf("Livro %d", i+1)
		mp.Livros[i] = Livro{
			Reservas: [][2]time.Time{},
			Nome:     nome,
			Secao:    secao,
		}
	}
	return mp
}

func decide_operacao() bool {
	var saida string = "2"
	for {
		fmt.Print("\nDeseja realizar uma operação de leitura (1) ou escrita (0)\n")
		fmt.Scan(&saida)
		//print(saida)
		if saida != "0" && saida != "1" {
			fmt.Print("ERRO! Selecione uma opção válida!")
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

	fmt.Printf("\nQual livro da biblioteca você gostaria de acessar? Selecione de 0 a %d\n", (QUANTIDADE_LIVROS - 1))
	fmt.Scan(&linha_escolhida)
	linha_escolhida_int, err := strconv.Atoi(linha_escolhida)
	for linha_escolhida_int >= QUANTIDADE_LIVROS || linha_escolhida_int < 0 || err != nil {
		fmt.Printf("Livro inexistente! Selecione um livro válido, de 0 a %d\n", (QUANTIDADE_LIVROS - 1))
		fmt.Scan(&linha_escolhida)
		linha_escolhida_int, err = strconv.Atoi(linha_escolhida)
	}
	//fmt.Print(cache_escolhida_int)
	return linha_escolhida_int
}
