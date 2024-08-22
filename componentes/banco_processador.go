package componentes

import (
	"MESI/constantes"
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

func (bp *BancoProcessadores) Verificar_MESI(linha int) (bool, MesiFlags, *Linha) {
	for _, p := range bp.BP {
		flag, linha_cache, err := p.Cachezinha.Status_Cache(linha)

		if err == nil && flag != I {
			return true, flag, linha_cache
		}
	}

	return false, 0, nil
}

func (bp *BancoProcessadores) Atualiza_Shared(linha int, cache_id int) {
	for i := 0; i < constantes.QUANTIDADE_USUARIOS; i++ {
		cache := &bp.BP[i].Cachezinha

		if cache.id_processador != cache_id {
			linha_cache := cache.Procura_Cache(linha)
			if linha_cache != nil {
				linha_cache.Mesi = S
			}
		}
	}
}

func (bp *BancoProcessadores) Atualiza_Invalid(linha int, cache_id int) {
	for i := 0; i < constantes.QUANTIDADE_USUARIOS; i++ {
		cache := &bp.BP[i].Cachezinha

		if cache.id_processador != cache_id {
			linha_cache := cache.Procura_Cache(linha)
			if linha_cache != nil {
				linha_cache.Mesi = I
			}
		}
	}
}

func (bp *BancoProcessadores) Atualiza_Shared_Exclusive(linha int, cache_id int) {
	var linha_existente *Linha

	for i := 0; i < constantes.QUANTIDADE_USUARIOS; i++ {
		cache := &bp.BP[i].Cachezinha

		if cache.id_processador != cache_id {
			linha_cache := cache.Procura_Cache(linha)

			if linha_cache != nil {
				if linha_existente != nil {
					println("mais de um")
					return
				}

				linha_existente = linha_cache
			}
		}
	}

	println(linha_existente)
	// panic("dsasddasdsa")

	linha_existente.Mesi = E
}

// func (bp *BancoProcessadores) Notifica_Caches(linha int, tag_nova MesiFlags, cache_id int) {
// 	for i := 0; i < constantes.QUANTIDADE_USUARIOS; i++ {
// 		fmt.Print("To no loop")
// 		cache := &bp.BP[i].Cachezinha

// 		if cache.id_processador != cache_id { // nao vai mudar a própria tag, mudar a dos outros.
// 			cache_index := cache.Procura_Cache(linha)
// 			if cache_index >= 0 { //tem nessa cache esse bloco, ver o que fazer com sua TAG_MESI.
// 				linha_analisada := cache.Linhas[cache_index]
// 				if linha_analisada.Mesi == E && tag_nova == S {
// 					fmt.Print("Achei! Mudando Tag!")
// 					//linha_analisada.Mesi = constantes.S //se eu fizer isso, nao muda. é uma cópia da referencia só.
// 					cache.Linhas[cache_index].Mesi = S
// 				} else if tag_nova == M {
// 					fmt.Print("@matheus_mds2!")
// 					cache.Linhas[cache_index].Mesi = I
// 				}

// 			}

// 		}

// 	}
// }
