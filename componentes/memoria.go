package componentes

import (
	"MESI/constantes"
	"fmt"
)

type MP struct { //pelo menos 50 posições;
	Livros [50]Livro
	// Tags   [10]MesiFlags //guardar na MP as tags, facilitar.
}

func PreencherLivros() MP {
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
		nome := fmt.Sprintf("Livro %d", i)
		mp.Livros[i] = Livro{
			Reservas: []Reserva{},
			Nome:     nome,
			Secao:    secao,
		}
	}

	// for i := range mp.Tags {
	// 	mp.Tags[i] = UNDEFINED
	// }

	return mp
}

func Printar_MP(memoria MP) {

	for i, livro := range memoria.Livros {
		fmt.Printf("Livro %d: Nome = %s, Seção = %s\n", i, livro.Nome, livro.Secao)
	}
}

func Transferir_MP_Cache(mp *MP, cache *Cache, bp *BancoProcessadores, bloco int) *Linha {
	linha_mp := [5]Livro{}

	for i := 0; i < constantes.TAMANHO_BLOCO; i++ {
		linha_mp[i] = mp.Livros[bloco*constantes.TAMANHO_BLOCO+i]
	}

	return cache.Carregar_Linha(linha_mp, bloco, mp, bp)
}

func Transferir_Cache_MP(mp *MP, linha_cache *Linha, bloco int) {
	for i := 0; i < constantes.TAMANHO_BLOCO; i++ {
		mp.Livros[bloco*constantes.TAMANHO_BLOCO+i] = linha_cache.Livros[i]
	}
}
