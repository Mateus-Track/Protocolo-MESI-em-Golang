package componentes

import (
	"errors"
	"fmt"
)

const QUANTIDADE_USUARIOS = 4

type BancoProcessadores struct {
	processadores [4]Processador
}

func InicializaBP() BancoProcessadores {
	bp := BancoProcessadores{
		processadores: [4]Processador{},
	}

	// Inicializar cada Processador dentro do Banco
	for i := range bp.processadores {
		bp.processadores[i] = InicializaProcessador(i)
	}

	return bp
}

func (bp *BancoProcessadores) SelecionarProcessador(i int) (*Processador, error) {
	if i < 0 || i >= len(bp.processadores) {
		return nil, errors.New("UndefinedCPU")
	}

	return &bp.processadores[i], nil
}

func (bp *BancoProcessadores) VerificarMESI(linha int) (bool, MesiFlags, *Linha) {
	for i := range bp.processadores {
		p := &bp.processadores[i]
		flag, linha_cache, err := p.Cachezinha.StatusCache(linha)

		if err == nil && flag != I {
			return true, flag, linha_cache
		}
	}

	return false, 0, nil
}

func (bp *BancoProcessadores) AtualizarShared(linha int, cache_id int) {
	for i := 0; i < len(bp.processadores); i++ {
		cache := &bp.processadores[i].Cachezinha

		if cache.id_processador != cache_id {
			linha_cache := cache.ProcurarLinha(linha)
			if linha_cache != nil {
				linha_cache.Mesi = S
			}
		}
	}
}

func (bp *BancoProcessadores) AtualizarInvalid(linha int, cache_id int) {
	for i := 0; i < len(bp.processadores); i++ {
		cache := &bp.processadores[i].Cachezinha

		if cache.id_processador != cache_id {
			linha_cache := cache.ProcurarLinha(linha)
			if linha_cache != nil {
				linha_cache.Mesi = I
			}
		}
	}
}

func (bp *BancoProcessadores) AtualizarSharedExclusive(linha int, cache_id int) {
	var linha_existente *Linha

	for i := 0; i < len(bp.processadores); i++ {
		cache := &bp.processadores[i].Cachezinha

		if cache.id_processador != cache_id {
			linha_cache := cache.ProcurarLinha(linha)
			fmt.Print(linha_cache)
			if linha_cache != nil {
				fmt.Printf("NEM ENTREI AQ MANEW")
				if linha_existente != nil {
					println("mais de um")
					return
				}

				linha_existente = linha_cache
			}
		}
	}

	println(linha_existente)

	linha_existente.Mesi = E
}
