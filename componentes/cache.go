package componentes

import (
	"MESI/constantes"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const QUANTIDADE_LINHAS_CACHE = 5

type Linha struct {
	Livros [5]Livro
	Bloco  int //saber se o bloco foi puxado pra cache ou nao.
	Mesi   MesiFlags
}

func InicializaLinha() Linha {
	linha := Linha{
		Livros: [5]Livro{},
		Bloco:  -1, // Valor inicial para o bloco
		Mesi:   I,  // Valor inicial para MESI, meti o loco aq pra n começar em algum.	}
	}

	for i := range linha.Livros {
		linha.Livros[i] = InicializaLivro()
	}
	return linha
}

func (l Linha) PrintLinha() {
	fmt.Println("Linha:")
	for i, livro := range l.Livros {
		fmt.Printf("  Livro %d: %s\n", i+1, livro.Nome)
		fmt.Printf("  Secao %d: %s\n", i+1, livro.Secao)
	}
	fmt.Printf("  Bloco: %d\n", l.Bloco)
	fmt.Printf("  MESI: %d\n", l.Mesi)
}

type Cache struct { //pelo menos 5 posições.
	id_processador int
	Linhas         [QUANTIDADE_LINHAS_CACHE]Linha //acredito que serão as mesmas linhas, só uma cópia da MP.
	Fila           []uint8
}

func InicializaCache(id int) Cache {
	cache := Cache{
		id_processador: id,
		Linhas:         [5]Linha{},
		Fila:           []uint8{},
	}

	// Inicializando cada Linha dentro da Cache
	for i := range cache.Linhas {
		cache.Linhas[i] = InicializaLinha()
	}

	return cache
}

func (cache *Cache) Procura_Cache(linha int) *Linha {
	bloco := linha / 5

	for i := 0; i < QUANTIDADE_LINHAS_CACHE; i++ {
		if cache.Linhas[i].Bloco == bloco {
			return &cache.Linhas[i]
		}
	}

	return nil
}

func (cache *Cache) Status_Cache(linha int) (MesiFlags, *Linha, error) {
	linha_cache := cache.Procura_Cache(linha)

	if linha_cache == nil {
		return 0, nil, errors.New("CacheUndefined")
	}

	return linha_cache.Mesi, linha_cache, nil
}

func (cache *Cache) Printa_Cache() { //sem a fila por enquanto.
	for i := 0; i < QUANTIDADE_LINHAS_CACHE; i++ {
		fmt.Printf("Linha da Cache de número %d:\n\n", i+1)
		fmt.Printf("MESI da linha = %d\n", cache.Linhas[i].Mesi)
		for j := 0; j < constantes.TAMANHO_BLOCO; j++ {
			fmt.Printf("Livro armazenado:\n")
			fmt.Printf("%s", cache.Linhas[i].Livros[j].ToString())
			//cache.Linhas[i].PrintLinha()
		}
		fmt.Printf("\n")
	}
}

func (cache *Cache) Carregar_Linha(livros [5]Livro, bloco int, mp *MP, bp *BancoProcessadores) *Linha {
	var posicao uint8

	linha_existe := cache.Procura_Cache(bloco * 5)

	if linha_existe != nil {
		linha_existe.Livros = livros
		return linha_existe
	}

	if cache.TemEspacoLivre() {
		disponiveis := cache.ValoresDisponiveis()
		numero_random := Gerar_Aleatorio(len(disponiveis))
		posicao = disponiveis[numero_random]
	} else { //tirar da fila, dar append no final lá, retornar pra mp.
		posicao = cache.Fila[0]
		primeiraLinha := &cache.Linhas[posicao]

		if primeiraLinha.Mesi == M {
			Transferir_Cache_MP(mp, primeiraLinha, primeiraLinha.Bloco)
		} else if primeiraLinha.Mesi == S {
			bp.Atualiza_Shared_Exclusive(bloco*5, cache.id_processador)
		}

		cache.Fila = cache.Fila[1:]
	}

	cache.Fila = append(cache.Fila, posicao)

	cache.Linhas[posicao].Bloco = bloco
	cache.Linhas[posicao].Livros = livros

	return &cache.Linhas[posicao]
}

func (cache *Cache) Read_Miss(linha int, mp *MP, bp *BancoProcessadores) {
	bloco := linha / 5
	encontrado, mesi_bp, linha_bp := bp.Verificar_MESI(linha)

	if !encontrado {
		linha_escrita := Transferir_MP_Cache(mp, cache, bp, bloco)
		linha_escrita.Mesi = E
		cache.Printa_Cache()
		return
	}

	switch mesi_bp {
	case M:
		fmt.Print("Achei um M mesmo.")
		linha_escrita := cache.Carregar_Linha(linha_bp.Livros, bloco, mp, bp)
		linha_escrita.Mesi = S

		//linha_nova := cache.Carregar_Linha(linha_bp.Livros, bloco, mp, bp)
		linha_bp.PrintLinha()
		linha_bp.Mesi = S
		linha_bp.PrintLinha()
		bp.BP[0].Cachezinha.Printa_Cache() //testando especificamente com escrever na CACHE 0 , puxar em outra, aí ele mostrou como ta na 0
		Transferir_Cache_MP(mp, linha_bp, bloco)

	case E:
		linha_escrita := Transferir_MP_Cache(mp, cache, bp, bloco)
		bp.Atualiza_Shared(linha, cache.id_processador)
		linha_escrita.Mesi = S

	case S:
		linha_escrita := Transferir_MP_Cache(mp, cache, bp, bloco)
		linha_escrita.Mesi = S

	default:
		panic("Flag Inválida")
	}

	cache.Printa_Cache()
}

func (cache *Cache) Read_Hit(linha int) {
	linha_cache := cache.Procura_Cache(linha)
	livro := linha_cache.Livros[linha%5]

	fmt.Printf("%s", livro.ToString())
	fmt.Printf("MESI do bloco: %d\n", linha_cache.Mesi)
}

func (cache *Cache) Write_Miss(linha int, reserva Reserva, mp *MP, bp *BancoProcessadores) {
	bloco := linha / 5
	encontrado, mesi_bp, linha_bp := bp.Verificar_MESI(linha)

	if !encontrado {
		linha_escrita := Transferir_MP_Cache(mp, cache, bp, bloco)
		linha_escrita.Mesi = E
	} else {
		switch mesi_bp {
		case M:
			linha_escrita := cache.Carregar_Linha(linha_bp.Livros, bloco, mp, bp)
			linha_escrita.Mesi = S
			Transferir_Cache_MP(mp, linha_bp, bloco)

		case E:
			linha_escrita := Transferir_MP_Cache(mp, cache, bp, bloco)
			bp.Atualiza_Shared(linha, cache.id_processador)
			linha_escrita.Mesi = S

		case S:
			linha_escrita := Transferir_MP_Cache(mp, cache, bp, bloco)
			linha_escrita.Mesi = S

		default:
			panic("Flag Inválida")
		}
	}

	linha_cache := cache.Procura_Cache(linha)
	livro := &linha_cache.Livros[linha%5]

	if linha_cache.Mesi == E || linha_cache.Mesi == S {
		linha_cache.Mesi = M
	}

	bp.Atualiza_Invalid(linha, cache.id_processador)

	livro.AdicionarReserva(reserva)
	cache.Printa_Cache()

	fmt.Printf("Escrita realizada:\n")
	fmt.Printf("%s", livro.ToString())
	fmt.Printf("MESI do bloco: %d\n", linha_cache.Mesi)
}

func (cache *Cache) Write_Hit(linha int, reserva Reserva, mp *MP, bp *BancoProcessadores) {
	linha_cache := cache.Procura_Cache(linha)
	livro := &linha_cache.Livros[linha%5]

	if linha_cache.Mesi == E || linha_cache.Mesi == S {
		linha_cache.Mesi = M
	}

	bp.Atualiza_Invalid(linha, cache.id_processador)

	livro.AdicionarReserva(reserva)

	fmt.Printf("Escrita realizada:\n")
	fmt.Printf("%s", livro.ToString())
	fmt.Printf("MESI do bloco: %d\n", linha_cache.Mesi)
}

func (c *Cache) TemEspacoLivre() bool {
	return len(c.Fila) < QUANTIDADE_LINHAS_CACHE
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
