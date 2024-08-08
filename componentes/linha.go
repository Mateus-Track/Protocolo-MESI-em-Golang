package componentes

type Linha struct {
	Livros [5]Livro
	Bloco  int //saber se o bloco foi puxado pra cache ou nao.
	Mesi   uint8
}
