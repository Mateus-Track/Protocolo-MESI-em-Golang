package componentes

import (
	"MESI/config"
	MESI "MESI/types"
	"errors"
	"fmt"
)

type BancoProcessadores struct {
	processadores [config.QUANTIDADE_PROCESSADORES]Processador
}

func InicializaBP() BancoProcessadores {
	bp := BancoProcessadores{
		processadores: [config.QUANTIDADE_PROCESSADORES]Processador{},
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

func (bp *BancoProcessadores) VerificarMESI(linha int) (bool, MESI.MesiFlags, *Linha) {
	for i := range bp.processadores {
		p := &bp.processadores[i]
		flag, linha_cache, err := p.Cachezinha.StatusCache(linha)

		if err == nil && flag != MESI.I {
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
				linha_cache.Mesi = MESI.S
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
				linha_cache.Mesi = MESI.I
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

	linha_existente.Mesi = MESI.E
}
