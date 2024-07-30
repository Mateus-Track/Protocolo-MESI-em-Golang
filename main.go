package main

import (
	"fmt"
	"time"
)

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
	Livros [4]Livro
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

}
