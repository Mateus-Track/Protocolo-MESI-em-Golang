package componentes

import "fmt"

type Linha struct {
	Livros [5]Livro
	Bloco  int //saber se o bloco foi puxado pra cache ou nao.
	Mesi   int8
}

func Linha_Conversao(linha int) int {
	return linha / 5
}

func (l Linha) PrintLinha() {
	fmt.Println("Linha:")
	for i, livro := range l.Livros {
		fmt.Printf("  Livro %d: %s\n", i+1, livro.Nome)
		fmt.Printf("  Secao %d: %s\n", i+1, livro.Secao)
	}
	fmt.Printf("  Bloco: %d\n", l.Bloco)
	fmt.Printf("  MESI: %d\n", l.Mesi)
}

func InicializaLinha() Linha {
	linha := Linha{
		Livros: [5]Livro{},
		Bloco:  -1, // Valor inicial para o bloco
		Mesi:   -1, // Valor inicial para MESI, meti o loco aq pra n começar em algum.	}
	}

	for i := range linha.Livros {
		linha.Livros[i] = InicializaLivro()
	}
	return linha
}
