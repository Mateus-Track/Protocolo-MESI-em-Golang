package componentes

import (
	"MESI/constantes"
	"fmt"
	"math/rand"
	"time"
)

type Cache struct { //pelo menos 5 posições.
	Linhas [constantes.QUANTIDADE_LINHAS_CACHE]Linha //acredito que serão as mesmas linhas, só uma cópia da MP.
	Fila   []uint8
}

func Procura_Cache(cache Cache, linha int) int { //verifica se o bloco está na cache.

	bloco := linha / 5
	i := 0
	for i < constantes.QUANTIDADE_LINHAS_CACHE {
		if cache.Linhas[i].Bloco == bloco {
			return i
		}

		i++
	}

	return -1
}

func Printa_Cache(cache Cache) { //sem a fila por enquanto.
	i := 0
	for i < constantes.QUANTIDADE_LINHAS_CACHE {
		cache.Linhas[i].PrintLinha()
		i++
	}

}

func InicializaCache() Cache {
	cache := Cache{
		Linhas: [5]Linha{},
		Fila:   []uint8{},
	}

	// Inicializando cada Linha dentro da Cache
	for i := range cache.Linhas {
		cache.Linhas[i] = InicializaLinha()
	}

	return cache
}

func Define_Transacao(encontrado, leitura bool) int16 { //não sei se é melhor passar ja o bloco e linha e chamar as transações ou fazer separado.
	switch {
	case encontrado && leitura:
		return constantes.RH
	case encontrado:
		return constantes.WH
	case leitura:
		return constantes.RM
		//Read_Miss(constantes.Cache_escolhida_int, bloco, linha)
	case !encontrado && !leitura:
		return constantes.WM
	default:
		panic("Erro, não encontrada essa transação.")
		// return -1
	}

}

func Realiza_Transacao(transacao int16, c *Cache, linha int, mp MP, cache_index int) {

}

func Read_Miss(c *Cache, linha int, mp *MP, bp *BancoProcessadores) { //por enquanto, n vou ver as TAGS, apenas puxar sem pensar.
	tag_nova := Verifica_MESI(c, linha, mp)
	if tag_nova == -10 {
		panic("Erro ao Verificar MESI da MP. Abortando.")
	}

	if tag_nova == constantes.S { //teste nesse caso
		Notifica_Caches(bp, linha, tag_nova, c)
	}

	bloco := linha / 5

	if c.TemEspacoLivre() {
		disponiveis := c.ValoresDisponiveis()
		numero_random := Gerar_Aleatorio(len(disponiveis))
		posicao := disponiveis[numero_random]
		c.Fila = append(c.Fila, posicao)
		TransferirMPCache(mp, c, (bloco * 5), posicao)
		c.Linhas[posicao].Bloco = linha / 5
		c.Linhas[posicao].Mesi = tag_nova
		Printa_Cache(*c)
	} else { //tirar da fila, dar append no final lá, retornar pra mp.
		primeiroElemento := c.Fila[0]
		c.Fila = c.Fila[1:]
		TransferirCacheMP(mp, c, (c.Linhas[primeiroElemento].Bloco * 5), primeiroElemento) //transferir de volta pra MP, fazer SÓ QUANDO HOUVE mudança.
		c.Fila = append(c.Fila, primeiroElemento)
		c.Linhas[primeiroElemento].Bloco = linha / 5
		TransferirMPCache(mp, c, (bloco * 5), primeiroElemento)
		Printa_Cache(*c)
	}

}

func Read_Hit(c *Cache, linha int) {

	index := Procura_Cache(*c, linha)
	livro := c.Linhas[index].Livros[linha%5]
	fmt.Printf("Livro: %s\n", livro.Nome)
	fmt.Printf("Secao: %s\n", livro.Secao)
	fmt.Printf("MESI do bloco: %d\n", c.Linhas[index].Mesi)

}

func Gerar_Aleatorio(quantidade int) int16 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return int16(r.Intn(quantidade))
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

func Verifica_MESI(c *Cache, linha int, mp *MP) int8 {
	bloco := linha / 5
	tag_bloco := mp.Tags[bloco]

	switch {
	case tag_bloco == -1:
		//colocar a tag de Exclusivo na MP também.
		mp.Tags[bloco] = constantes.E
		return constantes.E //nao foi puxado ainda, será exclusivo.
	case tag_bloco == constantes.E:
		mp.Tags[bloco] = constantes.S
		return constantes.S //E VAI TER QUE ENVIAR ISSO PARA o outro processador que contém esse dado, nao fiz ainda.
	case tag_bloco == constantes.S:
		return constantes.S //continua shared
	case tag_bloco == constantes.I:
		//vai ter que buscar no lugar correto.
	}

	return -10 //erro
}

func Notifica_Caches(bp *BancoProcessadores, linha int, tag_nova int8, c *Cache) {

	for i := 0; i < constantes.QUANTIDADE_USUARIOS; i++ {
		fmt.Print("To no loop")
		if i != constantes.Cache_escolhida_int { // nao vai mudar a própria tag, mudar a dos outros.
			cache_index := Procura_Cache(bp.BP[i].Cachezinha, linha)
			if cache_index >= 0 { //tem nessa cache esse bloco, ver o que fazer com sua TAG_MESI.
				linha_analisada := bp.BP[i].Cachezinha.Linhas[cache_index]
				if linha_analisada.Mesi == constantes.E && tag_nova == constantes.S {
					fmt.Print("Achei! Mudando Tag!")
					//linha_analisada.Mesi = constantes.S //se eu fizer isso, nao muda. é uma cópia da referencia só.
					bp.BP[i].Cachezinha.Linhas[cache_index].Mesi = constantes.S
				}

			}

		}

	}
}

// func Define_Transacao(encontrado bool, leitura bool) {
// 	if encontrado && leitura {
// 	} else if encontrado {
// 	} else if leitura {
// 	} else {
// 	}
// }
