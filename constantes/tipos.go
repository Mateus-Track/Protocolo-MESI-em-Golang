package constantes

const QUANTIDADE_USUARIOS = 4
const QUANTIDADE_LIVROS = 50
const QUANTIDADE_LINHAS_CACHE = 5
const TAMANHO_BLOCO = 5

const EXIT = false

var Cache_escolhida_int int

const (
	E = iota //0
	S
	M
	I
)

const ( //read hit, read miss...
	RH = iota //0
	RM
	WH
	WM
)
