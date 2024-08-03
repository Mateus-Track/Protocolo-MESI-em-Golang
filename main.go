package main

import (
	"fmt"
	"strconv"
	"time"
)

const QUANTIDADE_CACHES = 4
const EXIT = false

var cache_escolhida_int int

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
	fmt.Println("Carregando Memória Principal")
	mp := preencherLivros()
	printar_MP(mp)

	for {
		var status bool = verificacao()
		if !status {
			return //quitar do sistema
		}
		//fmt.Print(cache_escolhida_int)
	}

}

func escolher_cache() {
	var cache_escolhida string = strconv.Itoa(QUANTIDADE_CACHES)
	//inválido por padrão.
	fmt.Printf("\nQual usuário da biblioteca você gostaria de controlar? Selecione de 0 a %d\n", (QUANTIDADE_CACHES - 1))
	fmt.Scan(&cache_escolhida)
	var err error
	cache_escolhida_int, err = strconv.Atoi(cache_escolhida)
	for cache_escolhida_int >= QUANTIDADE_CACHES || cache_escolhida_int < 0 || err != nil {
		fmt.Printf("Usuário inexistente! Selecione um usuário válido, de 0 a %d\n", (QUANTIDADE_CACHES - 1))
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

func printar_MP(memoria MP) {

	for i, livro := range memoria.Livros {
		fmt.Printf("Livro %d: Nome = %s, Seção = %s\n", i+1, livro.Nome, livro.Secao)
	}
}
