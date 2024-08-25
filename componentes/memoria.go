package componentes

import (
	"fmt"
)

const TAMANHO_BLOCO = 5

//--Memoria---------------------------------------------------------------------

type Memoria struct { //pelo menos 50 posições;
	livros [50]Livro
	// Tags   [10]MesiFlags //guardar na MP as tags, facilitar.
}

func InicializaMemoria() Memoria {
	mp := Memoria{
		livros: [50]Livro{},
	}

	return mp
}

func (mp *Memoria) PreencherLivros() {
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

	for i := 0; i < 50; i++ {
		secao := secoes[i/5]
		nome := fmt.Sprintf("Livro %d", i)
		mp.livros[i] = Livro{
			Reservas: []Reserva{},
			Nome:     nome,
			Secao:    secao,
		}
	}
}

func (mp *Memoria) Print() {
	for i, livro := range mp.livros {
		fmt.Printf("Livro %d: Nome = %s, Seção = %s\n", i, livro.Nome, livro.Secao)
	}
}

func (mp *Memoria) GuardarLinha(bloco int, livros [TAMANHO_BLOCO]Livro) {
	for i := 0; i < TAMANHO_BLOCO; i++ {
		mp.livros[bloco*TAMANHO_BLOCO+1] = livros[i]
	}
}

func (mp *Memoria) CarregarLinha(bloco int) [TAMANHO_BLOCO]Livro {
	linha_mp := [TAMANHO_BLOCO]Livro{}

	for i := 0; i < TAMANHO_BLOCO; i++ {
		linha_mp[i] = mp.livros[bloco*TAMANHO_BLOCO+i]
	}

	return linha_mp
}

func (mp *Memoria) Transferir_MP_Cache(cache *Cache, bp *BancoProcessadores, bloco int) *Linha {
	linha_mp := [TAMANHO_BLOCO]Livro{}

	for i := 0; i < TAMANHO_BLOCO; i++ {
		linha_mp[i] = mp.livros[bloco*TAMANHO_BLOCO+i]
	}

	return cache.CarregarLinha(linha_mp, bloco, mp, bp)
}

func (mp *Memoria) Transferir_Cache_MP(linha_cache *Linha, bloco int) {
	for i := 0; i < TAMANHO_BLOCO; i++ {
		mp.livros[bloco*TAMANHO_BLOCO+i] = linha_cache.Livros[i]
	}
}
