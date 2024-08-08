package componentes

import (
	"MESI/constantes"
	"fmt"
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
