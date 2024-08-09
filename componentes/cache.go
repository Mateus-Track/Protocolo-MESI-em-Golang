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

func Procura_Cache(cache Cache, bloco int) bool { //verifica se o bloco está na cache.
	i := 0
	for i < constantes.QUANTIDADE_LINHAS_CACHE {
		if cache.Linhas[i].Bloco == bloco {
			fmt.Printf("Achou! - %d com %d", cache.Linhas[i].Bloco, bloco)
			return true
		}
		i++
	}

	return false
}

func Printa_Cache(cache Cache) { //sem a fila por enquanto.
	fmt.Println("Cache:")
	i := 0
	for i < constantes.QUANTIDADE_LINHAS_CACHE {
		cache.Linhas[i].PrintLinha()
		i++
	}
	fmt.Println("Fila da Cache: ")
	fmt.Println("Fila:", cache.Fila)

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

func Define_Transacao(encontrado, leitura bool, bloco, linha int) int16 { //não sei se é melhor passar ja o bloco e linha e chamar as transações ou fazer separado.
	switch {
	case encontrado && leitura:
		fmt.Println("Bloco encontrado na Cache! Read Hit")
		return constantes.RH
	case encontrado:
		fmt.Println("Bloco encontrado na Cache! Write Hit")
		return constantes.WH
	case leitura:
		fmt.Println("Bloco não encontrado! Read Miss")
		return constantes.RM
		//Read_Miss(constantes.Cache_escolhida_int, bloco, linha)
	case !encontrado && !leitura:
		fmt.Println("Bloco não encontrado! Write Miss")
		return constantes.WM
	default:
		panic("Erro, não encontrada essa transação.")
		// return -1
	}

}

func Realiza_Transacao(transacao int16, c *Cache, bloco, linha int, mp MP) {
	switch transacao {
	case constantes.RH:
		//Read_Hit()
	case constantes.RM:
		Read_Miss(c, bloco, linha, mp)
	case constantes.WH:
		//Write_Hit()
	case constantes.WM:
		//Write_Miss()
	}
}

func Read_Miss(c *Cache, bloco, linha int, mp MP) { //por enquanto, n vou ver as TAGS, apenas puxar sem pensar.
	index := bloco * 5

	if c.TemEspacoLivre() {
		disponiveis := c.ValoresDisponiveis()
		numero_random := Gerar_Aleatorio(len(disponiveis))
		posicao := disponiveis[numero_random]
		fmt.Print(posicao)
		c.Fila = append(c.Fila, posicao)
		TransferirMPCache(mp, c, index, posicao, constantes.TAMANHO_BLOCO)
		c.Linhas[posicao].Bloco = bloco
		Printa_Cache(*c)
	} else { //tirar da fila, dar append no final lá, retornar pra mp.
		fmt.Print("la ele")
	}

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

// func Define_Transacao(encontrado bool, leitura bool) {
// 	if encontrado && leitura {
// 		fmt.Println("Bloco encontrado na Cache! Read Hit")
// 	} else if encontrado {
// 		fmt.Println("Bloco encontrado na Cache! Write Hit")
// 	} else if leitura {
// 		fmt.Println("Bloco não encontrado! Read Miss")
// 	} else {
// 		fmt.Println("Bloco não encontrado! Write Miss")
// 	}
// }
