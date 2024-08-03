package main

import "time"

const QUANTIDADE_USUARIOS = 4
const QUANTIDADE_LIVROS = 50

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
	Linhas [5]Linha //acredito que serão as mesmas linhas, só uma cópia da MP.
}

type Linha struct {
	Livros [5]Livro
	Bloco  int //saber se o bloco foi puxado pra cache ou nao.
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
