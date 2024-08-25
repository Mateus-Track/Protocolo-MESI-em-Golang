package componentes

import (
	"MESI/config"
	m "MESI/models"
	"fmt"
)

//--Memoria---------------------------------------------------------------------

type Memoria struct { //pelo menos 50 posições;
	livros [config.QUANTIDADE_LIVROS]m.Livro
}

func InicializaMemoria() Memoria {
	mp := Memoria{
		livros: [config.QUANTIDADE_LIVROS]m.Livro{},
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

	for i := 0; i < config.QUANTIDADE_LIVROS; i++ {
		secao := secoes[i/config.TAMANHO_BLOCO]
		nome := fmt.Sprintf("Livro %d", i)
		mp.livros[i] = m.Livro{
			Reservas: []m.Reserva{},
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

func (mp *Memoria) GuardarLinha(bloco int, livros [config.TAMANHO_BLOCO]m.Livro) {
	for i := 0; i < config.TAMANHO_BLOCO; i++ {
		mp.livros[bloco*config.TAMANHO_BLOCO+1] = livros[i]
	}
}

func (mp *Memoria) CarregarLinha(bloco int) [config.TAMANHO_BLOCO]m.Livro {
	linha_mp := [config.TAMANHO_BLOCO]m.Livro{}

	for i := 0; i < config.TAMANHO_BLOCO; i++ {
		linha_mp[i] = mp.livros[bloco*config.TAMANHO_BLOCO+i]
	}

	return linha_mp
}

func (mp *Memoria) Transferir_MP_Cache(cache *Cache, bp *BancoProcessadores, bloco int) *Linha {
	linha_mp := [config.TAMANHO_BLOCO]m.Livro{}

	for i := 0; i < config.TAMANHO_BLOCO; i++ {
		linha_mp[i] = mp.livros[bloco*config.TAMANHO_BLOCO+i]
	}

	return cache.CarregarLinha(linha_mp, bloco, mp, bp)
}

func (mp *Memoria) Transferir_Cache_MP(linha_cache *Linha, bloco int) {
	for i := 0; i < config.TAMANHO_BLOCO; i++ {
		mp.livros[bloco*config.TAMANHO_BLOCO+i] = linha_cache.Livros[i]
	}
}
