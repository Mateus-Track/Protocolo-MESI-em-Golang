package componentes

type Linha struct {
	Livros [5]Livro
	Bloco  int //saber se o bloco foi puxado pra cache ou nao.
	Mesi   uint8
}

func Linha_Conversao(linha int) int {
	return linha / 5
}
