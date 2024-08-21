package componentes

import (
	"MESI/constantes"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Cache struct { //pelo menos 5 posições.
	id     int
	Linhas [constantes.QUANTIDADE_LINHAS_CACHE]Linha //acredito que serão as mesmas linhas, só uma cópia da MP.
	Fila   []uint8
}

func InicializaCache(id int) Cache {
	cache := Cache{
		id:     id,
		Linhas: [5]Linha{},
		Fila:   []uint8{},
	}

	// Inicializando cada Linha dentro da Cache
	for i := range cache.Linhas {
		cache.Linhas[i] = InicializaLinha()
	}

	return cache
}

func (cache *Cache) Procura_Cache(linha int) int {
	bloco := linha / 5

	for i := 0; i < constantes.QUANTIDADE_LINHAS_CACHE; i++ {
		if cache.Linhas[i].Bloco == bloco {
			return i
		}
	}

	return -1
}

func (cache *Cache) Printa_Cache() { //sem a fila por enquanto.
	i := 0
	for i < constantes.QUANTIDADE_LINHAS_CACHE {
		cache.Linhas[i].PrintLinha()
		i++
	}

}

func (cache *Cache) Read_Miss(linha int, mp *MP, bp *BancoProcessadores) { //por enquanto, n vou ver as TAGS, apenas puxar sem pensar.
	tag_nova, tag_err := cache.Verifica_MESI(linha, mp)
	if tag_err != nil {
		panic("Erro ao Verificar MESI da MP. Abortando.")
	}

	if tag_nova == S { //teste nesse caso
		bp.Notifica_Caches(linha, tag_nova, cache.id)
	}

	bloco := linha / 5

	if cache.TemEspacoLivre() {
		disponiveis := cache.ValoresDisponiveis()
		numero_random := Gerar_Aleatorio(len(disponiveis))
		posicao := disponiveis[numero_random]
		cache.Fila = append(cache.Fila, posicao)
		TransferirMPCache(mp, cache, (bloco * 5), posicao)
		cache.Linhas[posicao].Bloco = linha / 5
		cache.Linhas[posicao].Mesi = tag_nova
		cache.Printa_Cache()
	} else { //tirar da fila, dar append no final lá, retornar pra mp.
		primeiroElemento := cache.Fila[0]
		cache.Fila = cache.Fila[1:]
		TransferirCacheMP(mp, cache, (cache.Linhas[primeiroElemento].Bloco * 5), primeiroElemento) //transferir de volta pra MP, fazer SÓ QUANDO HOUVE mudança.
		cache.Fila = append(cache.Fila, primeiroElemento)
		cache.Linhas[primeiroElemento].Bloco = linha / 5
		TransferirMPCache(mp, cache, (bloco * 5), primeiroElemento)
		cache.Printa_Cache()
	}

}

func (cache *Cache) Read_Hit(linha int) {
	index := cache.Procura_Cache(linha)
	livro := cache.Linhas[index].Livros[linha%5]

	fmt.Printf("Livro: %s\n", livro.Nome)
	fmt.Printf("Secao: %s\n", livro.Secao)
	fmt.Printf("MESI do bloco: %d\n", cache.Linhas[index].Mesi)
}

func (cache *Cache) Write_Miss(linha int, mp *MP, bp *BancoProcessadores) {
	//tag_nova, tag_err := cache.Verifica_MESI(linha, mp) // nao precisa né? em Write a MP vai ficar "I" e a cache "M"
	bloco := linha / 5

	mp.Tags[bloco] = I
	tag_nova := M

	// if tag_nova == S { //teste nesse caso
	// 	bp.Notifica_Caches(linha, tag_nova, cache.id)
	// }

	if cache.TemEspacoLivre() {
		disponiveis := cache.ValoresDisponiveis()
		numero_random := Gerar_Aleatorio(len(disponiveis))
		posicao := disponiveis[numero_random]
		cache.Fila = append(cache.Fila, posicao)
		TransferirMPCache(mp, cache, (bloco * 5), posicao)
		cache.Linhas[posicao].Bloco = linha / 5
		cache.Linhas[posicao].Mesi = tag_nova // até aqui é igual nos READ MISS, dá pra transformar numa função tipo: "ManuseiaFilaCache(), antes do transferir, sla."
		//fazer escrita.
		cache.Printa_Cache()
	} else { //tirar da fila, dar append no final lá, retornar pra mp.
		primeiroElemento := cache.Fila[0]
		cache.Fila = cache.Fila[1:]
		TransferirCacheMP(mp, cache, (cache.Linhas[primeiroElemento].Bloco * 5), primeiroElemento) //transferir de volta pra MP, fazer SÓ QUANDO HOUVE mudança.
		cache.Fila = append(cache.Fila, primeiroElemento)
		cache.Linhas[primeiroElemento].Bloco = linha / 5
		TransferirMPCache(mp, cache, (bloco * 5), primeiroElemento)
		//fazer escrita.
		cache.Printa_Cache()
	}
}

func (cache *Cache) Write_Hit(linha int, mp *MP, bp *BancoProcessadores) {
	bloco := linha / 5
	index := cache.Procura_Cache(linha)
	livro := cache.Linhas[index].Livros[linha%5]

	if cache.Linhas[index].Mesi == E {
		cache.Linhas[index].Mesi = M
	}
	mp.Tags[bloco] = I

	//bp.Notifica_Caches(linha,M,index)

	//realizar escrita.
	fmt.Printf("Livro: %s\n", livro.Nome)
	fmt.Printf("Secao: %s\n", livro.Secao)
	fmt.Printf("MESI do bloco: %d\n", cache.Linhas[index].Mesi)
}

func (c *Cache) TemEspacoLivre() bool {
	return len(c.Fila) < constantes.QUANTIDADE_LINHAS_CACHE
}

func (c *Cache) ValoresDisponiveis() []uint8 { // n entendi totalmente.
	todosValores := []uint8{0, 1, 2, 3, 4}

	presente := make(map[uint8]bool)
	for _, v := range c.Fila {
		presente[v] = true
	}

	// Cria uma lista para armazenar valores disponíveis
	disponiveis := []uint8{}
	for _, v := range todosValores {
		if !presente[v] {
			disponiveis = append(disponiveis, v)
		}
	}

	return disponiveis
}

func (cache *Cache) Verifica_MESI(linha int, mp *MP) (MesiFlags, error) {
	bloco := linha / 5
	tag_bloco := mp.Tags[bloco]

	switch {
	case tag_bloco == UNDEFINED:
		//colocar a tag de Exclusivo na MP também.
		mp.Tags[bloco] = E
		return E, nil //nao foi puxado ainda, será exclusivo.
	case tag_bloco == E:
		mp.Tags[bloco] = S
		return S, nil //E VAI TER QUE ENVIAR ISSO PARA o outro processador que contém esse dado, nao fiz ainda.
	case tag_bloco == S:
		return S, nil //continua shared
	case tag_bloco == I:
		//vai ter que buscar no lugar correto.
	}

	return 0, errors.New("InvalidMesiFlag") //erro
}

func Gerar_Aleatorio(quantidade int) int16 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return int16(r.Intn(quantidade))
}

// func Define_Transacao(encontrado bool, leitura bool) {
// 	if encontrado && leitura {
// 	} else if encontrado {
// 	} else if leitura {
// 	} else {
// 	}
// }
