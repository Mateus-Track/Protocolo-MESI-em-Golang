package main

import "fmt"

func Printar_MP(memoria MP) {

	for i, livro := range memoria.Livros {
		fmt.Printf("Livro %d: Nome = %s, Seção = %s\n", i+1, livro.Nome, livro.Secao)
	}
}
