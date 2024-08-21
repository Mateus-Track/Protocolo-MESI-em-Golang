package componentes

import (
	"MESI/constantes"
	"fmt"
)

type BancoProcessadores struct {
	BP []Processador
}

func InicializaBP(QUANTIDADE_USUARIOS int) BancoProcessadores {
	bp := BancoProcessadores{
		BP: make([]Processador, QUANTIDADE_USUARIOS),
	}

	// Inicializar cada Processador dentro do Banco
	for i := range bp.BP {
		bp.BP[i] = InicializaProcessador(i)
	}

	return bp
}

func (bp *BancoProcessadores) Notifica_Caches(linha int, tag_nova MesiFlags, cache_id int) {
	for i := 0; i < constantes.QUANTIDADE_USUARIOS; i++ {
		fmt.Print("To no loop")
		cache := bp.BP[i].Cachezinha

		if cache.id != cache_id { // nao vai mudar a própria tag, mudar a dos outros.
			cache_index := cache.Procura_Cache(linha)
			if cache_index >= 0 { //tem nessa cache esse bloco, ver o que fazer com sua TAG_MESI.
				linha_analisada := cache.Linhas[cache_index]
				if linha_analisada.Mesi == E && tag_nova == S {
					fmt.Print("Achei! Mudando Tag!")
					//linha_analisada.Mesi = constantes.S //se eu fizer isso, nao muda. é uma cópia da referencia só.
					cache.Linhas[cache_index].Mesi = S
				}

			}

		}

	}
}
