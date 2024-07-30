package main

import (
	"fmt"
	"strconv"
	"time"
)

const QUANTIDADE_CACHES = 4
const EXIT = false

const (
	E = iota //0
	S
	M
	I
)

// type Cache_Acoes interface {
// 	Ler()
// }

type Cache struct { //pelo menos 5 posições.
	Linhas []Linha //acredito que serão as mesmas linhas, só uma cópia da MP.
}

type Linha struct {
	Livros [5]Livro
	Mesi   uint8
}

type MP struct { //pelo menos 50 posições;
	Livros [50]Livro
}

type Livro struct {
	Reservas [][2]time.Time
	Nome     string
	Secao    string
}

type Processadores struct {
	Caches []Cache
}

func main() {
	fmt.Println("Hello World!")

	for {
		var status bool = verificacao()
		if !status {
			return //quitar do sistema
		}
	}

}

func escolher_cache() {
	var cache_escolhida string = strconv.Itoa(QUANTIDADE_CACHES)
	//inválido por padrão.
	fmt.Printf("\nQual cache você gostaria de usar? Selecione de 0 a %d\n", (QUANTIDADE_CACHES - 1))
	fmt.Scan(&cache_escolhida)
	cache_escolhida_int, err := strconv.Atoi(cache_escolhida)
	for cache_escolhida_int >= QUANTIDADE_CACHES || cache_escolhida_int < 0 || err != nil {
		fmt.Printf("Cache inexistente! Selecione uma cache válida, de 0 a %d\n", (QUANTIDADE_CACHES - 1))
		fmt.Scan(&cache_escolhida)
	}
	fmt.Print(cache_escolhida)
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
