package componentes

import (
	"MESI/config"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//--Linha-----------------------------------------------------------------------

type Linha struct {
	Livros [config.LINHAS_CACHE]Livro
	Bloco  int //saber se o bloco foi puxado pra cache ou nao.
	Mesi   MesiFlags
}

func InicializaLinha() Linha {
	linha := Linha{
		Livros: [config.LINHAS_CACHE]Livro{},
		Bloco:  -1, // Valor inicial para o bloco
		Mesi:   I,  // Valor inicial para MESI, meti o loco aq pra n começar em algum.	}
	}

	for i := range linha.Livros {
		linha.Livros[i] = InicializaLivro()
	}
	return linha
}

func (l Linha) Print() {
	fmt.Println("Linha:")
	for i, livro := range l.Livros {
		fmt.Printf("  Livro %d: %s\n", i+1, livro.Nome)
		fmt.Printf("  Secao %d: %s\n", i+1, livro.Secao)
	}
	fmt.Printf("  Bloco: %d\n", l.Bloco)
	fmt.Printf("  MESI: %d\n", l.Mesi)
}

//--Cache-----------------------------------------------------------------------

type Cache struct { //pelo menos 5 posições.
	id_processador int
	Linhas         [config.LINHAS_CACHE]Linha //acredito que serão as mesmas linhas, só uma cópia da MP.
	Fila           []uint8
}

func InicializaCache(id int) Cache {
	cache := Cache{
		id_processador: id,
		Linhas:         [config.LINHAS_CACHE]Linha{},
		Fila:           []uint8{},
	}

	// Inicializando cada Linha dentro da Cache
	for i := range cache.Linhas {
		cache.Linhas[i] = InicializaLinha()
	}

	return cache
}

func (cache *Cache) ProcurarLinha(linha int) *Linha {
	bloco := linha / config.TAMANHO_BLOCO

	for i := 0; i < config.LINHAS_CACHE; i++ {
		if cache.Linhas[i].Bloco == bloco {
			return &cache.Linhas[i]
		}
	}

	return nil
}

func (cache *Cache) StatusCache(linha int) (MesiFlags, *Linha, error) {
	linha_cache := cache.ProcurarLinha(linha)

	if linha_cache == nil {
		return 0, nil, errors.New("CacheUndefined")
	}

	return linha_cache.Mesi, linha_cache, nil
}

func (cache *Cache) Print() { //sem a fila por enquanto.
	for i := 0; i < config.LINHAS_CACHE; i++ {
		fmt.Printf("Linha da Cache de número %d:\n\n", i+1)
		fmt.Printf("MESI da linha = %d\n", cache.Linhas[i].Mesi)

		for j := 0; j < config.TAMANHO_BLOCO; j++ {
			fmt.Printf("Livro armazenado:\n")
			fmt.Printf("%s", cache.Linhas[i].Livros[j].ToString())
		}
		fmt.Printf("\n")
	}
}

func (cache *Cache) CarregarLinha(livros [config.TAMANHO_BLOCO]Livro, bloco int, mp *Memoria, bp *BancoProcessadores) *Linha {
	var posicao uint8

	linha_existe := cache.ProcurarLinha(bloco * config.TAMANHO_BLOCO)

	if linha_existe != nil {
		linha_existe.Livros = livros
		return linha_existe
	}

	if cache.TemEspacoLivre() {
		disponiveis := cache.ValoresDisponiveis()
		numero_random := GerarAleatorio(len(disponiveis))
		posicao = disponiveis[numero_random]
	} else { //tirar da fila, dar append no final lá, retornar pra mp.
		posicao = cache.Fila[0]
		primeiraLinha := &cache.Linhas[posicao]
		linha_retornar_mp := (primeiraLinha.Bloco) * config.TAMANHO_BLOCO

		if primeiraLinha.Mesi == M {
			// mp.Transferir_Cache_MP(primeiraLinha, primeiraLinha.Bloco)
		} else if primeiraLinha.Mesi == S {
			bp.AtualizarSharedExclusive(linha_retornar_mp, cache.id_processador)
		}

		cache.Fila = cache.Fila[1:]
	}

	cache.Fila = append(cache.Fila, posicao)

	cache.Linhas[posicao].Bloco = bloco
	cache.Linhas[posicao].Livros = livros

	return &cache.Linhas[posicao]
}

func (cache *Cache) ReadMiss(linha int, mp *Memoria, bp *BancoProcessadores) {
	bloco := linha / config.TAMANHO_BLOCO

	var linha_escrita *Linha

	encontrado, mesi_bp, linha_bp := bp.VerificarMESI(linha)

	if !encontrado {
		linha_escrita = mp.Transferir_MP_Cache(cache, bp, bloco)
		linha_escrita.Mesi = E
		cache.Print()
		return
	}

	switch mesi_bp {
	case M:
		linha_escrita = cache.CarregarLinha(linha_bp.Livros, bloco, mp, bp)
		linha_escrita.Mesi = S

		linha_bp.Mesi = S
		mp.Transferir_Cache_MP(linha_bp, bloco)

	case E:
		linha_escrita = mp.Transferir_MP_Cache(cache, bp, bloco)
		bp.AtualizarShared(linha, cache.id_processador)
		linha_escrita.Mesi = S

	case S:
		linha_escrita = mp.Transferir_MP_Cache(cache, bp, bloco)
		linha_escrita.Mesi = S

	default:
		panic("Flag Inválida")
	}

	//cache.Printa_Cache()
	linha_escrita.Print()
}

func (cache *Cache) ReadHit(linha int) {
	linha_cache := cache.ProcurarLinha(linha)
	livro := linha_cache.Livros[linha%config.TAMANHO_BLOCO]

	fmt.Printf("%s", livro.ToString())
	fmt.Printf("MESI do bloco: %d\n", linha_cache.Mesi)
}

func (cache *Cache) WriteMiss(linha int, reserva Reserva, mp *Memoria, bp *BancoProcessadores) {
	bloco := linha / config.TAMANHO_BLOCO
	encontrado, mesi_bp, linha_bp := bp.VerificarMESI(linha)

	if !encontrado {
		linha_escrita := mp.Transferir_MP_Cache(cache, bp, bloco)
		linha_escrita.Mesi = E
	} else {
		switch mesi_bp {
		case M:
			linha_escrita := cache.CarregarLinha(linha_bp.Livros, bloco, mp, bp)
			linha_escrita.Mesi = S
			mp.Transferir_Cache_MP(linha_bp, bloco)

		case E:
			linha_escrita := mp.Transferir_MP_Cache(cache, bp, bloco)
			bp.AtualizarShared(linha, cache.id_processador)
			linha_escrita.Mesi = S

		case S:
			linha_escrita := mp.Transferir_MP_Cache(cache, bp, bloco)
			linha_escrita.Mesi = S

		default:
			panic("Flag Inválida")
		}
	}

	linha_cache := cache.ProcurarLinha(linha)
	livro := &linha_cache.Livros[linha%config.TAMANHO_BLOCO]

	if linha_cache.Mesi == E || linha_cache.Mesi == S {
		linha_cache.Mesi = M
	}

	bp.AtualizarInvalid(linha, cache.id_processador)

	livro.AdicionarReserva(reserva)
	cache.Print()

	fmt.Printf("Escrita realizada:\n")
	fmt.Printf("%s", livro.ToString())
	fmt.Printf("MESI do bloco: %d\n", linha_cache.Mesi)
}

func (cache *Cache) WriteHit(linha int, reserva Reserva, mp *Memoria, bp *BancoProcessadores) {
	linha_cache := cache.ProcurarLinha(linha)
	livro := &linha_cache.Livros[linha%config.TAMANHO_BLOCO]

	if linha_cache.Mesi == E || linha_cache.Mesi == S {
		linha_cache.Mesi = M
	}

	bp.AtualizarInvalid(linha, cache.id_processador)

	livro.AdicionarReserva(reserva)

	fmt.Printf("Escrita realizada:\n")
	fmt.Printf("%s", livro.ToString())
	fmt.Printf("MESI do bloco: %d\n", linha_cache.Mesi)
}

func (c *Cache) TemEspacoLivre() bool {
	return len(c.Fila) < config.LINHAS_CACHE
}

func (c *Cache) ValoresDisponiveis() []uint8 { // n entendi totalmente.
	todosValores := []uint8{0, 1, 2, 3, 4}

	presente := make(map[uint8]bool)
	for _, v := range c.Fila {
		presente[v] = true
	}

	// // Cria uma lista para armazenar valores disponíveis
	disponiveis := []uint8{}
	for _, v := range todosValores {
		if !presente[v] {
			disponiveis = append(disponiveis, v)
		}
	}

	return disponiveis
}

func GerarAleatorio(quantidade int) int16 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return int16(r.Intn(quantidade))
}
